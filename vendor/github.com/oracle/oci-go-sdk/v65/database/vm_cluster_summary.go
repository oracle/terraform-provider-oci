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

// VmClusterSummary Details of the Exadata Cloud@Customer VM cluster.
type VmClusterSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation starts.
	LastPatchHistoryEntryId *string `mandatory:"false" json:"lastPatchHistoryEntryId"`

	// The current state of the VM cluster.
	LifecycleState VmClusterSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The user-friendly name for the Exadata Cloud@Customer VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time that the VM cluster was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time zone of the Exadata infrastructure. For details, see Exadata Infrastructure Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// If true, database backup on local Exadata storage is configured for the VM cluster. If false, database backup on local Exadata storage is not available in the VM cluster.
	IsLocalBackupEnabled *bool `mandatory:"false" json:"isLocalBackupEnabled"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	ExadataInfrastructureId *string `mandatory:"false" json:"exadataInfrastructureId"`

	// If true, sparse disk group is configured for the VM cluster. If false, sparse disk group is not created.
	IsSparseDiskgroupEnabled *bool `mandatory:"false" json:"isSparseDiskgroupEnabled"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster network.
	VmClusterNetworkId *string `mandatory:"false" json:"vmClusterNetworkId"`

	// The number of enabled CPU cores.
	CpusEnabled *int `mandatory:"false" json:"cpusEnabled"`

	// The number of enabled OCPU cores.
	OcpusEnabled *float32 `mandatory:"false" json:"ocpusEnabled"`

	// The memory allocated in GBs.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The local node storage allocated in GBs.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// Size, in terabytes, of the DATA disk group.
	DataStorageSizeInTBs *float64 `mandatory:"false" json:"dataStorageSizeInTBs"`

	// Size of the DATA disk group in GBs.
	DataStorageSizeInGBs *float64 `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The shape of the Exadata infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance.
	Shape *string `mandatory:"false" json:"shape"`

	// The Oracle Grid Infrastructure software version for the VM cluster.
	GiVersion *string `mandatory:"false" json:"giVersion"`

	// Operating system version of the image.
	SystemVersion *string `mandatory:"false" json:"systemVersion"`

	// The public key portion of one or more key pairs used for SSH access to the VM cluster.
	SshPublicKeys []string `mandatory:"false" json:"sshPublicKeys"`

	// The Oracle license model that applies to the VM cluster. The default is LICENSE_INCLUDED.
	LicenseModel VmClusterSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The list of Db server.
	DbServers []string `mandatory:"false" json:"dbServers"`

	// The name of the availability domain that the VM cluster is located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	DataCollectionOptions *DataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`
}

func (m VmClusterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmClusterSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVmClusterSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVmClusterSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVmClusterSummaryLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetVmClusterSummaryLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VmClusterSummaryLifecycleStateEnum Enum with underlying type: string
type VmClusterSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for VmClusterSummaryLifecycleStateEnum
const (
	VmClusterSummaryLifecycleStateProvisioning          VmClusterSummaryLifecycleStateEnum = "PROVISIONING"
	VmClusterSummaryLifecycleStateAvailable             VmClusterSummaryLifecycleStateEnum = "AVAILABLE"
	VmClusterSummaryLifecycleStateUpdating              VmClusterSummaryLifecycleStateEnum = "UPDATING"
	VmClusterSummaryLifecycleStateTerminating           VmClusterSummaryLifecycleStateEnum = "TERMINATING"
	VmClusterSummaryLifecycleStateTerminated            VmClusterSummaryLifecycleStateEnum = "TERMINATED"
	VmClusterSummaryLifecycleStateFailed                VmClusterSummaryLifecycleStateEnum = "FAILED"
	VmClusterSummaryLifecycleStateMaintenanceInProgress VmClusterSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingVmClusterSummaryLifecycleStateEnum = map[string]VmClusterSummaryLifecycleStateEnum{
	"PROVISIONING":            VmClusterSummaryLifecycleStateProvisioning,
	"AVAILABLE":               VmClusterSummaryLifecycleStateAvailable,
	"UPDATING":                VmClusterSummaryLifecycleStateUpdating,
	"TERMINATING":             VmClusterSummaryLifecycleStateTerminating,
	"TERMINATED":              VmClusterSummaryLifecycleStateTerminated,
	"FAILED":                  VmClusterSummaryLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": VmClusterSummaryLifecycleStateMaintenanceInProgress,
}

var mappingVmClusterSummaryLifecycleStateEnumLowerCase = map[string]VmClusterSummaryLifecycleStateEnum{
	"provisioning":            VmClusterSummaryLifecycleStateProvisioning,
	"available":               VmClusterSummaryLifecycleStateAvailable,
	"updating":                VmClusterSummaryLifecycleStateUpdating,
	"terminating":             VmClusterSummaryLifecycleStateTerminating,
	"terminated":              VmClusterSummaryLifecycleStateTerminated,
	"failed":                  VmClusterSummaryLifecycleStateFailed,
	"maintenance_in_progress": VmClusterSummaryLifecycleStateMaintenanceInProgress,
}

// GetVmClusterSummaryLifecycleStateEnumValues Enumerates the set of values for VmClusterSummaryLifecycleStateEnum
func GetVmClusterSummaryLifecycleStateEnumValues() []VmClusterSummaryLifecycleStateEnum {
	values := make([]VmClusterSummaryLifecycleStateEnum, 0)
	for _, v := range mappingVmClusterSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for VmClusterSummaryLifecycleStateEnum
func GetVmClusterSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingVmClusterSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterSummaryLifecycleStateEnum(val string) (VmClusterSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingVmClusterSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VmClusterSummaryLicenseModelEnum Enum with underlying type: string
type VmClusterSummaryLicenseModelEnum string

// Set of constants representing the allowable values for VmClusterSummaryLicenseModelEnum
const (
	VmClusterSummaryLicenseModelLicenseIncluded     VmClusterSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	VmClusterSummaryLicenseModelBringYourOwnLicense VmClusterSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingVmClusterSummaryLicenseModelEnum = map[string]VmClusterSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       VmClusterSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": VmClusterSummaryLicenseModelBringYourOwnLicense,
}

var mappingVmClusterSummaryLicenseModelEnumLowerCase = map[string]VmClusterSummaryLicenseModelEnum{
	"license_included":       VmClusterSummaryLicenseModelLicenseIncluded,
	"bring_your_own_license": VmClusterSummaryLicenseModelBringYourOwnLicense,
}

// GetVmClusterSummaryLicenseModelEnumValues Enumerates the set of values for VmClusterSummaryLicenseModelEnum
func GetVmClusterSummaryLicenseModelEnumValues() []VmClusterSummaryLicenseModelEnum {
	values := make([]VmClusterSummaryLicenseModelEnum, 0)
	for _, v := range mappingVmClusterSummaryLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterSummaryLicenseModelEnumStringValues Enumerates the set of values in String for VmClusterSummaryLicenseModelEnum
func GetVmClusterSummaryLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingVmClusterSummaryLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterSummaryLicenseModelEnum(val string) (VmClusterSummaryLicenseModelEnum, bool) {
	enum, ok := mappingVmClusterSummaryLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
