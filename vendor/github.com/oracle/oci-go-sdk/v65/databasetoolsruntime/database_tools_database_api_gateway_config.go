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

// DatabaseToolsDatabaseApiGatewayConfig Simplified DatabaseToolsDatabaseApiGatewayConfig representation for splat usage only.
type DatabaseToolsDatabaseApiGatewayConfig struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools database API gateway config.
	Id *string `mandatory:"true" json:"id"`

	// The Database Tools database API gateway config type.
	Type DatabaseApiGatewayConfigTypeEnum `mandatory:"true" json:"type"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools database API gateway config.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Database Tools database API gateway config.
	LifecycleState DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the Database Tools database API gateway config was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Database Tools database API gateway config was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

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

func (m DatabaseToolsDatabaseApiGatewayConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsDatabaseApiGatewayConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseApiGatewayConfigTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDatabaseApiGatewayConfigTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum Enum with underlying type: string
type DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum
const (
	DatabaseToolsDatabaseApiGatewayConfigLifecycleStateActive  DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum = "ACTIVE"
	DatabaseToolsDatabaseApiGatewayConfigLifecycleStateDeleted DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum = "DELETED"
)

var mappingDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum = map[string]DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum{
	"ACTIVE":  DatabaseToolsDatabaseApiGatewayConfigLifecycleStateActive,
	"DELETED": DatabaseToolsDatabaseApiGatewayConfigLifecycleStateDeleted,
}

var mappingDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnumLowerCase = map[string]DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum{
	"active":  DatabaseToolsDatabaseApiGatewayConfigLifecycleStateActive,
	"deleted": DatabaseToolsDatabaseApiGatewayConfigLifecycleStateDeleted,
}

// GetDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnumValues Enumerates the set of values for DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum
func GetDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnumValues() []DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum {
	values := make([]DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnumStringValues Enumerates the set of values in String for DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum
func GetDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum(val string) (DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum, bool) {
	enum, ok := mappingDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
