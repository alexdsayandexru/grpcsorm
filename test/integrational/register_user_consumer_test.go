package main

import (
	"crypto/tls"
	"github.com/IBM/sarama"
	"gitlab.gid.team/sso/auth/sorm/configs"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/kafka"
	"log"
	"os"
	"os/signal"
	"strings"
	"testing"
)

func Test_RegisterUserConsumer0(t *testing.T) {
	config := configs.GetConfig()
	RunConsumer(config, config.TopicUserRegistration, 0)
}

func Test_RegisterUserConsumer1(t *testing.T) {
	config := configs.GetConfig()
	RunConsumer(config, config.TopicUserRegistration, 1)
}

func Test_RegisterUserConsumer2(t *testing.T) {
	config := configs.GetConfig()
	RunConsumer(config, config.TopicUserRegistration, 2)
}

func RunConsumer(config *configs.Config, topic string, partition int32) {
	var consumer *ConsumerSarama

	if config.EnabledTls {
		consumer = NewConsumerSarama(config.KafkaBrokers, topic,
			configs.GetTLS(config.CaCert, config.EnabledTls),
			config.SaslUserName, config.SaslPassword)
	} else {
		consumer = NewDefaultConsumerSarama(config.KafkaBrokers, topic)
	}

	consumer.Run(partition)
}

type ConsumerSarama struct {
	address string
	topic   string
	config  *sarama.Config
}

func NewConsumerSarama(address string, topic string, tls *tls.Config, saslUserName string, saslPassword string) *ConsumerSarama {
	config := kafka.NewConfig(tls, saslUserName, saslPassword)
	return &ConsumerSarama{
		address: address,
		topic:   topic,
		config:  config,
	}
}

func NewDefaultConsumerSarama(address string, topic string) *ConsumerSarama {
	config := kafka.NewDefaultConfig()
	return &ConsumerSarama{
		address: address,
		topic:   topic,
		config:  config,
	}
}

func (p *ConsumerSarama) Run(partition int32) {
	brokers := strings.Split(p.address, ",")

	consumer, err := sarama.NewConsumer(brokers, p.config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	client, _ := sarama.NewClient(brokers, p.config)
	offset, _ := client.GetOffset(p.topic, 0, sarama.OffsetNewest)

	partitionConsumer, err := consumer.ConsumePartition(p.topic, partition, offset-1)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("[%d] %s", msg.Offset, msg.Value)
		case <-signals:
			break ConsumerLoop
		}
	}
}
