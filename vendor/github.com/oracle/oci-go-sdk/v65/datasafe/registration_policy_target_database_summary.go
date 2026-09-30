// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RegistrationPolicyTargetDatabaseSummary Summary of a discovered database resource and its Data Safe target registered via a registration policy.
type RegistrationPolicyTargetDatabaseSummary struct {

	// The ID of the discovered database resource (for example, a Database or Pluggable Database) that is part of discovery.
	DiscoveredResourceId *string `mandatory:"true" json:"discoveredResourceId"`

	// The type of the discovered database resource.
	DiscoveredResourceType RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum `mandatory:"true" json:"discoveredResourceType"`

	// The ID of the Data Safe Target Database associated with the discovered resource.
	TargetDatabaseId *string `mandatory:"true" json:"targetDatabaseId"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`
}

func (m RegistrationPolicyTargetDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RegistrationPolicyTargetDatabaseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum(string(m.DiscoveredResourceType)); !ok && m.DiscoveredResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoveredResourceType: %s. Supported values are: %s.", m.DiscoveredResourceType, strings.Join(GetRegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum Enum with underlying type: string
type RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum string

// Set of constants representing the allowable values for RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum
const (
	RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeDatabase          RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum = "DATABASE"
	RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypePluggableDatabase RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum = "PLUGGABLE_DATABASE"
)

var mappingRegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum = map[string]RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum{
	"DATABASE":           RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeDatabase,
	"PLUGGABLE_DATABASE": RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypePluggableDatabase,
}

var mappingRegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnumLowerCase = map[string]RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum{
	"database":           RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeDatabase,
	"pluggable_database": RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypePluggableDatabase,
}

// GetRegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnumValues Enumerates the set of values for RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum
func GetRegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnumValues() []RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum {
	values := make([]RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum, 0)
	for _, v := range mappingRegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnumStringValues Enumerates the set of values in String for RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum
func GetRegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"PLUGGABLE_DATABASE",
	}
}

// GetMappingRegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum(val string) (RegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnum, bool) {
	enum, ok := mappingRegistrationPolicyTargetDatabaseSummaryDiscoveredResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
