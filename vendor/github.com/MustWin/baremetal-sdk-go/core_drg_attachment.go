package baremetal

import "net/http"

// DrgAttachment describes a Drg attachment to a Vcn
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DrgAttachment/
type DrgAttachment struct {
	ETaggedResource
	CompartmentID string `json:"compartmentId"`
	DisplayName   string `json:"displayName"`
	DrgID         string `json:"drgId"`
	ID            string `json:"id"`
	State         string `json:"lifecycleState"`
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
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DrgAttachment/CreateDrgAttachment
func (c *Client) CreateDrgAttachment(drgID, vcnID string, opts *CreateOptions) (res *DrgAttachment, e error) {
	required := struct {
		DrgID string `json:"drgId" url:"-"`
		VcnID string `json:"vcnId" url:"-"`
	}{
		DrgID: drgID,
		VcnID: vcnID,
	}

	details := &requestDetails{
		name:     resourceDrgAttachments,
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, details); e != nil {
		return
	}

	res = &DrgAttachment{}
	e = response.unmarshal(res)
	return
}

// GetDrgAttachment gets information about the specified drg attachment
//
// Seehttps://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DrgAttachment/GetDrgAttachment
func (c *Client) GetDrgAttachment(id string) (res *DrgAttachment, e error) {
	details := &requestDetails{
		name: resourceDrgAttachments,
		ids:  urlParts{id},
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &DrgAttachment{}
	e = resp.unmarshal(res)
	return
}

// DeleteDrgAttachment detaches a drg from its vcn
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DrgAttachment/DeleteDrgAttachment
func (c *Client) DeleteDrgAttachment(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		name:     resourceDrgAttachments,
		ids:      urlParts{id},
		optional: opts,
	}
	return c.coreApi.deleteRequest(details)
}

// ListDrgAttachments gets a list of the drgs in the specified compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DrgAttachment/ListDrgAttachments
func (c *Client) ListDrgAttachments(compartmentID string, opts *ListDrgAttachmentsOptions) (res *ListDrgAttachments, e error) {
	details := &requestDetails{
		name:     resourceDrgAttachments,
		required: listOCIDRequirement{compartmentID},
		optional: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &ListDrgAttachments{}
	e = resp.unmarshal(res)
	return
}
