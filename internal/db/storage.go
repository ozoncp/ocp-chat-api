package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx"
	//_ "github.com/mattes/migrate/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required for file://migrations dir
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type MigrateConf struct {
	// MigrationsURL is directory containing migration scripts.
	MigrationsURL string
	// MigrationRun is a flag: if true, we should run migration to particular version.
	MigrationRun bool
	// MigrationDBVersion is version of DB that we should migrate to.
	MigrationDBVersion uint
}

type DatabaseConf struct {
	Host     string
	Port     uint16
	User     string
	Password string
	DBName   string

	Timeout          time.Duration // needs definition!
	MaxAllowedPacket int           // needs definition!
	MultiStatements  bool          // needs definition!
}

func MigrateToVersion(ctx context.Context, psqlConf *pgx.ConnConfig, cfg *MigrateConf) error {
	logger := *zerolog.Ctx(ctx)

	connectionString := PostgresFormatDSN(psqlConf)
	// nolint:godox // ok fixmee
	// fixme попробовать найти способ работать всё-таки с одним и тем же соединением
	// он жестко требует драйвер с правами шире, чем sql.DB.Driver()
	p := &postgres.Postgres{}

	d, err := p.Open(connectionString)
	if err != nil {
		return errors.Wrap(err, "create migrate driver")
	}

	m, err := migrate.NewWithDatabaseInstance(
		cfg.MigrationsURL,
		"postgres",
		d,
	)
	defer func() {
		if err = d.Close(); err != nil {
			logger.Error().Err(err).Msg("close db after migration")
		}
	}()
	if err != nil {
		return errors.Wrap(err, "create migration")
	}

	if err := m.Migrate(cfg.MigrationDBVersion); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Info().Str("component", "migration").Msg("version is up-to-date, no changes")
			return nil
		}
		return errors.Wrapf(err, "migrate to version %d", cfg.MigrationDBVersion)
	}

	return nil
}

func InitAndCreateDB(ctx context.Context, dbCfg *DatabaseConf, migrateCfg *MigrateConf) (*sql.DB, error) {
	logger := *zerolog.Ctx(ctx)

	psqlConf := &pgx.ConnConfig{
		Host:     dbCfg.Host,
		Port:     dbCfg.Port,
		Database: dbCfg.DBName,
		User:     dbCfg.User,
		Password: dbCfg.Password,

		Logger:   nil,
		LogLevel: 0,
		Dial:     nil,
	}

	logger.Info().Msgf("connect to mysql with addr %s", psqlConf.Host)

	psqlDB, err := NewPostgresSQL(ctx, psqlConf)
	if err != nil {
		return nil, errors.Wrapf(err, "open db connection: %s", err)
	}

	// migrate db to appropriate version
	if migrateCfg.MigrationRun {
		logger.Info().Msgf("migration is required: to version %d", migrateCfg.MigrationDBVersion)
		if err := MigrateToVersion(ctx, psqlConf, migrateCfg); err != nil {
			return nil, errors.Wrapf(err, "migrate db to appropriate version: %d", migrateCfg.MigrationDBVersion)
		}
		logger.Info().Msgf("migration to version %d successful", migrateCfg.MigrationDBVersion)
	} else {
		logger.Info().Msg("migration is not required")
	}
	return psqlDB, nil
}

// PostgresFormatDSN  format: postgres://username:password@localhost:5432/database_name
func PostgresFormatDSN(c *pgx.ConnConfig) string {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Database)
	return psqlInfo
}

func NewPostgresSQL(ctx context.Context, psqlConfig *pgx.ConnConfig) (*sql.DB, error) {
	logger := *zerolog.Ctx(ctx)
	logger.Info().Str("component", "database").Str("stage", "create").Msgf("opening connection")

	sqlDB, err := sql.Open("postgres", PostgresFormatDSN(psqlConfig))
	if err != nil {
		return nil, errors.Wrap(err, "open database")
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, errors.Wrap(err, "ping database")
	}

	return sqlDB, nil
}
