// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Discovery and Monitoring Control API
//
// Use the Resource Discovery and Monitoring Control API to get details about monitored instances and perform actions. For more information, see Resource Discovery and Monitoring (https://docs.oracle.com/iaas/os-management/osms/osms-resource-discovery-monitoring.htm).
//

package appmgmtcontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MonitoredInstance Description of Monitored Instance.
type MonitoredInstance struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of monitored instance.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name of the monitored instance. It is binded to Compute Instance (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
	// DisplayName is fetched from Core Service API (https://docs.cloud.oracle.com/api/#/en/iaas/20160918/Instance/).
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Management Agent Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// Used to invoke manage operations on Management Agent Cloud Service.
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	// The time the MonitoredInstance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the MonitoredInstance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Monitoring status. Can be either enabled or disabled.
	MonitoringState MonitoredInstanceMonitoringStateEnum `mandatory:"false" json:"monitoringState,omitempty"`

	// The current state of the monitored instance.
	LifecycleState MonitoredInstanceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m MonitoredInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoredInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMonitoredInstanceMonitoringStateEnum(string(m.MonitoringState)); !ok && m.MonitoringState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MonitoringState: %s. Supported values are: %s.", m.MonitoringState, strings.Join(GetMonitoredInstanceMonitoringStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMonitoredInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMonitoredInstanceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MonitoredInstanceMonitoringStateEnum Enum with underlying type: string
type MonitoredInstanceMonitoringStateEnum string

// Set of constants representing the allowable values for MonitoredInstanceMonitoringStateEnum
const (
	MonitoredInstanceMonitoringStateEnabled  MonitoredInstanceMonitoringStateEnum = "ENABLED"
	MonitoredInstanceMonitoringStateDisabled MonitoredInstanceMonitoringStateEnum = "DISABLED"
)

var mappingMonitoredInstanceMonitoringStateEnum = map[string]MonitoredInstanceMonitoringStateEnum{
	"ENABLED":  MonitoredInstanceMonitoringStateEnabled,
	"DISABLED": MonitoredInstanceMonitoringStateDisabled,
}

var mappingMonitoredInstanceMonitoringStateEnumLowerCase = map[string]MonitoredInstanceMonitoringStateEnum{
	"enabled":  MonitoredInstanceMonitoringStateEnabled,
	"disabled": MonitoredInstanceMonitoringStateDisabled,
}

// GetMonitoredInstanceMonitoringStateEnumValues Enumerates the set of values for MonitoredInstanceMonitoringStateEnum
func GetMonitoredInstanceMonitoringStateEnumValues() []MonitoredInstanceMonitoringStateEnum {
	values := make([]MonitoredInstanceMonitoringStateEnum, 0)
	for _, v := range mappingMonitoredInstanceMonitoringStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitoredInstanceMonitoringStateEnumStringValues Enumerates the set of values in String for MonitoredInstanceMonitoringStateEnum
func GetMonitoredInstanceMonitoringStateEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingMonitoredInstanceMonitoringStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitoredInstanceMonitoringStateEnum(val string) (MonitoredInstanceMonitoringStateEnum, bool) {
	enum, ok := mappingMonitoredInstanceMonitoringStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MonitoredInstanceLifecycleStateEnum Enum with underlying type: string
type MonitoredInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for MonitoredInstanceLifecycleStateEnum
const (
	MonitoredInstanceLifecycleStateCreating MonitoredInstanceLifecycleStateEnum = "CREATING"
	MonitoredInstanceLifecycleStateUpdating MonitoredInstanceLifecycleStateEnum = "UPDATING"
	MonitoredInstanceLifecycleStateActive   MonitoredInstanceLifecycleStateEnum = "ACTIVE"
	MonitoredInstanceLifecycleStateInactive MonitoredInstanceLifecycleStateEnum = "INACTIVE"
	MonitoredInstanceLifecycleStateDeleting MonitoredInstanceLifecycleStateEnum = "DELETING"
	MonitoredInstanceLifecycleStateDeleted  MonitoredInstanceLifecycleStateEnum = "DELETED"
	MonitoredInstanceLifecycleStateFailed   MonitoredInstanceLifecycleStateEnum = "FAILED"
)

var mappingMonitoredInstanceLifecycleStateEnum = map[string]MonitoredInstanceLifecycleStateEnum{
	"CREATING": MonitoredInstanceLifecycleStateCreating,
	"UPDATING": MonitoredInstanceLifecycleStateUpdating,
	"ACTIVE":   MonitoredInstanceLifecycleStateActive,
	"INACTIVE": MonitoredInstanceLifecycleStateInactive,
	"DELETING": MonitoredInstanceLifecycleStateDeleting,
	"DELETED":  MonitoredInstanceLifecycleStateDeleted,
	"FAILED":   MonitoredInstanceLifecycleStateFailed,
}

var mappingMonitoredInstanceLifecycleStateEnumLowerCase = map[string]MonitoredInstanceLifecycleStateEnum{
	"creating": MonitoredInstanceLifecycleStateCreating,
	"updating": MonitoredInstanceLifecycleStateUpdating,
	"active":   MonitoredInstanceLifecycleStateActive,
	"inactive": MonitoredInstanceLifecycleStateInactive,
	"deleting": MonitoredInstanceLifecycleStateDeleting,
	"deleted":  MonitoredInstanceLifecycleStateDeleted,
	"failed":   MonitoredInstanceLifecycleStateFailed,
}

// GetMonitoredInstanceLifecycleStateEnumValues Enumerates the set of values for MonitoredInstanceLifecycleStateEnum
func GetMonitoredInstanceLifecycleStateEnumValues() []MonitoredInstanceLifecycleStateEnum {
	values := make([]MonitoredInstanceLifecycleStateEnum, 0)
	for _, v := range mappingMonitoredInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitoredInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for MonitoredInstanceLifecycleStateEnum
func GetMonitoredInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingMonitoredInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitoredInstanceLifecycleStateEnum(val string) (MonitoredInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingMonitoredInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
