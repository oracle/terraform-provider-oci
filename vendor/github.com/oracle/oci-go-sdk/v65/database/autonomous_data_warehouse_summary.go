// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// AutonomousDataWarehouseSummary **Deprecated.** See AutonomousDatabase for reference information about Autonomous Databases with the warehouse workload type.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type AutonomousDataWarehouseSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the database.
	LifecycleState AutonomousDataWarehouseSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The database name.
	DbName *string `mandatory:"true" json:"dbName"`

	// The number of CPU cores to be made available to the database.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The quantity of data in the database, in terabytes.
	DataStorageSizeInTBs *int `mandatory:"true" json:"dataStorageSizeInTBs"`

	// Information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the database was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The user-friendly name for the Autonomous Data Warehouse. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The URL of the Service Console for the Data Warehouse.
	ServiceConsoleUrl *string `mandatory:"false" json:"serviceConsoleUrl"`

	// The connection string used to connect to the Data Warehouse. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Data Warehouse for the password value.
	ConnectionStrings *AutonomousDataWarehouseConnectionStrings `mandatory:"false" json:"connectionStrings"`

	// The Oracle license model that applies to the Oracle Autonomous Data Warehouse. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel AutonomousDataWarehouseSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A valid Oracle Database version for Autonomous Data Warehouse.
	DbVersion *string `mandatory:"false" json:"dbVersion"`
}

