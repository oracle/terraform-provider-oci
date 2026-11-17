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

// TieringConfigurationDetails Tiering configuration details for a Lustre file system. Tiering archives eligible file
// data to Object Storage while keeping files visible at their original Lustre
// paths and restoring archived data transparently when accessed.
// When tiering is enabled and `policyMode` is `CUSTOM`, provide one or more policy rules.
// When tiering is disabled, custom policy rules are optional. If `policyMode` is `DEFAULT`,
// the service applies the built-in tiering policy rule, which archives files larger than 1 MB
// that have not been accessed for more than 30 days.
type TieringConfigurationDetails struct {

	// Whether tiering is enabled for the Lustre file system. If unspecified, defaults to `false`.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Whether the service default policy rule or custom policy rules are used. If unspecified,
	// defaults to `DEFAULT`.
	// When set to `DEFAULT`, the service applies the built-in tiering policy rule, which
	// archives files larger than 1 MB that have not been accessed for more than 30 days.
	// When set to `CUSTOM`, the service uses the policy rules specified in `rules`.
	PolicyMode TieringConfigurationDetailsPolicyModeEnum `mandatory:"false" json:"policyMode,omitempty"`

	// Custom tiering policy rules. Specify this only when `policyMode` is `CUSTOM`.
	// When `policyMode` is `DEFAULT`, omit this property; the service applies the
	// built-in tiering policy rule instead.
	// An enabled custom policy must contain 1 to 50 rules. When tiering is disabled,
	// custom policy rules are optional.
	Rules []TieringPolicyRule `mandatory:"false" json:"rules"`
}

func (m TieringConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TieringConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTieringConfigurationDetailsPolicyModeEnum(string(m.PolicyMode)); !ok && m.PolicyMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PolicyMode: %s. Supported values are: %s.", m.PolicyMode, strings.Join(GetTieringConfigurationDetailsPolicyModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TieringConfigurationDetailsPolicyModeEnum Enum with underlying type: string
type TieringConfigurationDetailsPolicyModeEnum string

// Set of constants representing the allowable values for TieringConfigurationDetailsPolicyModeEnum
const (
	TieringConfigurationDetailsPolicyModeDefault TieringConfigurationDetailsPolicyModeEnum = "DEFAULT"
	TieringConfigurationDetailsPolicyModeCustom  TieringConfigurationDetailsPolicyModeEnum = "CUSTOM"
)

var mappingTieringConfigurationDetailsPolicyModeEnum = map[string]TieringConfigurationDetailsPolicyModeEnum{
	"DEFAULT": TieringConfigurationDetailsPolicyModeDefault,
	"CUSTOM":  TieringConfigurationDetailsPolicyModeCustom,
}

var mappingTieringConfigurationDetailsPolicyModeEnumLowerCase = map[string]TieringConfigurationDetailsPolicyModeEnum{
	"default": TieringConfigurationDetailsPolicyModeDefault,
	"custom":  TieringConfigurationDetailsPolicyModeCustom,
}

// GetTieringConfigurationDetailsPolicyModeEnumValues Enumerates the set of values for TieringConfigurationDetailsPolicyModeEnum
func GetTieringConfigurationDetailsPolicyModeEnumValues() []TieringConfigurationDetailsPolicyModeEnum {
	values := make([]TieringConfigurationDetailsPolicyModeEnum, 0)
	for _, v := range mappingTieringConfigurationDetailsPolicyModeEnum {
		values = append(values, v)
	}
	return values
}

// GetTieringConfigurationDetailsPolicyModeEnumStringValues Enumerates the set of values in String for TieringConfigurationDetailsPolicyModeEnum
func GetTieringConfigurationDetailsPolicyModeEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"CUSTOM",
	}
}

// GetMappingTieringConfigurationDetailsPolicyModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTieringConfigurationDetailsPolicyModeEnum(val string) (TieringConfigurationDetailsPolicyModeEnum, bool) {
	enum, ok := mappingTieringConfigurationDetailsPolicyModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
