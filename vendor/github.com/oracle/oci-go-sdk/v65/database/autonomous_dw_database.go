// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutonomousDwDatabase An Oracle Autonomous Database.
type AutonomousDwDatabase struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the Autonomous Database.
	LifecycleState AutonomousDwDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The database name.
	DbName *string `mandatory:"true" json:"dbName"`

	// The quantity of data in the database, in terabytes.
	// The following points apply to Autonomous Databases on Serverless Infrastructure:
	// - This is an integer field whose value remains null when the data size is in GBs and cannot be converted to TBs (by dividing the GB value by 1024) without rounding error.
	// - To get the exact value of data storage size without rounding error, please see `dataStorageSizeInGBs` of Autonomous Database.
	DataStorageSizeInTBs *int `mandatory:"true" json:"dataStorageSizeInTBs"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// Information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// KMS key lifecycle details.
	KmsKeyLifecycleDetails *string `mandatory:"false" json:"kmsKeyLifecycleDetails"`

	EncryptionKey AutonomousDatabaseEncryptionKeyDetails `mandatory:"false" json:"encryptionKey"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The character set for the autonomous database.  The default is AL32UTF8. Allowed values are:
	// AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSC5601, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P1, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS
	CharacterSet *string `mandatory:"false" json:"characterSet"`

	// The national character set for the autonomous database.  The default is AL16UTF16. Allowed values are:
	// AL16UTF16 or UTF8.
	NcharacterSet *string `mandatory:"false" json:"ncharacterSet"`

	// The percentage of the System Global Area(SGA) assigned to In-Memory tables in Autonomous Database. This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform.
	InMemoryPercentage *int `mandatory:"false" json:"inMemoryPercentage"`

	// The area assigned to In-Memory tables in Autonomous Database.
	InMemoryAreaInGBs *int `mandatory:"false" json:"inMemoryAreaInGBs"`

	// The date and time when the next long-term backup would be created.
	NextLongTermBackupTimeStamp *common.SDKTime `mandatory:"false" json:"nextLongTermBackupTimeStamp"`

	LongTermBackupSchedule *LongTermBackUpScheduleDetails `mandatory:"false" json:"longTermBackupSchedule"`

	// Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isLocalDataGuardEnabled
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The date and time the Always Free database will be stopped because of inactivity. If this time is reached without any database activity, the database will automatically be put into the STOPPED state.
	TimeReclamationOfFreeAutonomousDatabase *common.SDKTime `mandatory:"false" json:"timeReclamationOfFreeAutonomousDatabase"`

	// The date and time the Always Free database will be automatically deleted because of inactivity. If the database is in the STOPPED state and without activity until this time, it will be deleted.
	TimeDeletionOfFreeAutonomousDatabase *common.SDKTime `mandatory:"false" json:"timeDeletionOfFreeAutonomousDatabase"`

	BackupConfig *AutonomousDatabaseBackupConfig `mandatory:"false" json:"backupConfig"`

	// Key History Entry.
	KeyHistoryEntry []AutonomousDatabaseKeyHistoryEntry `mandatory:"false" json:"keyHistoryEntry"`

	// Key History Entry.
	EncryptionKeyHistoryEntry []AutonomousDatabaseEncryptionKeyHistoryEntry `mandatory:"false" json:"encryptionKeyHistoryEntry"`

	// The number of CPU cores to be made available to the database. When the ECPU is selected, the value for cpuCoreCount is 0. For Autonomous Database on Dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `ocpuCount` parameter.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// Parameter that allows users to select an acceptable maximum data loss limit in seconds, up to which Automatic Failover will be triggered when necessary for a Local Autonomous Data Guard
	LocalAdgAutoFailoverMaxDataLossLimit *int `mandatory:"false" json:"localAdgAutoFailoverMaxDataLossLimit"`

	// The compute model of the Autonomous Database. This is required if using the `computeCount` parameter. If using `cpuCoreCount` then it is an error to specify `computeModel` to a non-null value. ECPU compute model is the recommended model and OCPU compute model is legacy.
	ComputeModel AutonomousDwDatabaseComputeModelEnum `mandatory:"false" json:"computeModel,omitempty"`

	// The compute amount (CPUs) available to the database. Minimum and maximum values depend on the compute model and whether the database is an Autonomous Database Serverless instance or an Autonomous Database on Dedicated Exadata Infrastructure.
	// The 'ECPU' compute model requires a minimum value of one, for databases in the elastic resource pool and minimum value of two, otherwise. Required when using the `computeModel` parameter. When using `cpuCoreCount` parameter, it is an error to specify computeCount to a non-null value. Providing `computeModel` and `computeCount` is the preferred method for both OCPU and ECPU.
	ComputeCount *float32 `mandatory:"false" json:"computeCount"`

	// Retention period, in days, for long-term backups
	BackupRetentionPeriodInDays *int `mandatory:"false" json:"backupRetentionPeriodInDays"`

	// The backup storage to the database.
	TotalBackupStorageSizeInGBs *float64 `mandatory:"false" json:"totalBackupStorageSizeInGBs"`

	// The number of OCPU cores to be made available to the database.
	// The following points apply:
	// - For Autonomous Databases on Dedicated Exadata Infrastructure, to provision less than 1 core, enter a fractional value in an increment of 0.1. For example, you can provision 0.3 or 0.4 cores, but not 0.35 cores. (Note that fractional OCPU values are not supported for Autonomous Database Serverless instances.)
	// - To provision cores, enter an integer between 1 and the maximum number of cores available for the infrastructure shape. For example, you can provision 2 cores or 3 cores, but not 2.5 cores. This applies to Autonomous Databases on both serverless and dedicated Exadata infrastructure.
	// - For Autonomous Database Serverless instances, this parameter is not used.
	// For Autonomous Databases on Dedicated Exadata Infrastructure, the maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbde/index.html) for shape details.
	// **Note:** This parameter cannot be used with the `cpuCoreCount` parameter.
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// An array of CPU values that an Autonomous Database can be scaled to.
	ProvisionableCpus []float32 `mandatory:"false" json:"provisionableCpus"`

	// The amount of memory (in GBs rounded off to nearest integer value) enabled per ECPU or OCPU. This is deprecated. Please refer to memoryPerComputeUnitInGBs for accurate value.
	MemoryPerOracleComputeUnitInGBs *int `mandatory:"false" json:"memoryPerOracleComputeUnitInGBs"`

	// The amount of memory (in GBs) to be enabled per OCPU or ECPU.
	MemoryPerComputeUnitInGBs *float32 `mandatory:"false" json:"memoryPerComputeUnitInGBs"`

	// The quantity of data in the database, in gigabytes.
	// For Autonomous Transaction Processing databases using ECPUs on Serverless Infrastructure, this value is always populated. In all the other cases, this value will be null and `dataStorageSizeInTBs` will be populated instead.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The storage space consumed by Autonomous Database in GBs.
	UsedDataStorageSizeInGBs *int `mandatory:"false" json:"usedDataStorageSizeInGBs"`

	// The infrastructure type this resource belongs to.
	InfrastructureType AutonomousDwDatabaseInfrastructureTypeEnum `mandatory:"false" json:"infrastructureType,omitempty"`

	// True if the database uses dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html).
	IsDedicated *bool `mandatory:"false" json:"isDedicated"`

	// The Autonomous Container Database OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Used only by Autonomous Database on Dedicated Exadata Infrastructure.
	AutonomousContainerDatabaseId *string `mandatory:"false" json:"autonomousContainerDatabaseId"`

	// Indicates if the Autonomous Database is backup retention locked.
	IsBackupRetentionLocked *bool `mandatory:"false" json:"isBackupRetentionLocked"`

	// The date and time the Autonomous Database was most recently undeleted.
	TimeUndeleted *common.SDKTime `mandatory:"false" json:"timeUndeleted"`

	// The earliest(min) date and time the Autonomous Database can be scheduled to upgrade to 23ai.
	TimeEarliestAvailableDbVersionUpgrade *common.SDKTime `mandatory:"false" json:"timeEarliestAvailableDbVersionUpgrade"`

	// The max date and time the Autonomous Database can be scheduled to upgrade to 23ai.
	TimeLatestAvailableDbVersionUpgrade *common.SDKTime `mandatory:"false" json:"timeLatestAvailableDbVersionUpgrade"`

	// The date and time the Autonomous Database scheduled to upgrade to 23ai.
	TimeScheduledDbVersionUpgrade *common.SDKTime `mandatory:"false" json:"timeScheduledDbVersionUpgrade"`

	// The date and time the Autonomous Database was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The user-friendly name for the Autonomous Database. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The URL of the Service Console for the Autonomous Database.
	ServiceConsoleUrl *string `mandatory:"false" json:"serviceConsoleUrl"`

	// The connection string used to connect to the Autonomous Database. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Database for the password value.
	ConnectionStrings *AutonomousDatabaseConnectionStrings `mandatory:"false" json:"connectionStrings"`

	ConnectionUrls *AutonomousDatabaseConnectionUrls `mandatory:"false" json:"connectionUrls"`

	// The Public URLs of Private Endpoint database for accessing Oracle Application Express (APEX) and SQL Developer Web with a browser from a Compute instance within your VCN or that has a direct connection to your VCN.
	PublicConnectionUrls *AutonomousDatabaseConnectionUrls `mandatory:"false" json:"publicConnectionUrls"`

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Oracle Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null. It is already set at the
	// Autonomous Exadata Infrastructure level. When provisioning an Autonomous Database Serverless  (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) database, if a value is not specified, the system defaults the value to `BRING_YOUR_OWN_LICENSE`. Bring your own license (BYOL) also allows you to select the DB edition using the optional parameter.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, dataStorageSizeInTBs, adminPassword, isMTLSConnectionRequired, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	LicenseModel AutonomousDwDatabaseLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The maximum number of CPUs allowed with a Bring Your Own License (BYOL), including those used for auto-scaling, disaster recovery, tools, etc. Any CPU usage above this limit is considered as License Included and billed.
	ByolComputeCountLimit *float32 `mandatory:"false" json:"byolComputeCountLimit"`

	// The amount of storage that has been used for Autonomous Databases in dedicated infrastructure, in terabytes.
	UsedDataStorageSizeInTBs *int `mandatory:"false" json:"usedDataStorageSizeInTBs"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Security Attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// - For Autonomous Database, setting this will disable public secure access to the database.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and the backup subnet.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The list of OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.oracle.com/iaas/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The private endpoint for the resource.
	PrivateEndpoint *string `mandatory:"false" json:"privateEndpoint"`

	// The public endpoint for the private endpoint enabled resource.
	PublicEndpoint *string `mandatory:"false" json:"publicEndpoint"`

	// The resource's private endpoint label.
	// - Setting the endpoint label to a non-empty string creates a private endpoint database.
	// - Resetting the endpoint label to an empty string, after the creation of the private endpoint database, changes the private endpoint database to a public endpoint database.
	// - Setting the endpoint label to a non-empty string value, updates to a new private endpoint database, when the database is disabled and re-enabled.
	// This setting cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, dbWorkload, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
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
	// - AJD - indicates an Autonomous JSON Database
	// - APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	DbWorkload AutonomousDwDatabaseDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// Autonomous Database for Developers are fixed-shape Autonomous Databases that developers can use to build and test new applications. On Serverless, these are low-cost and billed per instance, on Dedicated and Cloud@Customer there is no additional cost to create Developer databases. Developer databases come with limited resources and is not intended for large-scale testing and production deployments. When you need more compute or storage resources, you may upgrade to a full paid production database.
	IsDevTier *bool `mandatory:"false" json:"isDevTier"`

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
	// It's value would be `TRUE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses primary IP access control list (ACL) for standby.
	// It's value would be `FALSE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses different IP access control list (ACL) for standby compared to primary.
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

	// Information about Oracle APEX Application Development.
	ApexDetails *AutonomousDatabaseApex `mandatory:"false" json:"apexDetails"`

	// Indicates if auto scaling is enabled for the Autonomous Database CPU core count. The default value is `TRUE`.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// Status of the Data Safe registration for this Autonomous Database.
	DataSafeStatus AutonomousDwDatabaseDataSafeStatusEnum `mandatory:"false" json:"dataSafeStatus,omitempty"`

	// Status of Operations Insights for this Autonomous Database.
	OperationsInsightsStatus AutonomousDwDatabaseOperationsInsightsStatusEnum `mandatory:"false" json:"operationsInsightsStatus,omitempty"`

	// Status of Database Management for this Autonomous Database.
	DatabaseManagementStatus AutonomousDwDatabaseDatabaseManagementStatusEnum `mandatory:"false" json:"databaseManagementStatus,omitempty"`

	// The date and time when maintenance will begin.
	TimeMaintenanceBegin *common.SDKTime `mandatory:"false" json:"timeMaintenanceBegin"`

	// The date and time when maintenance will end.
	TimeMaintenanceEnd *common.SDKTime `mandatory:"false" json:"timeMaintenanceEnd"`

	// The component chosen for maintenance.
	MaintenanceTargetComponent *string `mandatory:"false" json:"maintenanceTargetComponent"`

	// Indicates if the Autonomous Database is a refreshable clone.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	IsRefreshableClone *bool `mandatory:"false" json:"isRefreshableClone"`

	// The date and time when last refresh happened.
	TimeOfLastRefresh *common.SDKTime `mandatory:"false" json:"timeOfLastRefresh"`

	// The refresh point timestamp (UTC). The refresh point is the time to which the database was most recently refreshed. Data created after the refresh point is not included in the refresh.
	TimeOfLastRefreshPoint *common.SDKTime `mandatory:"false" json:"timeOfLastRefreshPoint"`

	// The date and time of next refresh.
	TimeOfNextRefresh *common.SDKTime `mandatory:"false" json:"timeOfNextRefresh"`

	// Indicates the Autonomous Database mode. The database can be opened in `READ_ONLY` or `READ_WRITE` mode.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	OpenMode AutonomousDwDatabaseOpenModeEnum `mandatory:"false" json:"openMode,omitempty"`

	// The refresh status of the clone. REFRESHING indicates that the clone is currently being refreshed with data from the source Autonomous Database.
	RefreshableStatus AutonomousDwDatabaseRefreshableStatusEnum `mandatory:"false" json:"refreshableStatus,omitempty"`

	// The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.
	RefreshableMode AutonomousDwDatabaseRefreshableModeEnum `mandatory:"false" json:"refreshableMode,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that was cloned to create the current Autonomous Database.
	SourceId *string `mandatory:"false" json:"sourceId"`

	// The Autonomous Database permission level. Restricted mode allows access only by admin users.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	PermissionLevel AutonomousDwDatabasePermissionLevelEnum `mandatory:"false" json:"permissionLevel,omitempty"`

	// The timestamp of the last switchover operation for the Autonomous Database.
	TimeOfLastSwitchover *common.SDKTime `mandatory:"false" json:"timeOfLastSwitchover"`

	// The timestamp of the last failover operation.
	TimeOfLastFailover *common.SDKTime `mandatory:"false" json:"timeOfLastFailover"`

	// **Deprecated.** Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
	IsDataGuardEnabled *bool `mandatory:"false" json:"isDataGuardEnabled"`

	// Indicates the number of seconds of data loss for a Data Guard failover.
	FailedDataRecoveryInSeconds *int `mandatory:"false" json:"failedDataRecoveryInSeconds"`

	// **Deprecated** Autonomous Data Guard standby database details.
	StandbyDb *AutonomousDatabaseStandbySummary `mandatory:"false" json:"standbyDb"`

	// Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
	IsLocalDataGuardEnabled *bool `mandatory:"false" json:"isLocalDataGuardEnabled"`

	// Indicates whether the Autonomous Database has Cross Region Data Guard enabled. Not applicable to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
	IsRemoteDataGuardEnabled *bool `mandatory:"false" json:"isRemoteDataGuardEnabled"`

	LocalStandbyDb *AutonomousDatabaseStandbySummary `mandatory:"false" json:"localStandbyDb"`

	// The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled.
	Role AutonomousDwDatabaseRoleEnum `mandatory:"false" json:"role,omitempty"`

	// List of Oracle Database versions available for a database upgrade. If there are no version upgrades available, this list is empty.
	AvailableUpgradeVersions []string `mandatory:"false" json:"availableUpgradeVersions"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	// The wallet name for Oracle Key Vault.
	KeyStoreWalletName *string `mandatory:"false" json:"keyStoreWalletName"`

	// The frequency a refreshable clone is refreshed after auto-refresh is enabled. The minimum is 1 hour. The maximum is 7 days. The date and time that auto-refresh is enabled is controlled by the `timeOfAutoRefreshStart` parameter.
	AutoRefreshFrequencyInSeconds *int `mandatory:"false" json:"autoRefreshFrequencyInSeconds"`

	// The time, in seconds, the data of the refreshable clone lags the primary database at the point of refresh. The minimum is 0 minutes (0 mins means refresh to the latest available timestamp). The maximum is 7 days. The lag time increases after refreshing until the next data refresh happens.
	AutoRefreshPointLagInSeconds *int `mandatory:"false" json:"autoRefreshPointLagInSeconds"`

	// The the date and time that auto-refreshing will begin for an Autonomous Database refreshable clone. This value controls only the start time for the first refresh operation. Subsequent (ongoing) refresh operations have start times controlled by the value of the `autoRefreshFrequencyInSeconds` parameter.
	TimeOfAutoRefreshStart *common.SDKTime `mandatory:"false" json:"timeOfAutoRefreshStart"`

	// The list of regions that support the creation of an Autonomous Database clone or an Autonomous Data Guard standby database.
	SupportedRegionsToCloneTo []string `mandatory:"false" json:"supportedRegionsToCloneTo"`

	// Customer Contacts.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	// The date and time that Autonomous Data Guard was enabled for an Autonomous Database where the standby was provisioned in the same region as the primary database.
	TimeLocalDataGuardEnabled *common.SDKTime `mandatory:"false" json:"timeLocalDataGuardEnabled"`

	// **Deprecated.** The Autonomous Data Guard region type of the Autonomous Database. For Autonomous Database Serverless, Autonomous Data Guard associations have designated primary and standby regions, and these region types do not change when the database changes roles. The standby regions in Autonomous Data Guard associations can be the same region designated as the primary region, or they can be remote regions. Certain database administrative operations may be available only in the primary region of the Autonomous Data Guard association, and cannot be performed when the database using the primary role is operating in a remote Autonomous Data Guard standby region.
	DataguardRegionType AutonomousDwDatabaseDataguardRegionTypeEnum `mandatory:"false" json:"dataguardRegionType,omitempty"`

	// The date and time the Autonomous Data Guard role was switched for the Autonomous Database. For databases that have standbys in both the primary Data Guard region and a remote Data Guard standby region, this is the latest timestamp of either the database using the "primary" role in the primary Data Guard region, or database located in the remote Data Guard standby region.
	TimeDataGuardRoleChanged *common.SDKTime `mandatory:"false" json:"timeDataGuardRoleChanged"`

	// The list of OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of standby databases located in Autonomous Data Guard remote regions that are associated with the source database. Note that for Autonomous Database Serverless instances, standby databases located in the same region as the source primary database do not have OCIDs.
	PeerDbIds []string `mandatory:"false" json:"peerDbIds"`

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

	// The time the member joined the resource pool.
	TimeOfJoiningResourcePool *common.SDKTime `mandatory:"false" json:"timeOfJoiningResourcePool"`

	// The unique identifier for leader autonomous database OCID OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ResourcePoolLeaderId *string `mandatory:"false" json:"resourcePoolLeaderId"`

	ResourcePoolSummary *ResourcePoolSummary `mandatory:"false" json:"resourcePoolSummary"`

	// Indicates if the refreshable clone can be reconnected to its source database.
	IsReconnectCloneEnabled *bool `mandatory:"false" json:"isReconnectCloneEnabled"`

	// The time and date as an RFC3339 formatted string, e.g., 2022-01-01T12:00:00.000Z, to set the limit for a refreshable clone to be reconnected to its source database.
	TimeUntilReconnectCloneEnabled *common.SDKTime `mandatory:"false" json:"timeUntilReconnectCloneEnabled"`

	// The maintenance schedule type of the Autonomous Database Serverless. An EARLY maintenance schedule
	// follows a schedule applying patches prior to the REGULAR schedule. A REGULAR maintenance schedule follows the normal cycle
	AutonomousMaintenanceScheduleType AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum `mandatory:"false" json:"autonomousMaintenanceScheduleType,omitempty"`

	// The list of scheduled operations. Consists of values such as dayOfWeek, scheduledStartTime, scheduledStopTime.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	ScheduledOperations []ScheduledOperationDetails `mandatory:"false" json:"scheduledOperations"`

	// Indicates if auto scaling is enabled for the Autonomous Database storage. The default value is `FALSE`.
	IsAutoScalingForStorageEnabled *bool `mandatory:"false" json:"isAutoScalingForStorageEnabled"`

	// The amount of storage currently allocated for the database tables and billed for, rounded up. When auto-scaling is not enabled, this value is equal to the `dataStorageSizeInTBs` value. You can compare this value to the `actualUsedDataStorageSizeInTBs` value to determine if a manual shrink operation is appropriate for your allocated storage.
	// **Note:** Auto-scaling does not automatically decrease allocated storage when data is deleted from the database.
	AllocatedStorageSizeInTBs *float64 `mandatory:"false" json:"allocatedStorageSizeInTBs"`

	// The current amount of storage in use for user and system data, in terabytes (TB).
	ActualUsedDataStorageSizeInTBs *float64 `mandatory:"false" json:"actualUsedDataStorageSizeInTBs"`

	// The Oracle Database Edition that applies to the Autonomous databases.
	DatabaseEdition AutonomousDwDatabaseDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

	// The list of database tools details.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, isLocalDataGuardEnabled, or isFreeTier.
	DbToolsDetails []DatabaseTool `mandatory:"false" json:"dbToolsDetails"`

	// Indicates the local disaster recovery (DR) type of the Autonomous Database Serverless instance.
	// Autonomous Data Guard (ADG) DR type provides business critical DR with a faster recovery time objective (RTO) during failover or switchover.
	// Backup-based DR type provides lower cost DR with a slower RTO during failover or switchover.
	LocalDisasterRecoveryType DisasterRecoveryConfigurationDisasterRecoveryTypeEnum `mandatory:"false" json:"localDisasterRecoveryType,omitempty"`

	// **Deprecated.** The disaster recovery (DR) region type of the Autonomous Database. For Autonomous Database Serverless instances, DR associations have designated primary and standby regions. These region types do not change when the database changes roles. The standby region in DR associations can be the same region as the primary region, or they can be in a remote regions. Some database administration operations may be available only in the primary region of the DR association, and cannot be performed when the database using the primary role is operating in a remote region.
	DisasterRecoveryRegionType AutonomousDwDatabaseDisasterRecoveryRegionTypeEnum `mandatory:"false" json:"disasterRecoveryRegionType,omitempty"`

	// The date and time the Disaster Recovery role was switched for the standby Autonomous Database.
	TimeDisasterRecoveryRoleChanged *common.SDKTime `mandatory:"false" json:"timeDisasterRecoveryRoleChanged"`

	RemoteDisasterRecoveryConfiguration *DisasterRecoveryConfiguration `mandatory:"false" json:"remoteDisasterRecoveryConfiguration"`

	// Enabling SHARED server architecture enables a database server to allow many client processes to share very few server processes, thereby increasing the number of supported users.
	NetServicesArchitecture AutonomousDwDatabaseNetServicesArchitectureEnum `mandatory:"false" json:"netServicesArchitecture,omitempty"`

	// The availability domain where the Autonomous Database Serverless instance is located.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group of the Autonomous Serverless Database.
	ClusterPlacementGroupId *string `mandatory:"false" json:"clusterPlacementGroupId"`

	// A list of the source Autonomous Database's table space number(s) used to create this partial clone from the backup.
	CloneTableSpaceList []int `mandatory:"false" json:"cloneTableSpaceList"`

	// The Autonomous Database clone type.
	CloneType AutonomousDwDatabaseCloneTypeEnum `mandatory:"false" json:"cloneType,omitempty"`

	// Additional attributes for this resource. Each attribute is a simple key-value pair with no predefined name, type, or namespace.
	// Example: `{ "gcpAccountName": "gcpName" }`
	AdditionalAttributes map[string]string `mandatory:"false" json:"additionalAttributes"`
}

