// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.cloud.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DefaultConfiguration Default configurations for PostgreSQL database systems.
type DefaultConfiguration struct {

	// A unique identifier for the configuration.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the configuration.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time that the configuration was created, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the configuration.
	LifecycleState DefaultConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The name of the shape for the configuration.
	// Example: `VM.Standard.E4.Flex`
	Shape *string `mandatory:"true" json:"shape"`

	// CPU core count. Minimum value is 1.
	InstanceOcpuCount *int `mandatory:"true" json:"instanceOcpuCount"`

	// Memory size in gigabytes with 1GB increment.
	InstanceMemorySizeInGBs *int `mandatory:"true" json:"instanceMemorySizeInGBs"`

	// Version of the PostgreSQL database.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	ConfigurationDetails *DefaultConfigurationDetails `mandatory:"true" json:"configurationDetails"`

	// A description for the configuration.
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m DefaultConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefaultConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDefaultConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDefaultConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DefaultConfigurationLifecycleStateEnum Enum with underlying type: string
type DefaultConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for DefaultConfigurationLifecycleStateEnum
const (
	DefaultConfigurationLifecycleStateActive   DefaultConfigurationLifecycleStateEnum = "ACTIVE"
	DefaultConfigurationLifecycleStateInactive DefaultConfigurationLifecycleStateEnum = "INACTIVE"
	DefaultConfigurationLifecycleStateDeleting DefaultConfigurationLifecycleStateEnum = "DELETING"
	DefaultConfigurationLifecycleStateDeleted  DefaultConfigurationLifecycleStateEnum = "DELETED"
	DefaultConfigurationLifecycleStateFailed   DefaultConfigurationLifecycleStateEnum = "FAILED"
)

var mappingDefaultConfigurationLifecycleStateEnum = map[string]DefaultConfigurationLifecycleStateEnum{
	"ACTIVE":   DefaultConfigurationLifecycleStateActive,
	"INACTIVE": DefaultConfigurationLifecycleStateInactive,
	"DELETING": DefaultConfigurationLifecycleStateDeleting,
	"DELETED":  DefaultConfigurationLifecycleStateDeleted,
	"FAILED":   DefaultConfigurationLifecycleStateFailed,
}

var mappingDefaultConfigurationLifecycleStateEnumLowerCase = map[string]DefaultConfigurationLifecycleStateEnum{
	"active":   DefaultConfigurationLifecycleStateActive,
	"inactive": DefaultConfigurationLifecycleStateInactive,
	"deleting": DefaultConfigurationLifecycleStateDeleting,
	"deleted":  DefaultConfigurationLifecycleStateDeleted,
	"failed":   DefaultConfigurationLifecycleStateFailed,
}

// GetDefaultConfigurationLifecycleStateEnumValues Enumerates the set of values for DefaultConfigurationLifecycleStateEnum
func GetDefaultConfigurationLifecycleStateEnumValues() []DefaultConfigurationLifecycleStateEnum {
	values := make([]DefaultConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingDefaultConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDefaultConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for DefaultConfigurationLifecycleStateEnum
func GetDefaultConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDefaultConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDefaultConfigurationLifecycleStateEnum(val string) (DefaultConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingDefaultConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
