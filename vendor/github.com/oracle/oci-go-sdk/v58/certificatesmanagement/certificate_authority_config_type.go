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

// CertificateAuthorityConfigTypeEnum Enum with underlying type: string
type CertificateAuthorityConfigTypeEnum string

// Set of constants representing the allowable values for CertificateAuthorityConfigTypeEnum
const (
	CertificateAuthorityConfigTypeRootCaGeneratedInternally       CertificateAuthorityConfigTypeEnum = "ROOT_CA_GENERATED_INTERNALLY"
	CertificateAuthorityConfigTypeSubordinateCaIssuedByInternalCa CertificateAuthorityConfigTypeEnum = "SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA"
)

var mappingCertificateAuthorityConfigTypeEnum = map[string]CertificateAuthorityConfigTypeEnum{
	"ROOT_CA_GENERATED_INTERNALLY":         CertificateAuthorityConfigTypeRootCaGeneratedInternally,
	"SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA": CertificateAuthorityConfigTypeSubordinateCaIssuedByInternalCa,
}

// GetCertificateAuthorityConfigTypeEnumValues Enumerates the set of values for CertificateAuthorityConfigTypeEnum
func GetCertificateAuthorityConfigTypeEnumValues() []CertificateAuthorityConfigTypeEnum {
	values := make([]CertificateAuthorityConfigTypeEnum, 0)
	for _, v := range mappingCertificateAuthorityConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCertificateAuthorityConfigTypeEnumStringValues Enumerates the set of values in String for CertificateAuthorityConfigTypeEnum
func GetCertificateAuthorityConfigTypeEnumStringValues() []string {
	return []string{
		"ROOT_CA_GENERATED_INTERNALLY",
		"SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA",
	}
}

// GetMappingCertificateAuthorityConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertificateAuthorityConfigTypeEnum(val string) (CertificateAuthorityConfigTypeEnum, bool) {
	mappingCertificateAuthorityConfigTypeEnumIgnoreCase := make(map[string]CertificateAuthorityConfigTypeEnum)
	for k, v := range mappingCertificateAuthorityConfigTypeEnum {
		mappingCertificateAuthorityConfigTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCertificateAuthorityConfigTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
