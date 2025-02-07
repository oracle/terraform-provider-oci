// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// AutonomousDatabaseSummary An Oracle Autonomous Database.
//
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

	// The quantity of data in the database, in terabytes.
	// The following points apply to Autonomous Databases on Serverless Infrastructure:
	// - This is an integer field whose value remains null when the data size is in GBs and cannot be converted to TBs (by dividing the GB value by 1024) without rounding error.
	// - To get the exact value of data storage size without rounding error, please see `dataStorageSizeInGBs` of Autonomous Database.
	DataStorageSizeInTBs *int `mandatory:"true" json:"dataStorageSizeInTBs"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// Information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
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
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
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
	ComputeModel AutonomousDatabaseSummaryComputeModelEnum `mandatory:"false" json:"computeModel,omitempty"`

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

	// The amount of memory (in GBs) enabled per ECPU or OCPU.
	MemoryPerOracleComputeUnitInGBs *int `mandatory:"false" json:"memoryPerOracleComputeUnitInGBs"`

	// The quantity of data in the database, in gigabytes.
	// For Autonomous Transaction Processing databases using ECPUs on Serverless Infrastructure, this value is always populated. In all the other cases, this value will be null and `dataStorageSizeInTBs` will be populated instead.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The storage space consumed by Autonomous Database in GBs.
	UsedDataStorageSizeInGBs *int `mandatory:"false" json:"usedDataStorageSizeInGBs"`

	// The infrastructure type this resource belongs to.
	InfrastructureType AutonomousDatabaseSummaryInfrastructureTypeEnum `mandatory:"false" json:"infrastructureType,omitempty"`

	// True if the database uses dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html).
	IsDedicated *bool `mandatory:"false" json:"isDedicated"`

	// The Autonomous Container Database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). Used only by Autonomous Database on Dedicated Exadata Infrastructure.
	AutonomousContainerDatabaseId *string `mandatory:"false" json:"autonomousContainerDatabaseId"`

	// Indicates if the Autonomous Database is backup retention locked.
	IsBackupRetentionLocked *bool `mandatory:"false" json:"isBackupRetentionLocked"`

	// The date and time the Autonomous Database was most recently undeleted.
	TimeUndeleted *common.SDKTime `mandatory:"false" json:"timeUndeleted"`

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
	LicenseModel AutonomousDatabaseSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The maximum number of CPUs allowed with a Bring Your Own License (BYOL), including those used for auto-scaling, disaster recovery, tools, etc. Any CPU usage above this limit is considered as License Included and billed.
	ByolComputeCountLimit *float32 `mandatory:"false" json:"byolComputeCountLimit"`

	// The amount of storage that has been used for Autonomous Databases in dedicated infrastructure, in terabytes.
	UsedDataStorageSizeInTBs *int `mandatory:"false" json:"usedDataStorageSizeInTBs"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Security Attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// - For Autonomous Database, setting this will disable public secure access to the database.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and the backup subnet.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
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
	DbWorkload AutonomousDatabaseSummaryDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

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
	DataSafeStatus AutonomousDatabaseSummaryDataSafeStatusEnum `mandatory:"false" json:"dataSafeStatus,omitempty"`

	// Status of Operations Insights for this Autonomous Database.
	OperationsInsightsStatus AutonomousDatabaseSummaryOperationsInsightsStatusEnum `mandatory:"false" json:"operationsInsightsStatus,omitempty"`

	// Status of Database Management for this Autonomous Database.
	DatabaseManagementStatus AutonomousDatabaseSummaryDatabaseManagementStatusEnum `mandatory:"false" json:"databaseManagementStatus,omitempty"`

	// The date and time when maintenance will begin.
	TimeMaintenanceBegin *common.SDKTime `mandatory:"false" json:"timeMaintenanceBegin"`

	// The date and time when maintenance will end.
	TimeMaintenanceEnd *common.SDKTime `mandatory:"false" json:"timeMaintenanceEnd"`

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
	OpenMode AutonomousDatabaseSummaryOpenModeEnum `mandatory:"false" json:"openMode,omitempty"`

	// The refresh status of the clone. REFRESHING indicates that the clone is currently being refreshed with data from the source Autonomous Database.
	RefreshableStatus AutonomousDatabaseSummaryRefreshableStatusEnum `mandatory:"false" json:"refreshableStatus,omitempty"`

	// The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.
	RefreshableMode AutonomousDatabaseSummaryRefreshableModeEnum `mandatory:"false" json:"refreshableMode,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that was cloned to create the current Autonomous Database.
	SourceId *string `mandatory:"false" json:"sourceId"`

	// The Autonomous Database permission level. Restricted mode allows access only by admin users.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	PermissionLevel AutonomousDatabaseSummaryPermissionLevelEnum `mandatory:"false" json:"permissionLevel,omitempty"`

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
	Role AutonomousDatabaseSummaryRoleEnum `mandatory:"false" json:"role,omitempty"`

	// List of Oracle Database versions available for a database upgrade. If there are no version upgrades available, this list is empty.
	AvailableUpgradeVersions []string `mandatory:"false" json:"availableUpgradeVersions"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
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
	DataguardRegionType AutonomousDatabaseSummaryDataguardRegionTypeEnum `mandatory:"false" json:"dataguardRegionType,omitempty"`

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

	// The unique identifier for leader autonomous database OCID OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ResourcePoolLeaderId *string `mandatory:"false" json:"resourcePoolLeaderId"`

	ResourcePoolSummary *ResourcePoolSummary `mandatory:"false" json:"resourcePoolSummary"`

	// Indicates if the refreshable clone can be reconnected to its source database.
	IsReconnectCloneEnabled *bool `mandatory:"false" json:"isReconnectCloneEnabled"`

	// The time and date as an RFC3339 formatted string, e.g., 2022-01-01T12:00:00.000Z, to set the limit for a refreshable clone to be reconnected to its source database.
	TimeUntilReconnectCloneEnabled *common.SDKTime `mandatory:"false" json:"timeUntilReconnectCloneEnabled"`

	// The maintenance schedule type of the Autonomous Database Serverless. An EARLY maintenance schedule
	// follows a schedule applying patches prior to the REGULAR schedule. A REGULAR maintenance schedule follows the normal cycle
	AutonomousMaintenanceScheduleType AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum `mandatory:"false" json:"autonomousMaintenanceScheduleType,omitempty"`

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
	DatabaseEdition AutonomousDatabaseSummaryDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

	// The list of database tools details.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, isLocalDataGuardEnabled, or isFreeTier.
	DbToolsDetails []DatabaseTool `mandatory:"false" json:"dbToolsDetails"`

	// Indicates the local disaster recovery (DR) type of the Autonomous Database Serverless instance.
	// Autonomous Data Guard (ADG) DR type provides business critical DR with a faster recovery time objective (RTO) during failover or switchover.
	// Backup-based DR type provides lower cost DR with a slower RTO during failover or switchover.
	LocalDisasterRecoveryType DisasterRecoveryConfigurationDisasterRecoveryTypeEnum `mandatory:"false" json:"localDisasterRecoveryType,omitempty"`

	// **Deprecated.** The disaster recovery (DR) region type of the Autonomous Database. For Autonomous Database Serverless instances, DR associations have designated primary and standby regions. These region types do not change when the database changes roles. The standby region in DR associations can be the same region as the primary region, or they can be in a remote regions. Some database administration operations may be available only in the primary region of the DR association, and cannot be performed when the database using the primary role is operating in a remote region.
	DisasterRecoveryRegionType AutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum `mandatory:"false" json:"disasterRecoveryRegionType,omitempty"`

	// The date and time the Disaster Recovery role was switched for the standby Autonomous Database.
	TimeDisasterRecoveryRoleChanged *common.SDKTime `mandatory:"false" json:"timeDisasterRecoveryRoleChanged"`

	RemoteDisasterRecoveryConfiguration *DisasterRecoveryConfiguration `mandatory:"false" json:"remoteDisasterRecoveryConfiguration"`

	// Enabling SHARED server architecture enables a database server to allow many client processes to share very few server processes, thereby increasing the number of supported users.
	NetServicesArchitecture AutonomousDatabaseSummaryNetServicesArchitectureEnum `mandatory:"false" json:"netServicesArchitecture,omitempty"`

	// The availability domain where the Autonomous Database Serverless instance is located.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cluster placement group of the Autonomous Serverless Database.
	ClusterPlacementGroupId *string `mandatory:"false" json:"clusterPlacementGroupId"`

	// A list of the source Autonomous Database's table space number(s) used to create this partial clone from the backup.
	CloneTableSpaceList []int `mandatory:"false" json:"cloneTableSpaceList"`
}

