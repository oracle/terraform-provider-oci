// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ZprTag A zpr tag definition that belongs to a specific zpr tag namespace. "ZPR Defined tags" must be set up in your tenancy before
// you can apply them to resources.
// For more information, see Managing Tags and zpr tag Namespaces (https://docs.cloud.oracle.com/Content/Identity/Concepts/taggingoverview.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values
// using the API.
type ZprTag struct {

	// The OCID of the compartment that contains the zpr tag definition.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the zpr tag namespace that contains the zpr tag definition.
	ZprTagNamespaceId *string `mandatory:"true" json:"zprTagNamespaceId"`

	// The name of the zpr tag namespace that contains the zpr tag definition.
	ZprTagNamespaceName *string `mandatory:"true" json:"zprTagNamespaceName"`

	// The OCID of the zpr tag definition.
	Id *string `mandatory:"true" json:"id"`

	// The name assigned to the zpr tag during creation. This is the zpr tag key definition.
	// The name must be unique within the zpr tag namespace and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the zpr tag.
	Description *string `mandatory:"true" json:"description"`

	// Indicates whether the zpr tag is retired.
	// See Retiring Key Definitions and Namespace Definitions (https://docs.cloud.oracle.com/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm#retiringkeys).
	IsRetired *bool `mandatory:"true" json:"isRetired"`

	// Date and time the zpr tag was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The data type of the zpr tag.
	Type *string `mandatory:"false" json:"type"`

	// The tag's current state. After creating a tag, make sure its `lifecycleState` is ACTIVE before using it. After retiring a tag, make sure its `lifecycleState` is INACTIVE before using it. If you delete a tag, you cannot delete another zpr tag until the deleted tag's `lifecycleState` changes from DELETING to DELETED.
	LifecycleState ZprTagLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	Validator BaseTagDefinitionValidator `mandatory:"false" json:"validator"`
}

func (m ZprTag) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ZprTag) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingZprTagLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetZprTagLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ZprTag) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Type                *string                    `json:"type"`
		LifecycleState      ZprTagLifecycleStateEnum   `json:"lifecycleState"`
		Validator           basetagdefinitionvalidator `json:"validator"`
		CompartmentId       *string                    `json:"compartmentId"`
		ZprTagNamespaceId   *string                    `json:"zprTagNamespaceId"`
		ZprTagNamespaceName *string                    `json:"zprTagNamespaceName"`
		Id                  *string                    `json:"id"`
		Name                *string                    `json:"name"`
		Description         *string                    `json:"description"`
		IsRetired           *bool                      `json:"isRetired"`
		TimeCreated         *common.SDKTime            `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Type = model.Type

	m.LifecycleState = model.LifecycleState

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

	m.ZprTagNamespaceId = model.ZprTagNamespaceId

	m.ZprTagNamespaceName = model.ZprTagNamespaceName

	m.Id = model.Id

	m.Name = model.Name

	m.Description = model.Description

	m.IsRetired = model.IsRetired

	m.TimeCreated = model.TimeCreated

	return
}

// ZprTagLifecycleStateEnum Enum with underlying type: string
type ZprTagLifecycleStateEnum string

// Set of constants representing the allowable values for ZprTagLifecycleStateEnum
const (
	ZprTagLifecycleStateActive   ZprTagLifecycleStateEnum = "ACTIVE"
	ZprTagLifecycleStateInactive ZprTagLifecycleStateEnum = "INACTIVE"
	ZprTagLifecycleStateDeleting ZprTagLifecycleStateEnum = "DELETING"
	ZprTagLifecycleStateDeleted  ZprTagLifecycleStateEnum = "DELETED"
)

var mappingZprTagLifecycleStateEnum = map[string]ZprTagLifecycleStateEnum{
	"ACTIVE":   ZprTagLifecycleStateActive,
	"INACTIVE": ZprTagLifecycleStateInactive,
	"DELETING": ZprTagLifecycleStateDeleting,
	"DELETED":  ZprTagLifecycleStateDeleted,
}

var mappingZprTagLifecycleStateEnumLowerCase = map[string]ZprTagLifecycleStateEnum{
	"active":   ZprTagLifecycleStateActive,
	"inactive": ZprTagLifecycleStateInactive,
	"deleting": ZprTagLifecycleStateDeleting,
	"deleted":  ZprTagLifecycleStateDeleted,
}

// GetZprTagLifecycleStateEnumValues Enumerates the set of values for ZprTagLifecycleStateEnum
func GetZprTagLifecycleStateEnumValues() []ZprTagLifecycleStateEnum {
	values := make([]ZprTagLifecycleStateEnum, 0)
	for _, v := range mappingZprTagLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetZprTagLifecycleStateEnumStringValues Enumerates the set of values in String for ZprTagLifecycleStateEnum
func GetZprTagLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingZprTagLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingZprTagLifecycleStateEnum(val string) (ZprTagLifecycleStateEnum, bool) {
	enum, ok := mappingZprTagLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
