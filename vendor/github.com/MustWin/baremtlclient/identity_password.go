package baremtlsdk

// UIPassword is returned for change or create password operations.
import (
	"encoding/json"
	"net/http"
	"time"
)

type UpdateUIPasswordRequest struct {
	Password string `json:"password"`
}

//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#UIPassword
type UIPassword struct {
	NewPassword  string    `json:"password"`
	UserID       string    `json:"userId"`
	TimeCreated  time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
	State        string    `json:"state"`
	ETag         string    `json:"etag,omitempty"`
	OPCRequestID string    `json:"opc-request-id,omitempty"`
}

// CreateOrResetUIPassword - creates or resets password for user identified by
// userID. You MAY supply an idempotency token.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createOrResetUIPassword
func (c *Client) CreateOrResetUIPassword(password, userID string, opts ...Options) (newpassword *UIPassword, e error) {
	var headers http.Header
	if len(opts) > 0 && opts[0].OPCIdempotencyToken != "" {
		headers = http.Header{}
		headers.Set(headerOPCIdempotencyToken, opts[0].OPCIdempotencyToken)
	}

	url := buildIdentityURL(resourceUsers, nil, userID, uiPassword)
	request := UpdateUIPasswordRequest{
		Password: password,
	}

	var response *requestResponse
	if response, e = c.api.request(http.MethodPost, url, request, headers); e != nil {
		return
	}

	newpassword = &UIPassword{}
	if e = json.Unmarshal(response.body, newpassword); e != nil {
		return
	}

	if response.header != nil {
		newpassword.ETag = response.header.Get(headerIfMatch)
		newpassword.OPCRequestID = response.header.Get(headerOPCRequestID)
	}

	return
}

// UpdateUserUIPassword - Changes the password of a user identified by userID. An
// ETAG MAY be passed as an option for optimistic concurrency control.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createOrResetUIPassword
func (c *Client) UpdateUserUIPassword(newPassword, userID string, opts ...Options) (uipwd *UIPassword, e error) {
	var headers http.Header
	if len(opts) > 0 && opts[0].IfMatch != "" {
		headers = http.Header{}
		headers.Set(headerIfMatch, opts[0].IfMatch)
	}

	request := UpdateUIPasswordRequest{
		Password: newPassword,
	}

	url := buildIdentityURL(resourceUsers, nil, userID, uiPassword)

	var response *requestResponse
	if response, e = c.api.request(http.MethodPut, url, request, headers); e != nil {
		return
	}

	uipwd = &UIPassword{}
	if e = json.Unmarshal(response.body, uipwd); e != nil {
		return
	}

	if response.header != nil {
		uipwd.ETag = response.header.Get(headerIfMatch)
		uipwd.OPCRequestID = response.header.Get(headerOPCRequestID)
	}

	return
}
