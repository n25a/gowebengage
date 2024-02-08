package gowebengage

type CreatUserRequest struct {
	FirstName  string                 `json:"firstName"`
	LastName   string                 `json:"lastName"`
	BirthDate  string                 `json:"birthDate"`
	Gender     string                 `json:"gender"`
	Email      string                 `json:"email"`
	Phone      string                 `json:"phone"`
	Company    string                 `json:"company"`
	Attributes map[string]interface{} `json:"attributes"`
}
