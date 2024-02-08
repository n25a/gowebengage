package gowebengage

//////////////////////////////////////////////////////////////////////////////
//								  User										//
//////////////////////////////////////////////////////////////////////////////

type UserRequest struct {
	UserID     string                 `json:"userId"`
	FirstName  string                 `json:"firstName"`
	LastName   string                 `json:"lastName"`
	BirthDate  string                 `json:"birthDate"`
	Gender     string                 `json:"gender"`
	Email      string                 `json:"email"`
	Phone      string                 `json:"phone"`
	Company    string                 `json:"company"`
	Attributes map[string]interface{} `json:"attributes"`
}

//////////////////////////////////////////////////////////////////////////////
//								Campaign									//
//////////////////////////////////////////////////////////////////////////////

type TransactionalCampaignMessagesRequest struct {
	TTL          int    `json:"ttl"`
	UserId       string `json:"userId"`
	OverrideData struct {
		Context struct {
			Token map[string]interface{} `json:"token"`
		} `json:"context"`
	} `json:"overrideData"`
}

func (t *TransactionalCampaignMessagesRequest) SetData(ttl int, userID string, token map[string]interface{}) {
	t.TTL = ttl
	t.UserId = userID
	t.OverrideData.Context.Token = token
}

//////////////////////////////////////////////////////////////////////////////
//								Events										//
//////////////////////////////////////////////////////////////////////////////

type EventRequest struct {
	UserID    string                 `json:"userId"`
	EventName string                 `json:"eventName"`
	EventTime string                 `json:"eventTime"`
	EventData map[string]interface{} `json:"eventData"`
}
