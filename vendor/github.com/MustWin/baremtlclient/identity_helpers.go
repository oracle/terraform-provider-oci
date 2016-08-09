package baremtlsdk

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type CreateIdentityResourceRequest struct {
	CompartmentID string `json:"compartmentId"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}

type UpdateIdentityResourceRequest struct {
	Description string `json:"description"`
}

func (c *Client) createIdentityResource(resourceType resourceName, request CreateIdentityResourceRequest, headers http.Header) (resource *IdentityResource, e error) {
	urlStr := buildIdentityURL(resourceType, nil)

	var resp *requestResponse
	if resp, e = c.api.request(http.MethodPost, urlStr, request, headers); e != nil {
		return
	}

	resource = &IdentityResource{}
	e = json.Unmarshal(resp.body, resource)
	return

}

func (c *Client) getIdentityResource(resource resourceName, ids ...interface{}) (item *IdentityResource, e error) {
	url := buildIdentityURL(resource, nil, ids...)
	var getResp *requestResponse
	if getResp, e = c.api.getRequest(url, nil); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	item = &IdentityResource{}
	e = decoder.Decode(item)
	return

}

func (c *Client) listIdentityResources(resource resourceName, options ...ListOptions) (resp *ListResourceResponse, e error) {

	q := url.Values{}
	q.Set(queryCompartmentID, c.authInfo.tenancyOCID)
	if len(options) > 0 {
		q.Set(queryLimit, strconv.FormatUint(options[0].Limit, 10))
		q.Set(queryPage, options[0].Page)
	}

	resourceURL := buildIdentityURL(resource, q)

	var getResp *requestResponse
	if getResp, e = c.api.getRequest(resourceURL, nil); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	var items []IdentityResource
	if e = decoder.Decode(&items); e != nil {
		return
	}

	resp = &ListResourceResponse{
		Page:  getResp.header.Get(headerOPCNextPage),
		Items: items,
	}

	return
}

func getUpdateHeaders(options ...Options) http.Header {
	var headers http.Header
	if len(options) > 0 {
		if options[0].IfMatch != "" {
			headers := &http.Header{}
			headers.Set(headerIfMatch, options[0].IfMatch)
		}
	}
	return headers
}

func (c *Client) updateIdentityResource(resource resourceName, resourceID string, request interface{}, headers http.Header) (resp []byte, e error) {
	url := buildIdentityURL(resource, nil, resourceID)
	var r *requestResponse
	if r, e = c.api.request(http.MethodPut, url, request, headers); e != nil {
		return
	}
	resp = r.body
	return
}
