// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateAutonomousDatabaseBase Details to create an Oracle Autonomous Database.
// **Notes:**
// - To specify OCPU core count, you must use either `ocpuCount` or `cpuCoreCount`. You cannot use both parameters at the same time. For Autonomous Database Serverless instances, `ocpuCount` is not used.
// - To specify a storage allocation, you must use  either `dataStorageSizeInGBs` or `dataStorageSizeInTBs`.
// - See the individual parameter discriptions for more information on the OCPU and storage value parameters.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateAutonomousDatabaseBase interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the Autonomous Database.
	GetCompartmentId() *string

	// The character set for the autonomous database. The default is AL32UTF8. Allowed values for an Autonomous Database Serverless instance as as returned by List Autonomous Database Character Sets (https://docs.oracle.com/iaas/autonomous-database-serverless/doc/autonomous-character-set-selection.html)
	// For an Autonomous Database on dedicated infrastructure, the allowed values are:
	// AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSC5601, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P1, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS
	GetCharacterSet() *string

	// The character set for the Autonomous Database. The default is AL32UTF8. Use List Autonomous Database Character Sets (https://docs.oracle.com/iaas/autonomous-database-serverless/doc/autonomous-character-set-selection.html) to list the allowed values for an Autonomous Database Serverless instance.
	// For an Autonomous Database on dedicated Exadata infrastructure, the allowed values are:
	// AL16UTF16 or UTF8.
	GetNcharacterSet() *string

	// The database name. The name must begin with an alphabetic character and can contain a maximum of 30 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy. It is required in all cases except when creating a cross-region Autonomous Data Guard standby instance or a cross-region disaster recovery standby instance.
	GetDbName() *string

	// The number of CPU cores to be made available to the database. For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `ocpuCount` parameter.
	GetCpuCoreCount() *int

	// Retention period, in days, for long-term backups
	GetBackupRetentionPeriodInDays() *int

	// The compute model of the Autonomous Database. This is required if using the `computeCount` parameter. If using `cpuCoreCount` then it is an error to specify `computeModel` to a non-null value. ECPU compute model is the recommended model and OCPU compute model is legacy.
	GetComputeModel() CreateAutonomousDatabaseBaseComputeModelEnum

	// The compute amount (CPUs) available to the database. Minimum and maximum values depend on the compute model and whether the database is an Autonomous Database Serverless instance or an Autonomous Database on Dedicated Exadata Infrastructure.
	// For an Autonomous Database Serverless instance, the 'ECPU' compute model requires a minimum value of one, for databases in the elastic resource pool and minimum value of two, otherwise. Required when using the `computeModel` parameter. When using `cpuCoreCount` parameter, it is an error to specify computeCount to a non-null value. Providing `computeModel` and `computeCount` is the preferred method for both OCPU and ECPU.
	GetComputeCount() *float32

	// The number of OCPU cores to be made available to the database.
	// The following points apply:
	// - For Autonomous Databases on Dedicated Exadata infrastructure, to provision less than 1 core, enter a fractional value in an increment of 0.1. For example, you can provision 0.3 or 0.4 cores, but not 0.35 cores. (Note that fractional OCPU values are not supported for Autonomous Database Serverless instances.)
	// - To provision 1 or more cores, you must enter an integer between 1 and the maximum number of cores available for the infrastructure shape. For example, you can provision 2 cores or 3 cores, but not 2.5 cores. This applies to an Autonomous Database Serverless instance or an Autonomous Database on Dedicated Exadata Infrastructure.
	// - For Autonomous Database Serverless instances, this parameter is not used.
	// For Autonomous Databases on Dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `cpuCoreCount` parameter.
	GetOcpuCount() *float32

	// The Autonomous Database workload type. The following values are valid:
	// - OLTP - indicates an Autonomous Transaction Processing database
	// - DW - indicates an Autonomous Data Warehouse database
	// - AJD - indicates an Autonomous JSON Database
	// - APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	GetDbWorkload() CreateAutonomousDatabaseBaseDbWorkloadEnum

	// The size, in terabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed. For Autonomous Databases on dedicated Exadata infrastructure, the maximum storage value is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// A full Exadata service is allocated when the Autonomous Database size is set to the upper limit (384 TB).
	// **Note:** This parameter cannot be used with the `dataStorageSizeInGBs` parameter.
	GetDataStorageSizeInTBs() *int

	// The size, in gigabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed. The maximum storage value is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Notes**
	// - This parameter is only supported for dedicated Exadata infrastructure.
	// - This parameter cannot be used with the `dataStorageSizeInTBs` parameter.
	GetDataStorageSizeInGBs() *int

	// Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isLocalDataGuardEnabled
	GetIsFreeTier() *bool

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	GetKmsKeyId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	GetVaultId() *string

	// **Important** The `adminPassword` or `secretId` must be specified for all Autonomous Databases except for refreshable clones. The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing.
	// This cannot be used in conjunction with with OCI vault secrets (secretId).
	GetAdminPassword() *string

	// The user-friendly name for the Autonomous Database. The name does not have to be unique.
	GetDisplayName() *string

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Oracle Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null. It is already set at the
	// Autonomous Exadata Infrastructure level. When provisioning an Autonomous Database Serverless  (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) database, if a value is not specified, the system defaults the value to `BRING_YOUR_OWN_LICENSE`. Bring your own license (BYOL) also allows you to select the DB edition using the optional parameter.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, dataStorageSizeInTBs, adminPassword, isMTLSConnectionRequired, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	GetLicenseModel() CreateAutonomousDatabaseBaseLicenseModelEnum

	// If set to `TRUE`, indicates that an Autonomous Database preview version is being provisioned, and that the preview version's terms of service have been accepted. Note that preview version software is only available for Autonomous Database Serverless instances (https://docs.oracle.com/en/cloud/paas/autonomous-database/serverless/).
	GetIsPreviewVersionWithServiceTermsAccepted() *bool

	// Indicates if auto scaling is enabled for the Autonomous Database CPU core count. The default value is `TRUE`.
	GetIsAutoScalingEnabled() *bool

	// This project introduces Autonomous Database for Developers (ADB-Dev), a free tier on dedicated infrastructure, and Cloud@Customer for database development purposes. ADB-Dev enables ExaDB customers to experiment with ADB for free and incentivizes enterprises to use ADB for new development projects.Note that ADB-Dev have 4 CPU and 20GB of memory. For ADB-Dev , memory and CPU cannot be scaled
	GetIsDevTier() *bool

	// True if the database is on dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html).
	GetIsDedicated() *bool

	// The Autonomous Container Database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). Used only by Autonomous Database on Dedicated Exadata Infrastructure.
	GetAutonomousContainerDatabaseId() *string

	// The percentage of the System Global Area(SGA) assigned to In-Memory tables in Autonomous Database. This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform.
	GetInMemoryPercentage() *int

	// Indicates if the database-level access control is enabled.
	// If disabled, database access is defined by the network security rules.
	// If enabled, database access is restricted to the IP addresses defined by the rules specified with the `whitelistedIps` property. While specifying `whitelistedIps` rules is optional,
	//  if database-level access control is enabled and no rules are specified, the database will become inaccessible. The rules can be added later using the `UpdateAutonomousDatabase` API operation or edit option in console.
	// When creating a database clone, the desired access control setting should be specified. By default, database-level access control will be disabled for the clone.
	// This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform. For Autonomous Database Serverless instances, `whitelistedIps` is used.
	GetIsAccessControlEnabled() *bool

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
	GetWhitelistedIps() []string

	// This field will be null if the Autonomous Database is not Data Guard enabled or Access Control is disabled.
	// It's value would be `TRUE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses primary IP access control list (ACL) for standby.
	// It's value would be `FALSE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses different IP access control list (ACL) for standby compared to primary.
	GetArePrimaryWhitelistedIpsUsed() *bool

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
	GetStandbyWhitelistedIps() []string

	// **Deprecated.** Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
	GetIsDataGuardEnabled() *bool

	// Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
	GetIsLocalDataGuardEnabled() *bool

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// - For Autonomous Database, setting this will disable public secure access to the database.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and the backup subnet.
	GetSubnetId() *string

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	GetNsgIds() []string

	// The resource's private endpoint label.
	// - Setting the endpoint label to a non-empty string creates a private endpoint database.
	// - Resetting the endpoint label to an empty string, after the creation of the private endpoint database, changes the private endpoint database to a public endpoint database.
	// - Setting the endpoint label to a non-empty string value, updates to a new private endpoint database, when the database is disabled and re-enabled.
	// This setting cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, dbWorkload, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	GetPrivateEndpointLabel() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	GetDefinedTags() map[string]map[string]interface{}

	// The private endpoint Ip address for the resource.
	GetPrivateEndpointIp() *string

	// A valid Oracle Database version for Autonomous Database.
	GetDbVersion() *string

	// Customer Contacts.
	GetCustomerContacts() []CustomerContact

	// Specifies if the Autonomous Database requires mTLS connections.
	// This may not be updated in parallel with any of the following: licenseModel, databaseEdition, cpuCoreCount, computeCount, dataStorageSizeInTBs, whitelistedIps, openMode, permissionLevel, db-workload, privateEndpointLabel, nsgIds, customerContacts, dbVersion, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	// Service Change: The default value of the isMTLSConnectionRequired attribute will change from true to false on July 1, 2023 in the following APIs:
	// - CreateAutonomousDatabase
	// - GetAutonomousDatabase
	// - UpdateAutonomousDatabase
	// Details: Prior to the July 1, 2023 change, the isMTLSConnectionRequired attribute default value was true. This applies to Autonomous Database Serverless.
	// Does this impact me? If you use or maintain custom scripts or Terraform scripts referencing the CreateAutonomousDatabase, GetAutonomousDatabase, or UpdateAutonomousDatabase APIs, you want to check, and possibly modify, the scripts for the changed default value of the attribute. Should you choose not to leave your scripts unchanged, the API calls containing this attribute will continue to work, but the default value will switch from true to false.
	// How do I make this change? Using either OCI SDKs or command line tools, update your custom scripts to explicitly set the isMTLSConnectionRequired attribute to true.
	GetIsMtlsConnectionRequired() *bool

	// The unique identifier for leader autonomous database OCID OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	GetResourcePoolLeaderId() *string

	GetResourcePoolSummary() *ResourcePoolSummary

	// The maintenance schedule type of the Autonomous Database Serverless. An EARLY maintenance schedule
	// follows a schedule applying patches prior to the REGULAR schedule. A REGULAR maintenance schedule follows the normal cycle
	GetAutonomousMaintenanceScheduleType() CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum

	// The list of scheduled operations. Consists of values such as dayOfWeek, scheduledStartTime, scheduledStopTime.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	GetScheduledOperations() []ScheduledOperationDetails

	// Indicates if auto scaling is enabled for the Autonomous Database storage. The default value is `FALSE`.
	GetIsAutoScalingForStorageEnabled() *bool

	// The Oracle Database Edition that applies to the Autonomous databases.
	GetDatabaseEdition() AutonomousDatabaseSummaryDatabaseEditionEnum

	// The list of database tools details.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, isLocalDataGuardEnabled, or isFreeTier.
	GetDbToolsDetails() []DatabaseTool

	// The OCI vault secret [/Content/General/Concepts/identifiers.htm]OCID.
	// This cannot be used in conjunction with adminPassword.
	GetSecretId() *string

	// The version of the vault secret. If no version is specified, the latest version will be used.
	GetSecretVersionNumber() *int
}

