// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

// DBSystem described a dedicated bare metal instance running Oracle Linux 6.8.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbSystem/
type DBSystem struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	AvailabilityDomain string              `json:"availabilityDomain"`
	CompartmentID      string              `json:"compartmentId"`
	CPUCoreCount       uint64              `json:"cpuCoreCount"`
	DatabaseEdition    DatabaseEdition     `json:"databaseEdition"`
	DBHome             createDBHomeDetails `json:"dbHome"`
	DiskRedundancy     DiskRedundancy      `json:"diskRedundancy"`
	DisplayName        string              `json:"displayName"`
	Domain             string              `json:"domain"`
	Hostname           string              `json:"hostname"`
	ID                 string              `json:"id"`
	LifecycleDetails   string              `json:"lifecycleDetails"`
	State              string              `json:"lifecycleState"`
	ListenerPort       uint64              `json:"listenerPort"`
	Shape              string              `json:"shape"`
	SSHPublicKeys      []string            `json:"sshPublicKeys"`
	SubnetID           string              `json:"subnetId"`
	TimeCreated        Time                `json:"timeCreated"`
}

// ListDBSystems contains a list of DBSystems.
//
type ListDBSystems struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	DBSystems []DBSystem
}

func (l *ListDBSystems) GetList() interface{} {
	return &l.DBSystems
}

type createDatabaseDetails struct {
	AdminPassword string `header:"-" json:"adminPassword" url:"-"`
	DBName        string `header:"-" json:"dbName" url:"-"`
	CharacterSet  string `header:"-" json:"characterSet,omitempty" url:"-"`
	NCharacterSet string `header:"-" json:"ncharacterSet,omitempty" url:"-"`
}

type createDBHomeDetails struct {
	Database    createDatabaseDetails `header:"-" json:"database" url:"-"`
	DBVersion   string                `header:"-" json:"dbVersion" url:"-"`
	DisplayName string                `header:"-" json:"displayName,omitempty" url:"-"`
}

// NewCreateDBHomeDetails is used to create the optional DBHome argument to
// LaunchDBSystem.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/requests/CreateDbHomeDetails
func NewCreateDBHomeDetails(adminPassword, dbName, dbVersion, characterSet, nCharacterSet string, opts *DisplayNameOptions) (dbHome createDBHomeDetails) {
	dbHome = createDBHomeDetails{
		Database: createDatabaseDetails{
			AdminPassword: adminPassword,
			DBName:        dbName,
			CharacterSet:  characterSet,
			NCharacterSet: nCharacterSet,
		},
		DBVersion: dbVersion,
	}

	if opts != nil && opts.DisplayName != "" {
		dbHome.DisplayName = opts.DisplayName
	}
	return
}

// LaunchDBSystem launches a new DB System in the specified compartment and
// Availability Domain.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbSystem/LaunchDbSystem
func (c *Client) LaunchDBSystem(
	availabilityDomain, compartmentID, shape, subnetID string,
	sshPublicKeys []string,
	cpuCoreCount uint64,
	opts *LaunchDBSystemOptions,
) (res *DBSystem, e error) {
	required := struct {
		ocidRequirement
		AvailabilityDomain string   `header:"-" json:"availabilityDomain" url:"-"`
		CPUCoreCount       uint64   `header:"-" json:"cpuCoreCount" url:"-"`
		Shape              string   `header:"-" json:"shape" url:"-"`
		SSHPublicKeys      []string `header:"-" json:"sshPublicKeys" url:"-"`
		SubnetID           string   `header:"-" json:"subnetId" url:"-"`
	}{
		AvailabilityDomain: availabilityDomain,
		CPUCoreCount:       cpuCoreCount,
		Shape:              shape,
		SSHPublicKeys:      sshPublicKeys,
		SubnetID:           subnetID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceDBSystems,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.databaseApi.request(http.MethodPost, details); e != nil {
		return
	}

	res = &DBSystem{}
	e = resp.unmarshal(res)
	return
}

// GetDBSystem gets information about the specified DB System.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbSystem/GetDbSystem
func (c *Client) GetDBSystem(id string) (res *DBSystem, e error) {
	details := &requestDetails{
		name: resourceDBSystems,
		ids:  urlParts{id},
	}

	var resp *response
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	res = &DBSystem{}
	e = resp.unmarshal(res)
	return
}

// TerminateDBSystem terminates a DB System and permanently deletes it and any
// databases running on it.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbSystem/TerminateDbSystem
func (c *Client) TerminateDBSystem(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceDBSystems,
		optional: opts,
	}
	return c.databaseApi.deleteRequest(details)
}

// ListDBSystems gets a list of the DB Systems in the specified compartment.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbSystem/ListDbSystems
func (c *Client) ListDBSystems(compartmentID string, opts *ListOptions) (res *ListDBSystems, e error) {
	required := struct {
		listOCIDRequirement
	}{}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceDBSystems,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	res = &ListDBSystems{}
	e = resp.unmarshal(res)
	return
}