func (m AutonomousDwDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDwDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousDwDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDwDatabaseLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousDwDatabaseComputeModelEnum(string(m.ComputeModel)); !ok && m.ComputeModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeModel: %s. Supported values are: %s.", m.ComputeModel, strings.Join(GetAutonomousDwDatabaseComputeModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseInfrastructureTypeEnum(string(m.InfrastructureType)); !ok && m.InfrastructureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InfrastructureType: %s. Supported values are: %s.", m.InfrastructureType, strings.Join(GetAutonomousDwDatabaseInfrastructureTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetAutonomousDwDatabaseLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetAutonomousDwDatabaseDbWorkloadEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseDataSafeStatusEnum(string(m.DataSafeStatus)); !ok && m.DataSafeStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataSafeStatus: %s. Supported values are: %s.", m.DataSafeStatus, strings.Join(GetAutonomousDwDatabaseDataSafeStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseOperationsInsightsStatusEnum(string(m.OperationsInsightsStatus)); !ok && m.OperationsInsightsStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationsInsightsStatus: %s. Supported values are: %s.", m.OperationsInsightsStatus, strings.Join(GetAutonomousDwDatabaseOperationsInsightsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseDatabaseManagementStatusEnum(string(m.DatabaseManagementStatus)); !ok && m.DatabaseManagementStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseManagementStatus: %s. Supported values are: %s.", m.DatabaseManagementStatus, strings.Join(GetAutonomousDwDatabaseDatabaseManagementStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseOpenModeEnum(string(m.OpenMode)); !ok && m.OpenMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OpenMode: %s. Supported values are: %s.", m.OpenMode, strings.Join(GetAutonomousDwDatabaseOpenModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseRefreshableStatusEnum(string(m.RefreshableStatus)); !ok && m.RefreshableStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RefreshableStatus: %s. Supported values are: %s.", m.RefreshableStatus, strings.Join(GetAutonomousDwDatabaseRefreshableStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseRefreshableModeEnum(string(m.RefreshableMode)); !ok && m.RefreshableMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RefreshableMode: %s. Supported values are: %s.", m.RefreshableMode, strings.Join(GetAutonomousDwDatabaseRefreshableModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabasePermissionLevelEnum(string(m.PermissionLevel)); !ok && m.PermissionLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PermissionLevel: %s. Supported values are: %s.", m.PermissionLevel, strings.Join(GetAutonomousDwDatabasePermissionLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetAutonomousDwDatabaseRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseDataguardRegionTypeEnum(string(m.DataguardRegionType)); !ok && m.DataguardRegionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataguardRegionType: %s. Supported values are: %s.", m.DataguardRegionType, strings.Join(GetAutonomousDwDatabaseDataguardRegionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum(string(m.AutonomousMaintenanceScheduleType)); !ok && m.AutonomousMaintenanceScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutonomousMaintenanceScheduleType: %s. Supported values are: %s.", m.AutonomousMaintenanceScheduleType, strings.Join(GetAutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetAutonomousDwDatabaseDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDisasterRecoveryConfigurationDisasterRecoveryTypeEnum(string(m.LocalDisasterRecoveryType)); !ok && m.LocalDisasterRecoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LocalDisasterRecoveryType: %s. Supported values are: %s.", m.LocalDisasterRecoveryType, strings.Join(GetDisasterRecoveryConfigurationDisasterRecoveryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseDisasterRecoveryRegionTypeEnum(string(m.DisasterRecoveryRegionType)); !ok && m.DisasterRecoveryRegionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DisasterRecoveryRegionType: %s. Supported values are: %s.", m.DisasterRecoveryRegionType, strings.Join(GetAutonomousDwDatabaseDisasterRecoveryRegionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseNetServicesArchitectureEnum(string(m.NetServicesArchitecture)); !ok && m.NetServicesArchitecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetServicesArchitecture: %s. Supported values are: %s.", m.NetServicesArchitecture, strings.Join(GetAutonomousDwDatabaseNetServicesArchitectureEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDwDatabaseCloneTypeEnum(string(m.CloneType)); !ok && m.CloneType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CloneType: %s. Supported values are: %s.", m.CloneType, strings.Join(GetAutonomousDwDatabaseCloneTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AutonomousDwDatabase) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SubscriptionId                          *string                                                   `json:"subscriptionId"`
		LifecycleDetails                        *string                                                   `json:"lifecycleDetails"`
		KmsKeyId                                *string                                                   `json:"kmsKeyId"`
		VaultId                                 *string                                                   `json:"vaultId"`
		KmsKeyLifecycleDetails                  *string                                                   `json:"kmsKeyLifecycleDetails"`
		EncryptionKey                           autonomousdatabaseencryptionkeydetails                    `json:"encryptionKey"`
		KmsKeyVersionId                         *string                                                   `json:"kmsKeyVersionId"`
		CharacterSet                            *string                                                   `json:"characterSet"`
		NcharacterSet                           *string                                                   `json:"ncharacterSet"`
		InMemoryPercentage                      *int                                                      `json:"inMemoryPercentage"`
		InMemoryAreaInGBs                       *int                                                      `json:"inMemoryAreaInGBs"`
		NextLongTermBackupTimeStamp             *common.SDKTime                                           `json:"nextLongTermBackupTimeStamp"`
		LongTermBackupSchedule                  *LongTermBackUpScheduleDetails                            `json:"longTermBackupSchedule"`
		IsFreeTier                              *bool                                                     `json:"isFreeTier"`
		SystemTags                              map[string]map[string]interface{}                         `json:"systemTags"`
		TimeReclamationOfFreeAutonomousDatabase *common.SDKTime                                           `json:"timeReclamationOfFreeAutonomousDatabase"`
		TimeDeletionOfFreeAutonomousDatabase    *common.SDKTime                                           `json:"timeDeletionOfFreeAutonomousDatabase"`
		BackupConfig                            *AutonomousDatabaseBackupConfig                           `json:"backupConfig"`
		KeyHistoryEntry                         []AutonomousDatabaseKeyHistoryEntry                       `json:"keyHistoryEntry"`
		EncryptionKeyHistoryEntry               []AutonomousDatabaseEncryptionKeyHistoryEntry             `json:"encryptionKeyHistoryEntry"`
		CpuCoreCount                            *int                                                      `json:"cpuCoreCount"`
		LocalAdgAutoFailoverMaxDataLossLimit    *int                                                      `json:"localAdgAutoFailoverMaxDataLossLimit"`
		ComputeModel                            AutonomousDwDatabaseComputeModelEnum                      `json:"computeModel"`
		ComputeCount                            *float32                                                  `json:"computeCount"`
		BackupRetentionPeriodInDays             *int                                                      `json:"backupRetentionPeriodInDays"`
		TotalBackupStorageSizeInGBs             *float64                                                  `json:"totalBackupStorageSizeInGBs"`
		OcpuCount                               *float32                                                  `json:"ocpuCount"`
		ProvisionableCpus                       []float32                                                 `json:"provisionableCpus"`
		MemoryPerOracleComputeUnitInGBs         *int                                                      `json:"memoryPerOracleComputeUnitInGBs"`
		MemoryPerComputeUnitInGBs               *float32                                                  `json:"memoryPerComputeUnitInGBs"`
		DataStorageSizeInGBs                    *int                                                      `json:"dataStorageSizeInGBs"`
		UsedDataStorageSizeInGBs                *int                                                      `json:"usedDataStorageSizeInGBs"`
		InfrastructureType                      AutonomousDwDatabaseInfrastructureTypeEnum                `json:"infrastructureType"`
		IsDedicated                             *bool                                                     `json:"isDedicated"`
		AutonomousContainerDatabaseId           *string                                                   `json:"autonomousContainerDatabaseId"`
		IsBackupRetentionLocked                 *bool                                                     `json:"isBackupRetentionLocked"`
		TimeUndeleted                           *common.SDKTime                                           `json:"timeUndeleted"`
		TimeEarliestAvailableDbVersionUpgrade   *common.SDKTime                                           `json:"timeEarliestAvailableDbVersionUpgrade"`
		TimeLatestAvailableDbVersionUpgrade     *common.SDKTime                                           `json:"timeLatestAvailableDbVersionUpgrade"`
		TimeScheduledDbVersionUpgrade           *common.SDKTime                                           `json:"timeScheduledDbVersionUpgrade"`
		TimeCreated                             *common.SDKTime                                           `json:"timeCreated"`
		DisplayName                             *string                                                   `json:"displayName"`
		ServiceConsoleUrl                       *string                                                   `json:"serviceConsoleUrl"`
		ConnectionStrings                       *AutonomousDatabaseConnectionStrings                      `json:"connectionStrings"`
		ConnectionUrls                          *AutonomousDatabaseConnectionUrls                         `json:"connectionUrls"`
		PublicConnectionUrls                    *AutonomousDatabaseConnectionUrls                         `json:"publicConnectionUrls"`
		LicenseModel                            AutonomousDwDatabaseLicenseModelEnum                      `json:"licenseModel"`
		ByolComputeCountLimit                   *float32                                                  `json:"byolComputeCountLimit"`
		UsedDataStorageSizeInTBs                *int                                                      `json:"usedDataStorageSizeInTBs"`
		FreeformTags                            map[string]string                                         `json:"freeformTags"`
		DefinedTags                             map[string]map[string]interface{}                         `json:"definedTags"`
		SecurityAttributes                      map[string]map[string]interface{}                         `json:"securityAttributes"`
		SubnetId                                *string                                                   `json:"subnetId"`
		NsgIds                                  []string                                                  `json:"nsgIds"`
		PrivateEndpoint                         *string                                                   `json:"privateEndpoint"`
		PublicEndpoint                          *string                                                   `json:"publicEndpoint"`
		PrivateEndpointLabel                    *string                                                   `json:"privateEndpointLabel"`
		PrivateEndpointIp                       *string                                                   `json:"privateEndpointIp"`
		DbVersion                               *string                                                   `json:"dbVersion"`
		IsPreview                               *bool                                                     `json:"isPreview"`
		DbWorkload                              AutonomousDwDatabaseDbWorkloadEnum                        `json:"dbWorkload"`
		IsDevTier                               *bool                                                     `json:"isDevTier"`
		IsAccessControlEnabled                  *bool                                                     `json:"isAccessControlEnabled"`
		WhitelistedIps                          []string                                                  `json:"whitelistedIps"`
		ArePrimaryWhitelistedIpsUsed            *bool                                                     `json:"arePrimaryWhitelistedIpsUsed"`
		StandbyWhitelistedIps                   []string                                                  `json:"standbyWhitelistedIps"`
		ApexDetails                             *AutonomousDatabaseApex                                   `json:"apexDetails"`
		IsAutoScalingEnabled                    *bool                                                     `json:"isAutoScalingEnabled"`
		DataSafeStatus                          AutonomousDwDatabaseDataSafeStatusEnum                    `json:"dataSafeStatus"`
		OperationsInsightsStatus                AutonomousDwDatabaseOperationsInsightsStatusEnum          `json:"operationsInsightsStatus"`
		DatabaseManagementStatus                AutonomousDwDatabaseDatabaseManagementStatusEnum          `json:"databaseManagementStatus"`
		TimeMaintenanceBegin                    *common.SDKTime                                           `json:"timeMaintenanceBegin"`
		TimeMaintenanceEnd                      *common.SDKTime                                           `json:"timeMaintenanceEnd"`
		MaintenanceTargetComponent              *string                                                   `json:"maintenanceTargetComponent"`
		IsRefreshableClone                      *bool                                                     `json:"isRefreshableClone"`
		TimeOfLastRefresh                       *common.SDKTime                                           `json:"timeOfLastRefresh"`
		TimeOfLastRefreshPoint                  *common.SDKTime                                           `json:"timeOfLastRefreshPoint"`
		TimeOfNextRefresh                       *common.SDKTime                                           `json:"timeOfNextRefresh"`
		OpenMode                                AutonomousDwDatabaseOpenModeEnum                          `json:"openMode"`
		RefreshableStatus                       AutonomousDwDatabaseRefreshableStatusEnum                 `json:"refreshableStatus"`
		RefreshableMode                         AutonomousDwDatabaseRefreshableModeEnum                   `json:"refreshableMode"`
		SourceId                                *string                                                   `json:"sourceId"`
		PermissionLevel                         AutonomousDwDatabasePermissionLevelEnum                   `json:"permissionLevel"`
		TimeOfLastSwitchover                    *common.SDKTime                                           `json:"timeOfLastSwitchover"`
		TimeOfLastFailover                      *common.SDKTime                                           `json:"timeOfLastFailover"`
		IsDataGuardEnabled                      *bool                                                     `json:"isDataGuardEnabled"`
		FailedDataRecoveryInSeconds             *int                                                      `json:"failedDataRecoveryInSeconds"`
		StandbyDb                               *AutonomousDatabaseStandbySummary                         `json:"standbyDb"`
		IsLocalDataGuardEnabled                 *bool                                                     `json:"isLocalDataGuardEnabled"`
		IsRemoteDataGuardEnabled                *bool                                                     `json:"isRemoteDataGuardEnabled"`
		LocalStandbyDb                          *AutonomousDatabaseStandbySummary                         `json:"localStandbyDb"`
		Role                                    AutonomousDwDatabaseRoleEnum                              `json:"role"`
		AvailableUpgradeVersions                []string                                                  `json:"availableUpgradeVersions"`
		KeyStoreId                              *string                                                   `json:"keyStoreId"`
		KeyStoreWalletName                      *string                                                   `json:"keyStoreWalletName"`
		AutoRefreshFrequencyInSeconds           *int                                                      `json:"autoRefreshFrequencyInSeconds"`
		AutoRefreshPointLagInSeconds            *int                                                      `json:"autoRefreshPointLagInSeconds"`
		TimeOfAutoRefreshStart                  *common.SDKTime                                           `json:"timeOfAutoRefreshStart"`
		SupportedRegionsToCloneTo               []string                                                  `json:"supportedRegionsToCloneTo"`
		CustomerContacts                        []CustomerContact                                         `json:"customerContacts"`
		TimeLocalDataGuardEnabled               *common.SDKTime                                           `json:"timeLocalDataGuardEnabled"`
		DataguardRegionType                     AutonomousDwDatabaseDataguardRegionTypeEnum               `json:"dataguardRegionType"`
		TimeDataGuardRoleChanged                *common.SDKTime                                           `json:"timeDataGuardRoleChanged"`
		PeerDbIds                               []string                                                  `json:"peerDbIds"`
		IsMtlsConnectionRequired                *bool                                                     `json:"isMtlsConnectionRequired"`
		TimeOfJoiningResourcePool               *common.SDKTime                                           `json:"timeOfJoiningResourcePool"`
		ResourcePoolLeaderId                    *string                                                   `json:"resourcePoolLeaderId"`
		ResourcePoolSummary                     *ResourcePoolSummary                                      `json:"resourcePoolSummary"`
		IsReconnectCloneEnabled                 *bool                                                     `json:"isReconnectCloneEnabled"`
		TimeUntilReconnectCloneEnabled          *common.SDKTime                                           `json:"timeUntilReconnectCloneEnabled"`
		AutonomousMaintenanceScheduleType       AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum `json:"autonomousMaintenanceScheduleType"`
		ScheduledOperations                     []ScheduledOperationDetails                               `json:"scheduledOperations"`
		IsAutoScalingForStorageEnabled          *bool                                                     `json:"isAutoScalingForStorageEnabled"`
		AllocatedStorageSizeInTBs               *float64                                                  `json:"allocatedStorageSizeInTBs"`
		ActualUsedDataStorageSizeInTBs          *float64                                                  `json:"actualUsedDataStorageSizeInTBs"`
		DatabaseEdition                         AutonomousDwDatabaseDatabaseEditionEnum                   `json:"databaseEdition"`
		DbToolsDetails                          []DatabaseTool                                            `json:"dbToolsDetails"`
		LocalDisasterRecoveryType               DisasterRecoveryConfigurationDisasterRecoveryTypeEnum     `json:"localDisasterRecoveryType"`
		DisasterRecoveryRegionType              AutonomousDwDatabaseDisasterRecoveryRegionTypeEnum        `json:"disasterRecoveryRegionType"`
		TimeDisasterRecoveryRoleChanged         *common.SDKTime                                           `json:"timeDisasterRecoveryRoleChanged"`
		RemoteDisasterRecoveryConfiguration     *DisasterRecoveryConfiguration                            `json:"remoteDisasterRecoveryConfiguration"`
		NetServicesArchitecture                 AutonomousDwDatabaseNetServicesArchitectureEnum           `json:"netServicesArchitecture"`
		AvailabilityDomain                      *string                                                   `json:"availabilityDomain"`
		ClusterPlacementGroupId                 *string                                                   `json:"clusterPlacementGroupId"`
		CloneTableSpaceList                     []int                                                     `json:"cloneTableSpaceList"`
		CloneType                               AutonomousDwDatabaseCloneTypeEnum                         `json:"cloneType"`
		AdditionalAttributes                    map[string]string                                         `json:"additionalAttributes"`
		Id                                      *string                                                   `json:"id"`
		CompartmentId                           *string                                                   `json:"compartmentId"`
		LifecycleState                          AutonomousDwDatabaseLifecycleStateEnum                    `json:"lifecycleState"`
		DbName                                  *string                                                   `json:"dbName"`
		DataStorageSizeInTBs                    *int                                                      `json:"dataStorageSizeInTBs"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.SubscriptionId = model.SubscriptionId

	m.LifecycleDetails = model.LifecycleDetails

	m.KmsKeyId = model.KmsKeyId

	m.VaultId = model.VaultId

	m.KmsKeyLifecycleDetails = model.KmsKeyLifecycleDetails

	nn, e = model.EncryptionKey.UnmarshalPolymorphicJSON(model.EncryptionKey.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EncryptionKey = nn.(AutonomousDatabaseEncryptionKeyDetails)
	} else {
		m.EncryptionKey = nil
	}

	m.KmsKeyVersionId = model.KmsKeyVersionId

	m.CharacterSet = model.CharacterSet

	m.NcharacterSet = model.NcharacterSet

	m.InMemoryPercentage = model.InMemoryPercentage

	m.InMemoryAreaInGBs = model.InMemoryAreaInGBs

	m.NextLongTermBackupTimeStamp = model.NextLongTermBackupTimeStamp

	m.LongTermBackupSchedule = model.LongTermBackupSchedule

	m.IsFreeTier = model.IsFreeTier

	m.SystemTags = model.SystemTags

	m.TimeReclamationOfFreeAutonomousDatabase = model.TimeReclamationOfFreeAutonomousDatabase

	m.TimeDeletionOfFreeAutonomousDatabase = model.TimeDeletionOfFreeAutonomousDatabase

	m.BackupConfig = model.BackupConfig

	m.KeyHistoryEntry = make([]AutonomousDatabaseKeyHistoryEntry, len(model.KeyHistoryEntry))
	copy(m.KeyHistoryEntry, model.KeyHistoryEntry)
	m.EncryptionKeyHistoryEntry = make([]AutonomousDatabaseEncryptionKeyHistoryEntry, len(model.EncryptionKeyHistoryEntry))
	copy(m.EncryptionKeyHistoryEntry, model.EncryptionKeyHistoryEntry)
	m.CpuCoreCount = model.CpuCoreCount

	m.LocalAdgAutoFailoverMaxDataLossLimit = model.LocalAdgAutoFailoverMaxDataLossLimit

	m.ComputeModel = model.ComputeModel

	m.ComputeCount = model.ComputeCount

	m.BackupRetentionPeriodInDays = model.BackupRetentionPeriodInDays

	m.TotalBackupStorageSizeInGBs = model.TotalBackupStorageSizeInGBs

	m.OcpuCount = model.OcpuCount

	m.ProvisionableCpus = make([]float32, len(model.ProvisionableCpus))
	copy(m.ProvisionableCpus, model.ProvisionableCpus)
	m.MemoryPerOracleComputeUnitInGBs = model.MemoryPerOracleComputeUnitInGBs

	m.MemoryPerComputeUnitInGBs = model.MemoryPerComputeUnitInGBs

	m.DataStorageSizeInGBs = model.DataStorageSizeInGBs

	m.UsedDataStorageSizeInGBs = model.UsedDataStorageSizeInGBs

	m.InfrastructureType = model.InfrastructureType

	m.IsDedicated = model.IsDedicated

	m.AutonomousContainerDatabaseId = model.AutonomousContainerDatabaseId

	m.IsBackupRetentionLocked = model.IsBackupRetentionLocked

	m.TimeUndeleted = model.TimeUndeleted

	m.TimeEarliestAvailableDbVersionUpgrade = model.TimeEarliestAvailableDbVersionUpgrade

	m.TimeLatestAvailableDbVersionUpgrade = model.TimeLatestAvailableDbVersionUpgrade

	m.TimeScheduledDbVersionUpgrade = model.TimeScheduledDbVersionUpgrade

	m.TimeCreated = model.TimeCreated

	m.DisplayName = model.DisplayName

	m.ServiceConsoleUrl = model.ServiceConsoleUrl

	m.ConnectionStrings = model.ConnectionStrings

	m.ConnectionUrls = model.ConnectionUrls

	m.PublicConnectionUrls = model.PublicConnectionUrls

	m.LicenseModel = model.LicenseModel

	m.ByolComputeCountLimit = model.ByolComputeCountLimit

	m.UsedDataStorageSizeInTBs = model.UsedDataStorageSizeInTBs

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SecurityAttributes = model.SecurityAttributes

	m.SubnetId = model.SubnetId

	m.NsgIds = make([]string, len(model.NsgIds))
	copy(m.NsgIds, model.NsgIds)
	m.PrivateEndpoint = model.PrivateEndpoint

	m.PublicEndpoint = model.PublicEndpoint

	m.PrivateEndpointLabel = model.PrivateEndpointLabel

	m.PrivateEndpointIp = model.PrivateEndpointIp

	m.DbVersion = model.DbVersion

	m.IsPreview = model.IsPreview

	m.DbWorkload = model.DbWorkload

	m.IsDevTier = model.IsDevTier

	m.IsAccessControlEnabled = model.IsAccessControlEnabled

	m.WhitelistedIps = make([]string, len(model.WhitelistedIps))
	copy(m.WhitelistedIps, model.WhitelistedIps)
	m.ArePrimaryWhitelistedIpsUsed = model.ArePrimaryWhitelistedIpsUsed

	m.StandbyWhitelistedIps = make([]string, len(model.StandbyWhitelistedIps))
	copy(m.StandbyWhitelistedIps, model.StandbyWhitelistedIps)
	m.ApexDetails = model.ApexDetails

	m.IsAutoScalingEnabled = model.IsAutoScalingEnabled

	m.DataSafeStatus = model.DataSafeStatus

	m.OperationsInsightsStatus = model.OperationsInsightsStatus

	m.DatabaseManagementStatus = model.DatabaseManagementStatus

	m.TimeMaintenanceBegin = model.TimeMaintenanceBegin

	m.TimeMaintenanceEnd = model.TimeMaintenanceEnd

	m.MaintenanceTargetComponent = model.MaintenanceTargetComponent

	m.IsRefreshableClone = model.IsRefreshableClone

	m.TimeOfLastRefresh = model.TimeOfLastRefresh

	m.TimeOfLastRefreshPoint = model.TimeOfLastRefreshPoint

	m.TimeOfNextRefresh = model.TimeOfNextRefresh

	m.OpenMode = model.OpenMode

	m.RefreshableStatus = model.RefreshableStatus

	m.RefreshableMode = model.RefreshableMode

	m.SourceId = model.SourceId

	m.PermissionLevel = model.PermissionLevel

	m.TimeOfLastSwitchover = model.TimeOfLastSwitchover

	m.TimeOfLastFailover = model.TimeOfLastFailover

	m.IsDataGuardEnabled = model.IsDataGuardEnabled

	m.FailedDataRecoveryInSeconds = model.FailedDataRecoveryInSeconds

	m.StandbyDb = model.StandbyDb

	m.IsLocalDataGuardEnabled = model.IsLocalDataGuardEnabled

	m.IsRemoteDataGuardEnabled = model.IsRemoteDataGuardEnabled

	m.LocalStandbyDb = model.LocalStandbyDb

	m.Role = model.Role

	m.AvailableUpgradeVersions = make([]string, len(model.AvailableUpgradeVersions))
	copy(m.AvailableUpgradeVersions, model.AvailableUpgradeVersions)
	m.KeyStoreId = model.KeyStoreId

	m.KeyStoreWalletName = model.KeyStoreWalletName

	m.AutoRefreshFrequencyInSeconds = model.AutoRefreshFrequencyInSeconds

	m.AutoRefreshPointLagInSeconds = model.AutoRefreshPointLagInSeconds

	m.TimeOfAutoRefreshStart = model.TimeOfAutoRefreshStart

	m.SupportedRegionsToCloneTo = make([]string, len(model.SupportedRegionsToCloneTo))
	copy(m.SupportedRegionsToCloneTo, model.SupportedRegionsToCloneTo)
	m.CustomerContacts = make([]CustomerContact, len(model.CustomerContacts))
	copy(m.CustomerContacts, model.CustomerContacts)
	m.TimeLocalDataGuardEnabled = model.TimeLocalDataGuardEnabled

	m.DataguardRegionType = model.DataguardRegionType

	m.TimeDataGuardRoleChanged = model.TimeDataGuardRoleChanged

	m.PeerDbIds = make([]string, len(model.PeerDbIds))
	copy(m.PeerDbIds, model.PeerDbIds)
	m.IsMtlsConnectionRequired = model.IsMtlsConnectionRequired

	m.TimeOfJoiningResourcePool = model.TimeOfJoiningResourcePool

	m.ResourcePoolLeaderId = model.ResourcePoolLeaderId

	m.ResourcePoolSummary = model.ResourcePoolSummary

	m.IsReconnectCloneEnabled = model.IsReconnectCloneEnabled

	m.TimeUntilReconnectCloneEnabled = model.TimeUntilReconnectCloneEnabled

	m.AutonomousMaintenanceScheduleType = model.AutonomousMaintenanceScheduleType

	m.ScheduledOperations = make([]ScheduledOperationDetails, len(model.ScheduledOperations))
	copy(m.ScheduledOperations, model.ScheduledOperations)
	m.IsAutoScalingForStorageEnabled = model.IsAutoScalingForStorageEnabled

	m.AllocatedStorageSizeInTBs = model.AllocatedStorageSizeInTBs

	m.ActualUsedDataStorageSizeInTBs = model.ActualUsedDataStorageSizeInTBs

	m.DatabaseEdition = model.DatabaseEdition

	m.DbToolsDetails = make([]DatabaseTool, len(model.DbToolsDetails))
	copy(m.DbToolsDetails, model.DbToolsDetails)
	m.LocalDisasterRecoveryType = model.LocalDisasterRecoveryType

	m.DisasterRecoveryRegionType = model.DisasterRecoveryRegionType

	m.TimeDisasterRecoveryRoleChanged = model.TimeDisasterRecoveryRoleChanged

	m.RemoteDisasterRecoveryConfiguration = model.RemoteDisasterRecoveryConfiguration

	m.NetServicesArchitecture = model.NetServicesArchitecture

	m.AvailabilityDomain = model.AvailabilityDomain

	m.ClusterPlacementGroupId = model.ClusterPlacementGroupId

	m.CloneTableSpaceList = make([]int, len(model.CloneTableSpaceList))
	copy(m.CloneTableSpaceList, model.CloneTableSpaceList)
	m.CloneType = model.CloneType

	m.AdditionalAttributes = model.AdditionalAttributes

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.LifecycleState = model.LifecycleState

	m.DbName = model.DbName

	m.DataStorageSizeInTBs = model.DataStorageSizeInTBs

	return
}

// AutonomousDwDatabaseLifecycleStateEnum Enum with underlying type: string
type AutonomousDwDatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseLifecycleStateEnum
const (
	AutonomousDwDatabaseLifecycleStateProvisioning            AutonomousDwDatabaseLifecycleStateEnum = "PROVISIONING"
	AutonomousDwDatabaseLifecycleStateAvailable               AutonomousDwDatabaseLifecycleStateEnum = "AVAILABLE"
	AutonomousDwDatabaseLifecycleStateStopping                AutonomousDwDatabaseLifecycleStateEnum = "STOPPING"
	AutonomousDwDatabaseLifecycleStateStopped                 AutonomousDwDatabaseLifecycleStateEnum = "STOPPED"
	AutonomousDwDatabaseLifecycleStateStarting                AutonomousDwDatabaseLifecycleStateEnum = "STARTING"
	AutonomousDwDatabaseLifecycleStateTerminating             AutonomousDwDatabaseLifecycleStateEnum = "TERMINATING"
	AutonomousDwDatabaseLifecycleStateTerminated              AutonomousDwDatabaseLifecycleStateEnum = "TERMINATED"
	AutonomousDwDatabaseLifecycleStateUnavailable             AutonomousDwDatabaseLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDwDatabaseLifecycleStateRestoreInProgress       AutonomousDwDatabaseLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDwDatabaseLifecycleStateRestoreFailed           AutonomousDwDatabaseLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousDwDatabaseLifecycleStateBackupInProgress        AutonomousDwDatabaseLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDwDatabaseLifecycleStateScaleInProgress         AutonomousDwDatabaseLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDwDatabaseLifecycleStateAvailableNeedsAttention AutonomousDwDatabaseLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDwDatabaseLifecycleStateUpdating                AutonomousDwDatabaseLifecycleStateEnum = "UPDATING"
	AutonomousDwDatabaseLifecycleStateMaintenanceInProgress   AutonomousDwDatabaseLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	AutonomousDwDatabaseLifecycleStateRestarting              AutonomousDwDatabaseLifecycleStateEnum = "RESTARTING"
	AutonomousDwDatabaseLifecycleStateRecreating              AutonomousDwDatabaseLifecycleStateEnum = "RECREATING"
	AutonomousDwDatabaseLifecycleStateRoleChangeInProgress    AutonomousDwDatabaseLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousDwDatabaseLifecycleStateUpgrading               AutonomousDwDatabaseLifecycleStateEnum = "UPGRADING"
	AutonomousDwDatabaseLifecycleStateInaccessible            AutonomousDwDatabaseLifecycleStateEnum = "INACCESSIBLE"
	AutonomousDwDatabaseLifecycleStateStandby                 AutonomousDwDatabaseLifecycleStateEnum = "STANDBY"
)

var mappingAutonomousDwDatabaseLifecycleStateEnum = map[string]AutonomousDwDatabaseLifecycleStateEnum{
	"PROVISIONING":              AutonomousDwDatabaseLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDwDatabaseLifecycleStateAvailable,
	"STOPPING":                  AutonomousDwDatabaseLifecycleStateStopping,
	"STOPPED":                   AutonomousDwDatabaseLifecycleStateStopped,
	"STARTING":                  AutonomousDwDatabaseLifecycleStateStarting,
	"TERMINATING":               AutonomousDwDatabaseLifecycleStateTerminating,
	"TERMINATED":                AutonomousDwDatabaseLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDwDatabaseLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDwDatabaseLifecycleStateRestoreInProgress,
	"RESTORE_FAILED":            AutonomousDwDatabaseLifecycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":        AutonomousDwDatabaseLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDwDatabaseLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDwDatabaseLifecycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDwDatabaseLifecycleStateUpdating,
	"MAINTENANCE_IN_PROGRESS":   AutonomousDwDatabaseLifecycleStateMaintenanceInProgress,
	"RESTARTING":                AutonomousDwDatabaseLifecycleStateRestarting,
	"RECREATING":                AutonomousDwDatabaseLifecycleStateRecreating,
	"ROLE_CHANGE_IN_PROGRESS":   AutonomousDwDatabaseLifecycleStateRoleChangeInProgress,
	"UPGRADING":                 AutonomousDwDatabaseLifecycleStateUpgrading,
	"INACCESSIBLE":              AutonomousDwDatabaseLifecycleStateInaccessible,
	"STANDBY":                   AutonomousDwDatabaseLifecycleStateStandby,
}

var mappingAutonomousDwDatabaseLifecycleStateEnumLowerCase = map[string]AutonomousDwDatabaseLifecycleStateEnum{
	"provisioning":              AutonomousDwDatabaseLifecycleStateProvisioning,
	"available":                 AutonomousDwDatabaseLifecycleStateAvailable,
	"stopping":                  AutonomousDwDatabaseLifecycleStateStopping,
	"stopped":                   AutonomousDwDatabaseLifecycleStateStopped,
	"starting":                  AutonomousDwDatabaseLifecycleStateStarting,
	"terminating":               AutonomousDwDatabaseLifecycleStateTerminating,
	"terminated":                AutonomousDwDatabaseLifecycleStateTerminated,
	"unavailable":               AutonomousDwDatabaseLifecycleStateUnavailable,
	"restore_in_progress":       AutonomousDwDatabaseLifecycleStateRestoreInProgress,
	"restore_failed":            AutonomousDwDatabaseLifecycleStateRestoreFailed,
	"backup_in_progress":        AutonomousDwDatabaseLifecycleStateBackupInProgress,
	"scale_in_progress":         AutonomousDwDatabaseLifecycleStateScaleInProgress,
	"available_needs_attention": AutonomousDwDatabaseLifecycleStateAvailableNeedsAttention,
	"updating":                  AutonomousDwDatabaseLifecycleStateUpdating,
	"maintenance_in_progress":   AutonomousDwDatabaseLifecycleStateMaintenanceInProgress,
	"restarting":                AutonomousDwDatabaseLifecycleStateRestarting,
	"recreating":                AutonomousDwDatabaseLifecycleStateRecreating,
	"role_change_in_progress":   AutonomousDwDatabaseLifecycleStateRoleChangeInProgress,
	"upgrading":                 AutonomousDwDatabaseLifecycleStateUpgrading,
	"inaccessible":              AutonomousDwDatabaseLifecycleStateInaccessible,
	"standby":                   AutonomousDwDatabaseLifecycleStateStandby,
}

// GetAutonomousDwDatabaseLifecycleStateEnumValues Enumerates the set of values for AutonomousDwDatabaseLifecycleStateEnum
func GetAutonomousDwDatabaseLifecycleStateEnumValues() []AutonomousDwDatabaseLifecycleStateEnum {
	values := make([]AutonomousDwDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseLifecycleStateEnum
func GetAutonomousDwDatabaseLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"STOPPING",
		"STOPPED",
		"STARTING",
		"TERMINATING",
		"TERMINATED",
		"UNAVAILABLE",
		"RESTORE_IN_PROGRESS",
		"RESTORE_FAILED",
		"BACKUP_IN_PROGRESS",
		"SCALE_IN_PROGRESS",
		"AVAILABLE_NEEDS_ATTENTION",
		"UPDATING",
		"MAINTENANCE_IN_PROGRESS",
		"RESTARTING",
		"RECREATING",
		"ROLE_CHANGE_IN_PROGRESS",
		"UPGRADING",
		"INACCESSIBLE",
		"STANDBY",
	}
}

// GetMappingAutonomousDwDatabaseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseLifecycleStateEnum(val string) (AutonomousDwDatabaseLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseComputeModelEnum Enum with underlying type: string
type AutonomousDwDatabaseComputeModelEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseComputeModelEnum
const (
	AutonomousDwDatabaseComputeModelEcpu AutonomousDwDatabaseComputeModelEnum = "ECPU"
	AutonomousDwDatabaseComputeModelOcpu AutonomousDwDatabaseComputeModelEnum = "OCPU"
)

var mappingAutonomousDwDatabaseComputeModelEnum = map[string]AutonomousDwDatabaseComputeModelEnum{
	"ECPU": AutonomousDwDatabaseComputeModelEcpu,
	"OCPU": AutonomousDwDatabaseComputeModelOcpu,
}

var mappingAutonomousDwDatabaseComputeModelEnumLowerCase = map[string]AutonomousDwDatabaseComputeModelEnum{
	"ecpu": AutonomousDwDatabaseComputeModelEcpu,
	"ocpu": AutonomousDwDatabaseComputeModelOcpu,
}

// GetAutonomousDwDatabaseComputeModelEnumValues Enumerates the set of values for AutonomousDwDatabaseComputeModelEnum
func GetAutonomousDwDatabaseComputeModelEnumValues() []AutonomousDwDatabaseComputeModelEnum {
	values := make([]AutonomousDwDatabaseComputeModelEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseComputeModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseComputeModelEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseComputeModelEnum
func GetAutonomousDwDatabaseComputeModelEnumStringValues() []string {
	return []string{
		"ECPU",
		"OCPU",
	}
}

// GetMappingAutonomousDwDatabaseComputeModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseComputeModelEnum(val string) (AutonomousDwDatabaseComputeModelEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseComputeModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseInfrastructureTypeEnum Enum with underlying type: string
type AutonomousDwDatabaseInfrastructureTypeEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseInfrastructureTypeEnum
const (
	AutonomousDwDatabaseInfrastructureTypeCloud           AutonomousDwDatabaseInfrastructureTypeEnum = "CLOUD"
	AutonomousDwDatabaseInfrastructureTypeCloudAtCustomer AutonomousDwDatabaseInfrastructureTypeEnum = "CLOUD_AT_CUSTOMER"
)

var mappingAutonomousDwDatabaseInfrastructureTypeEnum = map[string]AutonomousDwDatabaseInfrastructureTypeEnum{
	"CLOUD":             AutonomousDwDatabaseInfrastructureTypeCloud,
	"CLOUD_AT_CUSTOMER": AutonomousDwDatabaseInfrastructureTypeCloudAtCustomer,
}

var mappingAutonomousDwDatabaseInfrastructureTypeEnumLowerCase = map[string]AutonomousDwDatabaseInfrastructureTypeEnum{
	"cloud":             AutonomousDwDatabaseInfrastructureTypeCloud,
	"cloud_at_customer": AutonomousDwDatabaseInfrastructureTypeCloudAtCustomer,
}

// GetAutonomousDwDatabaseInfrastructureTypeEnumValues Enumerates the set of values for AutonomousDwDatabaseInfrastructureTypeEnum
func GetAutonomousDwDatabaseInfrastructureTypeEnumValues() []AutonomousDwDatabaseInfrastructureTypeEnum {
	values := make([]AutonomousDwDatabaseInfrastructureTypeEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseInfrastructureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseInfrastructureTypeEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseInfrastructureTypeEnum
func GetAutonomousDwDatabaseInfrastructureTypeEnumStringValues() []string {
	return []string{
		"CLOUD",
		"CLOUD_AT_CUSTOMER",
	}
}

// GetMappingAutonomousDwDatabaseInfrastructureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseInfrastructureTypeEnum(val string) (AutonomousDwDatabaseInfrastructureTypeEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseInfrastructureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseLicenseModelEnum Enum with underlying type: string
type AutonomousDwDatabaseLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseLicenseModelEnum
const (
	AutonomousDwDatabaseLicenseModelLicenseIncluded     AutonomousDwDatabaseLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousDwDatabaseLicenseModelBringYourOwnLicense AutonomousDwDatabaseLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousDwDatabaseLicenseModelEnum = map[string]AutonomousDwDatabaseLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousDwDatabaseLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousDwDatabaseLicenseModelBringYourOwnLicense,
}

var mappingAutonomousDwDatabaseLicenseModelEnumLowerCase = map[string]AutonomousDwDatabaseLicenseModelEnum{
	"license_included":       AutonomousDwDatabaseLicenseModelLicenseIncluded,
	"bring_your_own_license": AutonomousDwDatabaseLicenseModelBringYourOwnLicense,
}

// GetAutonomousDwDatabaseLicenseModelEnumValues Enumerates the set of values for AutonomousDwDatabaseLicenseModelEnum
func GetAutonomousDwDatabaseLicenseModelEnumValues() []AutonomousDwDatabaseLicenseModelEnum {
	values := make([]AutonomousDwDatabaseLicenseModelEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseLicenseModelEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseLicenseModelEnum
func GetAutonomousDwDatabaseLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingAutonomousDwDatabaseLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseLicenseModelEnum(val string) (AutonomousDwDatabaseLicenseModelEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseDbWorkloadEnum Enum with underlying type: string
type AutonomousDwDatabaseDbWorkloadEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseDbWorkloadEnum
const (
	AutonomousDwDatabaseDbWorkloadOltp AutonomousDwDatabaseDbWorkloadEnum = "OLTP"
	AutonomousDwDatabaseDbWorkloadDw   AutonomousDwDatabaseDbWorkloadEnum = "DW"
	AutonomousDwDatabaseDbWorkloadAjd  AutonomousDwDatabaseDbWorkloadEnum = "AJD"
	AutonomousDwDatabaseDbWorkloadApex AutonomousDwDatabaseDbWorkloadEnum = "APEX"
)

var mappingAutonomousDwDatabaseDbWorkloadEnum = map[string]AutonomousDwDatabaseDbWorkloadEnum{
	"OLTP": AutonomousDwDatabaseDbWorkloadOltp,
	"DW":   AutonomousDwDatabaseDbWorkloadDw,
	"AJD":  AutonomousDwDatabaseDbWorkloadAjd,
	"APEX": AutonomousDwDatabaseDbWorkloadApex,
}

var mappingAutonomousDwDatabaseDbWorkloadEnumLowerCase = map[string]AutonomousDwDatabaseDbWorkloadEnum{
	"oltp": AutonomousDwDatabaseDbWorkloadOltp,
	"dw":   AutonomousDwDatabaseDbWorkloadDw,
	"ajd":  AutonomousDwDatabaseDbWorkloadAjd,
	"apex": AutonomousDwDatabaseDbWorkloadApex,
}

// GetAutonomousDwDatabaseDbWorkloadEnumValues Enumerates the set of values for AutonomousDwDatabaseDbWorkloadEnum
func GetAutonomousDwDatabaseDbWorkloadEnumValues() []AutonomousDwDatabaseDbWorkloadEnum {
	values := make([]AutonomousDwDatabaseDbWorkloadEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseDbWorkloadEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseDbWorkloadEnum
func GetAutonomousDwDatabaseDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
		"AJD",
		"APEX",
	}
}

// GetMappingAutonomousDwDatabaseDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseDbWorkloadEnum(val string) (AutonomousDwDatabaseDbWorkloadEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseDbWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseDataSafeStatusEnum Enum with underlying type: string
type AutonomousDwDatabaseDataSafeStatusEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseDataSafeStatusEnum
const (
	AutonomousDwDatabaseDataSafeStatusRegistering   AutonomousDwDatabaseDataSafeStatusEnum = "REGISTERING"
	AutonomousDwDatabaseDataSafeStatusRegistered    AutonomousDwDatabaseDataSafeStatusEnum = "REGISTERED"
	AutonomousDwDatabaseDataSafeStatusDeregistering AutonomousDwDatabaseDataSafeStatusEnum = "DEREGISTERING"
	AutonomousDwDatabaseDataSafeStatusNotRegistered AutonomousDwDatabaseDataSafeStatusEnum = "NOT_REGISTERED"
	AutonomousDwDatabaseDataSafeStatusFailed        AutonomousDwDatabaseDataSafeStatusEnum = "FAILED"
)

var mappingAutonomousDwDatabaseDataSafeStatusEnum = map[string]AutonomousDwDatabaseDataSafeStatusEnum{
	"REGISTERING":    AutonomousDwDatabaseDataSafeStatusRegistering,
	"REGISTERED":     AutonomousDwDatabaseDataSafeStatusRegistered,
	"DEREGISTERING":  AutonomousDwDatabaseDataSafeStatusDeregistering,
	"NOT_REGISTERED": AutonomousDwDatabaseDataSafeStatusNotRegistered,
	"FAILED":         AutonomousDwDatabaseDataSafeStatusFailed,
}

var mappingAutonomousDwDatabaseDataSafeStatusEnumLowerCase = map[string]AutonomousDwDatabaseDataSafeStatusEnum{
	"registering":    AutonomousDwDatabaseDataSafeStatusRegistering,
	"registered":     AutonomousDwDatabaseDataSafeStatusRegistered,
	"deregistering":  AutonomousDwDatabaseDataSafeStatusDeregistering,
	"not_registered": AutonomousDwDatabaseDataSafeStatusNotRegistered,
	"failed":         AutonomousDwDatabaseDataSafeStatusFailed,
}

// GetAutonomousDwDatabaseDataSafeStatusEnumValues Enumerates the set of values for AutonomousDwDatabaseDataSafeStatusEnum
func GetAutonomousDwDatabaseDataSafeStatusEnumValues() []AutonomousDwDatabaseDataSafeStatusEnum {
	values := make([]AutonomousDwDatabaseDataSafeStatusEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseDataSafeStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseDataSafeStatusEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseDataSafeStatusEnum
func GetAutonomousDwDatabaseDataSafeStatusEnumStringValues() []string {
	return []string{
		"REGISTERING",
		"REGISTERED",
		"DEREGISTERING",
		"NOT_REGISTERED",
		"FAILED",
	}
}

// GetMappingAutonomousDwDatabaseDataSafeStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseDataSafeStatusEnum(val string) (AutonomousDwDatabaseDataSafeStatusEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseDataSafeStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseOperationsInsightsStatusEnum Enum with underlying type: string
type AutonomousDwDatabaseOperationsInsightsStatusEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseOperationsInsightsStatusEnum
const (
	AutonomousDwDatabaseOperationsInsightsStatusEnabling        AutonomousDwDatabaseOperationsInsightsStatusEnum = "ENABLING"
	AutonomousDwDatabaseOperationsInsightsStatusEnabled         AutonomousDwDatabaseOperationsInsightsStatusEnum = "ENABLED"
	AutonomousDwDatabaseOperationsInsightsStatusDisabling       AutonomousDwDatabaseOperationsInsightsStatusEnum = "DISABLING"
	AutonomousDwDatabaseOperationsInsightsStatusNotEnabled      AutonomousDwDatabaseOperationsInsightsStatusEnum = "NOT_ENABLED"
	AutonomousDwDatabaseOperationsInsightsStatusFailedEnabling  AutonomousDwDatabaseOperationsInsightsStatusEnum = "FAILED_ENABLING"
	AutonomousDwDatabaseOperationsInsightsStatusFailedDisabling AutonomousDwDatabaseOperationsInsightsStatusEnum = "FAILED_DISABLING"
)

var mappingAutonomousDwDatabaseOperationsInsightsStatusEnum = map[string]AutonomousDwDatabaseOperationsInsightsStatusEnum{
	"ENABLING":         AutonomousDwDatabaseOperationsInsightsStatusEnabling,
	"ENABLED":          AutonomousDwDatabaseOperationsInsightsStatusEnabled,
	"DISABLING":        AutonomousDwDatabaseOperationsInsightsStatusDisabling,
	"NOT_ENABLED":      AutonomousDwDatabaseOperationsInsightsStatusNotEnabled,
	"FAILED_ENABLING":  AutonomousDwDatabaseOperationsInsightsStatusFailedEnabling,
	"FAILED_DISABLING": AutonomousDwDatabaseOperationsInsightsStatusFailedDisabling,
}

var mappingAutonomousDwDatabaseOperationsInsightsStatusEnumLowerCase = map[string]AutonomousDwDatabaseOperationsInsightsStatusEnum{
	"enabling":         AutonomousDwDatabaseOperationsInsightsStatusEnabling,
	"enabled":          AutonomousDwDatabaseOperationsInsightsStatusEnabled,
	"disabling":        AutonomousDwDatabaseOperationsInsightsStatusDisabling,
	"not_enabled":      AutonomousDwDatabaseOperationsInsightsStatusNotEnabled,
	"failed_enabling":  AutonomousDwDatabaseOperationsInsightsStatusFailedEnabling,
	"failed_disabling": AutonomousDwDatabaseOperationsInsightsStatusFailedDisabling,
}

// GetAutonomousDwDatabaseOperationsInsightsStatusEnumValues Enumerates the set of values for AutonomousDwDatabaseOperationsInsightsStatusEnum
func GetAutonomousDwDatabaseOperationsInsightsStatusEnumValues() []AutonomousDwDatabaseOperationsInsightsStatusEnum {
	values := make([]AutonomousDwDatabaseOperationsInsightsStatusEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseOperationsInsightsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseOperationsInsightsStatusEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseOperationsInsightsStatusEnum
func GetAutonomousDwDatabaseOperationsInsightsStatusEnumStringValues() []string {
	return []string{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"NOT_ENABLED",
		"FAILED_ENABLING",
		"FAILED_DISABLING",
	}
}

// GetMappingAutonomousDwDatabaseOperationsInsightsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseOperationsInsightsStatusEnum(val string) (AutonomousDwDatabaseOperationsInsightsStatusEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseOperationsInsightsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseDatabaseManagementStatusEnum Enum with underlying type: string
type AutonomousDwDatabaseDatabaseManagementStatusEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseDatabaseManagementStatusEnum
const (
	AutonomousDwDatabaseDatabaseManagementStatusEnabling        AutonomousDwDatabaseDatabaseManagementStatusEnum = "ENABLING"
	AutonomousDwDatabaseDatabaseManagementStatusEnabled         AutonomousDwDatabaseDatabaseManagementStatusEnum = "ENABLED"
	AutonomousDwDatabaseDatabaseManagementStatusDisabling       AutonomousDwDatabaseDatabaseManagementStatusEnum = "DISABLING"
	AutonomousDwDatabaseDatabaseManagementStatusNotEnabled      AutonomousDwDatabaseDatabaseManagementStatusEnum = "NOT_ENABLED"
	AutonomousDwDatabaseDatabaseManagementStatusFailedEnabling  AutonomousDwDatabaseDatabaseManagementStatusEnum = "FAILED_ENABLING"
	AutonomousDwDatabaseDatabaseManagementStatusFailedDisabling AutonomousDwDatabaseDatabaseManagementStatusEnum = "FAILED_DISABLING"
)

var mappingAutonomousDwDatabaseDatabaseManagementStatusEnum = map[string]AutonomousDwDatabaseDatabaseManagementStatusEnum{
	"ENABLING":         AutonomousDwDatabaseDatabaseManagementStatusEnabling,
	"ENABLED":          AutonomousDwDatabaseDatabaseManagementStatusEnabled,
	"DISABLING":        AutonomousDwDatabaseDatabaseManagementStatusDisabling,
	"NOT_ENABLED":      AutonomousDwDatabaseDatabaseManagementStatusNotEnabled,
	"FAILED_ENABLING":  AutonomousDwDatabaseDatabaseManagementStatusFailedEnabling,
	"FAILED_DISABLING": AutonomousDwDatabaseDatabaseManagementStatusFailedDisabling,
}

var mappingAutonomousDwDatabaseDatabaseManagementStatusEnumLowerCase = map[string]AutonomousDwDatabaseDatabaseManagementStatusEnum{
	"enabling":         AutonomousDwDatabaseDatabaseManagementStatusEnabling,
	"enabled":          AutonomousDwDatabaseDatabaseManagementStatusEnabled,
	"disabling":        AutonomousDwDatabaseDatabaseManagementStatusDisabling,
	"not_enabled":      AutonomousDwDatabaseDatabaseManagementStatusNotEnabled,
	"failed_enabling":  AutonomousDwDatabaseDatabaseManagementStatusFailedEnabling,
	"failed_disabling": AutonomousDwDatabaseDatabaseManagementStatusFailedDisabling,
}

// GetAutonomousDwDatabaseDatabaseManagementStatusEnumValues Enumerates the set of values for AutonomousDwDatabaseDatabaseManagementStatusEnum
func GetAutonomousDwDatabaseDatabaseManagementStatusEnumValues() []AutonomousDwDatabaseDatabaseManagementStatusEnum {
	values := make([]AutonomousDwDatabaseDatabaseManagementStatusEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseDatabaseManagementStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseDatabaseManagementStatusEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseDatabaseManagementStatusEnum
func GetAutonomousDwDatabaseDatabaseManagementStatusEnumStringValues() []string {
	return []string{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"NOT_ENABLED",
		"FAILED_ENABLING",
		"FAILED_DISABLING",
	}
}

// GetMappingAutonomousDwDatabaseDatabaseManagementStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseDatabaseManagementStatusEnum(val string) (AutonomousDwDatabaseDatabaseManagementStatusEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseDatabaseManagementStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseOpenModeEnum Enum with underlying type: string
type AutonomousDwDatabaseOpenModeEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseOpenModeEnum
const (
	AutonomousDwDatabaseOpenModeOnly  AutonomousDwDatabaseOpenModeEnum = "READ_ONLY"
	AutonomousDwDatabaseOpenModeWrite AutonomousDwDatabaseOpenModeEnum = "READ_WRITE"
)

var mappingAutonomousDwDatabaseOpenModeEnum = map[string]AutonomousDwDatabaseOpenModeEnum{
	"READ_ONLY":  AutonomousDwDatabaseOpenModeOnly,
	"READ_WRITE": AutonomousDwDatabaseOpenModeWrite,
}

var mappingAutonomousDwDatabaseOpenModeEnumLowerCase = map[string]AutonomousDwDatabaseOpenModeEnum{
	"read_only":  AutonomousDwDatabaseOpenModeOnly,
	"read_write": AutonomousDwDatabaseOpenModeWrite,
}

// GetAutonomousDwDatabaseOpenModeEnumValues Enumerates the set of values for AutonomousDwDatabaseOpenModeEnum
func GetAutonomousDwDatabaseOpenModeEnumValues() []AutonomousDwDatabaseOpenModeEnum {
	values := make([]AutonomousDwDatabaseOpenModeEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseOpenModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseOpenModeEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseOpenModeEnum
func GetAutonomousDwDatabaseOpenModeEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
	}
}

// GetMappingAutonomousDwDatabaseOpenModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseOpenModeEnum(val string) (AutonomousDwDatabaseOpenModeEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseOpenModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseRefreshableStatusEnum Enum with underlying type: string
type AutonomousDwDatabaseRefreshableStatusEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseRefreshableStatusEnum
const (
	AutonomousDwDatabaseRefreshableStatusRefreshing    AutonomousDwDatabaseRefreshableStatusEnum = "REFRESHING"
	AutonomousDwDatabaseRefreshableStatusNotRefreshing AutonomousDwDatabaseRefreshableStatusEnum = "NOT_REFRESHING"
)

var mappingAutonomousDwDatabaseRefreshableStatusEnum = map[string]AutonomousDwDatabaseRefreshableStatusEnum{
	"REFRESHING":     AutonomousDwDatabaseRefreshableStatusRefreshing,
	"NOT_REFRESHING": AutonomousDwDatabaseRefreshableStatusNotRefreshing,
}

var mappingAutonomousDwDatabaseRefreshableStatusEnumLowerCase = map[string]AutonomousDwDatabaseRefreshableStatusEnum{
	"refreshing":     AutonomousDwDatabaseRefreshableStatusRefreshing,
	"not_refreshing": AutonomousDwDatabaseRefreshableStatusNotRefreshing,
}

// GetAutonomousDwDatabaseRefreshableStatusEnumValues Enumerates the set of values for AutonomousDwDatabaseRefreshableStatusEnum
func GetAutonomousDwDatabaseRefreshableStatusEnumValues() []AutonomousDwDatabaseRefreshableStatusEnum {
	values := make([]AutonomousDwDatabaseRefreshableStatusEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseRefreshableStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseRefreshableStatusEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseRefreshableStatusEnum
func GetAutonomousDwDatabaseRefreshableStatusEnumStringValues() []string {
	return []string{
		"REFRESHING",
		"NOT_REFRESHING",
	}
}

// GetMappingAutonomousDwDatabaseRefreshableStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseRefreshableStatusEnum(val string) (AutonomousDwDatabaseRefreshableStatusEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseRefreshableStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseRefreshableModeEnum Enum with underlying type: string
type AutonomousDwDatabaseRefreshableModeEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseRefreshableModeEnum
const (
	AutonomousDwDatabaseRefreshableModeAutomatic AutonomousDwDatabaseRefreshableModeEnum = "AUTOMATIC"
	AutonomousDwDatabaseRefreshableModeManual    AutonomousDwDatabaseRefreshableModeEnum = "MANUAL"
)

var mappingAutonomousDwDatabaseRefreshableModeEnum = map[string]AutonomousDwDatabaseRefreshableModeEnum{
	"AUTOMATIC": AutonomousDwDatabaseRefreshableModeAutomatic,
	"MANUAL":    AutonomousDwDatabaseRefreshableModeManual,
}

var mappingAutonomousDwDatabaseRefreshableModeEnumLowerCase = map[string]AutonomousDwDatabaseRefreshableModeEnum{
	"automatic": AutonomousDwDatabaseRefreshableModeAutomatic,
	"manual":    AutonomousDwDatabaseRefreshableModeManual,
}

// GetAutonomousDwDatabaseRefreshableModeEnumValues Enumerates the set of values for AutonomousDwDatabaseRefreshableModeEnum
func GetAutonomousDwDatabaseRefreshableModeEnumValues() []AutonomousDwDatabaseRefreshableModeEnum {
	values := make([]AutonomousDwDatabaseRefreshableModeEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseRefreshableModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseRefreshableModeEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseRefreshableModeEnum
func GetAutonomousDwDatabaseRefreshableModeEnumStringValues() []string {
	return []string{
		"AUTOMATIC",
		"MANUAL",
	}
}

// GetMappingAutonomousDwDatabaseRefreshableModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseRefreshableModeEnum(val string) (AutonomousDwDatabaseRefreshableModeEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseRefreshableModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabasePermissionLevelEnum Enum with underlying type: string
type AutonomousDwDatabasePermissionLevelEnum string

// Set of constants representing the allowable values for AutonomousDwDatabasePermissionLevelEnum
const (
	AutonomousDwDatabasePermissionLevelRestricted   AutonomousDwDatabasePermissionLevelEnum = "RESTRICTED"
	AutonomousDwDatabasePermissionLevelUnrestricted AutonomousDwDatabasePermissionLevelEnum = "UNRESTRICTED"
)

var mappingAutonomousDwDatabasePermissionLevelEnum = map[string]AutonomousDwDatabasePermissionLevelEnum{
	"RESTRICTED":   AutonomousDwDatabasePermissionLevelRestricted,
	"UNRESTRICTED": AutonomousDwDatabasePermissionLevelUnrestricted,
}

var mappingAutonomousDwDatabasePermissionLevelEnumLowerCase = map[string]AutonomousDwDatabasePermissionLevelEnum{
	"restricted":   AutonomousDwDatabasePermissionLevelRestricted,
	"unrestricted": AutonomousDwDatabasePermissionLevelUnrestricted,
}

// GetAutonomousDwDatabasePermissionLevelEnumValues Enumerates the set of values for AutonomousDwDatabasePermissionLevelEnum
func GetAutonomousDwDatabasePermissionLevelEnumValues() []AutonomousDwDatabasePermissionLevelEnum {
	values := make([]AutonomousDwDatabasePermissionLevelEnum, 0)
	for _, v := range mappingAutonomousDwDatabasePermissionLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabasePermissionLevelEnumStringValues Enumerates the set of values in String for AutonomousDwDatabasePermissionLevelEnum
func GetAutonomousDwDatabasePermissionLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"UNRESTRICTED",
	}
}

// GetMappingAutonomousDwDatabasePermissionLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabasePermissionLevelEnum(val string) (AutonomousDwDatabasePermissionLevelEnum, bool) {
	enum, ok := mappingAutonomousDwDatabasePermissionLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseRoleEnum Enum with underlying type: string
type AutonomousDwDatabaseRoleEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseRoleEnum
const (
	AutonomousDwDatabaseRolePrimary         AutonomousDwDatabaseRoleEnum = "PRIMARY"
	AutonomousDwDatabaseRoleStandby         AutonomousDwDatabaseRoleEnum = "STANDBY"
	AutonomousDwDatabaseRoleDisabledStandby AutonomousDwDatabaseRoleEnum = "DISABLED_STANDBY"
	AutonomousDwDatabaseRoleBackupCopy      AutonomousDwDatabaseRoleEnum = "BACKUP_COPY"
	AutonomousDwDatabaseRoleSnapshotStandby AutonomousDwDatabaseRoleEnum = "SNAPSHOT_STANDBY"
)

var mappingAutonomousDwDatabaseRoleEnum = map[string]AutonomousDwDatabaseRoleEnum{
	"PRIMARY":          AutonomousDwDatabaseRolePrimary,
	"STANDBY":          AutonomousDwDatabaseRoleStandby,
	"DISABLED_STANDBY": AutonomousDwDatabaseRoleDisabledStandby,
	"BACKUP_COPY":      AutonomousDwDatabaseRoleBackupCopy,
	"SNAPSHOT_STANDBY": AutonomousDwDatabaseRoleSnapshotStandby,
}

var mappingAutonomousDwDatabaseRoleEnumLowerCase = map[string]AutonomousDwDatabaseRoleEnum{
	"primary":          AutonomousDwDatabaseRolePrimary,
	"standby":          AutonomousDwDatabaseRoleStandby,
	"disabled_standby": AutonomousDwDatabaseRoleDisabledStandby,
	"backup_copy":      AutonomousDwDatabaseRoleBackupCopy,
	"snapshot_standby": AutonomousDwDatabaseRoleSnapshotStandby,
}

// GetAutonomousDwDatabaseRoleEnumValues Enumerates the set of values for AutonomousDwDatabaseRoleEnum
func GetAutonomousDwDatabaseRoleEnumValues() []AutonomousDwDatabaseRoleEnum {
	values := make([]AutonomousDwDatabaseRoleEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseRoleEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseRoleEnum
func GetAutonomousDwDatabaseRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
		"BACKUP_COPY",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingAutonomousDwDatabaseRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseRoleEnum(val string) (AutonomousDwDatabaseRoleEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseDataguardRegionTypeEnum Enum with underlying type: string
type AutonomousDwDatabaseDataguardRegionTypeEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseDataguardRegionTypeEnum
const (
	AutonomousDwDatabaseDataguardRegionTypePrimaryDgRegion       AutonomousDwDatabaseDataguardRegionTypeEnum = "PRIMARY_DG_REGION"
	AutonomousDwDatabaseDataguardRegionTypeRemoteStandbyDgRegion AutonomousDwDatabaseDataguardRegionTypeEnum = "REMOTE_STANDBY_DG_REGION"
)

var mappingAutonomousDwDatabaseDataguardRegionTypeEnum = map[string]AutonomousDwDatabaseDataguardRegionTypeEnum{
	"PRIMARY_DG_REGION":        AutonomousDwDatabaseDataguardRegionTypePrimaryDgRegion,
	"REMOTE_STANDBY_DG_REGION": AutonomousDwDatabaseDataguardRegionTypeRemoteStandbyDgRegion,
}

var mappingAutonomousDwDatabaseDataguardRegionTypeEnumLowerCase = map[string]AutonomousDwDatabaseDataguardRegionTypeEnum{
	"primary_dg_region":        AutonomousDwDatabaseDataguardRegionTypePrimaryDgRegion,
	"remote_standby_dg_region": AutonomousDwDatabaseDataguardRegionTypeRemoteStandbyDgRegion,
}

// GetAutonomousDwDatabaseDataguardRegionTypeEnumValues Enumerates the set of values for AutonomousDwDatabaseDataguardRegionTypeEnum
func GetAutonomousDwDatabaseDataguardRegionTypeEnumValues() []AutonomousDwDatabaseDataguardRegionTypeEnum {
	values := make([]AutonomousDwDatabaseDataguardRegionTypeEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseDataguardRegionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseDataguardRegionTypeEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseDataguardRegionTypeEnum
func GetAutonomousDwDatabaseDataguardRegionTypeEnumStringValues() []string {
	return []string{
		"PRIMARY_DG_REGION",
		"REMOTE_STANDBY_DG_REGION",
	}
}

// GetMappingAutonomousDwDatabaseDataguardRegionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseDataguardRegionTypeEnum(val string) (AutonomousDwDatabaseDataguardRegionTypeEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseDataguardRegionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum Enum with underlying type: string
type AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum
const (
	AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEarly   AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum = "EARLY"
	AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeRegular AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum = "REGULAR"
)

var mappingAutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum = map[string]AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum{
	"EARLY":   AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEarly,
	"REGULAR": AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeRegular,
}

var mappingAutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnumLowerCase = map[string]AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum{
	"early":   AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEarly,
	"regular": AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeRegular,
}

// GetAutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnumValues Enumerates the set of values for AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum
func GetAutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnumValues() []AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum {
	values := make([]AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum
func GetAutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnumStringValues() []string {
	return []string{
		"EARLY",
		"REGULAR",
	}
}

// GetMappingAutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum(val string) (AutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseAutonomousMaintenanceScheduleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseDatabaseEditionEnum Enum with underlying type: string
type AutonomousDwDatabaseDatabaseEditionEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseDatabaseEditionEnum
const (
	AutonomousDwDatabaseDatabaseEditionStandardEdition   AutonomousDwDatabaseDatabaseEditionEnum = "STANDARD_EDITION"
	AutonomousDwDatabaseDatabaseEditionEnterpriseEdition AutonomousDwDatabaseDatabaseEditionEnum = "ENTERPRISE_EDITION"
)

var mappingAutonomousDwDatabaseDatabaseEditionEnum = map[string]AutonomousDwDatabaseDatabaseEditionEnum{
	"STANDARD_EDITION":   AutonomousDwDatabaseDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION": AutonomousDwDatabaseDatabaseEditionEnterpriseEdition,
}

var mappingAutonomousDwDatabaseDatabaseEditionEnumLowerCase = map[string]AutonomousDwDatabaseDatabaseEditionEnum{
	"standard_edition":   AutonomousDwDatabaseDatabaseEditionStandardEdition,
	"enterprise_edition": AutonomousDwDatabaseDatabaseEditionEnterpriseEdition,
}

// GetAutonomousDwDatabaseDatabaseEditionEnumValues Enumerates the set of values for AutonomousDwDatabaseDatabaseEditionEnum
func GetAutonomousDwDatabaseDatabaseEditionEnumValues() []AutonomousDwDatabaseDatabaseEditionEnum {
	values := make([]AutonomousDwDatabaseDatabaseEditionEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseDatabaseEditionEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseDatabaseEditionEnum
func GetAutonomousDwDatabaseDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
	}
}

// GetMappingAutonomousDwDatabaseDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseDatabaseEditionEnum(val string) (AutonomousDwDatabaseDatabaseEditionEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseDisasterRecoveryRegionTypeEnum Enum with underlying type: string
type AutonomousDwDatabaseDisasterRecoveryRegionTypeEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseDisasterRecoveryRegionTypeEnum
const (
	AutonomousDwDatabaseDisasterRecoveryRegionTypePrimary AutonomousDwDatabaseDisasterRecoveryRegionTypeEnum = "PRIMARY"
	AutonomousDwDatabaseDisasterRecoveryRegionTypeRemote  AutonomousDwDatabaseDisasterRecoveryRegionTypeEnum = "REMOTE"
)

var mappingAutonomousDwDatabaseDisasterRecoveryRegionTypeEnum = map[string]AutonomousDwDatabaseDisasterRecoveryRegionTypeEnum{
	"PRIMARY": AutonomousDwDatabaseDisasterRecoveryRegionTypePrimary,
	"REMOTE":  AutonomousDwDatabaseDisasterRecoveryRegionTypeRemote,
}

var mappingAutonomousDwDatabaseDisasterRecoveryRegionTypeEnumLowerCase = map[string]AutonomousDwDatabaseDisasterRecoveryRegionTypeEnum{
	"primary": AutonomousDwDatabaseDisasterRecoveryRegionTypePrimary,
	"remote":  AutonomousDwDatabaseDisasterRecoveryRegionTypeRemote,
}

// GetAutonomousDwDatabaseDisasterRecoveryRegionTypeEnumValues Enumerates the set of values for AutonomousDwDatabaseDisasterRecoveryRegionTypeEnum
func GetAutonomousDwDatabaseDisasterRecoveryRegionTypeEnumValues() []AutonomousDwDatabaseDisasterRecoveryRegionTypeEnum {
	values := make([]AutonomousDwDatabaseDisasterRecoveryRegionTypeEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseDisasterRecoveryRegionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseDisasterRecoveryRegionTypeEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseDisasterRecoveryRegionTypeEnum
func GetAutonomousDwDatabaseDisasterRecoveryRegionTypeEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"REMOTE",
	}
}

// GetMappingAutonomousDwDatabaseDisasterRecoveryRegionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseDisasterRecoveryRegionTypeEnum(val string) (AutonomousDwDatabaseDisasterRecoveryRegionTypeEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseDisasterRecoveryRegionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseNetServicesArchitectureEnum Enum with underlying type: string
type AutonomousDwDatabaseNetServicesArchitectureEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseNetServicesArchitectureEnum
const (
	AutonomousDwDatabaseNetServicesArchitectureDedicated AutonomousDwDatabaseNetServicesArchitectureEnum = "DEDICATED"
	AutonomousDwDatabaseNetServicesArchitectureShared    AutonomousDwDatabaseNetServicesArchitectureEnum = "SHARED"
)

var mappingAutonomousDwDatabaseNetServicesArchitectureEnum = map[string]AutonomousDwDatabaseNetServicesArchitectureEnum{
	"DEDICATED": AutonomousDwDatabaseNetServicesArchitectureDedicated,
	"SHARED":    AutonomousDwDatabaseNetServicesArchitectureShared,
}

var mappingAutonomousDwDatabaseNetServicesArchitectureEnumLowerCase = map[string]AutonomousDwDatabaseNetServicesArchitectureEnum{
	"dedicated": AutonomousDwDatabaseNetServicesArchitectureDedicated,
	"shared":    AutonomousDwDatabaseNetServicesArchitectureShared,
}

// GetAutonomousDwDatabaseNetServicesArchitectureEnumValues Enumerates the set of values for AutonomousDwDatabaseNetServicesArchitectureEnum
func GetAutonomousDwDatabaseNetServicesArchitectureEnumValues() []AutonomousDwDatabaseNetServicesArchitectureEnum {
	values := make([]AutonomousDwDatabaseNetServicesArchitectureEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseNetServicesArchitectureEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseNetServicesArchitectureEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseNetServicesArchitectureEnum
func GetAutonomousDwDatabaseNetServicesArchitectureEnumStringValues() []string {
	return []string{
		"DEDICATED",
		"SHARED",
	}
}

// GetMappingAutonomousDwDatabaseNetServicesArchitectureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseNetServicesArchitectureEnum(val string) (AutonomousDwDatabaseNetServicesArchitectureEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseNetServicesArchitectureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDwDatabaseCloneTypeEnum Enum with underlying type: string
type AutonomousDwDatabaseCloneTypeEnum string

// Set of constants representing the allowable values for AutonomousDwDatabaseCloneTypeEnum
const (
	AutonomousDwDatabaseCloneTypeFull     AutonomousDwDatabaseCloneTypeEnum = "FULL"
	AutonomousDwDatabaseCloneTypeMetadata AutonomousDwDatabaseCloneTypeEnum = "METADATA"
	AutonomousDwDatabaseCloneTypePartial  AutonomousDwDatabaseCloneTypeEnum = "PARTIAL"
)

var mappingAutonomousDwDatabaseCloneTypeEnum = map[string]AutonomousDwDatabaseCloneTypeEnum{
	"FULL":     AutonomousDwDatabaseCloneTypeFull,
	"METADATA": AutonomousDwDatabaseCloneTypeMetadata,
	"PARTIAL":  AutonomousDwDatabaseCloneTypePartial,
}

var mappingAutonomousDwDatabaseCloneTypeEnumLowerCase = map[string]AutonomousDwDatabaseCloneTypeEnum{
	"full":     AutonomousDwDatabaseCloneTypeFull,
	"metadata": AutonomousDwDatabaseCloneTypeMetadata,
	"partial":  AutonomousDwDatabaseCloneTypePartial,
}

// GetAutonomousDwDatabaseCloneTypeEnumValues Enumerates the set of values for AutonomousDwDatabaseCloneTypeEnum
func GetAutonomousDwDatabaseCloneTypeEnumValues() []AutonomousDwDatabaseCloneTypeEnum {
	values := make([]AutonomousDwDatabaseCloneTypeEnum, 0)
	for _, v := range mappingAutonomousDwDatabaseCloneTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDwDatabaseCloneTypeEnumStringValues Enumerates the set of values in String for AutonomousDwDatabaseCloneTypeEnum
func GetAutonomousDwDatabaseCloneTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"METADATA",
		"PARTIAL",
	}
}

// GetMappingAutonomousDwDatabaseCloneTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDwDatabaseCloneTypeEnum(val string) (AutonomousDwDatabaseCloneTypeEnum, bool) {
	enum, ok := mappingAutonomousDwDatabaseCloneTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
