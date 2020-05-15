// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousDatabaseSummary An Oracle Autonomous Database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type AutonomousDatabaseSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the Autonomous Database.
	LifecycleState AutonomousDatabaseSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The database name.
	DbName *string `mandatory:"true" json:"dbName"`

	// The number of OCPU cores to be made available to the database.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The quantity of data in the database, in terabytes.
	DataStorageSizeInTBs *int `mandatory:"true" json:"dataStorageSizeInTBs"`

	// Information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled.
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The date and time the Always Free database will be stopped because of inactivity. If this time is reached without any database activity, the database will automatically be put into the STOPPED state.
	TimeReclamationOfFreeAutonomousDatabase *common.SDKTime `mandatory:"false" json:"timeReclamationOfFreeAutonomousDatabase"`

	// The date and time the Always Free database will be automatically deleted because of inactivity. If the database is in the STOPPED state and without activity until this time, it will be deleted.
	TimeDeletionOfFreeAutonomousDatabase *common.SDKTime `mandatory:"false" json:"timeDeletionOfFreeAutonomousDatabase"`

	// True if the database uses dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm).
	IsDedicated *bool `mandatory:"false" json:"isDedicated"`

	// The Autonomous Container Database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	AutonomousContainerDatabaseId *string `mandatory:"false" json:"autonomousContainerDatabaseId"`

	// The date and time the Autonomous Database was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The user-friendly name for the Autonomous Database. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The URL of the Service Console for the Autonomous Database.
	ServiceConsoleUrl *string `mandatory:"false" json:"serviceConsoleUrl"`

	// The connection string used to connect to the Autonomous Database. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Database for the password value.
	ConnectionStrings *AutonomousDatabaseConnectionStrings `mandatory:"false" json:"connectionStrings"`

	ConnectionUrls *AutonomousDatabaseConnectionUrls `mandatory:"false" json:"connectionUrls"`

	// The Oracle license model that applies to the Oracle Autonomous Database. Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm), this attribute must be null because the attribute is already set at the
	// Autonomous Exadata Infrastructure level. When using shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`.
	LicenseModel AutonomousDatabaseSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The amount of storage that has been used, in terabytes.
	UsedDataStorageSizeInTBs *int `mandatory:"false" json:"usedDataStorageSizeInTBs"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// - For Autonomous Database, setting this will disable public secure access to the database.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and the backup subnet.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The private endpoint for the resource.
	PrivateEndpoint *string `mandatory:"false" json:"privateEndpoint"`

	// The private endpoint label for the resource.
	PrivateEndpointLabel *string `mandatory:"false" json:"privateEndpointLabel"`

	// The private endpoint Ip address for the resource.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// A valid Oracle Database version for Autonomous Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// Indicates if the Autonomous Database version is a preview version.
	IsPreview *bool `mandatory:"false" json:"isPreview"`

	// The Autonomous Database workload type. The following values are valid:
	// - OLTP - indicates an Autonomous Transaction Processing database
	// - DW - indicates an Autonomous Data Warehouse database
	DbWorkload AutonomousDatabaseSummaryDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// The client IP access control list (ACL). This feature is available for databases on shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) only.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. This is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID.
	// To add the whitelist VCN specific subnet or IP, use a semicoln ';' as a deliminator to add the VCN specific subnets or IPs.
	// Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw;1.1.1.1","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw;1.1.0.0/16"]`
	WhitelistedIps []string `mandatory:"false" json:"whitelistedIps"`

	// Indicates if auto scaling is enabled for the Autonomous Database CPU core count. Note that auto scaling is available for databases on shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) only.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// Status of the Data Safe registration for this Autonomous Database.
	DataSafeStatus AutonomousDatabaseSummaryDataSafeStatusEnum `mandatory:"false" json:"dataSafeStatus,omitempty"`

	// The date and time when maintenance will begin.
	TimeMaintenanceBegin *common.SDKTime `mandatory:"false" json:"timeMaintenanceBegin"`

	// The date and time when maintenance will end.
	TimeMaintenanceEnd *common.SDKTime `mandatory:"false" json:"timeMaintenanceEnd"`

	// List of Oracle Database versions available for a database upgrade. If there are no version upgrades available, this list is empty.
	AvailableUpgradeVersions []string `mandatory:"false" json:"availableUpgradeVersions"`
}

