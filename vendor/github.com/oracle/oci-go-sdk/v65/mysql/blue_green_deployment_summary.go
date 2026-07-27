// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BlueGreenDeploymentSummary List-safe summary of a blue/green deployment.
type BlueGreenDeploymentSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the blue/green deployment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current lifecycle state.
	LifecycleState BlueGreenDeploymentSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the deployment was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The display name of the blue/green deployment.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Additional lifecycle details.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Blue/original source DB system OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the deployment pair.
	SourceDbSystemId *string `mandatory:"false" json:"sourceDbSystemId"`

	// Green/target DB system OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the deployment pair.
	TargetDbSystemId *string `mandatory:"false" json:"targetDbSystemId"`

	// The DB system OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that currently owns the client-facing VIP and serves traffic.
	ActiveDbSystemId *string `mandatory:"false" json:"activeDbSystemId"`

	// Stage of the most recent switchover workflow.
	// `SWITCHOVER_FAILED` indicates terminal switchover failure.
	SwitchoverStatus BlueGreenDeploymentSwitchoverStatusEnum `mandatory:"false" json:"switchoverStatus,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The time the deployment was last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m BlueGreenDeploymentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BlueGreenDeploymentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBlueGreenDeploymentSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBlueGreenDeploymentSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBlueGreenDeploymentSwitchoverStatusEnum(string(m.SwitchoverStatus)); !ok && m.SwitchoverStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SwitchoverStatus: %s. Supported values are: %s.", m.SwitchoverStatus, strings.Join(GetBlueGreenDeploymentSwitchoverStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BlueGreenDeploymentSummaryLifecycleStateEnum Enum with underlying type: string
type BlueGreenDeploymentSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for BlueGreenDeploymentSummaryLifecycleStateEnum
const (
	BlueGreenDeploymentSummaryLifecycleStateCreating BlueGreenDeploymentSummaryLifecycleStateEnum = "CREATING"
	BlueGreenDeploymentSummaryLifecycleStateActive   BlueGreenDeploymentSummaryLifecycleStateEnum = "ACTIVE"
	BlueGreenDeploymentSummaryLifecycleStateUpdating BlueGreenDeploymentSummaryLifecycleStateEnum = "UPDATING"
	BlueGreenDeploymentSummaryLifecycleStateFailed   BlueGreenDeploymentSummaryLifecycleStateEnum = "FAILED"
	BlueGreenDeploymentSummaryLifecycleStateDeleting BlueGreenDeploymentSummaryLifecycleStateEnum = "DELETING"
	BlueGreenDeploymentSummaryLifecycleStateDeleted  BlueGreenDeploymentSummaryLifecycleStateEnum = "DELETED"
)

var mappingBlueGreenDeploymentSummaryLifecycleStateEnum = map[string]BlueGreenDeploymentSummaryLifecycleStateEnum{
	"CREATING": BlueGreenDeploymentSummaryLifecycleStateCreating,
	"ACTIVE":   BlueGreenDeploymentSummaryLifecycleStateActive,
	"UPDATING": BlueGreenDeploymentSummaryLifecycleStateUpdating,
	"FAILED":   BlueGreenDeploymentSummaryLifecycleStateFailed,
	"DELETING": BlueGreenDeploymentSummaryLifecycleStateDeleting,
	"DELETED":  BlueGreenDeploymentSummaryLifecycleStateDeleted,
}

var mappingBlueGreenDeploymentSummaryLifecycleStateEnumLowerCase = map[string]BlueGreenDeploymentSummaryLifecycleStateEnum{
	"creating": BlueGreenDeploymentSummaryLifecycleStateCreating,
	"active":   BlueGreenDeploymentSummaryLifecycleStateActive,
	"updating": BlueGreenDeploymentSummaryLifecycleStateUpdating,
	"failed":   BlueGreenDeploymentSummaryLifecycleStateFailed,
	"deleting": BlueGreenDeploymentSummaryLifecycleStateDeleting,
	"deleted":  BlueGreenDeploymentSummaryLifecycleStateDeleted,
}

// GetBlueGreenDeploymentSummaryLifecycleStateEnumValues Enumerates the set of values for BlueGreenDeploymentSummaryLifecycleStateEnum
func GetBlueGreenDeploymentSummaryLifecycleStateEnumValues() []BlueGreenDeploymentSummaryLifecycleStateEnum {
	values := make([]BlueGreenDeploymentSummaryLifecycleStateEnum, 0)
	for _, v := range mappingBlueGreenDeploymentSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBlueGreenDeploymentSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for BlueGreenDeploymentSummaryLifecycleStateEnum
func GetBlueGreenDeploymentSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingBlueGreenDeploymentSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBlueGreenDeploymentSummaryLifecycleStateEnum(val string) (BlueGreenDeploymentSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingBlueGreenDeploymentSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
