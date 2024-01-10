// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrPlanExecution The details of a DR plan execution.
type DrPlanExecution struct {

	// The OCID of the DR plan execution.
	// Example: `ocid1.drplanexecution.oc1..uniqueID`
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing this DR plan execution.
	// Example: `ocid1.compartment.oc1..uniqueID`
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the DR plan execution.
	// Example: `Execution - EBS Switchover PHX to IAD`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the DR plan.
	// Example: `ocid1.drplan.oc1..uniqueID`
	PlanId *string `mandatory:"true" json:"planId"`

	// The type of the DR plan executed.
	PlanExecutionType DrPlanExecutionTypeEnum `mandatory:"true" json:"planExecutionType"`

	ExecutionOptions DrPlanExecutionOptions `mandatory:"true" json:"executionOptions"`

	// The OCID of the DR protection group to which this DR plan execution belongs.
	// Example: `ocid1.drprotectiongroup.oc1..uniqueID`
	DrProtectionGroupId *string `mandatory:"true" json:"drProtectionGroupId"`

	// The OCID of peer DR protection group associated with this plan's
	// DR protection group.
	// Example: `ocid1.drprotectiongroup.oc1..uniqueID`
	PeerDrProtectionGroupId *string `mandatory:"true" json:"peerDrProtectionGroupId"`

	// The region of the peer DR protection group associated with this plan's
	// DR protection group.
	// Example: `us-ashburn-1`
	PeerRegion *string `mandatory:"true" json:"peerRegion"`

	LogLocation *ObjectStorageLogLocation `mandatory:"true" json:"logLocation"`

	// The date and time at which DR plan execution was created. An RFC3339 formatted datetime string.
	// Example: `2019-03-29T09:36:42Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when DR plan execution was last updated. An RFC3339 formatted datetime string.
	// Example: `2019-03-29T09:36:42Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A list of groups executed in this DR plan execution.
	GroupExecutions []DrPlanGroupExecution `mandatory:"true" json:"groupExecutions"`

	// The current state of the DR plan execution.
	LifecycleState DrPlanExecutionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time at which DR plan execution began. An RFC3339 formatted datetime string.
	// Example: `2019-03-29T09:36:42Z`
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time at which DR plan execution succeeded, failed, was paused, or was canceled.
	// An RFC3339 formatted datetime string.
	// Example: `2019-03-29T09:36:42Z`
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// The total duration in seconds taken to complete the DR plan execution.
	// Example: `750`
	ExecutionDurationInSec *int `mandatory:"false" json:"executionDurationInSec"`

	// A message describing the DR plan execution's current state in more detail.
	LifeCycleDetails *string `mandatory:"false" json:"lifeCycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DrPlanExecution) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrPlanExecution) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrPlanExecutionTypeEnum(string(m.PlanExecutionType)); !ok && m.PlanExecutionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanExecutionType: %s. Supported values are: %s.", m.PlanExecutionType, strings.Join(GetDrPlanExecutionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDrPlanExecutionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDrPlanExecutionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DrPlanExecution) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeStarted             *common.SDKTime                   `json:"timeStarted"`
		TimeEnded               *common.SDKTime                   `json:"timeEnded"`
		ExecutionDurationInSec  *int                              `json:"executionDurationInSec"`
		LifeCycleDetails        *string                           `json:"lifeCycleDetails"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		SystemTags              map[string]map[string]interface{} `json:"systemTags"`
		Id                      *string                           `json:"id"`
		CompartmentId           *string                           `json:"compartmentId"`
		DisplayName             *string                           `json:"displayName"`
		PlanId                  *string                           `json:"planId"`
		PlanExecutionType       DrPlanExecutionTypeEnum           `json:"planExecutionType"`
		ExecutionOptions        drplanexecutionoptions            `json:"executionOptions"`
		DrProtectionGroupId     *string                           `json:"drProtectionGroupId"`
		PeerDrProtectionGroupId *string                           `json:"peerDrProtectionGroupId"`
		PeerRegion              *string                           `json:"peerRegion"`
		LogLocation             *ObjectStorageLogLocation         `json:"logLocation"`
		TimeCreated             *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated             *common.SDKTime                   `json:"timeUpdated"`
		GroupExecutions         []DrPlanGroupExecution            `json:"groupExecutions"`
		LifecycleState          DrPlanExecutionLifecycleStateEnum `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeStarted = model.TimeStarted

	m.TimeEnded = model.TimeEnded

	m.ExecutionDurationInSec = model.ExecutionDurationInSec

	m.LifeCycleDetails = model.LifeCycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.PlanId = model.PlanId

	m.PlanExecutionType = model.PlanExecutionType

	nn, e = model.ExecutionOptions.UnmarshalPolymorphicJSON(model.ExecutionOptions.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ExecutionOptions = nn.(DrPlanExecutionOptions)
	} else {
		m.ExecutionOptions = nil
	}

	m.DrProtectionGroupId = model.DrProtectionGroupId

	m.PeerDrProtectionGroupId = model.PeerDrProtectionGroupId

	m.PeerRegion = model.PeerRegion

	m.LogLocation = model.LogLocation

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.GroupExecutions = make([]DrPlanGroupExecution, len(model.GroupExecutions))
	copy(m.GroupExecutions, model.GroupExecutions)
	m.LifecycleState = model.LifecycleState

	return
}
