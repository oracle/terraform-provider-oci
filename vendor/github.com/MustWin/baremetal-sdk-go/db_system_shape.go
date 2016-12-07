package baremetal

// DBSystemShape describes the shape of the DB System.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbSystemShape/
type DBSystemShape struct {
	AvailableCoreCount string `json:"availableCoreCount"`
	Name               string `json:"name"`
	Shape              string `json:"shape"`
}

// ListDBSystemShapes contains a list of DBSystemShapes.
type ListDBSystemShapes struct {
	ResourceContainer
	DBSystemShapes []DBSystemShape
}

// GetList returns the list of DBSystemShapes.
func (l *ListDBSystemShapes) GetList() interface{} {
	return &l.DBSystemShapes
}

// ListDBSystemShapes returns a set of DBSystemShapes.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbSystemShape/ListDbSystemShapes
func (c *Client) ListDBSystemShapes(
	availabilityDomain, compartmentID string,
	limit uint64,
	opts *PageListOptions,
) (resources *ListDBSystemShapes, e error) {

	required := struct {
		listOCIDRequirement
		AvailabilityDomain string `json:"-" url:"availabilityDomain"`
		Limit              uint64 `json:"-" url:"limit"`
	}{
		AvailabilityDomain: availabilityDomain,
		Limit:              limit,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceDBSystemShapes,
		optional: opts,
		required: required,
	}

	var resp *requestResponse
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	resources = &ListDBSystemShapes{}
	e = resp.unmarshal(resources)
	return
}
