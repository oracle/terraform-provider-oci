// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

type IdentityRegion struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	Key  string `json:"key"`
	Name string `json:"name"`
}

type ListRegions struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Regions []IdentityRegion
}

func (l *ListRegions) GetList() interface{} {
	return &l.Regions
}

// ListRegions returns a list of regions
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Region/ListRegions
func (c *Client) ListRegions() (resources *ListRegions, e error) {
	details := &requestDetails{
		name:     resourceRegions,
		required: listOCIDRequirement{c.authInfo.tenancyOCID},
	}

	var getResp *response
	if getResp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	resources = &ListRegions{}
	e = getResp.unmarshal(resources)
	return
}
