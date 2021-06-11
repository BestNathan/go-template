package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type WriterOption func(*Writer)

func WiterBalanceOption(b kafka.Balancer) WriterOption {
	return func(w *Writer) {
		w.Balancer = b
	}
}

func WiterRequiredAcksOption(ack kafka.RequiredAcks) WriterOption {
	return func(w *Writer) {
		w.RequiredAcks = ack
	}
}

func WiterCompressionOption(codec kafka.Compression) WriterOption {
	return func(w *Writer) {
		w.Compression = codec
	}
}

func WiterAddrOption(address ...string) WriterOption {
	return func(w *Writer) {
		w.Addr = kafka.TCP(address...)
	}
}

type Writer = kafka.Writer

func NewWriter(opts ...WriterOption) *Writer {
	w := &kafka.Writer{
		// Async:        true,
		BatchTimeout: 20 * time.Millisecond,
	}

	for _, opt := range opts {
		opt(w)
	}

	return w
}
