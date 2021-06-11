package kafka

import (
	"errors"
	"go-template/internal/modules/config"
	"go-template/internal/pkg/kafka"

	k "github.com/segmentio/kafka-go"
)

var km *KafkaModule

func init() {
	// config
	c := config.CM()

	// writer
	w := kafka.NewWriter(
		kafka.WiterAddrOption(c.Config().Kafka.Brokers),
		kafka.WiterBalanceOption(&k.LeastBytes{}),
	)

	// p
	km = &KafkaModule{
		p: kafka.NewProducer(kafka.ProducerWriterOption(w)),
	}
}

type KafkaModule struct {
	p *kafka.Producer
}

func New() *KafkaModule {
	if km == nil {
		panic(errors.New("kafka module is not inited"))
	}

	return km
}

func (km *KafkaModule) Producer() *kafka.Producer {
	return km.p
}
