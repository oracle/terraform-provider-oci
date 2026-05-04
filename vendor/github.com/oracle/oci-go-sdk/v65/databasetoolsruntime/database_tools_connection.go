// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsConnection Simplified DatabaseToolsConnection representation for splat usage only
type DatabaseToolsConnection struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools connection.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the Database Tools connection.
	LifecycleState DatabaseToolsConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the Database Tools connection was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the DatabaseToolsConnection was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The Database Tools connection type.
	Type ConnectionTypeEnum `mandatory:"true" json:"type"`

	// A message describing the current state in more detail. For example, this message can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DatabaseToolsConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseToolsConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabaseToolsConnectionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConnectionTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetConnectionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsConnectionLifecycleStateEnum Enum with underlying type: string
type DatabaseToolsConnectionLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseToolsConnectionLifecycleStateEnum
const (
	DatabaseToolsConnectionLifecycleStateCreating DatabaseToolsConnectionLifecycleStateEnum = "CREATING"
	DatabaseToolsConnectionLifecycleStateUpdating DatabaseToolsConnectionLifecycleStateEnum = "UPDATING"
	DatabaseToolsConnectionLifecycleStateActive   DatabaseToolsConnectionLifecycleStateEnum = "ACTIVE"
	DatabaseToolsConnectionLifecycleStateInactive DatabaseToolsConnectionLifecycleStateEnum = "INACTIVE"
	DatabaseToolsConnectionLifecycleStateDeleting DatabaseToolsConnectionLifecycleStateEnum = "DELETING"
	DatabaseToolsConnectionLifecycleStateDeleted  DatabaseToolsConnectionLifecycleStateEnum = "DELETED"
	DatabaseToolsConnectionLifecycleStateFailed   DatabaseToolsConnectionLifecycleStateEnum = "FAILED"
)

var mappingDatabaseToolsConnectionLifecycleStateEnum = map[string]DatabaseToolsConnectionLifecycleStateEnum{
	"CREATING": DatabaseToolsConnectionLifecycleStateCreating,
	"UPDATING": DatabaseToolsConnectionLifecycleStateUpdating,
	"ACTIVE":   DatabaseToolsConnectionLifecycleStateActive,
	"INACTIVE": DatabaseToolsConnectionLifecycleStateInactive,
	"DELETING": DatabaseToolsConnectionLifecycleStateDeleting,
	"DELETED":  DatabaseToolsConnectionLifecycleStateDeleted,
	"FAILED":   DatabaseToolsConnectionLifecycleStateFailed,
}

var mappingDatabaseToolsConnectionLifecycleStateEnumLowerCase = map[string]DatabaseToolsConnectionLifecycleStateEnum{
	"creating": DatabaseToolsConnectionLifecycleStateCreating,
	"updating": DatabaseToolsConnectionLifecycleStateUpdating,
	"active":   DatabaseToolsConnectionLifecycleStateActive,
	"inactive": DatabaseToolsConnectionLifecycleStateInactive,
	"deleting": DatabaseToolsConnectionLifecycleStateDeleting,
	"deleted":  DatabaseToolsConnectionLifecycleStateDeleted,
	"failed":   DatabaseToolsConnectionLifecycleStateFailed,
}

// GetDatabaseToolsConnectionLifecycleStateEnumValues Enumerates the set of values for DatabaseToolsConnectionLifecycleStateEnum
func GetDatabaseToolsConnectionLifecycleStateEnumValues() []DatabaseToolsConnectionLifecycleStateEnum {
	values := make([]DatabaseToolsConnectionLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseToolsConnectionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsConnectionLifecycleStateEnumStringValues Enumerates the set of values in String for DatabaseToolsConnectionLifecycleStateEnum
func GetDatabaseToolsConnectionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDatabaseToolsConnectionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsConnectionLifecycleStateEnum(val string) (DatabaseToolsConnectionLifecycleStateEnum, bool) {
	enum, ok := mappingDatabaseToolsConnectionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
