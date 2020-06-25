// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// DbSystem A DB System is the core logical unit of MySQL Database Service.
type DbSystem struct {

	// The OCID of the DB System.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the DB System. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment the DB System belongs in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the subnet the DB System is associated with.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Name of the MySQL Version in use for the DB System.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// Initial size of the data volume in GiBs that will be created and attached.
	DataStorageSizeInGBs *int `mandatory:"true" json:"dataStorageSizeInGBs"`

	// The current state of the DB System.
	LifecycleState DbSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	Maintenance *MaintenanceDetails `mandatory:"true" json:"maintenance"`

	// The date and time the DB System was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the DB System was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// User-provided data about the DB System.
	Description *string `mandatory:"false" json:"description"`

	// The Availability Domain where the primary DB System should be located.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The name of the Fault Domain the DB System is located in.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// The shape of the primary instances of the DB System. The shape
	// determines resources allocated to a DB System - CPU cores
	// and memory for VM shapes; CPU cores, memory and storage for non-VM
	// (or bare metal) shapes. To get a list of shapes, use (the
	// ListShapes operation.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	BackupPolicy *BackupPolicy `mandatory:"false" json:"backupPolicy"`

	Source DbSystemSource `mandatory:"false" json:"source"`

	// The OCID of the Configuration to be used for Instances in this DB System.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	// The hostname for the primary endpoint of the DB System. Used for DNS.
	// The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN)
	// (for example, "dbsystem-1" in FQDN "dbsystem-1.subnet123.vcn1.oraclevcn.com").
	// Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123.
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// The IP address the DB System is configured to listen on. A private
	// IP address of the primary endpoint of the DB System. Must be an
	// available IP address within the subnet's CIDR. This will be a
	// "dotted-quad" style IPv4 address.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The port for primary endpoint of the DB System to listen on.
	Port *int `mandatory:"false" json:"port"`

	// The network port on which X Plugin listens for TCP/IP connections. This is the X Plugin equivalent of port.
	PortX *int `mandatory:"false" json:"portX"`

	// The network endpoints available for this DB System.
	Endpoints []DbSystemEndpoint `mandatory:"false" json:"endpoints"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DbSystem) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DbSystem) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description          *string                           `json:"description"`
		AvailabilityDomain   *string                           `json:"availabilityDomain"`
		FaultDomain          *string                           `json:"faultDomain"`
		ShapeName            *string                           `json:"shapeName"`
		BackupPolicy         *BackupPolicy                     `json:"backupPolicy"`
		Source               dbsystemsource                    `json:"source"`
		ConfigurationId      *string                           `json:"configurationId"`
		HostnameLabel        *string                           `json:"hostnameLabel"`
		IpAddress            *string                           `json:"ipAddress"`
		Port                 *int                              `json:"port"`
		PortX                *int                              `json:"portX"`
		Endpoints            []DbSystemEndpoint                `json:"endpoints"`
		LifecycleDetails     *string                           `json:"lifecycleDetails"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
		Id                   *string                           `json:"id"`
		DisplayName          *string                           `json:"displayName"`
		CompartmentId        *string                           `json:"compartmentId"`
		SubnetId             *string                           `json:"subnetId"`
		MysqlVersion         *string                           `json:"mysqlVersion"`
		DataStorageSizeInGBs *int                              `json:"dataStorageSizeInGBs"`
		LifecycleState       DbSystemLifecycleStateEnum        `json:"lifecycleState"`
		Maintenance          *MaintenanceDetails               `json:"maintenance"`
		TimeCreated          *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated          *common.SDKTime                   `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.AvailabilityDomain = model.AvailabilityDomain

	m.FaultDomain = model.FaultDomain

	m.ShapeName = model.ShapeName

	m.BackupPolicy = model.BackupPolicy

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(DbSystemSource)
	} else {
		m.Source = nil
	}

	m.ConfigurationId = model.ConfigurationId

	m.HostnameLabel = model.HostnameLabel

	m.IpAddress = model.IpAddress

	m.Port = model.Port

	m.PortX = model.PortX

	m.Endpoints = make([]DbSystemEndpoint, len(model.Endpoints))
	for i, n := range model.Endpoints {
		m.Endpoints[i] = n
	}

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.SubnetId = model.SubnetId

	m.MysqlVersion = model.MysqlVersion

	m.DataStorageSizeInGBs = model.DataStorageSizeInGBs

	m.LifecycleState = model.LifecycleState

	m.Maintenance = model.Maintenance

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}

// DbSystemLifecycleStateEnum Enum with underlying type: string
type DbSystemLifecycleStateEnum string

// Set of constants representing the allowable values for DbSystemLifecycleStateEnum
const (
	DbSystemLifecycleStateCreating DbSystemLifecycleStateEnum = "CREATING"
	DbSystemLifecycleStateActive   DbSystemLifecycleStateEnum = "ACTIVE"
	DbSystemLifecycleStateInactive DbSystemLifecycleStateEnum = "INACTIVE"
	DbSystemLifecycleStateUpdating DbSystemLifecycleStateEnum = "UPDATING"
	DbSystemLifecycleStateDeleting DbSystemLifecycleStateEnum = "DELETING"
	DbSystemLifecycleStateDeleted  DbSystemLifecycleStateEnum = "DELETED"
	DbSystemLifecycleStateFailed   DbSystemLifecycleStateEnum = "FAILED"
)

var mappingDbSystemLifecycleState = map[string]DbSystemLifecycleStateEnum{
	"CREATING": DbSystemLifecycleStateCreating,
	"ACTIVE":   DbSystemLifecycleStateActive,
	"INACTIVE": DbSystemLifecycleStateInactive,
	"UPDATING": DbSystemLifecycleStateUpdating,
	"DELETING": DbSystemLifecycleStateDeleting,
	"DELETED":  DbSystemLifecycleStateDeleted,
	"FAILED":   DbSystemLifecycleStateFailed,
}

// GetDbSystemLifecycleStateEnumValues Enumerates the set of values for DbSystemLifecycleStateEnum
func GetDbSystemLifecycleStateEnumValues() []DbSystemLifecycleStateEnum {
	values := make([]DbSystemLifecycleStateEnum, 0)
	for _, v := range mappingDbSystemLifecycleState {
		values = append(values, v)
	}
	return values
}
