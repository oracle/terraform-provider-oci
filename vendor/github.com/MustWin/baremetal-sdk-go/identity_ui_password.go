package baremetal

import "net/http"

// UIPassword represents a user's temporary password.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/UIPassword/
type UIPassword struct {
	ETaggedResource
	InactiveStatus uint16 `json:"inactiveStatus"`
	Password       string `json:"password"`
	State          string `json:"lifecycleState"`
	TimeCreated    Time   `json:"timeCreated"`
	UserID         string `json:"userId"`
}

// CreateOrResetUIPassword creates or resets password for the user with userID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/UIPassword/CreateOrResetUIPassword
func (c *Client) CreateOrResetUIPassword(userID string, opts *RetryTokenOptions) (resource *UIPassword, e error) {
	details := &requestDetails{
		ids:      urlParts{userID},
		name:     resourceUsers,
		optional: opts,
	}

	var response *requestResponse
	if response, e = c.identityApi.request(http.MethodPost, details); e != nil {
		return
	}

	resource = &UIPassword{}
	e = response.unmarshal(resource)
	return
}
