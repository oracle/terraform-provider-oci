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

// ExadbVmClusterSummary Details of the Exadata VM cluster on Exascale Infrastructure. Applies to Exadata Database Service on Exascale Infrastructure only.
type ExadbVmClusterSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata VM cluster on Exascale Infrastructure.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the availability domain in which the Exadata VM cluster on Exascale Infrastructure is located.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure.
	BackupSubnetId *string `mandatory:"true" json:"backupSubnetId"`

	// The current state of the Exadata VM cluster on Exascale Infrastructure.
	LifecycleState ExadbVmClusterSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The number of nodes in the Exadata VM cluster on Exascale Infrastructure.
	NodeCount *int `mandatory:"true" json:"nodeCount"`

	// The shape of the Exadata VM cluster on Exascale Infrastructure resource
	Shape *string `mandatory:"true" json:"shape"`

	// The user-friendly name for the Exadata VM cluster on Exascale Infrastructure. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The hostname for the Exadata VM cluster on Exascale Infrastructure. The hostname must begin with an alphabetic character, and
	// can contain alphanumeric characters and hyphens (-). For Exadata systems, the maximum length of the hostname is 12 characters.
	// The maximum length of the combined hostname and domain is 63 characters.
	// **Note:** The hostname must be unique within the subnet. If it is not unique,
	// then the Exadata VM cluster on Exascale Infrastructure will fail to provision.
	Hostname *string `mandatory:"true" json:"hostname"`

	// A domain name used for the Exadata VM cluster on Exascale Infrastructure. If the Oracle-provided internet and VCN
	// resolver is enabled for the specified subnet, then the domain name for the subnet is used
	// (do not provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted.
	// Applies to Exadata Database Service on Exascale Infrastructure only.
	Domain *string `mandatory:"true" json:"domain"`

	// The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys"`

	// The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure.
	TotalECpuCount *int `mandatory:"true" json:"totalECpuCount"`

	// The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure.
	EnabledECpuCount *int `mandatory:"true" json:"enabledECpuCount"`

	VmFileSystemStorage *ExadbVmClusterStorageDetails `mandatory:"true" json:"vmFileSystemStorage"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
	ExascaleDbStorageVaultId *string `mandatory:"true" json:"exascaleDbStorageVaultId"`

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems.
	BackupNetworkNsgIds []string `mandatory:"false" json:"backupNetworkNsgIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last maintenance update history entry. This value is updated when a maintenance update starts.
	LastUpdateHistoryEntryId *string `mandatory:"false" json:"lastUpdateHistoryEntryId"`

	// The port number configured for the listener on the Exadata VM cluster on Exascale Infrastructure.
	ListenerPort *int64 `mandatory:"false" json:"listenerPort"`

	// The date and time that the Exadata VM cluster on Exascale Infrastructure was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time zone to use for the Exadata VM cluster on Exascale Infrastructure. For details, see Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The cluster name for Exadata VM cluster on Exascale Infrastructure. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive.
	ClusterName *string `mandatory:"false" json:"clusterName"`

	// A valid Oracle Grid Infrastructure (GI) software version.
	GiVersion *string `mandatory:"false" json:"giVersion"`

	// Grid Setup will be done using this grid image id.
	// The grid image id can be extracted from
	// 1. Obtain the supported major versions using API /20160918/giVersions?compartmentId=<compartmentId>&shape=EXADB_XS&availabilityDomain=<AD name>
	// 2. Replace {version} with one of the supported major versions and obtain the supported minor versions using
	// API /20160918/giVersions/{version}/minorVersions?compartmentId=<compartmentId>&shapeFamily=EXADB_XS&availabilityDomain=<AD name>
	GridImageId *string `mandatory:"false" json:"gridImageId"`

	// The type of Grid Image
	GridImageType ExadbVmClusterSummaryGridImageTypeEnum `mandatory:"false" json:"gridImageType,omitempty"`

	// Operating system version of the image.
	SystemVersion *string `mandatory:"false" json:"systemVersion"`

	// The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel ExadbVmClusterSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Single Client Access Name (SCAN) IP addresses associated with the Exadata VM cluster on Exascale Infrastructure.
	// SCAN IP addresses are typically used for load balancing and are not assigned to any interface.
	// Oracle Clusterware directs the requests to the appropriate nodes in the cluster.
	// **Note:** For a single-node DB system, this list is empty.
	ScanIpIds []string `mandatory:"false" json:"scanIpIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the virtual IP (VIP) addresses associated with the Exadata VM cluster on Exascale Infrastructure.
	// The Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the Exadata Cloud Service instance to
	// enable failover. If one node fails, then the VIP is reassigned to another active node in the cluster.
	VipIds []string `mandatory:"false" json:"vipIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DNS record for the SCAN IP addresses that are associated with the Exadata VM cluster on Exascale Infrastructure.
	ScanDnsRecordId *string `mandatory:"false" json:"scanDnsRecordId"`

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

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The FQDN of the DNS record for the SCAN IP addresses that are associated with the Exadata VM cluster on Exascale Infrastructure.
	ScanDnsName *string `mandatory:"false" json:"scanDnsName"`

	// The OCID of the zone with which the Exadata VM cluster on Exascale Infrastructure is associated.
	ZoneId *string `mandatory:"false" json:"zoneId"`

	// The TCP Single Client Access Name (SCAN) port. The default port is 1521.
	ScanListenerPortTcp *int `mandatory:"false" json:"scanListenerPortTcp"`

	// The Secured Communication (TCPS) protocol Single Client Access Name (SCAN) port. The default port is 2484.
	ScanListenerPortTcpSsl *int `mandatory:"false" json:"scanListenerPortTcpSsl"`

	// The private zone ID in which you want DNS records to be created.
	PrivateZoneId *string `mandatory:"false" json:"privateZoneId"`

	DataCollectionOptions *DataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`

	SnapshotFileSystemStorage *ExadbVmClusterStorageDetails `mandatory:"false" json:"snapshotFileSystemStorage"`

	TotalFileSystemStorage *ExadbVmClusterStorageDetails `mandatory:"false" json:"totalFileSystemStorage"`

	// The memory that you want to be allocated in GBs. Memory is calculated based on 11 GB per VM core reserved.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`
}

