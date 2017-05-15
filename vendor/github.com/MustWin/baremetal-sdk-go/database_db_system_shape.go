// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

// DBSystemShape describes the shape of the DB System.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbSystemShape/
type DBSystemShape struct {
	AvailableCoreCount uint64 `json:"availableCoreCount"`
	Name               string `json:"name"`
	Shape              string `json:"shape"`
}

// ListDBSystemShapes contains a list of DBSystemShapes.
type ListDBSystemShapes struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
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
	opts *ListOptions,
) (resources *ListDBSystemShapes, e error) {

	required := struct {
		listOCIDRequirement
		AvailabilityDomain string `header:"-" json:"-" url:"availabilityDomain"`
	}{
		AvailabilityDomain: availabilityDomain,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceDBSystemShapes,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	resources = &ListDBSystemShapes{}
	e = resp.unmarshal(resources)
	return
}
