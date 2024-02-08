package gowebengage

import (
	"context"
	"encoding/json"
	"fmt"
)

//////////////////////////////////////////////////////////////////////////////
//					create a new webengage user								//
//////////////////////////////////////////////////////////////////////////////

// CreateUser creates a new user in webengage.
func (w *webengage) CreateUser(ctx context.Context, apiKey, licenseCode,
	userID string, request UserRequest) error {

	body, err := json.Marshal(struct {
		UserID string `json:"userId"`
		UserRequest
	}{
		UserID:      userID,
		UserRequest: request,
	})
	if err != nil {
		return fmt.Errorf("unable to marshal request: %w", err)
	}

	url := fmt.Sprintf("/v1/accounts/%s/users", licenseCode)
	response, err := w.client.R().SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetAuthToken(apiKey).
		SetBody(body).
		Post(url)
	if err != nil {
		return fmt.Errorf("unable to send request: %w", err)
	}

	if response.StatusCode() != 200 && response.StatusCode() != 201 {
		var rsp WebengageResponseError
		if err := json.Unmarshal(response.Body(), &rsp); err != nil {
			return fmt.Errorf("unable to unmarshal response: %w", err)
		}
		return fmt.Errorf("unexpected status code: %d, message: %s",
			response.StatusCode(), rsp.Response.Message)
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////
//						update a new webengage user							//
//////////////////////////////////////////////////////////////////////////////

// UpdateUser creates a new user in webengage.
func (w *webengage) UpdateUser(ctx context.Context, apiKey, licenseCode,
	userID string, request UserRequest) error {

	body, err := json.Marshal(struct {
		UserID string `json:"userId"`
		UserRequest
	}{
		UserID:      userID,
		UserRequest: request,
	})
	if err != nil {
		return fmt.Errorf("unable to marshal request: %w", err)
	}

	url := fmt.Sprintf("/v1/accounts/%s/users", licenseCode)
	response, err := w.client.R().SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetAuthToken(apiKey).
		SetBody(body).
		Put(url)
	if err != nil {
		return fmt.Errorf("unable to send request: %w", err)
	}

	if response.StatusCode() != 200 {
		var rsp WebengageResponseError
		if err := json.Unmarshal(response.Body(), &rsp); err != nil {
			return fmt.Errorf("unable to unmarshal response: %w", err)
		}
		return fmt.Errorf("unexpected status code: %d, message: %s",
			response.StatusCode(), rsp.Response.Message)
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////
//						create a new webengage event						//
//////////////////////////////////////////////////////////////////////////////

// CreateEvent creates a new event in webengage.
func (w *webengage) CreateEvent(ctx context.Context, apiKey, licenseCode,
	userID string, eventName string, eventTime string, eventData map[string]interface{}) error {

	body, err := json.Marshal(struct {
		UserId    string                 `json:"userId"`
		EventName string                 `json:"eventName"`
		EventTime string                 `json:"eventTime"`
		EventData map[string]interface{} `json:"eventData"`
	}{
		UserId:    userID,
		EventName: eventName,
		EventTime: eventTime,
		EventData: eventData,
	})
	if err != nil {
		return fmt.Errorf("unable to marshal request: %w", err)
	}

	url := fmt.Sprintf("/v1/accounts/%s/events", licenseCode)
	response, err := w.client.R().SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetAuthToken(apiKey).
		SetBody(body).
		Post(url)
	if err != nil {
		return fmt.Errorf("unable to send request: %w", err)
	}

	if response.StatusCode() != 200 && response.StatusCode() != 201 {
		var rsp WebengageResponseError
		if err := json.Unmarshal(response.Body(), &rsp); err != nil {
			return fmt.Errorf("unable to unmarshal response: %w", err)
		}
		return fmt.Errorf("unexpected status code: %d, message: %s",
			response.StatusCode(), rsp.Response.Message)
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////
//			create a new webengage transactional campaign Messages			//
//////////////////////////////////////////////////////////////////////////////

// CreateTransactionalCampaignMessages creates a new transactional campaign in webengage.
func (w *webengage) CreateTransactionalCampaignMessages(ctx context.Context, apiKey, licenseCode,
	userID string, ttl int, token map[string]interface{}) (*TransactionalCampaignMessagesResponse, error) {

	var request TransactionalCampaignMessagesRequest
	request.SetData(ttl, userID, token)
	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal request: %w", err)
	}

	url := fmt.Sprintf("/v1/accounts/%s/events", licenseCode)
	response, err := w.client.R().SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetAuthToken(apiKey).
		SetBody(body).
		Post(url)
	if err != nil {
		return nil, fmt.Errorf("unable to send request: %w", err)
	}

	if response.StatusCode() != 200 && response.StatusCode() != 201 {
		var rsp WebengageResponseError
		if err := json.Unmarshal(response.Body(), &rsp); err != nil {
			return nil, fmt.Errorf("unable to unmarshal response: %w", err)
		}
		return nil, fmt.Errorf("unexpected status code: %d, message: %s",
			response.StatusCode(), rsp.Response.Message)
	}

	var rsp TransactionalCampaignMessagesResponse
	if err := json.Unmarshal(response.Body(), &rsp); err != nil {
		return nil, fmt.Errorf("request succeded send but unable to unmarshal response: %w", err)
	}

	return &rsp, nil
}
