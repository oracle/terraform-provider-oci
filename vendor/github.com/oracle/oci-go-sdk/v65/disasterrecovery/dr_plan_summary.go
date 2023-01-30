// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (FSDR) API to manage disaster recovery for business applications.
// FSDR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster recovery
// capabilities for all layers of an application stack, including infrastructure, middleware, database, and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrPlanSummary Summary information about a DR Plan Execution.
type DrPlanSummary struct {

	// The OCID of this DR Plan.
	// Example: `ocid1.drplan.oc1.iad.exampleocid2`
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the DR Plan.
	// Example: `ocid1.compartment.oc1..exampleocid1`
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of this DR Plan.
	// Example: `EBS Switchover PHX to IAD`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of this DR Plan.
	Type DrPlanTypeEnum `mandatory:"true" json:"type"`

	// The OCID of the DR Protection Group with which this DR Plan is associated.
	// Example: `ocid1.drprotectiongroup.oc1.iad.exampleocid2`
	DrProtectionGroupId *string `mandatory:"true" json:"drProtectionGroupId"`

	// The OCID of peer (remote) DR Protection Group associated with this plan execution's
	// DR Protection Group.
	// Example: `ocid1.drprotectiongroup.oc1.phx.exampleocid1`
	PeerDrProtectionGroupId *string `mandatory:"true" json:"peerDrProtectionGroupId"`

	// The region of the peer (remote) DR Protection Group.
	// Example: `us-phoenix-1`
	PeerRegion *string `mandatory:"true" json:"peerRegion"`

	// The date and time the DR Plan was created. An RFC3339 formatted datetime string.
	// Example: `2019-03-29T09:36:42Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the DR Plan was updated. An RFC3339 formatted datetime string.
	// Example: `2019-03-29T09:36:42Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the DR Plan.
	LifecycleState DrPlanLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A message describing the DR Plan's current state in more detail.
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

func (m DrPlanSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrPlanSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrPlanTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDrPlanTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDrPlanLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDrPlanLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
