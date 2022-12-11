package main_test

import (
	"testing"

	main "sqssample"
	"sqssample/adaptor"

	"github.com/stretchr/testify/assert"
)

func TestEnqueueAndDeQueue(t *testing.T) {
	t.Run("enqueueしたメッセージをdequeueできメッセージを受け取れる", func(t *testing.T) {
		s, err := adaptor.NewAdapter()
		assert.NoError(t, err)

		// sqsにenqueue
		err = main.SendMessageToSQS(s, "送りたいメッセージ ")
		assert.NoError(t, err)

		var message string
		// sqsからdequeue
		message, err = main.ReceiveMessageFromSQS(s)
		assert.NoError(t, err)

		assert.Equal(t, message, "送りたいメッセージ ")
	})
}
