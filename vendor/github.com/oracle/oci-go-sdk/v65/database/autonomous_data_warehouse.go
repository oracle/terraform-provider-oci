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

// AutonomousDataWarehouse **Deprecated.** See AutonomousDatabase for reference information about Autonomous Databases with the warehouse workload type.
type AutonomousDataWarehouse struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the database.
	LifecycleState AutonomousDataWarehouseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	LicenseModel AutonomousDataWarehouseLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A valid Oracle Database version for Autonomous Data Warehouse.
	DbVersion *string `mandatory:"false" json:"dbVersion"`
}

func (m AutonomousDataWarehouse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDataWarehouse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousDataWarehouseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDataWarehouseLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousDataWarehouseLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetAutonomousDataWarehouseLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDataWarehouseLifecycleStateEnum Enum with underlying type: string
type AutonomousDataWarehouseLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDataWarehouseLifecycleStateEnum
const (
	AutonomousDataWarehouseLifecycleStateProvisioning            AutonomousDataWarehouseLifecycleStateEnum = "PROVISIONING"
	AutonomousDataWarehouseLifecycleStateAvailable               AutonomousDataWarehouseLifecycleStateEnum = "AVAILABLE"
	AutonomousDataWarehouseLifecycleStateStopping                AutonomousDataWarehouseLifecycleStateEnum = "STOPPING"
	AutonomousDataWarehouseLifecycleStateStopped                 AutonomousDataWarehouseLifecycleStateEnum = "STOPPED"
	AutonomousDataWarehouseLifecycleStateStarting                AutonomousDataWarehouseLifecycleStateEnum = "STARTING"
	AutonomousDataWarehouseLifecycleStateTerminating             AutonomousDataWarehouseLifecycleStateEnum = "TERMINATING"
	AutonomousDataWarehouseLifecycleStateTerminated              AutonomousDataWarehouseLifecycleStateEnum = "TERMINATED"
	AutonomousDataWarehouseLifecycleStateUnavailable             AutonomousDataWarehouseLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDataWarehouseLifecycleStateRestoreInProgress       AutonomousDataWarehouseLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDataWarehouseLifecycleStateBackupInProgress        AutonomousDataWarehouseLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDataWarehouseLifecycleStateScaleInProgress         AutonomousDataWarehouseLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDataWarehouseLifecycleStateAvailableNeedsAttention AutonomousDataWarehouseLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDataWarehouseLifecycleStateUpdating                AutonomousDataWarehouseLifecycleStateEnum = "UPDATING"
)

var mappingAutonomousDataWarehouseLifecycleStateEnum = map[string]AutonomousDataWarehouseLifecycleStateEnum{
	"PROVISIONING":              AutonomousDataWarehouseLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDataWarehouseLifecycleStateAvailable,
	"STOPPING":                  AutonomousDataWarehouseLifecycleStateStopping,
	"STOPPED":                   AutonomousDataWarehouseLifecycleStateStopped,
	"STARTING":                  AutonomousDataWarehouseLifecycleStateStarting,
	"TERMINATING":               AutonomousDataWarehouseLifecycleStateTerminating,
	"TERMINATED":                AutonomousDataWarehouseLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDataWarehouseLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDataWarehouseLifecycleStateRestoreInProgress,
	"BACKUP_IN_PROGRESS":        AutonomousDataWarehouseLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDataWarehouseLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDataWarehouseLifecycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDataWarehouseLifecycleStateUpdating,
}

var mappingAutonomousDataWarehouseLifecycleStateEnumLowerCase = map[string]AutonomousDataWarehouseLifecycleStateEnum{
	"provisioning":              AutonomousDataWarehouseLifecycleStateProvisioning,
	"available":                 AutonomousDataWarehouseLifecycleStateAvailable,
	"stopping":                  AutonomousDataWarehouseLifecycleStateStopping,
	"stopped":                   AutonomousDataWarehouseLifecycleStateStopped,
	"starting":                  AutonomousDataWarehouseLifecycleStateStarting,
	"terminating":               AutonomousDataWarehouseLifecycleStateTerminating,
	"terminated":                AutonomousDataWarehouseLifecycleStateTerminated,
	"unavailable":               AutonomousDataWarehouseLifecycleStateUnavailable,
	"restore_in_progress":       AutonomousDataWarehouseLifecycleStateRestoreInProgress,
	"backup_in_progress":        AutonomousDataWarehouseLifecycleStateBackupInProgress,
	"scale_in_progress":         AutonomousDataWarehouseLifecycleStateScaleInProgress,
	"available_needs_attention": AutonomousDataWarehouseLifecycleStateAvailableNeedsAttention,
	"updating":                  AutonomousDataWarehouseLifecycleStateUpdating,
}

// GetAutonomousDataWarehouseLifecycleStateEnumValues Enumerates the set of values for AutonomousDataWarehouseLifecycleStateEnum
func GetAutonomousDataWarehouseLifecycleStateEnumValues() []AutonomousDataWarehouseLifecycleStateEnum {
	values := make([]AutonomousDataWarehouseLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDataWarehouseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDataWarehouseLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDataWarehouseLifecycleStateEnum
func GetAutonomousDataWarehouseLifecycleStateEnumStringValues() []string {
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

// GetMappingAutonomousDataWarehouseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDataWarehouseLifecycleStateEnum(val string) (AutonomousDataWarehouseLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDataWarehouseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDataWarehouseLicenseModelEnum Enum with underlying type: string
type AutonomousDataWarehouseLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousDataWarehouseLicenseModelEnum
const (
	AutonomousDataWarehouseLicenseModelLicenseIncluded     AutonomousDataWarehouseLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousDataWarehouseLicenseModelBringYourOwnLicense AutonomousDataWarehouseLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousDataWarehouseLicenseModelEnum = map[string]AutonomousDataWarehouseLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousDataWarehouseLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousDataWarehouseLicenseModelBringYourOwnLicense,
}

var mappingAutonomousDataWarehouseLicenseModelEnumLowerCase = map[string]AutonomousDataWarehouseLicenseModelEnum{
	"license_included":       AutonomousDataWarehouseLicenseModelLicenseIncluded,
	"bring_your_own_license": AutonomousDataWarehouseLicenseModelBringYourOwnLicense,
}

// GetAutonomousDataWarehouseLicenseModelEnumValues Enumerates the set of values for AutonomousDataWarehouseLicenseModelEnum
func GetAutonomousDataWarehouseLicenseModelEnumValues() []AutonomousDataWarehouseLicenseModelEnum {
	values := make([]AutonomousDataWarehouseLicenseModelEnum, 0)
	for _, v := range mappingAutonomousDataWarehouseLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDataWarehouseLicenseModelEnumStringValues Enumerates the set of values in String for AutonomousDataWarehouseLicenseModelEnum
func GetAutonomousDataWarehouseLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingAutonomousDataWarehouseLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDataWarehouseLicenseModelEnum(val string) (AutonomousDataWarehouseLicenseModelEnum, bool) {
	enum, ok := mappingAutonomousDataWarehouseLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
