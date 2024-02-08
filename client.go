package gowebengage

import "github.com/go-resty/resty/v2"

type Webengage interface {
}

type webengage struct {
	client *resty.Client
}

func NewWebengage(address string) Webengage {
	client := resty.New().
		SetAuthScheme("Bearer").
		SetBaseURL(address)

	return &webengage{
		client: client,
	}
}
