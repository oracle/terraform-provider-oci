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

// TieringConfigurationSummary Summary of tiering configuration for a Lustre file system.
type TieringConfigurationSummary struct {

	// Whether tiering is enabled for the Lustre file system.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Whether the built-in tiering policy rule or custom policy rules are used.
	PolicyMode TieringConfigurationSummaryPolicyModeEnum `mandatory:"true" json:"policyMode"`

	// Service-computed number of custom tiering policy rules.
	RuleCount *int `mandatory:"true" json:"ruleCount"`
}

func (m TieringConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TieringConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTieringConfigurationSummaryPolicyModeEnum(string(m.PolicyMode)); !ok && m.PolicyMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PolicyMode: %s. Supported values are: %s.", m.PolicyMode, strings.Join(GetTieringConfigurationSummaryPolicyModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TieringConfigurationSummaryPolicyModeEnum Enum with underlying type: string
type TieringConfigurationSummaryPolicyModeEnum string

// Set of constants representing the allowable values for TieringConfigurationSummaryPolicyModeEnum
const (
	TieringConfigurationSummaryPolicyModeDefault TieringConfigurationSummaryPolicyModeEnum = "DEFAULT"
	TieringConfigurationSummaryPolicyModeCustom  TieringConfigurationSummaryPolicyModeEnum = "CUSTOM"
)

var mappingTieringConfigurationSummaryPolicyModeEnum = map[string]TieringConfigurationSummaryPolicyModeEnum{
	"DEFAULT": TieringConfigurationSummaryPolicyModeDefault,
	"CUSTOM":  TieringConfigurationSummaryPolicyModeCustom,
}

var mappingTieringConfigurationSummaryPolicyModeEnumLowerCase = map[string]TieringConfigurationSummaryPolicyModeEnum{
	"default": TieringConfigurationSummaryPolicyModeDefault,
	"custom":  TieringConfigurationSummaryPolicyModeCustom,
}

// GetTieringConfigurationSummaryPolicyModeEnumValues Enumerates the set of values for TieringConfigurationSummaryPolicyModeEnum
func GetTieringConfigurationSummaryPolicyModeEnumValues() []TieringConfigurationSummaryPolicyModeEnum {
	values := make([]TieringConfigurationSummaryPolicyModeEnum, 0)
	for _, v := range mappingTieringConfigurationSummaryPolicyModeEnum {
		values = append(values, v)
	}
	return values
}

// GetTieringConfigurationSummaryPolicyModeEnumStringValues Enumerates the set of values in String for TieringConfigurationSummaryPolicyModeEnum
func GetTieringConfigurationSummaryPolicyModeEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"CUSTOM",
	}
}

// GetMappingTieringConfigurationSummaryPolicyModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTieringConfigurationSummaryPolicyModeEnum(val string) (TieringConfigurationSummaryPolicyModeEnum, bool) {
	enum, ok := mappingTieringConfigurationSummaryPolicyModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
