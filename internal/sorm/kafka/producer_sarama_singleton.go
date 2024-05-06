package kafka

import (
	"crypto/tls"
	"github.com/IBM/sarama"
	"strings"
	"sync"
)

type ProducerSaramaSingleton struct {
	address string
	topic   string
	config  *sarama.Config
}

func NewProducerSaramaSingleton(address string, topic string, tls *tls.Config, saslUserName string, saslPassword string) IProducer {
	config := NewConfig(tls, saslUserName, saslPassword)
	return &ProducerSaramaSingleton{
		address: address,
		topic:   topic,
		config:  config,
	}
}

func (p *ProducerSaramaSingleton) Send(data []byte) (bool, error) {
	producer, err := GetProducer(strings.Split(p.address, ","), p.config)
	if err != nil {
		return false, err
	}

	msg := sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(data),
	}
	if _, _, err = producer.SendMessage(&msg); err != nil {
		CloseProducer(producer)
	}
	return err == nil, err
}

var gProducer sarama.SyncProducer
var mu sync.Mutex

func GetProducer(addrs []string, config *sarama.Config) (sarama.SyncProducer, error) {
	var err error
	mu.Lock()
	defer mu.Unlock()

	if gProducer == nil {
		gProducer, err = sarama.NewSyncProducer(addrs, config)
		if err != nil {
			return nil, err
		}
	}
	return gProducer, nil
}

func CloseProducer(producer sarama.SyncProducer) {
	defer func() {
		recover()
	}()

	_ = producer.Close()
	mu.Lock()
	gProducer = nil
	mu.Unlock()
}