func (m AutonomousDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousDatabaseSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDatabaseSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousDatabaseSummaryComputeModelEnum(string(m.ComputeModel)); !ok && m.ComputeModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeModel: %s. Supported values are: %s.", m.ComputeModel, strings.Join(GetAutonomousDatabaseSummaryComputeModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryInfrastructureTypeEnum(string(m.InfrastructureType)); !ok && m.InfrastructureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InfrastructureType: %s. Supported values are: %s.", m.InfrastructureType, strings.Join(GetAutonomousDatabaseSummaryInfrastructureTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetAutonomousDatabaseSummaryLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetAutonomousDatabaseSummaryDbWorkloadEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryDataSafeStatusEnum(string(m.DataSafeStatus)); !ok && m.DataSafeStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataSafeStatus: %s. Supported values are: %s.", m.DataSafeStatus, strings.Join(GetAutonomousDatabaseSummaryDataSafeStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryOperationsInsightsStatusEnum(string(m.OperationsInsightsStatus)); !ok && m.OperationsInsightsStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationsInsightsStatus: %s. Supported values are: %s.", m.OperationsInsightsStatus, strings.Join(GetAutonomousDatabaseSummaryOperationsInsightsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryDatabaseManagementStatusEnum(string(m.DatabaseManagementStatus)); !ok && m.DatabaseManagementStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseManagementStatus: %s. Supported values are: %s.", m.DatabaseManagementStatus, strings.Join(GetAutonomousDatabaseSummaryDatabaseManagementStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryOpenModeEnum(string(m.OpenMode)); !ok && m.OpenMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OpenMode: %s. Supported values are: %s.", m.OpenMode, strings.Join(GetAutonomousDatabaseSummaryOpenModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryRefreshableStatusEnum(string(m.RefreshableStatus)); !ok && m.RefreshableStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RefreshableStatus: %s. Supported values are: %s.", m.RefreshableStatus, strings.Join(GetAutonomousDatabaseSummaryRefreshableStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryRefreshableModeEnum(string(m.RefreshableMode)); !ok && m.RefreshableMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RefreshableMode: %s. Supported values are: %s.", m.RefreshableMode, strings.Join(GetAutonomousDatabaseSummaryRefreshableModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryPermissionLevelEnum(string(m.PermissionLevel)); !ok && m.PermissionLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PermissionLevel: %s. Supported values are: %s.", m.PermissionLevel, strings.Join(GetAutonomousDatabaseSummaryPermissionLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetAutonomousDatabaseSummaryRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryDataguardRegionTypeEnum(string(m.DataguardRegionType)); !ok && m.DataguardRegionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataguardRegionType: %s. Supported values are: %s.", m.DataguardRegionType, strings.Join(GetAutonomousDatabaseSummaryDataguardRegionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum(string(m.AutonomousMaintenanceScheduleType)); !ok && m.AutonomousMaintenanceScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutonomousMaintenanceScheduleType: %s. Supported values are: %s.", m.AutonomousMaintenanceScheduleType, strings.Join(GetAutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetAutonomousDatabaseSummaryDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDisasterRecoveryConfigurationDisasterRecoveryTypeEnum(string(m.LocalDisasterRecoveryType)); !ok && m.LocalDisasterRecoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LocalDisasterRecoveryType: %s. Supported values are: %s.", m.LocalDisasterRecoveryType, strings.Join(GetDisasterRecoveryConfigurationDisasterRecoveryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum(string(m.DisasterRecoveryRegionType)); !ok && m.DisasterRecoveryRegionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DisasterRecoveryRegionType: %s. Supported values are: %s.", m.DisasterRecoveryRegionType, strings.Join(GetAutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryNetServicesArchitectureEnum(string(m.NetServicesArchitecture)); !ok && m.NetServicesArchitecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetServicesArchitecture: %s. Supported values are: %s.", m.NetServicesArchitecture, strings.Join(GetAutonomousDatabaseSummaryNetServicesArchitectureEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AutonomousDatabaseSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SubscriptionId                          *string                                                        `json:"subscriptionId"`
		LifecycleDetails                        *string                                                        `json:"lifecycleDetails"`
		KmsKeyId                                *string                                                        `json:"kmsKeyId"`
		VaultId                                 *string                                                        `json:"vaultId"`
		KmsKeyLifecycleDetails                  *string                                                        `json:"kmsKeyLifecycleDetails"`
		EncryptionKey                           autonomousdatabaseencryptionkeydetails                         `json:"encryptionKey"`
		KmsKeyVersionId                         *string                                                        `json:"kmsKeyVersionId"`
		CharacterSet                            *string                                                        `json:"characterSet"`
		NcharacterSet                           *string                                                        `json:"ncharacterSet"`
		InMemoryPercentage                      *int                                                           `json:"inMemoryPercentage"`
		InMemoryAreaInGBs                       *int                                                           `json:"inMemoryAreaInGBs"`
		NextLongTermBackupTimeStamp             *common.SDKTime                                                `json:"nextLongTermBackupTimeStamp"`
		LongTermBackupSchedule                  *LongTermBackUpScheduleDetails                                 `json:"longTermBackupSchedule"`
		IsFreeTier                              *bool                                                          `json:"isFreeTier"`
		SystemTags                              map[string]map[string]interface{}                              `json:"systemTags"`
		TimeReclamationOfFreeAutonomousDatabase *common.SDKTime                                                `json:"timeReclamationOfFreeAutonomousDatabase"`
		TimeDeletionOfFreeAutonomousDatabase    *common.SDKTime                                                `json:"timeDeletionOfFreeAutonomousDatabase"`
		BackupConfig                            *AutonomousDatabaseBackupConfig                                `json:"backupConfig"`
		KeyHistoryEntry                         []AutonomousDatabaseKeyHistoryEntry                            `json:"keyHistoryEntry"`
		EncryptionKeyHistoryEntry               []AutonomousDatabaseEncryptionKeyHistoryEntry                  `json:"encryptionKeyHistoryEntry"`
		CpuCoreCount                            *int                                                           `json:"cpuCoreCount"`
		LocalAdgAutoFailoverMaxDataLossLimit    *int                                                           `json:"localAdgAutoFailoverMaxDataLossLimit"`
		ComputeModel                            AutonomousDatabaseSummaryComputeModelEnum                      `json:"computeModel"`
		ComputeCount                            *float32                                                       `json:"computeCount"`
		BackupRetentionPeriodInDays             *int                                                           `json:"backupRetentionPeriodInDays"`
		TotalBackupStorageSizeInGBs             *float64                                                       `json:"totalBackupStorageSizeInGBs"`
		OcpuCount                               *float32                                                       `json:"ocpuCount"`
		ProvisionableCpus                       []float32                                                      `json:"provisionableCpus"`
		MemoryPerOracleComputeUnitInGBs         *int                                                           `json:"memoryPerOracleComputeUnitInGBs"`
		DataStorageSizeInGBs                    *int                                                           `json:"dataStorageSizeInGBs"`
		UsedDataStorageSizeInGBs                *int                                                           `json:"usedDataStorageSizeInGBs"`
		InfrastructureType                      AutonomousDatabaseSummaryInfrastructureTypeEnum                `json:"infrastructureType"`
		IsDedicated                             *bool                                                          `json:"isDedicated"`
		AutonomousContainerDatabaseId           *string                                                        `json:"autonomousContainerDatabaseId"`
		IsBackupRetentionLocked                 *bool                                                          `json:"isBackupRetentionLocked"`
		TimeUndeleted                           *common.SDKTime                                                `json:"timeUndeleted"`
		TimeCreated                             *common.SDKTime                                                `json:"timeCreated"`
		DisplayName                             *string                                                        `json:"displayName"`
		ServiceConsoleUrl                       *string                                                        `json:"serviceConsoleUrl"`
		ConnectionStrings                       *AutonomousDatabaseConnectionStrings                           `json:"connectionStrings"`
		ConnectionUrls                          *AutonomousDatabaseConnectionUrls                              `json:"connectionUrls"`
		PublicConnectionUrls                    *AutonomousDatabaseConnectionUrls                              `json:"publicConnectionUrls"`
		LicenseModel                            AutonomousDatabaseSummaryLicenseModelEnum                      `json:"licenseModel"`
		ByolComputeCountLimit                   *float32                                                       `json:"byolComputeCountLimit"`
		UsedDataStorageSizeInTBs                *int                                                           `json:"usedDataStorageSizeInTBs"`
		FreeformTags                            map[string]string                                              `json:"freeformTags"`
		DefinedTags                             map[string]map[string]interface{}                              `json:"definedTags"`
		SecurityAttributes                      map[string]map[string]interface{}                              `json:"securityAttributes"`
		SubnetId                                *string                                                        `json:"subnetId"`
		NsgIds                                  []string                                                       `json:"nsgIds"`
		PrivateEndpoint                         *string                                                        `json:"privateEndpoint"`
		PublicEndpoint                          *string                                                        `json:"publicEndpoint"`
		PrivateEndpointLabel                    *string                                                        `json:"privateEndpointLabel"`
		PrivateEndpointIp                       *string                                                        `json:"privateEndpointIp"`
		DbVersion                               *string                                                        `json:"dbVersion"`
		IsPreview                               *bool                                                          `json:"isPreview"`
		DbWorkload                              AutonomousDatabaseSummaryDbWorkloadEnum                        `json:"dbWorkload"`
		IsDevTier                               *bool                                                          `json:"isDevTier"`
		IsAccessControlEnabled                  *bool                                                          `json:"isAccessControlEnabled"`
		WhitelistedIps                          []string                                                       `json:"whitelistedIps"`
		ArePrimaryWhitelistedIpsUsed            *bool                                                          `json:"arePrimaryWhitelistedIpsUsed"`
		StandbyWhitelistedIps                   []string                                                       `json:"standbyWhitelistedIps"`
		ApexDetails                             *AutonomousDatabaseApex                                        `json:"apexDetails"`
		IsAutoScalingEnabled                    *bool                                                          `json:"isAutoScalingEnabled"`
		DataSafeStatus                          AutonomousDatabaseSummaryDataSafeStatusEnum                    `json:"dataSafeStatus"`
		OperationsInsightsStatus                AutonomousDatabaseSummaryOperationsInsightsStatusEnum          `json:"operationsInsightsStatus"`
		DatabaseManagementStatus                AutonomousDatabaseSummaryDatabaseManagementStatusEnum          `json:"databaseManagementStatus"`
		TimeMaintenanceBegin                    *common.SDKTime                                                `json:"timeMaintenanceBegin"`
		TimeMaintenanceEnd                      *common.SDKTime                                                `json:"timeMaintenanceEnd"`
		IsRefreshableClone                      *bool                                                          `json:"isRefreshableClone"`
		TimeOfLastRefresh                       *common.SDKTime                                                `json:"timeOfLastRefresh"`
		TimeOfLastRefreshPoint                  *common.SDKTime                                                `json:"timeOfLastRefreshPoint"`
		TimeOfNextRefresh                       *common.SDKTime                                                `json:"timeOfNextRefresh"`
		OpenMode                                AutonomousDatabaseSummaryOpenModeEnum                          `json:"openMode"`
		RefreshableStatus                       AutonomousDatabaseSummaryRefreshableStatusEnum                 `json:"refreshableStatus"`
		RefreshableMode                         AutonomousDatabaseSummaryRefreshableModeEnum                   `json:"refreshableMode"`
		SourceId                                *string                                                        `json:"sourceId"`
		PermissionLevel                         AutonomousDatabaseSummaryPermissionLevelEnum                   `json:"permissionLevel"`
		TimeOfLastSwitchover                    *common.SDKTime                                                `json:"timeOfLastSwitchover"`
		TimeOfLastFailover                      *common.SDKTime                                                `json:"timeOfLastFailover"`
		IsDataGuardEnabled                      *bool                                                          `json:"isDataGuardEnabled"`
		FailedDataRecoveryInSeconds             *int                                                           `json:"failedDataRecoveryInSeconds"`
		StandbyDb                               *AutonomousDatabaseStandbySummary                              `json:"standbyDb"`
		IsLocalDataGuardEnabled                 *bool                                                          `json:"isLocalDataGuardEnabled"`
		IsRemoteDataGuardEnabled                *bool                                                          `json:"isRemoteDataGuardEnabled"`
		LocalStandbyDb                          *AutonomousDatabaseStandbySummary                              `json:"localStandbyDb"`
		Role                                    AutonomousDatabaseSummaryRoleEnum                              `json:"role"`
		AvailableUpgradeVersions                []string                                                       `json:"availableUpgradeVersions"`
		KeyStoreId                              *string                                                        `json:"keyStoreId"`
		KeyStoreWalletName                      *string                                                        `json:"keyStoreWalletName"`
		AutoRefreshFrequencyInSeconds           *int                                                           `json:"autoRefreshFrequencyInSeconds"`
		AutoRefreshPointLagInSeconds            *int                                                           `json:"autoRefreshPointLagInSeconds"`
		TimeOfAutoRefreshStart                  *common.SDKTime                                                `json:"timeOfAutoRefreshStart"`
		SupportedRegionsToCloneTo               []string                                                       `json:"supportedRegionsToCloneTo"`
		CustomerContacts                        []CustomerContact                                              `json:"customerContacts"`
		TimeLocalDataGuardEnabled               *common.SDKTime                                                `json:"timeLocalDataGuardEnabled"`
		DataguardRegionType                     AutonomousDatabaseSummaryDataguardRegionTypeEnum               `json:"dataguardRegionType"`
		TimeDataGuardRoleChanged                *common.SDKTime                                                `json:"timeDataGuardRoleChanged"`
		PeerDbIds                               []string                                                       `json:"peerDbIds"`
		IsMtlsConnectionRequired                *bool                                                          `json:"isMtlsConnectionRequired"`
		TimeOfJoiningResourcePool               *common.SDKTime                                                `json:"timeOfJoiningResourcePool"`
		ResourcePoolLeaderId                    *string                                                        `json:"resourcePoolLeaderId"`
		ResourcePoolSummary                     *ResourcePoolSummary                                           `json:"resourcePoolSummary"`
		IsReconnectCloneEnabled                 *bool                                                          `json:"isReconnectCloneEnabled"`
		TimeUntilReconnectCloneEnabled          *common.SDKTime                                                `json:"timeUntilReconnectCloneEnabled"`
		AutonomousMaintenanceScheduleType       AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum `json:"autonomousMaintenanceScheduleType"`
		ScheduledOperations                     []ScheduledOperationDetails                                    `json:"scheduledOperations"`
		IsAutoScalingForStorageEnabled          *bool                                                          `json:"isAutoScalingForStorageEnabled"`
		AllocatedStorageSizeInTBs               *float64                                                       `json:"allocatedStorageSizeInTBs"`
		ActualUsedDataStorageSizeInTBs          *float64                                                       `json:"actualUsedDataStorageSizeInTBs"`
		DatabaseEdition                         AutonomousDatabaseSummaryDatabaseEditionEnum                   `json:"databaseEdition"`
		DbToolsDetails                          []DatabaseTool                                                 `json:"dbToolsDetails"`
		LocalDisasterRecoveryType               DisasterRecoveryConfigurationDisasterRecoveryTypeEnum          `json:"localDisasterRecoveryType"`
		DisasterRecoveryRegionType              AutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum        `json:"disasterRecoveryRegionType"`
		TimeDisasterRecoveryRoleChanged         *common.SDKTime                                                `json:"timeDisasterRecoveryRoleChanged"`
		RemoteDisasterRecoveryConfiguration     *DisasterRecoveryConfiguration                                 `json:"remoteDisasterRecoveryConfiguration"`
		NetServicesArchitecture                 AutonomousDatabaseSummaryNetServicesArchitectureEnum           `json:"netServicesArchitecture"`
		AvailabilityDomain                      *string                                                        `json:"availabilityDomain"`
		ClusterPlacementGroupId                 *string                                                        `json:"clusterPlacementGroupId"`
		CloneTableSpaceList                     []int                                                          `json:"cloneTableSpaceList"`
		Id                                      *string                                                        `json:"id"`
		CompartmentId                           *string                                                        `json:"compartmentId"`
		LifecycleState                          AutonomousDatabaseSummaryLifecycleStateEnum                    `json:"lifecycleState"`
		DbName                                  *string                                                        `json:"dbName"`
		DataStorageSizeInTBs                    *int                                                           `json:"dataStorageSizeInTBs"`
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

	m.DataStorageSizeInGBs = model.DataStorageSizeInGBs

	m.UsedDataStorageSizeInGBs = model.UsedDataStorageSizeInGBs

	m.InfrastructureType = model.InfrastructureType

	m.IsDedicated = model.IsDedicated

	m.AutonomousContainerDatabaseId = model.AutonomousContainerDatabaseId

	m.IsBackupRetentionLocked = model.IsBackupRetentionLocked

	m.TimeUndeleted = model.TimeUndeleted

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
	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.LifecycleState = model.LifecycleState

	m.DbName = model.DbName

	m.DataStorageSizeInTBs = model.DataStorageSizeInTBs

	return
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
	AutonomousDatabaseSummaryLifecycleStateRecreating              AutonomousDatabaseSummaryLifecycleStateEnum = "RECREATING"
	AutonomousDatabaseSummaryLifecycleStateRoleChangeInProgress    AutonomousDatabaseSummaryLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousDatabaseSummaryLifecycleStateUpgrading               AutonomousDatabaseSummaryLifecycleStateEnum = "UPGRADING"
	AutonomousDatabaseSummaryLifecycleStateInaccessible            AutonomousDatabaseSummaryLifecycleStateEnum = "INACCESSIBLE"
	AutonomousDatabaseSummaryLifecycleStateStandby                 AutonomousDatabaseSummaryLifecycleStateEnum = "STANDBY"
)

var mappingAutonomousDatabaseSummaryLifecycleStateEnum = map[string]AutonomousDatabaseSummaryLifecycleStateEnum{
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
	"RECREATING":                AutonomousDatabaseSummaryLifecycleStateRecreating,
	"ROLE_CHANGE_IN_PROGRESS":   AutonomousDatabaseSummaryLifecycleStateRoleChangeInProgress,
	"UPGRADING":                 AutonomousDatabaseSummaryLifecycleStateUpgrading,
	"INACCESSIBLE":              AutonomousDatabaseSummaryLifecycleStateInaccessible,
	"STANDBY":                   AutonomousDatabaseSummaryLifecycleStateStandby,
}

var mappingAutonomousDatabaseSummaryLifecycleStateEnumLowerCase = map[string]AutonomousDatabaseSummaryLifecycleStateEnum{
	"provisioning":              AutonomousDatabaseSummaryLifecycleStateProvisioning,
	"available":                 AutonomousDatabaseSummaryLifecycleStateAvailable,
	"stopping":                  AutonomousDatabaseSummaryLifecycleStateStopping,
	"stopped":                   AutonomousDatabaseSummaryLifecycleStateStopped,
	"starting":                  AutonomousDatabaseSummaryLifecycleStateStarting,
	"terminating":               AutonomousDatabaseSummaryLifecycleStateTerminating,
	"terminated":                AutonomousDatabaseSummaryLifecycleStateTerminated,
	"unavailable":               AutonomousDatabaseSummaryLifecycleStateUnavailable,
	"restore_in_progress":       AutonomousDatabaseSummaryLifecycleStateRestoreInProgress,
	"restore_failed":            AutonomousDatabaseSummaryLifecycleStateRestoreFailed,
	"backup_in_progress":        AutonomousDatabaseSummaryLifecycleStateBackupInProgress,
	"scale_in_progress":         AutonomousDatabaseSummaryLifecycleStateScaleInProgress,
	"available_needs_attention": AutonomousDatabaseSummaryLifecycleStateAvailableNeedsAttention,
	"updating":                  AutonomousDatabaseSummaryLifecycleStateUpdating,
	"maintenance_in_progress":   AutonomousDatabaseSummaryLifecycleStateMaintenanceInProgress,
	"restarting":                AutonomousDatabaseSummaryLifecycleStateRestarting,
	"recreating":                AutonomousDatabaseSummaryLifecycleStateRecreating,
	"role_change_in_progress":   AutonomousDatabaseSummaryLifecycleStateRoleChangeInProgress,
	"upgrading":                 AutonomousDatabaseSummaryLifecycleStateUpgrading,
	"inaccessible":              AutonomousDatabaseSummaryLifecycleStateInaccessible,
	"standby":                   AutonomousDatabaseSummaryLifecycleStateStandby,
}

// GetAutonomousDatabaseSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseSummaryLifecycleStateEnum
func GetAutonomousDatabaseSummaryLifecycleStateEnumValues() []AutonomousDatabaseSummaryLifecycleStateEnum {
	values := make([]AutonomousDatabaseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryLifecycleStateEnum
func GetAutonomousDatabaseSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingAutonomousDatabaseSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryLifecycleStateEnum(val string) (AutonomousDatabaseSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryComputeModelEnum Enum with underlying type: string
type AutonomousDatabaseSummaryComputeModelEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryComputeModelEnum
const (
	AutonomousDatabaseSummaryComputeModelEcpu AutonomousDatabaseSummaryComputeModelEnum = "ECPU"
	AutonomousDatabaseSummaryComputeModelOcpu AutonomousDatabaseSummaryComputeModelEnum = "OCPU"
)

var mappingAutonomousDatabaseSummaryComputeModelEnum = map[string]AutonomousDatabaseSummaryComputeModelEnum{
	"ECPU": AutonomousDatabaseSummaryComputeModelEcpu,
	"OCPU": AutonomousDatabaseSummaryComputeModelOcpu,
}

var mappingAutonomousDatabaseSummaryComputeModelEnumLowerCase = map[string]AutonomousDatabaseSummaryComputeModelEnum{
	"ecpu": AutonomousDatabaseSummaryComputeModelEcpu,
	"ocpu": AutonomousDatabaseSummaryComputeModelOcpu,
}

// GetAutonomousDatabaseSummaryComputeModelEnumValues Enumerates the set of values for AutonomousDatabaseSummaryComputeModelEnum
func GetAutonomousDatabaseSummaryComputeModelEnumValues() []AutonomousDatabaseSummaryComputeModelEnum {
	values := make([]AutonomousDatabaseSummaryComputeModelEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryComputeModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryComputeModelEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryComputeModelEnum
func GetAutonomousDatabaseSummaryComputeModelEnumStringValues() []string {
	return []string{
		"ECPU",
		"OCPU",
	}
}

// GetMappingAutonomousDatabaseSummaryComputeModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryComputeModelEnum(val string) (AutonomousDatabaseSummaryComputeModelEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryComputeModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryInfrastructureTypeEnum Enum with underlying type: string
type AutonomousDatabaseSummaryInfrastructureTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryInfrastructureTypeEnum
const (
	AutonomousDatabaseSummaryInfrastructureTypeCloud           AutonomousDatabaseSummaryInfrastructureTypeEnum = "CLOUD"
	AutonomousDatabaseSummaryInfrastructureTypeCloudAtCustomer AutonomousDatabaseSummaryInfrastructureTypeEnum = "CLOUD_AT_CUSTOMER"
)

var mappingAutonomousDatabaseSummaryInfrastructureTypeEnum = map[string]AutonomousDatabaseSummaryInfrastructureTypeEnum{
	"CLOUD":             AutonomousDatabaseSummaryInfrastructureTypeCloud,
	"CLOUD_AT_CUSTOMER": AutonomousDatabaseSummaryInfrastructureTypeCloudAtCustomer,
}

var mappingAutonomousDatabaseSummaryInfrastructureTypeEnumLowerCase = map[string]AutonomousDatabaseSummaryInfrastructureTypeEnum{
	"cloud":             AutonomousDatabaseSummaryInfrastructureTypeCloud,
	"cloud_at_customer": AutonomousDatabaseSummaryInfrastructureTypeCloudAtCustomer,
}

// GetAutonomousDatabaseSummaryInfrastructureTypeEnumValues Enumerates the set of values for AutonomousDatabaseSummaryInfrastructureTypeEnum
func GetAutonomousDatabaseSummaryInfrastructureTypeEnumValues() []AutonomousDatabaseSummaryInfrastructureTypeEnum {
	values := make([]AutonomousDatabaseSummaryInfrastructureTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryInfrastructureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryInfrastructureTypeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryInfrastructureTypeEnum
func GetAutonomousDatabaseSummaryInfrastructureTypeEnumStringValues() []string {
	return []string{
		"CLOUD",
		"CLOUD_AT_CUSTOMER",
	}
}

// GetMappingAutonomousDatabaseSummaryInfrastructureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryInfrastructureTypeEnum(val string) (AutonomousDatabaseSummaryInfrastructureTypeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryInfrastructureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryLicenseModelEnum Enum with underlying type: string
type AutonomousDatabaseSummaryLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryLicenseModelEnum
const (
	AutonomousDatabaseSummaryLicenseModelLicenseIncluded     AutonomousDatabaseSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousDatabaseSummaryLicenseModelBringYourOwnLicense AutonomousDatabaseSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousDatabaseSummaryLicenseModelEnum = map[string]AutonomousDatabaseSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousDatabaseSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousDatabaseSummaryLicenseModelBringYourOwnLicense,
}

var mappingAutonomousDatabaseSummaryLicenseModelEnumLowerCase = map[string]AutonomousDatabaseSummaryLicenseModelEnum{
	"license_included":       AutonomousDatabaseSummaryLicenseModelLicenseIncluded,
	"bring_your_own_license": AutonomousDatabaseSummaryLicenseModelBringYourOwnLicense,
}

// GetAutonomousDatabaseSummaryLicenseModelEnumValues Enumerates the set of values for AutonomousDatabaseSummaryLicenseModelEnum
func GetAutonomousDatabaseSummaryLicenseModelEnumValues() []AutonomousDatabaseSummaryLicenseModelEnum {
	values := make([]AutonomousDatabaseSummaryLicenseModelEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryLicenseModelEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryLicenseModelEnum
func GetAutonomousDatabaseSummaryLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingAutonomousDatabaseSummaryLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryLicenseModelEnum(val string) (AutonomousDatabaseSummaryLicenseModelEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryDbWorkloadEnum Enum with underlying type: string
type AutonomousDatabaseSummaryDbWorkloadEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryDbWorkloadEnum
const (
	AutonomousDatabaseSummaryDbWorkloadOltp AutonomousDatabaseSummaryDbWorkloadEnum = "OLTP"
	AutonomousDatabaseSummaryDbWorkloadDw   AutonomousDatabaseSummaryDbWorkloadEnum = "DW"
	AutonomousDatabaseSummaryDbWorkloadAjd  AutonomousDatabaseSummaryDbWorkloadEnum = "AJD"
	AutonomousDatabaseSummaryDbWorkloadApex AutonomousDatabaseSummaryDbWorkloadEnum = "APEX"
)

var mappingAutonomousDatabaseSummaryDbWorkloadEnum = map[string]AutonomousDatabaseSummaryDbWorkloadEnum{
	"OLTP": AutonomousDatabaseSummaryDbWorkloadOltp,
	"DW":   AutonomousDatabaseSummaryDbWorkloadDw,
	"AJD":  AutonomousDatabaseSummaryDbWorkloadAjd,
	"APEX": AutonomousDatabaseSummaryDbWorkloadApex,
}

var mappingAutonomousDatabaseSummaryDbWorkloadEnumLowerCase = map[string]AutonomousDatabaseSummaryDbWorkloadEnum{
	"oltp": AutonomousDatabaseSummaryDbWorkloadOltp,
	"dw":   AutonomousDatabaseSummaryDbWorkloadDw,
	"ajd":  AutonomousDatabaseSummaryDbWorkloadAjd,
	"apex": AutonomousDatabaseSummaryDbWorkloadApex,
}

// GetAutonomousDatabaseSummaryDbWorkloadEnumValues Enumerates the set of values for AutonomousDatabaseSummaryDbWorkloadEnum
func GetAutonomousDatabaseSummaryDbWorkloadEnumValues() []AutonomousDatabaseSummaryDbWorkloadEnum {
	values := make([]AutonomousDatabaseSummaryDbWorkloadEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryDbWorkloadEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryDbWorkloadEnum
func GetAutonomousDatabaseSummaryDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
		"AJD",
		"APEX",
	}
}

// GetMappingAutonomousDatabaseSummaryDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryDbWorkloadEnum(val string) (AutonomousDatabaseSummaryDbWorkloadEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryDbWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingAutonomousDatabaseSummaryDataSafeStatusEnum = map[string]AutonomousDatabaseSummaryDataSafeStatusEnum{
	"REGISTERING":    AutonomousDatabaseSummaryDataSafeStatusRegistering,
	"REGISTERED":     AutonomousDatabaseSummaryDataSafeStatusRegistered,
	"DEREGISTERING":  AutonomousDatabaseSummaryDataSafeStatusDeregistering,
	"NOT_REGISTERED": AutonomousDatabaseSummaryDataSafeStatusNotRegistered,
	"FAILED":         AutonomousDatabaseSummaryDataSafeStatusFailed,
}

var mappingAutonomousDatabaseSummaryDataSafeStatusEnumLowerCase = map[string]AutonomousDatabaseSummaryDataSafeStatusEnum{
	"registering":    AutonomousDatabaseSummaryDataSafeStatusRegistering,
	"registered":     AutonomousDatabaseSummaryDataSafeStatusRegistered,
	"deregistering":  AutonomousDatabaseSummaryDataSafeStatusDeregistering,
	"not_registered": AutonomousDatabaseSummaryDataSafeStatusNotRegistered,
	"failed":         AutonomousDatabaseSummaryDataSafeStatusFailed,
}

// GetAutonomousDatabaseSummaryDataSafeStatusEnumValues Enumerates the set of values for AutonomousDatabaseSummaryDataSafeStatusEnum
func GetAutonomousDatabaseSummaryDataSafeStatusEnumValues() []AutonomousDatabaseSummaryDataSafeStatusEnum {
	values := make([]AutonomousDatabaseSummaryDataSafeStatusEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryDataSafeStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryDataSafeStatusEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryDataSafeStatusEnum
func GetAutonomousDatabaseSummaryDataSafeStatusEnumStringValues() []string {
	return []string{
		"REGISTERING",
		"REGISTERED",
		"DEREGISTERING",
		"NOT_REGISTERED",
		"FAILED",
	}
}

// GetMappingAutonomousDatabaseSummaryDataSafeStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryDataSafeStatusEnum(val string) (AutonomousDatabaseSummaryDataSafeStatusEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryDataSafeStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryOperationsInsightsStatusEnum Enum with underlying type: string
type AutonomousDatabaseSummaryOperationsInsightsStatusEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryOperationsInsightsStatusEnum
const (
	AutonomousDatabaseSummaryOperationsInsightsStatusEnabling        AutonomousDatabaseSummaryOperationsInsightsStatusEnum = "ENABLING"
	AutonomousDatabaseSummaryOperationsInsightsStatusEnabled         AutonomousDatabaseSummaryOperationsInsightsStatusEnum = "ENABLED"
	AutonomousDatabaseSummaryOperationsInsightsStatusDisabling       AutonomousDatabaseSummaryOperationsInsightsStatusEnum = "DISABLING"
	AutonomousDatabaseSummaryOperationsInsightsStatusNotEnabled      AutonomousDatabaseSummaryOperationsInsightsStatusEnum = "NOT_ENABLED"
	AutonomousDatabaseSummaryOperationsInsightsStatusFailedEnabling  AutonomousDatabaseSummaryOperationsInsightsStatusEnum = "FAILED_ENABLING"
	AutonomousDatabaseSummaryOperationsInsightsStatusFailedDisabling AutonomousDatabaseSummaryOperationsInsightsStatusEnum = "FAILED_DISABLING"
)

var mappingAutonomousDatabaseSummaryOperationsInsightsStatusEnum = map[string]AutonomousDatabaseSummaryOperationsInsightsStatusEnum{
	"ENABLING":         AutonomousDatabaseSummaryOperationsInsightsStatusEnabling,
	"ENABLED":          AutonomousDatabaseSummaryOperationsInsightsStatusEnabled,
	"DISABLING":        AutonomousDatabaseSummaryOperationsInsightsStatusDisabling,
	"NOT_ENABLED":      AutonomousDatabaseSummaryOperationsInsightsStatusNotEnabled,
	"FAILED_ENABLING":  AutonomousDatabaseSummaryOperationsInsightsStatusFailedEnabling,
	"FAILED_DISABLING": AutonomousDatabaseSummaryOperationsInsightsStatusFailedDisabling,
}

var mappingAutonomousDatabaseSummaryOperationsInsightsStatusEnumLowerCase = map[string]AutonomousDatabaseSummaryOperationsInsightsStatusEnum{
	"enabling":         AutonomousDatabaseSummaryOperationsInsightsStatusEnabling,
	"enabled":          AutonomousDatabaseSummaryOperationsInsightsStatusEnabled,
	"disabling":        AutonomousDatabaseSummaryOperationsInsightsStatusDisabling,
	"not_enabled":      AutonomousDatabaseSummaryOperationsInsightsStatusNotEnabled,
	"failed_enabling":  AutonomousDatabaseSummaryOperationsInsightsStatusFailedEnabling,
	"failed_disabling": AutonomousDatabaseSummaryOperationsInsightsStatusFailedDisabling,
}

// GetAutonomousDatabaseSummaryOperationsInsightsStatusEnumValues Enumerates the set of values for AutonomousDatabaseSummaryOperationsInsightsStatusEnum
func GetAutonomousDatabaseSummaryOperationsInsightsStatusEnumValues() []AutonomousDatabaseSummaryOperationsInsightsStatusEnum {
	values := make([]AutonomousDatabaseSummaryOperationsInsightsStatusEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryOperationsInsightsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryOperationsInsightsStatusEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryOperationsInsightsStatusEnum
func GetAutonomousDatabaseSummaryOperationsInsightsStatusEnumStringValues() []string {
	return []string{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"NOT_ENABLED",
		"FAILED_ENABLING",
		"FAILED_DISABLING",
	}
}

// GetMappingAutonomousDatabaseSummaryOperationsInsightsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryOperationsInsightsStatusEnum(val string) (AutonomousDatabaseSummaryOperationsInsightsStatusEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryOperationsInsightsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryDatabaseManagementStatusEnum Enum with underlying type: string
type AutonomousDatabaseSummaryDatabaseManagementStatusEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryDatabaseManagementStatusEnum
const (
	AutonomousDatabaseSummaryDatabaseManagementStatusEnabling        AutonomousDatabaseSummaryDatabaseManagementStatusEnum = "ENABLING"
	AutonomousDatabaseSummaryDatabaseManagementStatusEnabled         AutonomousDatabaseSummaryDatabaseManagementStatusEnum = "ENABLED"
	AutonomousDatabaseSummaryDatabaseManagementStatusDisabling       AutonomousDatabaseSummaryDatabaseManagementStatusEnum = "DISABLING"
	AutonomousDatabaseSummaryDatabaseManagementStatusNotEnabled      AutonomousDatabaseSummaryDatabaseManagementStatusEnum = "NOT_ENABLED"
	AutonomousDatabaseSummaryDatabaseManagementStatusFailedEnabling  AutonomousDatabaseSummaryDatabaseManagementStatusEnum = "FAILED_ENABLING"
	AutonomousDatabaseSummaryDatabaseManagementStatusFailedDisabling AutonomousDatabaseSummaryDatabaseManagementStatusEnum = "FAILED_DISABLING"
)

var mappingAutonomousDatabaseSummaryDatabaseManagementStatusEnum = map[string]AutonomousDatabaseSummaryDatabaseManagementStatusEnum{
	"ENABLING":         AutonomousDatabaseSummaryDatabaseManagementStatusEnabling,
	"ENABLED":          AutonomousDatabaseSummaryDatabaseManagementStatusEnabled,
	"DISABLING":        AutonomousDatabaseSummaryDatabaseManagementStatusDisabling,
	"NOT_ENABLED":      AutonomousDatabaseSummaryDatabaseManagementStatusNotEnabled,
	"FAILED_ENABLING":  AutonomousDatabaseSummaryDatabaseManagementStatusFailedEnabling,
	"FAILED_DISABLING": AutonomousDatabaseSummaryDatabaseManagementStatusFailedDisabling,
}

var mappingAutonomousDatabaseSummaryDatabaseManagementStatusEnumLowerCase = map[string]AutonomousDatabaseSummaryDatabaseManagementStatusEnum{
	"enabling":         AutonomousDatabaseSummaryDatabaseManagementStatusEnabling,
	"enabled":          AutonomousDatabaseSummaryDatabaseManagementStatusEnabled,
	"disabling":        AutonomousDatabaseSummaryDatabaseManagementStatusDisabling,
	"not_enabled":      AutonomousDatabaseSummaryDatabaseManagementStatusNotEnabled,
	"failed_enabling":  AutonomousDatabaseSummaryDatabaseManagementStatusFailedEnabling,
	"failed_disabling": AutonomousDatabaseSummaryDatabaseManagementStatusFailedDisabling,
}

// GetAutonomousDatabaseSummaryDatabaseManagementStatusEnumValues Enumerates the set of values for AutonomousDatabaseSummaryDatabaseManagementStatusEnum
func GetAutonomousDatabaseSummaryDatabaseManagementStatusEnumValues() []AutonomousDatabaseSummaryDatabaseManagementStatusEnum {
	values := make([]AutonomousDatabaseSummaryDatabaseManagementStatusEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryDatabaseManagementStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryDatabaseManagementStatusEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryDatabaseManagementStatusEnum
func GetAutonomousDatabaseSummaryDatabaseManagementStatusEnumStringValues() []string {
	return []string{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"NOT_ENABLED",
		"FAILED_ENABLING",
		"FAILED_DISABLING",
	}
}

// GetMappingAutonomousDatabaseSummaryDatabaseManagementStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryDatabaseManagementStatusEnum(val string) (AutonomousDatabaseSummaryDatabaseManagementStatusEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryDatabaseManagementStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryOpenModeEnum Enum with underlying type: string
type AutonomousDatabaseSummaryOpenModeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryOpenModeEnum
const (
	AutonomousDatabaseSummaryOpenModeOnly  AutonomousDatabaseSummaryOpenModeEnum = "READ_ONLY"
	AutonomousDatabaseSummaryOpenModeWrite AutonomousDatabaseSummaryOpenModeEnum = "READ_WRITE"
)

var mappingAutonomousDatabaseSummaryOpenModeEnum = map[string]AutonomousDatabaseSummaryOpenModeEnum{
	"READ_ONLY":  AutonomousDatabaseSummaryOpenModeOnly,
	"READ_WRITE": AutonomousDatabaseSummaryOpenModeWrite,
}

var mappingAutonomousDatabaseSummaryOpenModeEnumLowerCase = map[string]AutonomousDatabaseSummaryOpenModeEnum{
	"read_only":  AutonomousDatabaseSummaryOpenModeOnly,
	"read_write": AutonomousDatabaseSummaryOpenModeWrite,
}

// GetAutonomousDatabaseSummaryOpenModeEnumValues Enumerates the set of values for AutonomousDatabaseSummaryOpenModeEnum
func GetAutonomousDatabaseSummaryOpenModeEnumValues() []AutonomousDatabaseSummaryOpenModeEnum {
	values := make([]AutonomousDatabaseSummaryOpenModeEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryOpenModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryOpenModeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryOpenModeEnum
func GetAutonomousDatabaseSummaryOpenModeEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
	}
}

// GetMappingAutonomousDatabaseSummaryOpenModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryOpenModeEnum(val string) (AutonomousDatabaseSummaryOpenModeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryOpenModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryRefreshableStatusEnum Enum with underlying type: string
type AutonomousDatabaseSummaryRefreshableStatusEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryRefreshableStatusEnum
const (
	AutonomousDatabaseSummaryRefreshableStatusRefreshing    AutonomousDatabaseSummaryRefreshableStatusEnum = "REFRESHING"
	AutonomousDatabaseSummaryRefreshableStatusNotRefreshing AutonomousDatabaseSummaryRefreshableStatusEnum = "NOT_REFRESHING"
)

var mappingAutonomousDatabaseSummaryRefreshableStatusEnum = map[string]AutonomousDatabaseSummaryRefreshableStatusEnum{
	"REFRESHING":     AutonomousDatabaseSummaryRefreshableStatusRefreshing,
	"NOT_REFRESHING": AutonomousDatabaseSummaryRefreshableStatusNotRefreshing,
}

var mappingAutonomousDatabaseSummaryRefreshableStatusEnumLowerCase = map[string]AutonomousDatabaseSummaryRefreshableStatusEnum{
	"refreshing":     AutonomousDatabaseSummaryRefreshableStatusRefreshing,
	"not_refreshing": AutonomousDatabaseSummaryRefreshableStatusNotRefreshing,
}

// GetAutonomousDatabaseSummaryRefreshableStatusEnumValues Enumerates the set of values for AutonomousDatabaseSummaryRefreshableStatusEnum
func GetAutonomousDatabaseSummaryRefreshableStatusEnumValues() []AutonomousDatabaseSummaryRefreshableStatusEnum {
	values := make([]AutonomousDatabaseSummaryRefreshableStatusEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryRefreshableStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryRefreshableStatusEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryRefreshableStatusEnum
func GetAutonomousDatabaseSummaryRefreshableStatusEnumStringValues() []string {
	return []string{
		"REFRESHING",
		"NOT_REFRESHING",
	}
}

// GetMappingAutonomousDatabaseSummaryRefreshableStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryRefreshableStatusEnum(val string) (AutonomousDatabaseSummaryRefreshableStatusEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryRefreshableStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryRefreshableModeEnum Enum with underlying type: string
type AutonomousDatabaseSummaryRefreshableModeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryRefreshableModeEnum
const (
	AutonomousDatabaseSummaryRefreshableModeAutomatic AutonomousDatabaseSummaryRefreshableModeEnum = "AUTOMATIC"
	AutonomousDatabaseSummaryRefreshableModeManual    AutonomousDatabaseSummaryRefreshableModeEnum = "MANUAL"
)

var mappingAutonomousDatabaseSummaryRefreshableModeEnum = map[string]AutonomousDatabaseSummaryRefreshableModeEnum{
	"AUTOMATIC": AutonomousDatabaseSummaryRefreshableModeAutomatic,
	"MANUAL":    AutonomousDatabaseSummaryRefreshableModeManual,
}

var mappingAutonomousDatabaseSummaryRefreshableModeEnumLowerCase = map[string]AutonomousDatabaseSummaryRefreshableModeEnum{
	"automatic": AutonomousDatabaseSummaryRefreshableModeAutomatic,
	"manual":    AutonomousDatabaseSummaryRefreshableModeManual,
}

// GetAutonomousDatabaseSummaryRefreshableModeEnumValues Enumerates the set of values for AutonomousDatabaseSummaryRefreshableModeEnum
func GetAutonomousDatabaseSummaryRefreshableModeEnumValues() []AutonomousDatabaseSummaryRefreshableModeEnum {
	values := make([]AutonomousDatabaseSummaryRefreshableModeEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryRefreshableModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryRefreshableModeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryRefreshableModeEnum
func GetAutonomousDatabaseSummaryRefreshableModeEnumStringValues() []string {
	return []string{
		"AUTOMATIC",
		"MANUAL",
	}
}

// GetMappingAutonomousDatabaseSummaryRefreshableModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryRefreshableModeEnum(val string) (AutonomousDatabaseSummaryRefreshableModeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryRefreshableModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryPermissionLevelEnum Enum with underlying type: string
type AutonomousDatabaseSummaryPermissionLevelEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryPermissionLevelEnum
const (
	AutonomousDatabaseSummaryPermissionLevelRestricted   AutonomousDatabaseSummaryPermissionLevelEnum = "RESTRICTED"
	AutonomousDatabaseSummaryPermissionLevelUnrestricted AutonomousDatabaseSummaryPermissionLevelEnum = "UNRESTRICTED"
)

var mappingAutonomousDatabaseSummaryPermissionLevelEnum = map[string]AutonomousDatabaseSummaryPermissionLevelEnum{
	"RESTRICTED":   AutonomousDatabaseSummaryPermissionLevelRestricted,
	"UNRESTRICTED": AutonomousDatabaseSummaryPermissionLevelUnrestricted,
}

var mappingAutonomousDatabaseSummaryPermissionLevelEnumLowerCase = map[string]AutonomousDatabaseSummaryPermissionLevelEnum{
	"restricted":   AutonomousDatabaseSummaryPermissionLevelRestricted,
	"unrestricted": AutonomousDatabaseSummaryPermissionLevelUnrestricted,
}

// GetAutonomousDatabaseSummaryPermissionLevelEnumValues Enumerates the set of values for AutonomousDatabaseSummaryPermissionLevelEnum
func GetAutonomousDatabaseSummaryPermissionLevelEnumValues() []AutonomousDatabaseSummaryPermissionLevelEnum {
	values := make([]AutonomousDatabaseSummaryPermissionLevelEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryPermissionLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryPermissionLevelEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryPermissionLevelEnum
func GetAutonomousDatabaseSummaryPermissionLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"UNRESTRICTED",
	}
}

// GetMappingAutonomousDatabaseSummaryPermissionLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryPermissionLevelEnum(val string) (AutonomousDatabaseSummaryPermissionLevelEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryPermissionLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryRoleEnum Enum with underlying type: string
type AutonomousDatabaseSummaryRoleEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryRoleEnum
const (
	AutonomousDatabaseSummaryRolePrimary         AutonomousDatabaseSummaryRoleEnum = "PRIMARY"
	AutonomousDatabaseSummaryRoleStandby         AutonomousDatabaseSummaryRoleEnum = "STANDBY"
	AutonomousDatabaseSummaryRoleDisabledStandby AutonomousDatabaseSummaryRoleEnum = "DISABLED_STANDBY"
	AutonomousDatabaseSummaryRoleBackupCopy      AutonomousDatabaseSummaryRoleEnum = "BACKUP_COPY"
	AutonomousDatabaseSummaryRoleSnapshotStandby AutonomousDatabaseSummaryRoleEnum = "SNAPSHOT_STANDBY"
)

var mappingAutonomousDatabaseSummaryRoleEnum = map[string]AutonomousDatabaseSummaryRoleEnum{
	"PRIMARY":          AutonomousDatabaseSummaryRolePrimary,
	"STANDBY":          AutonomousDatabaseSummaryRoleStandby,
	"DISABLED_STANDBY": AutonomousDatabaseSummaryRoleDisabledStandby,
	"BACKUP_COPY":      AutonomousDatabaseSummaryRoleBackupCopy,
	"SNAPSHOT_STANDBY": AutonomousDatabaseSummaryRoleSnapshotStandby,
}

var mappingAutonomousDatabaseSummaryRoleEnumLowerCase = map[string]AutonomousDatabaseSummaryRoleEnum{
	"primary":          AutonomousDatabaseSummaryRolePrimary,
	"standby":          AutonomousDatabaseSummaryRoleStandby,
	"disabled_standby": AutonomousDatabaseSummaryRoleDisabledStandby,
	"backup_copy":      AutonomousDatabaseSummaryRoleBackupCopy,
	"snapshot_standby": AutonomousDatabaseSummaryRoleSnapshotStandby,
}

// GetAutonomousDatabaseSummaryRoleEnumValues Enumerates the set of values for AutonomousDatabaseSummaryRoleEnum
func GetAutonomousDatabaseSummaryRoleEnumValues() []AutonomousDatabaseSummaryRoleEnum {
	values := make([]AutonomousDatabaseSummaryRoleEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryRoleEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryRoleEnum
func GetAutonomousDatabaseSummaryRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
		"BACKUP_COPY",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingAutonomousDatabaseSummaryRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryRoleEnum(val string) (AutonomousDatabaseSummaryRoleEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryDataguardRegionTypeEnum Enum with underlying type: string
type AutonomousDatabaseSummaryDataguardRegionTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryDataguardRegionTypeEnum
const (
	AutonomousDatabaseSummaryDataguardRegionTypePrimaryDgRegion       AutonomousDatabaseSummaryDataguardRegionTypeEnum = "PRIMARY_DG_REGION"
	AutonomousDatabaseSummaryDataguardRegionTypeRemoteStandbyDgRegion AutonomousDatabaseSummaryDataguardRegionTypeEnum = "REMOTE_STANDBY_DG_REGION"
)

var mappingAutonomousDatabaseSummaryDataguardRegionTypeEnum = map[string]AutonomousDatabaseSummaryDataguardRegionTypeEnum{
	"PRIMARY_DG_REGION":        AutonomousDatabaseSummaryDataguardRegionTypePrimaryDgRegion,
	"REMOTE_STANDBY_DG_REGION": AutonomousDatabaseSummaryDataguardRegionTypeRemoteStandbyDgRegion,
}

var mappingAutonomousDatabaseSummaryDataguardRegionTypeEnumLowerCase = map[string]AutonomousDatabaseSummaryDataguardRegionTypeEnum{
	"primary_dg_region":        AutonomousDatabaseSummaryDataguardRegionTypePrimaryDgRegion,
	"remote_standby_dg_region": AutonomousDatabaseSummaryDataguardRegionTypeRemoteStandbyDgRegion,
}

// GetAutonomousDatabaseSummaryDataguardRegionTypeEnumValues Enumerates the set of values for AutonomousDatabaseSummaryDataguardRegionTypeEnum
func GetAutonomousDatabaseSummaryDataguardRegionTypeEnumValues() []AutonomousDatabaseSummaryDataguardRegionTypeEnum {
	values := make([]AutonomousDatabaseSummaryDataguardRegionTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryDataguardRegionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryDataguardRegionTypeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryDataguardRegionTypeEnum
func GetAutonomousDatabaseSummaryDataguardRegionTypeEnumStringValues() []string {
	return []string{
		"PRIMARY_DG_REGION",
		"REMOTE_STANDBY_DG_REGION",
	}
}

// GetMappingAutonomousDatabaseSummaryDataguardRegionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryDataguardRegionTypeEnum(val string) (AutonomousDatabaseSummaryDataguardRegionTypeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryDataguardRegionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum Enum with underlying type: string
type AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum
const (
	AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEarly   AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum = "EARLY"
	AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeRegular AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum = "REGULAR"
)

var mappingAutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum = map[string]AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum{
	"EARLY":   AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEarly,
	"REGULAR": AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeRegular,
}

var mappingAutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnumLowerCase = map[string]AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum{
	"early":   AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEarly,
	"regular": AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeRegular,
}

// GetAutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnumValues Enumerates the set of values for AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum
func GetAutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnumValues() []AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum {
	values := make([]AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum
func GetAutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnumStringValues() []string {
	return []string{
		"EARLY",
		"REGULAR",
	}
}

// GetMappingAutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum(val string) (AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryDatabaseEditionEnum Enum with underlying type: string
type AutonomousDatabaseSummaryDatabaseEditionEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryDatabaseEditionEnum
const (
	AutonomousDatabaseSummaryDatabaseEditionStandardEdition   AutonomousDatabaseSummaryDatabaseEditionEnum = "STANDARD_EDITION"
	AutonomousDatabaseSummaryDatabaseEditionEnterpriseEdition AutonomousDatabaseSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION"
)

var mappingAutonomousDatabaseSummaryDatabaseEditionEnum = map[string]AutonomousDatabaseSummaryDatabaseEditionEnum{
	"STANDARD_EDITION":   AutonomousDatabaseSummaryDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION": AutonomousDatabaseSummaryDatabaseEditionEnterpriseEdition,
}

var mappingAutonomousDatabaseSummaryDatabaseEditionEnumLowerCase = map[string]AutonomousDatabaseSummaryDatabaseEditionEnum{
	"standard_edition":   AutonomousDatabaseSummaryDatabaseEditionStandardEdition,
	"enterprise_edition": AutonomousDatabaseSummaryDatabaseEditionEnterpriseEdition,
}

// GetAutonomousDatabaseSummaryDatabaseEditionEnumValues Enumerates the set of values for AutonomousDatabaseSummaryDatabaseEditionEnum
func GetAutonomousDatabaseSummaryDatabaseEditionEnumValues() []AutonomousDatabaseSummaryDatabaseEditionEnum {
	values := make([]AutonomousDatabaseSummaryDatabaseEditionEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryDatabaseEditionEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryDatabaseEditionEnum
func GetAutonomousDatabaseSummaryDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
	}
}

// GetMappingAutonomousDatabaseSummaryDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryDatabaseEditionEnum(val string) (AutonomousDatabaseSummaryDatabaseEditionEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum Enum with underlying type: string
type AutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum
const (
	AutonomousDatabaseSummaryDisasterRecoveryRegionTypePrimary AutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum = "PRIMARY"
	AutonomousDatabaseSummaryDisasterRecoveryRegionTypeRemote  AutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum = "REMOTE"
)

var mappingAutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum = map[string]AutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum{
	"PRIMARY": AutonomousDatabaseSummaryDisasterRecoveryRegionTypePrimary,
	"REMOTE":  AutonomousDatabaseSummaryDisasterRecoveryRegionTypeRemote,
}

var mappingAutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnumLowerCase = map[string]AutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum{
	"primary": AutonomousDatabaseSummaryDisasterRecoveryRegionTypePrimary,
	"remote":  AutonomousDatabaseSummaryDisasterRecoveryRegionTypeRemote,
}

// GetAutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnumValues Enumerates the set of values for AutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum
func GetAutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnumValues() []AutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum {
	values := make([]AutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum
func GetAutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"REMOTE",
	}
}

// GetMappingAutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum(val string) (AutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryDisasterRecoveryRegionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseSummaryNetServicesArchitectureEnum Enum with underlying type: string
type AutonomousDatabaseSummaryNetServicesArchitectureEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryNetServicesArchitectureEnum
const (
	AutonomousDatabaseSummaryNetServicesArchitectureDedicated AutonomousDatabaseSummaryNetServicesArchitectureEnum = "DEDICATED"
	AutonomousDatabaseSummaryNetServicesArchitectureShared    AutonomousDatabaseSummaryNetServicesArchitectureEnum = "SHARED"
)

var mappingAutonomousDatabaseSummaryNetServicesArchitectureEnum = map[string]AutonomousDatabaseSummaryNetServicesArchitectureEnum{
	"DEDICATED": AutonomousDatabaseSummaryNetServicesArchitectureDedicated,
	"SHARED":    AutonomousDatabaseSummaryNetServicesArchitectureShared,
}

var mappingAutonomousDatabaseSummaryNetServicesArchitectureEnumLowerCase = map[string]AutonomousDatabaseSummaryNetServicesArchitectureEnum{
	"dedicated": AutonomousDatabaseSummaryNetServicesArchitectureDedicated,
	"shared":    AutonomousDatabaseSummaryNetServicesArchitectureShared,
}

// GetAutonomousDatabaseSummaryNetServicesArchitectureEnumValues Enumerates the set of values for AutonomousDatabaseSummaryNetServicesArchitectureEnum
func GetAutonomousDatabaseSummaryNetServicesArchitectureEnumValues() []AutonomousDatabaseSummaryNetServicesArchitectureEnum {
	values := make([]AutonomousDatabaseSummaryNetServicesArchitectureEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryNetServicesArchitectureEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseSummaryNetServicesArchitectureEnumStringValues Enumerates the set of values in String for AutonomousDatabaseSummaryNetServicesArchitectureEnum
func GetAutonomousDatabaseSummaryNetServicesArchitectureEnumStringValues() []string {
	return []string{
		"DEDICATED",
		"SHARED",
	}
}

// GetMappingAutonomousDatabaseSummaryNetServicesArchitectureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseSummaryNetServicesArchitectureEnum(val string) (AutonomousDatabaseSummaryNetServicesArchitectureEnum, bool) {
	enum, ok := mappingAutonomousDatabaseSummaryNetServicesArchitectureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
