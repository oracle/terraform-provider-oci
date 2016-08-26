package baremetal

// APIKey is returned for operations that create or modify user API keys.
import (
	"net/http"
	"time"
)

type CreateAPIKeyRequest struct {
	Key string `json:"key"`
}

//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#ApiKey
type APIKey struct {
	KeyID        string    `json:"keyId"`
	KeyValue     string    `json:"keyValue"`
	Fingerprint  string    `json:"fingerprint"`
	UserID       string    `json:"userId"`
	TimeCreated  time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
	State        string    `json:"state"`
}

// ListAPIKeyResponses contains a list of API keys
type ListAPIKeyResponses struct {
	ResourceContainer
	Keys []APIKey
}

func (l *ListAPIKeyResponses) GetList() interface{} {
	return &l.Keys
}

// Deletes an API key belonging to a user.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#deleteApiKey
func (c *Client) DeleteAPIKey(userID, fingerprint string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceUsers,
		options: opts,
		ids:     urlParts{userID, apiKeys, fingerprint},
	}
	return c.identityApi.deleteRequest(reqOpts)
}

// ListAPIKeys returns information about a user's API keys.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listApiKeys
func (c *Client) ListAPIKeys(userID string) (response *ListAPIKeyResponses, e error) {
	reqOpts := &sdkRequestOptions{
		name: resourceUsers,
		ids:  urlParts{userID, apiKeys, "/"},
	}

	var getResp *requestResponse
	if getResp, e = c.identityApi.getRequest(reqOpts); e != nil {
		return
	}

	response = &ListAPIKeyResponses{}
	e = getResp.unmarshal(response)
	return
}

// UploadAPIKey - add an API signing key for user. The key must be an RSA public
// key in pem format.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#uploadApiKey
func (c *Client) UploadAPIKey(userID, key string, opts ...Options) (apiKey *APIKey, e error) {
	body := CreateAPIKeyRequest{
		Key: key,
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceUsers,
		options: opts,
		ids:     urlParts{userID, apiKeys, "/"},
	}

	var resp *requestResponse
	if resp, e = c.identityApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	apiKey = &APIKey{}
	e = resp.unmarshal(apiKey)
	return
}
