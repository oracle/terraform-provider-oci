// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Zero Trust Packet Routing Control Plane API
//
// Use the Zero Trust Packet Routing Control Plane API to manage ZPR configuration and policy. See the Zero Trust Packet Routing (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/home.htm) documentation for more information.
//

package zpr

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Configuration The configuration details of ZPR in the root compartment (the root compartment is the tenancy).
type Configuration struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ZprConfiguration.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the tenancy into which ZPR will be onboarded.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The enabled or disabled status of ZPR in tenancy.
	ZprStatus ConfigurationZprStatusEnum `mandatory:"true" json:"zprStatus"`

	// The date and time that ZPR was onboarded to the tenancy, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time that ZPR was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of ZPR in the tenancy.
	LifecycleState ConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A message that describes the current state of ZPR in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
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
	if _, ok := GetMappingConfigurationZprStatusEnum(string(m.ZprStatus)); !ok && m.ZprStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ZprStatus: %s. Supported values are: %s.", m.ZprStatus, strings.Join(GetConfigurationZprStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigurationZprStatusEnum Enum with underlying type: string
type ConfigurationZprStatusEnum string

// Set of constants representing the allowable values for ConfigurationZprStatusEnum
const (
	ConfigurationZprStatusEnabled  ConfigurationZprStatusEnum = "ENABLED"
	ConfigurationZprStatusDisabled ConfigurationZprStatusEnum = "DISABLED"
)

var mappingConfigurationZprStatusEnum = map[string]ConfigurationZprStatusEnum{
	"ENABLED":  ConfigurationZprStatusEnabled,
	"DISABLED": ConfigurationZprStatusDisabled,
}

var mappingConfigurationZprStatusEnumLowerCase = map[string]ConfigurationZprStatusEnum{
	"enabled":  ConfigurationZprStatusEnabled,
	"disabled": ConfigurationZprStatusDisabled,
}

// GetConfigurationZprStatusEnumValues Enumerates the set of values for ConfigurationZprStatusEnum
func GetConfigurationZprStatusEnumValues() []ConfigurationZprStatusEnum {
	values := make([]ConfigurationZprStatusEnum, 0)
	for _, v := range mappingConfigurationZprStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationZprStatusEnumStringValues Enumerates the set of values in String for ConfigurationZprStatusEnum
func GetConfigurationZprStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingConfigurationZprStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationZprStatusEnum(val string) (ConfigurationZprStatusEnum, bool) {
	enum, ok := mappingConfigurationZprStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConfigurationLifecycleStateEnum Enum with underlying type: string
type ConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for ConfigurationLifecycleStateEnum
const (
	ConfigurationLifecycleStateActive   ConfigurationLifecycleStateEnum = "ACTIVE"
	ConfigurationLifecycleStateCreating ConfigurationLifecycleStateEnum = "CREATING"
	ConfigurationLifecycleStateUpdating ConfigurationLifecycleStateEnum = "UPDATING"
	ConfigurationLifecycleStateDeleting ConfigurationLifecycleStateEnum = "DELETING"
	ConfigurationLifecycleStateDeleted  ConfigurationLifecycleStateEnum = "DELETED"
	ConfigurationLifecycleStateFailed   ConfigurationLifecycleStateEnum = "FAILED"
)

var mappingConfigurationLifecycleStateEnum = map[string]ConfigurationLifecycleStateEnum{
	"ACTIVE":   ConfigurationLifecycleStateActive,
	"CREATING": ConfigurationLifecycleStateCreating,
	"UPDATING": ConfigurationLifecycleStateUpdating,
	"DELETING": ConfigurationLifecycleStateDeleting,
	"DELETED":  ConfigurationLifecycleStateDeleted,
	"FAILED":   ConfigurationLifecycleStateFailed,
}

var mappingConfigurationLifecycleStateEnumLowerCase = map[string]ConfigurationLifecycleStateEnum{
	"active":   ConfigurationLifecycleStateActive,
	"creating": ConfigurationLifecycleStateCreating,
	"updating": ConfigurationLifecycleStateUpdating,
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
		"CREATING",
		"UPDATING",
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
