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

// UpdateSqlFirewallConfigDetails Details to update the SQL firewall config.
type UpdateSqlFirewallConfigDetails struct {

	// Specifies whether the firewall is enabled or disabled on the target database.
	Status UpdateSqlFirewallConfigDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Specifies whether Data Safe should automatically purge the violation logs
	// from the database after collecting the violation logs and persisting on Data Safe.
	ViolationLogAutoPurge UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum `mandatory:"false" json:"violationLogAutoPurge,omitempty"`

	// Specifies whether the firewall should include or exclude the database internal job activities.
	ExcludeJob UpdateSqlFirewallConfigDetailsExcludeJobEnum `mandatory:"false" json:"excludeJob,omitempty"`
}

func (m UpdateSqlFirewallConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSqlFirewallConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateSqlFirewallConfigDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUpdateSqlFirewallConfigDetailsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum(string(m.ViolationLogAutoPurge)); !ok && m.ViolationLogAutoPurge != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ViolationLogAutoPurge: %s. Supported values are: %s.", m.ViolationLogAutoPurge, strings.Join(GetUpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateSqlFirewallConfigDetailsExcludeJobEnum(string(m.ExcludeJob)); !ok && m.ExcludeJob != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExcludeJob: %s. Supported values are: %s.", m.ExcludeJob, strings.Join(GetUpdateSqlFirewallConfigDetailsExcludeJobEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateSqlFirewallConfigDetailsStatusEnum Enum with underlying type: string
type UpdateSqlFirewallConfigDetailsStatusEnum string

// Set of constants representing the allowable values for UpdateSqlFirewallConfigDetailsStatusEnum
const (
	UpdateSqlFirewallConfigDetailsStatusEnabled  UpdateSqlFirewallConfigDetailsStatusEnum = "ENABLED"
	UpdateSqlFirewallConfigDetailsStatusDisabled UpdateSqlFirewallConfigDetailsStatusEnum = "DISABLED"
)

var mappingUpdateSqlFirewallConfigDetailsStatusEnum = map[string]UpdateSqlFirewallConfigDetailsStatusEnum{
	"ENABLED":  UpdateSqlFirewallConfigDetailsStatusEnabled,
	"DISABLED": UpdateSqlFirewallConfigDetailsStatusDisabled,
}

var mappingUpdateSqlFirewallConfigDetailsStatusEnumLowerCase = map[string]UpdateSqlFirewallConfigDetailsStatusEnum{
	"enabled":  UpdateSqlFirewallConfigDetailsStatusEnabled,
	"disabled": UpdateSqlFirewallConfigDetailsStatusDisabled,
}

// GetUpdateSqlFirewallConfigDetailsStatusEnumValues Enumerates the set of values for UpdateSqlFirewallConfigDetailsStatusEnum
func GetUpdateSqlFirewallConfigDetailsStatusEnumValues() []UpdateSqlFirewallConfigDetailsStatusEnum {
	values := make([]UpdateSqlFirewallConfigDetailsStatusEnum, 0)
	for _, v := range mappingUpdateSqlFirewallConfigDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSqlFirewallConfigDetailsStatusEnumStringValues Enumerates the set of values in String for UpdateSqlFirewallConfigDetailsStatusEnum
func GetUpdateSqlFirewallConfigDetailsStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingUpdateSqlFirewallConfigDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSqlFirewallConfigDetailsStatusEnum(val string) (UpdateSqlFirewallConfigDetailsStatusEnum, bool) {
	enum, ok := mappingUpdateSqlFirewallConfigDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum Enum with underlying type: string
type UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum string

// Set of constants representing the allowable values for UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum
const (
	UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnabled  UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum = "ENABLED"
	UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeDisabled UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum = "DISABLED"
)

var mappingUpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum = map[string]UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum{
	"ENABLED":  UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnabled,
	"DISABLED": UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeDisabled,
}

var mappingUpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnumLowerCase = map[string]UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum{
	"enabled":  UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnabled,
	"disabled": UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeDisabled,
}

// GetUpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnumValues Enumerates the set of values for UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum
func GetUpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnumValues() []UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum {
	values := make([]UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum, 0)
	for _, v := range mappingUpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnumStringValues Enumerates the set of values in String for UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum
func GetUpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingUpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum(val string) (UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum, bool) {
	enum, ok := mappingUpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateSqlFirewallConfigDetailsExcludeJobEnum Enum with underlying type: string
type UpdateSqlFirewallConfigDetailsExcludeJobEnum string

// Set of constants representing the allowable values for UpdateSqlFirewallConfigDetailsExcludeJobEnum
const (
	UpdateSqlFirewallConfigDetailsExcludeJobExcluded UpdateSqlFirewallConfigDetailsExcludeJobEnum = "EXCLUDED"
	UpdateSqlFirewallConfigDetailsExcludeJobIncluded UpdateSqlFirewallConfigDetailsExcludeJobEnum = "INCLUDED"
)

var mappingUpdateSqlFirewallConfigDetailsExcludeJobEnum = map[string]UpdateSqlFirewallConfigDetailsExcludeJobEnum{
	"EXCLUDED": UpdateSqlFirewallConfigDetailsExcludeJobExcluded,
	"INCLUDED": UpdateSqlFirewallConfigDetailsExcludeJobIncluded,
}

var mappingUpdateSqlFirewallConfigDetailsExcludeJobEnumLowerCase = map[string]UpdateSqlFirewallConfigDetailsExcludeJobEnum{
	"excluded": UpdateSqlFirewallConfigDetailsExcludeJobExcluded,
	"included": UpdateSqlFirewallConfigDetailsExcludeJobIncluded,
}

// GetUpdateSqlFirewallConfigDetailsExcludeJobEnumValues Enumerates the set of values for UpdateSqlFirewallConfigDetailsExcludeJobEnum
func GetUpdateSqlFirewallConfigDetailsExcludeJobEnumValues() []UpdateSqlFirewallConfigDetailsExcludeJobEnum {
	values := make([]UpdateSqlFirewallConfigDetailsExcludeJobEnum, 0)
	for _, v := range mappingUpdateSqlFirewallConfigDetailsExcludeJobEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSqlFirewallConfigDetailsExcludeJobEnumStringValues Enumerates the set of values in String for UpdateSqlFirewallConfigDetailsExcludeJobEnum
func GetUpdateSqlFirewallConfigDetailsExcludeJobEnumStringValues() []string {
	return []string{
		"EXCLUDED",
		"INCLUDED",
	}
}

// GetMappingUpdateSqlFirewallConfigDetailsExcludeJobEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSqlFirewallConfigDetailsExcludeJobEnum(val string) (UpdateSqlFirewallConfigDetailsExcludeJobEnum, bool) {
	enum, ok := mappingUpdateSqlFirewallConfigDetailsExcludeJobEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
