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

// FirewallConfigDetails Details to update the SQL Firewall configuration.
type FirewallConfigDetails struct {

	// Specifies whether the firewall is enabled or disabled.
	Status FirewallConfigDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Specifies whether Data Safe should automatically purge the violation logs
	// from the database after collecting the violation logs and persisting them in Data Safe.
	ViolationLogAutoPurge FirewallConfigDetailsViolationLogAutoPurgeEnum `mandatory:"false" json:"violationLogAutoPurge,omitempty"`

	// Specifies whether the firewall should include or exclude the database internal job activities.
	ExcludeJob FirewallConfigDetailsExcludeJobEnum `mandatory:"false" json:"excludeJob,omitempty"`
}

func (m FirewallConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FirewallConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFirewallConfigDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetFirewallConfigDetailsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFirewallConfigDetailsViolationLogAutoPurgeEnum(string(m.ViolationLogAutoPurge)); !ok && m.ViolationLogAutoPurge != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ViolationLogAutoPurge: %s. Supported values are: %s.", m.ViolationLogAutoPurge, strings.Join(GetFirewallConfigDetailsViolationLogAutoPurgeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFirewallConfigDetailsExcludeJobEnum(string(m.ExcludeJob)); !ok && m.ExcludeJob != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExcludeJob: %s. Supported values are: %s.", m.ExcludeJob, strings.Join(GetFirewallConfigDetailsExcludeJobEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FirewallConfigDetailsStatusEnum Enum with underlying type: string
type FirewallConfigDetailsStatusEnum string

// Set of constants representing the allowable values for FirewallConfigDetailsStatusEnum
const (
	FirewallConfigDetailsStatusEnabled  FirewallConfigDetailsStatusEnum = "ENABLED"
	FirewallConfigDetailsStatusDisabled FirewallConfigDetailsStatusEnum = "DISABLED"
)

var mappingFirewallConfigDetailsStatusEnum = map[string]FirewallConfigDetailsStatusEnum{
	"ENABLED":  FirewallConfigDetailsStatusEnabled,
	"DISABLED": FirewallConfigDetailsStatusDisabled,
}

var mappingFirewallConfigDetailsStatusEnumLowerCase = map[string]FirewallConfigDetailsStatusEnum{
	"enabled":  FirewallConfigDetailsStatusEnabled,
	"disabled": FirewallConfigDetailsStatusDisabled,
}

// GetFirewallConfigDetailsStatusEnumValues Enumerates the set of values for FirewallConfigDetailsStatusEnum
func GetFirewallConfigDetailsStatusEnumValues() []FirewallConfigDetailsStatusEnum {
	values := make([]FirewallConfigDetailsStatusEnum, 0)
	for _, v := range mappingFirewallConfigDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetFirewallConfigDetailsStatusEnumStringValues Enumerates the set of values in String for FirewallConfigDetailsStatusEnum
func GetFirewallConfigDetailsStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingFirewallConfigDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFirewallConfigDetailsStatusEnum(val string) (FirewallConfigDetailsStatusEnum, bool) {
	enum, ok := mappingFirewallConfigDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FirewallConfigDetailsViolationLogAutoPurgeEnum Enum with underlying type: string
type FirewallConfigDetailsViolationLogAutoPurgeEnum string

// Set of constants representing the allowable values for FirewallConfigDetailsViolationLogAutoPurgeEnum
const (
	FirewallConfigDetailsViolationLogAutoPurgeEnabled  FirewallConfigDetailsViolationLogAutoPurgeEnum = "ENABLED"
	FirewallConfigDetailsViolationLogAutoPurgeDisabled FirewallConfigDetailsViolationLogAutoPurgeEnum = "DISABLED"
)

var mappingFirewallConfigDetailsViolationLogAutoPurgeEnum = map[string]FirewallConfigDetailsViolationLogAutoPurgeEnum{
	"ENABLED":  FirewallConfigDetailsViolationLogAutoPurgeEnabled,
	"DISABLED": FirewallConfigDetailsViolationLogAutoPurgeDisabled,
}

var mappingFirewallConfigDetailsViolationLogAutoPurgeEnumLowerCase = map[string]FirewallConfigDetailsViolationLogAutoPurgeEnum{
	"enabled":  FirewallConfigDetailsViolationLogAutoPurgeEnabled,
	"disabled": FirewallConfigDetailsViolationLogAutoPurgeDisabled,
}

// GetFirewallConfigDetailsViolationLogAutoPurgeEnumValues Enumerates the set of values for FirewallConfigDetailsViolationLogAutoPurgeEnum
func GetFirewallConfigDetailsViolationLogAutoPurgeEnumValues() []FirewallConfigDetailsViolationLogAutoPurgeEnum {
	values := make([]FirewallConfigDetailsViolationLogAutoPurgeEnum, 0)
	for _, v := range mappingFirewallConfigDetailsViolationLogAutoPurgeEnum {
		values = append(values, v)
	}
	return values
}

// GetFirewallConfigDetailsViolationLogAutoPurgeEnumStringValues Enumerates the set of values in String for FirewallConfigDetailsViolationLogAutoPurgeEnum
func GetFirewallConfigDetailsViolationLogAutoPurgeEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingFirewallConfigDetailsViolationLogAutoPurgeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFirewallConfigDetailsViolationLogAutoPurgeEnum(val string) (FirewallConfigDetailsViolationLogAutoPurgeEnum, bool) {
	enum, ok := mappingFirewallConfigDetailsViolationLogAutoPurgeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FirewallConfigDetailsExcludeJobEnum Enum with underlying type: string
type FirewallConfigDetailsExcludeJobEnum string

// Set of constants representing the allowable values for FirewallConfigDetailsExcludeJobEnum
const (
	FirewallConfigDetailsExcludeJobExcluded FirewallConfigDetailsExcludeJobEnum = "EXCLUDED"
	FirewallConfigDetailsExcludeJobIncluded FirewallConfigDetailsExcludeJobEnum = "INCLUDED"
)

var mappingFirewallConfigDetailsExcludeJobEnum = map[string]FirewallConfigDetailsExcludeJobEnum{
	"EXCLUDED": FirewallConfigDetailsExcludeJobExcluded,
	"INCLUDED": FirewallConfigDetailsExcludeJobIncluded,
}

var mappingFirewallConfigDetailsExcludeJobEnumLowerCase = map[string]FirewallConfigDetailsExcludeJobEnum{
	"excluded": FirewallConfigDetailsExcludeJobExcluded,
	"included": FirewallConfigDetailsExcludeJobIncluded,
}

// GetFirewallConfigDetailsExcludeJobEnumValues Enumerates the set of values for FirewallConfigDetailsExcludeJobEnum
func GetFirewallConfigDetailsExcludeJobEnumValues() []FirewallConfigDetailsExcludeJobEnum {
	values := make([]FirewallConfigDetailsExcludeJobEnum, 0)
	for _, v := range mappingFirewallConfigDetailsExcludeJobEnum {
		values = append(values, v)
	}
	return values
}

// GetFirewallConfigDetailsExcludeJobEnumStringValues Enumerates the set of values in String for FirewallConfigDetailsExcludeJobEnum
func GetFirewallConfigDetailsExcludeJobEnumStringValues() []string {
	return []string{
		"EXCLUDED",
		"INCLUDED",
	}
}

// GetMappingFirewallConfigDetailsExcludeJobEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFirewallConfigDetailsExcludeJobEnum(val string) (FirewallConfigDetailsExcludeJobEnum, bool) {
	enum, ok := mappingFirewallConfigDetailsExcludeJobEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
