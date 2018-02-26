// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "time"

type Database struct {
	ETagUnmarshaller
	OPCRequestIDUnmarshaller
	ocidRequirement
	CharacterSet     string    `json:"characterSet"`
	DBHomeID         string    `json:"dbHomeId"`
	DBName           string    `json:"dbName"`
	DBUniqueName     string    `json:"dbUniqueName"`
	DBWorkload       string    `json:"dbWorkload"`
	ID               string    `json:"id"`
	LifecycleDetails string    `json:"lifecycleDetails"`
	NcharacterSet    string    `json:"ncharacterSet"`
	PDBName          string    `json:"pdbName"`
	State            string    `json:"lifecycleState"`
	TimeCreated      time.Time `json:"timeCreated"`
}

type ListDatabases struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Databases []Database
}

func (l *ListDatabases) GetList() interface{} {
	return &l.Databases
}

// GetDatabase retrieves information about a Database
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/Database/GetDatabase
func (c *Client) GetDatabase(id string) (res *Database, e error) {
	details := &requestDetails{
		name: resourceDatabases,
		ids:  urlParts{id},
	}

	var resp *response
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	res = &Database{}
	e = resp.unmarshal(res)
	return
}

// ListDatabases returns a list of supported Oracle database versions. The request MAY contain optional paging arguments.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/Database/ListDatabases
func (c *Client) ListDatabases(compartmentID, dbHomeID string, limit uint64, opts *PageListOptions) (resources *ListDatabases, e error) {
	required := struct {
		listOCIDRequirement
		DBHomeID string `header:"-" json:"-" url:"dbHomeId"`
		Limit    uint64 `header:"-" json:"-" url:"limit"`
	}{
		DBHomeID: dbHomeID,
		Limit:    limit,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceDatabases,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	resources = &ListDatabases{}
	e = resp.unmarshal(resources)
	return
}

type CreateDatabaseDetails struct {
	AdminPassword string `header:"-" json:"adminPassword" url:"-"`
	DBName        string `header:"-" json:"dbName" url:"-"`
	DBWorkload    string `header:"-" json:"dbWorkload,omitempty" url:"-"`
	CharacterSet  string `header:"-" json:"characterSet,omitempty" url:"-"`
	NCharacterSet string `header:"-" json:"ncharacterSet,omitempty" url:"-"`
	PDBName       string `header:"-" json:"pdbName,omitempty" url:"-"`
}

func NewCreateDatabaseDetails(adminPassword, dbName string, opts *CreateDatabaseOptions) (db CreateDatabaseDetails) {
	db = CreateDatabaseDetails{
		AdminPassword: adminPassword,
		DBName:        dbName,
	}
	if opts != nil {
		if opts.DBWorkload != "" {
			db.DBWorkload = opts.DBWorkload
		}
		if opts.CharacterSet != "" {
			db.CharacterSet = opts.CharacterSet
		}
		if opts.NCharacterSet != "" {
			db.NCharacterSet = opts.NCharacterSet
		}
		if opts.PDBName != "" {
			db.PDBName = opts.PDBName
		}
	}
	return
}
