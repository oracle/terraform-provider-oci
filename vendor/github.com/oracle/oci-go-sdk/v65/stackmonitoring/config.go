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

// Config A configuration item that, for example defines whether resources of a specific type
// should be discovered automatically.
// In this case, the 'configType' is set to 'AUTO_PROMOTE' and additional fields like
// 'resourceType' and 'isEnabled' determine if such resources are to be discovered
// automatically (also referred to as 'Automatic Promotion').
type Config interface {

	// The Unique Oracle ID (OCID) that is immutable on creation.
	GetId() *string

	// The OCID of the compartment containing the configuration.
	GetCompartmentId() *string

	// The current state of the configuration.
	GetLifecycleState() ConfigLifecycleStateEnum

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The time the configuration was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the Config was updated.
	GetTimeUpdated() *common.SDKTime

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type config struct {
	JsonData       []byte
	DisplayName    *string                           `mandatory:"false" json:"displayName"`
	TimeCreated    *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated    *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	FreeformTags   map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags     map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id             *string                           `mandatory:"true" json:"id"`
	CompartmentId  *string                           `mandatory:"true" json:"compartmentId"`
	LifecycleState ConfigLifecycleStateEnum          `mandatory:"true" json:"lifecycleState"`
	ConfigType     string                            `json:"configType"`
}

// UnmarshalJSON unmarshals json
func (m *config) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfig config
	s := struct {
		Model Unmarshalerconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.LifecycleState = s.Model.LifecycleState
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.ConfigType = s.Model.ConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *config) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigType {
	case "LICENSE_ENTERPRISE_EXTENSIBILITY":
		mm := LicenseEnterpriseExtensibilityConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LICENSE_AUTO_ASSIGN":
		mm := LicenseAutoAssignConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AUTO_PROMOTE":
		mm := AutoPromoteConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Config: %s.", m.ConfigType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m config) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m config) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m config) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m config) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m config) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m config) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m config) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m config) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetLifecycleState returns LifecycleState
func (m config) GetLifecycleState() ConfigLifecycleStateEnum {
	return m.LifecycleState
}

func (m config) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m config) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConfigLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigLifecycleStateEnum Enum with underlying type: string
type ConfigLifecycleStateEnum string

// Set of constants representing the allowable values for ConfigLifecycleStateEnum
const (
	ConfigLifecycleStateCreating ConfigLifecycleStateEnum = "CREATING"
	ConfigLifecycleStateUpdating ConfigLifecycleStateEnum = "UPDATING"
	ConfigLifecycleStateActive   ConfigLifecycleStateEnum = "ACTIVE"
	ConfigLifecycleStateDeleting ConfigLifecycleStateEnum = "DELETING"
	ConfigLifecycleStateDeleted  ConfigLifecycleStateEnum = "DELETED"
	ConfigLifecycleStateFailed   ConfigLifecycleStateEnum = "FAILED"
)

var mappingConfigLifecycleStateEnum = map[string]ConfigLifecycleStateEnum{
	"CREATING": ConfigLifecycleStateCreating,
	"UPDATING": ConfigLifecycleStateUpdating,
	"ACTIVE":   ConfigLifecycleStateActive,
	"DELETING": ConfigLifecycleStateDeleting,
	"DELETED":  ConfigLifecycleStateDeleted,
	"FAILED":   ConfigLifecycleStateFailed,
}

var mappingConfigLifecycleStateEnumLowerCase = map[string]ConfigLifecycleStateEnum{
	"creating": ConfigLifecycleStateCreating,
	"updating": ConfigLifecycleStateUpdating,
	"active":   ConfigLifecycleStateActive,
	"deleting": ConfigLifecycleStateDeleting,
	"deleted":  ConfigLifecycleStateDeleted,
	"failed":   ConfigLifecycleStateFailed,
}

// GetConfigLifecycleStateEnumValues Enumerates the set of values for ConfigLifecycleStateEnum
func GetConfigLifecycleStateEnumValues() []ConfigLifecycleStateEnum {
	values := make([]ConfigLifecycleStateEnum, 0)
	for _, v := range mappingConfigLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigLifecycleStateEnumStringValues Enumerates the set of values in String for ConfigLifecycleStateEnum
func GetConfigLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingConfigLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigLifecycleStateEnum(val string) (ConfigLifecycleStateEnum, bool) {
	enum, ok := mappingConfigLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConfigConfigTypeEnum Enum with underlying type: string
type ConfigConfigTypeEnum string

// Set of constants representing the allowable values for ConfigConfigTypeEnum
const (
	ConfigConfigTypeAutoPromote                    ConfigConfigTypeEnum = "AUTO_PROMOTE"
	ConfigConfigTypeLicenseAutoAssign              ConfigConfigTypeEnum = "LICENSE_AUTO_ASSIGN"
	ConfigConfigTypeLicenseEnterpriseExtensibility ConfigConfigTypeEnum = "LICENSE_ENTERPRISE_EXTENSIBILITY"
)

var mappingConfigConfigTypeEnum = map[string]ConfigConfigTypeEnum{
	"AUTO_PROMOTE":                     ConfigConfigTypeAutoPromote,
	"LICENSE_AUTO_ASSIGN":              ConfigConfigTypeLicenseAutoAssign,
	"LICENSE_ENTERPRISE_EXTENSIBILITY": ConfigConfigTypeLicenseEnterpriseExtensibility,
}

var mappingConfigConfigTypeEnumLowerCase = map[string]ConfigConfigTypeEnum{
	"auto_promote":                     ConfigConfigTypeAutoPromote,
	"license_auto_assign":              ConfigConfigTypeLicenseAutoAssign,
	"license_enterprise_extensibility": ConfigConfigTypeLicenseEnterpriseExtensibility,
}

// GetConfigConfigTypeEnumValues Enumerates the set of values for ConfigConfigTypeEnum
func GetConfigConfigTypeEnumValues() []ConfigConfigTypeEnum {
	values := make([]ConfigConfigTypeEnum, 0)
	for _, v := range mappingConfigConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigConfigTypeEnumStringValues Enumerates the set of values in String for ConfigConfigTypeEnum
func GetConfigConfigTypeEnumStringValues() []string {
	return []string{
		"AUTO_PROMOTE",
		"LICENSE_AUTO_ASSIGN",
		"LICENSE_ENTERPRISE_EXTENSIBILITY",
	}
}

// GetMappingConfigConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigConfigTypeEnum(val string) (ConfigConfigTypeEnum, bool) {
	enum, ok := mappingConfigConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
