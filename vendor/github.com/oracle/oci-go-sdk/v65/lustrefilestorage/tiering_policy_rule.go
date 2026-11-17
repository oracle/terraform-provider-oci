// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage with Lustre API
//
// Use the File Storage with Lustre API to manage Lustre file systems and related resources. For more information, see File Storage with Lustre (https://docs.oracle.com/iaas/Content/lustre/home.htm).
//

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TieringPolicyRule A path-based tiering policy rule.
type TieringPolicyRule struct {

	// One or more absolute path prefixes in the Lustre file system to which this rule applies.
	// The service preserves the order of path prefixes in responses. Each path prefix can contain
	// up to 4096 characters. Each rule must contain at least one path prefix, and a tiering policy
	// can contain up to 100 path prefixes total across all rules.
	PathPrefixes []string `mandatory:"true" json:"pathPrefixes"`

	// The action to apply to files whose paths match this rule. `ARCHIVE` makes matching files eligible to be archived to Object Storage. `SKIP` excludes matching files from being archived.
	Action TieringPolicyRuleActionEnum `mandatory:"true" json:"action"`

	// Optional minimum file size in bytes for an ARCHIVE rule to apply. Must not be specified for SKIP rules.
	MinSizeInBytes *int64 `mandatory:"false" json:"minSizeInBytes"`

	// Required for ARCHIVE rules. Archive files that have not been accessed for this many days. Must not be specified for SKIP rules.
	ArchiveAfterDaysSinceLastAccess *int `mandatory:"false" json:"archiveAfterDaysSinceLastAccess"`

	// Required for ARCHIVE rules. Release archived file data from files that have not been accessed for this many days.
	// This value must be greater than `archiveAfterDaysSinceLastAccess`. Must not be specified for SKIP rules.
	ReleaseAfterDaysSinceLastAccess *int `mandatory:"false" json:"releaseAfterDaysSinceLastAccess"`
}

func (m TieringPolicyRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TieringPolicyRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTieringPolicyRuleActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetTieringPolicyRuleActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TieringPolicyRuleActionEnum Enum with underlying type: string
type TieringPolicyRuleActionEnum string

// Set of constants representing the allowable values for TieringPolicyRuleActionEnum
const (
	TieringPolicyRuleActionArchive TieringPolicyRuleActionEnum = "ARCHIVE"
	TieringPolicyRuleActionSkip    TieringPolicyRuleActionEnum = "SKIP"
)

var mappingTieringPolicyRuleActionEnum = map[string]TieringPolicyRuleActionEnum{
	"ARCHIVE": TieringPolicyRuleActionArchive,
	"SKIP":    TieringPolicyRuleActionSkip,
}

var mappingTieringPolicyRuleActionEnumLowerCase = map[string]TieringPolicyRuleActionEnum{
	"archive": TieringPolicyRuleActionArchive,
	"skip":    TieringPolicyRuleActionSkip,
}

// GetTieringPolicyRuleActionEnumValues Enumerates the set of values for TieringPolicyRuleActionEnum
func GetTieringPolicyRuleActionEnumValues() []TieringPolicyRuleActionEnum {
	values := make([]TieringPolicyRuleActionEnum, 0)
	for _, v := range mappingTieringPolicyRuleActionEnum {
		values = append(values, v)
	}
	return values
}

// GetTieringPolicyRuleActionEnumStringValues Enumerates the set of values in String for TieringPolicyRuleActionEnum
func GetTieringPolicyRuleActionEnumStringValues() []string {
	return []string{
		"ARCHIVE",
		"SKIP",
	}
}

// GetMappingTieringPolicyRuleActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTieringPolicyRuleActionEnum(val string) (TieringPolicyRuleActionEnum, bool) {
	enum, ok := mappingTieringPolicyRuleActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
