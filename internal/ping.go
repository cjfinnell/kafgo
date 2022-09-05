package internal

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func Ping(ctx context.Context, dialer *kafka.Dialer, bootstrap string) error {
	fmt.Printf("Attempting connection to %s ...", bootstrap)

	conn, err := kafka.DialContext(ctx, "tcp", bootstrap)
	if err != nil {
		return err
	}

	fmt.Println(" OK")

	return conn.Close()
}
