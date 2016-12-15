package baremetal

type SupportedOperation struct {
	ID string `json:"id"`
}

type ListSupportedOperations struct {
	RequestableResource
	SupportedOperations []SupportedOperation
}

func (l *ListSupportedOperations) GetList() interface{} {
	return &l.SupportedOperations
}

// ListDBSupportedOperations returns a list of supported operations.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/database/20160918/SupportedOperation/ListSupportedOperations
func (c *Client) ListSupportedOperations() (resources *ListSupportedOperations, e error) {

	details := &requestDetails{
		name: resourceDBSupportedOperations,
	}

	var response *requestResponse
	if response, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	resources = &ListSupportedOperations{}
	e = response.unmarshal(resources)
	return
}
