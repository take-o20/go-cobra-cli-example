package runner

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/take-o20/go-cli-example/pkg/client"
)

func NewRandomJokeRuuner() RandomJokeRunner {
	client := client.NewHttpClient()
	return RandomJokeRunner{
		client: client,
	}
}

type RandomJokeRunner struct {
	client client.HttpClient
}

func (runner *RandomJokeRunner) Run() error {
	url := "https://official-joke-api.appspot.com/jokes/random"
	res, err := runner.client.Get(url)
	if err != nil {
		return err
	}
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, []byte(res), "", "  "); err != nil {
		panic(err)
	}
	fmt.Println(dst.String())

	return nil
}
