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

// UnifiedAuditPolicyConfig The unified audit policy related configurations.
type UnifiedAuditPolicyConfig struct {

	// Specifies whether the Data Safe service account on the target database should be excluded in the unified audit policy.
	ExcludeDatasafeUser UnifiedAuditPolicyConfigExcludeDatasafeUserEnum `mandatory:"true" json:"excludeDatasafeUser"`
}

func (m UnifiedAuditPolicyConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAuditPolicyConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUnifiedAuditPolicyConfigExcludeDatasafeUserEnum(string(m.ExcludeDatasafeUser)); !ok && m.ExcludeDatasafeUser != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExcludeDatasafeUser: %s. Supported values are: %s.", m.ExcludeDatasafeUser, strings.Join(GetUnifiedAuditPolicyConfigExcludeDatasafeUserEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnifiedAuditPolicyConfigExcludeDatasafeUserEnum Enum with underlying type: string
type UnifiedAuditPolicyConfigExcludeDatasafeUserEnum string

// Set of constants representing the allowable values for UnifiedAuditPolicyConfigExcludeDatasafeUserEnum
const (
	UnifiedAuditPolicyConfigExcludeDatasafeUserEnabled  UnifiedAuditPolicyConfigExcludeDatasafeUserEnum = "ENABLED"
	UnifiedAuditPolicyConfigExcludeDatasafeUserDisabled UnifiedAuditPolicyConfigExcludeDatasafeUserEnum = "DISABLED"
)

var mappingUnifiedAuditPolicyConfigExcludeDatasafeUserEnum = map[string]UnifiedAuditPolicyConfigExcludeDatasafeUserEnum{
	"ENABLED":  UnifiedAuditPolicyConfigExcludeDatasafeUserEnabled,
	"DISABLED": UnifiedAuditPolicyConfigExcludeDatasafeUserDisabled,
}

var mappingUnifiedAuditPolicyConfigExcludeDatasafeUserEnumLowerCase = map[string]UnifiedAuditPolicyConfigExcludeDatasafeUserEnum{
	"enabled":  UnifiedAuditPolicyConfigExcludeDatasafeUserEnabled,
	"disabled": UnifiedAuditPolicyConfigExcludeDatasafeUserDisabled,
}

// GetUnifiedAuditPolicyConfigExcludeDatasafeUserEnumValues Enumerates the set of values for UnifiedAuditPolicyConfigExcludeDatasafeUserEnum
func GetUnifiedAuditPolicyConfigExcludeDatasafeUserEnumValues() []UnifiedAuditPolicyConfigExcludeDatasafeUserEnum {
	values := make([]UnifiedAuditPolicyConfigExcludeDatasafeUserEnum, 0)
	for _, v := range mappingUnifiedAuditPolicyConfigExcludeDatasafeUserEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAuditPolicyConfigExcludeDatasafeUserEnumStringValues Enumerates the set of values in String for UnifiedAuditPolicyConfigExcludeDatasafeUserEnum
func GetUnifiedAuditPolicyConfigExcludeDatasafeUserEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingUnifiedAuditPolicyConfigExcludeDatasafeUserEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAuditPolicyConfigExcludeDatasafeUserEnum(val string) (UnifiedAuditPolicyConfigExcludeDatasafeUserEnum, bool) {
	enum, ok := mappingUnifiedAuditPolicyConfigExcludeDatasafeUserEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
