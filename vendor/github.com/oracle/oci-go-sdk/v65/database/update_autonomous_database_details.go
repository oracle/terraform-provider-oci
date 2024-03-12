// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAutonomousDatabaseDetails Details to update an Oracle Autonomous Database.
// **Notes**
// - To specify OCPU core count, you must use either `ocpuCount` or `cpuCoreCount`. You cannot use both parameters at the same time. For Autonomous Database Serverless instances, `ocpuCount` is not used.
// - To specify a storage allocation, you must use  either `dataStorageSizeInGBs` or `dataStorageSizeInTBs`.
// - See the individual parameter discriptions for more information on the OCPU and storage value parameters.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateAutonomousDatabaseDetails struct {

	// Retention period, in days, for long-term backups
	BackupRetentionPeriodInDays *int `mandatory:"false" json:"backupRetentionPeriodInDays"`

	// The compute model of the Autonomous Database. This is required if using the `computeCount` parameter. If using `cpuCoreCount` then it is an error to specify `computeModel` to a non-null value. ECPU compute model is the recommended model and OCPU compute model is legacy.
	ComputeModel UpdateAutonomousDatabaseDetailsComputeModelEnum `mandatory:"false" json:"computeModel,omitempty"`

	// The percentage of the System Global Area(SGA) assigned to In-Memory tables in Autonomous Database. This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform.
	InMemoryPercentage *int `mandatory:"false" json:"inMemoryPercentage"`

	// Parameter that allows users to select an acceptable maximum data loss limit in seconds, up to which Automatic Failover will be triggered when necessary for a Local Autonomous Data Guard
	LocalAdgAutoFailoverMaxDataLossLimit *int `mandatory:"false" json:"localAdgAutoFailoverMaxDataLossLimit"`

	// The number of CPUs to be made available to the Autonomous Database.<br>
	// For Autonomous Databases on Dedicated Exadata Infrastructure:
	// - The CPU type (OCPUs or ECPUs) is determined by the parent Autonomous Exadata VM Cluster's compute model.
	// - It is suggested to use 'computeCount' parameter if you want to use fractional value to provision less than 1 core.
	// **Note:** This parameter cannot be used with the `ocpuCount` or `computeCount` parameter.
	// This cannot be updated in parallel with any of the following: licenseModel, databaseEdition, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	LongTermBackupSchedule *LongTermBackUpScheduleDetails `mandatory:"false" json:"longTermBackupSchedule"`

	// This project introduces Autonomous Database for Developers (ADB-Dev), a free tier on dedicated infrastructure, and Cloud@Customer for database development purposes. ADB-Dev enables ExaDB customers to experiment with ADB for free and incentivizes enterprises to use ADB for new development projects.Note that ADB-Dev have 4 CPU and 20GB of memory. For ADB-Dev , memory and CPU cannot be scaled
	IsDevTier *bool `mandatory:"false" json:"isDevTier"`

	// The compute amount (CPUs) available to the database. Minimum and maximum values depend on the compute model and whether the database is an Autonomous Database Serverless instance or an Autonomous Database on Dedicated Exadata Infrastructure.
	// For an Autonomous Database Serverless instance, the 'ECPU' compute model requires a minimum value of one, for databases in the elastic resource pool and minimum value of two, otherwise. Required when using the `computeModel` parameter. When using `cpuCoreCount` parameter, it is an error to specify computeCount to a non-null value. Providing `computeModel` and `computeCount` is the preferred method for both OCPU and ECPU.
	// This cannot be updated in parallel with any of the following: licenseModel, databaseEdition, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	ComputeCount *float32 `mandatory:"false" json:"computeCount"`

	// The number of OCPU cores to be made available to the Autonomous Database.
	// For Autonomous Databases on Dedicated Exadata Infrastructure, you can specify a fractional value for this parameter. Fractional values are not supported for Autonomous Database Serverless instances.
	// For Autonomous Database Serverless instances, this parameter is not used.
	// To provision less than 1 core, enter a fractional value in an increment of 0.1. To provision 1 or more cores, you must enter an integer between 1 and the maximum number of cores available to the infrastructure shape. For example, you can provision 0.3 or 0.4 cores, but not 0.35 cores. Likewise, you can provision 2 cores or 3 cores, but not 2.5 cores. The maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbde/index.html) for shape details.
	// **Note:** This parameter cannot be used with the `cpuCoreCount` parameter.
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// The size, in terabytes, of the data volume that will be created and attached to the database. For Autonomous Databases on dedicated Exadata infrastructure, the maximum storage value is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// A full Exadata service is allocated when the Autonomous Database size is set to the upper limit (384 TB).
	// **Note:** This parameter cannot be used with the `dataStorageSizeInGBs` parameter.
	// This cannot be updated in parallel with any of the following: licenseModel, databaseEdition, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	DataStorageSizeInTBs *int `mandatory:"false" json:"dataStorageSizeInTBs"`

	// Applies to dedicated Exadata infrastructure only.
	// The size, in gigabytes, of the data volume that will be created and attached to the database. The maximum storage value depends on the system shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `dataStorageSizeInTBs` parameter.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The user-friendly name for the Autonomous Database. The name does not have to be unique. The display name can only be updated for Autonomous Databases using dedicated Exadata Infrastructure. This parameter may not be updated in parallel with dbVersion.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isLocalDataGuardEnabled
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing. It must be different from the last four passwords and it must not be a password used within the last 24 hours.
	// This cannot be used in conjunction with with OCI vault secrets (secretId).
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, whitelistedIps, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, or isFreeTier.
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

	// New name for this Autonomous Database.
	// For Autonomous Databases on Dedicated Exadata Infrastructure, the name must begin with an alphabetic character, and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
	// For Autonomous Database Serverless instances, the name must begin with an alphabetic character, and can contain a maximum of 30 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails.
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
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	DbWorkload UpdateAutonomousDatabaseDetailsDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Oracle Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null. It is already set at the
	// Autonomous Exadata Infrastructure level. When provisioning an Autonomous Database Serverless  (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) database, if a value is not specified, the system defaults the value to `BRING_YOUR_OWN_LICENSE`. Bring your own license (BYOL) also allows you to select the DB edition using the optional parameter.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, dataStorageSizeInTBs, adminPassword, isMTLSConnectionRequired, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	LicenseModel UpdateAutonomousDatabaseDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Indicates if the database-level access control is enabled.
	// If disabled, database access is defined by the network security rules.
	// If enabled, database access is restricted to the IP addresses defined by the rules specified with the `whitelistedIps` property. While specifying `whitelistedIps` rules is optional,
	//  if database-level access control is enabled and no rules are specified, the database will become inaccessible. The rules can be added later using the `UpdateAutonomousDatabase` API operation or edit option in console.
	// When creating a database clone, the desired access control setting should be specified. By default, database-level access control will be disabled for the clone.
	// This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform. For Autonomous Database Serverless instances, `whitelistedIps` is used.
	IsAccessControlEnabled *bool `mandatory:"false" json:"isAccessControlEnabled"`

	// The client IP access control list (ACL). This feature is available for Autonomous Database Serverless  (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) and on Exadata Cloud@Customer.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.
	// If `arePrimaryWhitelistedIpsUsed` is 'TRUE' then Autonomous Database uses this primary's IP access control list (ACL) for the disaster recovery peer called `standbywhitelistedips`.
	// For Autonomous Database Serverless, this is an array of CIDR (classless inter-domain routing) notations for a subnet or VCN OCID (virtual cloud network Oracle Cloud ID).
	// Multiple IPs and VCN OCIDs should be separate strings separated by commas, but if it’s other configurations that need multiple pieces of information then its each piece is connected with semicolon (;) as a delimiter.
	// Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]`
	// For Exadata Cloud@Customer, this is an array of IP addresses or CIDR notations.
	// Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`
	// For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	WhitelistedIps []string `mandatory:"false" json:"whitelistedIps"`

	// This field will be null if the Autonomous Database is not Data Guard enabled or Access Control is disabled.
	// `TRUE` if the Autonomous Database has Data Guard and Access Control enabled, and the Autonomous Database uses the primary's IP access control list (ACL) for standby.
	// `FALSE` if the Autonomous Database has Data Guard and Access Control enabled, and the Autonomous Database uses a different IP access control list (ACL) for standby compared to primary.
	ArePrimaryWhitelistedIpsUsed *bool `mandatory:"false" json:"arePrimaryWhitelistedIpsUsed"`

	// The client IP access control list (ACL). This feature is available for Autonomous Database Serverless  (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) and on Exadata Cloud@Customer.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.
	// If `arePrimaryWhitelistedIpsUsed` is 'TRUE' then Autonomous Database uses this primary's IP access control list (ACL) for the disaster recovery peer called `standbywhitelistedips`.
	// For Autonomous Database Serverless, this is an array of CIDR (classless inter-domain routing) notations for a subnet or VCN OCID (virtual cloud network Oracle Cloud ID).
	// Multiple IPs and VCN OCIDs should be separate strings separated by commas, but if it’s other configurations that need multiple pieces of information then its each piece is connected with semicolon (;) as a delimiter.
	// Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]`
	// For Exadata Cloud@Customer, this is an array of IP addresses or CIDR notations.
	// Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`
	// For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	StandbyWhitelistedIps []string `mandatory:"false" json:"standbyWhitelistedIps"`

	// Indicates whether auto scaling is enabled for the Autonomous Database CPU core count. Setting to `TRUE` enables auto scaling. Setting to `FALSE` disables auto scaling. The default value is `TRUE`. Auto scaling is only available for Autonomous Database Serverless instances (https://docs.oracle.com/en/cloud/paas/autonomous-database/serverless/).
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// Indicates if the Autonomous Database is a refreshable clone.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	IsRefreshableClone *bool `mandatory:"false" json:"isRefreshableClone"`

	// The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.
	RefreshableMode UpdateAutonomousDatabaseDetailsRefreshableModeEnum `mandatory:"false" json:"refreshableMode,omitempty"`

	// Indicates whether the Autonomous Database has a local (in-region) standby database. Not applicable when creating a cross-region Autonomous Data Guard associations, or to
	// Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
	// To create a local standby, set to `TRUE`. To delete a local standby, set to `FALSE`. For more information on using Autonomous Data Guard on an Autonomous Database Serverless instance (local and cross-region) , see About Standby Databases (https://docs.oracle.com/en/cloud/paas/autonomous-database/adbsa/autonomous-data-guard-about.html#GUID-045AD017-8120-4BDC-AF58-7430FFE28D2B)
	// To enable cross-region Autonomous Data Guard on an Autonomous Database Serverless instance, see Enable Autonomous Data Guard (https://docs-uat.us.oracle.com/en/cloud/paas/autonomous-database/adbsa/autonomous-data-guard-update-type.html#GUID-967ED737-4A05-4D6E-A7CA-C3F21ACF9BF0).
	// This cannot be updated in parallel with any of the following: isMTLSRequired, dbWorkload, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	IsLocalDataGuardEnabled *bool `mandatory:"false" json:"isLocalDataGuardEnabled"`

	// ** Deprecated. ** Indicates whether the Autonomous Database has a local (in-region) standby database. Not applicable when creating a cross-region Autonomous Data Guard associations, or to
	// Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
	// To create a local standby, set to `TRUE`. To delete a local standby, set to `FALSE`. For more information on using Autonomous Data Guard on an Autonomous Database Serverless instance (local and cross-region) , see About Standby Databases (https://docs.oracle.com/en/cloud/paas/autonomous-database/adbsa/autonomous-data-guard-about.html#GUID-045AD017-8120-4BDC-AF58-7430FFE28D2B)
	// To enable cross-region Autonomous Data Guard on an Autonomous Database Serverless instance, see Enable Autonomous Data Guard (https://docs-uat.us.oracle.com/en/cloud/paas/autonomous-database/adbsa/autonomous-data-guard-update-type.html#GUID-967ED737-4A05-4D6E-A7CA-C3F21ACF9BF0).
	// To delete a cross-region standby database, provide the `peerDbId` for the standby database in a remote region, and set `isDataGuardEnabled` to `FALSE`.
	IsDataGuardEnabled *bool `mandatory:"false" json:"isDataGuardEnabled"`

	// The database OCID(/Content/General/Concepts/identifiers.htm) of the Disaster Recovery peer (source Primary) database, which is located in a different (remote) region from the current peer database.
	// To create or delete a local (in-region) standby, see the `isDataGuardEnabled` parameter.
	PeerDbId *string `mandatory:"false" json:"peerDbId"`

	// A valid Oracle Database version for Autonomous Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// Indicates the Autonomous Database mode. The database can be opened in `READ_ONLY` or `READ_WRITE` mode.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	OpenMode UpdateAutonomousDatabaseDetailsOpenModeEnum `mandatory:"false" json:"openMode,omitempty"`

	// The Autonomous Database permission level. Restricted mode allows access only by admin users.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
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

	// The resource's private endpoint label.
	// - Setting the endpoint label to a non-empty string creates a private endpoint database.
	// - Resetting the endpoint label to an empty string, after the creation of the private endpoint database, changes the private endpoint database to a public endpoint database.
	// - Setting the endpoint label to a non-empty string value, updates to a new private endpoint database, when the database is disabled and re-enabled.
	// This setting cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, dbWorkload, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	PrivateEndpointLabel *string `mandatory:"false" json:"privateEndpointLabel"`

	// The private endpoint Ip address for the resource.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Customer Contacts. Setting this to an empty list removes all customer contacts of an Oracle
	// This cannot be updated in parallel with any of the following: isMTLSConnectionRequired, scheduledOperations, or dbToolsDetails.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	// Specifies if the Autonomous Database requires mTLS connections.
	// This may not be updated in parallel with any of the following: licenseModel, databaseEdition, cpuCoreCount, computeCount, dataStorageSizeInTBs, whitelistedIps, openMode, permissionLevel, db-workload, privateEndpointLabel, nsgIds, customerContacts, dbVersion, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	// Service Change: The default value of the isMTLSConnectionRequired attribute will change from true to false on July 1, 2023 in the following APIs:
	// - CreateAutonomousDatabase
	// - GetAutonomousDatabase
	// - UpdateAutonomousDatabase
	// Details: Prior to the July 1, 2023 change, the isMTLSConnectionRequired attribute default value was true. This applies to Autonomous Database Serverless.
	// Does this impact me? If you use or maintain custom scripts or Terraform scripts referencing the CreateAutonomousDatabase, GetAutonomousDatabase, or UpdateAutonomousDatabase APIs, you want to check, and possibly modify, the scripts for the changed default value of the attribute. Should you choose not to leave your scripts unchanged, the API calls containing this attribute will continue to work, but the default value will switch from true to false.
	// How do I make this change? Using either OCI SDKs or command line tools, update your custom scripts to explicitly set the isMTLSConnectionRequired attribute to true.
	IsMtlsConnectionRequired *bool `mandatory:"false" json:"isMtlsConnectionRequired"`

	// The unique identifier for leader autonomous database OCID OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ResourcePoolLeaderId *string `mandatory:"false" json:"resourcePoolLeaderId"`

	ResourcePoolSummary *ResourcePoolSummary `mandatory:"false" json:"resourcePoolSummary"`

	// The list of scheduled operations. Consists of values such as dayOfWeek, scheduledStartTime, scheduledStopTime.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	ScheduledOperations []ScheduledOperationDetails `mandatory:"false" json:"scheduledOperations"`

	// Indicates if auto scaling is enabled for the Autonomous Database storage. The default value is `FALSE`.
	IsAutoScalingForStorageEnabled *bool `mandatory:"false" json:"isAutoScalingForStorageEnabled"`

	// The Oracle Database Edition that applies to the Autonomous databases.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	DatabaseEdition AutonomousDatabaseSummaryDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

	// The list of database tools details.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, isLocalDataGuardEnabled, or isFreeTier.
	DbToolsDetails []DatabaseTool `mandatory:"false" json:"dbToolsDetails"`

	// The OCI vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
	SecretId *string `mandatory:"false" json:"secretId"`

	// The version of the vault secret. If no version is specified, the latest version will be used.
	SecretVersionNumber *int `mandatory:"false" json:"secretVersionNumber"`
}

