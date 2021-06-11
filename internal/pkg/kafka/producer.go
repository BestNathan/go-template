package kafka

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

// producer
var DefaultTimeout = time.Second * 60
var ErrProducerHasNoWriter = errors.New("no kafka writer")

type ProducerOption func(*Producer)

func ProducerTimeoutOption(timeout int) ProducerOption {
	return func(p *Producer) {
		p.timeout = time.Millisecond * time.Duration(timeout)
	}
}

func ProducerWriterOption(writer *kafka.Writer) ProducerOption {
	return func(p *Producer) {
		p.writer = writer
	}
}

type Producer struct {
	writer  *kafka.Writer
	timeout time.Duration
}

func NewProducer(opts ...ProducerOption) *Producer {
	p := &Producer{
		timeout: DefaultTimeout,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p Producer) PublishMessage(ctx context.Context, msg *kafka.Message) error {
	if p.writer == nil {
		return ErrProducerHasNoWriter
	}
	return p.writer.WriteMessages(ctx, *msg)
}

func (p Producer) PublishMessages(ctx context.Context, msg []kafka.Message) error {
	if p.writer == nil {
		return ErrProducerHasNoWriter
	}
	return p.writer.WriteMessages(ctx, msg...)
}

func (p Producer) PublishWithContext(ctx context.Context, topic string, key, data []byte) error {
	msg := &kafka.Message{
		Topic: topic,
		Key:   key,
		Value: data,
	}
	return p.PublishMessage(ctx, msg)
}

func (p Producer) Publish(topic string, key, data []byte) error {
	if p.timeout != 0 {
		ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
		defer cancel()
		return p.PublishWithContext(ctx, topic, key, data)
	}
	return p.PublishWithContext(context.TODO(), topic, key, data)
}

func (p Producer) PublishJsonMessage(topic string, key []byte, m interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	if p.timeout != 0 {
		ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
		defer cancel()
		return p.PublishWithContext(ctx, topic, key, data)
	}
	return p.PublishWithContext(context.TODO(), topic, key, data)
}

func (p Producer) RandomKey() []byte {
	key := make([]byte, 4)
	n := rand.Int31()
	binary.BigEndian.PutUint32(key, uint32(n))
	return key
}
