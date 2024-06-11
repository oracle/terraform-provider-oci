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

// QuotaRuleSummary Summary information for a principal's usage and quota rule.
type QuotaRuleSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file System.
	FileSystemId *string `mandatory:"true" json:"fileSystemId"`

	// The type of the owner of this quota rule and usage.
	PrincipalType QuotaRuleSummaryPrincipalTypeEnum `mandatory:"true" json:"principalType"`

	// An identifier for the owner of this usage and quota rule. Unix-like operating systems use this integer value to
	// identify a user or group to manage access control.
	PrincipalId *int `mandatory:"true" json:"principalId"`

	// The usage value corresponds to this principal. The unit is Byte.
	UsageInBytes *int64 `mandatory:"true" json:"usageInBytes"`

	// The identifier of the quota rule. It is the base64 encoded string of the tuple <principalId, principalType, isHardQuota>.
	Id *string `mandatory:"false" json:"id"`

	// The flag is an identifier to tell whether the quota rule will be enforced.
	// If `isHardQuota` is false, the quota rule will be enforced so the usage cannot exceed the hard quota limit.
	// If `isHardQuota` is true, usage can exceed the soft quota limit. An alarm or notification will be sent to
	// the customer, if the specific usage exceeds.
	IsHardQuota *bool `mandatory:"false" json:"isHardQuota"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `UserXYZ's quota`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The value of the quota rule. The unit is Gigabyte.
	QuotaLimitInGigabytes *int `mandatory:"false" json:"quotaLimitInGigabytes"`

	// The date and time the quota rule was started, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the quota rule was last updated, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// An option to only display the users or groups that violate their quota rules.
	// If `areViolatorsOnly` is false, the list result will display all the quota and usage report.
	// If `areViolatorsOnly` is true, the list result will only display the quota and usage report for
	// the users or groups that violate their quota rules.
	AreViolatorsOnly *bool `mandatory:"false" json:"areViolatorsOnly"`
}

func (m QuotaRuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QuotaRuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQuotaRuleSummaryPrincipalTypeEnum(string(m.PrincipalType)); !ok && m.PrincipalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrincipalType: %s. Supported values are: %s.", m.PrincipalType, strings.Join(GetQuotaRuleSummaryPrincipalTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QuotaRuleSummaryPrincipalTypeEnum Enum with underlying type: string
type QuotaRuleSummaryPrincipalTypeEnum string

// Set of constants representing the allowable values for QuotaRuleSummaryPrincipalTypeEnum
const (
	QuotaRuleSummaryPrincipalTypeFileSystemLevel QuotaRuleSummaryPrincipalTypeEnum = "FILE_SYSTEM_LEVEL"
	QuotaRuleSummaryPrincipalTypeDefaultGroup    QuotaRuleSummaryPrincipalTypeEnum = "DEFAULT_GROUP"
	QuotaRuleSummaryPrincipalTypeDefaultUser     QuotaRuleSummaryPrincipalTypeEnum = "DEFAULT_USER"
	QuotaRuleSummaryPrincipalTypeIndividualGroup QuotaRuleSummaryPrincipalTypeEnum = "INDIVIDUAL_GROUP"
	QuotaRuleSummaryPrincipalTypeIndividualUser  QuotaRuleSummaryPrincipalTypeEnum = "INDIVIDUAL_USER"
)

var mappingQuotaRuleSummaryPrincipalTypeEnum = map[string]QuotaRuleSummaryPrincipalTypeEnum{
	"FILE_SYSTEM_LEVEL": QuotaRuleSummaryPrincipalTypeFileSystemLevel,
	"DEFAULT_GROUP":     QuotaRuleSummaryPrincipalTypeDefaultGroup,
	"DEFAULT_USER":      QuotaRuleSummaryPrincipalTypeDefaultUser,
	"INDIVIDUAL_GROUP":  QuotaRuleSummaryPrincipalTypeIndividualGroup,
	"INDIVIDUAL_USER":   QuotaRuleSummaryPrincipalTypeIndividualUser,
}

var mappingQuotaRuleSummaryPrincipalTypeEnumLowerCase = map[string]QuotaRuleSummaryPrincipalTypeEnum{
	"file_system_level": QuotaRuleSummaryPrincipalTypeFileSystemLevel,
	"default_group":     QuotaRuleSummaryPrincipalTypeDefaultGroup,
	"default_user":      QuotaRuleSummaryPrincipalTypeDefaultUser,
	"individual_group":  QuotaRuleSummaryPrincipalTypeIndividualGroup,
	"individual_user":   QuotaRuleSummaryPrincipalTypeIndividualUser,
}

// GetQuotaRuleSummaryPrincipalTypeEnumValues Enumerates the set of values for QuotaRuleSummaryPrincipalTypeEnum
func GetQuotaRuleSummaryPrincipalTypeEnumValues() []QuotaRuleSummaryPrincipalTypeEnum {
	values := make([]QuotaRuleSummaryPrincipalTypeEnum, 0)
	for _, v := range mappingQuotaRuleSummaryPrincipalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetQuotaRuleSummaryPrincipalTypeEnumStringValues Enumerates the set of values in String for QuotaRuleSummaryPrincipalTypeEnum
func GetQuotaRuleSummaryPrincipalTypeEnumStringValues() []string {
	return []string{
		"FILE_SYSTEM_LEVEL",
		"DEFAULT_GROUP",
		"DEFAULT_USER",
		"INDIVIDUAL_GROUP",
		"INDIVIDUAL_USER",
	}
}

// GetMappingQuotaRuleSummaryPrincipalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQuotaRuleSummaryPrincipalTypeEnum(val string) (QuotaRuleSummaryPrincipalTypeEnum, bool) {
	enum, ok := mappingQuotaRuleSummaryPrincipalTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
