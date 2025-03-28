// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TagDefault Tag defaults let you specify a default tag (tagnamespace.tag="value") to apply to all resource types
// in a specified compartment. The tag default is applied at the time the resource is created. Resources
// that exist in the compartment before you create the tag default are not tagged. The `TagDefault` object
// specifies the tag and compartment details.
// Tag defaults are inherited by child compartments. This means that if you set a tag default on the root compartment
// for a tenancy, all resources that are created in the tenancy are tagged. For more information about
// using tag defaults, see Managing Tag Defaults (https://docs.oracle.com/iaas/Content/Tagging/Tasks/managingtagdefaults.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator.
type TagDefault struct {

	// The OCID of the tag default.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment. The tag default applies to all new resources that get created in the
	// compartment. Resources that existed before the tag default was created are not tagged.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the tag namespace that contains the tag definition.
	TagNamespaceId *string `mandatory:"true" json:"tagNamespaceId"`

	// The OCID of the tag definition. The tag default will always assign a default value for this tag definition.
	TagDefinitionId *string `mandatory:"true" json:"tagDefinitionId"`

	// The name used in the tag definition. This field is informational in the context of the tag default.
	TagDefinitionName *string `mandatory:"true" json:"tagDefinitionName"`

	// The default value for the tag definition. This will be applied to all resources created in the compartment.
	Value *string `mandatory:"true" json:"value"`

	// Date and time the `TagDefault` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// If you specify that a value is required, a value is set during resource creation (either by the
	// user creating the resource or another tag defualt). If no value is set, resource creation is
	// blocked.
	// * If the `isRequired` flag is set to "true", the value is set during resource creation.
	// * If the `isRequired` flag is set to "false", the value you enter is set during resource creation.
	// Example: `false`
	IsRequired *bool `mandatory:"true" json:"isRequired"`

	// The tag default's current state. After creating a `TagDefault`, make sure its `lifecycleState` is ACTIVE before using it.
	LifecycleState TagDefaultLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m TagDefault) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TagDefault) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTagDefaultLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTagDefaultLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TagDefaultLifecycleStateEnum Enum with underlying type: string
type TagDefaultLifecycleStateEnum string

// Set of constants representing the allowable values for TagDefaultLifecycleStateEnum
const (
	TagDefaultLifecycleStateActive TagDefaultLifecycleStateEnum = "ACTIVE"
)

var mappingTagDefaultLifecycleStateEnum = map[string]TagDefaultLifecycleStateEnum{
	"ACTIVE": TagDefaultLifecycleStateActive,
}

var mappingTagDefaultLifecycleStateEnumLowerCase = map[string]TagDefaultLifecycleStateEnum{
	"active": TagDefaultLifecycleStateActive,
}

// GetTagDefaultLifecycleStateEnumValues Enumerates the set of values for TagDefaultLifecycleStateEnum
func GetTagDefaultLifecycleStateEnumValues() []TagDefaultLifecycleStateEnum {
	values := make([]TagDefaultLifecycleStateEnum, 0)
	for _, v := range mappingTagDefaultLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTagDefaultLifecycleStateEnumStringValues Enumerates the set of values in String for TagDefaultLifecycleStateEnum
func GetTagDefaultLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
	}
}

// GetMappingTagDefaultLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTagDefaultLifecycleStateEnum(val string) (TagDefaultLifecycleStateEnum, bool) {
	enum, ok := mappingTagDefaultLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
