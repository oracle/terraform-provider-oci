package baremetal

import (
	"net/http"
)

type PrivateIP struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	AvailabilityDomain string `json:"availabilityDomain"`
	CompartmentID      string `json:"compartmentId"`
	DisplayName        string `json:"displayName"`
	HostnameLabel      string `json:"hostnameLabel"`
	ID                 string `json:"id"`
	IPAddress          string `json:"ipAddress"`
	IsPrimary          bool   `json:"isPrimary"`
	SubnetID           string `json:"subnetId"`
	TimeCreated        Time   `json:"timeCreated"`
	VnicID             string `json:"vnicId"`
}

type ListPrivateIPs struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	PrivateIPs []PrivateIP
}

func (l *ListPrivateIPs) GetList() interface{} {
	return &l.PrivateIPs
}

func (c *Client) CreatePrivateIP(vnicID string, opts *CreatePrivateIPOptions) (privateIP *PrivateIP, e error) {
	required := struct {
		VnicId string `header:"-" json:"vnicId" url:"-"`
	}{
		VnicId: vnicID,
	}

	details := &requestDetails{
		name:     resourcePrivateIPs,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	privateIP = &PrivateIP{}
	e = resp.unmarshal(privateIP)
	return
}

func (c *Client) GetPrivateIP(id string) (privateIP *PrivateIP, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourcePrivateIPs,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	privateIP = &PrivateIP{}
	e = resp.unmarshal(privateIP)
	return
}

func (c *Client) UpdatePrivateIP(id string, opts *UpdatePrivateIPOptions) (privateIP *PrivateIP, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourcePrivateIPs,
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	privateIP = &PrivateIP{}
	e = resp.unmarshal(privateIP)
	return
}

func (c *Client) DeletePrivateIP(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourcePrivateIPs,
		optional: opts,
	}
	return c.coreApi.deleteRequest(details)
}

func (c *Client) ListPrivateIPs(opts *ListPrivateIPsOptions) (privateIPs *ListPrivateIPs, e error) {
	details := &requestDetails{
		name:     resourcePrivateIPs,
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	privateIPs = &ListPrivateIPs{}
	e = resp.unmarshal(privateIPs)
	return
}
