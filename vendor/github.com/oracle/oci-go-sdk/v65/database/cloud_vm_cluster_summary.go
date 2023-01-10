// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// CloudVmClusterSummary Details of the cloud VM cluster. Applies to Exadata Cloud Service instances only.
type CloudVmClusterSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud VM cluster.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the availability domain that the cloud Exadata infrastructure resource is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet associated with the cloud VM cluster.
	// **Subnet Restrictions:**
	// - For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and backup subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The model name of the Exadata hardware running the cloud VM cluster.
	Shape *string `mandatory:"true" json:"shape"`

	// The current state of the cloud VM cluster.
	LifecycleState CloudVmClusterSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The user-friendly name for the cloud VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The hostname for the cloud VM cluster.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The domain name for the cloud VM cluster.
	Domain *string `mandatory:"true" json:"domain"`

	// The number of CPU cores enabled on the cloud VM cluster.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
	CloudExadataInfrastructureId *string `mandatory:"true" json:"cloudExadataInfrastructureId"`

	// The public key portion of one or more key pairs used for SSH access to the cloud VM cluster.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the cloud VM cluster.
	// **Subnet Restriction:** See the subnet restrictions information for **subnetId**.
	BackupSubnetId *string `mandatory:"false" json:"backupSubnetId"`

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems.
	BackupNetworkNsgIds []string `mandatory:"false" json:"backupNetworkNsgIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last maintenance update history entry. This value is updated when a maintenance update starts.
	LastUpdateHistoryEntryId *string `mandatory:"false" json:"lastUpdateHistoryEntryId"`

	// The port number configured for the listener on the cloud VM cluster.
	ListenerPort *int64 `mandatory:"false" json:"listenerPort"`

	// The number of nodes in the cloud VM cluster.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The storage allocation for the disk group, in gigabytes (GB).
	StorageSizeInGBs *int `mandatory:"false" json:"storageSizeInGBs"`

	// The date and time that the cloud VM cluster was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time zone of the cloud VM cluster. For details, see Exadata Infrastructure Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The number of OCPU cores to enable on the cloud VM cluster. Only 1 decimal place is allowed for the fractional part.
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// The memory to be allocated in GBs.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The local node storage to be allocated in GBs.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// The data disk group size to be allocated in TBs.
	DataStorageSizeInTBs *float64 `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The list of Db servers.
	DbServers []string `mandatory:"false" json:"dbServers"`

	// The cluster name for cloud VM cluster. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive.
	ClusterName *string `mandatory:"false" json:"clusterName"`

	// The percentage assigned to DATA storage (user data and database files).
	// The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Accepted values are 35, 40, 60 and 80. The default is 80 percent assigned to DATA storage. See Storage Configuration (https://docs.cloud.oracle.com/Content/Database/Concepts/exaoverview.htm#Exadata) in the Exadata documentation for details on the impact of the configuration settings on storage.
	DataStoragePercentage *int `mandatory:"false" json:"dataStoragePercentage"`

	// If true, database backup on local Exadata storage is configured for the cloud VM cluster. If false, database backup on local Exadata storage is not available in the cloud VM cluster.
	IsLocalBackupEnabled *bool `mandatory:"false" json:"isLocalBackupEnabled"`

	// If true, sparse disk group is configured for the cloud VM cluster. If false, sparse disk group is not created.
	IsSparseDiskgroupEnabled *bool `mandatory:"false" json:"isSparseDiskgroupEnabled"`

	// A valid Oracle Grid Infrastructure (GI) software version.
	GiVersion *string `mandatory:"false" json:"giVersion"`

	// Operating system version of the image.
	SystemVersion *string `mandatory:"false" json:"systemVersion"`

	// The Oracle license model that applies to the cloud VM cluster. The default is LICENSE_INCLUDED.
	LicenseModel CloudVmClusterSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The type of redundancy configured for the cloud Vm cluster.
	// NORMAL is 2-way redundancy.
	// HIGH is 3-way redundancy.
	DiskRedundancy CloudVmClusterSummaryDiskRedundancyEnum `mandatory:"false" json:"diskRedundancy,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Single Client Access Name (SCAN) IP addresses associated with the cloud VM cluster.
	// SCAN IP addresses are typically used for load balancing and are not assigned to any interface.
	// Oracle Clusterware directs the requests to the appropriate nodes in the cluster.
	// **Note:** For a single-node DB system, this list is empty.
	ScanIpIds []string `mandatory:"false" json:"scanIpIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the virtual IP (VIP) addresses associated with the cloud VM cluster.
	// The Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the Exadata Cloud Service instance to
	// enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.
	// **Note:** For a single-node DB system, this list is empty.
	VipIds []string `mandatory:"false" json:"vipIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DNS record for the SCAN IP addresses that are associated with the cloud VM cluster.
	ScanDnsRecordId *string `mandatory:"false" json:"scanDnsRecordId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The FQDN of the DNS record for the SCAN IP addresses that are associated with the cloud VM cluster.
	ScanDnsName *string `mandatory:"false" json:"scanDnsName"`

	// The OCID of the zone the cloud VM cluster is associated with.
	ZoneId *string `mandatory:"false" json:"zoneId"`

	// The TCP Single Client Access Name (SCAN) port. The default port is 1521.
	ScanListenerPortTcp *int `mandatory:"false" json:"scanListenerPortTcp"`

	// The TCPS Single Client Access Name (SCAN) port. The default port is 2484.
	ScanListenerPortTcpSsl *int `mandatory:"false" json:"scanListenerPortTcpSsl"`

	DataCollectionOptions *DataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`
}

func (m CloudVmClusterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudVmClusterSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudVmClusterSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCloudVmClusterSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCloudVmClusterSummaryLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCloudVmClusterSummaryLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCloudVmClusterSummaryDiskRedundancyEnum(string(m.DiskRedundancy)); !ok && m.DiskRedundancy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiskRedundancy: %s. Supported values are: %s.", m.DiskRedundancy, strings.Join(GetCloudVmClusterSummaryDiskRedundancyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudVmClusterSummaryLifecycleStateEnum Enum with underlying type: string
type CloudVmClusterSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for CloudVmClusterSummaryLifecycleStateEnum
const (
	CloudVmClusterSummaryLifecycleStateProvisioning          CloudVmClusterSummaryLifecycleStateEnum = "PROVISIONING"
	CloudVmClusterSummaryLifecycleStateAvailable             CloudVmClusterSummaryLifecycleStateEnum = "AVAILABLE"
	CloudVmClusterSummaryLifecycleStateUpdating              CloudVmClusterSummaryLifecycleStateEnum = "UPDATING"
	CloudVmClusterSummaryLifecycleStateTerminating           CloudVmClusterSummaryLifecycleStateEnum = "TERMINATING"
	CloudVmClusterSummaryLifecycleStateTerminated            CloudVmClusterSummaryLifecycleStateEnum = "TERMINATED"
	CloudVmClusterSummaryLifecycleStateFailed                CloudVmClusterSummaryLifecycleStateEnum = "FAILED"
	CloudVmClusterSummaryLifecycleStateMaintenanceInProgress CloudVmClusterSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingCloudVmClusterSummaryLifecycleStateEnum = map[string]CloudVmClusterSummaryLifecycleStateEnum{
	"PROVISIONING":            CloudVmClusterSummaryLifecycleStateProvisioning,
	"AVAILABLE":               CloudVmClusterSummaryLifecycleStateAvailable,
	"UPDATING":                CloudVmClusterSummaryLifecycleStateUpdating,
	"TERMINATING":             CloudVmClusterSummaryLifecycleStateTerminating,
	"TERMINATED":              CloudVmClusterSummaryLifecycleStateTerminated,
	"FAILED":                  CloudVmClusterSummaryLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": CloudVmClusterSummaryLifecycleStateMaintenanceInProgress,
}

var mappingCloudVmClusterSummaryLifecycleStateEnumLowerCase = map[string]CloudVmClusterSummaryLifecycleStateEnum{
	"provisioning":            CloudVmClusterSummaryLifecycleStateProvisioning,
	"available":               CloudVmClusterSummaryLifecycleStateAvailable,
	"updating":                CloudVmClusterSummaryLifecycleStateUpdating,
	"terminating":             CloudVmClusterSummaryLifecycleStateTerminating,
	"terminated":              CloudVmClusterSummaryLifecycleStateTerminated,
	"failed":                  CloudVmClusterSummaryLifecycleStateFailed,
	"maintenance_in_progress": CloudVmClusterSummaryLifecycleStateMaintenanceInProgress,
}

// GetCloudVmClusterSummaryLifecycleStateEnumValues Enumerates the set of values for CloudVmClusterSummaryLifecycleStateEnum
func GetCloudVmClusterSummaryLifecycleStateEnumValues() []CloudVmClusterSummaryLifecycleStateEnum {
	values := make([]CloudVmClusterSummaryLifecycleStateEnum, 0)
	for _, v := range mappingCloudVmClusterSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudVmClusterSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for CloudVmClusterSummaryLifecycleStateEnum
func GetCloudVmClusterSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"MAINTENANCE_IN_PROGRESS",
	}
}

// GetMappingCloudVmClusterSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudVmClusterSummaryLifecycleStateEnum(val string) (CloudVmClusterSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingCloudVmClusterSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CloudVmClusterSummaryLicenseModelEnum Enum with underlying type: string
type CloudVmClusterSummaryLicenseModelEnum string

// Set of constants representing the allowable values for CloudVmClusterSummaryLicenseModelEnum
const (
	CloudVmClusterSummaryLicenseModelLicenseIncluded     CloudVmClusterSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	CloudVmClusterSummaryLicenseModelBringYourOwnLicense CloudVmClusterSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCloudVmClusterSummaryLicenseModelEnum = map[string]CloudVmClusterSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       CloudVmClusterSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CloudVmClusterSummaryLicenseModelBringYourOwnLicense,
}

var mappingCloudVmClusterSummaryLicenseModelEnumLowerCase = map[string]CloudVmClusterSummaryLicenseModelEnum{
	"license_included":       CloudVmClusterSummaryLicenseModelLicenseIncluded,
	"bring_your_own_license": CloudVmClusterSummaryLicenseModelBringYourOwnLicense,
}

// GetCloudVmClusterSummaryLicenseModelEnumValues Enumerates the set of values for CloudVmClusterSummaryLicenseModelEnum
func GetCloudVmClusterSummaryLicenseModelEnumValues() []CloudVmClusterSummaryLicenseModelEnum {
	values := make([]CloudVmClusterSummaryLicenseModelEnum, 0)
	for _, v := range mappingCloudVmClusterSummaryLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudVmClusterSummaryLicenseModelEnumStringValues Enumerates the set of values in String for CloudVmClusterSummaryLicenseModelEnum
func GetCloudVmClusterSummaryLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCloudVmClusterSummaryLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudVmClusterSummaryLicenseModelEnum(val string) (CloudVmClusterSummaryLicenseModelEnum, bool) {
	enum, ok := mappingCloudVmClusterSummaryLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CloudVmClusterSummaryDiskRedundancyEnum Enum with underlying type: string
type CloudVmClusterSummaryDiskRedundancyEnum string

// Set of constants representing the allowable values for CloudVmClusterSummaryDiskRedundancyEnum
const (
	CloudVmClusterSummaryDiskRedundancyHigh   CloudVmClusterSummaryDiskRedundancyEnum = "HIGH"
	CloudVmClusterSummaryDiskRedundancyNormal CloudVmClusterSummaryDiskRedundancyEnum = "NORMAL"
)

var mappingCloudVmClusterSummaryDiskRedundancyEnum = map[string]CloudVmClusterSummaryDiskRedundancyEnum{
	"HIGH":   CloudVmClusterSummaryDiskRedundancyHigh,
	"NORMAL": CloudVmClusterSummaryDiskRedundancyNormal,
}

var mappingCloudVmClusterSummaryDiskRedundancyEnumLowerCase = map[string]CloudVmClusterSummaryDiskRedundancyEnum{
	"high":   CloudVmClusterSummaryDiskRedundancyHigh,
	"normal": CloudVmClusterSummaryDiskRedundancyNormal,
}

// GetCloudVmClusterSummaryDiskRedundancyEnumValues Enumerates the set of values for CloudVmClusterSummaryDiskRedundancyEnum
func GetCloudVmClusterSummaryDiskRedundancyEnumValues() []CloudVmClusterSummaryDiskRedundancyEnum {
	values := make([]CloudVmClusterSummaryDiskRedundancyEnum, 0)
	for _, v := range mappingCloudVmClusterSummaryDiskRedundancyEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudVmClusterSummaryDiskRedundancyEnumStringValues Enumerates the set of values in String for CloudVmClusterSummaryDiskRedundancyEnum
func GetCloudVmClusterSummaryDiskRedundancyEnumStringValues() []string {
	return []string{
		"HIGH",
		"NORMAL",
	}
}

// GetMappingCloudVmClusterSummaryDiskRedundancyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudVmClusterSummaryDiskRedundancyEnum(val string) (CloudVmClusterSummaryDiskRedundancyEnum, bool) {
	enum, ok := mappingCloudVmClusterSummaryDiskRedundancyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
