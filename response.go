package gowebengage

type WebengageResponseError struct {
	Response struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"response"`
}
