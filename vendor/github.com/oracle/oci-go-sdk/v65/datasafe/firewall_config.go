// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// FirewallConfig The SQL Firewall related configurations.
type FirewallConfig struct {

	// Specifies if the firewall is enabled or disabled.
	Status FirewallConfigStatusEnum `mandatory:"true" json:"status"`

	// Specifies whether Data Safe should automatically purge the violation logs
	// from the database after collecting the violation logs and persisting on Data Safe.
	ViolationLogAutoPurge FirewallConfigViolationLogAutoPurgeEnum `mandatory:"true" json:"violationLogAutoPurge"`

	// Specifies whether the firewall should include or exclude the database internal job activities.
	ExcludeJob FirewallConfigExcludeJobEnum `mandatory:"false" json:"excludeJob,omitempty"`

	// The date and time the firewall configuration was last updated, in the format defined by RFC3339.
	TimeStatusUpdated *common.SDKTime `mandatory:"false" json:"timeStatusUpdated"`
}

func (m FirewallConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FirewallConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFirewallConfigStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetFirewallConfigStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFirewallConfigViolationLogAutoPurgeEnum(string(m.ViolationLogAutoPurge)); !ok && m.ViolationLogAutoPurge != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ViolationLogAutoPurge: %s. Supported values are: %s.", m.ViolationLogAutoPurge, strings.Join(GetFirewallConfigViolationLogAutoPurgeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingFirewallConfigExcludeJobEnum(string(m.ExcludeJob)); !ok && m.ExcludeJob != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExcludeJob: %s. Supported values are: %s.", m.ExcludeJob, strings.Join(GetFirewallConfigExcludeJobEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FirewallConfigStatusEnum Enum with underlying type: string
type FirewallConfigStatusEnum string

// Set of constants representing the allowable values for FirewallConfigStatusEnum
const (
	FirewallConfigStatusEnabled  FirewallConfigStatusEnum = "ENABLED"
	FirewallConfigStatusDisabled FirewallConfigStatusEnum = "DISABLED"
)

var mappingFirewallConfigStatusEnum = map[string]FirewallConfigStatusEnum{
	"ENABLED":  FirewallConfigStatusEnabled,
	"DISABLED": FirewallConfigStatusDisabled,
}

var mappingFirewallConfigStatusEnumLowerCase = map[string]FirewallConfigStatusEnum{
	"enabled":  FirewallConfigStatusEnabled,
	"disabled": FirewallConfigStatusDisabled,
}

// GetFirewallConfigStatusEnumValues Enumerates the set of values for FirewallConfigStatusEnum
func GetFirewallConfigStatusEnumValues() []FirewallConfigStatusEnum {
	values := make([]FirewallConfigStatusEnum, 0)
	for _, v := range mappingFirewallConfigStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetFirewallConfigStatusEnumStringValues Enumerates the set of values in String for FirewallConfigStatusEnum
func GetFirewallConfigStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingFirewallConfigStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFirewallConfigStatusEnum(val string) (FirewallConfigStatusEnum, bool) {
	enum, ok := mappingFirewallConfigStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FirewallConfigViolationLogAutoPurgeEnum Enum with underlying type: string
type FirewallConfigViolationLogAutoPurgeEnum string

// Set of constants representing the allowable values for FirewallConfigViolationLogAutoPurgeEnum
const (
	FirewallConfigViolationLogAutoPurgeEnabled  FirewallConfigViolationLogAutoPurgeEnum = "ENABLED"
	FirewallConfigViolationLogAutoPurgeDisabled FirewallConfigViolationLogAutoPurgeEnum = "DISABLED"
)

var mappingFirewallConfigViolationLogAutoPurgeEnum = map[string]FirewallConfigViolationLogAutoPurgeEnum{
	"ENABLED":  FirewallConfigViolationLogAutoPurgeEnabled,
	"DISABLED": FirewallConfigViolationLogAutoPurgeDisabled,
}

var mappingFirewallConfigViolationLogAutoPurgeEnumLowerCase = map[string]FirewallConfigViolationLogAutoPurgeEnum{
	"enabled":  FirewallConfigViolationLogAutoPurgeEnabled,
	"disabled": FirewallConfigViolationLogAutoPurgeDisabled,
}

// GetFirewallConfigViolationLogAutoPurgeEnumValues Enumerates the set of values for FirewallConfigViolationLogAutoPurgeEnum
func GetFirewallConfigViolationLogAutoPurgeEnumValues() []FirewallConfigViolationLogAutoPurgeEnum {
	values := make([]FirewallConfigViolationLogAutoPurgeEnum, 0)
	for _, v := range mappingFirewallConfigViolationLogAutoPurgeEnum {
		values = append(values, v)
	}
	return values
}

// GetFirewallConfigViolationLogAutoPurgeEnumStringValues Enumerates the set of values in String for FirewallConfigViolationLogAutoPurgeEnum
func GetFirewallConfigViolationLogAutoPurgeEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingFirewallConfigViolationLogAutoPurgeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFirewallConfigViolationLogAutoPurgeEnum(val string) (FirewallConfigViolationLogAutoPurgeEnum, bool) {
	enum, ok := mappingFirewallConfigViolationLogAutoPurgeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FirewallConfigExcludeJobEnum Enum with underlying type: string
type FirewallConfigExcludeJobEnum string

// Set of constants representing the allowable values for FirewallConfigExcludeJobEnum
const (
	FirewallConfigExcludeJobExcluded FirewallConfigExcludeJobEnum = "EXCLUDED"
	FirewallConfigExcludeJobIncluded FirewallConfigExcludeJobEnum = "INCLUDED"
)

var mappingFirewallConfigExcludeJobEnum = map[string]FirewallConfigExcludeJobEnum{
	"EXCLUDED": FirewallConfigExcludeJobExcluded,
	"INCLUDED": FirewallConfigExcludeJobIncluded,
}

var mappingFirewallConfigExcludeJobEnumLowerCase = map[string]FirewallConfigExcludeJobEnum{
	"excluded": FirewallConfigExcludeJobExcluded,
	"included": FirewallConfigExcludeJobIncluded,
}

// GetFirewallConfigExcludeJobEnumValues Enumerates the set of values for FirewallConfigExcludeJobEnum
func GetFirewallConfigExcludeJobEnumValues() []FirewallConfigExcludeJobEnum {
	values := make([]FirewallConfigExcludeJobEnum, 0)
	for _, v := range mappingFirewallConfigExcludeJobEnum {
		values = append(values, v)
	}
	return values
}

// GetFirewallConfigExcludeJobEnumStringValues Enumerates the set of values in String for FirewallConfigExcludeJobEnum
func GetFirewallConfigExcludeJobEnumStringValues() []string {
	return []string{
		"EXCLUDED",
		"INCLUDED",
	}
}

// GetMappingFirewallConfigExcludeJobEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFirewallConfigExcludeJobEnum(val string) (FirewallConfigExcludeJobEnum, bool) {
	enum, ok := mappingFirewallConfigExcludeJobEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
