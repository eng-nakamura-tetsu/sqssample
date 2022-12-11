package main

import (
	adaptor "sqssample/adapter"
)

func main() {
	// 省略
}

const queueName = "prepared"

func SendMessageToSQS(s *adaptor.Adapter, message string) error {
	err := s.Enqueue(queueName, message)
	if err != nil {
		return err
	}
	return nil
}

func ReceiveMessageFromSQS(s *adaptor.Adapter) (string, error) {
	messages, err := s.Dequeue(queueName)
	if err != nil {
		return "", err
	}

	if len(messages) == 0 {
		return "", nil
	}

	message := *messages[0].Body
	return message, nil
}
