package base

type Kafka interface {
	SendMessage(message string)
}
