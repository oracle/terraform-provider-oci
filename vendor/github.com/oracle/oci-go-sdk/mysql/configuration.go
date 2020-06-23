// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Configuration The set of MySQL variables to be used when deploying a MySQL Database Service DB System.
type Configuration struct {

	// The OCID of the Configuration.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the Compartment the Configuration exists in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the associated Shape.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The Configuration type, DEFAULT or CUSTOM.
	Type ConfigurationTypeEnum `mandatory:"true" json:"type"`

	// The date and time the Configuration was created, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the Configuration was last updated, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the Configuration.
	LifecycleState ConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	Variables *ConfigurationVariables `mandatory:"true" json:"variables"`

	// User-provided data about the Configuration.
	Description *string `mandatory:"false" json:"description"`

	// The display name of the Configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the Configuration from which this Configuration is
	// "derived". This is entirely a metadata relationship. There is no
	// relation between the values in this Configuration and its parent.
	ParentConfigurationId *string `mandatory:"false" json:"parentConfigurationId"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Configuration) String() string {
	return common.PointerString(m)
}

// ConfigurationTypeEnum Enum with underlying type: string
type ConfigurationTypeEnum string

// Set of constants representing the allowable values for ConfigurationTypeEnum
const (
	ConfigurationTypeDefault ConfigurationTypeEnum = "DEFAULT"
	ConfigurationTypeCustom  ConfigurationTypeEnum = "CUSTOM"
)

var mappingConfigurationType = map[string]ConfigurationTypeEnum{
	"DEFAULT": ConfigurationTypeDefault,
	"CUSTOM":  ConfigurationTypeCustom,
}

// GetConfigurationTypeEnumValues Enumerates the set of values for ConfigurationTypeEnum
func GetConfigurationTypeEnumValues() []ConfigurationTypeEnum {
	values := make([]ConfigurationTypeEnum, 0)
	for _, v := range mappingConfigurationType {
		values = append(values, v)
	}
	return values
}

// ConfigurationLifecycleStateEnum Enum with underlying type: string
type ConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for ConfigurationLifecycleStateEnum
const (
	ConfigurationLifecycleStateActive  ConfigurationLifecycleStateEnum = "ACTIVE"
	ConfigurationLifecycleStateDeleted ConfigurationLifecycleStateEnum = "DELETED"
)

var mappingConfigurationLifecycleState = map[string]ConfigurationLifecycleStateEnum{
	"ACTIVE":  ConfigurationLifecycleStateActive,
	"DELETED": ConfigurationLifecycleStateDeleted,
}

// GetConfigurationLifecycleStateEnumValues Enumerates the set of values for ConfigurationLifecycleStateEnum
func GetConfigurationLifecycleStateEnumValues() []ConfigurationLifecycleStateEnum {
	values := make([]ConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingConfigurationLifecycleState {
		values = append(values, v)
	}
	return values
}
