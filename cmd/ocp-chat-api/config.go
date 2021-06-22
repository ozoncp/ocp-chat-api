package main

import "time"

const defaultTransportProtocol = "tcp"

const (
	defaultSQLMultiStatements  = true
	defaultSQLTimeout          = 5 * time.Second
	defaultSQLMaxAllowedPacket = 4096
)

const (
	defaultKafkaTopic = "chats"
	defaultKafkaHost  = "kafka-1"
	defaultKafkaPort  = "9092"
)

const (
	defaultStorageFlusherPeriod   = 10 * time.Second
	defaultStorageFlusherCapacity = 1000
)

const (
	defaultSQLHost     = "postgres"
	defaultSQLPort     = 5432
	defaultSQLLogin    = "postgres"
	defaultSQLPassword = "example"
	defaultSQLDBName   = "my_my_db"
)

const (
	defaultMigrationsURL      = "file://./migrations"
	defaultMigrationRun       = false
	defaultMigrationDBVersion = 1
)

type DatabaseConfig struct {
	Host         string `envconfig:"SQL_HOST" required:"true"`
	Port         uint16 `envconfig:"SQL_PORT" `
	User         string `envconfig:"SQL_LOGIN" required:"true"`
	Password     string `envconfig:"SQL_PASSWORD" required:"true"`
	DatabaseName string `envconfig:"SQL_DB_NAME" required:"true"`

	Timeout          time.Duration // needs definition!
	MaxAllowedPacket int           // needs definition!
	MultiStatements  bool          // needs definition!

	// MigrationsURL is directory containing migration scripts.
	MigrationsURL string `envconfig:"MIGRATION_FILES_LOCATION"`
	// MigrationRun is a flag: if true, we should run migration to particular version.
	MigrationRun bool `envconfig:"MIGRATION_RUN"`
	// MigrationDBVersion is version of DB that we should migrate to.
	MigrationDBVersion uint `envconfig:"MIGRATION_DB_VERSION"`
}

type KafkaConfig struct {
	Host  string `envconfig:"KAFKA_HOST"    required:"true"`
	Port  string `envconfig:"KAFKA_PORT" `
	Topic string `envconfig:"KAFKA_TOPIC"   required:"true"`

	// MigrationsURL is directory containing migration scripts.
	MigrationsURL string `envconfig:"MIGRATION_FILES_LOCATION"`
	// MigrationRun is a flag: if true, we should run migration to particular version.
	MigrationRun bool `envconfig:"MIGRATION_RUN"`
	// MigrationDBVersion is version of DB that we should migrate to.
	MigrationDBVersion uint `envconfig:"MIGRATION_DB_VERSION"`
}

type Config struct {
	HTTPAddr               string `envconfig:"HTTP_ADDR"`
	GRPCAddr               string `envconfig:"GRPC_ADDR"`
	DatabaseCfg            DatabaseConfig
	KafkaCfg               KafkaConfig
	StorageFlusherPeriod   time.Duration `envconfig:"STORAGE_FLUSHER_PERIOD"`
	StorageFlusherCapacity int64         `envconfig:"STORAGE_FLUSHER_CAPACITY"`
}

func NewDefaultConfig() *Config {
	return &Config{
		HTTPAddr: ":8080",
		GRPCAddr: ":5300",
		DatabaseCfg: DatabaseConfig{
			Host:         defaultSQLHost,
			Port:         defaultSQLPort,
			User:         defaultSQLLogin,
			Password:     defaultSQLPassword,
			DatabaseName: defaultSQLDBName,

			MultiStatements:  defaultSQLMultiStatements,
			Timeout:          defaultSQLTimeout,
			MaxAllowedPacket: defaultSQLMaxAllowedPacket,

			MigrationsURL:      defaultMigrationsURL,
			MigrationRun:       defaultMigrationRun,
			MigrationDBVersion: defaultMigrationDBVersion,
		},
		KafkaCfg: KafkaConfig{
			Host:  defaultKafkaHost,
			Port:  defaultKafkaPort,
			Topic: defaultKafkaTopic,
		},
		StorageFlusherPeriod:   defaultStorageFlusherPeriod,
		StorageFlusherCapacity: defaultStorageFlusherCapacity,
	}
}
