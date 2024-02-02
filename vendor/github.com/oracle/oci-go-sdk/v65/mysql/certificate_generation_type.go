// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// CertificateGenerationTypeEnum Enum with underlying type: string
type CertificateGenerationTypeEnum string

// Set of constants representing the allowable values for CertificateGenerationTypeEnum
const (
	CertificateGenerationTypeSystem CertificateGenerationTypeEnum = "SYSTEM"
	CertificateGenerationTypeByoc   CertificateGenerationTypeEnum = "BYOC"
)

var mappingCertificateGenerationTypeEnum = map[string]CertificateGenerationTypeEnum{
	"SYSTEM": CertificateGenerationTypeSystem,
	"BYOC":   CertificateGenerationTypeByoc,
}

var mappingCertificateGenerationTypeEnumLowerCase = map[string]CertificateGenerationTypeEnum{
	"system": CertificateGenerationTypeSystem,
	"byoc":   CertificateGenerationTypeByoc,
}

// GetCertificateGenerationTypeEnumValues Enumerates the set of values for CertificateGenerationTypeEnum
func GetCertificateGenerationTypeEnumValues() []CertificateGenerationTypeEnum {
	values := make([]CertificateGenerationTypeEnum, 0)
	for _, v := range mappingCertificateGenerationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCertificateGenerationTypeEnumStringValues Enumerates the set of values in String for CertificateGenerationTypeEnum
func GetCertificateGenerationTypeEnumStringValues() []string {
	return []string{
		"SYSTEM",
		"BYOC",
	}
}

// GetMappingCertificateGenerationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertificateGenerationTypeEnum(val string) (CertificateGenerationTypeEnum, bool) {
	enum, ok := mappingCertificateGenerationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
