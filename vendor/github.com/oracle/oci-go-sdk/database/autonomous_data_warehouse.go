// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousDataWarehouse An Oracle Autonomous Data Warehouse.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type AutonomousDataWarehouse struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of CPU cores to be made available to the database.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The quantity of data in the database, in terabytes.
	DataStorageSizeInTBs *int `mandatory:"true" json:"dataStorageSizeInTBs"`

	// The database name.
	DbName *string `mandatory:"true" json:"dbName"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the database.
	LifecycleState AutonomousDataWarehouseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The connection string used to connect to the Data Warehouse. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Data Warehouse for the password value.
	ConnectionStrings *AutonomousDataWarehouseConnectionStrings `mandatory:"false" json:"connectionStrings"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The user-friendly name for the Autonomous Data Warehouse. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The Oracle license model that applies to the Oracle Autonomous Data Warehouse. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel AutonomousDataWarehouseLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The URL of the Service Console for the Data Warehouse.
	ServiceConsoleUrl *string `mandatory:"false" json:"serviceConsoleUrl"`

	// The date and time the database was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m AutonomousDataWarehouse) String() string {
	return common.PointerString(m)
}

// AutonomousDataWarehouseLicenseModelEnum Enum with underlying type: string
type AutonomousDataWarehouseLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousDataWarehouseLicenseModel
const (
	AutonomousDataWarehouseLicenseModelLicenseIncluded     AutonomousDataWarehouseLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousDataWarehouseLicenseModelBringYourOwnLicense AutonomousDataWarehouseLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousDataWarehouseLicenseModel = map[string]AutonomousDataWarehouseLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousDataWarehouseLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousDataWarehouseLicenseModelBringYourOwnLicense,
}

// GetAutonomousDataWarehouseLicenseModelEnumValues Enumerates the set of values for AutonomousDataWarehouseLicenseModel
func GetAutonomousDataWarehouseLicenseModelEnumValues() []AutonomousDataWarehouseLicenseModelEnum {
	values := make([]AutonomousDataWarehouseLicenseModelEnum, 0)
	for _, v := range mappingAutonomousDataWarehouseLicenseModel {
		values = append(values, v)
	}
	return values
}

// AutonomousDataWarehouseLifecycleStateEnum Enum with underlying type: string
type AutonomousDataWarehouseLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDataWarehouseLifecycleState
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
)

var mappingAutonomousDataWarehouseLifecycleState = map[string]AutonomousDataWarehouseLifecycleStateEnum{
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
}

// GetAutonomousDataWarehouseLifecycleStateEnumValues Enumerates the set of values for AutonomousDataWarehouseLifecycleState
func GetAutonomousDataWarehouseLifecycleStateEnumValues() []AutonomousDataWarehouseLifecycleStateEnum {
	values := make([]AutonomousDataWarehouseLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDataWarehouseLifecycleState {
		values = append(values, v)
	}
	return values
}
