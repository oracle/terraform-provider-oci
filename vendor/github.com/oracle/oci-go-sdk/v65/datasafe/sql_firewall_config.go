// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// SqlFirewallConfig The SQL firewall related configurations.
type SqlFirewallConfig struct {

	// Specifies if the firewall is enabled or disabled on the target database.
	Status SqlFirewallConfigStatusEnum `mandatory:"true" json:"status"`

	// Specifies whether Data Safe should automatically purge the violation logs
	// from the database after collecting the violation logs and persisting on Data Safe.
	ViolationLogAutoPurge SqlFirewallConfigViolationLogAutoPurgeEnum `mandatory:"true" json:"violationLogAutoPurge"`

	// Specifies whether the firewall should include or exclude the database internal job activities.
	ExcludeJob SqlFirewallConfigExcludeJobEnum `mandatory:"false" json:"excludeJob,omitempty"`

	// The most recent time when the firewall status is updated, in the format defined by RFC3339.
	TimeStatusUpdated *common.SDKTime `mandatory:"false" json:"timeStatusUpdated"`
}

func (m SqlFirewallConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlFirewallConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlFirewallConfigStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSqlFirewallConfigStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlFirewallConfigViolationLogAutoPurgeEnum(string(m.ViolationLogAutoPurge)); !ok && m.ViolationLogAutoPurge != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ViolationLogAutoPurge: %s. Supported values are: %s.", m.ViolationLogAutoPurge, strings.Join(GetSqlFirewallConfigViolationLogAutoPurgeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSqlFirewallConfigExcludeJobEnum(string(m.ExcludeJob)); !ok && m.ExcludeJob != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExcludeJob: %s. Supported values are: %s.", m.ExcludeJob, strings.Join(GetSqlFirewallConfigExcludeJobEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlFirewallConfigStatusEnum Enum with underlying type: string
type SqlFirewallConfigStatusEnum string

// Set of constants representing the allowable values for SqlFirewallConfigStatusEnum
const (
	SqlFirewallConfigStatusEnabled  SqlFirewallConfigStatusEnum = "ENABLED"
	SqlFirewallConfigStatusDisabled SqlFirewallConfigStatusEnum = "DISABLED"
)

var mappingSqlFirewallConfigStatusEnum = map[string]SqlFirewallConfigStatusEnum{
	"ENABLED":  SqlFirewallConfigStatusEnabled,
	"DISABLED": SqlFirewallConfigStatusDisabled,
}

var mappingSqlFirewallConfigStatusEnumLowerCase = map[string]SqlFirewallConfigStatusEnum{
	"enabled":  SqlFirewallConfigStatusEnabled,
	"disabled": SqlFirewallConfigStatusDisabled,
}

// GetSqlFirewallConfigStatusEnumValues Enumerates the set of values for SqlFirewallConfigStatusEnum
func GetSqlFirewallConfigStatusEnumValues() []SqlFirewallConfigStatusEnum {
	values := make([]SqlFirewallConfigStatusEnum, 0)
	for _, v := range mappingSqlFirewallConfigStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallConfigStatusEnumStringValues Enumerates the set of values in String for SqlFirewallConfigStatusEnum
func GetSqlFirewallConfigStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingSqlFirewallConfigStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallConfigStatusEnum(val string) (SqlFirewallConfigStatusEnum, bool) {
	enum, ok := mappingSqlFirewallConfigStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlFirewallConfigViolationLogAutoPurgeEnum Enum with underlying type: string
type SqlFirewallConfigViolationLogAutoPurgeEnum string

// Set of constants representing the allowable values for SqlFirewallConfigViolationLogAutoPurgeEnum
const (
	SqlFirewallConfigViolationLogAutoPurgeEnabled  SqlFirewallConfigViolationLogAutoPurgeEnum = "ENABLED"
	SqlFirewallConfigViolationLogAutoPurgeDisabled SqlFirewallConfigViolationLogAutoPurgeEnum = "DISABLED"
)

var mappingSqlFirewallConfigViolationLogAutoPurgeEnum = map[string]SqlFirewallConfigViolationLogAutoPurgeEnum{
	"ENABLED":  SqlFirewallConfigViolationLogAutoPurgeEnabled,
	"DISABLED": SqlFirewallConfigViolationLogAutoPurgeDisabled,
}

var mappingSqlFirewallConfigViolationLogAutoPurgeEnumLowerCase = map[string]SqlFirewallConfigViolationLogAutoPurgeEnum{
	"enabled":  SqlFirewallConfigViolationLogAutoPurgeEnabled,
	"disabled": SqlFirewallConfigViolationLogAutoPurgeDisabled,
}

// GetSqlFirewallConfigViolationLogAutoPurgeEnumValues Enumerates the set of values for SqlFirewallConfigViolationLogAutoPurgeEnum
func GetSqlFirewallConfigViolationLogAutoPurgeEnumValues() []SqlFirewallConfigViolationLogAutoPurgeEnum {
	values := make([]SqlFirewallConfigViolationLogAutoPurgeEnum, 0)
	for _, v := range mappingSqlFirewallConfigViolationLogAutoPurgeEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallConfigViolationLogAutoPurgeEnumStringValues Enumerates the set of values in String for SqlFirewallConfigViolationLogAutoPurgeEnum
func GetSqlFirewallConfigViolationLogAutoPurgeEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingSqlFirewallConfigViolationLogAutoPurgeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallConfigViolationLogAutoPurgeEnum(val string) (SqlFirewallConfigViolationLogAutoPurgeEnum, bool) {
	enum, ok := mappingSqlFirewallConfigViolationLogAutoPurgeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlFirewallConfigExcludeJobEnum Enum with underlying type: string
type SqlFirewallConfigExcludeJobEnum string

// Set of constants representing the allowable values for SqlFirewallConfigExcludeJobEnum
const (
	SqlFirewallConfigExcludeJobExcluded SqlFirewallConfigExcludeJobEnum = "EXCLUDED"
	SqlFirewallConfigExcludeJobIncluded SqlFirewallConfigExcludeJobEnum = "INCLUDED"
)

var mappingSqlFirewallConfigExcludeJobEnum = map[string]SqlFirewallConfigExcludeJobEnum{
	"EXCLUDED": SqlFirewallConfigExcludeJobExcluded,
	"INCLUDED": SqlFirewallConfigExcludeJobIncluded,
}

var mappingSqlFirewallConfigExcludeJobEnumLowerCase = map[string]SqlFirewallConfigExcludeJobEnum{
	"excluded": SqlFirewallConfigExcludeJobExcluded,
	"included": SqlFirewallConfigExcludeJobIncluded,
}

// GetSqlFirewallConfigExcludeJobEnumValues Enumerates the set of values for SqlFirewallConfigExcludeJobEnum
func GetSqlFirewallConfigExcludeJobEnumValues() []SqlFirewallConfigExcludeJobEnum {
	values := make([]SqlFirewallConfigExcludeJobEnum, 0)
	for _, v := range mappingSqlFirewallConfigExcludeJobEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallConfigExcludeJobEnumStringValues Enumerates the set of values in String for SqlFirewallConfigExcludeJobEnum
func GetSqlFirewallConfigExcludeJobEnumStringValues() []string {
	return []string{
		"EXCLUDED",
		"INCLUDED",
	}
}

// GetMappingSqlFirewallConfigExcludeJobEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallConfigExcludeJobEnum(val string) (SqlFirewallConfigExcludeJobEnum, bool) {
	enum, ok := mappingSqlFirewallConfigExcludeJobEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
