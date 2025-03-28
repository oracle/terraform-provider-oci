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

// CreateAutonomousDatabaseCloneDetails Details to create an Oracle Autonomous Database by cloning an existing Autonomous Database.
type CreateAutonomousDatabaseCloneDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of the Autonomous Database.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that you will clone to create a new Autonomous Database.
	SourceId *string `mandatory:"true" json:"sourceId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// The character set for the autonomous database. The default is AL32UTF8. Allowed values for an Autonomous Database Serverless instance as as returned by List Autonomous Database Character Sets (https://docs.oracle.com/iaas/autonomous-database-serverless/doc/autonomous-character-set-selection.html)
	// For an Autonomous Database on dedicated infrastructure, the allowed values are:
	// AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSC5601, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P1, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS
	CharacterSet *string `mandatory:"false" json:"characterSet"`

	// The character set for the Autonomous Database. The default is AL32UTF8. Use List Autonomous Database Character Sets (https://docs.oracle.com/iaas/autonomous-database-serverless/doc/autonomous-character-set-selection.html) to list the allowed values for an Autonomous Database Serverless instance.
	// For an Autonomous Database on dedicated Exadata infrastructure, the allowed values are:
	// AL16UTF16 or UTF8.
	NcharacterSet *string `mandatory:"false" json:"ncharacterSet"`

	// The database name. The name must begin with an alphabetic character and can contain a maximum of 30 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy. It is required in all cases except when creating a cross-region Autonomous Data Guard standby instance or a cross-region disaster recovery standby instance.
	DbName *string `mandatory:"false" json:"dbName"`

	// The number of CPU cores to be made available to the database. For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `ocpuCount` parameter.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// Retention period, in days, for long-term backups
	BackupRetentionPeriodInDays *int `mandatory:"false" json:"backupRetentionPeriodInDays"`

	// The compute amount (CPUs) available to the database. Minimum and maximum values depend on the compute model and whether the database is an Autonomous Database Serverless instance or an Autonomous Database on Dedicated Exadata Infrastructure.
	// The 'ECPU' compute model requires a minimum value of one, for databases in the elastic resource pool and minimum value of two, otherwise. Required when using the `computeModel` parameter. When using `cpuCoreCount` parameter, it is an error to specify computeCount to a non-null value. Providing `computeModel` and `computeCount` is the preferred method for both OCPU and ECPU.
	ComputeCount *float32 `mandatory:"false" json:"computeCount"`

	// The number of OCPU cores to be made available to the database.
	// The following points apply:
	// - For Autonomous Databases on Dedicated Exadata infrastructure, to provision less than 1 core, enter a fractional value in an increment of 0.1. For example, you can provision 0.3 or 0.4 cores, but not 0.35 cores. (Note that fractional OCPU values are not supported for Autonomous Database Serverless instances.)
	// - To provision 1 or more cores, you must enter an integer between 1 and the maximum number of cores available for the infrastructure shape. For example, you can provision 2 cores or 3 cores, but not 2.5 cores. This applies to an Autonomous Database Serverless instance or an Autonomous Database on Dedicated Exadata Infrastructure.
	// - For Autonomous Database Serverless instances, this parameter is not used.
	// For Autonomous Databases on Dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `cpuCoreCount` parameter.
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// The size, in terabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed. For Autonomous Databases on dedicated Exadata infrastructure, the maximum storage value is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// A full Exadata service is allocated when the Autonomous Database size is set to the upper limit (384 TB).
	// **Note:** This parameter cannot be used with the `dataStorageSizeInGBs` parameter.
	DataStorageSizeInTBs *int `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The size, in gigabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed. The maximum storage value is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Notes**
	// - This parameter is only supported for dedicated Exadata infrastructure.
	// - This parameter cannot be used with the `dataStorageSizeInTBs` parameter.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isLocalDataGuardEnabled
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	EncryptionKey AutonomousDatabaseEncryptionKeyDetails `mandatory:"false" json:"encryptionKey"`

	// **Important** The `adminPassword` or `secretId` must be specified for all Autonomous Databases except for refreshable clones. The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing.
	// This cannot be used in conjunction with with OCI vault secrets (secretId).
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

	// The user-friendly name for the Autonomous Database. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The maximum number of CPUs allowed with a Bring Your Own License (BYOL), including those used for auto-scaling, disaster recovery, tools, etc. Any CPU usage above this limit is considered as License Included and billed.
	ByolComputeCountLimit *float32 `mandatory:"false" json:"byolComputeCountLimit"`

	// If set to `TRUE`, indicates that an Autonomous Database preview version is being provisioned, and that the preview version's terms of service have been accepted. Note that preview version software is only available for Autonomous Database Serverless instances (https://docs.oracle.com/en/cloud/paas/autonomous-database/serverless/).
	IsPreviewVersionWithServiceTermsAccepted *bool `mandatory:"false" json:"isPreviewVersionWithServiceTermsAccepted"`

	// Indicates if auto scaling is enabled for the Autonomous Database CPU core count. The default value is `TRUE`.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// Autonomous Database for Developers are fixed-shape Autonomous Databases that developers can use to build and test new applications. On Serverless, these are low-cost and billed per instance, on Dedicated and Cloud@Customer there is no additional cost to create Developer databases. Developer databases come with limited resources and is not intended for large-scale testing and production deployments. When you need more compute or storage resources, you may upgrade to a full paid production database.
	IsDevTier *bool `mandatory:"false" json:"isDevTier"`

	// True if the database is on dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html).
	IsDedicated *bool `mandatory:"false" json:"isDedicated"`

	// The Autonomous Container Database OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Used only by Autonomous Database on Dedicated Exadata Infrastructure.
	AutonomousContainerDatabaseId *string `mandatory:"false" json:"autonomousContainerDatabaseId"`

	// The percentage of the System Global Area(SGA) assigned to In-Memory tables in Autonomous Database. This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform.
	InMemoryPercentage *int `mandatory:"false" json:"inMemoryPercentage"`

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

	// **Deprecated.** Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
	IsDataGuardEnabled *bool `mandatory:"false" json:"isDataGuardEnabled"`

	// Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
	IsLocalDataGuardEnabled *bool `mandatory:"false" json:"isLocalDataGuardEnabled"`

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

	// The resource's private endpoint label.
	// - Setting the endpoint label to a non-empty string creates a private endpoint database.
	// - Resetting the endpoint label to an empty string, after the creation of the private endpoint database, changes the private endpoint database to a public endpoint database.
	// - Setting the endpoint label to a non-empty string value, updates to a new private endpoint database, when the database is disabled and re-enabled.
	// This setting cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, dbWorkload, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	PrivateEndpointLabel *string `mandatory:"false" json:"privateEndpointLabel"`

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

	// The private endpoint Ip address for the resource.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// A valid Oracle Database version for Autonomous Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// Customer Contacts.
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

	// The unique identifier for leader autonomous database OCID OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ResourcePoolLeaderId *string `mandatory:"false" json:"resourcePoolLeaderId"`

	ResourcePoolSummary *ResourcePoolSummary `mandatory:"false" json:"resourcePoolSummary"`

	// The list of scheduled operations. Consists of values such as dayOfWeek, scheduledStartTime, scheduledStopTime.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	ScheduledOperations []ScheduledOperationDetails `mandatory:"false" json:"scheduledOperations"`

	// Indicates if auto scaling is enabled for the Autonomous Database storage. The default value is `FALSE`.
	IsAutoScalingForStorageEnabled *bool `mandatory:"false" json:"isAutoScalingForStorageEnabled"`

	// The list of database tools details.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, isLocalDataGuardEnabled, or isFreeTier.
	DbToolsDetails []DatabaseTool `mandatory:"false" json:"dbToolsDetails"`

	// True if the Autonomous Database is backup retention locked.
	IsBackupRetentionLocked *bool `mandatory:"false" json:"isBackupRetentionLocked"`

	// The OCI vault secret [/Content/General/Concepts/identifiers.htm]OCID.
	// This cannot be used in conjunction with adminPassword.
	SecretId *string `mandatory:"false" json:"secretId"`

	// The version of the vault secret. If no version is specified, the latest version will be used.
	SecretVersionNumber *int `mandatory:"false" json:"secretVersionNumber"`

	// The Autonomous Database clone type.
	CloneType CreateAutonomousDatabaseCloneDetailsCloneTypeEnum `mandatory:"true" json:"cloneType"`

	// The Oracle Database Edition that applies to the Autonomous databases. This parameter accepts options `STANDARD_EDITION` and `ENTERPRISE_EDITION`.
	DatabaseEdition AutonomousDatabaseSummaryDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

	// The compute model of the Autonomous Database. This is required if using the `computeCount` parameter. If using `cpuCoreCount` then it is an error to specify `computeModel` to a non-null value. ECPU compute model is the recommended model and OCPU compute model is legacy.
	ComputeModel CreateAutonomousDatabaseBaseComputeModelEnum `mandatory:"false" json:"computeModel,omitempty"`

	// The Autonomous Database workload type. The following values are valid:
	// - OLTP - indicates an Autonomous Transaction Processing database
	// - DW - indicates an Autonomous Data Warehouse database
	// - AJD - indicates an Autonomous JSON Database
	// - APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	DbWorkload CreateAutonomousDatabaseBaseDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Oracle Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null. It is already set at the
	// Autonomous Exadata Infrastructure level. When provisioning an Autonomous Database Serverless  (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) database, if a value is not specified, the system defaults the value to `BRING_YOUR_OWN_LICENSE`. Bring your own license (BYOL) also allows you to select the DB edition using the optional parameter.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, dataStorageSizeInTBs, adminPassword, isMTLSConnectionRequired, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	LicenseModel CreateAutonomousDatabaseBaseLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The maintenance schedule type of the Autonomous Database Serverless. An EARLY maintenance schedule
	// follows a schedule applying patches prior to the REGULAR schedule. A REGULAR maintenance schedule follows the normal cycle
	AutonomousMaintenanceScheduleType CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum `mandatory:"false" json:"autonomousMaintenanceScheduleType,omitempty"`
}

