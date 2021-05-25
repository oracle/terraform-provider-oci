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

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty.
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
}

func (m CloudVmClusterSummary) String() string {
	return common.PointerString(m)
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

var mappingCloudVmClusterSummaryLifecycleState = map[string]CloudVmClusterSummaryLifecycleStateEnum{
	"PROVISIONING":            CloudVmClusterSummaryLifecycleStateProvisioning,
	"AVAILABLE":               CloudVmClusterSummaryLifecycleStateAvailable,
	"UPDATING":                CloudVmClusterSummaryLifecycleStateUpdating,
	"TERMINATING":             CloudVmClusterSummaryLifecycleStateTerminating,
	"TERMINATED":              CloudVmClusterSummaryLifecycleStateTerminated,
	"FAILED":                  CloudVmClusterSummaryLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": CloudVmClusterSummaryLifecycleStateMaintenanceInProgress,
}

// GetCloudVmClusterSummaryLifecycleStateEnumValues Enumerates the set of values for CloudVmClusterSummaryLifecycleStateEnum
func GetCloudVmClusterSummaryLifecycleStateEnumValues() []CloudVmClusterSummaryLifecycleStateEnum {
	values := make([]CloudVmClusterSummaryLifecycleStateEnum, 0)
	for _, v := range mappingCloudVmClusterSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// CloudVmClusterSummaryLicenseModelEnum Enum with underlying type: string
type CloudVmClusterSummaryLicenseModelEnum string

// Set of constants representing the allowable values for CloudVmClusterSummaryLicenseModelEnum
const (
	CloudVmClusterSummaryLicenseModelLicenseIncluded     CloudVmClusterSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	CloudVmClusterSummaryLicenseModelBringYourOwnLicense CloudVmClusterSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCloudVmClusterSummaryLicenseModel = map[string]CloudVmClusterSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       CloudVmClusterSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CloudVmClusterSummaryLicenseModelBringYourOwnLicense,
}

// GetCloudVmClusterSummaryLicenseModelEnumValues Enumerates the set of values for CloudVmClusterSummaryLicenseModelEnum
func GetCloudVmClusterSummaryLicenseModelEnumValues() []CloudVmClusterSummaryLicenseModelEnum {
	values := make([]CloudVmClusterSummaryLicenseModelEnum, 0)
	for _, v := range mappingCloudVmClusterSummaryLicenseModel {
		values = append(values, v)
	}
	return values
}

// CloudVmClusterSummaryDiskRedundancyEnum Enum with underlying type: string
type CloudVmClusterSummaryDiskRedundancyEnum string

// Set of constants representing the allowable values for CloudVmClusterSummaryDiskRedundancyEnum
const (
	CloudVmClusterSummaryDiskRedundancyHigh   CloudVmClusterSummaryDiskRedundancyEnum = "HIGH"
	CloudVmClusterSummaryDiskRedundancyNormal CloudVmClusterSummaryDiskRedundancyEnum = "NORMAL"
)

var mappingCloudVmClusterSummaryDiskRedundancy = map[string]CloudVmClusterSummaryDiskRedundancyEnum{
	"HIGH":   CloudVmClusterSummaryDiskRedundancyHigh,
	"NORMAL": CloudVmClusterSummaryDiskRedundancyNormal,
}

// GetCloudVmClusterSummaryDiskRedundancyEnumValues Enumerates the set of values for CloudVmClusterSummaryDiskRedundancyEnum
func GetCloudVmClusterSummaryDiskRedundancyEnumValues() []CloudVmClusterSummaryDiskRedundancyEnum {
	values := make([]CloudVmClusterSummaryDiskRedundancyEnum, 0)
	for _, v := range mappingCloudVmClusterSummaryDiskRedundancy {
		values = append(values, v)
	}
	return values
}