type createautonomousdatabasebase struct {
	JsonData                                 []byte
	CharacterSet                             *string                                                           `mandatory:"false" json:"characterSet"`
	NcharacterSet                            *string                                                           `mandatory:"false" json:"ncharacterSet"`
	DbName                                   *string                                                           `mandatory:"false" json:"dbName"`
	CpuCoreCount                             *int                                                              `mandatory:"false" json:"cpuCoreCount"`
	BackupRetentionPeriodInDays              *int                                                              `mandatory:"false" json:"backupRetentionPeriodInDays"`
	ComputeModel                             CreateAutonomousDatabaseBaseComputeModelEnum                      `mandatory:"false" json:"computeModel,omitempty"`
	ComputeCount                             *float32                                                          `mandatory:"false" json:"computeCount"`
	OcpuCount                                *float32                                                          `mandatory:"false" json:"ocpuCount"`
	DbWorkload                               CreateAutonomousDatabaseBaseDbWorkloadEnum                        `mandatory:"false" json:"dbWorkload,omitempty"`
	DataStorageSizeInTBs                     *int                                                              `mandatory:"false" json:"dataStorageSizeInTBs"`
	DataStorageSizeInGBs                     *int                                                              `mandatory:"false" json:"dataStorageSizeInGBs"`
	IsFreeTier                               *bool                                                             `mandatory:"false" json:"isFreeTier"`
	KmsKeyId                                 *string                                                           `mandatory:"false" json:"kmsKeyId"`
	VaultId                                  *string                                                           `mandatory:"false" json:"vaultId"`
	AdminPassword                            *string                                                           `mandatory:"false" json:"adminPassword"`
	DisplayName                              *string                                                           `mandatory:"false" json:"displayName"`
	LicenseModel                             CreateAutonomousDatabaseBaseLicenseModelEnum                      `mandatory:"false" json:"licenseModel,omitempty"`
	IsPreviewVersionWithServiceTermsAccepted *bool                                                             `mandatory:"false" json:"isPreviewVersionWithServiceTermsAccepted"`
	IsAutoScalingEnabled                     *bool                                                             `mandatory:"false" json:"isAutoScalingEnabled"`
	IsDevTier                                *bool                                                             `mandatory:"false" json:"isDevTier"`
	IsDedicated                              *bool                                                             `mandatory:"false" json:"isDedicated"`
	AutonomousContainerDatabaseId            *string                                                           `mandatory:"false" json:"autonomousContainerDatabaseId"`
	InMemoryPercentage                       *int                                                              `mandatory:"false" json:"inMemoryPercentage"`
	IsAccessControlEnabled                   *bool                                                             `mandatory:"false" json:"isAccessControlEnabled"`
	WhitelistedIps                           []string                                                          `mandatory:"false" json:"whitelistedIps"`
	ArePrimaryWhitelistedIpsUsed             *bool                                                             `mandatory:"false" json:"arePrimaryWhitelistedIpsUsed"`
	StandbyWhitelistedIps                    []string                                                          `mandatory:"false" json:"standbyWhitelistedIps"`
	IsDataGuardEnabled                       *bool                                                             `mandatory:"false" json:"isDataGuardEnabled"`
	IsLocalDataGuardEnabled                  *bool                                                             `mandatory:"false" json:"isLocalDataGuardEnabled"`
	SubnetId                                 *string                                                           `mandatory:"false" json:"subnetId"`
	NsgIds                                   []string                                                          `mandatory:"false" json:"nsgIds"`
	PrivateEndpointLabel                     *string                                                           `mandatory:"false" json:"privateEndpointLabel"`
	FreeformTags                             map[string]string                                                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                              map[string]map[string]interface{}                                 `mandatory:"false" json:"definedTags"`
	PrivateEndpointIp                        *string                                                           `mandatory:"false" json:"privateEndpointIp"`
	DbVersion                                *string                                                           `mandatory:"false" json:"dbVersion"`
	CustomerContacts                         []CustomerContact                                                 `mandatory:"false" json:"customerContacts"`
	IsMtlsConnectionRequired                 *bool                                                             `mandatory:"false" json:"isMtlsConnectionRequired"`
	ResourcePoolLeaderId                     *string                                                           `mandatory:"false" json:"resourcePoolLeaderId"`
	ResourcePoolSummary                      *ResourcePoolSummary                                              `mandatory:"false" json:"resourcePoolSummary"`
	AutonomousMaintenanceScheduleType        CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum `mandatory:"false" json:"autonomousMaintenanceScheduleType,omitempty"`
	ScheduledOperations                      []ScheduledOperationDetails                                       `mandatory:"false" json:"scheduledOperations"`
	IsAutoScalingForStorageEnabled           *bool                                                             `mandatory:"false" json:"isAutoScalingForStorageEnabled"`
	DatabaseEdition                          AutonomousDatabaseSummaryDatabaseEditionEnum                      `mandatory:"false" json:"databaseEdition,omitempty"`
	DbToolsDetails                           []DatabaseTool                                                    `mandatory:"false" json:"dbToolsDetails"`
	SecretId                                 *string                                                           `mandatory:"false" json:"secretId"`
	SecretVersionNumber                      *int                                                              `mandatory:"false" json:"secretVersionNumber"`
	CompartmentId                            *string                                                           `mandatory:"true" json:"compartmentId"`
	Source                                   string                                                            `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *createautonomousdatabasebase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateautonomousdatabasebase createautonomousdatabasebase
	s := struct {
		Model Unmarshalercreateautonomousdatabasebase
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.CharacterSet = s.Model.CharacterSet
	m.NcharacterSet = s.Model.NcharacterSet
	m.DbName = s.Model.DbName
	m.CpuCoreCount = s.Model.CpuCoreCount
	m.BackupRetentionPeriodInDays = s.Model.BackupRetentionPeriodInDays
	m.ComputeModel = s.Model.ComputeModel
	m.ComputeCount = s.Model.ComputeCount
	m.OcpuCount = s.Model.OcpuCount
	m.DbWorkload = s.Model.DbWorkload
	m.DataStorageSizeInTBs = s.Model.DataStorageSizeInTBs
	m.DataStorageSizeInGBs = s.Model.DataStorageSizeInGBs
	m.IsFreeTier = s.Model.IsFreeTier
	m.KmsKeyId = s.Model.KmsKeyId
	m.VaultId = s.Model.VaultId
	m.AdminPassword = s.Model.AdminPassword
	m.DisplayName = s.Model.DisplayName
	m.LicenseModel = s.Model.LicenseModel
	m.IsPreviewVersionWithServiceTermsAccepted = s.Model.IsPreviewVersionWithServiceTermsAccepted
	m.IsAutoScalingEnabled = s.Model.IsAutoScalingEnabled
	m.IsDevTier = s.Model.IsDevTier
	m.IsDedicated = s.Model.IsDedicated
	m.AutonomousContainerDatabaseId = s.Model.AutonomousContainerDatabaseId
	m.InMemoryPercentage = s.Model.InMemoryPercentage
	m.IsAccessControlEnabled = s.Model.IsAccessControlEnabled
	m.WhitelistedIps = s.Model.WhitelistedIps
	m.ArePrimaryWhitelistedIpsUsed = s.Model.ArePrimaryWhitelistedIpsUsed
	m.StandbyWhitelistedIps = s.Model.StandbyWhitelistedIps
	m.IsDataGuardEnabled = s.Model.IsDataGuardEnabled
	m.IsLocalDataGuardEnabled = s.Model.IsLocalDataGuardEnabled
	m.SubnetId = s.Model.SubnetId
	m.NsgIds = s.Model.NsgIds
	m.PrivateEndpointLabel = s.Model.PrivateEndpointLabel
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.PrivateEndpointIp = s.Model.PrivateEndpointIp
	m.DbVersion = s.Model.DbVersion
	m.CustomerContacts = s.Model.CustomerContacts
	m.IsMtlsConnectionRequired = s.Model.IsMtlsConnectionRequired
	m.ResourcePoolLeaderId = s.Model.ResourcePoolLeaderId
	m.ResourcePoolSummary = s.Model.ResourcePoolSummary
	m.AutonomousMaintenanceScheduleType = s.Model.AutonomousMaintenanceScheduleType
	m.ScheduledOperations = s.Model.ScheduledOperations
	m.IsAutoScalingForStorageEnabled = s.Model.IsAutoScalingForStorageEnabled
	m.DatabaseEdition = s.Model.DatabaseEdition
	m.DbToolsDetails = s.Model.DbToolsDetails
	m.SecretId = s.Model.SecretId
	m.SecretVersionNumber = s.Model.SecretVersionNumber
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createautonomousdatabasebase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "DATABASE":
		mm := CreateAutonomousDatabaseCloneDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLONE_TO_REFRESHABLE":
		mm := CreateRefreshableAutonomousDatabaseCloneDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BACKUP_FROM_ID":
		mm := CreateAutonomousDatabaseFromBackupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CROSS_REGION_DISASTER_RECOVERY":
		mm := CreateCrossRegionDisasterRecoveryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BACKUP_FROM_TIMESTAMP":
		mm := CreateAutonomousDatabaseFromBackupTimestampDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CROSS_REGION_DATAGUARD":
		mm := CreateCrossRegionAutonomousDatabaseDataGuardDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := CreateAutonomousDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateAutonomousDatabaseBase: %s.", m.Source)
		return *m, nil
	}
}

// GetCharacterSet returns CharacterSet
func (m createautonomousdatabasebase) GetCharacterSet() *string {
	return m.CharacterSet
}

// GetNcharacterSet returns NcharacterSet
func (m createautonomousdatabasebase) GetNcharacterSet() *string {
	return m.NcharacterSet
}

// GetDbName returns DbName
func (m createautonomousdatabasebase) GetDbName() *string {
	return m.DbName
}

// GetCpuCoreCount returns CpuCoreCount
func (m createautonomousdatabasebase) GetCpuCoreCount() *int {
	return m.CpuCoreCount
}

// GetBackupRetentionPeriodInDays returns BackupRetentionPeriodInDays
func (m createautonomousdatabasebase) GetBackupRetentionPeriodInDays() *int {
	return m.BackupRetentionPeriodInDays
}

// GetComputeModel returns ComputeModel
func (m createautonomousdatabasebase) GetComputeModel() CreateAutonomousDatabaseBaseComputeModelEnum {
	return m.ComputeModel
}

// GetComputeCount returns ComputeCount
func (m createautonomousdatabasebase) GetComputeCount() *float32 {
	return m.ComputeCount
}

// GetOcpuCount returns OcpuCount
func (m createautonomousdatabasebase) GetOcpuCount() *float32 {
	return m.OcpuCount
}

// GetDbWorkload returns DbWorkload
func (m createautonomousdatabasebase) GetDbWorkload() CreateAutonomousDatabaseBaseDbWorkloadEnum {
	return m.DbWorkload
}

// GetDataStorageSizeInTBs returns DataStorageSizeInTBs
func (m createautonomousdatabasebase) GetDataStorageSizeInTBs() *int {
	return m.DataStorageSizeInTBs
}

// GetDataStorageSizeInGBs returns DataStorageSizeInGBs
func (m createautonomousdatabasebase) GetDataStorageSizeInGBs() *int {
	return m.DataStorageSizeInGBs
}

// GetIsFreeTier returns IsFreeTier
func (m createautonomousdatabasebase) GetIsFreeTier() *bool {
	return m.IsFreeTier
}

// GetKmsKeyId returns KmsKeyId
func (m createautonomousdatabasebase) GetKmsKeyId() *string {
	return m.KmsKeyId
}

// GetVaultId returns VaultId
func (m createautonomousdatabasebase) GetVaultId() *string {
	return m.VaultId
}

// GetAdminPassword returns AdminPassword
func (m createautonomousdatabasebase) GetAdminPassword() *string {
	return m.AdminPassword
}

// GetDisplayName returns DisplayName
func (m createautonomousdatabasebase) GetDisplayName() *string {
	return m.DisplayName
}

// GetLicenseModel returns LicenseModel
func (m createautonomousdatabasebase) GetLicenseModel() CreateAutonomousDatabaseBaseLicenseModelEnum {
	return m.LicenseModel
}

// GetIsPreviewVersionWithServiceTermsAccepted returns IsPreviewVersionWithServiceTermsAccepted
func (m createautonomousdatabasebase) GetIsPreviewVersionWithServiceTermsAccepted() *bool {
	return m.IsPreviewVersionWithServiceTermsAccepted
}

// GetIsAutoScalingEnabled returns IsAutoScalingEnabled
func (m createautonomousdatabasebase) GetIsAutoScalingEnabled() *bool {
	return m.IsAutoScalingEnabled
}

// GetIsDevTier returns IsDevTier
func (m createautonomousdatabasebase) GetIsDevTier() *bool {
	return m.IsDevTier
}

// GetIsDedicated returns IsDedicated
func (m createautonomousdatabasebase) GetIsDedicated() *bool {
	return m.IsDedicated
}

// GetAutonomousContainerDatabaseId returns AutonomousContainerDatabaseId
func (m createautonomousdatabasebase) GetAutonomousContainerDatabaseId() *string {
	return m.AutonomousContainerDatabaseId
}

// GetInMemoryPercentage returns InMemoryPercentage
func (m createautonomousdatabasebase) GetInMemoryPercentage() *int {
	return m.InMemoryPercentage
}

// GetIsAccessControlEnabled returns IsAccessControlEnabled
func (m createautonomousdatabasebase) GetIsAccessControlEnabled() *bool {
	return m.IsAccessControlEnabled
}

// GetWhitelistedIps returns WhitelistedIps
func (m createautonomousdatabasebase) GetWhitelistedIps() []string {
	return m.WhitelistedIps
}

// GetArePrimaryWhitelistedIpsUsed returns ArePrimaryWhitelistedIpsUsed
func (m createautonomousdatabasebase) GetArePrimaryWhitelistedIpsUsed() *bool {
	return m.ArePrimaryWhitelistedIpsUsed
}

// GetStandbyWhitelistedIps returns StandbyWhitelistedIps
func (m createautonomousdatabasebase) GetStandbyWhitelistedIps() []string {
	return m.StandbyWhitelistedIps
}

// GetIsDataGuardEnabled returns IsDataGuardEnabled
func (m createautonomousdatabasebase) GetIsDataGuardEnabled() *bool {
	return m.IsDataGuardEnabled
}

// GetIsLocalDataGuardEnabled returns IsLocalDataGuardEnabled
func (m createautonomousdatabasebase) GetIsLocalDataGuardEnabled() *bool {
	return m.IsLocalDataGuardEnabled
}

// GetSubnetId returns SubnetId
func (m createautonomousdatabasebase) GetSubnetId() *string {
	return m.SubnetId
}

// GetNsgIds returns NsgIds
func (m createautonomousdatabasebase) GetNsgIds() []string {
	return m.NsgIds
}

// GetPrivateEndpointLabel returns PrivateEndpointLabel
func (m createautonomousdatabasebase) GetPrivateEndpointLabel() *string {
	return m.PrivateEndpointLabel
}

// GetFreeformTags returns FreeformTags
func (m createautonomousdatabasebase) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createautonomousdatabasebase) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetPrivateEndpointIp returns PrivateEndpointIp
func (m createautonomousdatabasebase) GetPrivateEndpointIp() *string {
	return m.PrivateEndpointIp
}

// GetDbVersion returns DbVersion
func (m createautonomousdatabasebase) GetDbVersion() *string {
	return m.DbVersion
}

// GetCustomerContacts returns CustomerContacts
func (m createautonomousdatabasebase) GetCustomerContacts() []CustomerContact {
	return m.CustomerContacts
}

// GetIsMtlsConnectionRequired returns IsMtlsConnectionRequired
func (m createautonomousdatabasebase) GetIsMtlsConnectionRequired() *bool {
	return m.IsMtlsConnectionRequired
}

// GetResourcePoolLeaderId returns ResourcePoolLeaderId
func (m createautonomousdatabasebase) GetResourcePoolLeaderId() *string {
	return m.ResourcePoolLeaderId
}

// GetResourcePoolSummary returns ResourcePoolSummary
func (m createautonomousdatabasebase) GetResourcePoolSummary() *ResourcePoolSummary {
	return m.ResourcePoolSummary
}

// GetAutonomousMaintenanceScheduleType returns AutonomousMaintenanceScheduleType
func (m createautonomousdatabasebase) GetAutonomousMaintenanceScheduleType() CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum {
	return m.AutonomousMaintenanceScheduleType
}

// GetScheduledOperations returns ScheduledOperations
func (m createautonomousdatabasebase) GetScheduledOperations() []ScheduledOperationDetails {
	return m.ScheduledOperations
}

// GetIsAutoScalingForStorageEnabled returns IsAutoScalingForStorageEnabled
func (m createautonomousdatabasebase) GetIsAutoScalingForStorageEnabled() *bool {
	return m.IsAutoScalingForStorageEnabled
}

// GetDatabaseEdition returns DatabaseEdition
func (m createautonomousdatabasebase) GetDatabaseEdition() AutonomousDatabaseSummaryDatabaseEditionEnum {
	return m.DatabaseEdition
}

// GetDbToolsDetails returns DbToolsDetails
func (m createautonomousdatabasebase) GetDbToolsDetails() []DatabaseTool {
	return m.DbToolsDetails
}

// GetSecretId returns SecretId
func (m createautonomousdatabasebase) GetSecretId() *string {
	return m.SecretId
}

// GetSecretVersionNumber returns SecretVersionNumber
func (m createautonomousdatabasebase) GetSecretVersionNumber() *int {
	return m.SecretVersionNumber
}

// GetCompartmentId returns CompartmentId
func (m createautonomousdatabasebase) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m createautonomousdatabasebase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createautonomousdatabasebase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateAutonomousDatabaseBaseComputeModelEnum(string(m.ComputeModel)); !ok && m.ComputeModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeModel: %s. Supported values are: %s.", m.ComputeModel, strings.Join(GetCreateAutonomousDatabaseBaseComputeModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousDatabaseBaseDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetCreateAutonomousDatabaseBaseDbWorkloadEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousDatabaseBaseLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCreateAutonomousDatabaseBaseLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum(string(m.AutonomousMaintenanceScheduleType)); !ok && m.AutonomousMaintenanceScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutonomousMaintenanceScheduleType: %s. Supported values are: %s.", m.AutonomousMaintenanceScheduleType, strings.Join(GetCreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetAutonomousDatabaseSummaryDatabaseEditionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateAutonomousDatabaseBaseComputeModelEnum Enum with underlying type: string
type CreateAutonomousDatabaseBaseComputeModelEnum string

// Set of constants representing the allowable values for CreateAutonomousDatabaseBaseComputeModelEnum
const (
	CreateAutonomousDatabaseBaseComputeModelEcpu CreateAutonomousDatabaseBaseComputeModelEnum = "ECPU"
	CreateAutonomousDatabaseBaseComputeModelOcpu CreateAutonomousDatabaseBaseComputeModelEnum = "OCPU"
)

var mappingCreateAutonomousDatabaseBaseComputeModelEnum = map[string]CreateAutonomousDatabaseBaseComputeModelEnum{
	"ECPU": CreateAutonomousDatabaseBaseComputeModelEcpu,
	"OCPU": CreateAutonomousDatabaseBaseComputeModelOcpu,
}

var mappingCreateAutonomousDatabaseBaseComputeModelEnumLowerCase = map[string]CreateAutonomousDatabaseBaseComputeModelEnum{
	"ecpu": CreateAutonomousDatabaseBaseComputeModelEcpu,
	"ocpu": CreateAutonomousDatabaseBaseComputeModelOcpu,
}

// GetCreateAutonomousDatabaseBaseComputeModelEnumValues Enumerates the set of values for CreateAutonomousDatabaseBaseComputeModelEnum
func GetCreateAutonomousDatabaseBaseComputeModelEnumValues() []CreateAutonomousDatabaseBaseComputeModelEnum {
	values := make([]CreateAutonomousDatabaseBaseComputeModelEnum, 0)
	for _, v := range mappingCreateAutonomousDatabaseBaseComputeModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousDatabaseBaseComputeModelEnumStringValues Enumerates the set of values in String for CreateAutonomousDatabaseBaseComputeModelEnum
func GetCreateAutonomousDatabaseBaseComputeModelEnumStringValues() []string {
	return []string{
		"ECPU",
		"OCPU",
	}
}

// GetMappingCreateAutonomousDatabaseBaseComputeModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousDatabaseBaseComputeModelEnum(val string) (CreateAutonomousDatabaseBaseComputeModelEnum, bool) {
	enum, ok := mappingCreateAutonomousDatabaseBaseComputeModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousDatabaseBaseDbWorkloadEnum Enum with underlying type: string
type CreateAutonomousDatabaseBaseDbWorkloadEnum string

// Set of constants representing the allowable values for CreateAutonomousDatabaseBaseDbWorkloadEnum
const (
	CreateAutonomousDatabaseBaseDbWorkloadOltp CreateAutonomousDatabaseBaseDbWorkloadEnum = "OLTP"
	CreateAutonomousDatabaseBaseDbWorkloadDw   CreateAutonomousDatabaseBaseDbWorkloadEnum = "DW"
	CreateAutonomousDatabaseBaseDbWorkloadAjd  CreateAutonomousDatabaseBaseDbWorkloadEnum = "AJD"
	CreateAutonomousDatabaseBaseDbWorkloadApex CreateAutonomousDatabaseBaseDbWorkloadEnum = "APEX"
)

var mappingCreateAutonomousDatabaseBaseDbWorkloadEnum = map[string]CreateAutonomousDatabaseBaseDbWorkloadEnum{
	"OLTP": CreateAutonomousDatabaseBaseDbWorkloadOltp,
	"DW":   CreateAutonomousDatabaseBaseDbWorkloadDw,
	"AJD":  CreateAutonomousDatabaseBaseDbWorkloadAjd,
	"APEX": CreateAutonomousDatabaseBaseDbWorkloadApex,
}

var mappingCreateAutonomousDatabaseBaseDbWorkloadEnumLowerCase = map[string]CreateAutonomousDatabaseBaseDbWorkloadEnum{
	"oltp": CreateAutonomousDatabaseBaseDbWorkloadOltp,
	"dw":   CreateAutonomousDatabaseBaseDbWorkloadDw,
	"ajd":  CreateAutonomousDatabaseBaseDbWorkloadAjd,
	"apex": CreateAutonomousDatabaseBaseDbWorkloadApex,
}

// GetCreateAutonomousDatabaseBaseDbWorkloadEnumValues Enumerates the set of values for CreateAutonomousDatabaseBaseDbWorkloadEnum
func GetCreateAutonomousDatabaseBaseDbWorkloadEnumValues() []CreateAutonomousDatabaseBaseDbWorkloadEnum {
	values := make([]CreateAutonomousDatabaseBaseDbWorkloadEnum, 0)
	for _, v := range mappingCreateAutonomousDatabaseBaseDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousDatabaseBaseDbWorkloadEnumStringValues Enumerates the set of values in String for CreateAutonomousDatabaseBaseDbWorkloadEnum
func GetCreateAutonomousDatabaseBaseDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
		"AJD",
		"APEX",
	}
}

// GetMappingCreateAutonomousDatabaseBaseDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousDatabaseBaseDbWorkloadEnum(val string) (CreateAutonomousDatabaseBaseDbWorkloadEnum, bool) {
	enum, ok := mappingCreateAutonomousDatabaseBaseDbWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousDatabaseBaseLicenseModelEnum Enum with underlying type: string
type CreateAutonomousDatabaseBaseLicenseModelEnum string

// Set of constants representing the allowable values for CreateAutonomousDatabaseBaseLicenseModelEnum
const (
	CreateAutonomousDatabaseBaseLicenseModelLicenseIncluded     CreateAutonomousDatabaseBaseLicenseModelEnum = "LICENSE_INCLUDED"
	CreateAutonomousDatabaseBaseLicenseModelBringYourOwnLicense CreateAutonomousDatabaseBaseLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateAutonomousDatabaseBaseLicenseModelEnum = map[string]CreateAutonomousDatabaseBaseLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateAutonomousDatabaseBaseLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateAutonomousDatabaseBaseLicenseModelBringYourOwnLicense,
}

var mappingCreateAutonomousDatabaseBaseLicenseModelEnumLowerCase = map[string]CreateAutonomousDatabaseBaseLicenseModelEnum{
	"license_included":       CreateAutonomousDatabaseBaseLicenseModelLicenseIncluded,
	"bring_your_own_license": CreateAutonomousDatabaseBaseLicenseModelBringYourOwnLicense,
}

// GetCreateAutonomousDatabaseBaseLicenseModelEnumValues Enumerates the set of values for CreateAutonomousDatabaseBaseLicenseModelEnum
func GetCreateAutonomousDatabaseBaseLicenseModelEnumValues() []CreateAutonomousDatabaseBaseLicenseModelEnum {
	values := make([]CreateAutonomousDatabaseBaseLicenseModelEnum, 0)
	for _, v := range mappingCreateAutonomousDatabaseBaseLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousDatabaseBaseLicenseModelEnumStringValues Enumerates the set of values in String for CreateAutonomousDatabaseBaseLicenseModelEnum
func GetCreateAutonomousDatabaseBaseLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCreateAutonomousDatabaseBaseLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousDatabaseBaseLicenseModelEnum(val string) (CreateAutonomousDatabaseBaseLicenseModelEnum, bool) {
	enum, ok := mappingCreateAutonomousDatabaseBaseLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum Enum with underlying type: string
type CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum string

// Set of constants representing the allowable values for CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum
const (
	CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEarly   CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum = "EARLY"
	CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeRegular CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum = "REGULAR"
)

var mappingCreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum = map[string]CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum{
	"EARLY":   CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEarly,
	"REGULAR": CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeRegular,
}

var mappingCreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnumLowerCase = map[string]CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum{
	"early":   CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEarly,
	"regular": CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeRegular,
}

// GetCreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnumValues Enumerates the set of values for CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum
func GetCreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnumValues() []CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum {
	values := make([]CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum, 0)
	for _, v := range mappingCreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnumStringValues Enumerates the set of values in String for CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum
func GetCreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnumStringValues() []string {
	return []string{
		"EARLY",
		"REGULAR",
	}
}

// GetMappingCreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum(val string) (CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum, bool) {
	enum, ok := mappingCreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousDatabaseBaseSourceEnum Enum with underlying type: string
type CreateAutonomousDatabaseBaseSourceEnum string

// Set of constants representing the allowable values for CreateAutonomousDatabaseBaseSourceEnum
const (
	CreateAutonomousDatabaseBaseSourceNone                        CreateAutonomousDatabaseBaseSourceEnum = "NONE"
	CreateAutonomousDatabaseBaseSourceDatabase                    CreateAutonomousDatabaseBaseSourceEnum = "DATABASE"
	CreateAutonomousDatabaseBaseSourceBackupFromId                CreateAutonomousDatabaseBaseSourceEnum = "BACKUP_FROM_ID"
	CreateAutonomousDatabaseBaseSourceBackupFromTimestamp         CreateAutonomousDatabaseBaseSourceEnum = "BACKUP_FROM_TIMESTAMP"
	CreateAutonomousDatabaseBaseSourceCloneToRefreshable          CreateAutonomousDatabaseBaseSourceEnum = "CLONE_TO_REFRESHABLE"
	CreateAutonomousDatabaseBaseSourceCrossRegionDataguard        CreateAutonomousDatabaseBaseSourceEnum = "CROSS_REGION_DATAGUARD"
	CreateAutonomousDatabaseBaseSourceCrossRegionDisasterRecovery CreateAutonomousDatabaseBaseSourceEnum = "CROSS_REGION_DISASTER_RECOVERY"
)

var mappingCreateAutonomousDatabaseBaseSourceEnum = map[string]CreateAutonomousDatabaseBaseSourceEnum{
	"NONE":                           CreateAutonomousDatabaseBaseSourceNone,
	"DATABASE":                       CreateAutonomousDatabaseBaseSourceDatabase,
	"BACKUP_FROM_ID":                 CreateAutonomousDatabaseBaseSourceBackupFromId,
	"BACKUP_FROM_TIMESTAMP":          CreateAutonomousDatabaseBaseSourceBackupFromTimestamp,
	"CLONE_TO_REFRESHABLE":           CreateAutonomousDatabaseBaseSourceCloneToRefreshable,
	"CROSS_REGION_DATAGUARD":         CreateAutonomousDatabaseBaseSourceCrossRegionDataguard,
	"CROSS_REGION_DISASTER_RECOVERY": CreateAutonomousDatabaseBaseSourceCrossRegionDisasterRecovery,
}

var mappingCreateAutonomousDatabaseBaseSourceEnumLowerCase = map[string]CreateAutonomousDatabaseBaseSourceEnum{
	"none":                           CreateAutonomousDatabaseBaseSourceNone,
	"database":                       CreateAutonomousDatabaseBaseSourceDatabase,
	"backup_from_id":                 CreateAutonomousDatabaseBaseSourceBackupFromId,
	"backup_from_timestamp":          CreateAutonomousDatabaseBaseSourceBackupFromTimestamp,
	"clone_to_refreshable":           CreateAutonomousDatabaseBaseSourceCloneToRefreshable,
	"cross_region_dataguard":         CreateAutonomousDatabaseBaseSourceCrossRegionDataguard,
	"cross_region_disaster_recovery": CreateAutonomousDatabaseBaseSourceCrossRegionDisasterRecovery,
}

// GetCreateAutonomousDatabaseBaseSourceEnumValues Enumerates the set of values for CreateAutonomousDatabaseBaseSourceEnum
func GetCreateAutonomousDatabaseBaseSourceEnumValues() []CreateAutonomousDatabaseBaseSourceEnum {
	values := make([]CreateAutonomousDatabaseBaseSourceEnum, 0)
	for _, v := range mappingCreateAutonomousDatabaseBaseSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousDatabaseBaseSourceEnumStringValues Enumerates the set of values in String for CreateAutonomousDatabaseBaseSourceEnum
func GetCreateAutonomousDatabaseBaseSourceEnumStringValues() []string {
	return []string{
		"NONE",
		"DATABASE",
		"BACKUP_FROM_ID",
		"BACKUP_FROM_TIMESTAMP",
		"CLONE_TO_REFRESHABLE",
		"CROSS_REGION_DATAGUARD",
		"CROSS_REGION_DISASTER_RECOVERY",
	}
}

// GetMappingCreateAutonomousDatabaseBaseSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousDatabaseBaseSourceEnum(val string) (CreateAutonomousDatabaseBaseSourceEnum, bool) {
	enum, ok := mappingCreateAutonomousDatabaseBaseSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
