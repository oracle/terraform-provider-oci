package baremetal

type DBVersion struct {
	Version string `json:"version"`
}

type ListDBVersions struct {
	ResourceContainer
	DBVersions []DBVersion
}

func (l *ListDBVersions) GetList() interface{} {
	return &l.DBVersions
}

// ListVersions returns a list of supported Oracle database versions. The request MAY contain optional paging arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/database/20160918/DbVersion/ListDbVersions
func (c *Client) ListDBVersions(limit uint64, opts *PageListOptions) (resources *ListDBVersions, e error) {
	required := struct {
		listOCIDRequirement
		Limit uint64 `json:"-" url:"limit"`
	}{
		Limit: limit,
	}
	required.CompartmentID = c.authInfo.tenancyOCID

	details := &requestDetails{
		name:     resourceDBVersions,
		optional: opts, // Page is optional, limit is required and covered by the next line as well
		required: required,
	}

	var response *requestResponse
	if response, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	resources = &ListDBVersions{}
	e = response.unmarshal(resources)
	return
}
