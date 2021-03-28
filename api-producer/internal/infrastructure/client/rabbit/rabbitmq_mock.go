package rabbit

type RabbitMQMock struct {
	PublishInvokedCount int
	PublishFn           func(*[]string, string, *string) error
	FakeQueue           []string
}

func (self *RabbitMQMock) Publish(queueName string, message string) (err error) {
	self.PublishInvokedCount++
	return self.PublishFn(&self.FakeQueue, queueName, &message)
}
