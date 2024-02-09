// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddQuotaRuleDetails Details for creating a quota rule in the file system.
type AddQuotaRuleDetails struct {

	// The type of the owner of this quota rule and usage.
	PrincipalType AddQuotaRuleDetailsPrincipalTypeEnum `mandatory:"true" json:"principalType"`

	// An identifier for the owner of this usage and quota rule. Unix-like operating systems use this integer value to
	// identify a user or group to manage access control.
	PrincipalId *int `mandatory:"true" json:"principalId"`

	// The flag is an identifier to tell whether the quota rule will be enforced.
	// If `isHardQuota` is true, the quota rule will be enforced so the write will be blocked if usage
	// exceeds the hard quota limit.
	// If `isHardQuota` is false, usage can exceed the soft quota limit. An alarm or notification will be sent to
	// the customer, if the specific usage exceeds.
	IsHardQuota *bool `mandatory:"true" json:"isHardQuota"`

	// The value of the quota rule. The unit is Gigabyte.
	QuotaLimitInGigabytes *int `mandatory:"true" json:"quotaLimitInGigabytes"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `UserXYZ's quota`
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m AddQuotaRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddQuotaRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddQuotaRuleDetailsPrincipalTypeEnum(string(m.PrincipalType)); !ok && m.PrincipalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrincipalType: %s. Supported values are: %s.", m.PrincipalType, strings.Join(GetAddQuotaRuleDetailsPrincipalTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddQuotaRuleDetailsPrincipalTypeEnum Enum with underlying type: string
type AddQuotaRuleDetailsPrincipalTypeEnum string

// Set of constants representing the allowable values for AddQuotaRuleDetailsPrincipalTypeEnum
const (
	AddQuotaRuleDetailsPrincipalTypeFileSystemLevel AddQuotaRuleDetailsPrincipalTypeEnum = "FILE_SYSTEM_LEVEL"
	AddQuotaRuleDetailsPrincipalTypeDefaultGroup    AddQuotaRuleDetailsPrincipalTypeEnum = "DEFAULT_GROUP"
	AddQuotaRuleDetailsPrincipalTypeDefaultUser     AddQuotaRuleDetailsPrincipalTypeEnum = "DEFAULT_USER"
	AddQuotaRuleDetailsPrincipalTypeIndividualGroup AddQuotaRuleDetailsPrincipalTypeEnum = "INDIVIDUAL_GROUP"
	AddQuotaRuleDetailsPrincipalTypeIndividualUser  AddQuotaRuleDetailsPrincipalTypeEnum = "INDIVIDUAL_USER"
)

var mappingAddQuotaRuleDetailsPrincipalTypeEnum = map[string]AddQuotaRuleDetailsPrincipalTypeEnum{
	"FILE_SYSTEM_LEVEL": AddQuotaRuleDetailsPrincipalTypeFileSystemLevel,
	"DEFAULT_GROUP":     AddQuotaRuleDetailsPrincipalTypeDefaultGroup,
	"DEFAULT_USER":      AddQuotaRuleDetailsPrincipalTypeDefaultUser,
	"INDIVIDUAL_GROUP":  AddQuotaRuleDetailsPrincipalTypeIndividualGroup,
	"INDIVIDUAL_USER":   AddQuotaRuleDetailsPrincipalTypeIndividualUser,
}

var mappingAddQuotaRuleDetailsPrincipalTypeEnumLowerCase = map[string]AddQuotaRuleDetailsPrincipalTypeEnum{
	"file_system_level": AddQuotaRuleDetailsPrincipalTypeFileSystemLevel,
	"default_group":     AddQuotaRuleDetailsPrincipalTypeDefaultGroup,
	"default_user":      AddQuotaRuleDetailsPrincipalTypeDefaultUser,
	"individual_group":  AddQuotaRuleDetailsPrincipalTypeIndividualGroup,
	"individual_user":   AddQuotaRuleDetailsPrincipalTypeIndividualUser,
}

// GetAddQuotaRuleDetailsPrincipalTypeEnumValues Enumerates the set of values for AddQuotaRuleDetailsPrincipalTypeEnum
func GetAddQuotaRuleDetailsPrincipalTypeEnumValues() []AddQuotaRuleDetailsPrincipalTypeEnum {
	values := make([]AddQuotaRuleDetailsPrincipalTypeEnum, 0)
	for _, v := range mappingAddQuotaRuleDetailsPrincipalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddQuotaRuleDetailsPrincipalTypeEnumStringValues Enumerates the set of values in String for AddQuotaRuleDetailsPrincipalTypeEnum
func GetAddQuotaRuleDetailsPrincipalTypeEnumStringValues() []string {
	return []string{
		"FILE_SYSTEM_LEVEL",
		"DEFAULT_GROUP",
		"DEFAULT_USER",
		"INDIVIDUAL_GROUP",
		"INDIVIDUAL_USER",
	}
}

// GetMappingAddQuotaRuleDetailsPrincipalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddQuotaRuleDetailsPrincipalTypeEnum(val string) (AddQuotaRuleDetailsPrincipalTypeEnum, bool) {
	enum, ok := mappingAddQuotaRuleDetailsPrincipalTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
