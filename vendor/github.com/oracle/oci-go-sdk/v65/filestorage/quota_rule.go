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

// QuotaRule A rule that can restrict the logical space that a user or group can consume in a file system.
type QuotaRule struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file System.
	FileSystemId *string `mandatory:"true" json:"fileSystemId"`

	// The flag is an identifier to tell whether the quota rule will be enforced.
	// If `isHardQuota` is false, the quota rule will be enforced so the usage cannot exceed the hard quota limit.
	// If `isHardQuota` is true, usage can exceed the soft quota limit. An alarm or notification will be sent to
	// the customer, if the specific usage exceeds.
	IsHardQuota *bool `mandatory:"true" json:"isHardQuota"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `UserXYZ's quota`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The value of the quota rule. The unit is Gigabyte.
	QuotaLimitInGigabytes *int `mandatory:"true" json:"quotaLimitInGigabytes"`

	// The date and time the quota rule was started, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the quota rule was last updated, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The type of the owner of this quota rule and usage.
	PrincipalType QuotaRulePrincipalTypeEnum `mandatory:"false" json:"principalType,omitempty"`

	// An identifier for the owner of this usage and quota rule. Unix-like operating systems use this integer value to
	// identify a user or group to manage access control.
	PrincipalId *int `mandatory:"false" json:"principalId"`
}

func (m QuotaRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QuotaRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingQuotaRulePrincipalTypeEnum(string(m.PrincipalType)); !ok && m.PrincipalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrincipalType: %s. Supported values are: %s.", m.PrincipalType, strings.Join(GetQuotaRulePrincipalTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QuotaRulePrincipalTypeEnum Enum with underlying type: string
type QuotaRulePrincipalTypeEnum string

// Set of constants representing the allowable values for QuotaRulePrincipalTypeEnum
const (
	QuotaRulePrincipalTypeFileSystemLevel QuotaRulePrincipalTypeEnum = "FILE_SYSTEM_LEVEL"
	QuotaRulePrincipalTypeDefaultGroup    QuotaRulePrincipalTypeEnum = "DEFAULT_GROUP"
	QuotaRulePrincipalTypeDefaultUser     QuotaRulePrincipalTypeEnum = "DEFAULT_USER"
	QuotaRulePrincipalTypeIndividualGroup QuotaRulePrincipalTypeEnum = "INDIVIDUAL_GROUP"
	QuotaRulePrincipalTypeIndividualUser  QuotaRulePrincipalTypeEnum = "INDIVIDUAL_USER"
)

var mappingQuotaRulePrincipalTypeEnum = map[string]QuotaRulePrincipalTypeEnum{
	"FILE_SYSTEM_LEVEL": QuotaRulePrincipalTypeFileSystemLevel,
	"DEFAULT_GROUP":     QuotaRulePrincipalTypeDefaultGroup,
	"DEFAULT_USER":      QuotaRulePrincipalTypeDefaultUser,
	"INDIVIDUAL_GROUP":  QuotaRulePrincipalTypeIndividualGroup,
	"INDIVIDUAL_USER":   QuotaRulePrincipalTypeIndividualUser,
}

var mappingQuotaRulePrincipalTypeEnumLowerCase = map[string]QuotaRulePrincipalTypeEnum{
	"file_system_level": QuotaRulePrincipalTypeFileSystemLevel,
	"default_group":     QuotaRulePrincipalTypeDefaultGroup,
	"default_user":      QuotaRulePrincipalTypeDefaultUser,
	"individual_group":  QuotaRulePrincipalTypeIndividualGroup,
	"individual_user":   QuotaRulePrincipalTypeIndividualUser,
}

// GetQuotaRulePrincipalTypeEnumValues Enumerates the set of values for QuotaRulePrincipalTypeEnum
func GetQuotaRulePrincipalTypeEnumValues() []QuotaRulePrincipalTypeEnum {
	values := make([]QuotaRulePrincipalTypeEnum, 0)
	for _, v := range mappingQuotaRulePrincipalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetQuotaRulePrincipalTypeEnumStringValues Enumerates the set of values in String for QuotaRulePrincipalTypeEnum
func GetQuotaRulePrincipalTypeEnumStringValues() []string {
	return []string{
		"FILE_SYSTEM_LEVEL",
		"DEFAULT_GROUP",
		"DEFAULT_USER",
		"INDIVIDUAL_GROUP",
		"INDIVIDUAL_USER",
	}
}

// GetMappingQuotaRulePrincipalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQuotaRulePrincipalTypeEnum(val string) (QuotaRulePrincipalTypeEnum, bool) {
	enum, ok := mappingQuotaRulePrincipalTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
