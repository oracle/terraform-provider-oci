// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Configuration PostgreSQL configuration for a database system.
type Configuration struct {

	// A unique identifier for the configuration. Immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the configuration. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the configuration.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time that the configuration was created, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the configuration.
	LifecycleState ConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Version of the PostgreSQL database.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// The name of the shape for the configuration.
	// Example: `VM.Standard.E4.Flex`
	Shape *string `mandatory:"true" json:"shape"`

	// CPU core count.
	// It's value is set to 0 if configuration is for a flexible shape.
	InstanceOcpuCount *int `mandatory:"true" json:"instanceOcpuCount"`

	// Memory size in gigabytes with 1GB increment.
	// It's value is set to 0 if configuration is for a flexible shape.
	InstanceMemorySizeInGBs *int `mandatory:"true" json:"instanceMemorySizeInGBs"`

	ConfigurationDetails *ConfigurationDetails `mandatory:"true" json:"configurationDetails"`

	// A description for the configuration.
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The type of configuration. Either user-created or a default configuration.
	ConfigType ConfigurationConfigTypeEnum `mandatory:"false" json:"configType,omitempty"`

	// Whether the configuration supports flexible shapes.
	IsFlexible *bool `mandatory:"false" json:"isFlexible"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Configuration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Configuration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingConfigurationConfigTypeEnum(string(m.ConfigType)); !ok && m.ConfigType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigType: %s. Supported values are: %s.", m.ConfigType, strings.Join(GetConfigurationConfigTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigurationLifecycleStateEnum Enum with underlying type: string
type ConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for ConfigurationLifecycleStateEnum
const (
	ConfigurationLifecycleStateActive   ConfigurationLifecycleStateEnum = "ACTIVE"
	ConfigurationLifecycleStateDeleting ConfigurationLifecycleStateEnum = "DELETING"
	ConfigurationLifecycleStateDeleted  ConfigurationLifecycleStateEnum = "DELETED"
	ConfigurationLifecycleStateFailed   ConfigurationLifecycleStateEnum = "FAILED"
)

var mappingConfigurationLifecycleStateEnum = map[string]ConfigurationLifecycleStateEnum{
	"ACTIVE":   ConfigurationLifecycleStateActive,
	"DELETING": ConfigurationLifecycleStateDeleting,
	"DELETED":  ConfigurationLifecycleStateDeleted,
	"FAILED":   ConfigurationLifecycleStateFailed,
}

var mappingConfigurationLifecycleStateEnumLowerCase = map[string]ConfigurationLifecycleStateEnum{
	"active":   ConfigurationLifecycleStateActive,
	"deleting": ConfigurationLifecycleStateDeleting,
	"deleted":  ConfigurationLifecycleStateDeleted,
	"failed":   ConfigurationLifecycleStateFailed,
}

// GetConfigurationLifecycleStateEnumValues Enumerates the set of values for ConfigurationLifecycleStateEnum
func GetConfigurationLifecycleStateEnumValues() []ConfigurationLifecycleStateEnum {
	values := make([]ConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for ConfigurationLifecycleStateEnum
func GetConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationLifecycleStateEnum(val string) (ConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConfigurationConfigTypeEnum Enum with underlying type: string
type ConfigurationConfigTypeEnum string

// Set of constants representing the allowable values for ConfigurationConfigTypeEnum
const (
	ConfigurationConfigTypeDefault ConfigurationConfigTypeEnum = "DEFAULT"
	ConfigurationConfigTypeCustom  ConfigurationConfigTypeEnum = "CUSTOM"
	ConfigurationConfigTypeCopied  ConfigurationConfigTypeEnum = "COPIED"
)

var mappingConfigurationConfigTypeEnum = map[string]ConfigurationConfigTypeEnum{
	"DEFAULT": ConfigurationConfigTypeDefault,
	"CUSTOM":  ConfigurationConfigTypeCustom,
	"COPIED":  ConfigurationConfigTypeCopied,
}

var mappingConfigurationConfigTypeEnumLowerCase = map[string]ConfigurationConfigTypeEnum{
	"default": ConfigurationConfigTypeDefault,
	"custom":  ConfigurationConfigTypeCustom,
	"copied":  ConfigurationConfigTypeCopied,
}

// GetConfigurationConfigTypeEnumValues Enumerates the set of values for ConfigurationConfigTypeEnum
func GetConfigurationConfigTypeEnumValues() []ConfigurationConfigTypeEnum {
	values := make([]ConfigurationConfigTypeEnum, 0)
	for _, v := range mappingConfigurationConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationConfigTypeEnumStringValues Enumerates the set of values in String for ConfigurationConfigTypeEnum
func GetConfigurationConfigTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"CUSTOM",
		"COPIED",
	}
}

// GetMappingConfigurationConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationConfigTypeEnum(val string) (ConfigurationConfigTypeEnum, bool) {
	enum, ok := mappingConfigurationConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
