package kafka

import (
	"go-template/internal/pkg/time"

	"github.com/google/uuid"
)

type KafkaConfig struct {
	Brokers string `mapstructure:"brokers"`
}

type KafkaMessageOption func(*KafkaMessage)

func KafkaMessageTraceIdOption(id string) KafkaMessageOption {
	return func(km *KafkaMessage) {
		km.TraceId = id
	}
}

type KafkaMessage struct {
	Action    string         `json:"action"`
	Data      interface{}    `json:"data"`
	Timestamp time.Timestamp `json:"timestamp"`
	MsgId     string         `json:"msgId"`
	TraceId   string         `json:"traceId"`
}

func NewKafkaMessage(action string, d interface{}, opts ...KafkaMessageOption) *KafkaMessage {
	km := &KafkaMessage{
		Action:    action,
		Data:      d,
		Timestamp: time.New(),
		MsgId:     uuid.NewString(),
		TraceId:   uuid.NewString(),
	}

	for _, opt := range opts {
		opt(km)
	}

	return km
}

func (km *KafkaMessage) SetAction(action string) *KafkaMessage {
	km.Action = action
	return km
}

func (km *KafkaMessage) SetData(data interface{}) *KafkaMessage {
	km.Data = data
	return km
}

func (km *KafkaMessage) SetTimestamp(ts time.Timestamp) *KafkaMessage {
	km.Timestamp = ts
	return km
}

func (km *KafkaMessage) SetMsgId(m string) *KafkaMessage {
	km.MsgId = m
	return km
}

func (km *KafkaMessage) SetTraceId(t string) *KafkaMessage {
	km.TraceId = t
	return km
}
