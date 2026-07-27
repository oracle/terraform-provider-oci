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

// BlueGreenDeployment A blue/green deployment resource.
type BlueGreenDeployment struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the blue/green deployment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current lifecycle state.
	LifecycleState BlueGreenDeploymentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

	TargetDbSystemDetails *BlueGreenDeploymentTargetDbSystemDetails `mandatory:"false" json:"targetDbSystemDetails"`

	// Replication channel OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ReplicationChannelId *string `mandatory:"false" json:"replicationChannelId"`

	// SSL mode used for the replication channel created by the blue/green workflow.
	SslMode SslModeEnum `mandatory:"false" json:"sslMode,omitempty"`
}

func (m BlueGreenDeployment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BlueGreenDeployment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBlueGreenDeploymentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBlueGreenDeploymentLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBlueGreenDeploymentSwitchoverStatusEnum(string(m.SwitchoverStatus)); !ok && m.SwitchoverStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SwitchoverStatus: %s. Supported values are: %s.", m.SwitchoverStatus, strings.Join(GetBlueGreenDeploymentSwitchoverStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSslModeEnum(string(m.SslMode)); !ok && m.SslMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SslMode: %s. Supported values are: %s.", m.SslMode, strings.Join(GetSslModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BlueGreenDeploymentLifecycleStateEnum Enum with underlying type: string
type BlueGreenDeploymentLifecycleStateEnum string

// Set of constants representing the allowable values for BlueGreenDeploymentLifecycleStateEnum
const (
	BlueGreenDeploymentLifecycleStateCreating BlueGreenDeploymentLifecycleStateEnum = "CREATING"
	BlueGreenDeploymentLifecycleStateActive   BlueGreenDeploymentLifecycleStateEnum = "ACTIVE"
	BlueGreenDeploymentLifecycleStateUpdating BlueGreenDeploymentLifecycleStateEnum = "UPDATING"
	BlueGreenDeploymentLifecycleStateFailed   BlueGreenDeploymentLifecycleStateEnum = "FAILED"
	BlueGreenDeploymentLifecycleStateDeleting BlueGreenDeploymentLifecycleStateEnum = "DELETING"
	BlueGreenDeploymentLifecycleStateDeleted  BlueGreenDeploymentLifecycleStateEnum = "DELETED"
)

var mappingBlueGreenDeploymentLifecycleStateEnum = map[string]BlueGreenDeploymentLifecycleStateEnum{
	"CREATING": BlueGreenDeploymentLifecycleStateCreating,
	"ACTIVE":   BlueGreenDeploymentLifecycleStateActive,
	"UPDATING": BlueGreenDeploymentLifecycleStateUpdating,
	"FAILED":   BlueGreenDeploymentLifecycleStateFailed,
	"DELETING": BlueGreenDeploymentLifecycleStateDeleting,
	"DELETED":  BlueGreenDeploymentLifecycleStateDeleted,
}

var mappingBlueGreenDeploymentLifecycleStateEnumLowerCase = map[string]BlueGreenDeploymentLifecycleStateEnum{
	"creating": BlueGreenDeploymentLifecycleStateCreating,
	"active":   BlueGreenDeploymentLifecycleStateActive,
	"updating": BlueGreenDeploymentLifecycleStateUpdating,
	"failed":   BlueGreenDeploymentLifecycleStateFailed,
	"deleting": BlueGreenDeploymentLifecycleStateDeleting,
	"deleted":  BlueGreenDeploymentLifecycleStateDeleted,
}

// GetBlueGreenDeploymentLifecycleStateEnumValues Enumerates the set of values for BlueGreenDeploymentLifecycleStateEnum
func GetBlueGreenDeploymentLifecycleStateEnumValues() []BlueGreenDeploymentLifecycleStateEnum {
	values := make([]BlueGreenDeploymentLifecycleStateEnum, 0)
	for _, v := range mappingBlueGreenDeploymentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBlueGreenDeploymentLifecycleStateEnumStringValues Enumerates the set of values in String for BlueGreenDeploymentLifecycleStateEnum
func GetBlueGreenDeploymentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingBlueGreenDeploymentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBlueGreenDeploymentLifecycleStateEnum(val string) (BlueGreenDeploymentLifecycleStateEnum, bool) {
	enum, ok := mappingBlueGreenDeploymentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