// GetSubscriptionId returns SubscriptionId
func (m CreateAutonomousDatabaseCloneDetails) GetSubscriptionId() *string {
	return m.SubscriptionId
}

// GetCompartmentId returns CompartmentId
func (m CreateAutonomousDatabaseCloneDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetCharacterSet returns CharacterSet
func (m CreateAutonomousDatabaseCloneDetails) GetCharacterSet() *string {
	return m.CharacterSet
}

// GetNcharacterSet returns NcharacterSet
func (m CreateAutonomousDatabaseCloneDetails) GetNcharacterSet() *string {
	return m.NcharacterSet
}

// GetDbName returns DbName
func (m CreateAutonomousDatabaseCloneDetails) GetDbName() *string {
	return m.DbName
}

// GetCpuCoreCount returns CpuCoreCount
func (m CreateAutonomousDatabaseCloneDetails) GetCpuCoreCount() *int {
	return m.CpuCoreCount
}

// GetBackupRetentionPeriodInDays returns BackupRetentionPeriodInDays
func (m CreateAutonomousDatabaseCloneDetails) GetBackupRetentionPeriodInDays() *int {
	return m.BackupRetentionPeriodInDays
}

// GetComputeModel returns ComputeModel
func (m CreateAutonomousDatabaseCloneDetails) GetComputeModel() CreateAutonomousDatabaseBaseComputeModelEnum {
	return m.ComputeModel
}

// GetComputeCount returns ComputeCount
func (m CreateAutonomousDatabaseCloneDetails) GetComputeCount() *float32 {
	return m.ComputeCount
}

// GetOcpuCount returns OcpuCount
func (m CreateAutonomousDatabaseCloneDetails) GetOcpuCount() *float32 {
	return m.OcpuCount
}

// GetDbWorkload returns DbWorkload
func (m CreateAutonomousDatabaseCloneDetails) GetDbWorkload() CreateAutonomousDatabaseBaseDbWorkloadEnum {
	return m.DbWorkload
}

// GetDataStorageSizeInTBs returns DataStorageSizeInTBs
func (m CreateAutonomousDatabaseCloneDetails) GetDataStorageSizeInTBs() *int {
	return m.DataStorageSizeInTBs
}

// GetDataStorageSizeInGBs returns DataStorageSizeInGBs
func (m CreateAutonomousDatabaseCloneDetails) GetDataStorageSizeInGBs() *int {
	return m.DataStorageSizeInGBs
}

// GetIsFreeTier returns IsFreeTier
func (m CreateAutonomousDatabaseCloneDetails) GetIsFreeTier() *bool {
	return m.IsFreeTier
}

// GetKmsKeyId returns KmsKeyId
func (m CreateAutonomousDatabaseCloneDetails) GetKmsKeyId() *string {
	return m.KmsKeyId
}

// GetVaultId returns VaultId
func (m CreateAutonomousDatabaseCloneDetails) GetVaultId() *string {
	return m.VaultId
}

// GetEncryptionKey returns EncryptionKey
func (m CreateAutonomousDatabaseCloneDetails) GetEncryptionKey() AutonomousDatabaseEncryptionKeyDetails {
	return m.EncryptionKey
}

// GetAdminPassword returns AdminPassword
func (m CreateAutonomousDatabaseCloneDetails) GetAdminPassword() *string {
	return m.AdminPassword
}

// GetDisplayName returns DisplayName
func (m CreateAutonomousDatabaseCloneDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetLicenseModel returns LicenseModel
func (m CreateAutonomousDatabaseCloneDetails) GetLicenseModel() CreateAutonomousDatabaseBaseLicenseModelEnum {
	return m.LicenseModel
}

// GetByolComputeCountLimit returns ByolComputeCountLimit
func (m CreateAutonomousDatabaseCloneDetails) GetByolComputeCountLimit() *float32 {
	return m.ByolComputeCountLimit
}

// GetIsPreviewVersionWithServiceTermsAccepted returns IsPreviewVersionWithServiceTermsAccepted
func (m CreateAutonomousDatabaseCloneDetails) GetIsPreviewVersionWithServiceTermsAccepted() *bool {
	return m.IsPreviewVersionWithServiceTermsAccepted
}

// GetIsAutoScalingEnabled returns IsAutoScalingEnabled
func (m CreateAutonomousDatabaseCloneDetails) GetIsAutoScalingEnabled() *bool {
	return m.IsAutoScalingEnabled
}

// GetIsDevTier returns IsDevTier
func (m CreateAutonomousDatabaseCloneDetails) GetIsDevTier() *bool {
	return m.IsDevTier
}

// GetIsDedicated returns IsDedicated
func (m CreateAutonomousDatabaseCloneDetails) GetIsDedicated() *bool {
	return m.IsDedicated
}

// GetAutonomousContainerDatabaseId returns AutonomousContainerDatabaseId
func (m CreateAutonomousDatabaseCloneDetails) GetAutonomousContainerDatabaseId() *string {
	return m.AutonomousContainerDatabaseId
}

// GetInMemoryPercentage returns InMemoryPercentage
func (m CreateAutonomousDatabaseCloneDetails) GetInMemoryPercentage() *int {
	return m.InMemoryPercentage
}

// GetIsAccessControlEnabled returns IsAccessControlEnabled
func (m CreateAutonomousDatabaseCloneDetails) GetIsAccessControlEnabled() *bool {
	return m.IsAccessControlEnabled
}

// GetWhitelistedIps returns WhitelistedIps
func (m CreateAutonomousDatabaseCloneDetails) GetWhitelistedIps() []string {
	return m.WhitelistedIps
}

// GetArePrimaryWhitelistedIpsUsed returns ArePrimaryWhitelistedIpsUsed
func (m CreateAutonomousDatabaseCloneDetails) GetArePrimaryWhitelistedIpsUsed() *bool {
	return m.ArePrimaryWhitelistedIpsUsed
}

// GetStandbyWhitelistedIps returns StandbyWhitelistedIps
func (m CreateAutonomousDatabaseCloneDetails) GetStandbyWhitelistedIps() []string {
	return m.StandbyWhitelistedIps
}

// GetIsDataGuardEnabled returns IsDataGuardEnabled
func (m CreateAutonomousDatabaseCloneDetails) GetIsDataGuardEnabled() *bool {
	return m.IsDataGuardEnabled
}

// GetIsLocalDataGuardEnabled returns IsLocalDataGuardEnabled
func (m CreateAutonomousDatabaseCloneDetails) GetIsLocalDataGuardEnabled() *bool {
	return m.IsLocalDataGuardEnabled
}

// GetSubnetId returns SubnetId
func (m CreateAutonomousDatabaseCloneDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetNsgIds returns NsgIds
func (m CreateAutonomousDatabaseCloneDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetPrivateEndpointLabel returns PrivateEndpointLabel
func (m CreateAutonomousDatabaseCloneDetails) GetPrivateEndpointLabel() *string {
	return m.PrivateEndpointLabel
}

// GetFreeformTags returns FreeformTags
func (m CreateAutonomousDatabaseCloneDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateAutonomousDatabaseCloneDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSecurityAttributes returns SecurityAttributes
func (m CreateAutonomousDatabaseCloneDetails) GetSecurityAttributes() map[string]map[string]interface{} {
	return m.SecurityAttributes
}

// GetPrivateEndpointIp returns PrivateEndpointIp
func (m CreateAutonomousDatabaseCloneDetails) GetPrivateEndpointIp() *string {
	return m.PrivateEndpointIp
}

// GetDbVersion returns DbVersion
func (m CreateAutonomousDatabaseCloneDetails) GetDbVersion() *string {
	return m.DbVersion
}

// GetCustomerContacts returns CustomerContacts
func (m CreateAutonomousDatabaseCloneDetails) GetCustomerContacts() []CustomerContact {
	return m.CustomerContacts
}

// GetIsMtlsConnectionRequired returns IsMtlsConnectionRequired
func (m CreateAutonomousDatabaseCloneDetails) GetIsMtlsConnectionRequired() *bool {
	return m.IsMtlsConnectionRequired
}

// GetResourcePoolLeaderId returns ResourcePoolLeaderId
func (m CreateAutonomousDatabaseCloneDetails) GetResourcePoolLeaderId() *string {
	return m.ResourcePoolLeaderId
}

// GetResourcePoolSummary returns ResourcePoolSummary
func (m CreateAutonomousDatabaseCloneDetails) GetResourcePoolSummary() *ResourcePoolSummary {
	return m.ResourcePoolSummary
}

// GetAutonomousMaintenanceScheduleType returns AutonomousMaintenanceScheduleType
func (m CreateAutonomousDatabaseCloneDetails) GetAutonomousMaintenanceScheduleType() CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum {
	return m.AutonomousMaintenanceScheduleType
}

// GetScheduledOperations returns ScheduledOperations
func (m CreateAutonomousDatabaseCloneDetails) GetScheduledOperations() []ScheduledOperationDetails {
	return m.ScheduledOperations
}

// GetIsAutoScalingForStorageEnabled returns IsAutoScalingForStorageEnabled
func (m CreateAutonomousDatabaseCloneDetails) GetIsAutoScalingForStorageEnabled() *bool {
	return m.IsAutoScalingForStorageEnabled
}

// GetDatabaseEdition returns DatabaseEdition
func (m CreateAutonomousDatabaseCloneDetails) GetDatabaseEdition() AutonomousDatabaseSummaryDatabaseEditionEnum {
	return m.DatabaseEdition
}

// GetDbToolsDetails returns DbToolsDetails
func (m CreateAutonomousDatabaseCloneDetails) GetDbToolsDetails() []DatabaseTool {
	return m.DbToolsDetails
}

// GetIsBackupRetentionLocked returns IsBackupRetentionLocked
func (m CreateAutonomousDatabaseCloneDetails) GetIsBackupRetentionLocked() *bool {
	return m.IsBackupRetentionLocked
}

// GetSecretId returns SecretId
func (m CreateAutonomousDatabaseCloneDetails) GetSecretId() *string {
	return m.SecretId
}

// GetSecretVersionNumber returns SecretVersionNumber
func (m CreateAutonomousDatabaseCloneDetails) GetSecretVersionNumber() *int {
	return m.SecretVersionNumber
}

func (m CreateAutonomousDatabaseCloneDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAutonomousDatabaseCloneDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateAutonomousDatabaseCloneDetailsCloneTypeEnum(string(m.CloneType)); !ok && m.CloneType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CloneType: %s. Supported values are: %s.", m.CloneType, strings.Join(GetCreateAutonomousDatabaseCloneDetailsCloneTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousDatabaseSummaryDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetAutonomousDatabaseSummaryDatabaseEditionEnumStringValues(), ",")))
	}
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
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateAutonomousDatabaseCloneDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAutonomousDatabaseCloneDetails CreateAutonomousDatabaseCloneDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateAutonomousDatabaseCloneDetails
	}{
		"DATABASE",
		(MarshalTypeCreateAutonomousDatabaseCloneDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateAutonomousDatabaseCloneDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SubscriptionId                           *string                                                           `json:"subscriptionId"`
		CharacterSet                             *string                                                           `json:"characterSet"`
		NcharacterSet                            *string                                                           `json:"ncharacterSet"`
		DbName                                   *string                                                           `json:"dbName"`
		CpuCoreCount                             *int                                                              `json:"cpuCoreCount"`
		BackupRetentionPeriodInDays              *int                                                              `json:"backupRetentionPeriodInDays"`
		ComputeModel                             CreateAutonomousDatabaseBaseComputeModelEnum                      `json:"computeModel"`
		ComputeCount                             *float32                                                          `json:"computeCount"`
		OcpuCount                                *float32                                                          `json:"ocpuCount"`
		DbWorkload                               CreateAutonomousDatabaseBaseDbWorkloadEnum                        `json:"dbWorkload"`
		DataStorageSizeInTBs                     *int                                                              `json:"dataStorageSizeInTBs"`
		DataStorageSizeInGBs                     *int                                                              `json:"dataStorageSizeInGBs"`
		IsFreeTier                               *bool                                                             `json:"isFreeTier"`
		KmsKeyId                                 *string                                                           `json:"kmsKeyId"`
		VaultId                                  *string                                                           `json:"vaultId"`
		EncryptionKey                            autonomousdatabaseencryptionkeydetails                            `json:"encryptionKey"`
		AdminPassword                            *string                                                           `json:"adminPassword"`
		DisplayName                              *string                                                           `json:"displayName"`
		LicenseModel                             CreateAutonomousDatabaseBaseLicenseModelEnum                      `json:"licenseModel"`
		ByolComputeCountLimit                    *float32                                                          `json:"byolComputeCountLimit"`
		IsPreviewVersionWithServiceTermsAccepted *bool                                                             `json:"isPreviewVersionWithServiceTermsAccepted"`
		IsAutoScalingEnabled                     *bool                                                             `json:"isAutoScalingEnabled"`
		IsDevTier                                *bool                                                             `json:"isDevTier"`
		IsDedicated                              *bool                                                             `json:"isDedicated"`
		AutonomousContainerDatabaseId            *string                                                           `json:"autonomousContainerDatabaseId"`
		InMemoryPercentage                       *int                                                              `json:"inMemoryPercentage"`
		IsAccessControlEnabled                   *bool                                                             `json:"isAccessControlEnabled"`
		WhitelistedIps                           []string                                                          `json:"whitelistedIps"`
		ArePrimaryWhitelistedIpsUsed             *bool                                                             `json:"arePrimaryWhitelistedIpsUsed"`
		StandbyWhitelistedIps                    []string                                                          `json:"standbyWhitelistedIps"`
		IsDataGuardEnabled                       *bool                                                             `json:"isDataGuardEnabled"`
		IsLocalDataGuardEnabled                  *bool                                                             `json:"isLocalDataGuardEnabled"`
		SubnetId                                 *string                                                           `json:"subnetId"`
		NsgIds                                   []string                                                          `json:"nsgIds"`
		PrivateEndpointLabel                     *string                                                           `json:"privateEndpointLabel"`
		FreeformTags                             map[string]string                                                 `json:"freeformTags"`
		DefinedTags                              map[string]map[string]interface{}                                 `json:"definedTags"`
		SecurityAttributes                       map[string]map[string]interface{}                                 `json:"securityAttributes"`
		PrivateEndpointIp                        *string                                                           `json:"privateEndpointIp"`
		DbVersion                                *string                                                           `json:"dbVersion"`
		CustomerContacts                         []CustomerContact                                                 `json:"customerContacts"`
		IsMtlsConnectionRequired                 *bool                                                             `json:"isMtlsConnectionRequired"`
		ResourcePoolLeaderId                     *string                                                           `json:"resourcePoolLeaderId"`
		ResourcePoolSummary                      *ResourcePoolSummary                                              `json:"resourcePoolSummary"`
		AutonomousMaintenanceScheduleType        CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum `json:"autonomousMaintenanceScheduleType"`
		ScheduledOperations                      []ScheduledOperationDetails                                       `json:"scheduledOperations"`
		IsAutoScalingForStorageEnabled           *bool                                                             `json:"isAutoScalingForStorageEnabled"`
		DatabaseEdition                          AutonomousDatabaseSummaryDatabaseEditionEnum                      `json:"databaseEdition"`
		DbToolsDetails                           []DatabaseTool                                                    `json:"dbToolsDetails"`
		IsBackupRetentionLocked                  *bool                                                             `json:"isBackupRetentionLocked"`
		SecretId                                 *string                                                           `json:"secretId"`
		SecretVersionNumber                      *int                                                              `json:"secretVersionNumber"`
		CompartmentId                            *string                                                           `json:"compartmentId"`
		SourceId                                 *string                                                           `json:"sourceId"`
		CloneType                                CreateAutonomousDatabaseCloneDetailsCloneTypeEnum                 `json:"cloneType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.SubscriptionId = model.SubscriptionId

	m.CharacterSet = model.CharacterSet

	m.NcharacterSet = model.NcharacterSet

	m.DbName = model.DbName

	m.CpuCoreCount = model.CpuCoreCount

	m.BackupRetentionPeriodInDays = model.BackupRetentionPeriodInDays

	m.ComputeModel = model.ComputeModel

	m.ComputeCount = model.ComputeCount

	m.OcpuCount = model.OcpuCount

	m.DbWorkload = model.DbWorkload

	m.DataStorageSizeInTBs = model.DataStorageSizeInTBs

	m.DataStorageSizeInGBs = model.DataStorageSizeInGBs

	m.IsFreeTier = model.IsFreeTier

	m.KmsKeyId = model.KmsKeyId

	m.VaultId = model.VaultId

	nn, e = model.EncryptionKey.UnmarshalPolymorphicJSON(model.EncryptionKey.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EncryptionKey = nn.(AutonomousDatabaseEncryptionKeyDetails)
	} else {
		m.EncryptionKey = nil
	}

	m.AdminPassword = model.AdminPassword

	m.DisplayName = model.DisplayName

	m.LicenseModel = model.LicenseModel

	m.ByolComputeCountLimit = model.ByolComputeCountLimit

	m.IsPreviewVersionWithServiceTermsAccepted = model.IsPreviewVersionWithServiceTermsAccepted

	m.IsAutoScalingEnabled = model.IsAutoScalingEnabled

	m.IsDevTier = model.IsDevTier

	m.IsDedicated = model.IsDedicated

	m.AutonomousContainerDatabaseId = model.AutonomousContainerDatabaseId

	m.InMemoryPercentage = model.InMemoryPercentage

	m.IsAccessControlEnabled = model.IsAccessControlEnabled

	m.WhitelistedIps = make([]string, len(model.WhitelistedIps))
	copy(m.WhitelistedIps, model.WhitelistedIps)
	m.ArePrimaryWhitelistedIpsUsed = model.ArePrimaryWhitelistedIpsUsed

	m.StandbyWhitelistedIps = make([]string, len(model.StandbyWhitelistedIps))
	copy(m.StandbyWhitelistedIps, model.StandbyWhitelistedIps)
	m.IsDataGuardEnabled = model.IsDataGuardEnabled

	m.IsLocalDataGuardEnabled = model.IsLocalDataGuardEnabled

	m.SubnetId = model.SubnetId

	m.NsgIds = make([]string, len(model.NsgIds))
	copy(m.NsgIds, model.NsgIds)
	m.PrivateEndpointLabel = model.PrivateEndpointLabel

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SecurityAttributes = model.SecurityAttributes

	m.PrivateEndpointIp = model.PrivateEndpointIp

	m.DbVersion = model.DbVersion

	m.CustomerContacts = make([]CustomerContact, len(model.CustomerContacts))
	copy(m.CustomerContacts, model.CustomerContacts)
	m.IsMtlsConnectionRequired = model.IsMtlsConnectionRequired

	m.ResourcePoolLeaderId = model.ResourcePoolLeaderId

	m.ResourcePoolSummary = model.ResourcePoolSummary

	m.AutonomousMaintenanceScheduleType = model.AutonomousMaintenanceScheduleType

	m.ScheduledOperations = make([]ScheduledOperationDetails, len(model.ScheduledOperations))
	copy(m.ScheduledOperations, model.ScheduledOperations)
	m.IsAutoScalingForStorageEnabled = model.IsAutoScalingForStorageEnabled

	m.DatabaseEdition = model.DatabaseEdition

	m.DbToolsDetails = make([]DatabaseTool, len(model.DbToolsDetails))
	copy(m.DbToolsDetails, model.DbToolsDetails)
	m.IsBackupRetentionLocked = model.IsBackupRetentionLocked

	m.SecretId = model.SecretId

	m.SecretVersionNumber = model.SecretVersionNumber

	m.CompartmentId = model.CompartmentId

	m.SourceId = model.SourceId

	m.CloneType = model.CloneType

	return
}

// CreateAutonomousDatabaseCloneDetailsCloneTypeEnum Enum with underlying type: string
type CreateAutonomousDatabaseCloneDetailsCloneTypeEnum string

// Set of constants representing the allowable values for CreateAutonomousDatabaseCloneDetailsCloneTypeEnum
const (
	CreateAutonomousDatabaseCloneDetailsCloneTypeFull     CreateAutonomousDatabaseCloneDetailsCloneTypeEnum = "FULL"
	CreateAutonomousDatabaseCloneDetailsCloneTypeMetadata CreateAutonomousDatabaseCloneDetailsCloneTypeEnum = "METADATA"
	CreateAutonomousDatabaseCloneDetailsCloneTypePartial  CreateAutonomousDatabaseCloneDetailsCloneTypeEnum = "PARTIAL"
)

var mappingCreateAutonomousDatabaseCloneDetailsCloneTypeEnum = map[string]CreateAutonomousDatabaseCloneDetailsCloneTypeEnum{
	"FULL":     CreateAutonomousDatabaseCloneDetailsCloneTypeFull,
	"METADATA": CreateAutonomousDatabaseCloneDetailsCloneTypeMetadata,
	"PARTIAL":  CreateAutonomousDatabaseCloneDetailsCloneTypePartial,
}

var mappingCreateAutonomousDatabaseCloneDetailsCloneTypeEnumLowerCase = map[string]CreateAutonomousDatabaseCloneDetailsCloneTypeEnum{
	"full":     CreateAutonomousDatabaseCloneDetailsCloneTypeFull,
	"metadata": CreateAutonomousDatabaseCloneDetailsCloneTypeMetadata,
	"partial":  CreateAutonomousDatabaseCloneDetailsCloneTypePartial,
}

// GetCreateAutonomousDatabaseCloneDetailsCloneTypeEnumValues Enumerates the set of values for CreateAutonomousDatabaseCloneDetailsCloneTypeEnum
func GetCreateAutonomousDatabaseCloneDetailsCloneTypeEnumValues() []CreateAutonomousDatabaseCloneDetailsCloneTypeEnum {
	values := make([]CreateAutonomousDatabaseCloneDetailsCloneTypeEnum, 0)
	for _, v := range mappingCreateAutonomousDatabaseCloneDetailsCloneTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousDatabaseCloneDetailsCloneTypeEnumStringValues Enumerates the set of values in String for CreateAutonomousDatabaseCloneDetailsCloneTypeEnum
func GetCreateAutonomousDatabaseCloneDetailsCloneTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"METADATA",
		"PARTIAL",
	}
}

// GetMappingCreateAutonomousDatabaseCloneDetailsCloneTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousDatabaseCloneDetailsCloneTypeEnum(val string) (CreateAutonomousDatabaseCloneDetailsCloneTypeEnum, bool) {
	enum, ok := mappingCreateAutonomousDatabaseCloneDetailsCloneTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
