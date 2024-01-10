// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Environment Description of the source environment.
type Environment struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Environment identifier, which can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time when the source environment was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the source environment.
	LifecycleState EnvironmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The time when the source environment was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Environment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Environment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEnvironmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEnvironmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EnvironmentLifecycleStateEnum Enum with underlying type: string
type EnvironmentLifecycleStateEnum string

// Set of constants representing the allowable values for EnvironmentLifecycleStateEnum
const (
	EnvironmentLifecycleStateCreating EnvironmentLifecycleStateEnum = "CREATING"
	EnvironmentLifecycleStateUpdating EnvironmentLifecycleStateEnum = "UPDATING"
	EnvironmentLifecycleStateActive   EnvironmentLifecycleStateEnum = "ACTIVE"
	EnvironmentLifecycleStateDeleting EnvironmentLifecycleStateEnum = "DELETING"
	EnvironmentLifecycleStateDeleted  EnvironmentLifecycleStateEnum = "DELETED"
	EnvironmentLifecycleStateFailed   EnvironmentLifecycleStateEnum = "FAILED"
)

var mappingEnvironmentLifecycleStateEnum = map[string]EnvironmentLifecycleStateEnum{
	"CREATING": EnvironmentLifecycleStateCreating,
	"UPDATING": EnvironmentLifecycleStateUpdating,
	"ACTIVE":   EnvironmentLifecycleStateActive,
	"DELETING": EnvironmentLifecycleStateDeleting,
	"DELETED":  EnvironmentLifecycleStateDeleted,
	"FAILED":   EnvironmentLifecycleStateFailed,
}

var mappingEnvironmentLifecycleStateEnumLowerCase = map[string]EnvironmentLifecycleStateEnum{
	"creating": EnvironmentLifecycleStateCreating,
	"updating": EnvironmentLifecycleStateUpdating,
	"active":   EnvironmentLifecycleStateActive,
	"deleting": EnvironmentLifecycleStateDeleting,
	"deleted":  EnvironmentLifecycleStateDeleted,
	"failed":   EnvironmentLifecycleStateFailed,
}

// GetEnvironmentLifecycleStateEnumValues Enumerates the set of values for EnvironmentLifecycleStateEnum
func GetEnvironmentLifecycleStateEnumValues() []EnvironmentLifecycleStateEnum {
	values := make([]EnvironmentLifecycleStateEnum, 0)
	for _, v := range mappingEnvironmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEnvironmentLifecycleStateEnumStringValues Enumerates the set of values in String for EnvironmentLifecycleStateEnum
func GetEnvironmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingEnvironmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnvironmentLifecycleStateEnum(val string) (EnvironmentLifecycleStateEnum, bool) {
	enum, ok := mappingEnvironmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