func (m AutonomousDataWarehouseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDataWarehouseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousDataWarehouseSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDataWarehouseSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousDataWarehouseSummaryLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetAutonomousDataWarehouseSummaryLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDataWarehouseSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousDataWarehouseSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDataWarehouseSummaryLifecycleStateEnum
const (
	AutonomousDataWarehouseSummaryLifecycleStateProvisioning            AutonomousDataWarehouseSummaryLifecycleStateEnum = "PROVISIONING"
	AutonomousDataWarehouseSummaryLifecycleStateAvailable               AutonomousDataWarehouseSummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousDataWarehouseSummaryLifecycleStateStopping                AutonomousDataWarehouseSummaryLifecycleStateEnum = "STOPPING"
	AutonomousDataWarehouseSummaryLifecycleStateStopped                 AutonomousDataWarehouseSummaryLifecycleStateEnum = "STOPPED"
	AutonomousDataWarehouseSummaryLifecycleStateStarting                AutonomousDataWarehouseSummaryLifecycleStateEnum = "STARTING"
	AutonomousDataWarehouseSummaryLifecycleStateTerminating             AutonomousDataWarehouseSummaryLifecycleStateEnum = "TERMINATING"
	AutonomousDataWarehouseSummaryLifecycleStateTerminated              AutonomousDataWarehouseSummaryLifecycleStateEnum = "TERMINATED"
	AutonomousDataWarehouseSummaryLifecycleStateUnavailable             AutonomousDataWarehouseSummaryLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDataWarehouseSummaryLifecycleStateRestoreInProgress       AutonomousDataWarehouseSummaryLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDataWarehouseSummaryLifecycleStateBackupInProgress        AutonomousDataWarehouseSummaryLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDataWarehouseSummaryLifecycleStateScaleInProgress         AutonomousDataWarehouseSummaryLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDataWarehouseSummaryLifecycleStateAvailableNeedsAttention AutonomousDataWarehouseSummaryLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDataWarehouseSummaryLifecycleStateUpdating                AutonomousDataWarehouseSummaryLifecycleStateEnum = "UPDATING"
)

var mappingAutonomousDataWarehouseSummaryLifecycleStateEnum = map[string]AutonomousDataWarehouseSummaryLifecycleStateEnum{
	"PROVISIONING":              AutonomousDataWarehouseSummaryLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDataWarehouseSummaryLifecycleStateAvailable,
	"STOPPING":                  AutonomousDataWarehouseSummaryLifecycleStateStopping,
	"STOPPED":                   AutonomousDataWarehouseSummaryLifecycleStateStopped,
	"STARTING":                  AutonomousDataWarehouseSummaryLifecycleStateStarting,
	"TERMINATING":               AutonomousDataWarehouseSummaryLifecycleStateTerminating,
	"TERMINATED":                AutonomousDataWarehouseSummaryLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDataWarehouseSummaryLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDataWarehouseSummaryLifecycleStateRestoreInProgress,
	"BACKUP_IN_PROGRESS":        AutonomousDataWarehouseSummaryLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDataWarehouseSummaryLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDataWarehouseSummaryLifecycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDataWarehouseSummaryLifecycleStateUpdating,
}

var mappingAutonomousDataWarehouseSummaryLifecycleStateEnumLowerCase = map[string]AutonomousDataWarehouseSummaryLifecycleStateEnum{
	"provisioning":              AutonomousDataWarehouseSummaryLifecycleStateProvisioning,
	"available":                 AutonomousDataWarehouseSummaryLifecycleStateAvailable,
	"stopping":                  AutonomousDataWarehouseSummaryLifecycleStateStopping,
	"stopped":                   AutonomousDataWarehouseSummaryLifecycleStateStopped,
	"starting":                  AutonomousDataWarehouseSummaryLifecycleStateStarting,
	"terminating":               AutonomousDataWarehouseSummaryLifecycleStateTerminating,
	"terminated":                AutonomousDataWarehouseSummaryLifecycleStateTerminated,
	"unavailable":               AutonomousDataWarehouseSummaryLifecycleStateUnavailable,
	"restore_in_progress":       AutonomousDataWarehouseSummaryLifecycleStateRestoreInProgress,
	"backup_in_progress":        AutonomousDataWarehouseSummaryLifecycleStateBackupInProgress,
	"scale_in_progress":         AutonomousDataWarehouseSummaryLifecycleStateScaleInProgress,
	"available_needs_attention": AutonomousDataWarehouseSummaryLifecycleStateAvailableNeedsAttention,
	"updating":                  AutonomousDataWarehouseSummaryLifecycleStateUpdating,
}

// GetAutonomousDataWarehouseSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousDataWarehouseSummaryLifecycleStateEnum
func GetAutonomousDataWarehouseSummaryLifecycleStateEnumValues() []AutonomousDataWarehouseSummaryLifecycleStateEnum {
	values := make([]AutonomousDataWarehouseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDataWarehouseSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDataWarehouseSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDataWarehouseSummaryLifecycleStateEnum
func GetAutonomousDataWarehouseSummaryLifecycleStateEnumStringValues() []string {
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
		"BACKUP_IN_PROGRESS",
		"SCALE_IN_PROGRESS",
		"AVAILABLE_NEEDS_ATTENTION",
		"UPDATING",
	}
}

// GetMappingAutonomousDataWarehouseSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDataWarehouseSummaryLifecycleStateEnum(val string) (AutonomousDataWarehouseSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDataWarehouseSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDataWarehouseSummaryLicenseModelEnum Enum with underlying type: string
type AutonomousDataWarehouseSummaryLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousDataWarehouseSummaryLicenseModelEnum
const (
	AutonomousDataWarehouseSummaryLicenseModelLicenseIncluded     AutonomousDataWarehouseSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousDataWarehouseSummaryLicenseModelBringYourOwnLicense AutonomousDataWarehouseSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousDataWarehouseSummaryLicenseModelEnum = map[string]AutonomousDataWarehouseSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousDataWarehouseSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousDataWarehouseSummaryLicenseModelBringYourOwnLicense,
}

var mappingAutonomousDataWarehouseSummaryLicenseModelEnumLowerCase = map[string]AutonomousDataWarehouseSummaryLicenseModelEnum{
	"license_included":       AutonomousDataWarehouseSummaryLicenseModelLicenseIncluded,
	"bring_your_own_license": AutonomousDataWarehouseSummaryLicenseModelBringYourOwnLicense,
}

// GetAutonomousDataWarehouseSummaryLicenseModelEnumValues Enumerates the set of values for AutonomousDataWarehouseSummaryLicenseModelEnum
func GetAutonomousDataWarehouseSummaryLicenseModelEnumValues() []AutonomousDataWarehouseSummaryLicenseModelEnum {
	values := make([]AutonomousDataWarehouseSummaryLicenseModelEnum, 0)
	for _, v := range mappingAutonomousDataWarehouseSummaryLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDataWarehouseSummaryLicenseModelEnumStringValues Enumerates the set of values in String for AutonomousDataWarehouseSummaryLicenseModelEnum
func GetAutonomousDataWarehouseSummaryLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingAutonomousDataWarehouseSummaryLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDataWarehouseSummaryLicenseModelEnum(val string) (AutonomousDataWarehouseSummaryLicenseModelEnum, bool) {
	enum, ok := mappingAutonomousDataWarehouseSummaryLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
