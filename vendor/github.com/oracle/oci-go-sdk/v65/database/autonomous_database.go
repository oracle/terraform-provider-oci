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

// AutonomousDatabase An Oracle Autonomous Database.
type AutonomousDatabase struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the Autonomous Database.
	LifecycleState AutonomousDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The database name.
	DbName *string `mandatory:"true" json:"dbName"`

	// The quantity of data in the database, in terabytes.
	DataStorageSizeInTBs *int `mandatory:"true" json:"dataStorageSizeInTBs"`

	// Information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// KMS key lifecycle details.
	KmsKeyLifecycleDetails *string `mandatory:"false" json:"kmsKeyLifecycleDetails"`

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

	// The number of CPU cores to be made available to the database. When the ECPU is selected, the value for cpuCoreCount is 0. For Autonomous Database on Dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `ocpuCount` parameter.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// Parameter that allows users to select an acceptable maximum data loss limit in seconds, up to which Automatic Failover will be triggered when necessary for a Local Autonomous Data Guard
	LocalAdgAutoFailoverMaxDataLossLimit *int `mandatory:"false" json:"localAdgAutoFailoverMaxDataLossLimit"`

	// The compute model of the Autonomous Database. This is required if using the `computeCount` parameter. If using `cpuCoreCount` then it is an error to specify `computeModel` to a non-null value. ECPU compute model is the recommended model and OCPU compute model is legacy.
	ComputeModel AutonomousDatabaseComputeModelEnum `mandatory:"false" json:"computeModel,omitempty"`

	// The compute amount (CPUs) available to the database. Minimum and maximum values depend on the compute model and whether the database is an Autonomous Database Serverless instance or an Autonomous Database on Dedicated Exadata Infrastructure.
	// For an Autonomous Database Serverless instance, the 'ECPU' compute model requires a minimum value of one, for databases in the elastic resource pool and minimum value of two, otherwise. Required when using the `computeModel` parameter. When using `cpuCoreCount` parameter, it is an error to specify computeCount to a non-null value. Providing `computeModel` and `computeCount` is the preferred method for both OCPU and ECPU.
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
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The storage space consumed by Autonomous Database in GBs.
	UsedDataStorageSizeInGBs *int `mandatory:"false" json:"usedDataStorageSizeInGBs"`

	// The infrastructure type this resource belongs to.
	InfrastructureType AutonomousDatabaseInfrastructureTypeEnum `mandatory:"false" json:"infrastructureType,omitempty"`

	// True if the database uses dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html).
	IsDedicated *bool `mandatory:"false" json:"isDedicated"`

	// The Autonomous Container Database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). Used only by Autonomous Database on Dedicated Exadata Infrastructure.
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

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Oracle Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null. It is already set at the
	// Autonomous Exadata Infrastructure level. When provisioning an Autonomous Database Serverless  (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) database, if a value is not specified, the system defaults the value to `BRING_YOUR_OWN_LICENSE`. Bring your own license (BYOL) also allows you to select the DB edition using the optional parameter.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, dataStorageSizeInTBs, adminPassword, isMTLSConnectionRequired, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	LicenseModel AutonomousDatabaseLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

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
	DbWorkload AutonomousDatabaseDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// This project introduces Autonomous Database for Developers (ADB-Dev), a free tier on dedicated infrastructure, and Cloud@Customer for database development purposes. ADB-Dev enables ExaDB customers to experiment with ADB for free and incentivizes enterprises to use ADB for new development projects.Note that ADB-Dev have 4 CPU and 20GB of memory. For ADB-Dev , memory and CPU cannot be scaled
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
	DataSafeStatus AutonomousDatabaseDataSafeStatusEnum `mandatory:"false" json:"dataSafeStatus,omitempty"`

	// Status of Operations Insights for this Autonomous Database.
	OperationsInsightsStatus AutonomousDatabaseOperationsInsightsStatusEnum `mandatory:"false" json:"operationsInsightsStatus,omitempty"`

	// Status of Database Management for this Autonomous Database.
	DatabaseManagementStatus AutonomousDatabaseDatabaseManagementStatusEnum `mandatory:"false" json:"databaseManagementStatus,omitempty"`

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
	OpenMode AutonomousDatabaseOpenModeEnum `mandatory:"false" json:"openMode,omitempty"`

	// The refresh status of the clone. REFRESHING indicates that the clone is currently being refreshed with data from the source Autonomous Database.
	RefreshableStatus AutonomousDatabaseRefreshableStatusEnum `mandatory:"false" json:"refreshableStatus,omitempty"`

	// The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.
	RefreshableMode AutonomousDatabaseRefreshableModeEnum `mandatory:"false" json:"refreshableMode,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that was cloned to create the current Autonomous Database.
	SourceId *string `mandatory:"false" json:"sourceId"`

	// The Autonomous Database permission level. Restricted mode allows access only by admin users.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	PermissionLevel AutonomousDatabasePermissionLevelEnum `mandatory:"false" json:"permissionLevel,omitempty"`

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
	Role AutonomousDatabaseRoleEnum `mandatory:"false" json:"role,omitempty"`

	// List of Oracle Database versions available for a database upgrade. If there are no version upgrades available, this list is empty.
	AvailableUpgradeVersions []string `mandatory:"false" json:"availableUpgradeVersions"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	// The wallet name for Oracle Key Vault.
	KeyStoreWalletName *string `mandatory:"false" json:"keyStoreWalletName"`

	// The list of regions that support the creation of an Autonomous Database clone or an Autonomous Data Guard standby database.
	SupportedRegionsToCloneTo []string `mandatory:"false" json:"supportedRegionsToCloneTo"`

	// Customer Contacts.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	// The date and time that Autonomous Data Guard was enabled for an Autonomous Database where the standby was provisioned in the same region as the primary database.
	TimeLocalDataGuardEnabled *common.SDKTime `mandatory:"false" json:"timeLocalDataGuardEnabled"`

	// The Autonomous Data Guard region type of the Autonomous Database. For Autonomous Database Serverless, Autonomous Data Guard associations have designated primary and standby regions, and these region types do not change when the database changes roles. The standby regions in Autonomous Data Guard associations can be the same region designated as the primary region, or they can be remote regions. Certain database administrative operations may be available only in the primary region of the Autonomous Data Guard association, and cannot be performed when the database using the primary role is operating in a remote Autonomous Data Guard standby region.
	DataguardRegionType AutonomousDatabaseDataguardRegionTypeEnum `mandatory:"false" json:"dataguardRegionType,omitempty"`

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
	AutonomousMaintenanceScheduleType AutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum `mandatory:"false" json:"autonomousMaintenanceScheduleType,omitempty"`

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
	DatabaseEdition AutonomousDatabaseDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

	// The list of database tools details.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, isLocalDataGuardEnabled, or isFreeTier.
	DbToolsDetails []DatabaseTool `mandatory:"false" json:"dbToolsDetails"`

	// Indicates the local disaster recovery (DR) type of the Autonomous Database Serverless instance.
	// Autonomous Data Guard (ADG) DR type provides business critical DR with a faster recovery time objective (RTO) during failover or switchover.
	// Backup-based DR type provides lower cost DR with a slower RTO during failover or switchover.
	LocalDisasterRecoveryType DisasterRecoveryConfigurationDisasterRecoveryTypeEnum `mandatory:"false" json:"localDisasterRecoveryType,omitempty"`

	// The disaster recovery (DR) region type of the Autonomous Database. For Autonomous Database Serverless instances, DR associations have designated primary and standby regions. These region types do not change when the database changes roles. The standby region in DR associations can be the same region as the primary region, or they can be in a remote regions. Some database administration operations may be available only in the primary region of the DR association, and cannot be performed when the database using the primary role is operating in a remote region.
	DisasterRecoveryRegionType AutonomousDatabaseDisasterRecoveryRegionTypeEnum `mandatory:"false" json:"disasterRecoveryRegionType,omitempty"`

	// The date and time the Disaster Recovery role was switched for the standby Autonomous Database.
	TimeDisasterRecoveryRoleChanged *common.SDKTime `mandatory:"false" json:"timeDisasterRecoveryRoleChanged"`

	RemoteDisasterRecoveryConfiguration *DisasterRecoveryConfiguration `mandatory:"false" json:"remoteDisasterRecoveryConfiguration"`

	// Enabling SHARED server architecture enables a database server to allow many client processes to share very few server processes, thereby increasing the number of supported users.
	NetServicesArchitecture AutonomousDatabaseNetServicesArchitectureEnum `mandatory:"false" json:"netServicesArchitecture,omitempty"`
}

