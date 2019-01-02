package messaging

type Producer interface {
	PublishMessages(topic string, msgs []string)
	Close()
}
