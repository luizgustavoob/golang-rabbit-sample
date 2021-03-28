package domain

type Messaging interface {
	Publish(queueName string, message string) error
}