func (m AutonomousDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDatabaseLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousDatabaseComputeModelEnum(string(m.ComputeModel)); !ok && m.ComputeModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeModel: %s. Supported values are: %s.", m.ComputeModel, strings.Join(GetAutonomousDatabaseComputeModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseInfrastructureTypeEnum(string(m.InfrastructureType)); !ok && m.InfrastructureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InfrastructureType: %s. Supported values are: %s.", m.InfrastructureType, strings.Join(GetAutonomousDatabaseInfrastructureTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetAutonomousDatabaseLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetAutonomousDatabaseDbWorkloadEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseDataSafeStatusEnum(string(m.DataSafeStatus)); !ok && m.DataSafeStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataSafeStatus: %s. Supported values are: %s.", m.DataSafeStatus, strings.Join(GetAutonomousDatabaseDataSafeStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseOperationsInsightsStatusEnum(string(m.OperationsInsightsStatus)); !ok && m.OperationsInsightsStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationsInsightsStatus: %s. Supported values are: %s.", m.OperationsInsightsStatus, strings.Join(GetAutonomousDatabaseOperationsInsightsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseDatabaseManagementStatusEnum(string(m.DatabaseManagementStatus)); !ok && m.DatabaseManagementStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseManagementStatus: %s. Supported values are: %s.", m.DatabaseManagementStatus, strings.Join(GetAutonomousDatabaseDatabaseManagementStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseOpenModeEnum(string(m.OpenMode)); !ok && m.OpenMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OpenMode: %s. Supported values are: %s.", m.OpenMode, strings.Join(GetAutonomousDatabaseOpenModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseRefreshableStatusEnum(string(m.RefreshableStatus)); !ok && m.RefreshableStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RefreshableStatus: %s. Supported values are: %s.", m.RefreshableStatus, strings.Join(GetAutonomousDatabaseRefreshableStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseRefreshableModeEnum(string(m.RefreshableMode)); !ok && m.RefreshableMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RefreshableMode: %s. Supported values are: %s.", m.RefreshableMode, strings.Join(GetAutonomousDatabaseRefreshableModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabasePermissionLevelEnum(string(m.PermissionLevel)); !ok && m.PermissionLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PermissionLevel: %s. Supported values are: %s.", m.PermissionLevel, strings.Join(GetAutonomousDatabasePermissionLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetAutonomousDatabaseRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseDataguardRegionTypeEnum(string(m.DataguardRegionType)); !ok && m.DataguardRegionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataguardRegionType: %s. Supported values are: %s.", m.DataguardRegionType, strings.Join(GetAutonomousDatabaseDataguardRegionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum(string(m.AutonomousMaintenanceScheduleType)); !ok && m.AutonomousMaintenanceScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutonomousMaintenanceScheduleType: %s. Supported values are: %s.", m.AutonomousMaintenanceScheduleType, strings.Join(GetAutonomousDatabaseAutonomousMaintenanceScheduleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetAutonomousDatabaseDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDisasterRecoveryConfigurationDisasterRecoveryTypeEnum(string(m.LocalDisasterRecoveryType)); !ok && m.LocalDisasterRecoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LocalDisasterRecoveryType: %s. Supported values are: %s.", m.LocalDisasterRecoveryType, strings.Join(GetDisasterRecoveryConfigurationDisasterRecoveryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseDisasterRecoveryRegionTypeEnum(string(m.DisasterRecoveryRegionType)); !ok && m.DisasterRecoveryRegionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DisasterRecoveryRegionType: %s. Supported values are: %s.", m.DisasterRecoveryRegionType, strings.Join(GetAutonomousDatabaseDisasterRecoveryRegionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseNetServicesArchitectureEnum(string(m.NetServicesArchitecture)); !ok && m.NetServicesArchitecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetServicesArchitecture: %s. Supported values are: %s.", m.NetServicesArchitecture, strings.Join(GetAutonomousDatabaseNetServicesArchitectureEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDatabaseLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseLifecycleStateEnum
const (
	AutonomousDatabaseLifecycleStateProvisioning            AutonomousDatabaseLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseLifecycleStateAvailable               AutonomousDatabaseLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseLifecycleStateStopping                AutonomousDatabaseLifecycleStateEnum = "STOPPING"
	AutonomousDatabaseLifecycleStateStopped                 AutonomousDatabaseLifecycleStateEnum = "STOPPED"
	AutonomousDatabaseLifecycleStateStarting                AutonomousDatabaseLifecycleStateEnum = "STARTING"
	AutonomousDatabaseLifecycleStateTerminating             AutonomousDatabaseLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseLifecycleStateTerminated              AutonomousDatabaseLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseLifecycleStateUnavailable             AutonomousDatabaseLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDatabaseLifecycleStateRestoreInProgress       AutonomousDatabaseLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDatabaseLifecycleStateRestoreFailed           AutonomousDatabaseLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousDatabaseLifecycleStateBackupInProgress        AutonomousDatabaseLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDatabaseLifecycleStateScaleInProgress         AutonomousDatabaseLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDatabaseLifecycleStateAvailableNeedsAttention AutonomousDatabaseLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDatabaseLifecycleStateUpdating                AutonomousDatabaseLifecycleStateEnum = "UPDATING"
	AutonomousDatabaseLifecycleStateMaintenanceInProgress   AutonomousDatabaseLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	AutonomousDatabaseLifecycleStateRestarting              AutonomousDatabaseLifecycleStateEnum = "RESTARTING"
	AutonomousDatabaseLifecycleStateRecreating              AutonomousDatabaseLifecycleStateEnum = "RECREATING"
	AutonomousDatabaseLifecycleStateRoleChangeInProgress    AutonomousDatabaseLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousDatabaseLifecycleStateUpgrading               AutonomousDatabaseLifecycleStateEnum = "UPGRADING"
	AutonomousDatabaseLifecycleStateInaccessible            AutonomousDatabaseLifecycleStateEnum = "INACCESSIBLE"
	AutonomousDatabaseLifecycleStateStandby                 AutonomousDatabaseLifecycleStateEnum = "STANDBY"
)

var mappingAutonomousDatabaseLifecycleStateEnum = map[string]AutonomousDatabaseLifecycleStateEnum{
	"PROVISIONING":              AutonomousDatabaseLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDatabaseLifecycleStateAvailable,
	"STOPPING":                  AutonomousDatabaseLifecycleStateStopping,
	"STOPPED":                   AutonomousDatabaseLifecycleStateStopped,
	"STARTING":                  AutonomousDatabaseLifecycleStateStarting,
	"TERMINATING":               AutonomousDatabaseLifecycleStateTerminating,
	"TERMINATED":                AutonomousDatabaseLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDatabaseLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDatabaseLifecycleStateRestoreInProgress,
	"RESTORE_FAILED":            AutonomousDatabaseLifecycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":        AutonomousDatabaseLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDatabaseLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDatabaseLifecycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDatabaseLifecycleStateUpdating,
	"MAINTENANCE_IN_PROGRESS":   AutonomousDatabaseLifecycleStateMaintenanceInProgress,
	"RESTARTING":                AutonomousDatabaseLifecycleStateRestarting,
	"RECREATING":                AutonomousDatabaseLifecycleStateRecreating,
	"ROLE_CHANGE_IN_PROGRESS":   AutonomousDatabaseLifecycleStateRoleChangeInProgress,
	"UPGRADING":                 AutonomousDatabaseLifecycleStateUpgrading,
	"INACCESSIBLE":              AutonomousDatabaseLifecycleStateInaccessible,
	"STANDBY":                   AutonomousDatabaseLifecycleStateStandby,
}

var mappingAutonomousDatabaseLifecycleStateEnumLowerCase = map[string]AutonomousDatabaseLifecycleStateEnum{
	"provisioning":              AutonomousDatabaseLifecycleStateProvisioning,
	"available":                 AutonomousDatabaseLifecycleStateAvailable,
	"stopping":                  AutonomousDatabaseLifecycleStateStopping,
	"stopped":                   AutonomousDatabaseLifecycleStateStopped,
	"starting":                  AutonomousDatabaseLifecycleStateStarting,
	"terminating":               AutonomousDatabaseLifecycleStateTerminating,
	"terminated":                AutonomousDatabaseLifecycleStateTerminated,
	"unavailable":               AutonomousDatabaseLifecycleStateUnavailable,
	"restore_in_progress":       AutonomousDatabaseLifecycleStateRestoreInProgress,
	"restore_failed":            AutonomousDatabaseLifecycleStateRestoreFailed,
	"backup_in_progress":        AutonomousDatabaseLifecycleStateBackupInProgress,
	"scale_in_progress":         AutonomousDatabaseLifecycleStateScaleInProgress,
	"available_needs_attention": AutonomousDatabaseLifecycleStateAvailableNeedsAttention,
	"updating":                  AutonomousDatabaseLifecycleStateUpdating,
	"maintenance_in_progress":   AutonomousDatabaseLifecycleStateMaintenanceInProgress,
	"restarting":                AutonomousDatabaseLifecycleStateRestarting,
	"recreating":                AutonomousDatabaseLifecycleStateRecreating,
	"role_change_in_progress":   AutonomousDatabaseLifecycleStateRoleChangeInProgress,
	"upgrading":                 AutonomousDatabaseLifecycleStateUpgrading,
	"inaccessible":              AutonomousDatabaseLifecycleStateInaccessible,
	"standby":                   AutonomousDatabaseLifecycleStateStandby,
}

// GetAutonomousDatabaseLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseLifecycleStateEnum
func GetAutonomousDatabaseLifecycleStateEnumValues() []AutonomousDatabaseLifecycleStateEnum {
	values := make([]AutonomousDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDatabaseLifecycleStateEnum
func GetAutonomousDatabaseLifecycleStateEnumStringValues() []string {
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

// GetMappingAutonomousDatabaseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseLifecycleStateEnum(val string) (AutonomousDatabaseLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDatabaseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseComputeModelEnum Enum with underlying type: string
type AutonomousDatabaseComputeModelEnum string

// Set of constants representing the allowable values for AutonomousDatabaseComputeModelEnum
const (
	AutonomousDatabaseComputeModelEcpu AutonomousDatabaseComputeModelEnum = "ECPU"
	AutonomousDatabaseComputeModelOcpu AutonomousDatabaseComputeModelEnum = "OCPU"
)

var mappingAutonomousDatabaseComputeModelEnum = map[string]AutonomousDatabaseComputeModelEnum{
	"ECPU": AutonomousDatabaseComputeModelEcpu,
	"OCPU": AutonomousDatabaseComputeModelOcpu,
}

var mappingAutonomousDatabaseComputeModelEnumLowerCase = map[string]AutonomousDatabaseComputeModelEnum{
	"ecpu": AutonomousDatabaseComputeModelEcpu,
	"ocpu": AutonomousDatabaseComputeModelOcpu,
}

// GetAutonomousDatabaseComputeModelEnumValues Enumerates the set of values for AutonomousDatabaseComputeModelEnum
func GetAutonomousDatabaseComputeModelEnumValues() []AutonomousDatabaseComputeModelEnum {
	values := make([]AutonomousDatabaseComputeModelEnum, 0)
	for _, v := range mappingAutonomousDatabaseComputeModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseComputeModelEnumStringValues Enumerates the set of values in String for AutonomousDatabaseComputeModelEnum
func GetAutonomousDatabaseComputeModelEnumStringValues() []string {
	return []string{
		"ECPU",
		"OCPU",
	}
}

// GetMappingAutonomousDatabaseComputeModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseComputeModelEnum(val string) (AutonomousDatabaseComputeModelEnum, bool) {
	enum, ok := mappingAutonomousDatabaseComputeModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseInfrastructureTypeEnum Enum with underlying type: string
type AutonomousDatabaseInfrastructureTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseInfrastructureTypeEnum
const (
	AutonomousDatabaseInfrastructureTypeCloud           AutonomousDatabaseInfrastructureTypeEnum = "CLOUD"
	AutonomousDatabaseInfrastructureTypeCloudAtCustomer AutonomousDatabaseInfrastructureTypeEnum = "CLOUD_AT_CUSTOMER"
)

var mappingAutonomousDatabaseInfrastructureTypeEnum = map[string]AutonomousDatabaseInfrastructureTypeEnum{
	"CLOUD":             AutonomousDatabaseInfrastructureTypeCloud,
	"CLOUD_AT_CUSTOMER": AutonomousDatabaseInfrastructureTypeCloudAtCustomer,
}

var mappingAutonomousDatabaseInfrastructureTypeEnumLowerCase = map[string]AutonomousDatabaseInfrastructureTypeEnum{
	"cloud":             AutonomousDatabaseInfrastructureTypeCloud,
	"cloud_at_customer": AutonomousDatabaseInfrastructureTypeCloudAtCustomer,
}

// GetAutonomousDatabaseInfrastructureTypeEnumValues Enumerates the set of values for AutonomousDatabaseInfrastructureTypeEnum
func GetAutonomousDatabaseInfrastructureTypeEnumValues() []AutonomousDatabaseInfrastructureTypeEnum {
	values := make([]AutonomousDatabaseInfrastructureTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseInfrastructureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseInfrastructureTypeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseInfrastructureTypeEnum
func GetAutonomousDatabaseInfrastructureTypeEnumStringValues() []string {
	return []string{
		"CLOUD",
		"CLOUD_AT_CUSTOMER",
	}
}

// GetMappingAutonomousDatabaseInfrastructureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseInfrastructureTypeEnum(val string) (AutonomousDatabaseInfrastructureTypeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseInfrastructureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseLicenseModelEnum Enum with underlying type: string
type AutonomousDatabaseLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousDatabaseLicenseModelEnum
const (
	AutonomousDatabaseLicenseModelLicenseIncluded     AutonomousDatabaseLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousDatabaseLicenseModelBringYourOwnLicense AutonomousDatabaseLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousDatabaseLicenseModelEnum = map[string]AutonomousDatabaseLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousDatabaseLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousDatabaseLicenseModelBringYourOwnLicense,
}

var mappingAutonomousDatabaseLicenseModelEnumLowerCase = map[string]AutonomousDatabaseLicenseModelEnum{
	"license_included":       AutonomousDatabaseLicenseModelLicenseIncluded,
	"bring_your_own_license": AutonomousDatabaseLicenseModelBringYourOwnLicense,
}

// GetAutonomousDatabaseLicenseModelEnumValues Enumerates the set of values for AutonomousDatabaseLicenseModelEnum
func GetAutonomousDatabaseLicenseModelEnumValues() []AutonomousDatabaseLicenseModelEnum {
	values := make([]AutonomousDatabaseLicenseModelEnum, 0)
	for _, v := range mappingAutonomousDatabaseLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseLicenseModelEnumStringValues Enumerates the set of values in String for AutonomousDatabaseLicenseModelEnum
func GetAutonomousDatabaseLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingAutonomousDatabaseLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseLicenseModelEnum(val string) (AutonomousDatabaseLicenseModelEnum, bool) {
	enum, ok := mappingAutonomousDatabaseLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseDbWorkloadEnum Enum with underlying type: string
type AutonomousDatabaseDbWorkloadEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDbWorkloadEnum
const (
	AutonomousDatabaseDbWorkloadOltp AutonomousDatabaseDbWorkloadEnum = "OLTP"
	AutonomousDatabaseDbWorkloadDw   AutonomousDatabaseDbWorkloadEnum = "DW"
	AutonomousDatabaseDbWorkloadAjd  AutonomousDatabaseDbWorkloadEnum = "AJD"
	AutonomousDatabaseDbWorkloadApex AutonomousDatabaseDbWorkloadEnum = "APEX"
)

var mappingAutonomousDatabaseDbWorkloadEnum = map[string]AutonomousDatabaseDbWorkloadEnum{
	"OLTP": AutonomousDatabaseDbWorkloadOltp,
	"DW":   AutonomousDatabaseDbWorkloadDw,
	"AJD":  AutonomousDatabaseDbWorkloadAjd,
	"APEX": AutonomousDatabaseDbWorkloadApex,
}

var mappingAutonomousDatabaseDbWorkloadEnumLowerCase = map[string]AutonomousDatabaseDbWorkloadEnum{
	"oltp": AutonomousDatabaseDbWorkloadOltp,
	"dw":   AutonomousDatabaseDbWorkloadDw,
	"ajd":  AutonomousDatabaseDbWorkloadAjd,
	"apex": AutonomousDatabaseDbWorkloadApex,
}

// GetAutonomousDatabaseDbWorkloadEnumValues Enumerates the set of values for AutonomousDatabaseDbWorkloadEnum
func GetAutonomousDatabaseDbWorkloadEnumValues() []AutonomousDatabaseDbWorkloadEnum {
	values := make([]AutonomousDatabaseDbWorkloadEnum, 0)
	for _, v := range mappingAutonomousDatabaseDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseDbWorkloadEnumStringValues Enumerates the set of values in String for AutonomousDatabaseDbWorkloadEnum
func GetAutonomousDatabaseDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
		"AJD",
		"APEX",
	}
}

// GetMappingAutonomousDatabaseDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseDbWorkloadEnum(val string) (AutonomousDatabaseDbWorkloadEnum, bool) {
	enum, ok := mappingAutonomousDatabaseDbWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseDataSafeStatusEnum Enum with underlying type: string
type AutonomousDatabaseDataSafeStatusEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDataSafeStatusEnum
const (
	AutonomousDatabaseDataSafeStatusRegistering   AutonomousDatabaseDataSafeStatusEnum = "REGISTERING"
	AutonomousDatabaseDataSafeStatusRegistered    AutonomousDatabaseDataSafeStatusEnum = "REGISTERED"
	AutonomousDatabaseDataSafeStatusDeregistering AutonomousDatabaseDataSafeStatusEnum = "DEREGISTERING"
	AutonomousDatabaseDataSafeStatusNotRegistered AutonomousDatabaseDataSafeStatusEnum = "NOT_REGISTERED"
	AutonomousDatabaseDataSafeStatusFailed        AutonomousDatabaseDataSafeStatusEnum = "FAILED"
)

var mappingAutonomousDatabaseDataSafeStatusEnum = map[string]AutonomousDatabaseDataSafeStatusEnum{
	"REGISTERING":    AutonomousDatabaseDataSafeStatusRegistering,
	"REGISTERED":     AutonomousDatabaseDataSafeStatusRegistered,
	"DEREGISTERING":  AutonomousDatabaseDataSafeStatusDeregistering,
	"NOT_REGISTERED": AutonomousDatabaseDataSafeStatusNotRegistered,
	"FAILED":         AutonomousDatabaseDataSafeStatusFailed,
}

var mappingAutonomousDatabaseDataSafeStatusEnumLowerCase = map[string]AutonomousDatabaseDataSafeStatusEnum{
	"registering":    AutonomousDatabaseDataSafeStatusRegistering,
	"registered":     AutonomousDatabaseDataSafeStatusRegistered,
	"deregistering":  AutonomousDatabaseDataSafeStatusDeregistering,
	"not_registered": AutonomousDatabaseDataSafeStatusNotRegistered,
	"failed":         AutonomousDatabaseDataSafeStatusFailed,
}

// GetAutonomousDatabaseDataSafeStatusEnumValues Enumerates the set of values for AutonomousDatabaseDataSafeStatusEnum
func GetAutonomousDatabaseDataSafeStatusEnumValues() []AutonomousDatabaseDataSafeStatusEnum {
	values := make([]AutonomousDatabaseDataSafeStatusEnum, 0)
	for _, v := range mappingAutonomousDatabaseDataSafeStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseDataSafeStatusEnumStringValues Enumerates the set of values in String for AutonomousDatabaseDataSafeStatusEnum
func GetAutonomousDatabaseDataSafeStatusEnumStringValues() []string {
	return []string{
		"REGISTERING",
		"REGISTERED",
		"DEREGISTERING",
		"NOT_REGISTERED",
		"FAILED",
	}
}

// GetMappingAutonomousDatabaseDataSafeStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseDataSafeStatusEnum(val string) (AutonomousDatabaseDataSafeStatusEnum, bool) {
	enum, ok := mappingAutonomousDatabaseDataSafeStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseOperationsInsightsStatusEnum Enum with underlying type: string
type AutonomousDatabaseOperationsInsightsStatusEnum string

// Set of constants representing the allowable values for AutonomousDatabaseOperationsInsightsStatusEnum
const (
	AutonomousDatabaseOperationsInsightsStatusEnabling        AutonomousDatabaseOperationsInsightsStatusEnum = "ENABLING"
	AutonomousDatabaseOperationsInsightsStatusEnabled         AutonomousDatabaseOperationsInsightsStatusEnum = "ENABLED"
	AutonomousDatabaseOperationsInsightsStatusDisabling       AutonomousDatabaseOperationsInsightsStatusEnum = "DISABLING"
	AutonomousDatabaseOperationsInsightsStatusNotEnabled      AutonomousDatabaseOperationsInsightsStatusEnum = "NOT_ENABLED"
	AutonomousDatabaseOperationsInsightsStatusFailedEnabling  AutonomousDatabaseOperationsInsightsStatusEnum = "FAILED_ENABLING"
	AutonomousDatabaseOperationsInsightsStatusFailedDisabling AutonomousDatabaseOperationsInsightsStatusEnum = "FAILED_DISABLING"
)

var mappingAutonomousDatabaseOperationsInsightsStatusEnum = map[string]AutonomousDatabaseOperationsInsightsStatusEnum{
	"ENABLING":         AutonomousDatabaseOperationsInsightsStatusEnabling,
	"ENABLED":          AutonomousDatabaseOperationsInsightsStatusEnabled,
	"DISABLING":        AutonomousDatabaseOperationsInsightsStatusDisabling,
	"NOT_ENABLED":      AutonomousDatabaseOperationsInsightsStatusNotEnabled,
	"FAILED_ENABLING":  AutonomousDatabaseOperationsInsightsStatusFailedEnabling,
	"FAILED_DISABLING": AutonomousDatabaseOperationsInsightsStatusFailedDisabling,
}

var mappingAutonomousDatabaseOperationsInsightsStatusEnumLowerCase = map[string]AutonomousDatabaseOperationsInsightsStatusEnum{
	"enabling":         AutonomousDatabaseOperationsInsightsStatusEnabling,
	"enabled":          AutonomousDatabaseOperationsInsightsStatusEnabled,
	"disabling":        AutonomousDatabaseOperationsInsightsStatusDisabling,
	"not_enabled":      AutonomousDatabaseOperationsInsightsStatusNotEnabled,
	"failed_enabling":  AutonomousDatabaseOperationsInsightsStatusFailedEnabling,
	"failed_disabling": AutonomousDatabaseOperationsInsightsStatusFailedDisabling,
}

// GetAutonomousDatabaseOperationsInsightsStatusEnumValues Enumerates the set of values for AutonomousDatabaseOperationsInsightsStatusEnum
func GetAutonomousDatabaseOperationsInsightsStatusEnumValues() []AutonomousDatabaseOperationsInsightsStatusEnum {
	values := make([]AutonomousDatabaseOperationsInsightsStatusEnum, 0)
	for _, v := range mappingAutonomousDatabaseOperationsInsightsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseOperationsInsightsStatusEnumStringValues Enumerates the set of values in String for AutonomousDatabaseOperationsInsightsStatusEnum
func GetAutonomousDatabaseOperationsInsightsStatusEnumStringValues() []string {
	return []string{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"NOT_ENABLED",
		"FAILED_ENABLING",
		"FAILED_DISABLING",
	}
}

// GetMappingAutonomousDatabaseOperationsInsightsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseOperationsInsightsStatusEnum(val string) (AutonomousDatabaseOperationsInsightsStatusEnum, bool) {
	enum, ok := mappingAutonomousDatabaseOperationsInsightsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseDatabaseManagementStatusEnum Enum with underlying type: string
type AutonomousDatabaseDatabaseManagementStatusEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDatabaseManagementStatusEnum
const (
	AutonomousDatabaseDatabaseManagementStatusEnabling        AutonomousDatabaseDatabaseManagementStatusEnum = "ENABLING"
	AutonomousDatabaseDatabaseManagementStatusEnabled         AutonomousDatabaseDatabaseManagementStatusEnum = "ENABLED"
	AutonomousDatabaseDatabaseManagementStatusDisabling       AutonomousDatabaseDatabaseManagementStatusEnum = "DISABLING"
	AutonomousDatabaseDatabaseManagementStatusNotEnabled      AutonomousDatabaseDatabaseManagementStatusEnum = "NOT_ENABLED"
	AutonomousDatabaseDatabaseManagementStatusFailedEnabling  AutonomousDatabaseDatabaseManagementStatusEnum = "FAILED_ENABLING"
	AutonomousDatabaseDatabaseManagementStatusFailedDisabling AutonomousDatabaseDatabaseManagementStatusEnum = "FAILED_DISABLING"
)

var mappingAutonomousDatabaseDatabaseManagementStatusEnum = map[string]AutonomousDatabaseDatabaseManagementStatusEnum{
	"ENABLING":         AutonomousDatabaseDatabaseManagementStatusEnabling,
	"ENABLED":          AutonomousDatabaseDatabaseManagementStatusEnabled,
	"DISABLING":        AutonomousDatabaseDatabaseManagementStatusDisabling,
	"NOT_ENABLED":      AutonomousDatabaseDatabaseManagementStatusNotEnabled,
	"FAILED_ENABLING":  AutonomousDatabaseDatabaseManagementStatusFailedEnabling,
	"FAILED_DISABLING": AutonomousDatabaseDatabaseManagementStatusFailedDisabling,
}

var mappingAutonomousDatabaseDatabaseManagementStatusEnumLowerCase = map[string]AutonomousDatabaseDatabaseManagementStatusEnum{
	"enabling":         AutonomousDatabaseDatabaseManagementStatusEnabling,
	"enabled":          AutonomousDatabaseDatabaseManagementStatusEnabled,
	"disabling":        AutonomousDatabaseDatabaseManagementStatusDisabling,
	"not_enabled":      AutonomousDatabaseDatabaseManagementStatusNotEnabled,
	"failed_enabling":  AutonomousDatabaseDatabaseManagementStatusFailedEnabling,
	"failed_disabling": AutonomousDatabaseDatabaseManagementStatusFailedDisabling,
}

// GetAutonomousDatabaseDatabaseManagementStatusEnumValues Enumerates the set of values for AutonomousDatabaseDatabaseManagementStatusEnum
func GetAutonomousDatabaseDatabaseManagementStatusEnumValues() []AutonomousDatabaseDatabaseManagementStatusEnum {
	values := make([]AutonomousDatabaseDatabaseManagementStatusEnum, 0)
	for _, v := range mappingAutonomousDatabaseDatabaseManagementStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseDatabaseManagementStatusEnumStringValues Enumerates the set of values in String for AutonomousDatabaseDatabaseManagementStatusEnum
func GetAutonomousDatabaseDatabaseManagementStatusEnumStringValues() []string {
	return []string{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"NOT_ENABLED",
		"FAILED_ENABLING",
		"FAILED_DISABLING",
	}
}

// GetMappingAutonomousDatabaseDatabaseManagementStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseDatabaseManagementStatusEnum(val string) (AutonomousDatabaseDatabaseManagementStatusEnum, bool) {
	enum, ok := mappingAutonomousDatabaseDatabaseManagementStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseOpenModeEnum Enum with underlying type: string
type AutonomousDatabaseOpenModeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseOpenModeEnum
const (
	AutonomousDatabaseOpenModeOnly  AutonomousDatabaseOpenModeEnum = "READ_ONLY"
	AutonomousDatabaseOpenModeWrite AutonomousDatabaseOpenModeEnum = "READ_WRITE"
)

var mappingAutonomousDatabaseOpenModeEnum = map[string]AutonomousDatabaseOpenModeEnum{
	"READ_ONLY":  AutonomousDatabaseOpenModeOnly,
	"READ_WRITE": AutonomousDatabaseOpenModeWrite,
}

var mappingAutonomousDatabaseOpenModeEnumLowerCase = map[string]AutonomousDatabaseOpenModeEnum{
	"read_only":  AutonomousDatabaseOpenModeOnly,
	"read_write": AutonomousDatabaseOpenModeWrite,
}

// GetAutonomousDatabaseOpenModeEnumValues Enumerates the set of values for AutonomousDatabaseOpenModeEnum
func GetAutonomousDatabaseOpenModeEnumValues() []AutonomousDatabaseOpenModeEnum {
	values := make([]AutonomousDatabaseOpenModeEnum, 0)
	for _, v := range mappingAutonomousDatabaseOpenModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseOpenModeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseOpenModeEnum
func GetAutonomousDatabaseOpenModeEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
	}
}

// GetMappingAutonomousDatabaseOpenModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseOpenModeEnum(val string) (AutonomousDatabaseOpenModeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseOpenModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseRefreshableStatusEnum Enum with underlying type: string
type AutonomousDatabaseRefreshableStatusEnum string

// Set of constants representing the allowable values for AutonomousDatabaseRefreshableStatusEnum
const (
	AutonomousDatabaseRefreshableStatusRefreshing    AutonomousDatabaseRefreshableStatusEnum = "REFRESHING"
	AutonomousDatabaseRefreshableStatusNotRefreshing AutonomousDatabaseRefreshableStatusEnum = "NOT_REFRESHING"
)

var mappingAutonomousDatabaseRefreshableStatusEnum = map[string]AutonomousDatabaseRefreshableStatusEnum{
	"REFRESHING":     AutonomousDatabaseRefreshableStatusRefreshing,
	"NOT_REFRESHING": AutonomousDatabaseRefreshableStatusNotRefreshing,
}

var mappingAutonomousDatabaseRefreshableStatusEnumLowerCase = map[string]AutonomousDatabaseRefreshableStatusEnum{
	"refreshing":     AutonomousDatabaseRefreshableStatusRefreshing,
	"not_refreshing": AutonomousDatabaseRefreshableStatusNotRefreshing,
}

// GetAutonomousDatabaseRefreshableStatusEnumValues Enumerates the set of values for AutonomousDatabaseRefreshableStatusEnum
func GetAutonomousDatabaseRefreshableStatusEnumValues() []AutonomousDatabaseRefreshableStatusEnum {
	values := make([]AutonomousDatabaseRefreshableStatusEnum, 0)
	for _, v := range mappingAutonomousDatabaseRefreshableStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseRefreshableStatusEnumStringValues Enumerates the set of values in String for AutonomousDatabaseRefreshableStatusEnum
func GetAutonomousDatabaseRefreshableStatusEnumStringValues() []string {
	return []string{
		"REFRESHING",
		"NOT_REFRESHING",
	}
}

// GetMappingAutonomousDatabaseRefreshableStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseRefreshableStatusEnum(val string) (AutonomousDatabaseRefreshableStatusEnum, bool) {
	enum, ok := mappingAutonomousDatabaseRefreshableStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseRefreshableModeEnum Enum with underlying type: string
type AutonomousDatabaseRefreshableModeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseRefreshableModeEnum
const (
	AutonomousDatabaseRefreshableModeAutomatic AutonomousDatabaseRefreshableModeEnum = "AUTOMATIC"
	AutonomousDatabaseRefreshableModeManual    AutonomousDatabaseRefreshableModeEnum = "MANUAL"
)

var mappingAutonomousDatabaseRefreshableModeEnum = map[string]AutonomousDatabaseRefreshableModeEnum{
	"AUTOMATIC": AutonomousDatabaseRefreshableModeAutomatic,
	"MANUAL":    AutonomousDatabaseRefreshableModeManual,
}

var mappingAutonomousDatabaseRefreshableModeEnumLowerCase = map[string]AutonomousDatabaseRefreshableModeEnum{
	"automatic": AutonomousDatabaseRefreshableModeAutomatic,
	"manual":    AutonomousDatabaseRefreshableModeManual,
}

// GetAutonomousDatabaseRefreshableModeEnumValues Enumerates the set of values for AutonomousDatabaseRefreshableModeEnum
func GetAutonomousDatabaseRefreshableModeEnumValues() []AutonomousDatabaseRefreshableModeEnum {
	values := make([]AutonomousDatabaseRefreshableModeEnum, 0)
	for _, v := range mappingAutonomousDatabaseRefreshableModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseRefreshableModeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseRefreshableModeEnum
func GetAutonomousDatabaseRefreshableModeEnumStringValues() []string {
	return []string{
		"AUTOMATIC",
		"MANUAL",
	}
}

// GetMappingAutonomousDatabaseRefreshableModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseRefreshableModeEnum(val string) (AutonomousDatabaseRefreshableModeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseRefreshableModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabasePermissionLevelEnum Enum with underlying type: string
type AutonomousDatabasePermissionLevelEnum string

// Set of constants representing the allowable values for AutonomousDatabasePermissionLevelEnum
const (
	AutonomousDatabasePermissionLevelRestricted   AutonomousDatabasePermissionLevelEnum = "RESTRICTED"
	AutonomousDatabasePermissionLevelUnrestricted AutonomousDatabasePermissionLevelEnum = "UNRESTRICTED"
)

var mappingAutonomousDatabasePermissionLevelEnum = map[string]AutonomousDatabasePermissionLevelEnum{
	"RESTRICTED":   AutonomousDatabasePermissionLevelRestricted,
	"UNRESTRICTED": AutonomousDatabasePermissionLevelUnrestricted,
}

var mappingAutonomousDatabasePermissionLevelEnumLowerCase = map[string]AutonomousDatabasePermissionLevelEnum{
	"restricted":   AutonomousDatabasePermissionLevelRestricted,
	"unrestricted": AutonomousDatabasePermissionLevelUnrestricted,
}

// GetAutonomousDatabasePermissionLevelEnumValues Enumerates the set of values for AutonomousDatabasePermissionLevelEnum
func GetAutonomousDatabasePermissionLevelEnumValues() []AutonomousDatabasePermissionLevelEnum {
	values := make([]AutonomousDatabasePermissionLevelEnum, 0)
	for _, v := range mappingAutonomousDatabasePermissionLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabasePermissionLevelEnumStringValues Enumerates the set of values in String for AutonomousDatabasePermissionLevelEnum
func GetAutonomousDatabasePermissionLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"UNRESTRICTED",
	}
}

// GetMappingAutonomousDatabasePermissionLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabasePermissionLevelEnum(val string) (AutonomousDatabasePermissionLevelEnum, bool) {
	enum, ok := mappingAutonomousDatabasePermissionLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseRoleEnum Enum with underlying type: string
type AutonomousDatabaseRoleEnum string

// Set of constants representing the allowable values for AutonomousDatabaseRoleEnum
const (
	AutonomousDatabaseRolePrimary         AutonomousDatabaseRoleEnum = "PRIMARY"
	AutonomousDatabaseRoleStandby         AutonomousDatabaseRoleEnum = "STANDBY"
	AutonomousDatabaseRoleDisabledStandby AutonomousDatabaseRoleEnum = "DISABLED_STANDBY"
	AutonomousDatabaseRoleBackupCopy      AutonomousDatabaseRoleEnum = "BACKUP_COPY"
	AutonomousDatabaseRoleSnapshotStandby AutonomousDatabaseRoleEnum = "SNAPSHOT_STANDBY"
)

var mappingAutonomousDatabaseRoleEnum = map[string]AutonomousDatabaseRoleEnum{
	"PRIMARY":          AutonomousDatabaseRolePrimary,
	"STANDBY":          AutonomousDatabaseRoleStandby,
	"DISABLED_STANDBY": AutonomousDatabaseRoleDisabledStandby,
	"BACKUP_COPY":      AutonomousDatabaseRoleBackupCopy,
	"SNAPSHOT_STANDBY": AutonomousDatabaseRoleSnapshotStandby,
}

var mappingAutonomousDatabaseRoleEnumLowerCase = map[string]AutonomousDatabaseRoleEnum{
	"primary":          AutonomousDatabaseRolePrimary,
	"standby":          AutonomousDatabaseRoleStandby,
	"disabled_standby": AutonomousDatabaseRoleDisabledStandby,
	"backup_copy":      AutonomousDatabaseRoleBackupCopy,
	"snapshot_standby": AutonomousDatabaseRoleSnapshotStandby,
}

// GetAutonomousDatabaseRoleEnumValues Enumerates the set of values for AutonomousDatabaseRoleEnum
func GetAutonomousDatabaseRoleEnumValues() []AutonomousDatabaseRoleEnum {
	values := make([]AutonomousDatabaseRoleEnum, 0)
	for _, v := range mappingAutonomousDatabaseRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseRoleEnumStringValues Enumerates the set of values in String for AutonomousDatabaseRoleEnum
func GetAutonomousDatabaseRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
		"BACKUP_COPY",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingAutonomousDatabaseRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseRoleEnum(val string) (AutonomousDatabaseRoleEnum, bool) {
	enum, ok := mappingAutonomousDatabaseRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseDataguardRegionTypeEnum Enum with underlying type: string
type AutonomousDatabaseDataguardRegionTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDataguardRegionTypeEnum
const (
	AutonomousDatabaseDataguardRegionTypePrimaryDgRegion       AutonomousDatabaseDataguardRegionTypeEnum = "PRIMARY_DG_REGION"
	AutonomousDatabaseDataguardRegionTypeRemoteStandbyDgRegion AutonomousDatabaseDataguardRegionTypeEnum = "REMOTE_STANDBY_DG_REGION"
)

var mappingAutonomousDatabaseDataguardRegionTypeEnum = map[string]AutonomousDatabaseDataguardRegionTypeEnum{
	"PRIMARY_DG_REGION":        AutonomousDatabaseDataguardRegionTypePrimaryDgRegion,
	"REMOTE_STANDBY_DG_REGION": AutonomousDatabaseDataguardRegionTypeRemoteStandbyDgRegion,
}

var mappingAutonomousDatabaseDataguardRegionTypeEnumLowerCase = map[string]AutonomousDatabaseDataguardRegionTypeEnum{
	"primary_dg_region":        AutonomousDatabaseDataguardRegionTypePrimaryDgRegion,
	"remote_standby_dg_region": AutonomousDatabaseDataguardRegionTypeRemoteStandbyDgRegion,
}

// GetAutonomousDatabaseDataguardRegionTypeEnumValues Enumerates the set of values for AutonomousDatabaseDataguardRegionTypeEnum
func GetAutonomousDatabaseDataguardRegionTypeEnumValues() []AutonomousDatabaseDataguardRegionTypeEnum {
	values := make([]AutonomousDatabaseDataguardRegionTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseDataguardRegionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseDataguardRegionTypeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseDataguardRegionTypeEnum
func GetAutonomousDatabaseDataguardRegionTypeEnumStringValues() []string {
	return []string{
		"PRIMARY_DG_REGION",
		"REMOTE_STANDBY_DG_REGION",
	}
}

// GetMappingAutonomousDatabaseDataguardRegionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseDataguardRegionTypeEnum(val string) (AutonomousDatabaseDataguardRegionTypeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseDataguardRegionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum Enum with underlying type: string
type AutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum
const (
	AutonomousDatabaseAutonomousMaintenanceScheduleTypeEarly   AutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum = "EARLY"
	AutonomousDatabaseAutonomousMaintenanceScheduleTypeRegular AutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum = "REGULAR"
)

var mappingAutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum = map[string]AutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum{
	"EARLY":   AutonomousDatabaseAutonomousMaintenanceScheduleTypeEarly,
	"REGULAR": AutonomousDatabaseAutonomousMaintenanceScheduleTypeRegular,
}

var mappingAutonomousDatabaseAutonomousMaintenanceScheduleTypeEnumLowerCase = map[string]AutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum{
	"early":   AutonomousDatabaseAutonomousMaintenanceScheduleTypeEarly,
	"regular": AutonomousDatabaseAutonomousMaintenanceScheduleTypeRegular,
}

// GetAutonomousDatabaseAutonomousMaintenanceScheduleTypeEnumValues Enumerates the set of values for AutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum
func GetAutonomousDatabaseAutonomousMaintenanceScheduleTypeEnumValues() []AutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum {
	values := make([]AutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseAutonomousMaintenanceScheduleTypeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum
func GetAutonomousDatabaseAutonomousMaintenanceScheduleTypeEnumStringValues() []string {
	return []string{
		"EARLY",
		"REGULAR",
	}
}

// GetMappingAutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum(val string) (AutonomousDatabaseAutonomousMaintenanceScheduleTypeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseAutonomousMaintenanceScheduleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseDatabaseEditionEnum Enum with underlying type: string
type AutonomousDatabaseDatabaseEditionEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDatabaseEditionEnum
const (
	AutonomousDatabaseDatabaseEditionStandardEdition   AutonomousDatabaseDatabaseEditionEnum = "STANDARD_EDITION"
	AutonomousDatabaseDatabaseEditionEnterpriseEdition AutonomousDatabaseDatabaseEditionEnum = "ENTERPRISE_EDITION"
)

var mappingAutonomousDatabaseDatabaseEditionEnum = map[string]AutonomousDatabaseDatabaseEditionEnum{
	"STANDARD_EDITION":   AutonomousDatabaseDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION": AutonomousDatabaseDatabaseEditionEnterpriseEdition,
}

var mappingAutonomousDatabaseDatabaseEditionEnumLowerCase = map[string]AutonomousDatabaseDatabaseEditionEnum{
	"standard_edition":   AutonomousDatabaseDatabaseEditionStandardEdition,
	"enterprise_edition": AutonomousDatabaseDatabaseEditionEnterpriseEdition,
}

// GetAutonomousDatabaseDatabaseEditionEnumValues Enumerates the set of values for AutonomousDatabaseDatabaseEditionEnum
func GetAutonomousDatabaseDatabaseEditionEnumValues() []AutonomousDatabaseDatabaseEditionEnum {
	values := make([]AutonomousDatabaseDatabaseEditionEnum, 0)
	for _, v := range mappingAutonomousDatabaseDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseDatabaseEditionEnumStringValues Enumerates the set of values in String for AutonomousDatabaseDatabaseEditionEnum
func GetAutonomousDatabaseDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
	}
}

// GetMappingAutonomousDatabaseDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseDatabaseEditionEnum(val string) (AutonomousDatabaseDatabaseEditionEnum, bool) {
	enum, ok := mappingAutonomousDatabaseDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseDisasterRecoveryRegionTypeEnum Enum with underlying type: string
type AutonomousDatabaseDisasterRecoveryRegionTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDisasterRecoveryRegionTypeEnum
const (
	AutonomousDatabaseDisasterRecoveryRegionTypePrimary AutonomousDatabaseDisasterRecoveryRegionTypeEnum = "PRIMARY"
	AutonomousDatabaseDisasterRecoveryRegionTypeRemote  AutonomousDatabaseDisasterRecoveryRegionTypeEnum = "REMOTE"
)

var mappingAutonomousDatabaseDisasterRecoveryRegionTypeEnum = map[string]AutonomousDatabaseDisasterRecoveryRegionTypeEnum{
	"PRIMARY": AutonomousDatabaseDisasterRecoveryRegionTypePrimary,
	"REMOTE":  AutonomousDatabaseDisasterRecoveryRegionTypeRemote,
}

var mappingAutonomousDatabaseDisasterRecoveryRegionTypeEnumLowerCase = map[string]AutonomousDatabaseDisasterRecoveryRegionTypeEnum{
	"primary": AutonomousDatabaseDisasterRecoveryRegionTypePrimary,
	"remote":  AutonomousDatabaseDisasterRecoveryRegionTypeRemote,
}

// GetAutonomousDatabaseDisasterRecoveryRegionTypeEnumValues Enumerates the set of values for AutonomousDatabaseDisasterRecoveryRegionTypeEnum
func GetAutonomousDatabaseDisasterRecoveryRegionTypeEnumValues() []AutonomousDatabaseDisasterRecoveryRegionTypeEnum {
	values := make([]AutonomousDatabaseDisasterRecoveryRegionTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseDisasterRecoveryRegionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseDisasterRecoveryRegionTypeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseDisasterRecoveryRegionTypeEnum
func GetAutonomousDatabaseDisasterRecoveryRegionTypeEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"REMOTE",
	}
}

// GetMappingAutonomousDatabaseDisasterRecoveryRegionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseDisasterRecoveryRegionTypeEnum(val string) (AutonomousDatabaseDisasterRecoveryRegionTypeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseDisasterRecoveryRegionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseNetServicesArchitectureEnum Enum with underlying type: string
type AutonomousDatabaseNetServicesArchitectureEnum string

// Set of constants representing the allowable values for AutonomousDatabaseNetServicesArchitectureEnum
const (
	AutonomousDatabaseNetServicesArchitectureDedicated AutonomousDatabaseNetServicesArchitectureEnum = "DEDICATED"
	AutonomousDatabaseNetServicesArchitectureShared    AutonomousDatabaseNetServicesArchitectureEnum = "SHARED"
)

var mappingAutonomousDatabaseNetServicesArchitectureEnum = map[string]AutonomousDatabaseNetServicesArchitectureEnum{
	"DEDICATED": AutonomousDatabaseNetServicesArchitectureDedicated,
	"SHARED":    AutonomousDatabaseNetServicesArchitectureShared,
}

var mappingAutonomousDatabaseNetServicesArchitectureEnumLowerCase = map[string]AutonomousDatabaseNetServicesArchitectureEnum{
	"dedicated": AutonomousDatabaseNetServicesArchitectureDedicated,
	"shared":    AutonomousDatabaseNetServicesArchitectureShared,
}

// GetAutonomousDatabaseNetServicesArchitectureEnumValues Enumerates the set of values for AutonomousDatabaseNetServicesArchitectureEnum
func GetAutonomousDatabaseNetServicesArchitectureEnumValues() []AutonomousDatabaseNetServicesArchitectureEnum {
	values := make([]AutonomousDatabaseNetServicesArchitectureEnum, 0)
	for _, v := range mappingAutonomousDatabaseNetServicesArchitectureEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseNetServicesArchitectureEnumStringValues Enumerates the set of values in String for AutonomousDatabaseNetServicesArchitectureEnum
func GetAutonomousDatabaseNetServicesArchitectureEnumStringValues() []string {
	return []string{
		"DEDICATED",
		"SHARED",
	}
}

// GetMappingAutonomousDatabaseNetServicesArchitectureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseNetServicesArchitectureEnum(val string) (AutonomousDatabaseNetServicesArchitectureEnum, bool) {
	enum, ok := mappingAutonomousDatabaseNetServicesArchitectureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