func (m AutonomousDatabaseSummary) String() string {
	return common.PointerString(m)
}

// AutonomousDatabaseSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryLifecycleStateEnum
const (
	AutonomousDatabaseSummaryLifecycleStateProvisioning            AutonomousDatabaseSummaryLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseSummaryLifecycleStateAvailable               AutonomousDatabaseSummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseSummaryLifecycleStateStopping                AutonomousDatabaseSummaryLifecycleStateEnum = "STOPPING"
	AutonomousDatabaseSummaryLifecycleStateStopped                 AutonomousDatabaseSummaryLifecycleStateEnum = "STOPPED"
	AutonomousDatabaseSummaryLifecycleStateStarting                AutonomousDatabaseSummaryLifecycleStateEnum = "STARTING"
	AutonomousDatabaseSummaryLifecycleStateTerminating             AutonomousDatabaseSummaryLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseSummaryLifecycleStateTerminated              AutonomousDatabaseSummaryLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseSummaryLifecycleStateUnavailable             AutonomousDatabaseSummaryLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDatabaseSummaryLifecycleStateRestoreInProgress       AutonomousDatabaseSummaryLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDatabaseSummaryLifecycleStateRestoreFailed           AutonomousDatabaseSummaryLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousDatabaseSummaryLifecycleStateBackupInProgress        AutonomousDatabaseSummaryLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDatabaseSummaryLifecycleStateScaleInProgress         AutonomousDatabaseSummaryLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDatabaseSummaryLifecycleStateAvailableNeedsAttention AutonomousDatabaseSummaryLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDatabaseSummaryLifecycleStateUpdating                AutonomousDatabaseSummaryLifecycleStateEnum = "UPDATING"
	AutonomousDatabaseSummaryLifecycleStateMaintenanceInProgress   AutonomousDatabaseSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	AutonomousDatabaseSummaryLifecycleStateRestarting              AutonomousDatabaseSummaryLifecycleStateEnum = "RESTARTING"
	AutonomousDatabaseSummaryLifecycleStateUpgrading               AutonomousDatabaseSummaryLifecycleStateEnum = "UPGRADING"
)

var mappingAutonomousDatabaseSummaryLifecycleState = map[string]AutonomousDatabaseSummaryLifecycleStateEnum{
	"PROVISIONING":              AutonomousDatabaseSummaryLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDatabaseSummaryLifecycleStateAvailable,
	"STOPPING":                  AutonomousDatabaseSummaryLifecycleStateStopping,
	"STOPPED":                   AutonomousDatabaseSummaryLifecycleStateStopped,
	"STARTING":                  AutonomousDatabaseSummaryLifecycleStateStarting,
	"TERMINATING":               AutonomousDatabaseSummaryLifecycleStateTerminating,
	"TERMINATED":                AutonomousDatabaseSummaryLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDatabaseSummaryLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDatabaseSummaryLifecycleStateRestoreInProgress,
	"RESTORE_FAILED":            AutonomousDatabaseSummaryLifecycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":        AutonomousDatabaseSummaryLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDatabaseSummaryLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDatabaseSummaryLifecycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDatabaseSummaryLifecycleStateUpdating,
	"MAINTENANCE_IN_PROGRESS":   AutonomousDatabaseSummaryLifecycleStateMaintenanceInProgress,
	"RESTARTING":                AutonomousDatabaseSummaryLifecycleStateRestarting,
	"UPGRADING":                 AutonomousDatabaseSummaryLifecycleStateUpgrading,
}

// GetAutonomousDatabaseSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseSummaryLifecycleStateEnum
func GetAutonomousDatabaseSummaryLifecycleStateEnumValues() []AutonomousDatabaseSummaryLifecycleStateEnum {
	values := make([]AutonomousDatabaseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryLicenseModelEnum Enum with underlying type: string
type AutonomousDatabaseSummaryLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryLicenseModelEnum
const (
	AutonomousDatabaseSummaryLicenseModelLicenseIncluded     AutonomousDatabaseSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousDatabaseSummaryLicenseModelBringYourOwnLicense AutonomousDatabaseSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousDatabaseSummaryLicenseModel = map[string]AutonomousDatabaseSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousDatabaseSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousDatabaseSummaryLicenseModelBringYourOwnLicense,
}

// GetAutonomousDatabaseSummaryLicenseModelEnumValues Enumerates the set of values for AutonomousDatabaseSummaryLicenseModelEnum
func GetAutonomousDatabaseSummaryLicenseModelEnumValues() []AutonomousDatabaseSummaryLicenseModelEnum {
	values := make([]AutonomousDatabaseSummaryLicenseModelEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryLicenseModel {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryDbWorkloadEnum Enum with underlying type: string
type AutonomousDatabaseSummaryDbWorkloadEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryDbWorkloadEnum
const (
	AutonomousDatabaseSummaryDbWorkloadOltp AutonomousDatabaseSummaryDbWorkloadEnum = "OLTP"
	AutonomousDatabaseSummaryDbWorkloadDw   AutonomousDatabaseSummaryDbWorkloadEnum = "DW"
)

var mappingAutonomousDatabaseSummaryDbWorkload = map[string]AutonomousDatabaseSummaryDbWorkloadEnum{
	"OLTP": AutonomousDatabaseSummaryDbWorkloadOltp,
	"DW":   AutonomousDatabaseSummaryDbWorkloadDw,
}

// GetAutonomousDatabaseSummaryDbWorkloadEnumValues Enumerates the set of values for AutonomousDatabaseSummaryDbWorkloadEnum
func GetAutonomousDatabaseSummaryDbWorkloadEnumValues() []AutonomousDatabaseSummaryDbWorkloadEnum {
	values := make([]AutonomousDatabaseSummaryDbWorkloadEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryDbWorkload {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryDataSafeStatusEnum Enum with underlying type: string
type AutonomousDatabaseSummaryDataSafeStatusEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryDataSafeStatusEnum
const (
	AutonomousDatabaseSummaryDataSafeStatusRegistering   AutonomousDatabaseSummaryDataSafeStatusEnum = "REGISTERING"
	AutonomousDatabaseSummaryDataSafeStatusRegistered    AutonomousDatabaseSummaryDataSafeStatusEnum = "REGISTERED"
	AutonomousDatabaseSummaryDataSafeStatusDeregistering AutonomousDatabaseSummaryDataSafeStatusEnum = "DEREGISTERING"
	AutonomousDatabaseSummaryDataSafeStatusNotRegistered AutonomousDatabaseSummaryDataSafeStatusEnum = "NOT_REGISTERED"
	AutonomousDatabaseSummaryDataSafeStatusFailed        AutonomousDatabaseSummaryDataSafeStatusEnum = "FAILED"
)

var mappingAutonomousDatabaseSummaryDataSafeStatus = map[string]AutonomousDatabaseSummaryDataSafeStatusEnum{
	"REGISTERING":    AutonomousDatabaseSummaryDataSafeStatusRegistering,
	"REGISTERED":     AutonomousDatabaseSummaryDataSafeStatusRegistered,
	"DEREGISTERING":  AutonomousDatabaseSummaryDataSafeStatusDeregistering,
	"NOT_REGISTERED": AutonomousDatabaseSummaryDataSafeStatusNotRegistered,
	"FAILED":         AutonomousDatabaseSummaryDataSafeStatusFailed,
}

// GetAutonomousDatabaseSummaryDataSafeStatusEnumValues Enumerates the set of values for AutonomousDatabaseSummaryDataSafeStatusEnum
func GetAutonomousDatabaseSummaryDataSafeStatusEnumValues() []AutonomousDatabaseSummaryDataSafeStatusEnum {
	values := make([]AutonomousDatabaseSummaryDataSafeStatusEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryDataSafeStatus {
		values = append(values, v)
	}
	return values
}
