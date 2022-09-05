package internal

import (
	"crypto/tls"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

type DialerOption func(*kafka.Dialer)

func WithSASLPlain(username, password string) DialerOption {
	return func(d *kafka.Dialer) {
		d.SASLMechanism = plain.Mechanism{
			Username: username,
			Password: password,
		}
	}
}

func NewDialer(opts ...DialerOption) *kafka.Dialer {
	d := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS:       &tls.Config{},
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}
