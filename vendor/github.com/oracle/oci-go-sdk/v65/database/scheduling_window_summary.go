// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SchedulingWindowSummary Details of a Scheduling Window.
type SchedulingWindowSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Scheduling Window.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Scheduling Window. Valid states are CREATING, ACTIVE, UPDATING, FAILED, DELETING and DELETED.
	LifecycleState SchedulingWindowSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	WindowPreference *WindowPreferenceDetail `mandatory:"true" json:"windowPreference"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Scheduling Policy.
	SchedulingPolicyId *string `mandatory:"true" json:"schedulingPolicyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The user-friendly name for the Scheduling Window. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time of the next upcoming window associated within the schedulingWindow is planned to start.
	TimeNextSchedulingWindowStarts *common.SDKTime `mandatory:"false" json:"timeNextSchedulingWindowStarts"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the Scheduling Window was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last date and time that the Scheduling Window was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m SchedulingWindowSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SchedulingWindowSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSchedulingWindowSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSchedulingWindowSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SchedulingWindowSummaryLifecycleStateEnum Enum with underlying type: string
type SchedulingWindowSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for SchedulingWindowSummaryLifecycleStateEnum
const (
	SchedulingWindowSummaryLifecycleStateCreating  SchedulingWindowSummaryLifecycleStateEnum = "CREATING"
	SchedulingWindowSummaryLifecycleStateAvailable SchedulingWindowSummaryLifecycleStateEnum = "AVAILABLE"
	SchedulingWindowSummaryLifecycleStateUpdating  SchedulingWindowSummaryLifecycleStateEnum = "UPDATING"
	SchedulingWindowSummaryLifecycleStateFailed    SchedulingWindowSummaryLifecycleStateEnum = "FAILED"
	SchedulingWindowSummaryLifecycleStateDeleting  SchedulingWindowSummaryLifecycleStateEnum = "DELETING"
	SchedulingWindowSummaryLifecycleStateDeleted   SchedulingWindowSummaryLifecycleStateEnum = "DELETED"
)

var mappingSchedulingWindowSummaryLifecycleStateEnum = map[string]SchedulingWindowSummaryLifecycleStateEnum{
	"CREATING":  SchedulingWindowSummaryLifecycleStateCreating,
	"AVAILABLE": SchedulingWindowSummaryLifecycleStateAvailable,
	"UPDATING":  SchedulingWindowSummaryLifecycleStateUpdating,
	"FAILED":    SchedulingWindowSummaryLifecycleStateFailed,
	"DELETING":  SchedulingWindowSummaryLifecycleStateDeleting,
	"DELETED":   SchedulingWindowSummaryLifecycleStateDeleted,
}

var mappingSchedulingWindowSummaryLifecycleStateEnumLowerCase = map[string]SchedulingWindowSummaryLifecycleStateEnum{
	"creating":  SchedulingWindowSummaryLifecycleStateCreating,
	"available": SchedulingWindowSummaryLifecycleStateAvailable,
	"updating":  SchedulingWindowSummaryLifecycleStateUpdating,
	"failed":    SchedulingWindowSummaryLifecycleStateFailed,
	"deleting":  SchedulingWindowSummaryLifecycleStateDeleting,
	"deleted":   SchedulingWindowSummaryLifecycleStateDeleted,
}

// GetSchedulingWindowSummaryLifecycleStateEnumValues Enumerates the set of values for SchedulingWindowSummaryLifecycleStateEnum
func GetSchedulingWindowSummaryLifecycleStateEnumValues() []SchedulingWindowSummaryLifecycleStateEnum {
	values := make([]SchedulingWindowSummaryLifecycleStateEnum, 0)
	for _, v := range mappingSchedulingWindowSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulingWindowSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for SchedulingWindowSummaryLifecycleStateEnum
func GetSchedulingWindowSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"AVAILABLE",
		"UPDATING",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingSchedulingWindowSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulingWindowSummaryLifecycleStateEnum(val string) (SchedulingWindowSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingSchedulingWindowSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
