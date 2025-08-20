// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateQuotaRuleDetails Details for creating a quota rule in the file system.
type CreateQuotaRuleDetails struct {

	// The type of the owner of this quota rule and usage.
	PrincipalType CreateQuotaRuleDetailsPrincipalTypeEnum `mandatory:"true" json:"principalType"`

	// Whether the quota rule will be enforced.
	// If `isHardQuota` is true, the quota rule is enforced so that the write is blocked if usage
	// exceeds the hard quota limit.
	// If `isHardQuota` is false, writes succeed even if usage exceeds the soft quota limit, but the quota rule is violated.
	IsHardQuota *bool `mandatory:"true" json:"isHardQuota"`

	// The value of the quota rule in gigabytes.
	QuotaLimitInGigabytes *int `mandatory:"true" json:"quotaLimitInGigabytes"`

	// An identifier for the user or the group associated with quota rule and usage. UNIX-like operating systems use this integer value to
	// identify a user or group to manage access control.
	PrincipalId *int `mandatory:"false" json:"principalId"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `UserXYZ's quota`
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CreateQuotaRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateQuotaRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateQuotaRuleDetailsPrincipalTypeEnum(string(m.PrincipalType)); !ok && m.PrincipalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrincipalType: %s. Supported values are: %s.", m.PrincipalType, strings.Join(GetCreateQuotaRuleDetailsPrincipalTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateQuotaRuleDetailsPrincipalTypeEnum Enum with underlying type: string
type CreateQuotaRuleDetailsPrincipalTypeEnum string

// Set of constants representing the allowable values for CreateQuotaRuleDetailsPrincipalTypeEnum
const (
	CreateQuotaRuleDetailsPrincipalTypeFileSystemLevel CreateQuotaRuleDetailsPrincipalTypeEnum = "FILE_SYSTEM_LEVEL"
	CreateQuotaRuleDetailsPrincipalTypeDefaultGroup    CreateQuotaRuleDetailsPrincipalTypeEnum = "DEFAULT_GROUP"
	CreateQuotaRuleDetailsPrincipalTypeDefaultUser     CreateQuotaRuleDetailsPrincipalTypeEnum = "DEFAULT_USER"
	CreateQuotaRuleDetailsPrincipalTypeIndividualGroup CreateQuotaRuleDetailsPrincipalTypeEnum = "INDIVIDUAL_GROUP"
	CreateQuotaRuleDetailsPrincipalTypeIndividualUser  CreateQuotaRuleDetailsPrincipalTypeEnum = "INDIVIDUAL_USER"
)

var mappingCreateQuotaRuleDetailsPrincipalTypeEnum = map[string]CreateQuotaRuleDetailsPrincipalTypeEnum{
	"FILE_SYSTEM_LEVEL": CreateQuotaRuleDetailsPrincipalTypeFileSystemLevel,
	"DEFAULT_GROUP":     CreateQuotaRuleDetailsPrincipalTypeDefaultGroup,
	"DEFAULT_USER":      CreateQuotaRuleDetailsPrincipalTypeDefaultUser,
	"INDIVIDUAL_GROUP":  CreateQuotaRuleDetailsPrincipalTypeIndividualGroup,
	"INDIVIDUAL_USER":   CreateQuotaRuleDetailsPrincipalTypeIndividualUser,
}

var mappingCreateQuotaRuleDetailsPrincipalTypeEnumLowerCase = map[string]CreateQuotaRuleDetailsPrincipalTypeEnum{
	"file_system_level": CreateQuotaRuleDetailsPrincipalTypeFileSystemLevel,
	"default_group":     CreateQuotaRuleDetailsPrincipalTypeDefaultGroup,
	"default_user":      CreateQuotaRuleDetailsPrincipalTypeDefaultUser,
	"individual_group":  CreateQuotaRuleDetailsPrincipalTypeIndividualGroup,
	"individual_user":   CreateQuotaRuleDetailsPrincipalTypeIndividualUser,
}

// GetCreateQuotaRuleDetailsPrincipalTypeEnumValues Enumerates the set of values for CreateQuotaRuleDetailsPrincipalTypeEnum
func GetCreateQuotaRuleDetailsPrincipalTypeEnumValues() []CreateQuotaRuleDetailsPrincipalTypeEnum {
	values := make([]CreateQuotaRuleDetailsPrincipalTypeEnum, 0)
	for _, v := range mappingCreateQuotaRuleDetailsPrincipalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateQuotaRuleDetailsPrincipalTypeEnumStringValues Enumerates the set of values in String for CreateQuotaRuleDetailsPrincipalTypeEnum
func GetCreateQuotaRuleDetailsPrincipalTypeEnumStringValues() []string {
	return []string{
		"FILE_SYSTEM_LEVEL",
		"DEFAULT_GROUP",
		"DEFAULT_USER",
		"INDIVIDUAL_GROUP",
		"INDIVIDUAL_USER",
	}
}

// GetMappingCreateQuotaRuleDetailsPrincipalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateQuotaRuleDetailsPrincipalTypeEnum(val string) (CreateQuotaRuleDetailsPrincipalTypeEnum, bool) {
	enum, ok := mappingCreateQuotaRuleDetailsPrincipalTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
