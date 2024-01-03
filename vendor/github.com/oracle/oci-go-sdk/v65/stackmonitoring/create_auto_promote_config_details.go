// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateAutoPromoteConfigDetails The details of an AUTO_PROMOTE configuration.
type CreateAutoPromoteConfigDetails struct {

	// Compartment in which the configuration is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// True if automatic promotion is enabled, false if it is not enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The display name of the configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The type of resource to configure for automatic promotion.
	ResourceType CreateAutoPromoteConfigDetailsResourceTypeEnum `mandatory:"true" json:"resourceType"`
}

// GetDisplayName returns DisplayName
func (m CreateAutoPromoteConfigDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateAutoPromoteConfigDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateAutoPromoteConfigDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateAutoPromoteConfigDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateAutoPromoteConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAutoPromoteConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateAutoPromoteConfigDetailsResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetCreateAutoPromoteConfigDetailsResourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateAutoPromoteConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAutoPromoteConfigDetails CreateAutoPromoteConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateAutoPromoteConfigDetails
	}{
		"AUTO_PROMOTE",
		(MarshalTypeCreateAutoPromoteConfigDetails)(m),
	}

	return json.Marshal(&s)
}

// CreateAutoPromoteConfigDetailsResourceTypeEnum Enum with underlying type: string
type CreateAutoPromoteConfigDetailsResourceTypeEnum string

// Set of constants representing the allowable values for CreateAutoPromoteConfigDetailsResourceTypeEnum
const (
	CreateAutoPromoteConfigDetailsResourceTypeHost CreateAutoPromoteConfigDetailsResourceTypeEnum = "HOST"
)

var mappingCreateAutoPromoteConfigDetailsResourceTypeEnum = map[string]CreateAutoPromoteConfigDetailsResourceTypeEnum{
	"HOST": CreateAutoPromoteConfigDetailsResourceTypeHost,
}

var mappingCreateAutoPromoteConfigDetailsResourceTypeEnumLowerCase = map[string]CreateAutoPromoteConfigDetailsResourceTypeEnum{
	"host": CreateAutoPromoteConfigDetailsResourceTypeHost,
}

// GetCreateAutoPromoteConfigDetailsResourceTypeEnumValues Enumerates the set of values for CreateAutoPromoteConfigDetailsResourceTypeEnum
func GetCreateAutoPromoteConfigDetailsResourceTypeEnumValues() []CreateAutoPromoteConfigDetailsResourceTypeEnum {
	values := make([]CreateAutoPromoteConfigDetailsResourceTypeEnum, 0)
	for _, v := range mappingCreateAutoPromoteConfigDetailsResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutoPromoteConfigDetailsResourceTypeEnumStringValues Enumerates the set of values in String for CreateAutoPromoteConfigDetailsResourceTypeEnum
func GetCreateAutoPromoteConfigDetailsResourceTypeEnumStringValues() []string {
	return []string{
		"HOST",
	}
}

// GetMappingCreateAutoPromoteConfigDetailsResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutoPromoteConfigDetailsResourceTypeEnum(val string) (CreateAutoPromoteConfigDetailsResourceTypeEnum, bool) {
	enum, ok := mappingCreateAutoPromoteConfigDetailsResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
