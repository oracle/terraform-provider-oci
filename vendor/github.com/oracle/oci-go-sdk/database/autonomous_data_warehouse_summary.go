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

// AutonomousDataWarehouseSummary An Oracle Autonomous Data Warehouse.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type AutonomousDataWarehouseSummary struct {

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
	LifecycleState AutonomousDataWarehouseSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	LicenseModel AutonomousDataWarehouseSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The URL of the Service Console for the Data Warehouse.
	ServiceConsoleUrl *string `mandatory:"false" json:"serviceConsoleUrl"`

	// The date and time the database was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m AutonomousDataWarehouseSummary) String() string {
	return common.PointerString(m)
}

// AutonomousDataWarehouseSummaryLicenseModelEnum Enum with underlying type: string
type AutonomousDataWarehouseSummaryLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousDataWarehouseSummaryLicenseModel
const (
	AutonomousDataWarehouseSummaryLicenseModelLicenseIncluded     AutonomousDataWarehouseSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousDataWarehouseSummaryLicenseModelBringYourOwnLicense AutonomousDataWarehouseSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousDataWarehouseSummaryLicenseModel = map[string]AutonomousDataWarehouseSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousDataWarehouseSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousDataWarehouseSummaryLicenseModelBringYourOwnLicense,
}

// GetAutonomousDataWarehouseSummaryLicenseModelEnumValues Enumerates the set of values for AutonomousDataWarehouseSummaryLicenseModel
func GetAutonomousDataWarehouseSummaryLicenseModelEnumValues() []AutonomousDataWarehouseSummaryLicenseModelEnum {
	values := make([]AutonomousDataWarehouseSummaryLicenseModelEnum, 0)
	for _, v := range mappingAutonomousDataWarehouseSummaryLicenseModel {
		values = append(values, v)
	}
	return values
}

// AutonomousDataWarehouseSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousDataWarehouseSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDataWarehouseSummaryLifecycleState
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
)

var mappingAutonomousDataWarehouseSummaryLifecycleState = map[string]AutonomousDataWarehouseSummaryLifecycleStateEnum{
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
}

// GetAutonomousDataWarehouseSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousDataWarehouseSummaryLifecycleState
func GetAutonomousDataWarehouseSummaryLifecycleStateEnumValues() []AutonomousDataWarehouseSummaryLifecycleStateEnum {
	values := make([]AutonomousDataWarehouseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDataWarehouseSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
