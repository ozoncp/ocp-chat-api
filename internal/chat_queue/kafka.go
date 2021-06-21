package chat_queue

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/utils"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"os"
	"time"
)

type Consumer struct {
	kafkaConsumer sarama.Consumer
	kafkaReader   *kafka.Reader
	batchSize     int
	topic         string
}

func NewKafkaConsumer(consumer sarama.Consumer, reader *kafka.Reader, batchSize int, topic string) *Consumer {
	return &Consumer{
		kafkaConsumer: consumer,
		kafkaReader:   reader,
		batchSize:     batchSize,
		topic:         topic,
	}
}

func (c *Consumer) ReadChatsBatch(ctx context.Context, batchSize int) ([]*chat.Chat, error) {
	logger := utils.LoggerFromCtxOrCreate(ctx)

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
			return nil, errors.Wrapf(err, "extract json from message: %+v", m.Value)
		}
		chats = append(chats, ch)
		logger.Info().Msgf("chat added: %+v", ch)
	}
	return chats, nil
}

func (c *Consumer) Run(ctx context.Context) error {
	logger := utils.LoggerFromCtxOrCreate(ctx)
	//ticker := time.NewTicker(1 * time.Second)

	pl, err := c.kafkaConsumer.Partitions(c.topic)
	if err != nil {
		return errors.Wrap(err, "partition")
	}

	initialOffset := sarama.OffsetOldest

	var chats []*chat.Chat // todo channel -- queue

	for _, pt := range pl {
		pc, err := c.kafkaConsumer.ConsumePartition(c.topic, pt, initialOffset)
		if err != nil {
			logger.Err(err).Msgf("cannot consume partition. topic: %+v, pl %+v, pt %+v, offset %+v", c.topic, pl, pt, initialOffset)
			time.Sleep(1 * time.Second)
			continue
		}
		for m := range pc.Messages() {
			var ch *chat.Chat
			err = json.NewDecoder(bytes.NewReader(m.Value)).Decode(ch)
			if err != nil {
				logger.Warn().Err(err).Msgf("extract json from message: %+v", string(m.Value))
				continue
			}
			chats = append(chats, ch)
			logger.Info().Msgf("chat added: %+v", ch)
		}
	}
	// todo ctx.Done()

	return nil

	//
	//for {
	//	select {
	//	case <-ctx.Done():
	//		return errors.Wrap(ctx.Err(), "finish buffering saver by context done")
	//	case <- ticker.C:
	//		logger.Warn().Msg("no messages for whole second")
	//	default:
	//		gotChats, err := c.ReadChatsBatch(ctx, c.batchSize)
	//		if err != nil {
	//			return errors.Wrap(err, "flush by ticker")
	//		}
	//		for _, ch:= range gotChats {
	//			logger.Info().Msgf("chat: %+v", ch)
	//		}
	//	}
	//}
}

func messageRecieved(mess *sarama.ConsumerMessage) {
	_, _ = fmt.Fprintf(os.Stderr, "%v\n", string(mess.Value))
}
