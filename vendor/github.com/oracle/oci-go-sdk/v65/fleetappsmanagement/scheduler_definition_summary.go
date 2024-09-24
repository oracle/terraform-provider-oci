// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SchedulerDefinitionSummary Summary of the SchedulerDefinition.
type SchedulerDefinitionSummary struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Tenancy OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Associated region
	ResourceRegion *string `mandatory:"false" json:"resourceRegion"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Scheduled date for the next run of the Job.
	TimeOfNextRun *common.SDKTime `mandatory:"false" json:"timeOfNextRun"`

	Schedule *Schedule `mandatory:"false" json:"schedule"`

	// Count of Action Groups affected by the Schedule.
	CountOfAffectedActionGroups *int `mandatory:"false" json:"countOfAffectedActionGroups"`

	// Count of Resources affected by the Schedule
	CountOfAffectedResources *int `mandatory:"false" json:"countOfAffectedResources"`

	// Count of Targets affected by the Schedule
	CountOfAffectedTargets *int `mandatory:"false" json:"countOfAffectedTargets"`

	// All ActionGroup Types part of the schedule.
	ActionGroupTypes []LifeCycleActionGroupTypeEnum `mandatory:"false" json:"actionGroupTypes,omitempty"`

	// All application types part of the schedule for ENVIRONMENT ActionGroup Type.
	ApplicationTypes []string `mandatory:"false" json:"applicationTypes"`

	// All products part of the schedule for PRODUCT ActionGroup Type.
	Products []string `mandatory:"false" json:"products"`

	// All LifeCycle Operations part of the schedule
	LifecycleOperations []string `mandatory:"false" json:"lifecycleOperations"`

	// The current state of the SchedulerDefinition.
	LifecycleState SchedulerDefinitionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SchedulerDefinitionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SchedulerDefinitionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.ActionGroupTypes {
		if _, ok := GetMappingLifeCycleActionGroupTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionGroupTypes: %s. Supported values are: %s.", val, strings.Join(GetLifeCycleActionGroupTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSchedulerDefinitionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSchedulerDefinitionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
