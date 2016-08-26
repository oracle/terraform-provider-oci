package baremetal

import "net/http"

// DrgAttachment describes a Drg attachment to a Vcn
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#DrgAttachment
type DrgAttachment struct {
	ETaggedResource
	CompartmentID string `json:"compartmentId"`
	DisplayName   string `json:"displayName"`
	DrgID         string `json:"drgId"`
	ID            string `json:"id"`
	State         string `json:"state"`
	TimeCreated   Time   `json:"timeCreated"`
	VcnID         string `json:"vcnId"`
}

// ListDrgAttachments contains a list of volume attachments
//
type ListDrgAttachments struct {
	ResourceContainer
	DrgAttachments []DrgAttachment
}

func (l *ListDrgAttachments) GetList() interface{} {
	return &l.DrgAttachments
}

// CreateDrgAttachment attaches a drg to a vcn.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#createDrgAttachment
func (c *Client) CreateDrgAttachment(drgID, vcnID string, opts ...Options) (res *DrgAttachment, e error) {
	body := struct {
		DisplayName string `json:"displayName,omitempty"`
		DrgID       string `json:"drgId"`
		VcnID       string `json:"vcnId"`
	}{
		DrgID: drgID,
		VcnID: vcnID,
	}

	if len(opts) > 0 {
		body.DisplayName = opts[0].DisplayName
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceDrgAttachments,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	res = &DrgAttachment{}
	e = response.unmarshal(res)
	return
}

// GetDrgAttachment gets information about the specified drg attachment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#getDrgAttachment
func (c *Client) GetDrgAttachment(id string, opts ...Options) (res *DrgAttachment, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceDrgAttachments,
		options: opts,
		ids:     urlParts{id},
	}
	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	res = &DrgAttachment{}
	e = resp.unmarshal(res)
	return
}

// DeleteDrgAttachment detaches a drg from its vcn
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#deleteDrgAttachment
func (c *Client) DeleteDrgAttachment(id string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceDrgAttachments,
		options: opts,
		ids:     urlParts{id},
	}
	return c.coreApi.deleteRequest(reqOpts)
}

// ListDrgAttachments gets a list of the drgs in the specified compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#listDrgAttachments
func (c *Client) ListDrgAttachments(compartmentID string, opts ...Options) (res *ListDrgAttachments, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceDrgAttachments,
		ocid:    compartmentID,
		options: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	res = &ListDrgAttachments{}
	e = resp.unmarshal(res)
	return
}
