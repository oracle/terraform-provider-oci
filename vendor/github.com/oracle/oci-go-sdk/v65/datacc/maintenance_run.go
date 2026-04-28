// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MaintenanceRun Details of a maintenance run.
type MaintenanceRun struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance run.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the maintenance run.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the maintenance run.
	LifecycleState MaintenanceRunLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the maintenance run is scheduled to occur.
	TimeScheduled *common.SDKTime `mandatory:"true" json:"timeScheduled"`

	// Description of the maintenance run.
	Description *string `mandatory:"false" json:"description"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the maintenance run was originally scheduled.
	InitialScheduledTime *common.SDKTime `mandatory:"false" json:"initialScheduledTime"`

	// The date and time the maintenance run starts.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the maintenance run was completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// The type of the target resource on which the maintenance run occurs.
	TargetResourceType TargetResourceTypeEnumEnum `mandatory:"false" json:"targetResourceType,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure on which the maintenance run execution occurs.
	InfrastructureId *string `mandatory:"false" json:"infrastructureId"`

	// Maintenance type.
	MaintenanceType MaintenanceTypeEnumEnum `mandatory:"false" json:"maintenanceType,omitempty"`

	// The unique identifier of the patch. The identifier string includes the patch type and the version of the Database Infrastructure.
	PatchIdentifier *string `mandatory:"false" json:"patchIdentifier"`

	// Maintenance run sub-type.
	MaintenanceSubtype MaintenanceSubtypeEnumEnum `mandatory:"false" json:"maintenanceSubtype,omitempty"`

	// Database Infrastructure patching mode, either "ROLLING" or "NONROLLING". Default value is ROLLING.
	PatchingMode PatchingModeEnumEnum `mandatory:"false" json:"patchingMode,omitempty"`

	// Contain the patch failure count.
	PatchFailureCount *int `mandatory:"false" json:"patchFailureCount"`

	// The source software version for the Oracle infrastructure.
	SourceVersion *string `mandatory:"false" json:"sourceVersion"`

	// The target software version for the Database Infrastructure patching operation.
	TargetVersion *string `mandatory:"false" json:"targetVersion"`

	// If true, enables the configuration of a custom action timeout (waiting period) between compute servers patching operations.
	IsCustomActionTimeoutEnabled *bool `mandatory:"false" json:"isCustomActionTimeoutEnabled"`

	// Determines the amount of time the system will wait before the start of each compute server patching operation.
	// Supported values are 15 to 120 minutes.
	CustomActionTimeoutInMins *int `mandatory:"false" json:"customActionTimeoutInMins"`

	// The status of the patching operation.
	PatchingStatus PatchingStatusEnumEnum `mandatory:"false" json:"patchingStatus,omitempty"`

	// The time when the patching operation started.
	PatchingStartTime *common.SDKTime `mandatory:"false" json:"patchingStartTime"`

	// The time when the patching operation ended.
	PatchingEndTime *common.SDKTime `mandatory:"false" json:"patchingEndTime"`

	// The estimated total time required in minutes for all patching operations (compute servers, storage).
	EstimatedPatchingTime *int `mandatory:"false" json:"estimatedPatchingTime"`

	// The name of the current infrastructure component that is getting patched.
	CurrentPatchingComponent *string `mandatory:"false" json:"currentPatchingComponent"`

	// The estimated start time of the next infrastructure component patching operation.
	EstimatedComponentPatchingStartTime *common.SDKTime `mandatory:"false" json:"estimatedComponentPatchingStartTime"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The total time taken by corresponding resource activity in minutes.
	TotalTimeTakenInMins *int `mandatory:"false" json:"totalTimeTakenInMins"`
}

func (m MaintenanceRun) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaintenanceRun) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaintenanceRunLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMaintenanceRunLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingTargetResourceTypeEnumEnum(string(m.TargetResourceType)); !ok && m.TargetResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetResourceType: %s. Supported values are: %s.", m.TargetResourceType, strings.Join(GetTargetResourceTypeEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceTypeEnumEnum(string(m.MaintenanceType)); !ok && m.MaintenanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceType: %s. Supported values are: %s.", m.MaintenanceType, strings.Join(GetMaintenanceTypeEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceSubtypeEnumEnum(string(m.MaintenanceSubtype)); !ok && m.MaintenanceSubtype != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceSubtype: %s. Supported values are: %s.", m.MaintenanceSubtype, strings.Join(GetMaintenanceSubtypeEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchingModeEnumEnum(string(m.PatchingMode)); !ok && m.PatchingMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchingMode: %s. Supported values are: %s.", m.PatchingMode, strings.Join(GetPatchingModeEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchingStatusEnumEnum(string(m.PatchingStatus)); !ok && m.PatchingStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchingStatus: %s. Supported values are: %s.", m.PatchingStatus, strings.Join(GetPatchingStatusEnumEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
