// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousVmClusterSummary Details of the Autonomous VM cluster.
type AutonomousVmClusterSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous VM cluster.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Autonomous VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Autonomous VM cluster.
	LifecycleState AutonomousVmClusterSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	ExadataInfrastructureId *string `mandatory:"true" json:"exadataInfrastructureId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster network.
	VmClusterNetworkId *string `mandatory:"true" json:"vmClusterNetworkId"`

	// The date and time that the Autonomous VM cluster was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time zone to use for the Autonomous VM cluster. For details, see DB System Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// If true, database backup on local Exadata storage is configured for the Autonomous VM cluster. If false, database backup on local Exadata storage is not available in the Autonomous VM cluster.
	IsLocalBackupEnabled *bool `mandatory:"false" json:"isLocalBackupEnabled"`

	// The number of enabled CPU cores.
	CpusEnabled *int `mandatory:"false" json:"cpusEnabled"`

	// The numnber of CPU cores available.
	AvailableCpus *int `mandatory:"false" json:"availableCpus"`

	// The memory allocated in GBs.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The local node storage allocated in GBs.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// The total data storage allocated in TBs
	DataStorageSizeInTBs *float64 `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The data storage available in TBs
	AvailableDataStorageSizeInTBs *float64 `mandatory:"false" json:"availableDataStorageSizeInTBs"`

	// The Oracle license model that applies to the Autonomous VM cluster. The default is LICENSE_INCLUDED.
	LicenseModel AutonomousVmClusterSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AutonomousVmClusterSummary) String() string {
	return common.PointerString(m)
}

// AutonomousVmClusterSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousVmClusterSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousVmClusterSummaryLifecycleStateEnum
const (
	AutonomousVmClusterSummaryLifecycleStateProvisioning          AutonomousVmClusterSummaryLifecycleStateEnum = "PROVISIONING"
	AutonomousVmClusterSummaryLifecycleStateAvailable             AutonomousVmClusterSummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousVmClusterSummaryLifecycleStateUpdating              AutonomousVmClusterSummaryLifecycleStateEnum = "UPDATING"
	AutonomousVmClusterSummaryLifecycleStateTerminating           AutonomousVmClusterSummaryLifecycleStateEnum = "TERMINATING"
	AutonomousVmClusterSummaryLifecycleStateTerminated            AutonomousVmClusterSummaryLifecycleStateEnum = "TERMINATED"
	AutonomousVmClusterSummaryLifecycleStateFailed                AutonomousVmClusterSummaryLifecycleStateEnum = "FAILED"
	AutonomousVmClusterSummaryLifecycleStateMaintenanceInProgress AutonomousVmClusterSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingAutonomousVmClusterSummaryLifecycleState = map[string]AutonomousVmClusterSummaryLifecycleStateEnum{
	"PROVISIONING":            AutonomousVmClusterSummaryLifecycleStateProvisioning,
	"AVAILABLE":               AutonomousVmClusterSummaryLifecycleStateAvailable,
	"UPDATING":                AutonomousVmClusterSummaryLifecycleStateUpdating,
	"TERMINATING":             AutonomousVmClusterSummaryLifecycleStateTerminating,
	"TERMINATED":              AutonomousVmClusterSummaryLifecycleStateTerminated,
	"FAILED":                  AutonomousVmClusterSummaryLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": AutonomousVmClusterSummaryLifecycleStateMaintenanceInProgress,
}

// GetAutonomousVmClusterSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousVmClusterSummaryLifecycleStateEnum
func GetAutonomousVmClusterSummaryLifecycleStateEnumValues() []AutonomousVmClusterSummaryLifecycleStateEnum {
	values := make([]AutonomousVmClusterSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousVmClusterSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousVmClusterSummaryLicenseModelEnum Enum with underlying type: string
type AutonomousVmClusterSummaryLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousVmClusterSummaryLicenseModelEnum
const (
	AutonomousVmClusterSummaryLicenseModelLicenseIncluded     AutonomousVmClusterSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousVmClusterSummaryLicenseModelBringYourOwnLicense AutonomousVmClusterSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousVmClusterSummaryLicenseModel = map[string]AutonomousVmClusterSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousVmClusterSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousVmClusterSummaryLicenseModelBringYourOwnLicense,
}

// GetAutonomousVmClusterSummaryLicenseModelEnumValues Enumerates the set of values for AutonomousVmClusterSummaryLicenseModelEnum
func GetAutonomousVmClusterSummaryLicenseModelEnumValues() []AutonomousVmClusterSummaryLicenseModelEnum {
	values := make([]AutonomousVmClusterSummaryLicenseModelEnum, 0)
	for _, v := range mappingAutonomousVmClusterSummaryLicenseModel {
		values = append(values, v)
	}
	return values
}
