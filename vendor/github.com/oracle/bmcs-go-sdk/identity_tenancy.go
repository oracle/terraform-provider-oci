// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

type Tenancy struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	Description   string `json:"description"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	HomeRegionKey string `json:"homeRegion"`
}

// GetTenancy returns the tenancy identified by id.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Tenancy/GetTenancy
func (c *Client) GetTenancy(id string) (res *Tenancy, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourceTenancies,
	}

	var resp *response
	if resp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	res = &Tenancy{}
	e = resp.unmarshal(res)
	return
}
