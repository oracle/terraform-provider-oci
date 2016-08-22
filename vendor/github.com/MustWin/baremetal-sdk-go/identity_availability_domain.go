package baremetal

// AvailablityDomain contains name and then tenancy ID that an

// availability domain belongs to.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#AvailabilityDomain
type AvailabilityDomain struct {
	Name          string `json:"name"`
	CompartmentID string `json:"compartmentId"`
}

type AvailabilityDomainList struct {
	ResourceContainer
	AvailabilityDomains []AvailabilityDomain
}

func (l *AvailabilityDomainList) GetList() interface{} {
	return &l.AvailabilityDomains
}

// ListAvailablityDomains lists availability domains in a user's root tenancy.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listAvailabilityDomains
func (c *Client) ListAvailablityDomains(compartmentID string) (ads *AvailabilityDomainList, e error) {
	reqOpts := &sdkRequestOptions{
		name: resourceAvailabilityDomains,
		ocid: compartmentID,
	}

	var getResp *requestResponse
	if getResp, e = c.identityApi.getRequest(reqOpts); e != nil {
		return
	}

	ads = &AvailabilityDomainList{}
	e = getResp.unmarshal(ads)
	return
}
