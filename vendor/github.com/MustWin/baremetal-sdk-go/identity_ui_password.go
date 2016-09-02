package baremetal

import "net/http"

// UIPassword represents a user's temporary password.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/UIPassword/
type UIPassword struct {
	ETaggedResource
	InactiveStatus uint16 `json:"inactiveStatus"`
	State          string `json:"lifecycleState"`
	Password       string `json:"password"`
	TimeCreated    Time   `json:"timeCreated"`
	UserID         string `json:"userId"`
}

// CreateOrResetUIPassword creates or resets password for the user with userID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/UIPassword/CreateOrResetUIPassword
func (c *Client) CreateOrResetUIPassword(userID string, opts ...Options) (resource *UIPassword, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceUsers,
		ids:     urlParts{userID},
		options: opts,
	}

	var response *requestResponse
	if response, e = c.identityApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	resource = &UIPassword{}
	e = response.unmarshal(resource)
	return
}
