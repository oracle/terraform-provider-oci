// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// UpdateAutonomousDatabaseDetails Details to update an Oracle Autonomous Database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateAutonomousDatabaseDetails struct {

	// The number of CPU cores to be made available to the database.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The size, in terabytes, of the data volume that will be attached to the database.
	DataStorageSizeInTBs *int `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The user-friendly name for the Autonomous Database. The name does not have to be unique. Can only be updated for Autonomous Databases
	// using dedicated Exadata infrastructure.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled.
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing. It must be different from the last four passwords and it must not be a password used within the last 24 hours.
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

	// New name for this Autonomous Database.
	// For databases using dedicated Exadata infrastructure, the name must begin with an alphabetic character, and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
	// For databases using shared Exadata infrastructure, the name must begin with an alphabetic character, and can contain a maximum of 14 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
	DbName *string `mandatory:"false" json:"dbName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The Autonomous Database workload type. The following values are valid:
	// - OLTP - indicates an Autonomous Transaction Processing database
	// - DW - indicates an Autonomous Data Warehouse database
	// - AJD - indicates an Autonomous JSON Database
	// - APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type.
	DbWorkload UpdateAutonomousDatabaseDetailsDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm), this attribute must be null because the attribute is already set at the
	// Autonomous Exadata Infrastructure level. When using shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`.
	LicenseModel UpdateAutonomousDatabaseDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Indicates if the database-level access control is enabled.
	// If disabled, database access is defined by the network security rules.
	// If enabled, database access is restricted to the IP addresses defined by the rules specified with the `whitelistedIps` property. While specifying `whitelistedIps` rules is optional,
	//  if database-level access control is enabled and no rules are specified, the database will become inaccessible. The rules can be added later using the `UpdateAutonomousDatabase` API operation or edit option in console.
	// When creating a database clone, the desired access control setting should be specified. By default, database-level access control will be disabled for the clone.
	// This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform.
	IsAccessControlEnabled *bool `mandatory:"false" json:"isAccessControlEnabled"`

	// The client IP access control list (ACL). This feature is available for autonomous databases on shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) and on Exadata Cloud@Customer.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.
	// For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID.
	// Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs.
	// Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]`
	// For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations.
	// Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`
	// For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.
	WhitelistedIps []string `mandatory:"false" json:"whitelistedIps"`

	// This field will be null if the Autonomous Database is not Data Guard enabled or Access Control is disabled.
	// It's value would be `TRUE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses primary IP access control list (ACL) for standby.
	// It's value would be `FALSE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses different IP access control list (ACL) for standby compared to primary.
	ArePrimaryWhitelistedIpsUsed *bool `mandatory:"false" json:"arePrimaryWhitelistedIpsUsed"`

	// The client IP access control list (ACL). This feature is available for autonomous databases on shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) and on Exadata Cloud@Customer.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.
	// For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID.
	// Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs.
	// Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]`
	// For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations.
	// Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`
	// For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.
	StandbyWhitelistedIps []string `mandatory:"false" json:"standbyWhitelistedIps"`

	// Indicates whether to enable or disable auto scaling for the Autonomous Database OCPU core count. Setting to `true` enables auto scaling. Setting to `false` disables auto scaling. The default value is true. Auto scaling is available for databases on shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) only.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// Indicates whether the Autonomous Database is a refreshable clone.
	IsRefreshableClone *bool `mandatory:"false" json:"isRefreshableClone"`

	// The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.
	RefreshableMode UpdateAutonomousDatabaseDetailsRefreshableModeEnum `mandatory:"false" json:"refreshableMode,omitempty"`

	// Indicates whether the Autonomous Database has Data Guard enabled.
	IsDataGuardEnabled *bool `mandatory:"false" json:"isDataGuardEnabled"`

	// A valid Oracle Database version for Autonomous Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// The `DATABASE OPEN` mode. You can open the database in `READ_ONLY` or `READ_WRITE` mode.
	OpenMode UpdateAutonomousDatabaseDetailsOpenModeEnum `mandatory:"false" json:"openMode,omitempty"`

	// The Autonomous Database permission level. Restricted mode allows access only to admin users.
	PermissionLevel UpdateAutonomousDatabaseDetailsPermissionLevelEnum `mandatory:"false" json:"permissionLevel,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// - For Autonomous Database, setting this will disable public secure access to the database.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and the backup subnet.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The private endpoint label for the resource. Setting this to an empty string, after the private endpoint database gets created, will change the same private endpoint database to the public endpoint database.
	PrivateEndpointLabel *string `mandatory:"false" json:"privateEndpointLabel"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Customer Contacts. Setting this to an empty list removes all customer contacts of an Oracle Autonomous Database.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`
}

func (m UpdateAutonomousDatabaseDetails) String() string {
	return common.PointerString(m)
}

// UpdateAutonomousDatabaseDetailsDbWorkloadEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsDbWorkloadEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsDbWorkloadEnum
const (
	UpdateAutonomousDatabaseDetailsDbWorkloadOltp UpdateAutonomousDatabaseDetailsDbWorkloadEnum = "OLTP"
	UpdateAutonomousDatabaseDetailsDbWorkloadDw   UpdateAutonomousDatabaseDetailsDbWorkloadEnum = "DW"
	UpdateAutonomousDatabaseDetailsDbWorkloadAjd  UpdateAutonomousDatabaseDetailsDbWorkloadEnum = "AJD"
	UpdateAutonomousDatabaseDetailsDbWorkloadApex UpdateAutonomousDatabaseDetailsDbWorkloadEnum = "APEX"
)

var mappingUpdateAutonomousDatabaseDetailsDbWorkload = map[string]UpdateAutonomousDatabaseDetailsDbWorkloadEnum{
	"OLTP": UpdateAutonomousDatabaseDetailsDbWorkloadOltp,
	"DW":   UpdateAutonomousDatabaseDetailsDbWorkloadDw,
	"AJD":  UpdateAutonomousDatabaseDetailsDbWorkloadAjd,
	"APEX": UpdateAutonomousDatabaseDetailsDbWorkloadApex,
}

// GetUpdateAutonomousDatabaseDetailsDbWorkloadEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsDbWorkloadEnum
func GetUpdateAutonomousDatabaseDetailsDbWorkloadEnumValues() []UpdateAutonomousDatabaseDetailsDbWorkloadEnum {
	values := make([]UpdateAutonomousDatabaseDetailsDbWorkloadEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsDbWorkload {
		values = append(values, v)
	}
	return values
}

// UpdateAutonomousDatabaseDetailsLicenseModelEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsLicenseModelEnum
const (
	UpdateAutonomousDatabaseDetailsLicenseModelLicenseIncluded     UpdateAutonomousDatabaseDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateAutonomousDatabaseDetailsLicenseModelBringYourOwnLicense UpdateAutonomousDatabaseDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateAutonomousDatabaseDetailsLicenseModel = map[string]UpdateAutonomousDatabaseDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateAutonomousDatabaseDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateAutonomousDatabaseDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateAutonomousDatabaseDetailsLicenseModelEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsLicenseModelEnum
func GetUpdateAutonomousDatabaseDetailsLicenseModelEnumValues() []UpdateAutonomousDatabaseDetailsLicenseModelEnum {
	values := make([]UpdateAutonomousDatabaseDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsLicenseModel {
		values = append(values, v)
	}
	return values
}

// UpdateAutonomousDatabaseDetailsRefreshableModeEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsRefreshableModeEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsRefreshableModeEnum
const (
	UpdateAutonomousDatabaseDetailsRefreshableModeAutomatic UpdateAutonomousDatabaseDetailsRefreshableModeEnum = "AUTOMATIC"
	UpdateAutonomousDatabaseDetailsRefreshableModeManual    UpdateAutonomousDatabaseDetailsRefreshableModeEnum = "MANUAL"
)

var mappingUpdateAutonomousDatabaseDetailsRefreshableMode = map[string]UpdateAutonomousDatabaseDetailsRefreshableModeEnum{
	"AUTOMATIC": UpdateAutonomousDatabaseDetailsRefreshableModeAutomatic,
	"MANUAL":    UpdateAutonomousDatabaseDetailsRefreshableModeManual,
}

// GetUpdateAutonomousDatabaseDetailsRefreshableModeEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsRefreshableModeEnum
func GetUpdateAutonomousDatabaseDetailsRefreshableModeEnumValues() []UpdateAutonomousDatabaseDetailsRefreshableModeEnum {
	values := make([]UpdateAutonomousDatabaseDetailsRefreshableModeEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsRefreshableMode {
		values = append(values, v)
	}
	return values
}

// UpdateAutonomousDatabaseDetailsOpenModeEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsOpenModeEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsOpenModeEnum
const (
	UpdateAutonomousDatabaseDetailsOpenModeOnly  UpdateAutonomousDatabaseDetailsOpenModeEnum = "READ_ONLY"
	UpdateAutonomousDatabaseDetailsOpenModeWrite UpdateAutonomousDatabaseDetailsOpenModeEnum = "READ_WRITE"
)

var mappingUpdateAutonomousDatabaseDetailsOpenMode = map[string]UpdateAutonomousDatabaseDetailsOpenModeEnum{
	"READ_ONLY":  UpdateAutonomousDatabaseDetailsOpenModeOnly,
	"READ_WRITE": UpdateAutonomousDatabaseDetailsOpenModeWrite,
}

// GetUpdateAutonomousDatabaseDetailsOpenModeEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsOpenModeEnum
func GetUpdateAutonomousDatabaseDetailsOpenModeEnumValues() []UpdateAutonomousDatabaseDetailsOpenModeEnum {
	values := make([]UpdateAutonomousDatabaseDetailsOpenModeEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsOpenMode {
		values = append(values, v)
	}
	return values
}

// UpdateAutonomousDatabaseDetailsPermissionLevelEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsPermissionLevelEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsPermissionLevelEnum
const (
	UpdateAutonomousDatabaseDetailsPermissionLevelRestricted   UpdateAutonomousDatabaseDetailsPermissionLevelEnum = "RESTRICTED"
	UpdateAutonomousDatabaseDetailsPermissionLevelUnrestricted UpdateAutonomousDatabaseDetailsPermissionLevelEnum = "UNRESTRICTED"
)

var mappingUpdateAutonomousDatabaseDetailsPermissionLevel = map[string]UpdateAutonomousDatabaseDetailsPermissionLevelEnum{
	"RESTRICTED":   UpdateAutonomousDatabaseDetailsPermissionLevelRestricted,
	"UNRESTRICTED": UpdateAutonomousDatabaseDetailsPermissionLevelUnrestricted,
}

// GetUpdateAutonomousDatabaseDetailsPermissionLevelEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsPermissionLevelEnum
func GetUpdateAutonomousDatabaseDetailsPermissionLevelEnumValues() []UpdateAutonomousDatabaseDetailsPermissionLevelEnum {
	values := make([]UpdateAutonomousDatabaseDetailsPermissionLevelEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsPermissionLevel {
		values = append(values, v)
	}
	return values
}
