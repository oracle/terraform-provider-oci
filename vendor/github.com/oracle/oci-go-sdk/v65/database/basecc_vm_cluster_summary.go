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

// BaseccVmClusterSummary Details of the Base Cloud@Customer VM cluster.
type BaseccVmClusterSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Basecc VM cluster on BICC Infrastructure.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the VM cluster.
	LifecycleState BaseccVmClusterSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	DatabaseEdition BaseccVmClusterSummaryDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

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

	// The Boot Additional Vm Storage Size in GB, to be allocated for the /u01 partition for the Basecc VM cluster on BICC Infrastructure.
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
	LicenseModel BaseccVmClusterSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Operating system version of the image.
	SystemVersion *string `mandatory:"false" json:"systemVersion"`

	// The vmcluster type for the Base Cloud@Customer VM cluster.
	VmClusterType BaseccVmClusterSummaryVmClusterTypeEnum `mandatory:"false" json:"vmClusterType,omitempty"`

	CloudAutomationUpdateDetails *CloudAutomationUpdateDetails `mandatory:"false" json:"cloudAutomationUpdateDetails"`
}

func (m BaseccVmClusterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BaseccVmClusterSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBaseccVmClusterSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBaseccVmClusterSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBaseccVmClusterSummaryDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetBaseccVmClusterSummaryDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseccVmClusterSummaryLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetBaseccVmClusterSummaryLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseccVmClusterSummaryVmClusterTypeEnum(string(m.VmClusterType)); !ok && m.VmClusterType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VmClusterType: %s. Supported values are: %s.", m.VmClusterType, strings.Join(GetBaseccVmClusterSummaryVmClusterTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseccVmClusterSummaryLifecycleStateEnum Enum with underlying type: string
type BaseccVmClusterSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for BaseccVmClusterSummaryLifecycleStateEnum
const (
	BaseccVmClusterSummaryLifecycleStateProvisioning          BaseccVmClusterSummaryLifecycleStateEnum = "PROVISIONING"
	BaseccVmClusterSummaryLifecycleStateAvailable             BaseccVmClusterSummaryLifecycleStateEnum = "AVAILABLE"
	BaseccVmClusterSummaryLifecycleStateUpdating              BaseccVmClusterSummaryLifecycleStateEnum = "UPDATING"
	BaseccVmClusterSummaryLifecycleStateTerminating           BaseccVmClusterSummaryLifecycleStateEnum = "TERMINATING"
	BaseccVmClusterSummaryLifecycleStateTerminated            BaseccVmClusterSummaryLifecycleStateEnum = "TERMINATED"
	BaseccVmClusterSummaryLifecycleStateFailed                BaseccVmClusterSummaryLifecycleStateEnum = "FAILED"
	BaseccVmClusterSummaryLifecycleStateMaintenanceInProgress BaseccVmClusterSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingBaseccVmClusterSummaryLifecycleStateEnum = map[string]BaseccVmClusterSummaryLifecycleStateEnum{
	"PROVISIONING":            BaseccVmClusterSummaryLifecycleStateProvisioning,
	"AVAILABLE":               BaseccVmClusterSummaryLifecycleStateAvailable,
	"UPDATING":                BaseccVmClusterSummaryLifecycleStateUpdating,
	"TERMINATING":             BaseccVmClusterSummaryLifecycleStateTerminating,
	"TERMINATED":              BaseccVmClusterSummaryLifecycleStateTerminated,
	"FAILED":                  BaseccVmClusterSummaryLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": BaseccVmClusterSummaryLifecycleStateMaintenanceInProgress,
}

var mappingBaseccVmClusterSummaryLifecycleStateEnumLowerCase = map[string]BaseccVmClusterSummaryLifecycleStateEnum{
	"provisioning":            BaseccVmClusterSummaryLifecycleStateProvisioning,
	"available":               BaseccVmClusterSummaryLifecycleStateAvailable,
	"updating":                BaseccVmClusterSummaryLifecycleStateUpdating,
	"terminating":             BaseccVmClusterSummaryLifecycleStateTerminating,
	"terminated":              BaseccVmClusterSummaryLifecycleStateTerminated,
	"failed":                  BaseccVmClusterSummaryLifecycleStateFailed,
	"maintenance_in_progress": BaseccVmClusterSummaryLifecycleStateMaintenanceInProgress,
}

// GetBaseccVmClusterSummaryLifecycleStateEnumValues Enumerates the set of values for BaseccVmClusterSummaryLifecycleStateEnum
func GetBaseccVmClusterSummaryLifecycleStateEnumValues() []BaseccVmClusterSummaryLifecycleStateEnum {
	values := make([]BaseccVmClusterSummaryLifecycleStateEnum, 0)
	for _, v := range mappingBaseccVmClusterSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for BaseccVmClusterSummaryLifecycleStateEnum
func GetBaseccVmClusterSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingBaseccVmClusterSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterSummaryLifecycleStateEnum(val string) (BaseccVmClusterSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingBaseccVmClusterSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterSummaryDatabaseEditionEnum Enum with underlying type: string
type BaseccVmClusterSummaryDatabaseEditionEnum string

// Set of constants representing the allowable values for BaseccVmClusterSummaryDatabaseEditionEnum
const (
	BaseccVmClusterSummaryDatabaseEditionStandardEdition                     BaseccVmClusterSummaryDatabaseEditionEnum = "STANDARD_EDITION"
	BaseccVmClusterSummaryDatabaseEditionEnterpriseEdition                   BaseccVmClusterSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION"
	BaseccVmClusterSummaryDatabaseEditionEnterpriseEditionHighPerformance    BaseccVmClusterSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	BaseccVmClusterSummaryDatabaseEditionEnterpriseEditionExtremePerformance BaseccVmClusterSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
	BaseccVmClusterSummaryDatabaseEditionEnterpriseEditionDeveloper          BaseccVmClusterSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION_DEVELOPER"
)

var mappingBaseccVmClusterSummaryDatabaseEditionEnum = map[string]BaseccVmClusterSummaryDatabaseEditionEnum{
	"STANDARD_EDITION":                       BaseccVmClusterSummaryDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     BaseccVmClusterSummaryDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    BaseccVmClusterSummaryDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": BaseccVmClusterSummaryDatabaseEditionEnterpriseEditionExtremePerformance,
	"ENTERPRISE_EDITION_DEVELOPER":           BaseccVmClusterSummaryDatabaseEditionEnterpriseEditionDeveloper,
}

var mappingBaseccVmClusterSummaryDatabaseEditionEnumLowerCase = map[string]BaseccVmClusterSummaryDatabaseEditionEnum{
	"standard_edition":                       BaseccVmClusterSummaryDatabaseEditionStandardEdition,
	"enterprise_edition":                     BaseccVmClusterSummaryDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    BaseccVmClusterSummaryDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": BaseccVmClusterSummaryDatabaseEditionEnterpriseEditionExtremePerformance,
	"enterprise_edition_developer":           BaseccVmClusterSummaryDatabaseEditionEnterpriseEditionDeveloper,
}

// GetBaseccVmClusterSummaryDatabaseEditionEnumValues Enumerates the set of values for BaseccVmClusterSummaryDatabaseEditionEnum
func GetBaseccVmClusterSummaryDatabaseEditionEnumValues() []BaseccVmClusterSummaryDatabaseEditionEnum {
	values := make([]BaseccVmClusterSummaryDatabaseEditionEnum, 0)
	for _, v := range mappingBaseccVmClusterSummaryDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterSummaryDatabaseEditionEnumStringValues Enumerates the set of values in String for BaseccVmClusterSummaryDatabaseEditionEnum
func GetBaseccVmClusterSummaryDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
		"ENTERPRISE_EDITION_DEVELOPER",
	}
}

// GetMappingBaseccVmClusterSummaryDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterSummaryDatabaseEditionEnum(val string) (BaseccVmClusterSummaryDatabaseEditionEnum, bool) {
	enum, ok := mappingBaseccVmClusterSummaryDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterSummaryLicenseModelEnum Enum with underlying type: string
type BaseccVmClusterSummaryLicenseModelEnum string

// Set of constants representing the allowable values for BaseccVmClusterSummaryLicenseModelEnum
const (
	BaseccVmClusterSummaryLicenseModelLicenseIncluded     BaseccVmClusterSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	BaseccVmClusterSummaryLicenseModelBringYourOwnLicense BaseccVmClusterSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingBaseccVmClusterSummaryLicenseModelEnum = map[string]BaseccVmClusterSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       BaseccVmClusterSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": BaseccVmClusterSummaryLicenseModelBringYourOwnLicense,
}

var mappingBaseccVmClusterSummaryLicenseModelEnumLowerCase = map[string]BaseccVmClusterSummaryLicenseModelEnum{
	"license_included":       BaseccVmClusterSummaryLicenseModelLicenseIncluded,
	"bring_your_own_license": BaseccVmClusterSummaryLicenseModelBringYourOwnLicense,
}

// GetBaseccVmClusterSummaryLicenseModelEnumValues Enumerates the set of values for BaseccVmClusterSummaryLicenseModelEnum
func GetBaseccVmClusterSummaryLicenseModelEnumValues() []BaseccVmClusterSummaryLicenseModelEnum {
	values := make([]BaseccVmClusterSummaryLicenseModelEnum, 0)
	for _, v := range mappingBaseccVmClusterSummaryLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterSummaryLicenseModelEnumStringValues Enumerates the set of values in String for BaseccVmClusterSummaryLicenseModelEnum
func GetBaseccVmClusterSummaryLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingBaseccVmClusterSummaryLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterSummaryLicenseModelEnum(val string) (BaseccVmClusterSummaryLicenseModelEnum, bool) {
	enum, ok := mappingBaseccVmClusterSummaryLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterSummaryVmClusterTypeEnum Enum with underlying type: string
type BaseccVmClusterSummaryVmClusterTypeEnum string

// Set of constants representing the allowable values for BaseccVmClusterSummaryVmClusterTypeEnum
const (
	BaseccVmClusterSummaryVmClusterTypeRegular   BaseccVmClusterSummaryVmClusterTypeEnum = "REGULAR"
	BaseccVmClusterSummaryVmClusterTypeDeveloper BaseccVmClusterSummaryVmClusterTypeEnum = "DEVELOPER"
)

var mappingBaseccVmClusterSummaryVmClusterTypeEnum = map[string]BaseccVmClusterSummaryVmClusterTypeEnum{
	"REGULAR":   BaseccVmClusterSummaryVmClusterTypeRegular,
	"DEVELOPER": BaseccVmClusterSummaryVmClusterTypeDeveloper,
}

var mappingBaseccVmClusterSummaryVmClusterTypeEnumLowerCase = map[string]BaseccVmClusterSummaryVmClusterTypeEnum{
	"regular":   BaseccVmClusterSummaryVmClusterTypeRegular,
	"developer": BaseccVmClusterSummaryVmClusterTypeDeveloper,
}

// GetBaseccVmClusterSummaryVmClusterTypeEnumValues Enumerates the set of values for BaseccVmClusterSummaryVmClusterTypeEnum
func GetBaseccVmClusterSummaryVmClusterTypeEnumValues() []BaseccVmClusterSummaryVmClusterTypeEnum {
	values := make([]BaseccVmClusterSummaryVmClusterTypeEnum, 0)
	for _, v := range mappingBaseccVmClusterSummaryVmClusterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterSummaryVmClusterTypeEnumStringValues Enumerates the set of values in String for BaseccVmClusterSummaryVmClusterTypeEnum
func GetBaseccVmClusterSummaryVmClusterTypeEnumStringValues() []string {
	return []string{
		"REGULAR",
		"DEVELOPER",
	}
}

// GetMappingBaseccVmClusterSummaryVmClusterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterSummaryVmClusterTypeEnum(val string) (BaseccVmClusterSummaryVmClusterTypeEnum, bool) {
	enum, ok := mappingBaseccVmClusterSummaryVmClusterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
