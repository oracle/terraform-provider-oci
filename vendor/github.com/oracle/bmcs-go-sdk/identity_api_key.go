// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

// APIKey is returned for operations that create or modify user API keys.
import (
	"time"
)

// APIKey is a PEM-format RSA credential for securing requests to the Oracle Bare Metal Cloud Services REST API.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/ApiKey/
type APIKey struct {
	KeyID        string    `json:"keyId"`
	KeyValue     string    `json:"keyValue"`
	Fingerprint  string    `json:"fingerprint"`
	UserID       string    `json:"userId"`
	TimeCreated  time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
	State        string    `json:"lifecycleState"`
}

// ListAPIKeyResponses contains a list of API keys
type ListAPIKeyResponses struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Keys []APIKey
}

func (l *ListAPIKeyResponses) GetList() interface{} {
	return &l.Keys
}

// DeleteAPIKey deletes an API key belonging to a user.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/ApiKey/DeleteApiKey
func (c *Client) DeleteAPIKey(userID, fingerprint string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{userID, apiKeys, fingerprint},
		name:     resourceUsers,
		optional: opts,
	}

	return c.identityApi.deleteRequest(details)
}

// ListAPIKeys returns information about a user's API keys.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/ApiKey/ListApiKeys
func (c *Client) ListAPIKeys(userID string) (resp *ListAPIKeyResponses, e error) {
	details := &requestDetails{
		ids:  urlParts{userID, apiKeys, "/"},
		name: resourceUsers,
	}

	var getResp *response
	if getResp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	resp = &ListAPIKeyResponses{}
	e = getResp.unmarshal(resp)
	return
}

// UploadAPIKey - add an API signing key for user. The key must be an RSA public
// key in pem format.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/ApiKey/UploadApiKey
func (c *Client) UploadAPIKey(userID, key string, opts *RetryTokenOptions) (apiKey *APIKey, e error) {
	required := struct {
		Key string `header:"-" json:"key" url:"-"`
	}{
		Key: key,
	}

	details := &requestDetails{
		ids:      urlParts{userID, apiKeys, "/"},
		name:     resourceUsers,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.identityApi.postRequest(details); e != nil {
		return
	}

	apiKey = &APIKey{}
	e = resp.unmarshal(apiKey)
	return
}
