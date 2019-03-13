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

// TagDefault A document that specifies a default value for a Tag Definition for all resource types created in a Compartment.
// Tag Defaults are inherited by child compartments. This means that if you set a Tag Default on the root Compartment
// for a tenancy, all resources are guaranteed to be created with the referenced Tag Definition applied.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator.
type TagDefault struct {

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
	LifecycleState TagDefaultLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m TagDefault) String() string {
	return common.PointerString(m)
}

// TagDefaultLifecycleStateEnum Enum with underlying type: string
type TagDefaultLifecycleStateEnum string

// Set of constants representing the allowable values for TagDefaultLifecycleStateEnum
const (
	TagDefaultLifecycleStateActive TagDefaultLifecycleStateEnum = "ACTIVE"
)

var mappingTagDefaultLifecycleState = map[string]TagDefaultLifecycleStateEnum{
	"ACTIVE": TagDefaultLifecycleStateActive,
}

// GetTagDefaultLifecycleStateEnumValues Enumerates the set of values for TagDefaultLifecycleStateEnum
func GetTagDefaultLifecycleStateEnumValues() []TagDefaultLifecycleStateEnum {
	values := make([]TagDefaultLifecycleStateEnum, 0)
	for _, v := range mappingTagDefaultLifecycleState {
		values = append(values, v)
	}
	return values
}
