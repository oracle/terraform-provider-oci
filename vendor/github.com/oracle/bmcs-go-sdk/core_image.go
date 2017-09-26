// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

// Image describes a boot disk image for launching an instance
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Image/
type Image struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	BaseImageID            string `json:"baseImageId"`
	CompartmentID          string `json:"compartmentId"`
	CreateImageAllowed     bool   `json:"createImageAllowed"`
	DisplayName            string `json:"displayName"`
	ID                     string `json:"id"`
	State                  string `json:"lifecycleState"`
	OperatingSystem        string `json:"operatingSystem"`
	OperatingSystemVersion string `json:"operatingSystemVersion"`
	TimeCreated            Time   `json:"timeCreated"`
}

// ListImages contains a list of images
//
type ListImages struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Images []Image
}

func (l *ListImages) GetList() interface{} {
	return &l.Images
}

// CreateImage is used to create an image
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Image/CreateImage
func (c *Client) CreateImage(compartmentID, instanceID string, opts *CreateOptions) (res *Image, e error) {
	required := struct {
		ocidRequirement
		InstanceID string `header:"-" json:"instanceId" url:"-"`
	}{
		InstanceID: instanceID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceImages,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	res = &Image{}
	e = resp.unmarshal(res)
	return
}

// GetImage retrieves information about an image
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Image/GetImage
func (c *Client) GetImage(id string) (res *Image, e error) {
	details := &requestDetails{
		name: resourceImages,
		ids:  urlParts{id},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &Image{}
	e = resp.unmarshal(res)
	return
}

// UpdateImage updates an images display name
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Image/UpdateImage
func (c *Client) UpdateImage(id string, opts *UpdateOptions) (res *Image, e error) {
	details := &requestDetails{
		name:     resourceImages,
		ids:      urlParts{id},
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &Image{}
	e = resp.unmarshal(res)
	return
}

// DeleteImage removes an image
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Image/DeleteImage
func (c *Client) DeleteImage(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		name:     resourceImages,
		ids:      urlParts{id},
		optional: opts,
	}
	return c.coreApi.deleteRequest(details)
}

// ListImages returns a list of images
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Image/ListImages
func (c *Client) ListImages(compartmentID string, opts *ListImagesOptions) (res *ListImages, e error) {
	details := &requestDetails{
		name:     resourceImages,
		required: listOCIDRequirement{compartmentID},
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &ListImages{}
	e = resp.unmarshal(res)
	return
}
