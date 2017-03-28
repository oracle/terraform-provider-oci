// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

type SupportedOperation struct {
	ID string `json:"id"`
}

type ListSupportedOperations struct {
	OPCRequestIDUnmarshaller
	SupportedOperations []SupportedOperation
}

func (l *ListSupportedOperations) GetList() interface{} {
	return &l.SupportedOperations
}

// ListSupportedOperations returns a list of supported operations.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/database/20160918/SupportedOperation/ListSupportedOperations
// FIXME: this endpoint is no longer documented, is it still available? [jh]
func (c *Client) ListSupportedOperations() (resources *ListSupportedOperations, e error) {

	details := &requestDetails{
		name: resourceDBSupportedOperations,
	}

	var resp *response
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	resources = &ListSupportedOperations{}
	e = resp.unmarshal(resources)
	return
}
