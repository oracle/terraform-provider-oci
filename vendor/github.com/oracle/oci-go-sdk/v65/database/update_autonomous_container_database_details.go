// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateAutonomousContainerDatabaseDetails Describes the modification parameters for the Autonomous Container Database.
type UpdateAutonomousContainerDatabaseDetails struct {

	// The display name for the Autonomous Container Database.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Customer Contacts. Setting this to an empty list removes all customer contacts.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	// The OKV End Point Group name for the Autonomous Container Database.
	OkvEndPointGroupName *string `mandatory:"false" json:"okvEndPointGroupName"`

	// Database Patch model preference.
	PatchModel UpdateAutonomousContainerDatabaseDetailsPatchModelEnum `mandatory:"false" json:"patchModel,omitempty"`

	MaintenanceWindowDetails *MaintenanceWindow `mandatory:"false" json:"maintenanceWindowDetails"`

	// The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database.
	// This value represents the number of days before schedlued maintenance of the primary database.
	StandbyMaintenanceBufferInDays *int `mandatory:"false" json:"standbyMaintenanceBufferInDays"`

	// The next maintenance version preference.
	VersionPreference UpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum `mandatory:"false" json:"versionPreference,omitempty"`

	// Indicates if an automatic DST Time Zone file update is enabled for the Autonomous Container Database. If enabled along with Release Update, patching will be done in a Non-Rolling manner.
	IsDstFileUpdateEnabled *bool `mandatory:"false" json:"isDstFileUpdateEnabled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	BackupConfig *AutonomousContainerDatabaseBackupConfig `mandatory:"false" json:"backupConfig"`

	// The CPU value beyond which an Autonomous AI Database will be opened across multiple nodes. The default value of this attribute is 16 for OCPUs and 64 for ECPUs.
	DbSplitThreshold *int `mandatory:"false" json:"dbSplitThreshold"`

	// The percentage of CPUs reserved across nodes to support node failover. Allowed values are 0%, 25%, 50%, 75%, and 100%, with 50% being the default option.
	VmFailoverReservation *int `mandatory:"false" json:"vmFailoverReservation"`

	// Determines whether an Autonomous AI Database must be opened across a minimum or maximum of nodes. By default, Minimum nodes is selected.
	DistributionAffinity UpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum `mandatory:"false" json:"distributionAffinity,omitempty"`

	// Enabling SHARED server architecture enables a database server to allow many client processes to share very few server processes, thereby increasing the number of supported users.
	NetServicesArchitecture UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum `mandatory:"false" json:"netServicesArchitecture,omitempty"`
}

func (m UpdateAutonomousContainerDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAutonomousContainerDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateAutonomousContainerDatabaseDetailsPatchModelEnum(string(m.PatchModel)); !ok && m.PatchModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchModel: %s. Supported values are: %s.", m.PatchModel, strings.Join(GetUpdateAutonomousContainerDatabaseDetailsPatchModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum(string(m.VersionPreference)); !ok && m.VersionPreference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VersionPreference: %s. Supported values are: %s.", m.VersionPreference, strings.Join(GetUpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum(string(m.DistributionAffinity)); !ok && m.DistributionAffinity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DistributionAffinity: %s. Supported values are: %s.", m.DistributionAffinity, strings.Join(GetUpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum(string(m.NetServicesArchitecture)); !ok && m.NetServicesArchitecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetServicesArchitecture: %s. Supported values are: %s.", m.NetServicesArchitecture, strings.Join(GetUpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateAutonomousContainerDatabaseDetailsPatchModelEnum Enum with underlying type: string
type UpdateAutonomousContainerDatabaseDetailsPatchModelEnum string

// Set of constants representing the allowable values for UpdateAutonomousContainerDatabaseDetailsPatchModelEnum
const (
	UpdateAutonomousContainerDatabaseDetailsPatchModelUpdates         UpdateAutonomousContainerDatabaseDetailsPatchModelEnum = "RELEASE_UPDATES"
	UpdateAutonomousContainerDatabaseDetailsPatchModelUpdateRevisions UpdateAutonomousContainerDatabaseDetailsPatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingUpdateAutonomousContainerDatabaseDetailsPatchModelEnum = map[string]UpdateAutonomousContainerDatabaseDetailsPatchModelEnum{
	"RELEASE_UPDATES":          UpdateAutonomousContainerDatabaseDetailsPatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": UpdateAutonomousContainerDatabaseDetailsPatchModelUpdateRevisions,
}

var mappingUpdateAutonomousContainerDatabaseDetailsPatchModelEnumLowerCase = map[string]UpdateAutonomousContainerDatabaseDetailsPatchModelEnum{
	"release_updates":          UpdateAutonomousContainerDatabaseDetailsPatchModelUpdates,
	"release_update_revisions": UpdateAutonomousContainerDatabaseDetailsPatchModelUpdateRevisions,
}

// GetUpdateAutonomousContainerDatabaseDetailsPatchModelEnumValues Enumerates the set of values for UpdateAutonomousContainerDatabaseDetailsPatchModelEnum
func GetUpdateAutonomousContainerDatabaseDetailsPatchModelEnumValues() []UpdateAutonomousContainerDatabaseDetailsPatchModelEnum {
	values := make([]UpdateAutonomousContainerDatabaseDetailsPatchModelEnum, 0)
	for _, v := range mappingUpdateAutonomousContainerDatabaseDetailsPatchModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousContainerDatabaseDetailsPatchModelEnumStringValues Enumerates the set of values in String for UpdateAutonomousContainerDatabaseDetailsPatchModelEnum
func GetUpdateAutonomousContainerDatabaseDetailsPatchModelEnumStringValues() []string {
	return []string{
		"RELEASE_UPDATES",
		"RELEASE_UPDATE_REVISIONS",
	}
}

// GetMappingUpdateAutonomousContainerDatabaseDetailsPatchModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAutonomousContainerDatabaseDetailsPatchModelEnum(val string) (UpdateAutonomousContainerDatabaseDetailsPatchModelEnum, bool) {
	enum, ok := mappingUpdateAutonomousContainerDatabaseDetailsPatchModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum Enum with underlying type: string
type UpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum string

// Set of constants representing the allowable values for UpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum
const (
	UpdateAutonomousContainerDatabaseDetailsVersionPreferenceNextReleaseUpdate   UpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum = "NEXT_RELEASE_UPDATE"
	UpdateAutonomousContainerDatabaseDetailsVersionPreferenceLatestReleaseUpdate UpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum = "LATEST_RELEASE_UPDATE"
)

var mappingUpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum = map[string]UpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum{
	"NEXT_RELEASE_UPDATE":   UpdateAutonomousContainerDatabaseDetailsVersionPreferenceNextReleaseUpdate,
	"LATEST_RELEASE_UPDATE": UpdateAutonomousContainerDatabaseDetailsVersionPreferenceLatestReleaseUpdate,
}

var mappingUpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnumLowerCase = map[string]UpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum{
	"next_release_update":   UpdateAutonomousContainerDatabaseDetailsVersionPreferenceNextReleaseUpdate,
	"latest_release_update": UpdateAutonomousContainerDatabaseDetailsVersionPreferenceLatestReleaseUpdate,
}

// GetUpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnumValues Enumerates the set of values for UpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum
func GetUpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnumValues() []UpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum {
	values := make([]UpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum, 0)
	for _, v := range mappingUpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnumStringValues Enumerates the set of values in String for UpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum
func GetUpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnumStringValues() []string {
	return []string{
		"NEXT_RELEASE_UPDATE",
		"LATEST_RELEASE_UPDATE",
	}
}

// GetMappingUpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum(val string) (UpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnum, bool) {
	enum, ok := mappingUpdateAutonomousContainerDatabaseDetailsVersionPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum Enum with underlying type: string
type UpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum string

// Set of constants representing the allowable values for UpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum
const (
	UpdateAutonomousContainerDatabaseDetailsDistributionAffinityMinimumDistribution UpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum = "MINIMUM_DISTRIBUTION"
	UpdateAutonomousContainerDatabaseDetailsDistributionAffinityMaximumDistribution UpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum = "MAXIMUM_DISTRIBUTION"
)

var mappingUpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum = map[string]UpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum{
	"MINIMUM_DISTRIBUTION": UpdateAutonomousContainerDatabaseDetailsDistributionAffinityMinimumDistribution,
	"MAXIMUM_DISTRIBUTION": UpdateAutonomousContainerDatabaseDetailsDistributionAffinityMaximumDistribution,
}

var mappingUpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnumLowerCase = map[string]UpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum{
	"minimum_distribution": UpdateAutonomousContainerDatabaseDetailsDistributionAffinityMinimumDistribution,
	"maximum_distribution": UpdateAutonomousContainerDatabaseDetailsDistributionAffinityMaximumDistribution,
}

// GetUpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnumValues Enumerates the set of values for UpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum
func GetUpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnumValues() []UpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum {
	values := make([]UpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum, 0)
	for _, v := range mappingUpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnumStringValues Enumerates the set of values in String for UpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum
func GetUpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnumStringValues() []string {
	return []string{
		"MINIMUM_DISTRIBUTION",
		"MAXIMUM_DISTRIBUTION",
	}
}

// GetMappingUpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum(val string) (UpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnum, bool) {
	enum, ok := mappingUpdateAutonomousContainerDatabaseDetailsDistributionAffinityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum Enum with underlying type: string
type UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum string

// Set of constants representing the allowable values for UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum
const (
	UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureDedicated UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum = "DEDICATED"
	UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureShared    UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum = "SHARED"
	UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureDrcp      UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum = "DRCP"
)

var mappingUpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum = map[string]UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum{
	"DEDICATED": UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureDedicated,
	"SHARED":    UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureShared,
	"DRCP":      UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureDrcp,
}

var mappingUpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnumLowerCase = map[string]UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum{
	"dedicated": UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureDedicated,
	"shared":    UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureShared,
	"drcp":      UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureDrcp,
}

// GetUpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnumValues Enumerates the set of values for UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum
func GetUpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnumValues() []UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum {
	values := make([]UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum, 0)
	for _, v := range mappingUpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnumStringValues Enumerates the set of values in String for UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum
func GetUpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnumStringValues() []string {
	return []string{
		"DEDICATED",
		"SHARED",
		"DRCP",
	}
}

// GetMappingUpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum(val string) (UpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum, bool) {
	enum, ok := mappingUpdateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
