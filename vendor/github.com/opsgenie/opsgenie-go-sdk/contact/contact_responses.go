package contact

// Create contact response structure
type CreateContactResponse struct {
	Id 	string 	`json:"id"`
	Status 	string 	`json:"status"`
	Code 	int 	`json:"code"`
}

// Update contact response structure
type UpdateContactResponse struct {
	Id 	string 	`json:"id"`
	Status 	string 	`json:"status"`
	Code 	int 	`json:"code"`
}

// Delete contact response structure
type DeleteContactResponse struct {
	Status 	string 	`json:"status"`
	Code 	int 	`json:"code"`
}

// Get contact response structure
type GetContactResponse struct {
	Id 		string 	`json:"id,omitempty"`
	Method 		string 	`json:"method,omitempty"`
	To 		string 	`json:"to,omitempty"`
	DisabledReason 	string 	`json:"disabledReason, omitempty"`
	Enabled 	bool 	`json:"enabled, omitempty"`
}