func (m ExadbVmClusterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadbVmClusterSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExadbVmClusterSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExadbVmClusterSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExadbVmClusterSummaryGridImageTypeEnum(string(m.GridImageType)); !ok && m.GridImageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GridImageType: %s. Supported values are: %s.", m.GridImageType, strings.Join(GetExadbVmClusterSummaryGridImageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadbVmClusterSummaryLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetExadbVmClusterSummaryLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExadbVmClusterSummaryLifecycleStateEnum Enum with underlying type: string
type ExadbVmClusterSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ExadbVmClusterSummaryLifecycleStateEnum
const (
	ExadbVmClusterSummaryLifecycleStateProvisioning          ExadbVmClusterSummaryLifecycleStateEnum = "PROVISIONING"
	ExadbVmClusterSummaryLifecycleStateAvailable             ExadbVmClusterSummaryLifecycleStateEnum = "AVAILABLE"
	ExadbVmClusterSummaryLifecycleStateUpdating              ExadbVmClusterSummaryLifecycleStateEnum = "UPDATING"
	ExadbVmClusterSummaryLifecycleStateTerminating           ExadbVmClusterSummaryLifecycleStateEnum = "TERMINATING"
	ExadbVmClusterSummaryLifecycleStateTerminated            ExadbVmClusterSummaryLifecycleStateEnum = "TERMINATED"
	ExadbVmClusterSummaryLifecycleStateFailed                ExadbVmClusterSummaryLifecycleStateEnum = "FAILED"
	ExadbVmClusterSummaryLifecycleStateMaintenanceInProgress ExadbVmClusterSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingExadbVmClusterSummaryLifecycleStateEnum = map[string]ExadbVmClusterSummaryLifecycleStateEnum{
	"PROVISIONING":            ExadbVmClusterSummaryLifecycleStateProvisioning,
	"AVAILABLE":               ExadbVmClusterSummaryLifecycleStateAvailable,
	"UPDATING":                ExadbVmClusterSummaryLifecycleStateUpdating,
	"TERMINATING":             ExadbVmClusterSummaryLifecycleStateTerminating,
	"TERMINATED":              ExadbVmClusterSummaryLifecycleStateTerminated,
	"FAILED":                  ExadbVmClusterSummaryLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": ExadbVmClusterSummaryLifecycleStateMaintenanceInProgress,
}

var mappingExadbVmClusterSummaryLifecycleStateEnumLowerCase = map[string]ExadbVmClusterSummaryLifecycleStateEnum{
	"provisioning":            ExadbVmClusterSummaryLifecycleStateProvisioning,
	"available":               ExadbVmClusterSummaryLifecycleStateAvailable,
	"updating":                ExadbVmClusterSummaryLifecycleStateUpdating,
	"terminating":             ExadbVmClusterSummaryLifecycleStateTerminating,
	"terminated":              ExadbVmClusterSummaryLifecycleStateTerminated,
	"failed":                  ExadbVmClusterSummaryLifecycleStateFailed,
	"maintenance_in_progress": ExadbVmClusterSummaryLifecycleStateMaintenanceInProgress,
}

// GetExadbVmClusterSummaryLifecycleStateEnumValues Enumerates the set of values for ExadbVmClusterSummaryLifecycleStateEnum
func GetExadbVmClusterSummaryLifecycleStateEnumValues() []ExadbVmClusterSummaryLifecycleStateEnum {
	values := make([]ExadbVmClusterSummaryLifecycleStateEnum, 0)
	for _, v := range mappingExadbVmClusterSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ExadbVmClusterSummaryLifecycleStateEnum
func GetExadbVmClusterSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingExadbVmClusterSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterSummaryLifecycleStateEnum(val string) (ExadbVmClusterSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingExadbVmClusterSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadbVmClusterSummaryGridImageTypeEnum Enum with underlying type: string
type ExadbVmClusterSummaryGridImageTypeEnum string

// Set of constants representing the allowable values for ExadbVmClusterSummaryGridImageTypeEnum
const (
	ExadbVmClusterSummaryGridImageTypeReleaseUpdate ExadbVmClusterSummaryGridImageTypeEnum = "RELEASE_UPDATE"
	ExadbVmClusterSummaryGridImageTypeCustomImage   ExadbVmClusterSummaryGridImageTypeEnum = "CUSTOM_IMAGE"
)

var mappingExadbVmClusterSummaryGridImageTypeEnum = map[string]ExadbVmClusterSummaryGridImageTypeEnum{
	"RELEASE_UPDATE": ExadbVmClusterSummaryGridImageTypeReleaseUpdate,
	"CUSTOM_IMAGE":   ExadbVmClusterSummaryGridImageTypeCustomImage,
}

var mappingExadbVmClusterSummaryGridImageTypeEnumLowerCase = map[string]ExadbVmClusterSummaryGridImageTypeEnum{
	"release_update": ExadbVmClusterSummaryGridImageTypeReleaseUpdate,
	"custom_image":   ExadbVmClusterSummaryGridImageTypeCustomImage,
}

// GetExadbVmClusterSummaryGridImageTypeEnumValues Enumerates the set of values for ExadbVmClusterSummaryGridImageTypeEnum
func GetExadbVmClusterSummaryGridImageTypeEnumValues() []ExadbVmClusterSummaryGridImageTypeEnum {
	values := make([]ExadbVmClusterSummaryGridImageTypeEnum, 0)
	for _, v := range mappingExadbVmClusterSummaryGridImageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterSummaryGridImageTypeEnumStringValues Enumerates the set of values in String for ExadbVmClusterSummaryGridImageTypeEnum
func GetExadbVmClusterSummaryGridImageTypeEnumStringValues() []string {
	return []string{
		"RELEASE_UPDATE",
		"CUSTOM_IMAGE",
	}
}

// GetMappingExadbVmClusterSummaryGridImageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterSummaryGridImageTypeEnum(val string) (ExadbVmClusterSummaryGridImageTypeEnum, bool) {
	enum, ok := mappingExadbVmClusterSummaryGridImageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadbVmClusterSummaryLicenseModelEnum Enum with underlying type: string
type ExadbVmClusterSummaryLicenseModelEnum string

// Set of constants representing the allowable values for ExadbVmClusterSummaryLicenseModelEnum
const (
	ExadbVmClusterSummaryLicenseModelLicenseIncluded     ExadbVmClusterSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	ExadbVmClusterSummaryLicenseModelBringYourOwnLicense ExadbVmClusterSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingExadbVmClusterSummaryLicenseModelEnum = map[string]ExadbVmClusterSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       ExadbVmClusterSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": ExadbVmClusterSummaryLicenseModelBringYourOwnLicense,
}

var mappingExadbVmClusterSummaryLicenseModelEnumLowerCase = map[string]ExadbVmClusterSummaryLicenseModelEnum{
	"license_included":       ExadbVmClusterSummaryLicenseModelLicenseIncluded,
	"bring_your_own_license": ExadbVmClusterSummaryLicenseModelBringYourOwnLicense,
}

// GetExadbVmClusterSummaryLicenseModelEnumValues Enumerates the set of values for ExadbVmClusterSummaryLicenseModelEnum
func GetExadbVmClusterSummaryLicenseModelEnumValues() []ExadbVmClusterSummaryLicenseModelEnum {
	values := make([]ExadbVmClusterSummaryLicenseModelEnum, 0)
	for _, v := range mappingExadbVmClusterSummaryLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterSummaryLicenseModelEnumStringValues Enumerates the set of values in String for ExadbVmClusterSummaryLicenseModelEnum
func GetExadbVmClusterSummaryLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingExadbVmClusterSummaryLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterSummaryLicenseModelEnum(val string) (ExadbVmClusterSummaryLicenseModelEnum, bool) {
	enum, ok := mappingExadbVmClusterSummaryLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
