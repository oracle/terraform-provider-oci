// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"strings"
)

// ValidationTypeEnum Enum with underlying type: string
type ValidationTypeEnum string

// Set of constants representing the allowable values for ValidationTypeEnum
const (
	ValidationTypeDomainValidation       ValidationTypeEnum = "DOMAIN_VALIDATION"
	ValidationTypeOrganizationValidation ValidationTypeEnum = "ORGANIZATION_VALIDATION"
)

var mappingValidationTypeEnum = map[string]ValidationTypeEnum{
	"DOMAIN_VALIDATION":       ValidationTypeDomainValidation,
	"ORGANIZATION_VALIDATION": ValidationTypeOrganizationValidation,
}

var mappingValidationTypeEnumLowerCase = map[string]ValidationTypeEnum{
	"domain_validation":       ValidationTypeDomainValidation,
	"organization_validation": ValidationTypeOrganizationValidation,
}

// GetValidationTypeEnumValues Enumerates the set of values for ValidationTypeEnum
func GetValidationTypeEnumValues() []ValidationTypeEnum {
	values := make([]ValidationTypeEnum, 0)
	for _, v := range mappingValidationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetValidationTypeEnumStringValues Enumerates the set of values in String for ValidationTypeEnum
func GetValidationTypeEnumStringValues() []string {
	return []string{
		"DOMAIN_VALIDATION",
		"ORGANIZATION_VALIDATION",
	}
}

// GetMappingValidationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingValidationTypeEnum(val string) (ValidationTypeEnum, bool) {
	enum, ok := mappingValidationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
