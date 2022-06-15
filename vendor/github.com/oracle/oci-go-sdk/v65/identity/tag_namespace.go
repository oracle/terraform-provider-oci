// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TagNamespace A managed container for defined tags. A tag namespace is unique in a tenancy. For more information,
// see Managing Tags and Tag Namespaces (https://docs.cloud.oracle.com/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values
// using the API.
type TagNamespace struct {

	// The OCID of the tag namespace.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the tag namespace.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the tag namespace. It must be unique across all tag namespaces in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the tag namespace.
	Description *string `mandatory:"true" json:"description"`

	// Whether the tag namespace is retired.
	// See Retiring Key Definitions and Namespace Definitions (https://docs.cloud.oracle.com/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm#retiringkeys).
	IsRetired *bool `mandatory:"true" json:"isRetired"`

	// Date and time the tagNamespace was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The tagnamespace's current state. After creating a tagnamespace, make sure its `lifecycleState` is ACTIVE before using it. After retiring a tagnamespace, make sure its `lifecycleState` is INACTIVE before using it.
	LifecycleState TagNamespaceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m TagNamespace) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TagNamespace) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTagNamespaceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTagNamespaceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TagNamespaceLifecycleStateEnum Enum with underlying type: string
type TagNamespaceLifecycleStateEnum string

// Set of constants representing the allowable values for TagNamespaceLifecycleStateEnum
const (
	TagNamespaceLifecycleStateActive   TagNamespaceLifecycleStateEnum = "ACTIVE"
	TagNamespaceLifecycleStateInactive TagNamespaceLifecycleStateEnum = "INACTIVE"
	TagNamespaceLifecycleStateDeleting TagNamespaceLifecycleStateEnum = "DELETING"
	TagNamespaceLifecycleStateDeleted  TagNamespaceLifecycleStateEnum = "DELETED"
)

var mappingTagNamespaceLifecycleStateEnum = map[string]TagNamespaceLifecycleStateEnum{
	"ACTIVE":   TagNamespaceLifecycleStateActive,
	"INACTIVE": TagNamespaceLifecycleStateInactive,
	"DELETING": TagNamespaceLifecycleStateDeleting,
	"DELETED":  TagNamespaceLifecycleStateDeleted,
}

var mappingTagNamespaceLifecycleStateEnumLowerCase = map[string]TagNamespaceLifecycleStateEnum{
	"active":   TagNamespaceLifecycleStateActive,
	"inactive": TagNamespaceLifecycleStateInactive,
	"deleting": TagNamespaceLifecycleStateDeleting,
	"deleted":  TagNamespaceLifecycleStateDeleted,
}

// GetTagNamespaceLifecycleStateEnumValues Enumerates the set of values for TagNamespaceLifecycleStateEnum
func GetTagNamespaceLifecycleStateEnumValues() []TagNamespaceLifecycleStateEnum {
	values := make([]TagNamespaceLifecycleStateEnum, 0)
	for _, v := range mappingTagNamespaceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTagNamespaceLifecycleStateEnumStringValues Enumerates the set of values in String for TagNamespaceLifecycleStateEnum
func GetTagNamespaceLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingTagNamespaceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTagNamespaceLifecycleStateEnum(val string) (TagNamespaceLifecycleStateEnum, bool) {
	enum, ok := mappingTagNamespaceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
