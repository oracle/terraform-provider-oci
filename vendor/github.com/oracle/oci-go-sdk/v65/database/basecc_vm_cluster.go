// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BaseccVmCluster Details of the VM cluster resource. Applies to Base Cloud@Customer instances only.
type BaseccVmCluster struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Basecc VM cluster on BICC Infrastructure.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the VM cluster.
	LifecycleState BaseccVmClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time that the Basecc VM cluster on BICC Infrastructure was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The name of the availability domain in which the Basecc VM cluster on BICC Infrastructure is located.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of CPU cores to enable for the Basecc VM cluster.
	CpusEnabled *int `mandatory:"true" json:"cpusEnabled"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Oracle Base Infrastructure.
	BaseInfrastructureId *string `mandatory:"true" json:"baseInfrastructureId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Base Virtual Machine Network Id.
	BaseVmClusterNetworkId *string `mandatory:"true" json:"baseVmClusterNetworkId"`

	// A valid Oracle Grid Infrastructure (GI) software version.
	GiVersion *string `mandatory:"true" json:"giVersion"`

	// The user-friendly name for the Basecc VM cluster on BICC Infrastructure. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The public key portion of one or more key pairs used for SSH access to the Basecc VM cluster on BICC Infrastructure.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation starts.
	LastPatchHistoryEntryId *string `mandatory:"false" json:"lastPatchHistoryEntryId"`

	// The date and time that the Basecc VM cluster was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The Oracle Database Edition that applies to all the databases on the DB system.
	// Exadata DB systems and 2-node RAC DB systems require ENTERPRISE_EDITION_EXTREME_PERFORMANCE.
	DatabaseEdition BaseccVmClusterDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

	// The number of nodes in the Basecc VM cluster on BICC Infrastructure.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The list of base server.
	DbServers []string `mandatory:"false" json:"dbServers"`

	DataCollectionOptions *DataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`

	// The description for Basecc VM Cluster.
	Description *string `mandatory:"false" json:"description"`

	// The time zone to use for the Basecc VM cluster on BICC Infrastructure. For details, see Time Zones (https://docs.oracle.com/iaas/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The cluster name for Basecc VM cluster on BICC Infrastructure. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive.
	ClusterName *string `mandatory:"false" json:"clusterName"`

	// The Boot disk group size to be allocated in GBs for the Basecc VM cluster on BICC Infrastructure.
	BootStorageSizeInGBs *int `mandatory:"false" json:"bootStorageSizeInGBs"`

	// The Additional Vm Storage Size in GB, to be allocated for the /u01 partition for the Basecc VM cluster on BICC Infrastructure.
	AdditionalVmStorageSizeInGBs *int `mandatory:"false" json:"additionalVmStorageSizeInGBs"`

	// The total storage allocated in GBs.
	TotalStorageSizeInGBs *int `mandatory:"false" json:"totalStorageSizeInGBs"`

	// The Data Disk Group size in GB for the Basecc VM cluster on BICC Infrastructure.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The Reco Disk Group size in GB for the Basecc VM cluster on BICC Infrastructure.
	RecoStorageSizeInGBs *int `mandatory:"false" json:"recoStorageSizeInGBs"`

	// The memory to be allocated per VM in GBs. The default is 8GB per core.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The Oracle license model that applies to the Basecc VM cluster on BICC Infrastructure. The default is LICENSE_INCLUDED.
	LicenseModel BaseccVmClusterLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Operating system version of the image.
	SystemVersion *string `mandatory:"false" json:"systemVersion"`

	// The vmcluster type for the Base Cloud@Customer VM cluster.
	VmClusterType BaseccVmClusterVmClusterTypeEnum `mandatory:"false" json:"vmClusterType,omitempty"`

	CloudAutomationUpdateDetails *CloudAutomationUpdateDetails `mandatory:"false" json:"cloudAutomationUpdateDetails"`
}

func (m BaseccVmCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BaseccVmCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBaseccVmClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBaseccVmClusterLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBaseccVmClusterDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetBaseccVmClusterDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseccVmClusterLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetBaseccVmClusterLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseccVmClusterVmClusterTypeEnum(string(m.VmClusterType)); !ok && m.VmClusterType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VmClusterType: %s. Supported values are: %s.", m.VmClusterType, strings.Join(GetBaseccVmClusterVmClusterTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseccVmClusterLifecycleStateEnum Enum with underlying type: string
type BaseccVmClusterLifecycleStateEnum string

// Set of constants representing the allowable values for BaseccVmClusterLifecycleStateEnum
const (
	BaseccVmClusterLifecycleStateProvisioning          BaseccVmClusterLifecycleStateEnum = "PROVISIONING"
	BaseccVmClusterLifecycleStateAvailable             BaseccVmClusterLifecycleStateEnum = "AVAILABLE"
	BaseccVmClusterLifecycleStateUpdating              BaseccVmClusterLifecycleStateEnum = "UPDATING"
	BaseccVmClusterLifecycleStateTerminating           BaseccVmClusterLifecycleStateEnum = "TERMINATING"
	BaseccVmClusterLifecycleStateTerminated            BaseccVmClusterLifecycleStateEnum = "TERMINATED"
	BaseccVmClusterLifecycleStateFailed                BaseccVmClusterLifecycleStateEnum = "FAILED"
	BaseccVmClusterLifecycleStateMaintenanceInProgress BaseccVmClusterLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingBaseccVmClusterLifecycleStateEnum = map[string]BaseccVmClusterLifecycleStateEnum{
	"PROVISIONING":            BaseccVmClusterLifecycleStateProvisioning,
	"AVAILABLE":               BaseccVmClusterLifecycleStateAvailable,
	"UPDATING":                BaseccVmClusterLifecycleStateUpdating,
	"TERMINATING":             BaseccVmClusterLifecycleStateTerminating,
	"TERMINATED":              BaseccVmClusterLifecycleStateTerminated,
	"FAILED":                  BaseccVmClusterLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": BaseccVmClusterLifecycleStateMaintenanceInProgress,
}

var mappingBaseccVmClusterLifecycleStateEnumLowerCase = map[string]BaseccVmClusterLifecycleStateEnum{
	"provisioning":            BaseccVmClusterLifecycleStateProvisioning,
	"available":               BaseccVmClusterLifecycleStateAvailable,
	"updating":                BaseccVmClusterLifecycleStateUpdating,
	"terminating":             BaseccVmClusterLifecycleStateTerminating,
	"terminated":              BaseccVmClusterLifecycleStateTerminated,
	"failed":                  BaseccVmClusterLifecycleStateFailed,
	"maintenance_in_progress": BaseccVmClusterLifecycleStateMaintenanceInProgress,
}

// GetBaseccVmClusterLifecycleStateEnumValues Enumerates the set of values for BaseccVmClusterLifecycleStateEnum
func GetBaseccVmClusterLifecycleStateEnumValues() []BaseccVmClusterLifecycleStateEnum {
	values := make([]BaseccVmClusterLifecycleStateEnum, 0)
	for _, v := range mappingBaseccVmClusterLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterLifecycleStateEnumStringValues Enumerates the set of values in String for BaseccVmClusterLifecycleStateEnum
func GetBaseccVmClusterLifecycleStateEnumStringValues() []string {
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

// GetMappingBaseccVmClusterLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterLifecycleStateEnum(val string) (BaseccVmClusterLifecycleStateEnum, bool) {
	enum, ok := mappingBaseccVmClusterLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterDatabaseEditionEnum Enum with underlying type: string
type BaseccVmClusterDatabaseEditionEnum string

// Set of constants representing the allowable values for BaseccVmClusterDatabaseEditionEnum
const (
	BaseccVmClusterDatabaseEditionStandardEdition                     BaseccVmClusterDatabaseEditionEnum = "STANDARD_EDITION"
	BaseccVmClusterDatabaseEditionEnterpriseEdition                   BaseccVmClusterDatabaseEditionEnum = "ENTERPRISE_EDITION"
	BaseccVmClusterDatabaseEditionEnterpriseEditionHighPerformance    BaseccVmClusterDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	BaseccVmClusterDatabaseEditionEnterpriseEditionExtremePerformance BaseccVmClusterDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
	BaseccVmClusterDatabaseEditionEnterpriseEditionDeveloper          BaseccVmClusterDatabaseEditionEnum = "ENTERPRISE_EDITION_DEVELOPER"
)

var mappingBaseccVmClusterDatabaseEditionEnum = map[string]BaseccVmClusterDatabaseEditionEnum{
	"STANDARD_EDITION":                       BaseccVmClusterDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     BaseccVmClusterDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    BaseccVmClusterDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": BaseccVmClusterDatabaseEditionEnterpriseEditionExtremePerformance,
	"ENTERPRISE_EDITION_DEVELOPER":           BaseccVmClusterDatabaseEditionEnterpriseEditionDeveloper,
}

var mappingBaseccVmClusterDatabaseEditionEnumLowerCase = map[string]BaseccVmClusterDatabaseEditionEnum{
	"standard_edition":                       BaseccVmClusterDatabaseEditionStandardEdition,
	"enterprise_edition":                     BaseccVmClusterDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    BaseccVmClusterDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": BaseccVmClusterDatabaseEditionEnterpriseEditionExtremePerformance,
	"enterprise_edition_developer":           BaseccVmClusterDatabaseEditionEnterpriseEditionDeveloper,
}

// GetBaseccVmClusterDatabaseEditionEnumValues Enumerates the set of values for BaseccVmClusterDatabaseEditionEnum
func GetBaseccVmClusterDatabaseEditionEnumValues() []BaseccVmClusterDatabaseEditionEnum {
	values := make([]BaseccVmClusterDatabaseEditionEnum, 0)
	for _, v := range mappingBaseccVmClusterDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterDatabaseEditionEnumStringValues Enumerates the set of values in String for BaseccVmClusterDatabaseEditionEnum
func GetBaseccVmClusterDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
		"ENTERPRISE_EDITION_DEVELOPER",
	}
}

// GetMappingBaseccVmClusterDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterDatabaseEditionEnum(val string) (BaseccVmClusterDatabaseEditionEnum, bool) {
	enum, ok := mappingBaseccVmClusterDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterLicenseModelEnum Enum with underlying type: string
type BaseccVmClusterLicenseModelEnum string

// Set of constants representing the allowable values for BaseccVmClusterLicenseModelEnum
const (
	BaseccVmClusterLicenseModelLicenseIncluded     BaseccVmClusterLicenseModelEnum = "LICENSE_INCLUDED"
	BaseccVmClusterLicenseModelBringYourOwnLicense BaseccVmClusterLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingBaseccVmClusterLicenseModelEnum = map[string]BaseccVmClusterLicenseModelEnum{
	"LICENSE_INCLUDED":       BaseccVmClusterLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": BaseccVmClusterLicenseModelBringYourOwnLicense,
}

var mappingBaseccVmClusterLicenseModelEnumLowerCase = map[string]BaseccVmClusterLicenseModelEnum{
	"license_included":       BaseccVmClusterLicenseModelLicenseIncluded,
	"bring_your_own_license": BaseccVmClusterLicenseModelBringYourOwnLicense,
}

// GetBaseccVmClusterLicenseModelEnumValues Enumerates the set of values for BaseccVmClusterLicenseModelEnum
func GetBaseccVmClusterLicenseModelEnumValues() []BaseccVmClusterLicenseModelEnum {
	values := make([]BaseccVmClusterLicenseModelEnum, 0)
	for _, v := range mappingBaseccVmClusterLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterLicenseModelEnumStringValues Enumerates the set of values in String for BaseccVmClusterLicenseModelEnum
func GetBaseccVmClusterLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingBaseccVmClusterLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterLicenseModelEnum(val string) (BaseccVmClusterLicenseModelEnum, bool) {
	enum, ok := mappingBaseccVmClusterLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterVmClusterTypeEnum Enum with underlying type: string
type BaseccVmClusterVmClusterTypeEnum string

// Set of constants representing the allowable values for BaseccVmClusterVmClusterTypeEnum
const (
	BaseccVmClusterVmClusterTypeRegular   BaseccVmClusterVmClusterTypeEnum = "REGULAR"
	BaseccVmClusterVmClusterTypeDeveloper BaseccVmClusterVmClusterTypeEnum = "DEVELOPER"
)

var mappingBaseccVmClusterVmClusterTypeEnum = map[string]BaseccVmClusterVmClusterTypeEnum{
	"REGULAR":   BaseccVmClusterVmClusterTypeRegular,
	"DEVELOPER": BaseccVmClusterVmClusterTypeDeveloper,
}

var mappingBaseccVmClusterVmClusterTypeEnumLowerCase = map[string]BaseccVmClusterVmClusterTypeEnum{
	"regular":   BaseccVmClusterVmClusterTypeRegular,
	"developer": BaseccVmClusterVmClusterTypeDeveloper,
}

// GetBaseccVmClusterVmClusterTypeEnumValues Enumerates the set of values for BaseccVmClusterVmClusterTypeEnum
func GetBaseccVmClusterVmClusterTypeEnumValues() []BaseccVmClusterVmClusterTypeEnum {
	values := make([]BaseccVmClusterVmClusterTypeEnum, 0)
	for _, v := range mappingBaseccVmClusterVmClusterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterVmClusterTypeEnumStringValues Enumerates the set of values in String for BaseccVmClusterVmClusterTypeEnum
func GetBaseccVmClusterVmClusterTypeEnumStringValues() []string {
	return []string{
		"REGULAR",
		"DEVELOPER",
	}
}

// GetMappingBaseccVmClusterVmClusterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterVmClusterTypeEnum(val string) (BaseccVmClusterVmClusterTypeEnum, bool) {
	enum, ok := mappingBaseccVmClusterVmClusterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
