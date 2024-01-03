// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReplicationSchedule Replication schedule.
type ReplicationSchedule struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the replication schedule.
	Id *string `mandatory:"true" json:"id"`

	// A name of the replication schedule.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Recurrence specification for the replication schedule execution.
	ExecutionRecurrences *string `mandatory:"true" json:"executionRecurrences"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the replication schedule exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Current state of the replication schedule.
	LifecycleState ReplicationScheduleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The detailed state of the replication schedule.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The time when the replication schedule was created in RFC3339 format.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the replication schedule was last updated in RFC3339 format.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ReplicationSchedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReplicationSchedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReplicationScheduleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetReplicationScheduleLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReplicationScheduleLifecycleStateEnum Enum with underlying type: string
type ReplicationScheduleLifecycleStateEnum string

// Set of constants representing the allowable values for ReplicationScheduleLifecycleStateEnum
const (
	ReplicationScheduleLifecycleStateCreating       ReplicationScheduleLifecycleStateEnum = "CREATING"
	ReplicationScheduleLifecycleStateUpdating       ReplicationScheduleLifecycleStateEnum = "UPDATING"
	ReplicationScheduleLifecycleStateNeedsAttention ReplicationScheduleLifecycleStateEnum = "NEEDS_ATTENTION"
	ReplicationScheduleLifecycleStateActive         ReplicationScheduleLifecycleStateEnum = "ACTIVE"
	ReplicationScheduleLifecycleStateDeleting       ReplicationScheduleLifecycleStateEnum = "DELETING"
	ReplicationScheduleLifecycleStateDeleted        ReplicationScheduleLifecycleStateEnum = "DELETED"
	ReplicationScheduleLifecycleStateFailed         ReplicationScheduleLifecycleStateEnum = "FAILED"
)

var mappingReplicationScheduleLifecycleStateEnum = map[string]ReplicationScheduleLifecycleStateEnum{
	"CREATING":        ReplicationScheduleLifecycleStateCreating,
	"UPDATING":        ReplicationScheduleLifecycleStateUpdating,
	"NEEDS_ATTENTION": ReplicationScheduleLifecycleStateNeedsAttention,
	"ACTIVE":          ReplicationScheduleLifecycleStateActive,
	"DELETING":        ReplicationScheduleLifecycleStateDeleting,
	"DELETED":         ReplicationScheduleLifecycleStateDeleted,
	"FAILED":          ReplicationScheduleLifecycleStateFailed,
}

var mappingReplicationScheduleLifecycleStateEnumLowerCase = map[string]ReplicationScheduleLifecycleStateEnum{
	"creating":        ReplicationScheduleLifecycleStateCreating,
	"updating":        ReplicationScheduleLifecycleStateUpdating,
	"needs_attention": ReplicationScheduleLifecycleStateNeedsAttention,
	"active":          ReplicationScheduleLifecycleStateActive,
	"deleting":        ReplicationScheduleLifecycleStateDeleting,
	"deleted":         ReplicationScheduleLifecycleStateDeleted,
	"failed":          ReplicationScheduleLifecycleStateFailed,
}

// GetReplicationScheduleLifecycleStateEnumValues Enumerates the set of values for ReplicationScheduleLifecycleStateEnum
func GetReplicationScheduleLifecycleStateEnumValues() []ReplicationScheduleLifecycleStateEnum {
	values := make([]ReplicationScheduleLifecycleStateEnum, 0)
	for _, v := range mappingReplicationScheduleLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicationScheduleLifecycleStateEnumStringValues Enumerates the set of values in String for ReplicationScheduleLifecycleStateEnum
func GetReplicationScheduleLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"NEEDS_ATTENTION",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingReplicationScheduleLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicationScheduleLifecycleStateEnum(val string) (ReplicationScheduleLifecycleStateEnum, bool) {
	enum, ok := mappingReplicationScheduleLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
