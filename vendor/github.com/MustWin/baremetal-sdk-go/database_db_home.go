// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"time"
)

type DBHome struct {
	ocidRequirement
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	DBSystemID  string    `json:"dbSystemId"`
	DBVersion   string    `json:"dbVersion"`
	DisplayName string    `json:"displayName"`
	ID          string    `json:"id"`
	State       string    `json:"lifecycleState"`
	TimeCreated time.Time `json:"timeCreated"`
}

type ListDBHomes struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	DBHomes []DBHome
}

func (l *ListDBHomes) GetList() interface{} {
	return &l.DBHomes
}

// GetDBHome retrieves information about a DBHome
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbHome/GetDbHome
func (c *Client) GetDBHome(id string) (res *DBHome, e error) {
	details := &requestDetails{
		name: resourceDBHomes,
		ids:  urlParts{id},
	}

	var resp *response
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	res = &DBHome{}
	e = resp.unmarshal(res)
	return
}

// ListDBHomes returns a list of database homes in the specified DB System. The request MAY contain optional paging arguments.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbHome/ListDbHomes
func (c *Client) ListDBHomes(compartmentID, dbSystemID string, opts *ListOptions) (resources *ListDBHomes, e error) {
	required := struct {
		listOCIDRequirement
		DBSystemID string `header:"-" json:"-" url:"dbSystemId"`
	}{
		DBSystemID: dbSystemID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceDBHomes,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	resources = &ListDBHomes{}
	e = resp.unmarshal(resources)
	return
}

type CreateDBHomeDetails struct {
	Database    CreateDatabaseDetails `header:"-" json:"database" url:"-"`
	DBVersion   string                `header:"-" json:"dbVersion" url:"-"`
	DisplayName string                `header:"-" json:"displayName,omitempty" url:"-"`
}

// NewCreateDBHomeDetails is used to create the DBHome argument to
// LaunchDBSystem.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/requests/CreateDbHomeDetails
func NewCreateDBHomeDetails(createDatabaseDetails CreateDatabaseDetails, dbVersion string, opts *CreateDBHomeOptions) (dbHome CreateDBHomeDetails) {
	dbHome = CreateDBHomeDetails{
		Database:  createDatabaseDetails,
		DBVersion: dbVersion,
	}

	if opts != nil {
		if opts.DisplayName != "" {
			dbHome.DisplayName = opts.DisplayName
		}
	}
	return
}
