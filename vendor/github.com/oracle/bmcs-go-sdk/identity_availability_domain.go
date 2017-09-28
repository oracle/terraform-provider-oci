// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

// AvailablityDomain contains name and then tenancy ID that an
// availability domain belongs to.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/AvailabilityDomain/
type AvailabilityDomain struct {
	Name          string `json:"name"`
	CompartmentID string `json:"compartmentId"`
}

// ListAvailabilityDomains contains a list AvailabilityDomain
type ListAvailabilityDomains struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	AvailabilityDomains []AvailabilityDomain
}

func (l *ListAvailabilityDomains) GetList() interface{} {
	return &l.AvailabilityDomains
}

// ListAvailabilityDomains lists availability domains in a user's root tenancy.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/AvailabilityDomain/ListAvailabilityDomains
func (c *Client) ListAvailabilityDomains(compartmentID string) (ads *ListAvailabilityDomains, e error) {
	details := &requestDetails{
		name:     resourceAvailabilityDomains,
		required: listOCIDRequirement{CompartmentID: compartmentID},
	}

	var getResp *response
	if getResp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	ads = &ListAvailabilityDomains{}
	e = getResp.unmarshal(ads)
	return
}
