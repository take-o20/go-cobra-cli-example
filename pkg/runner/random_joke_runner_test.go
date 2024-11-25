package runner

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/take-o20/go-cli-example/pkg/client"
)

func TestRandomJokeRunner_Succeeded(t *testing.T) {
	mockResponse := `
{
  "type": "programming",
  "setup": "Knock-knock.",
  "punchline": "A race condition. Who is there?",
  "id": 362
}`
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := client.NewMockHttpClient(ctrl)
	mockClient.EXPECT().
		Get("https://official-joke-api.appspot.com/jokes/random").
		Return(mockResponse, nil)

	runner := RandomJokeRunner{
		client: mockClient,
	}
	err := runner.Run()
	assert.NoError(t, err)
}

func TestRandomJokeRunner_Failed(t *testing.T) {
	mockResponse := ""
	mockErr := errors.New("mock error")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := client.NewMockHttpClient(ctrl)
	mockClient.EXPECT().
		Get("https://official-joke-api.appspot.com/jokes/random").
		Return(mockResponse, mockErr)

	runner := RandomJokeRunner{
		client: mockClient,
	}
	err := runner.Run()
	assert.Error(t, err)
}
