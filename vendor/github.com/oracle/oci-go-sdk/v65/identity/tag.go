// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Tag A tag definition that belongs to a specific tag namespace.  "Defined tags" must be set up in your tenancy before
// you can apply them to resources.
// For more information, see Managing Tags and Tag Namespaces (https://docs.oracle.com/iaas/Content/Identity/Concepts/taggingoverview.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values
// using the API.
type Tag struct {

	// The OCID of the compartment that contains the tag definition.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the namespace that contains the tag definition.
	TagNamespaceId *string `mandatory:"true" json:"tagNamespaceId"`

	// The name of the tag namespace that contains the tag definition.
	TagNamespaceName *string `mandatory:"true" json:"tagNamespaceName"`

	// The OCID of the tag definition.
	Id *string `mandatory:"true" json:"id"`

	// The name assigned to the tag during creation. This is the tag key definition.
	// The name must be unique within the tag namespace and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the tag.
	Description *string `mandatory:"true" json:"description"`

	// Indicates whether the tag is retired.
	// See Retiring Key Definitions and Namespace Definitions (https://docs.oracle.com/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm#retiringkeys).
	IsRetired *bool `mandatory:"true" json:"isRetired"`

	// Date and time the tag was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The tag's current state. After creating a tag, make sure its `lifecycleState` is ACTIVE before using it. After retiring a tag, make sure its `lifecycleState` is INACTIVE before using it. If you delete a tag, you cannot delete another tag until the deleted tag's `lifecycleState` changes from DELETING to DELETED.
	LifecycleState TagLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Indicates whether the tag is enabled for cost tracking.
	IsCostTracking *bool `mandatory:"false" json:"isCostTracking"`

	Validator BaseTagDefinitionValidator `mandatory:"false" json:"validator"`
}

func (m Tag) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Tag) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTagLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTagLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Tag) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		LifecycleState   TagLifecycleStateEnum             `json:"lifecycleState"`
		IsCostTracking   *bool                             `json:"isCostTracking"`
		Validator        basetagdefinitionvalidator        `json:"validator"`
		CompartmentId    *string                           `json:"compartmentId"`
		TagNamespaceId   *string                           `json:"tagNamespaceId"`
		TagNamespaceName *string                           `json:"tagNamespaceName"`
		Id               *string                           `json:"id"`
		Name             *string                           `json:"name"`
		Description      *string                           `json:"description"`
		IsRetired        *bool                             `json:"isRetired"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.LifecycleState = model.LifecycleState

	m.IsCostTracking = model.IsCostTracking

	nn, e = model.Validator.UnmarshalPolymorphicJSON(model.Validator.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Validator = nn.(BaseTagDefinitionValidator)
	} else {
		m.Validator = nil
	}

	m.CompartmentId = model.CompartmentId

	m.TagNamespaceId = model.TagNamespaceId

	m.TagNamespaceName = model.TagNamespaceName

	m.Id = model.Id

	m.Name = model.Name

	m.Description = model.Description

	m.IsRetired = model.IsRetired

	m.TimeCreated = model.TimeCreated

	return
}

// TagLifecycleStateEnum Enum with underlying type: string
type TagLifecycleStateEnum string

// Set of constants representing the allowable values for TagLifecycleStateEnum
const (
	TagLifecycleStateActive   TagLifecycleStateEnum = "ACTIVE"
	TagLifecycleStateInactive TagLifecycleStateEnum = "INACTIVE"
	TagLifecycleStateDeleting TagLifecycleStateEnum = "DELETING"
	TagLifecycleStateDeleted  TagLifecycleStateEnum = "DELETED"
)

var mappingTagLifecycleStateEnum = map[string]TagLifecycleStateEnum{
	"ACTIVE":   TagLifecycleStateActive,
	"INACTIVE": TagLifecycleStateInactive,
	"DELETING": TagLifecycleStateDeleting,
	"DELETED":  TagLifecycleStateDeleted,
}

var mappingTagLifecycleStateEnumLowerCase = map[string]TagLifecycleStateEnum{
	"active":   TagLifecycleStateActive,
	"inactive": TagLifecycleStateInactive,
	"deleting": TagLifecycleStateDeleting,
	"deleted":  TagLifecycleStateDeleted,
}

// GetTagLifecycleStateEnumValues Enumerates the set of values for TagLifecycleStateEnum
func GetTagLifecycleStateEnumValues() []TagLifecycleStateEnum {
	values := make([]TagLifecycleStateEnum, 0)
	for _, v := range mappingTagLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTagLifecycleStateEnumStringValues Enumerates the set of values in String for TagLifecycleStateEnum
func GetTagLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingTagLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTagLifecycleStateEnum(val string) (TagLifecycleStateEnum, bool) {
	enum, ok := mappingTagLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
