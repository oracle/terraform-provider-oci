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

// UnifiedAuditPolicyConfigDetails The unified audit policy related configurations.
type UnifiedAuditPolicyConfigDetails struct {

	// Specifies whether the Data Safe service account on the target database should be excluded in the unified audit policy.
	ExcludeDatasafeUser UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum `mandatory:"false" json:"excludeDatasafeUser,omitempty"`
}

func (m UnifiedAuditPolicyConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAuditPolicyConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum(string(m.ExcludeDatasafeUser)); !ok && m.ExcludeDatasafeUser != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExcludeDatasafeUser: %s. Supported values are: %s.", m.ExcludeDatasafeUser, strings.Join(GetUnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum Enum with underlying type: string
type UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum string

// Set of constants representing the allowable values for UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum
const (
	UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnabled  UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum = "ENABLED"
	UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserDisabled UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum = "DISABLED"
)

var mappingUnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum = map[string]UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum{
	"ENABLED":  UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnabled,
	"DISABLED": UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserDisabled,
}

var mappingUnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnumLowerCase = map[string]UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum{
	"enabled":  UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnabled,
	"disabled": UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserDisabled,
}

// GetUnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnumValues Enumerates the set of values for UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum
func GetUnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnumValues() []UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum {
	values := make([]UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum, 0)
	for _, v := range mappingUnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnumStringValues Enumerates the set of values in String for UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum
func GetUnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingUnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum(val string) (UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum, bool) {
	enum, ok := mappingUnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
