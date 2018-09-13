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

// AutonomousDatabase An Oracle Autonomous Database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type AutonomousDatabase struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of CPU cores to be made available to the database.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The quantity of data in the database, in terabytes.
	DataStorageSizeInTBs *int `mandatory:"true" json:"dataStorageSizeInTBs"`

	// The database name.
	DbName *string `mandatory:"true" json:"dbName"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the database.
	LifecycleState AutonomousDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The connection string used to connect to the Autonomous Database. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Database for the password value.
	ConnectionStrings *AutonomousDatabaseConnectionStrings `mandatory:"false" json:"connectionStrings"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The user-friendly name for the Autonomous Database. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The Oracle license model that applies to the Oracle Autonomous Database. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel AutonomousDatabaseLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The URL of the Service Console for the Autonomous Database.
	ServiceConsoleUrl *string `mandatory:"false" json:"serviceConsoleUrl"`

	// The date and time the database was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m AutonomousDatabase) String() string {
	return common.PointerString(m)
}

// AutonomousDatabaseLicenseModelEnum Enum with underlying type: string
type AutonomousDatabaseLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousDatabaseLicenseModel
const (
	AutonomousDatabaseLicenseModelLicenseIncluded     AutonomousDatabaseLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousDatabaseLicenseModelBringYourOwnLicense AutonomousDatabaseLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousDatabaseLicenseModel = map[string]AutonomousDatabaseLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousDatabaseLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousDatabaseLicenseModelBringYourOwnLicense,
}

// GetAutonomousDatabaseLicenseModelEnumValues Enumerates the set of values for AutonomousDatabaseLicenseModel
func GetAutonomousDatabaseLicenseModelEnumValues() []AutonomousDatabaseLicenseModelEnum {
	values := make([]AutonomousDatabaseLicenseModelEnum, 0)
	for _, v := range mappingAutonomousDatabaseLicenseModel {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseLifecycleState
const (
	AutonomousDatabaseLifecycleStateProvisioning            AutonomousDatabaseLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseLifecycleStateAvailable               AutonomousDatabaseLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseLifecycleStateStopping                AutonomousDatabaseLifecycleStateEnum = "STOPPING"
	AutonomousDatabaseLifecycleStateStopped                 AutonomousDatabaseLifecycleStateEnum = "STOPPED"
	AutonomousDatabaseLifecycleStateStarting                AutonomousDatabaseLifecycleStateEnum = "STARTING"
	AutonomousDatabaseLifecycleStateTerminating             AutonomousDatabaseLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseLifecycleStateTerminated              AutonomousDatabaseLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseLifecycleStateUnavailable             AutonomousDatabaseLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDatabaseLifecycleStateRestoreInProgress       AutonomousDatabaseLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDatabaseLifecycleStateBackupInProgress        AutonomousDatabaseLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDatabaseLifecycleStateScaleInProgress         AutonomousDatabaseLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDatabaseLifecycleStateAvailableNeedsAttention AutonomousDatabaseLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
)

var mappingAutonomousDatabaseLifecycleState = map[string]AutonomousDatabaseLifecycleStateEnum{
	"PROVISIONING":              AutonomousDatabaseLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDatabaseLifecycleStateAvailable,
	"STOPPING":                  AutonomousDatabaseLifecycleStateStopping,
	"STOPPED":                   AutonomousDatabaseLifecycleStateStopped,
	"STARTING":                  AutonomousDatabaseLifecycleStateStarting,
	"TERMINATING":               AutonomousDatabaseLifecycleStateTerminating,
	"TERMINATED":                AutonomousDatabaseLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDatabaseLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDatabaseLifecycleStateRestoreInProgress,
	"BACKUP_IN_PROGRESS":        AutonomousDatabaseLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDatabaseLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDatabaseLifecycleStateAvailableNeedsAttention,
}

// GetAutonomousDatabaseLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseLifecycleState
func GetAutonomousDatabaseLifecycleStateEnumValues() []AutonomousDatabaseLifecycleStateEnum {
	values := make([]AutonomousDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseLifecycleState {
		values = append(values, v)
	}
	return values
}
