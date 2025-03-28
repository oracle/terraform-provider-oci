// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrPlanExecutionSummary The summary of a DR plan execution.
type DrPlanExecutionSummary struct {

	// The OCID of the DR plan execution.
	// Example: `ocid1.drplanexecution.oc1..uniqueID`
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing this plan execution.
	// Example: `ocid1.compartment.oc1..uniqueID`
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the DR plan execution.
	// Example: `Execution - EBS Switchover PHX to IAD`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the DR plan for this DR plan execution.
	// Example: `ocid1.drplan.oc1..uniqueID`
	PlanId *string `mandatory:"true" json:"planId"`

	// The type of the DR plan execution.
	PlanExecutionType DrPlanExecutionTypeEnum `mandatory:"true" json:"planExecutionType"`

	// The OCID of the DR protection group to which this DR plan execution belongs.
	// Example: `ocid1.drprotectiongroup.oc1..uniqueID`
	DrProtectionGroupId *string `mandatory:"true" json:"drProtectionGroupId"`

	// The OCID of peer DR protection group associated with this DR plan execution's
	// DR protection group.
	// Example: `ocid1.drprotectiongroup.oc1..uniqueID`
	PeerDrProtectionGroupId *string `mandatory:"true" json:"peerDrProtectionGroupId"`

	// The region of the peer DR protection group associated with this DR plan execution's
	// DR protection group.
	// Example: `us-ashburn-1`
	PeerRegion *string `mandatory:"true" json:"peerRegion"`

	LogLocation *ObjectStorageLogLocation `mandatory:"true" json:"logLocation"`

	// The date and time at which DR plan execution was created. An RFC3339 formatted datetime string.
	// Example: `2019-03-29T09:36:42Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when this DR plan execution was last updated.
	// Example: `2019-03-29T09:36:42Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the DR plan execution.
	LifecycleState DrPlanExecutionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time at which DR plan execution began. An RFC3339 formatted datetime string.
	// Example: `2019-03-29T09:36:42Z`
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time at which DR plan execution succeeded, failed, was paused, or canceled.
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

func (m DrPlanExecutionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrPlanExecutionSummary) ValidateEnumValue() (bool, error) {
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
