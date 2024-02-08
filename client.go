package gowebengage

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// Webengage is the interface for the webengage client.
type Webengage interface {
	CreateUser(ctx context.Context, apiKey string, licenseCode string, request UserRequest) error
	CreateBulkUser(ctx context.Context, apiKey string, licenseCode string, request []UserRequest) error
	UpdateUser(ctx context.Context, apiKey, licenseCode, userID string, request UserRequest) error
	CreateEvent(ctx context.Context, apiKey, licenseCode, userID string, eventName string, eventTime string,
		eventData map[string]interface{}) error
	CreateBulkEvent(ctx context.Context, apiKey string, licenseCode string, request []EventRequest) error
	CreateTransactionalCampaignMessages(ctx context.Context, apiKey, licenseCode, userID string, ttl int,
		token map[string]interface{}) (*TransactionalCampaignMessagesResponse, error)
	GetSurvey(ctx context.Context, apiKey, surveyResponseID string) (*SurveyResponse, error)
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
