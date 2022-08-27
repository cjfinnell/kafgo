package internal

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func Ping(ctx context.Context, bootstrap string) error {
	fmt.Printf("Attempting connection to %s ...", bootstrap)

	conn, err := kafka.DialContext(ctx, "tcp", bootstrap)
	if err != nil {
		return err
	}

	fmt.Println(" OK")

	return conn.Close()
}

func PingSASL(ctx context.Context, bootstrap, username, password string) error {
	mechanism := plain.Mechanism{
		Username: username,
		Password: password,
	}

	dialer := kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	fmt.Printf("Attempting connection to %s ...", bootstrap)

	conn, err := dialer.DialContext(ctx, "tcp", bootstrap)
	if err != nil {
		return err
	}

	fmt.Println(" OK")

	return conn.Close()
}
