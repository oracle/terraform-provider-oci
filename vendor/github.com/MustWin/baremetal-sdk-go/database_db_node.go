package baremetal

import (
	"net/http"
	"time"
)

type DBNode struct {
	ETaggedResource
	DBSystemID  string    `json:"dbSystemId"`
	Hostname    string    `json:"hostname"`
	ID          string    `json:"id"`
	State       string    `json:"lifecycleState"`
	TimeCreated time.Time `json:"timeCreated"`
	VnicID      string    `json:"vnicId"`
}

type ListDBNodes struct {
	ResourceContainer
	DBNodes []DBNode
}

func (l *ListDBNodes) GetList() interface{} {
	return &l.DBNodes
}

// GetDBNode retrieves information about a DBNode
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbNode/GetDbNode
func (c *Client) GetDBNode(id string) (res *DBNode, e error) {
	details := &requestDetails{
		name: resourceDBNodes,
		ids:  urlParts{id},
	}

	var resp *requestResponse
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	res = &DBNode{}
	e = resp.unmarshal(res)
	return
}

// DBNodeAction starts, stops, or resets a compute instance identified by
// instanceID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/database/20160918/DbNode/DbNodeAction
func (c *Client) DBNodeAction(id string, action DBNodeAction, opts *HeaderOptions) (inst *DBNode, e error) {
	required := struct {
		Action string `json:"-" url:"action"`
	}{
		Action: string(action),
	}

	details := &requestDetails{
		name:     resourceDBNodes,
		ids:      urlParts{id},
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.databaseApi.request(http.MethodPost, details); e != nil {
		return
	}

	inst = &DBNode{}
	e = response.unmarshal(inst)
	return
}

// ListDBNodes returns a list of database nodes in the specified DB System. The request MAY contain optional paging arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/database/20160918/DbNode/ListDbNodes
func (c *Client) ListDBNodes(compartmentID, dbSystemID string, limit uint64, opts *PageListOptions) (resources *ListDBNodes, e error) {
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
		name:     resourceDBNodes,
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	resources = &ListDBNodes{}
	e = response.unmarshal(resources)
	return
}
