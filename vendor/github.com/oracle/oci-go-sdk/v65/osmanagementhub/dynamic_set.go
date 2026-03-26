// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DynamicSet An object that defines the dynamic set.
type DynamicSet struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dynamic set.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the dynamic set.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// User-friendly name for the dynamic set.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the dynamic set was created (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the event.
	LifecycleState DynamicSetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// User-specified description for the dynamic set.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the dynamic set was last updated (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Number of scheduled jobs currently targeting this dynamic set.
	ScheduledJobCount *string `mandatory:"false" json:"scheduledJobCount"`

	// The list of compartment details.
	TargetCompartments []TargetCompartmentDetails `mandatory:"false" json:"targetCompartments"`

	// Include either any or all attributes.
	MatchType MatchTypeEnum `mandatory:"false" json:"matchType,omitempty"`

	MatchingRule *MatchingRule `mandatory:"false" json:"matchingRule"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DynamicSet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DynamicSet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDynamicSetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDynamicSetLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMatchTypeEnum(string(m.MatchType)); !ok && m.MatchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MatchType: %s. Supported values are: %s.", m.MatchType, strings.Join(GetMatchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DynamicSetLifecycleStateEnum Enum with underlying type: string
type DynamicSetLifecycleStateEnum string

// Set of constants representing the allowable values for DynamicSetLifecycleStateEnum
const (
	DynamicSetLifecycleStateCreating DynamicSetLifecycleStateEnum = "CREATING"
	DynamicSetLifecycleStateUpdating DynamicSetLifecycleStateEnum = "UPDATING"
	DynamicSetLifecycleStateActive   DynamicSetLifecycleStateEnum = "ACTIVE"
	DynamicSetLifecycleStateDeleting DynamicSetLifecycleStateEnum = "DELETING"
	DynamicSetLifecycleStateDeleted  DynamicSetLifecycleStateEnum = "DELETED"
	DynamicSetLifecycleStateFailed   DynamicSetLifecycleStateEnum = "FAILED"
)

var mappingDynamicSetLifecycleStateEnum = map[string]DynamicSetLifecycleStateEnum{
	"CREATING": DynamicSetLifecycleStateCreating,
	"UPDATING": DynamicSetLifecycleStateUpdating,
	"ACTIVE":   DynamicSetLifecycleStateActive,
	"DELETING": DynamicSetLifecycleStateDeleting,
	"DELETED":  DynamicSetLifecycleStateDeleted,
	"FAILED":   DynamicSetLifecycleStateFailed,
}

var mappingDynamicSetLifecycleStateEnumLowerCase = map[string]DynamicSetLifecycleStateEnum{
	"creating": DynamicSetLifecycleStateCreating,
	"updating": DynamicSetLifecycleStateUpdating,
	"active":   DynamicSetLifecycleStateActive,
	"deleting": DynamicSetLifecycleStateDeleting,
	"deleted":  DynamicSetLifecycleStateDeleted,
	"failed":   DynamicSetLifecycleStateFailed,
}

// GetDynamicSetLifecycleStateEnumValues Enumerates the set of values for DynamicSetLifecycleStateEnum
func GetDynamicSetLifecycleStateEnumValues() []DynamicSetLifecycleStateEnum {
	values := make([]DynamicSetLifecycleStateEnum, 0)
	for _, v := range mappingDynamicSetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDynamicSetLifecycleStateEnumStringValues Enumerates the set of values in String for DynamicSetLifecycleStateEnum
func GetDynamicSetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDynamicSetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDynamicSetLifecycleStateEnum(val string) (DynamicSetLifecycleStateEnum, bool) {
	enum, ok := mappingDynamicSetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
