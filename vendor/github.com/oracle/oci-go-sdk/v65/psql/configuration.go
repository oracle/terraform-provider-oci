// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// A description of the PGSQL Control Plane API
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Configuration Db system Postgresql Configuration
type Configuration struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Config display name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Config compartment identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time Configuration was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Configuration.
	LifecycleState ConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Compute Shape Name like VM.Standard3.Flex.
	Shape *string `mandatory:"true" json:"shape"`

	// CPU cpuCoreCount. Min value is 1. Max value depends on the shape.
	InstanceOcpuCount *int `mandatory:"true" json:"instanceOcpuCount"`

	// Memory Size in GB with 1GB increment. Min value matches the cpuCoreCount. Max value depends on the shape.
	InstanceMemorySizeInGBs *int `mandatory:"true" json:"instanceMemorySizeInGBs"`

	// Version of the Postgresql DB
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	ConfigurationDetails *ConfigurationDetails `mandatory:"true" json:"configurationDetails"`

	// Config description
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

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
