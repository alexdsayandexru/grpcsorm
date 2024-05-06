package kafka

type EmptyProducer struct{}

func (p *EmptyProducer) Send([]byte) (bool, error) {
	return true, nil
}

func NewEmptyProducer() IProducer {
	return &EmptyProducer{}
}
