// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

type DBHome struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	CompartmentID           string `json:"compartmentId"`
	DBSystemID              string `json:"dbSystemId"`
	DBVersion               string `json:"dbVersion"`
	DisplayName             string `json:"displayName"`
	ID                      string `json:"id"`
	LastPatchHistoryEntryID string `json:"lastPatchHistoryEntryId"`
	State                   string `json:"lifecycleState"`
	TimeCreated             Time   `json:"timeCreated"`
}

type CreateDBHomeOptions struct {
	CreateOptions
}

type UpdateDBHomeOptions struct {
	UpdateOptions
	DBVersion string `header:"-" json:"dbVersion,omitempty" url:"-"`
}

type ListDBHomesOptions struct {
	ListOptions
}

type ListDBHomes struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	DBHomes []DBHome
}

func (l *ListDBHomes) GetList() interface{} {
	return &l.DBHomes
}

type CreateDatabaseDetails struct {
	AdminPassword string `header:"-" json:"adminPassword" url:"-"`
	CharacterSet  string `header:"-" json:"characterSet,omitempty" url:"-"`
	DBName        string `header:"-" json:"dbName" url:"-"`
	DBWorkload    string `header:"-" json:"dbWorkload,omitempty" url:"-"`
	NcharacterSet string `header:"-" json:"ncharacterSet,omitempty" url:"-"`
	PDBName       string `header:"-" json:"pdbName,omitempty" url:"-"`
}

func (c *Client) CreateDBHome(database *CreateDatabaseDetails, dbSystemID string, dbVersion string, opts *CreateDBHomeOptions) (res *DBHome, e error) {
	required := struct {
		Database   *CreateDatabaseDetails `header:"-" json:"database" url:"-"`
		DBSystemID string                 `header:"-" json:"dbSystemId" url:"-"`
		DBVersion  string                 `header:"-" json:"dbVersion" url:"-"`
	}{}
	required.Database = database
	required.DBSystemID = dbSystemID
	required.DBVersion = dbVersion

	details := &requestDetails{
		name:     resourceDBHomes,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.databaseApi.postRequest(details); e != nil {
		return
	}

	res = &DBHome{}
	e = resp.unmarshal(res)
	return
}

func (c *Client) GetDBHome(dbHomeID string) (res *DBHome, e error) {
	required := struct {
	}{}

	details := &requestDetails{
		name: resourceDBHomes,
		ids: urlParts{
			dbHomeID,
		},
		required: required,
	}

	var resp *response
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	res = &DBHome{}
	e = resp.unmarshal(res)
	return
}

func (c *Client) UpdateDBHome(dbHomeID string, opts *UpdateDBHomeOptions) (res *DBHome, e error) {
	required := struct {
	}{}

	details := &requestDetails{
		name: resourceDBHomes,
		ids: urlParts{
			dbHomeID,
		},
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.databaseApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &DBHome{}
	e = resp.unmarshal(res)
	return
}

func (c *Client) DeleteDBHome(dbHomeID string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		name: resourceDBHomes,
		ids: urlParts{
			dbHomeID,
		},
		optional: opts,
	}
	return c.databaseApi.deleteRequest(details)
}

func (c *Client) ListDBHomes(compartmentID string, dbSystemID string, opts *ListDBHomesOptions) (res *ListDBHomes, e error) {
	required := struct {
		CompartmentID string `header:"-" json:"-" url:"compartmentId"`
		DBSystemID    string `header:"-" json:"-" url:"dbSystemId"`
	}{}
	required.CompartmentID = compartmentID
	required.DBSystemID = dbSystemID

	details := &requestDetails{
		name:     resourceDBHomes,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	res = &ListDBHomes{}
	e = resp.unmarshal(res)
	return
}

// The following section is manually added to the generated code above. Still used in a few places.

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
