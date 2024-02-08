package gowebengage

type WebengageResponseError struct {
	Response struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"response"`
}

//////////////////////////////////////////////////////////////////////////////
//								Campaign									//
//////////////////////////////////////////////////////////////////////////////

type TransactionalCampaignMessagesResponse struct {
	Response struct {
		Data struct {
			TxnID        string `json:"txnId"`
			ExperimentID string `json:"experimentId"`
			UserID       string `json:"userId"`
			TTL          int    `json:"ttl"`
		} `json:"data"`
	} `json:"response"`
}

//////////////////////////////////////////////////////////////////////////////
//								Survey										//
//////////////////////////////////////////////////////////////////////////////

type SurveyResponse struct {
	Response struct {
		Data    map[string]interface{} `json:"data"`
		Message string                 `json:"message"`
		Status  string                 `json:"status"`
	} `json:"response"`
}
