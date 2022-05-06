// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDbSystemDetails Details required to update a DB System.
type UpdateDbSystemDetails struct {

	// The user-friendly name for the DB System. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// User-provided data about the DB System.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the subnet the DB System is associated with.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Specifies if the DB System is highly available.
	// Set to true to enable high availability. Two secondary MySQL instances are created and placed in the unused
	// availability or fault domains, depending on your region and subnet type.
	// Set to false to disable high availability. The secondary MySQL instances are removed and the MySQL instance
	// in the preferred location is used.
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

	// The shape of the DB System. The shape determines resources
	// allocated to the DB System - CPU cores and memory for VM
	// shapes; CPU cores, memory and storage for non-VM (or bare metal)
	// shapes. To get a list of shapes, use the
	// ListShapes
	// operation.
	// Changes in Shape will result in a downtime as the MySQL DB System is
	// migrated to the new Compute instance.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// The specific MySQL version identifier.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	// The OCID of the Configuration to be used for Instances in this DB System.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	// The username for the administrative user for the MySQL Instance.
	AdminUsername *string `mandatory:"false" json:"adminUsername"`

	// The password for the administrative user. The password must be
	// between 8 and 32 characters long, and must contain at least 1
	// numeric character, 1 lowercase character, 1 uppercase character, and
	// 1 special (nonalphanumeric) character.
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

	// Expands the DB System's storage to the specified value. Only supports values larger than the current DB System's
	// storage size.
	// DB Systems with initial storage of 400 GB or less can be expanded up to 32 TB. DB Systems with initial storage
	// larger than 400 GB can be expanded up to 64 TB.
	// It is not possible to decrease data storage size.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The hostname for the primary endpoint of the DB System. Used for DNS.
	// The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN)
	// (for example, "dbsystem-1" in FQDN "dbsystem-1.subnet123.vcn1.oraclevcn.com").
	// Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123.
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// The IP address the DB System should be configured to listen on the provided subnet.
	// It must be a free private IP address within the subnet's CIDR. If you don't specify a
	// value, Oracle automatically assigns a private IP address from the subnet. This should
	// be a "dotted-quad" style IPv4 address.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The port for primary endpoint of the DB System to listen on.
	Port *int `mandatory:"false" json:"port"`

	// The TCP network port on which X Plugin listens for connections. This is the X Plugin equivalent of port.
	PortX *int `mandatory:"false" json:"portX"`

	BackupPolicy *UpdateBackupPolicyDetails `mandatory:"false" json:"backupPolicy"`

	Maintenance *UpdateMaintenanceDetails `mandatory:"false" json:"maintenance"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	DeletionPolicy *UpdateDeletionPolicyDetails `mandatory:"false" json:"deletionPolicy"`

	// Whether to run the DB System with InnoDB Redo Logs and the Double Write Buffer enabled or disabled,
	// and whether to enable or disable syncing of the Binary Logs.
	CrashRecovery CrashRecoveryStatusEnum `mandatory:"false" json:"crashRecovery,omitempty"`
}

func (m UpdateDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCrashRecoveryStatusEnum(string(m.CrashRecovery)); !ok && m.CrashRecovery != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CrashRecovery: %s. Supported values are: %s.", m.CrashRecovery, strings.Join(GetCrashRecoveryStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
