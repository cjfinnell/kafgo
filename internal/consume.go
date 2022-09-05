package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context, dialer *kafka.Dialer, bootstrap, topic string) error {
	fmt.Printf("Attempting connection to %s ...", bootstrap)

	groupID := fmt.Sprintf("kafgo-%d", time.Now().Unix())

	readerConfig := kafka.ReaderConfig{
		Brokers: []string{bootstrap},
		GroupID: groupID,
		Topic:   topic,
		Dialer:  dialer,
	}
	err := readerConfig.Validate()
	if err != nil {
		return err
	}

	reader := kafka.NewReader(readerConfig)

	fmt.Println(" OK")

	defer func() {
		if err := reader.Close(); err != nil {
			fmt.Printf("Error! %s: %s", topic, err)
		}
	}()

	fmt.Printf("Consuming from %s ...\n", topic)

	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}

		fmt.Printf("%v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