func (m UpdateAutonomousDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAutonomousDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateAutonomousDatabaseDetailsComputeModelEnum(string(m.ComputeModel)); !ok && m.ComputeModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeModel: %s. Supported values are: %s.", m.ComputeModel, strings.Join(GetUpdateAutonomousDatabaseDetailsComputeModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateAutonomousDatabaseDetailsDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetUpdateAutonomousDatabaseDetailsDbWorkloadEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateAutonomousDatabaseDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetUpdateAutonomousDatabaseDetailsLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateAutonomousDatabaseDetailsRefreshableModeEnum(string(m.RefreshableMode)); !ok && m.RefreshableMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RefreshableMode: %s. Supported values are: %s.", m.RefreshableMode, strings.Join(GetUpdateAutonomousDatabaseDetailsRefreshableModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateAutonomousDatabaseDetailsOpenModeEnum(string(m.OpenMode)); !ok && m.OpenMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OpenMode: %s. Supported values are: %s.", m.OpenMode, strings.Join(GetUpdateAutonomousDatabaseDetailsOpenModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateAutonomousDatabaseDetailsPermissionLevelEnum(string(m.PermissionLevel)); !ok && m.PermissionLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PermissionLevel: %s. Supported values are: %s.", m.PermissionLevel, strings.Join(GetUpdateAutonomousDatabaseDetailsPermissionLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetAutonomousDatabaseSummaryDatabaseEditionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateAutonomousDatabaseDetailsComputeModelEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsComputeModelEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsComputeModelEnum
const (
	UpdateAutonomousDatabaseDetailsComputeModelEcpu UpdateAutonomousDatabaseDetailsComputeModelEnum = "ECPU"
	UpdateAutonomousDatabaseDetailsComputeModelOcpu UpdateAutonomousDatabaseDetailsComputeModelEnum = "OCPU"
)

var mappingUpdateAutonomousDatabaseDetailsComputeModelEnum = map[string]UpdateAutonomousDatabaseDetailsComputeModelEnum{
	"ECPU": UpdateAutonomousDatabaseDetailsComputeModelEcpu,
	"OCPU": UpdateAutonomousDatabaseDetailsComputeModelOcpu,
}

var mappingUpdateAutonomousDatabaseDetailsComputeModelEnumLowerCase = map[string]UpdateAutonomousDatabaseDetailsComputeModelEnum{
	"ecpu": UpdateAutonomousDatabaseDetailsComputeModelEcpu,
	"ocpu": UpdateAutonomousDatabaseDetailsComputeModelOcpu,
}

// GetUpdateAutonomousDatabaseDetailsComputeModelEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsComputeModelEnum
func GetUpdateAutonomousDatabaseDetailsComputeModelEnumValues() []UpdateAutonomousDatabaseDetailsComputeModelEnum {
	values := make([]UpdateAutonomousDatabaseDetailsComputeModelEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsComputeModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousDatabaseDetailsComputeModelEnumStringValues Enumerates the set of values in String for UpdateAutonomousDatabaseDetailsComputeModelEnum
func GetUpdateAutonomousDatabaseDetailsComputeModelEnumStringValues() []string {
	return []string{
		"ECPU",
		"OCPU",
	}
}

// GetMappingUpdateAutonomousDatabaseDetailsComputeModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAutonomousDatabaseDetailsComputeModelEnum(val string) (UpdateAutonomousDatabaseDetailsComputeModelEnum, bool) {
	enum, ok := mappingUpdateAutonomousDatabaseDetailsComputeModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingUpdateAutonomousDatabaseDetailsDbWorkloadEnum = map[string]UpdateAutonomousDatabaseDetailsDbWorkloadEnum{
	"OLTP": UpdateAutonomousDatabaseDetailsDbWorkloadOltp,
	"DW":   UpdateAutonomousDatabaseDetailsDbWorkloadDw,
	"AJD":  UpdateAutonomousDatabaseDetailsDbWorkloadAjd,
	"APEX": UpdateAutonomousDatabaseDetailsDbWorkloadApex,
}

var mappingUpdateAutonomousDatabaseDetailsDbWorkloadEnumLowerCase = map[string]UpdateAutonomousDatabaseDetailsDbWorkloadEnum{
	"oltp": UpdateAutonomousDatabaseDetailsDbWorkloadOltp,
	"dw":   UpdateAutonomousDatabaseDetailsDbWorkloadDw,
	"ajd":  UpdateAutonomousDatabaseDetailsDbWorkloadAjd,
	"apex": UpdateAutonomousDatabaseDetailsDbWorkloadApex,
}

// GetUpdateAutonomousDatabaseDetailsDbWorkloadEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsDbWorkloadEnum
func GetUpdateAutonomousDatabaseDetailsDbWorkloadEnumValues() []UpdateAutonomousDatabaseDetailsDbWorkloadEnum {
	values := make([]UpdateAutonomousDatabaseDetailsDbWorkloadEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousDatabaseDetailsDbWorkloadEnumStringValues Enumerates the set of values in String for UpdateAutonomousDatabaseDetailsDbWorkloadEnum
func GetUpdateAutonomousDatabaseDetailsDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
		"AJD",
		"APEX",
	}
}

// GetMappingUpdateAutonomousDatabaseDetailsDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAutonomousDatabaseDetailsDbWorkloadEnum(val string) (UpdateAutonomousDatabaseDetailsDbWorkloadEnum, bool) {
	enum, ok := mappingUpdateAutonomousDatabaseDetailsDbWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateAutonomousDatabaseDetailsLicenseModelEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsLicenseModelEnum
const (
	UpdateAutonomousDatabaseDetailsLicenseModelLicenseIncluded     UpdateAutonomousDatabaseDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateAutonomousDatabaseDetailsLicenseModelBringYourOwnLicense UpdateAutonomousDatabaseDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateAutonomousDatabaseDetailsLicenseModelEnum = map[string]UpdateAutonomousDatabaseDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateAutonomousDatabaseDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateAutonomousDatabaseDetailsLicenseModelBringYourOwnLicense,
}

var mappingUpdateAutonomousDatabaseDetailsLicenseModelEnumLowerCase = map[string]UpdateAutonomousDatabaseDetailsLicenseModelEnum{
	"license_included":       UpdateAutonomousDatabaseDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": UpdateAutonomousDatabaseDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateAutonomousDatabaseDetailsLicenseModelEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsLicenseModelEnum
func GetUpdateAutonomousDatabaseDetailsLicenseModelEnumValues() []UpdateAutonomousDatabaseDetailsLicenseModelEnum {
	values := make([]UpdateAutonomousDatabaseDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousDatabaseDetailsLicenseModelEnumStringValues Enumerates the set of values in String for UpdateAutonomousDatabaseDetailsLicenseModelEnum
func GetUpdateAutonomousDatabaseDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingUpdateAutonomousDatabaseDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAutonomousDatabaseDetailsLicenseModelEnum(val string) (UpdateAutonomousDatabaseDetailsLicenseModelEnum, bool) {
	enum, ok := mappingUpdateAutonomousDatabaseDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateAutonomousDatabaseDetailsRefreshableModeEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsRefreshableModeEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsRefreshableModeEnum
const (
	UpdateAutonomousDatabaseDetailsRefreshableModeAutomatic UpdateAutonomousDatabaseDetailsRefreshableModeEnum = "AUTOMATIC"
	UpdateAutonomousDatabaseDetailsRefreshableModeManual    UpdateAutonomousDatabaseDetailsRefreshableModeEnum = "MANUAL"
)

var mappingUpdateAutonomousDatabaseDetailsRefreshableModeEnum = map[string]UpdateAutonomousDatabaseDetailsRefreshableModeEnum{
	"AUTOMATIC": UpdateAutonomousDatabaseDetailsRefreshableModeAutomatic,
	"MANUAL":    UpdateAutonomousDatabaseDetailsRefreshableModeManual,
}

var mappingUpdateAutonomousDatabaseDetailsRefreshableModeEnumLowerCase = map[string]UpdateAutonomousDatabaseDetailsRefreshableModeEnum{
	"automatic": UpdateAutonomousDatabaseDetailsRefreshableModeAutomatic,
	"manual":    UpdateAutonomousDatabaseDetailsRefreshableModeManual,
}

// GetUpdateAutonomousDatabaseDetailsRefreshableModeEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsRefreshableModeEnum
func GetUpdateAutonomousDatabaseDetailsRefreshableModeEnumValues() []UpdateAutonomousDatabaseDetailsRefreshableModeEnum {
	values := make([]UpdateAutonomousDatabaseDetailsRefreshableModeEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsRefreshableModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousDatabaseDetailsRefreshableModeEnumStringValues Enumerates the set of values in String for UpdateAutonomousDatabaseDetailsRefreshableModeEnum
func GetUpdateAutonomousDatabaseDetailsRefreshableModeEnumStringValues() []string {
	return []string{
		"AUTOMATIC",
		"MANUAL",
	}
}

// GetMappingUpdateAutonomousDatabaseDetailsRefreshableModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAutonomousDatabaseDetailsRefreshableModeEnum(val string) (UpdateAutonomousDatabaseDetailsRefreshableModeEnum, bool) {
	enum, ok := mappingUpdateAutonomousDatabaseDetailsRefreshableModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateAutonomousDatabaseDetailsOpenModeEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsOpenModeEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsOpenModeEnum
const (
	UpdateAutonomousDatabaseDetailsOpenModeOnly  UpdateAutonomousDatabaseDetailsOpenModeEnum = "READ_ONLY"
	UpdateAutonomousDatabaseDetailsOpenModeWrite UpdateAutonomousDatabaseDetailsOpenModeEnum = "READ_WRITE"
)

var mappingUpdateAutonomousDatabaseDetailsOpenModeEnum = map[string]UpdateAutonomousDatabaseDetailsOpenModeEnum{
	"READ_ONLY":  UpdateAutonomousDatabaseDetailsOpenModeOnly,
	"READ_WRITE": UpdateAutonomousDatabaseDetailsOpenModeWrite,
}

var mappingUpdateAutonomousDatabaseDetailsOpenModeEnumLowerCase = map[string]UpdateAutonomousDatabaseDetailsOpenModeEnum{
	"read_only":  UpdateAutonomousDatabaseDetailsOpenModeOnly,
	"read_write": UpdateAutonomousDatabaseDetailsOpenModeWrite,
}

// GetUpdateAutonomousDatabaseDetailsOpenModeEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsOpenModeEnum
func GetUpdateAutonomousDatabaseDetailsOpenModeEnumValues() []UpdateAutonomousDatabaseDetailsOpenModeEnum {
	values := make([]UpdateAutonomousDatabaseDetailsOpenModeEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsOpenModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousDatabaseDetailsOpenModeEnumStringValues Enumerates the set of values in String for UpdateAutonomousDatabaseDetailsOpenModeEnum
func GetUpdateAutonomousDatabaseDetailsOpenModeEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
	}
}

// GetMappingUpdateAutonomousDatabaseDetailsOpenModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAutonomousDatabaseDetailsOpenModeEnum(val string) (UpdateAutonomousDatabaseDetailsOpenModeEnum, bool) {
	enum, ok := mappingUpdateAutonomousDatabaseDetailsOpenModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateAutonomousDatabaseDetailsPermissionLevelEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsPermissionLevelEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsPermissionLevelEnum
const (
	UpdateAutonomousDatabaseDetailsPermissionLevelRestricted   UpdateAutonomousDatabaseDetailsPermissionLevelEnum = "RESTRICTED"
	UpdateAutonomousDatabaseDetailsPermissionLevelUnrestricted UpdateAutonomousDatabaseDetailsPermissionLevelEnum = "UNRESTRICTED"
)

var mappingUpdateAutonomousDatabaseDetailsPermissionLevelEnum = map[string]UpdateAutonomousDatabaseDetailsPermissionLevelEnum{
	"RESTRICTED":   UpdateAutonomousDatabaseDetailsPermissionLevelRestricted,
	"UNRESTRICTED": UpdateAutonomousDatabaseDetailsPermissionLevelUnrestricted,
}

var mappingUpdateAutonomousDatabaseDetailsPermissionLevelEnumLowerCase = map[string]UpdateAutonomousDatabaseDetailsPermissionLevelEnum{
	"restricted":   UpdateAutonomousDatabaseDetailsPermissionLevelRestricted,
	"unrestricted": UpdateAutonomousDatabaseDetailsPermissionLevelUnrestricted,
}

// GetUpdateAutonomousDatabaseDetailsPermissionLevelEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsPermissionLevelEnum
func GetUpdateAutonomousDatabaseDetailsPermissionLevelEnumValues() []UpdateAutonomousDatabaseDetailsPermissionLevelEnum {
	values := make([]UpdateAutonomousDatabaseDetailsPermissionLevelEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsPermissionLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousDatabaseDetailsPermissionLevelEnumStringValues Enumerates the set of values in String for UpdateAutonomousDatabaseDetailsPermissionLevelEnum
func GetUpdateAutonomousDatabaseDetailsPermissionLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"UNRESTRICTED",
	}
}

// GetMappingUpdateAutonomousDatabaseDetailsPermissionLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAutonomousDatabaseDetailsPermissionLevelEnum(val string) (UpdateAutonomousDatabaseDetailsPermissionLevelEnum, bool) {
	enum, ok := mappingUpdateAutonomousDatabaseDetailsPermissionLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
