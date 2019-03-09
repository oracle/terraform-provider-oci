// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TagDefaultSummary A document that summarizes the default value for a Tag Definition for all resource types created in a Compartment.
type TagDefaultSummary struct {

	// The OCID of the Tag Default.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the Compartment. The Tag Default will apply to any resource contained in this Compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the Tag Namespace that contains the Tag Definition.
	TagNamespaceId *string `mandatory:"true" json:"tagNamespaceId"`

	// The OCID of the Tag Definition. The Tag Default will always assign a default value for this Tag Definition.
	TagDefinitionId *string `mandatory:"true" json:"tagDefinitionId"`

	// The name used in the Tag Definition. This field is informational in the context of the Tag Default.
	TagDefinitionName *string `mandatory:"true" json:"tagDefinitionName"`

	// The default value for the Tag Definition. This will be applied to all resources created in the Compartment.
	Value *string `mandatory:"true" json:"value"`

	// Date and time the `TagDefault` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The tag default's current state. After creating a tagdefault, make sure its `lifecycleState` is ACTIVE before using it.
	LifecycleState TagDefaultSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m TagDefaultSummary) String() string {
	return common.PointerString(m)
}

// TagDefaultSummaryLifecycleStateEnum Enum with underlying type: string
type TagDefaultSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for TagDefaultSummaryLifecycleStateEnum
const (
	TagDefaultSummaryLifecycleStateActive TagDefaultSummaryLifecycleStateEnum = "ACTIVE"
)

var mappingTagDefaultSummaryLifecycleState = map[string]TagDefaultSummaryLifecycleStateEnum{
	"ACTIVE": TagDefaultSummaryLifecycleStateActive,
}

// GetTagDefaultSummaryLifecycleStateEnumValues Enumerates the set of values for TagDefaultSummaryLifecycleStateEnum
func GetTagDefaultSummaryLifecycleStateEnumValues() []TagDefaultSummaryLifecycleStateEnum {
	values := make([]TagDefaultSummaryLifecycleStateEnum, 0)
	for _, v := range mappingTagDefaultSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
