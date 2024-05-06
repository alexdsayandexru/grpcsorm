package kafka

type IProducer interface {
	Send(data []byte) (bool, error)
}
