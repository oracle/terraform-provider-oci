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

// MaintenanceExecutionSummary Summary of a maintenance executions.
type MaintenanceExecutionSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance run execution.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the maintenance run execution.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The state of the maintenance run execution.
	LifecycleState MaintenanceRunLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Description of the maintenance run execution.
	Description *string `mandatory:"false" json:"description"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the maintenance run execution started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the maintenance run was completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// The type of the target resource on which the maintenance run execution occurred.
	TargetResourceType TargetResourceTypeEnumEnum `mandatory:"false" json:"targetResourceType,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure on which the maintenance run execution occurred.
	InfrastructureId *string `mandatory:"false" json:"infrastructureId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance run to which this maintenance execution belongs.
	MaintenanceRunId *string `mandatory:"false" json:"maintenanceRunId"`

	// Maintenance type.
	MaintenanceType MaintenanceTypeEnumEnum `mandatory:"false" json:"maintenanceType,omitempty"`

	// Maintenance run execution sub-type.
	MaintenanceSubtype MaintenanceSubtypeEnumEnum `mandatory:"false" json:"maintenanceSubtype,omitempty"`

	// The patching mode for the maintenance run that is being executed.
	PatchingMode PatchingModeEnumEnum `mandatory:"false" json:"patchingMode,omitempty"`

	// The source software version for the Oracle infrastructure.
	SourceVersion *string `mandatory:"false" json:"sourceVersion"`

	// The target software version for the Database Infrastructure patching operation.
	TargetVersion *string `mandatory:"false" json:"targetVersion"`

	// At the time of execution whether the custom action time out is enabled for the maintenance run that is being executed.
	IsCustomActionTimeoutEnabled *bool `mandatory:"false" json:"isCustomActionTimeoutEnabled"`

	// Determines the amount of time the system will wait before the start of each compute server patching operation.
	// Supported values are 15 to 120 minutes.
	CustomActionTimeoutInMins *int `mandatory:"false" json:"customActionTimeoutInMins"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the work request executed by this execution.
	WorkflowId *string `mandatory:"false" json:"workflowId"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The total time taken by this execution in minutes.
	TotalTimeTakenInMins *int `mandatory:"false" json:"totalTimeTakenInMins"`
}

func (m MaintenanceExecutionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaintenanceExecutionSummary) ValidateEnumValue() (bool, error) {
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
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
