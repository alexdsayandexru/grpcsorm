package kafka

import (
	"crypto/tls"
	"github.com/IBM/sarama"
	"testing"
)

func Test_CloceProducer(t *testing.T) {
	producer, _ := GetProducer([]string{""}, &sarama.Config{})
	CloseProducer(producer)
}

func Test_Config(t *testing.T) {
	if NewConfig(&tls.Config{}, "", "") == nil {
		t.Error()
	}
	if NewDefaultConfig() == nil {
		t.Error()
	}
}

func Test_EmptyProducer(t *testing.T) {
	if ok, err := NewEmptyProducer().Send(nil); !ok {
		t.Error(err)
	}
}

func Test_DefaultProducerSarama(t *testing.T) {
	_, _ = NewDefaultProducerSarama("haproxy-kafka-dev.gid.team:39091", "").Send(nil)
	_, _ = NewDefaultProducerSarama("localhost:9092", "dev.idp.sorm.user-migrate.0").Send(nil)
}

func Test_ProducerSaramaSingleton(t *testing.T) {
	p := NewProducerSaramaSingleton("", "", &tls.Config{}, "", "")
	_, _ = p.Send(nil)
}

func Test_XDGSCRAMClient(t *testing.T) {
	client := &XDGSCRAMClient{HashGeneratorFcn: SHA512}
	if err := client.Begin("", "", ""); err != nil {
		t.Error(err)
	}
	if _, err := client.Step(""); err != nil {
		t.Error(err)
	}
	if ok := client.Done(); ok {
		t.Error()
	}
}
