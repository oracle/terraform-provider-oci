// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

// DrgAttachment describes a Drg attachment to a Vcn
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DrgAttachment/
type DrgAttachment struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
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
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	DrgAttachments []DrgAttachment
}

func (l *ListDrgAttachments) GetList() interface{} {
	return &l.DrgAttachments
}

// CreateDrgAttachment attaches a drg to a vcn.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DrgAttachment/CreateDrgAttachment
func (c *Client) CreateDrgAttachment(drgID, vcnID string, opts *CreateOptions) (res *DrgAttachment, e error) {
	required := struct {
		DrgID string `header:"-" json:"drgId" url:"-"`
		VcnID string `header:"-" json:"vcnId" url:"-"`
	}{
		DrgID: drgID,
		VcnID: vcnID,
	}

	details := &requestDetails{
		name:     resourceDrgAttachments,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	res = &DrgAttachment{}
	e = resp.unmarshal(res)
	return
}

// GetDrgAttachment gets information about the specified drg attachment
//
// Seehttps://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DrgAttachment/GetDrgAttachment
func (c *Client) GetDrgAttachment(id string) (res *DrgAttachment, e error) {
	details := &requestDetails{
		name: resourceDrgAttachments,
		ids:  urlParts{id},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &DrgAttachment{}
	e = resp.unmarshal(res)
	return
}

// UpdateDrgAttachment updates the display name for the specified DrgAttachment.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DrgAttachment/UpdateDrgAttachment
func (c *Client) UpdateDrgAttachment(id string, opts *IfMatchDisplayNameOptions) (drg *DrgAttachment, e error) {
	details := &requestDetails{
		name:     resourceDrgAttachments,
		ids:      urlParts{id},
		optional: opts,
	}
	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	drg = &DrgAttachment{}
	e = resp.unmarshal(drg)

	return
}

// DeleteDrgAttachment detaches a drg from its vcn
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DrgAttachment/DeleteDrgAttachment
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
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DrgAttachment/ListDrgAttachments
func (c *Client) ListDrgAttachments(compartmentID string, opts *ListDrgAttachmentsOptions) (res *ListDrgAttachments, e error) {
	details := &requestDetails{
		name:     resourceDrgAttachments,
		required: listOCIDRequirement{compartmentID},
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &ListDrgAttachments{}
	e = resp.unmarshal(res)
	return
}
