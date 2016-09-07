package baremetal

import "net/http"

// Image describes a boot disk image for launching an instance
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Image/
type Image struct {
	ETaggedResource
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
	ResourceContainer
	Images []Image
}

func (l *ListImages) GetList() interface{} {
	return &l.Images
}

// CreateImage is used to create an image
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Image/CreateImage
func (c *Client) CreateImage(compartmentID, instanceID string, opts ...Options) (res *Image, e error) {
	body := struct {
		CompartmentID string `json:"compartmentId"`
		DisplayName   string `json:"displayName,omitempty"`
		InstanceID    string `json:"instanceId"`
	}{
		CompartmentID: compartmentID,
		InstanceID:    instanceID,
	}
	if len(opts) > 0 {
		body.DisplayName = opts[0].DisplayName
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceImages,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	res = &Image{}
	e = response.unmarshal(res)
	return
}

// GetImage retrieves information about an image
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Image/GetImage
func (c *Client) GetImage(id string, opts ...Options) (res *Image, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceImages,
		options: opts,
		ids:     urlParts{id},
	}
	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	res = &Image{}
	e = resp.unmarshal(res)
	return
}

// UpdateImage updates an images display name
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Image/UpdateImage
func (c *Client) UpdateImage(id string, opts ...Options) (res *Image, e error) {
	body := struct {
		DisplayName string `json:"displayName,omitempty"`
	}{}
	if len(opts) > 0 {
		body.DisplayName = opts[0].DisplayName
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceImages,
		options: opts,
		ids:     urlParts{id},
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPut, reqOpts); e != nil {
		return
	}

	res = &Image{}
	e = response.unmarshal(res)
	return
}

// DeleteImage removes an image
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Image/DeleteImage
func (c *Client) DeleteImage(id string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceImages,
		options: opts,
		ids:     urlParts{id},
	}
	return c.coreApi.deleteRequest(reqOpts)
}

// ListImages returns a list of images
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Image/ListImages
func (c *Client) ListImages(compartmentID string, opts ...Options) (res *ListImages, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceImages,
		ocid:    compartmentID,
		options: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	res = &ListImages{}
	e = resp.unmarshal(res)
	return
}
