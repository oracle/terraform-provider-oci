// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDbSystemDetails Details required to create a DB System.
type CreateDbSystemDetails struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the shape. The shape determines the resources allocated
	// - CPU cores and memory for VM shapes; CPU cores, memory and storage
	// for non-VM (or bare metal) shapes. To get a list of shapes, use the
	// ListShapes operation.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The OCID of the subnet the DB System is associated with.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The user-friendly name for the DB System. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// User-provided data about the DB System.
	Description *string `mandatory:"false" json:"description"`

	// Specifies if the DB System is highly available.
	// When creating a DB System with High Availability, three instances
	// are created and placed according to your region- and
	// subnet-type. The secondaries are placed automatically in the other
	// two availability or fault domains.  You can choose the preferred
	// location of your primary instance, only.
	IsHighlyAvailable *bool `mandatory:"false" json:"isHighlyAvailable"`

	// The availability domain on which to deploy the Read/Write endpoint. This defines the preferred primary instance.
	// In a failover scenario, the Read/Write endpoint is redirected to one of the other availability domains
	// and the MySQL instance in that domain is promoted to the primary instance.
	// This redirection does not affect the IP address of the DB System in any way.
	// For a standalone DB System, this defines the availability domain in which the DB System is placed.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The fault domain on which to deploy the Read/Write endpoint. This defines the preferred primary instance.
	// In a failover scenario, the Read/Write endpoint is redirected to one of the other fault domains
	// and the MySQL instance in that domain is promoted to the primary instance.
	// This redirection does not affect the IP address of the DB System in any way.
	// For a standalone DB System, this defines the fault domain in which the DB System is placed.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// The OCID of the Configuration to be used for this DB System.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	// The specific MySQL version identifier.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	// The username for the administrative user.
	AdminUsername *string `mandatory:"false" json:"adminUsername"`

	// The password for the administrative user. The password must be
	// between 8 and 32 characters long, and must contain at least 1
	// numeric character, 1 lowercase character, 1 uppercase character, and
	// 1 special (nonalphanumeric) character.
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

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

	DeletionPolicy *CreateDeletionPolicyDetails `mandatory:"false" json:"deletionPolicy"`

	// Whether to run the DB System with InnoDB Redo Logs and the Double Write Buffer enabled or disabled,
	// and whether to enable or disable syncing of the Binary Logs.
	CrashRecovery CrashRecoveryStatusEnum `mandatory:"false" json:"crashRecovery,omitempty"`
}

func (m CreateDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCrashRecoveryStatusEnum(string(m.CrashRecovery)); !ok && m.CrashRecovery != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CrashRecovery: %s. Supported values are: %s.", m.CrashRecovery, strings.Join(GetCrashRecoveryStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDbSystemDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                           `json:"displayName"`
		Description          *string                           `json:"description"`
		IsHighlyAvailable    *bool                             `json:"isHighlyAvailable"`
		AvailabilityDomain   *string                           `json:"availabilityDomain"`
		FaultDomain          *string                           `json:"faultDomain"`
		ConfigurationId      *string                           `json:"configurationId"`
		MysqlVersion         *string                           `json:"mysqlVersion"`
		AdminUsername        *string                           `json:"adminUsername"`
		AdminPassword        *string                           `json:"adminPassword"`
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
		DeletionPolicy       *CreateDeletionPolicyDetails      `json:"deletionPolicy"`
		CrashRecovery        CrashRecoveryStatusEnum           `json:"crashRecovery"`
		CompartmentId        *string                           `json:"compartmentId"`
		ShapeName            *string                           `json:"shapeName"`
		SubnetId             *string                           `json:"subnetId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.IsHighlyAvailable = model.IsHighlyAvailable

	m.AvailabilityDomain = model.AvailabilityDomain

	m.FaultDomain = model.FaultDomain

	m.ConfigurationId = model.ConfigurationId

	m.MysqlVersion = model.MysqlVersion

	m.AdminUsername = model.AdminUsername

	m.AdminPassword = model.AdminPassword

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

	m.DeletionPolicy = model.DeletionPolicy

	m.CrashRecovery = model.CrashRecovery

	m.CompartmentId = model.CompartmentId

	m.ShapeName = model.ShapeName

	m.SubnetId = model.SubnetId

	return
}
