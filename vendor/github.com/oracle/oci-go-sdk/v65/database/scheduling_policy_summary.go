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

// SchedulingPolicySummary Details of a Scheduling Policy.
type SchedulingPolicySummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Scheduling Policy.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Scheduling Policy. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Scheduling Policy. Valid states are CREATING, NEEDS_ATTENTION, ACTIVE, UPDATING, FAILED, DELETING and DELETED.
	LifecycleState SchedulingPolicySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The cadence period.
	Cadence SchedulingPolicySummaryCadenceEnum `mandatory:"true" json:"cadence"`

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

func (m SchedulingPolicySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SchedulingPolicySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSchedulingPolicySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSchedulingPolicySummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchedulingPolicySummaryCadenceEnum(string(m.Cadence)); !ok && m.Cadence != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Cadence: %s. Supported values are: %s.", m.Cadence, strings.Join(GetSchedulingPolicySummaryCadenceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SchedulingPolicySummaryLifecycleStateEnum Enum with underlying type: string
type SchedulingPolicySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for SchedulingPolicySummaryLifecycleStateEnum
const (
	SchedulingPolicySummaryLifecycleStateCreating       SchedulingPolicySummaryLifecycleStateEnum = "CREATING"
	SchedulingPolicySummaryLifecycleStateNeedsAttention SchedulingPolicySummaryLifecycleStateEnum = "NEEDS_ATTENTION"
	SchedulingPolicySummaryLifecycleStateAvailable      SchedulingPolicySummaryLifecycleStateEnum = "AVAILABLE"
	SchedulingPolicySummaryLifecycleStateUpdating       SchedulingPolicySummaryLifecycleStateEnum = "UPDATING"
	SchedulingPolicySummaryLifecycleStateFailed         SchedulingPolicySummaryLifecycleStateEnum = "FAILED"
	SchedulingPolicySummaryLifecycleStateDeleting       SchedulingPolicySummaryLifecycleStateEnum = "DELETING"
	SchedulingPolicySummaryLifecycleStateDeleted        SchedulingPolicySummaryLifecycleStateEnum = "DELETED"
)

var mappingSchedulingPolicySummaryLifecycleStateEnum = map[string]SchedulingPolicySummaryLifecycleStateEnum{
	"CREATING":        SchedulingPolicySummaryLifecycleStateCreating,
	"NEEDS_ATTENTION": SchedulingPolicySummaryLifecycleStateNeedsAttention,
	"AVAILABLE":       SchedulingPolicySummaryLifecycleStateAvailable,
	"UPDATING":        SchedulingPolicySummaryLifecycleStateUpdating,
	"FAILED":          SchedulingPolicySummaryLifecycleStateFailed,
	"DELETING":        SchedulingPolicySummaryLifecycleStateDeleting,
	"DELETED":         SchedulingPolicySummaryLifecycleStateDeleted,
}

var mappingSchedulingPolicySummaryLifecycleStateEnumLowerCase = map[string]SchedulingPolicySummaryLifecycleStateEnum{
	"creating":        SchedulingPolicySummaryLifecycleStateCreating,
	"needs_attention": SchedulingPolicySummaryLifecycleStateNeedsAttention,
	"available":       SchedulingPolicySummaryLifecycleStateAvailable,
	"updating":        SchedulingPolicySummaryLifecycleStateUpdating,
	"failed":          SchedulingPolicySummaryLifecycleStateFailed,
	"deleting":        SchedulingPolicySummaryLifecycleStateDeleting,
	"deleted":         SchedulingPolicySummaryLifecycleStateDeleted,
}

// GetSchedulingPolicySummaryLifecycleStateEnumValues Enumerates the set of values for SchedulingPolicySummaryLifecycleStateEnum
func GetSchedulingPolicySummaryLifecycleStateEnumValues() []SchedulingPolicySummaryLifecycleStateEnum {
	values := make([]SchedulingPolicySummaryLifecycleStateEnum, 0)
	for _, v := range mappingSchedulingPolicySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulingPolicySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for SchedulingPolicySummaryLifecycleStateEnum
func GetSchedulingPolicySummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingSchedulingPolicySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulingPolicySummaryLifecycleStateEnum(val string) (SchedulingPolicySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingSchedulingPolicySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchedulingPolicySummaryCadenceEnum Enum with underlying type: string
type SchedulingPolicySummaryCadenceEnum string

// Set of constants representing the allowable values for SchedulingPolicySummaryCadenceEnum
const (
	SchedulingPolicySummaryCadenceHalfyearly SchedulingPolicySummaryCadenceEnum = "HALFYEARLY"
	SchedulingPolicySummaryCadenceQuarterly  SchedulingPolicySummaryCadenceEnum = "QUARTERLY"
	SchedulingPolicySummaryCadenceMonthly    SchedulingPolicySummaryCadenceEnum = "MONTHLY"
)

var mappingSchedulingPolicySummaryCadenceEnum = map[string]SchedulingPolicySummaryCadenceEnum{
	"HALFYEARLY": SchedulingPolicySummaryCadenceHalfyearly,
	"QUARTERLY":  SchedulingPolicySummaryCadenceQuarterly,
	"MONTHLY":    SchedulingPolicySummaryCadenceMonthly,
}

var mappingSchedulingPolicySummaryCadenceEnumLowerCase = map[string]SchedulingPolicySummaryCadenceEnum{
	"halfyearly": SchedulingPolicySummaryCadenceHalfyearly,
	"quarterly":  SchedulingPolicySummaryCadenceQuarterly,
	"monthly":    SchedulingPolicySummaryCadenceMonthly,
}

// GetSchedulingPolicySummaryCadenceEnumValues Enumerates the set of values for SchedulingPolicySummaryCadenceEnum
func GetSchedulingPolicySummaryCadenceEnumValues() []SchedulingPolicySummaryCadenceEnum {
	values := make([]SchedulingPolicySummaryCadenceEnum, 0)
	for _, v := range mappingSchedulingPolicySummaryCadenceEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulingPolicySummaryCadenceEnumStringValues Enumerates the set of values in String for SchedulingPolicySummaryCadenceEnum
func GetSchedulingPolicySummaryCadenceEnumStringValues() []string {
	return []string{
		"HALFYEARLY",
		"QUARTERLY",
		"MONTHLY",
	}
}

// GetMappingSchedulingPolicySummaryCadenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulingPolicySummaryCadenceEnum(val string) (SchedulingPolicySummaryCadenceEnum, bool) {
	enum, ok := mappingSchedulingPolicySummaryCadenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
