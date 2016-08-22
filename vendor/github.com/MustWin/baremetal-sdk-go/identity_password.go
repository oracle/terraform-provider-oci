package baremetal

// UIPassword is returned for change or create password operations.
import (
	"net/http"
	"time"
)

type UpdateUIPasswordRequest struct {
	Password string `json:"password"`
}

//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#UIPassword
type UIPassword struct {
	ETaggedResource
	NewPassword  string    `json:"password"`
	UserID       string    `json:"userId"`
	TimeCreated  time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
	State        string    `json:"state"`
}

// CreateOrResetUIPassword - creates or resets password for user identified by
// userID. You MAY supply an idempotency token.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createOrResetUIPassword
func (c *Client) CreateOrResetUIPassword(password, userID string, opts ...Options) (resource *UIPassword, e error) {
	body := UpdateUIPasswordRequest{
		Password: password,
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceUsers,
		ids:     urlParts{userID, uiPassword},
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

// UpdateUserUIPassword - Changes the password of a user identified by userID. An
// ETAG MAY be passed as an option for optimistic concurrency control.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createOrResetUIPassword
func (c *Client) UpdateUserUIPassword(newPassword, userID string, opts ...Options) (resource *UIPassword, e error) {
	body := UpdateUIPasswordRequest{
		Password: newPassword,
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceUsers,
		options: opts,
		ids:     urlParts{userID, uiPassword},
	}

	var response *requestResponse
	if response, e = c.identityApi.request(http.MethodPut, reqOpts); e != nil {
		return
	}

	resource = &UIPassword{}
	e = response.unmarshal(resource)
	return
}
