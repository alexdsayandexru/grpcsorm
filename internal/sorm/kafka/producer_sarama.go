package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"strings"
	"sync"
)

type ProducerSarama struct {
	address  string
	topic    string
	config   *sarama.Config
	mu       sync.Mutex
	producer sarama.SyncProducer
}

func NewDefaultProducerSarama(address string, topic string) IProducer {
	config := NewDefaultConfig()
	return &ProducerSarama{
		address: address,
		topic:   topic,
		config:  config,
	}
}

func (p *ProducerSarama) Send(data []byte) (bool, error) {
	producer, err := sarama.NewSyncProducer(strings.Split(p.address, ","), p.config)
	if err != nil {
		return false, err
	}
	defer func() {
		if err := producer.Close(); err != nil {
			fmt.Println("Error call producer.Close():", err)
		}
	}()

	msg := sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = producer.SendMessage(&msg)

	return err == nil, err
}
