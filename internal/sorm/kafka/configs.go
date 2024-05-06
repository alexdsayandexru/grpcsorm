package kafka

import (
	"crypto/tls"
	"github.com/IBM/sarama"
	"time"
)

func NewConfig(tls *tls.Config, saslUserName string, saslPassword string) *sarama.Config {
	config := sarama.NewConfig()
	config.Net.TLS.Enable = true
	config.Net.TLS.Config = tls
	config.Net.TLS.Config.InsecureSkipVerify = true
	config.Net.SASL.Enable = len(saslUserName) > 0
	config.Net.SASL.User = saslUserName
	config.Net.SASL.Password = saslPassword
	config.Net.SASL.Handshake = len(saslUserName) > 0
	config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
	config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Net.DialTimeout = 10 * time.Second
	config.Net.WriteTimeout = 10 * time.Second
	return config
}

func NewDefaultConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Net.TLS.Enable = false
	config.Net.SASL.Enable = false
	config.Net.SASL.Handshake = false
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	return config
}
