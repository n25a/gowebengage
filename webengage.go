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
	userID string, request CreatUserRequest) error {

	body, err := json.Marshal(struct {
		UserID string `json:"userId"`
		CreatUserRequest
	}{
		UserID:           userID,
		CreatUserRequest: request,
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

	if response.StatusCode() != 200 {
		response.Body()
		var rsp WebengageResponseError
		if err := json.Unmarshal(response.Body(), &rsp); err != nil {
			return fmt.Errorf("unable to unmarshal response: %w", err)
		}
		return fmt.Errorf("unexpected status code: %d, message: %s",
			response.StatusCode(), rsp.Response.Message)
	}

	return nil
}
