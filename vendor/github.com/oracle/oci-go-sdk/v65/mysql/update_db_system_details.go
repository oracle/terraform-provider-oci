// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

	// Network Security Group OCIDs used for the VNIC attachment.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Security Attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see ZPR Artifacts (https://docs.oracle.com/en-us/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// The database mode indicating the types of statements that will be allowed to run in the DB system.
	// This mode will apply only to statements run by user connections. Replicated write statements will continue
	// to be allowed regardless of the DatabaseMode.
	//   - READ_WRITE: allow running read and write statements on the DB system;
	//   - READ_ONLY: only allow running read statements on the DB system.
	DatabaseMode DbSystemDatabaseModeEnum `mandatory:"false" json:"databaseMode,omitempty"`

	// The access mode indicating if the database access will be restricted only to administrators or not:
	//  - UNRESTRICTED: the access to the database is not restricted;
	//  - RESTRICTED: the access will be allowed only to users with specific privileges;
	//    RESTRICTED will correspond to setting the MySQL system variable
	//    offline_mode (https://dev.mysql.com/doc/en/server-system-variables.html#sysvar_offline_mode) to ON.
	AccessMode DbSystemAccessModeEnum `mandatory:"false" json:"accessMode,omitempty"`

	Rest *UpdateRestDetails `mandatory:"false" json:"rest"`

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
	// DB Systems with an initial storage size of 400 GB or less can be expanded up to 32 TB.
	// DB Systems with an initial storage size between 401-800 GB can be expanded up to 64 TB.
	// DB Systems with an initial storage size between 801-1200 GB can be expanded up to 96 TB.
	// DB Systems with an initial storage size of 1201 GB or more can be expanded up to 128 TB.
	// It is not possible to decrease data storage size.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	DataStorage *DataStorageDetails `mandatory:"false" json:"dataStorage"`

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

	// Whether to enable monitoring via the Database Management service.
	DatabaseManagement DatabaseManagementStatusEnum `mandatory:"false" json:"databaseManagement,omitempty"`

	SecureConnections *SecureConnectionDetails `mandatory:"false" json:"secureConnections"`

	EncryptData *EncryptDataDetails `mandatory:"false" json:"encryptData"`

	// The list of customer email addresses that receive information from Oracle about the specified OCI DB System resource.
	// Oracle uses these email addresses to send notifications about planned and unplanned software maintenance updates, information about system hardware, and other information needed by administrators.
	// Up to 10 email addresses can be added to the customer contacts for a DB System.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	ReadEndpoint *UpdateReadEndpointDetails `mandatory:"false" json:"readEndpoint"`
}

func (m UpdateDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDbSystemDatabaseModeEnum(string(m.DatabaseMode)); !ok && m.DatabaseMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseMode: %s. Supported values are: %s.", m.DatabaseMode, strings.Join(GetDbSystemDatabaseModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbSystemAccessModeEnum(string(m.AccessMode)); !ok && m.AccessMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessMode: %s. Supported values are: %s.", m.AccessMode, strings.Join(GetDbSystemAccessModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCrashRecoveryStatusEnum(string(m.CrashRecovery)); !ok && m.CrashRecovery != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CrashRecovery: %s. Supported values are: %s.", m.CrashRecovery, strings.Join(GetCrashRecoveryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseManagementStatusEnum(string(m.DatabaseManagement)); !ok && m.DatabaseManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseManagement: %s. Supported values are: %s.", m.DatabaseManagement, strings.Join(GetDatabaseManagementStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
