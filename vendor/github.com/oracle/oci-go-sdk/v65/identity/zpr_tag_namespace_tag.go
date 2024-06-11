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

// ZprTagNamespaceTag A ZPR tag definition that belongs to a specific ZPR tag namespace. ZPR tags must be created in your tenancy before
// you can apply them to resources.
// For more information, see Managing Tags and ZPR tag Namespaces (https://docs.cloud.oracle.com/Content/Identity/Concepts/taggingoverview.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values
// using the API.
type ZprTagNamespaceTag struct {

	// The OCID of the compartment that contains the ZPR tag definition.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the ZPR tag namespace that contains the ZPR tag definition.
	ZprTagNamespaceId *string `mandatory:"true" json:"zprTagNamespaceId"`

	// The name of the ZPR tag namespace that contains the ZPR tag definition.
	ZprTagNamespaceName *string `mandatory:"true" json:"zprTagNamespaceName"`

	// The OCID of the ZPR tag definition.
	Id *string `mandatory:"true" json:"id"`

	// The name assigned to the ZPR tag during creation. This is the ZPR tag key definition.
	// The name must be unique within the ZPR tag namespace and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the ZPR tag.
	Description *string `mandatory:"true" json:"description"`

	// Indicates whether the ZPR tag is retired.
	// See Retiring Key Definitions and Namespace Definitions (https://docs.cloud.oracle.com/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm#retiringkeys).
	IsRetired *bool `mandatory:"true" json:"isRetired"`

	// Date and time the ZPR tag was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The data type of the ZPR tag.
	Type *string `mandatory:"false" json:"type"`

	// The ZPR tag's current state. After creating a ZPR tag, make sure its `lifecycleState` is ACTIVE before using it. After retiring a ZPR tag, make sure its `lifecycleState` is INACTIVE before using it. If you delete a ZPR tag, you cannot delete another ZPR tag until the deleted tag's `lifecycleState` changes from DELETING to DELETED.
	LifecycleState ZprTagNamespaceTagLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	Validator BaseTagDefinitionValidator `mandatory:"false" json:"validator"`
}

func (m ZprTagNamespaceTag) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ZprTagNamespaceTag) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingZprTagNamespaceTagLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetZprTagNamespaceTagLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ZprTagNamespaceTag) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Type                *string                              `json:"type"`
		LifecycleState      ZprTagNamespaceTagLifecycleStateEnum `json:"lifecycleState"`
		Validator           basetagdefinitionvalidator           `json:"validator"`
		CompartmentId       *string                              `json:"compartmentId"`
		ZprTagNamespaceId   *string                              `json:"zprTagNamespaceId"`
		ZprTagNamespaceName *string                              `json:"zprTagNamespaceName"`
		Id                  *string                              `json:"id"`
		Name                *string                              `json:"name"`
		Description         *string                              `json:"description"`
		IsRetired           *bool                                `json:"isRetired"`
		TimeCreated         *common.SDKTime                      `json:"timeCreated"`
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

// ZprTagNamespaceTagLifecycleStateEnum Enum with underlying type: string
type ZprTagNamespaceTagLifecycleStateEnum string

// Set of constants representing the allowable values for ZprTagNamespaceTagLifecycleStateEnum
const (
	ZprTagNamespaceTagLifecycleStateActive   ZprTagNamespaceTagLifecycleStateEnum = "ACTIVE"
	ZprTagNamespaceTagLifecycleStateInactive ZprTagNamespaceTagLifecycleStateEnum = "INACTIVE"
	ZprTagNamespaceTagLifecycleStateDeleting ZprTagNamespaceTagLifecycleStateEnum = "DELETING"
	ZprTagNamespaceTagLifecycleStateDeleted  ZprTagNamespaceTagLifecycleStateEnum = "DELETED"
)

var mappingZprTagNamespaceTagLifecycleStateEnum = map[string]ZprTagNamespaceTagLifecycleStateEnum{
	"ACTIVE":   ZprTagNamespaceTagLifecycleStateActive,
	"INACTIVE": ZprTagNamespaceTagLifecycleStateInactive,
	"DELETING": ZprTagNamespaceTagLifecycleStateDeleting,
	"DELETED":  ZprTagNamespaceTagLifecycleStateDeleted,
}

var mappingZprTagNamespaceTagLifecycleStateEnumLowerCase = map[string]ZprTagNamespaceTagLifecycleStateEnum{
	"active":   ZprTagNamespaceTagLifecycleStateActive,
	"inactive": ZprTagNamespaceTagLifecycleStateInactive,
	"deleting": ZprTagNamespaceTagLifecycleStateDeleting,
	"deleted":  ZprTagNamespaceTagLifecycleStateDeleted,
}

// GetZprTagNamespaceTagLifecycleStateEnumValues Enumerates the set of values for ZprTagNamespaceTagLifecycleStateEnum
func GetZprTagNamespaceTagLifecycleStateEnumValues() []ZprTagNamespaceTagLifecycleStateEnum {
	values := make([]ZprTagNamespaceTagLifecycleStateEnum, 0)
	for _, v := range mappingZprTagNamespaceTagLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetZprTagNamespaceTagLifecycleStateEnumStringValues Enumerates the set of values in String for ZprTagNamespaceTagLifecycleStateEnum
func GetZprTagNamespaceTagLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingZprTagNamespaceTagLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingZprTagNamespaceTagLifecycleStateEnum(val string) (ZprTagNamespaceTagLifecycleStateEnum, bool) {
	enum, ok := mappingZprTagNamespaceTagLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
