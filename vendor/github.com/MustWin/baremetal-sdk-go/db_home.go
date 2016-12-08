package baremetal

import (
	"time"
)

type DBHome struct {
	ocidRequirement
	ETaggedResource
	DBSystemID  string    `json:"dbSystemId"`
	DBVersion   string    `json:"dbVersion"`
	DisplayName string    `json:"displayName"`
	ID          string    `json:"id"`
	State       string    `json:"lifecycleState"`
	TimeCreated time.Time `json:"timeCreated"`
}

type ListDBHomes struct {
	ResourceContainer
	DBHomes []DBHome
}

func (l *ListDBHomes) GetList() interface{} {
	return &l.DBHomes
}

// GetDBNode retrieves information about a DBHome
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/database/20160918/DbHome/GetDbHome
func (c *Client) GetDBHome(id string) (res *DBHome, e error) {
	details := &requestDetails{
		name: resourceDBHomes,
		ids:  urlParts{id},
	}

	var resp *requestResponse
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	res = &DBHome{}
	e = resp.unmarshal(res)
	return
}

// ListDBHomes returns a list of database homes in the specified DB System. The request MAY contain optional paging arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/database/20160918/DbHome/ListDbHomes
func (c *Client) ListDBHomes(compartmentID, dbSystemID string, limit uint64, opts *PageListOptions) (resources *ListDBHomes, e error) {
	required := struct {
		listOCIDRequirement
		DBSystemID string `json:"-" url:"dbSystemId"`
		Limit      uint64 `json:"-" url:"limit"`
	}{
		DBSystemID: dbSystemID,
		Limit:      limit,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceDBHomes,
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	resources = &ListDBHomes{}
	e = response.unmarshal(resources)
	return
}
