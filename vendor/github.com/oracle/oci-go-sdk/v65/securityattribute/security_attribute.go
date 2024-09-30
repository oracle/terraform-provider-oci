// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Security Attribute API
//
// Use the Security Attributes API to manage security attributes and security attribute namespaces. For more information, see the documentation for Security Attributes (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attributes.htm) and Security Attribute Nampespaces (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm).
//

package securityattribute

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecurityAttribute A security attribute that belongs to a specific security attribute namespace. Security attributes must be created in a tenancy before
// a user can apply them to resources.
// For more information, see Managing Security Attributes (https://docs.cloud.oracle.com/Content/zero-trust-packet-routing/managing-security-attributes.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values
// using the API.
type SecurityAttribute struct {

	// The OCID of the compartment that contains the security attribute definition.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the security attribute namespace that contains the security attribute definition.
	SecurityAttributeNamespaceId *string `mandatory:"true" json:"securityAttributeNamespaceId"`

	// The name of the security attribute namespace that contains the security attribute.
	SecurityAttributeNamespaceName *string `mandatory:"true" json:"securityAttributeNamespaceName"`

	// The OCID of the security attribute definition.
	Id *string `mandatory:"true" json:"id"`

	// The name assigned to the security attribute during creation. This is the security attribute key.
	// The name must be unique within the security attribute namespace and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description of the security attribute.
	Description *string `mandatory:"true" json:"description"`

	// Indicates whether the security attribute is retired.
	// See Managing Security Attribute Namespaces (https://docs.cloud.oracle.com/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm).
	IsRetired *bool `mandatory:"true" json:"isRetired"`

	// Date and time the security attribute was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The data type of the security attribute.
	Type *string `mandatory:"false" json:"type"`

	// The security attribute's current state. After creating a security attribute, make sure its `lifecycleState` is ACTIVE before using it. After retiring a security attribute, make sure its `lifecycleState` is INACTIVE before using it. If you delete a security attribute, you cannot delete another security attribute until the deleted tag's `lifecycleState` changes from DELETING to DELETED.
	LifecycleState SecurityAttributeLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	Validator BaseSecurityAttributeValidator `mandatory:"false" json:"validator"`
}

func (m SecurityAttribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityAttribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSecurityAttributeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSecurityAttributeLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SecurityAttribute) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Type                           *string                             `json:"type"`
		LifecycleState                 SecurityAttributeLifecycleStateEnum `json:"lifecycleState"`
		Validator                      basesecurityattributevalidator      `json:"validator"`
		CompartmentId                  *string                             `json:"compartmentId"`
		SecurityAttributeNamespaceId   *string                             `json:"securityAttributeNamespaceId"`
		SecurityAttributeNamespaceName *string                             `json:"securityAttributeNamespaceName"`
		Id                             *string                             `json:"id"`
		Name                           *string                             `json:"name"`
		Description                    *string                             `json:"description"`
		IsRetired                      *bool                               `json:"isRetired"`
		TimeCreated                    *common.SDKTime                     `json:"timeCreated"`
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
		m.Validator = nn.(BaseSecurityAttributeValidator)
	} else {
		m.Validator = nil
	}

	m.CompartmentId = model.CompartmentId

	m.SecurityAttributeNamespaceId = model.SecurityAttributeNamespaceId

	m.SecurityAttributeNamespaceName = model.SecurityAttributeNamespaceName

	m.Id = model.Id

	m.Name = model.Name

	m.Description = model.Description

	m.IsRetired = model.IsRetired

	m.TimeCreated = model.TimeCreated

	return
}

// SecurityAttributeLifecycleStateEnum Enum with underlying type: string
type SecurityAttributeLifecycleStateEnum string

// Set of constants representing the allowable values for SecurityAttributeLifecycleStateEnum
const (
	SecurityAttributeLifecycleStateActive   SecurityAttributeLifecycleStateEnum = "ACTIVE"
	SecurityAttributeLifecycleStateInactive SecurityAttributeLifecycleStateEnum = "INACTIVE"
	SecurityAttributeLifecycleStateDeleting SecurityAttributeLifecycleStateEnum = "DELETING"
	SecurityAttributeLifecycleStateDeleted  SecurityAttributeLifecycleStateEnum = "DELETED"
)

var mappingSecurityAttributeLifecycleStateEnum = map[string]SecurityAttributeLifecycleStateEnum{
	"ACTIVE":   SecurityAttributeLifecycleStateActive,
	"INACTIVE": SecurityAttributeLifecycleStateInactive,
	"DELETING": SecurityAttributeLifecycleStateDeleting,
	"DELETED":  SecurityAttributeLifecycleStateDeleted,
}

var mappingSecurityAttributeLifecycleStateEnumLowerCase = map[string]SecurityAttributeLifecycleStateEnum{
	"active":   SecurityAttributeLifecycleStateActive,
	"inactive": SecurityAttributeLifecycleStateInactive,
	"deleting": SecurityAttributeLifecycleStateDeleting,
	"deleted":  SecurityAttributeLifecycleStateDeleted,
}

// GetSecurityAttributeLifecycleStateEnumValues Enumerates the set of values for SecurityAttributeLifecycleStateEnum
func GetSecurityAttributeLifecycleStateEnumValues() []SecurityAttributeLifecycleStateEnum {
	values := make([]SecurityAttributeLifecycleStateEnum, 0)
	for _, v := range mappingSecurityAttributeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityAttributeLifecycleStateEnumStringValues Enumerates the set of values in String for SecurityAttributeLifecycleStateEnum
func GetSecurityAttributeLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingSecurityAttributeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityAttributeLifecycleStateEnum(val string) (SecurityAttributeLifecycleStateEnum, bool) {
	enum, ok := mappingSecurityAttributeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
