package baremtlsdk

// CreatePolicy creates a new policy.
import (
	"bytes"
	"encoding/json"
	"net/http"
)

type CreatePolicyRequest struct {
	CreateIdentityResourceRequest
	Statements []string `json:"statements"`
}

type UpdatePolicyRequest struct {
	UpdateIdentityResourceRequest
	Statements []string `json:"statements"`
}

// Policy returned by GetPolicy and other policy related methods.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#Policy
type Policy struct {
	IdentityResource
	Statements []string `json:"statements"`
}

// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createPolicy
func (c *Client) CreatePolicy(policyName, policyDescription string, statements []string, options ...Options) (policy *Policy, e error) {
	urlStr := buildIdentityURL(resourcePolicies, nil)
	var headers http.Header

	if len(options) > 0 {
		if options[0].OPCIdempotencyToken != "" {
			headers = http.Header{}
			headers.Set(headerOPCIdempotencyToken, options[0].OPCIdempotencyToken)
		}
	}

	var request CreatePolicyRequest
	request.CompartmentID = c.authInfo.tenancyOCID
	request.Name = policyName
	request.Description = policyDescription
	request.Statements = statements

	var resp *requestResponse
	if resp, e = c.api.request(http.MethodPost, urlStr, request, headers); e != nil {
		return
	}

	policy = &Policy{}
	e = json.Unmarshal(resp.body, policy)
	return

}

// DeletePolicy removes a policy identified by policyID. Optionally pass an
// etag for optmistic concurrency control.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#DeletePolicy
func (c *Client) DeletePolicy(policyID string, opts ...Options) (e error) {
	var headers http.Header
	if len(opts) > 0 {
		if opts[0].IfMatch != "" {
			headers = http.Header{}
			headers.Set(headerIfMatch, opts[0].IfMatch)
		}
	}

	url := buildIdentityURL(resourcePolicies, nil, policyID)

	return c.api.deleteRequest(url, headers)

}

// GetPolicy returns a policy identified by a policyID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#getPolicy
func (c *Client) GetPolicy(policyID string) (policy *Policy, e error) {

	url := buildIdentityURL(resourcePolicies, nil, policyID)

	var getResp *requestResponse
	if getResp, e = c.api.getRequest(url, nil); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	policy = &Policy{}
	e = decoder.Decode(policy)
	return

}

// UpdatePolicy can be called to modify the description and statements of an existing policy.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#updatePolicy
func (c *Client) UpdatePolicy(policyID, policyDescription string, policyStatements []string, opts ...Options) (policy *Policy, e error) {
	var headers http.Header
	if len(opts) > 0 {
		if opts[0].IfMatch != "" {
			headers = http.Header{}
			headers.Set(headerIfMatch, opts[0].IfMatch)
		}
	}

	var request UpdatePolicyRequest
	request.Description = policyDescription
	request.Statements = policyStatements

	var resp []byte
	if resp, e = c.updateIdentityResource(resourcePolicies, policyID, request, headers); e != nil {
		return
	}

	policy = &Policy{}
	e = json.Unmarshal(resp, policy)
	return
}
