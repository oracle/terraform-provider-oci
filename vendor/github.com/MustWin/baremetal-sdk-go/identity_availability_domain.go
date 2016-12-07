package baremetal

// AvailablityDomain contains name and then tenancy ID that an

// availability domain belongs to.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/AvailabilityDomain/
type AvailabilityDomain struct {
	Name          string `json:"name"`
	CompartmentID string `json:"compartmentId"`
}

// ListAvailabilityDomains contains a list AvailabilityDomain
type ListAvailabilityDomains struct {
	ResourceContainer
	AvailabilityDomains []AvailabilityDomain
}

func (l *ListAvailabilityDomains) GetList() interface{} {
	return &l.AvailabilityDomains
}

// ListAvailablityDomains lists availability domains in a user's root tenancy.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/AvailabilityDomain/ListAvailabilityDomains
func (c *Client) ListAvailabilityDomains(compartmentID string) (ads *ListAvailabilityDomains, e error) {
	details := &requestDetails{
		name:     resourceAvailabilityDomains,
		required: listOCIDRequirement{CompartmentID: compartmentID},
	}

	var getResp *requestResponse
	if getResp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	ads = &ListAvailabilityDomains{}
	e = getResp.unmarshal(ads)
	return
}
