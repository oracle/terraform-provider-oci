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

// SchedulingPolicy Details of a Scheduling Policy.
type SchedulingPolicy struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Scheduling Policy.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Scheduling Policy. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Scheduling Policy. Valid states are CREATING, NEEDS_ATTENTION, ACTIVE, UPDATING, FAILED, DELETING and DELETED.
	LifecycleState SchedulingPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The cadence period.
	Cadence SchedulingPolicyCadenceEnum `mandatory:"true" json:"cadence"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the Scheduling Policy was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last date and time that the Scheduling Policy was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Start of the month to be followed during the cadence period.
	CadenceStartMonth *Month `mandatory:"false" json:"cadenceStartMonth"`

	// The date and time of the next scheduling window associated with the schedulingPolicy is planned to start.
	TimeNextWindowStarts *common.SDKTime `mandatory:"false" json:"timeNextWindowStarts"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m SchedulingPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SchedulingPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSchedulingPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSchedulingPolicyLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchedulingPolicyCadenceEnum(string(m.Cadence)); !ok && m.Cadence != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Cadence: %s. Supported values are: %s.", m.Cadence, strings.Join(GetSchedulingPolicyCadenceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SchedulingPolicyLifecycleStateEnum Enum with underlying type: string
type SchedulingPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for SchedulingPolicyLifecycleStateEnum
const (
	SchedulingPolicyLifecycleStateCreating       SchedulingPolicyLifecycleStateEnum = "CREATING"
	SchedulingPolicyLifecycleStateNeedsAttention SchedulingPolicyLifecycleStateEnum = "NEEDS_ATTENTION"
	SchedulingPolicyLifecycleStateAvailable      SchedulingPolicyLifecycleStateEnum = "AVAILABLE"
	SchedulingPolicyLifecycleStateUpdating       SchedulingPolicyLifecycleStateEnum = "UPDATING"
	SchedulingPolicyLifecycleStateFailed         SchedulingPolicyLifecycleStateEnum = "FAILED"
	SchedulingPolicyLifecycleStateDeleting       SchedulingPolicyLifecycleStateEnum = "DELETING"
	SchedulingPolicyLifecycleStateDeleted        SchedulingPolicyLifecycleStateEnum = "DELETED"
)

var mappingSchedulingPolicyLifecycleStateEnum = map[string]SchedulingPolicyLifecycleStateEnum{
	"CREATING":        SchedulingPolicyLifecycleStateCreating,
	"NEEDS_ATTENTION": SchedulingPolicyLifecycleStateNeedsAttention,
	"AVAILABLE":       SchedulingPolicyLifecycleStateAvailable,
	"UPDATING":        SchedulingPolicyLifecycleStateUpdating,
	"FAILED":          SchedulingPolicyLifecycleStateFailed,
	"DELETING":        SchedulingPolicyLifecycleStateDeleting,
	"DELETED":         SchedulingPolicyLifecycleStateDeleted,
}

var mappingSchedulingPolicyLifecycleStateEnumLowerCase = map[string]SchedulingPolicyLifecycleStateEnum{
	"creating":        SchedulingPolicyLifecycleStateCreating,
	"needs_attention": SchedulingPolicyLifecycleStateNeedsAttention,
	"available":       SchedulingPolicyLifecycleStateAvailable,
	"updating":        SchedulingPolicyLifecycleStateUpdating,
	"failed":          SchedulingPolicyLifecycleStateFailed,
	"deleting":        SchedulingPolicyLifecycleStateDeleting,
	"deleted":         SchedulingPolicyLifecycleStateDeleted,
}

// GetSchedulingPolicyLifecycleStateEnumValues Enumerates the set of values for SchedulingPolicyLifecycleStateEnum
func GetSchedulingPolicyLifecycleStateEnumValues() []SchedulingPolicyLifecycleStateEnum {
	values := make([]SchedulingPolicyLifecycleStateEnum, 0)
	for _, v := range mappingSchedulingPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulingPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for SchedulingPolicyLifecycleStateEnum
func GetSchedulingPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"NEEDS_ATTENTION",
		"AVAILABLE",
		"UPDATING",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingSchedulingPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulingPolicyLifecycleStateEnum(val string) (SchedulingPolicyLifecycleStateEnum, bool) {
	enum, ok := mappingSchedulingPolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchedulingPolicyCadenceEnum Enum with underlying type: string
type SchedulingPolicyCadenceEnum string

// Set of constants representing the allowable values for SchedulingPolicyCadenceEnum
const (
	SchedulingPolicyCadenceHalfyearly SchedulingPolicyCadenceEnum = "HALFYEARLY"
	SchedulingPolicyCadenceQuarterly  SchedulingPolicyCadenceEnum = "QUARTERLY"
	SchedulingPolicyCadenceMonthly    SchedulingPolicyCadenceEnum = "MONTHLY"
)

var mappingSchedulingPolicyCadenceEnum = map[string]SchedulingPolicyCadenceEnum{
	"HALFYEARLY": SchedulingPolicyCadenceHalfyearly,
	"QUARTERLY":  SchedulingPolicyCadenceQuarterly,
	"MONTHLY":    SchedulingPolicyCadenceMonthly,
}

var mappingSchedulingPolicyCadenceEnumLowerCase = map[string]SchedulingPolicyCadenceEnum{
	"halfyearly": SchedulingPolicyCadenceHalfyearly,
	"quarterly":  SchedulingPolicyCadenceQuarterly,
	"monthly":    SchedulingPolicyCadenceMonthly,
}

// GetSchedulingPolicyCadenceEnumValues Enumerates the set of values for SchedulingPolicyCadenceEnum
func GetSchedulingPolicyCadenceEnumValues() []SchedulingPolicyCadenceEnum {
	values := make([]SchedulingPolicyCadenceEnum, 0)
	for _, v := range mappingSchedulingPolicyCadenceEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulingPolicyCadenceEnumStringValues Enumerates the set of values in String for SchedulingPolicyCadenceEnum
func GetSchedulingPolicyCadenceEnumStringValues() []string {
	return []string{
		"HALFYEARLY",
		"QUARTERLY",
		"MONTHLY",
	}
}

// GetMappingSchedulingPolicyCadenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulingPolicyCadenceEnum(val string) (SchedulingPolicyCadenceEnum, bool) {
	enum, ok := mappingSchedulingPolicyCadenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
