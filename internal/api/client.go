package api

type Client interface {
	Get() (*JokeResponse, error)
}