package runner

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/take-o20/go-cli-example/pkg/client"
)

type mockHttpClient struct {
	mockResponse string
	mockError    error
}

func (m *mockHttpClient) Get(string) (string, error) {
	return m.mockResponse, m.mockError
}

func createMockHttpClinet(t *testing.T, mockResponse string, mockError error) client.HttpClient {
	t.Helper()
	return &mockHttpClient{
		mockResponse: mockResponse,
		mockError:    mockError,
	}
}

func TestRandomJokeRunner_Succeeded(t *testing.T) {
	testResponse := `
{
  "type": "programming",
  "setup": "Knock-knock.",
  "punchline": "A race condition. Who is there?",
  "id": 362
}`
	mockClient := createMockHttpClinet(t, testResponse, nil)
	runner := RandomJokeRunner{
		client: mockClient,
	}
	err := runner.Run()
	assert.NoError(t, err)
}

func TestRandomJokeRunner_Failed(t *testing.T) {
	testResponse := ""
	mockErr := errors.New("mock error")
	mockClient := createMockHttpClinet(t, testResponse, mockErr)
	runner := RandomJokeRunner{
		client: mockClient,
	}
	err := runner.Run()
	assert.Error(t, err)
}
