package client

type HttpClient interface {
	Get(string) (string, error)
}
