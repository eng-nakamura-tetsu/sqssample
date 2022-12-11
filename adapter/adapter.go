package adaptor

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Adapter struct {
	sqs *sqs.SQS
}

func NewAdapter() (*Adapter, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("ap-northeast-1"),
		Endpoint: aws.String("http://localhost:4566"),
	})
	if err != nil {
		return nil, err
	}
	return &Adapter{
		sqs: sqs.New(sess),
	}, nil
}

func (s *Adapter) Enqueue(queueName string, message string) error {
	i := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueName),
		MessageBody: aws.String(message),
	}
	_, err := s.sqs.SendMessage(i)
	if err != nil {
		return err
	}
	return nil
}

func (s *Adapter) Dequeue(queueName string) ([]*sqs.Message, error) {
	i := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(queueName),
	}
	o, err := s.sqs.ReceiveMessage(i)
	if err != nil {
		return nil, err
	}
	return o.Messages, nil
}
