// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "time"

// VnicAttachment Vnic information for a particular instance
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VnicAttachment/
type VnicAttachment struct {
	AvailabilityDomain string    `json:"availabilityDomain"`
	CompartmentID      string    `json:"compartmentId"`
	DisplayName        string    `json:"displayName"`
	ID                 string    `json:"id"`
	InstanceID         string    `json:"instanceId"`
	State              string    `json:"lifecycleState"`
	SubnetID           string    `json:"subnetId"`
	TimeCreated        time.Time `json:"TimeCreated"`
	VnicID             string    `json:"vnicId"`
}

// ListVnicAttachments list of VnicAttachments as well as optional OPCNextPage which
// can be used to pass as the Page field of CoreOptions in subsequent List calls.
// In conjunction with Limit is used in paginating results.
// OPCRequestID is used to identify the request for support issues.
type ListVnicAttachments struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Attachments []VnicAttachment
}

func (l *ListVnicAttachments) GetList() interface{} {
	return &l.Attachments
}

// ListVnicAttachments returns a list of VnicAttachments with matching compartmentID
// and optionally instanceId, vnicId, and/or availabilityDomain. Optional parameters
// are assigned to the optional CoreOptions argument.  Page and Limit can also
// be supplied to support pagination.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VnicAttachment/ListVnicAttachments
func (c *Client) ListVnicAttachments(compartmentID string, opts *ListVnicAttachmentsOptions) (res *ListVnicAttachments, e error) {
	details := &requestDetails{
		name:     resourceVnicAttachments,
		optional: opts,
		required: ocidRequirement{compartmentID},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &ListVnicAttachments{}
	e = resp.unmarshal(res)
	return
}
