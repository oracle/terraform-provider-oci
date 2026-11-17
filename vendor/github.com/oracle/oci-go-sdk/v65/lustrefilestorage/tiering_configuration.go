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

// TieringConfiguration Tiering configuration for a Lustre file system. Tiering archives eligible file
// data to Object Storage while keeping files visible at their original Lustre
// paths, and restores archived data transparently when accessed.
// When `policyMode` is `CUSTOM`, this object includes the custom policy rules
// configured for the file system. When `policyMode` is `DEFAULT`, the service
// applies the built-in tiering policy rule, which archives files larger than 1 MB that
// have not been accessed for more than 30 days.
type TieringConfiguration struct {

	// Whether tiering is enabled for the Lustre file system.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Whether the service default policy rule or custom policy rules are used.
	// When set to `DEFAULT`, the service applies the built-in tiering policy rule, which
	// archives files larger than 1 MB that have not been accessed for more than 30 days.
	// When set to `CUSTOM`, the service uses the policy rules configured in `rules`.
	PolicyMode TieringConfigurationPolicyModeEnum `mandatory:"true" json:"policyMode"`

	// Service-computed number of custom tiering policy rules.
	RuleCount *int `mandatory:"true" json:"ruleCount"`

	// Custom tiering policy rules configured for the Lustre file system.
	// This property is only applicable when `policyMode` is `CUSTOM`. When `policyMode`
	// is `DEFAULT`, the service applies the built-in tiering policy rule instead.
	// An enabled custom policy contains 1 to 50 rules. When tiering is disabled,
	// custom policy rules are optional.
	Rules []TieringPolicyRule `mandatory:"false" json:"rules"`
}

func (m TieringConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TieringConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTieringConfigurationPolicyModeEnum(string(m.PolicyMode)); !ok && m.PolicyMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PolicyMode: %s. Supported values are: %s.", m.PolicyMode, strings.Join(GetTieringConfigurationPolicyModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TieringConfigurationPolicyModeEnum Enum with underlying type: string
type TieringConfigurationPolicyModeEnum string

// Set of constants representing the allowable values for TieringConfigurationPolicyModeEnum
const (
	TieringConfigurationPolicyModeDefault TieringConfigurationPolicyModeEnum = "DEFAULT"
	TieringConfigurationPolicyModeCustom  TieringConfigurationPolicyModeEnum = "CUSTOM"
)

var mappingTieringConfigurationPolicyModeEnum = map[string]TieringConfigurationPolicyModeEnum{
	"DEFAULT": TieringConfigurationPolicyModeDefault,
	"CUSTOM":  TieringConfigurationPolicyModeCustom,
}

var mappingTieringConfigurationPolicyModeEnumLowerCase = map[string]TieringConfigurationPolicyModeEnum{
	"default": TieringConfigurationPolicyModeDefault,
	"custom":  TieringConfigurationPolicyModeCustom,
}

// GetTieringConfigurationPolicyModeEnumValues Enumerates the set of values for TieringConfigurationPolicyModeEnum
func GetTieringConfigurationPolicyModeEnumValues() []TieringConfigurationPolicyModeEnum {
	values := make([]TieringConfigurationPolicyModeEnum, 0)
	for _, v := range mappingTieringConfigurationPolicyModeEnum {
		values = append(values, v)
	}
	return values
}

// GetTieringConfigurationPolicyModeEnumStringValues Enumerates the set of values in String for TieringConfigurationPolicyModeEnum
func GetTieringConfigurationPolicyModeEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"CUSTOM",
	}
}

// GetMappingTieringConfigurationPolicyModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTieringConfigurationPolicyModeEnum(val string) (TieringConfigurationPolicyModeEnum, bool) {
	enum, ok := mappingTieringConfigurationPolicyModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
