package chat_queue

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/utils"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	kafkaConnection *kafka.Conn
	kafkaReader     *kafka.Reader
}

func NewKafkaConsumer(conn *kafka.Conn, reader *kafka.Reader) *Consumer {
	return &Consumer{
		kafkaConnection: conn,
		kafkaReader:     reader,
	}
}

func (c *Consumer) ReadChatsBatchBad(ctx context.Context) ([]*chat.Chat, error) { // DO NOT USE, BAD AND RETURNS EMPTY ARRAY

	logger := utils.LoggerFromCtxOrCreate(ctx)
	batch := c.kafkaConnection.ReadBatch(10e3, 1e6) // fetch 1KB min, 1MB max

	var err error
	defer func() {
		errClose := batch.Close()
		if errClose != nil {
			if err == nil {
				err = errClose
			} else {
				logger.Warn().Msg("close batch")
			}
		}
	}()

	b := make([]byte, 10e3) // 1KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
	}
	return []*chat.Chat{}, nil
}

func (c *Consumer) ReadChatsBatch(ctx context.Context, batchSize int) ([]*chat.Chat, error) {
	logger := utils.LoggerFromCtxOrCreate(ctx)
	batch := c.kafkaConnection.ReadBatch(10e3, 1e6) // fetch 1KB min, 1MB max

	var err error
	defer func() {
		errClose := batch.Close()
		if errClose != nil {
			if err == nil {
				err = errClose
			} else {
				logger.Warn().Msg("close batch")
			}
		}
	}()

	var chats []*chat.Chat
	for {
		if len(chats) >= batchSize {
			break
		}

		m, err := c.kafkaReader.ReadMessage(ctx) // todo move to fetch + commit logic, this is not reliable
		if err != nil {
			break
		}

		var ch *chat.Chat
		err = json.NewDecoder(bytes.NewReader(m.Value)).Decode(ch)
		if err != nil {
			return nil, errors.Wrap(err, "extract json from message")
		}
		chats = append(chats, ch)
	}
	return chats, nil
}
