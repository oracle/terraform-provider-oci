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

// AutoPromoteConfigDetails A configuration of the AUTO_PROMOTE type, consists of a resource type and a boolean value
// that determines if this resource needs to be automatically promoted/discovered.
// For example, when a Management Agent registration event occurs and if isEnabled is TRUE for
// a HOST resource type, a HOST resource will be automatically discovered using that Management Agent.
type AutoPromoteConfigDetails struct {

	// The Unique Oracle ID (OCID) that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the configuration.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// True if automatic promotion is enabled, false if it is not enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the configuration was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Config was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The type of resource to configure for automatic promotion.
	ResourceType AutoPromoteConfigDetailsResourceTypeEnum `mandatory:"true" json:"resourceType"`

	// The current state of the configuration.
	LifecycleState ConfigLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m AutoPromoteConfigDetails) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m AutoPromoteConfigDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m AutoPromoteConfigDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m AutoPromoteConfigDetails) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m AutoPromoteConfigDetails) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m AutoPromoteConfigDetails) GetLifecycleState() ConfigLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m AutoPromoteConfigDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m AutoPromoteConfigDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m AutoPromoteConfigDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m AutoPromoteConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoPromoteConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutoPromoteConfigDetailsResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetAutoPromoteConfigDetailsResourceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConfigLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AutoPromoteConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAutoPromoteConfigDetails AutoPromoteConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeAutoPromoteConfigDetails
	}{
		"AUTO_PROMOTE",
		(MarshalTypeAutoPromoteConfigDetails)(m),
	}

	return json.Marshal(&s)
}

// AutoPromoteConfigDetailsResourceTypeEnum Enum with underlying type: string
type AutoPromoteConfigDetailsResourceTypeEnum string

// Set of constants representing the allowable values for AutoPromoteConfigDetailsResourceTypeEnum
const (
	AutoPromoteConfigDetailsResourceTypeHost AutoPromoteConfigDetailsResourceTypeEnum = "HOST"
)

var mappingAutoPromoteConfigDetailsResourceTypeEnum = map[string]AutoPromoteConfigDetailsResourceTypeEnum{
	"HOST": AutoPromoteConfigDetailsResourceTypeHost,
}

var mappingAutoPromoteConfigDetailsResourceTypeEnumLowerCase = map[string]AutoPromoteConfigDetailsResourceTypeEnum{
	"host": AutoPromoteConfigDetailsResourceTypeHost,
}

// GetAutoPromoteConfigDetailsResourceTypeEnumValues Enumerates the set of values for AutoPromoteConfigDetailsResourceTypeEnum
func GetAutoPromoteConfigDetailsResourceTypeEnumValues() []AutoPromoteConfigDetailsResourceTypeEnum {
	values := make([]AutoPromoteConfigDetailsResourceTypeEnum, 0)
	for _, v := range mappingAutoPromoteConfigDetailsResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoPromoteConfigDetailsResourceTypeEnumStringValues Enumerates the set of values in String for AutoPromoteConfigDetailsResourceTypeEnum
func GetAutoPromoteConfigDetailsResourceTypeEnumStringValues() []string {
	return []string{
		"HOST",
	}
}

// GetMappingAutoPromoteConfigDetailsResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoPromoteConfigDetailsResourceTypeEnum(val string) (AutoPromoteConfigDetailsResourceTypeEnum, bool) {
	enum, ok := mappingAutoPromoteConfigDetailsResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
