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

// CreateDbSystemDetails Details required to create a DB System.
type CreateDbSystemDetails struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the Configuration to be used for this DB System.
	ConfigurationId *string `mandatory:"true" json:"configurationId"`

	// The name of the shape. The shape determines the resources allocated
	// - CPU cores and memory for VM shapes; CPU cores, memory and storage
	// for non-VM (or bare metal) shapes. To get a list of shapes, use the
	// ListShapes operation.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The OCID of the subnet the DB System is associated with.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The username for the administrative user.
	AdminUsername *string `mandatory:"true" json:"adminUsername"`

	// The password for the administrative user. The password must be
	// between 8 and 32 characters long, and must contain at least 1
	// numeric character, 1 lowercase character, 1 uppercase character, and
	// 1 special (nonalphanumeric) character.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The user-friendly name for the DB System. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// User-provided data about the DB System.
	Description *string `mandatory:"false" json:"description"`

	// The Availability Domain where the primary instance should be located.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The name of the Fault Domain the DB System is located in.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// The specific MySQL version identifier.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	// Initial size of the data volume in GBs that will be created and attached.
	// Keep in mind that this only specifies the size of the database data volume,
	// the log volume for the database will be scaled appropriately with its shape.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The hostname for the primary endpoint of the DB System. Used for DNS.
	// The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN)
	// (for example, "dbsystem-1" in FQDN "dbsystem-1.subnet123.vcn1.oraclevcn.com").
	// Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123.
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// The IP address the DB System is configured to listen on.
	// A private IP address of your choice to assign to the primary endpoint of the DB System.
	// Must be an available IP address within the subnet's CIDR. If you don't specify a value,
	// Oracle automatically assigns a private IP address from the subnet. This should be a
	// "dotted-quad" style IPv4 address.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The port for primary endpoint of the DB System to listen on.
	Port *int `mandatory:"false" json:"port"`

	// The TCP network port on which X Plugin listens for connections. This is the X Plugin equivalent of port.
	PortX *int `mandatory:"false" json:"portX"`

	BackupPolicy *CreateBackupPolicyDetails `mandatory:"false" json:"backupPolicy"`

	Source CreateDbSystemSourceDetails `mandatory:"false" json:"source"`

	Maintenance *CreateMaintenanceDetails `mandatory:"false" json:"maintenance"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDbSystemDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateDbSystemDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                           `json:"displayName"`
		Description          *string                           `json:"description"`
		AvailabilityDomain   *string                           `json:"availabilityDomain"`
		FaultDomain          *string                           `json:"faultDomain"`
		MysqlVersion         *string                           `json:"mysqlVersion"`
		DataStorageSizeInGBs *int                              `json:"dataStorageSizeInGBs"`
		HostnameLabel        *string                           `json:"hostnameLabel"`
		IpAddress            *string                           `json:"ipAddress"`
		Port                 *int                              `json:"port"`
		PortX                *int                              `json:"portX"`
		BackupPolicy         *CreateBackupPolicyDetails        `json:"backupPolicy"`
		Source               createdbsystemsourcedetails       `json:"source"`
		Maintenance          *CreateMaintenanceDetails         `json:"maintenance"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId        *string                           `json:"compartmentId"`
		ConfigurationId      *string                           `json:"configurationId"`
		ShapeName            *string                           `json:"shapeName"`
		SubnetId             *string                           `json:"subnetId"`
		AdminUsername        *string                           `json:"adminUsername"`
		AdminPassword        *string                           `json:"adminPassword"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.AvailabilityDomain = model.AvailabilityDomain

	m.FaultDomain = model.FaultDomain

	m.MysqlVersion = model.MysqlVersion

	m.DataStorageSizeInGBs = model.DataStorageSizeInGBs

	m.HostnameLabel = model.HostnameLabel

	m.IpAddress = model.IpAddress

	m.Port = model.Port

	m.PortX = model.PortX

	m.BackupPolicy = model.BackupPolicy

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(CreateDbSystemSourceDetails)
	} else {
		m.Source = nil
	}

	m.Maintenance = model.Maintenance

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.ConfigurationId = model.ConfigurationId

	m.ShapeName = model.ShapeName

	m.SubnetId = model.SubnetId

	m.AdminUsername = model.AdminUsername

	m.AdminPassword = model.AdminPassword

	return
}
