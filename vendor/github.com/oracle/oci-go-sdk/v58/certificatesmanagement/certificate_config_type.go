// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// CertificateConfigTypeEnum Enum with underlying type: string
type CertificateConfigTypeEnum string

// Set of constants representing the allowable values for CertificateConfigTypeEnum
const (
	CertificateConfigTypeIssuedByInternalCa                  CertificateConfigTypeEnum = "ISSUED_BY_INTERNAL_CA"
	CertificateConfigTypeManagedExternallyIssuedByInternalCa CertificateConfigTypeEnum = "MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA"
	CertificateConfigTypeImported                            CertificateConfigTypeEnum = "IMPORTED"
)

var mappingCertificateConfigTypeEnum = map[string]CertificateConfigTypeEnum{
	"ISSUED_BY_INTERNAL_CA":                    CertificateConfigTypeIssuedByInternalCa,
	"MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA": CertificateConfigTypeManagedExternallyIssuedByInternalCa,
	"IMPORTED": CertificateConfigTypeImported,
}

// GetCertificateConfigTypeEnumValues Enumerates the set of values for CertificateConfigTypeEnum
func GetCertificateConfigTypeEnumValues() []CertificateConfigTypeEnum {
	values := make([]CertificateConfigTypeEnum, 0)
	for _, v := range mappingCertificateConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCertificateConfigTypeEnumStringValues Enumerates the set of values in String for CertificateConfigTypeEnum
func GetCertificateConfigTypeEnumStringValues() []string {
	return []string{
		"ISSUED_BY_INTERNAL_CA",
		"MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA",
		"IMPORTED",
	}
}

// GetMappingCertificateConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertificateConfigTypeEnum(val string) (CertificateConfigTypeEnum, bool) {
	mappingCertificateConfigTypeEnumIgnoreCase := make(map[string]CertificateConfigTypeEnum)
	for k, v := range mappingCertificateConfigTypeEnum {
		mappingCertificateConfigTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCertificateConfigTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
