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

// CertificateProfileTypeEnum Enum with underlying type: string
type CertificateProfileTypeEnum string

// Set of constants representing the allowable values for CertificateProfileTypeEnum
const (
	CertificateProfileTypeTlsServerOrClient CertificateProfileTypeEnum = "TLS_SERVER_OR_CLIENT"
	CertificateProfileTypeTlsServer         CertificateProfileTypeEnum = "TLS_SERVER"
	CertificateProfileTypeTlsClient         CertificateProfileTypeEnum = "TLS_CLIENT"
	CertificateProfileTypeTlsCodeSign       CertificateProfileTypeEnum = "TLS_CODE_SIGN"
)

var mappingCertificateProfileTypeEnum = map[string]CertificateProfileTypeEnum{
	"TLS_SERVER_OR_CLIENT": CertificateProfileTypeTlsServerOrClient,
	"TLS_SERVER":           CertificateProfileTypeTlsServer,
	"TLS_CLIENT":           CertificateProfileTypeTlsClient,
	"TLS_CODE_SIGN":        CertificateProfileTypeTlsCodeSign,
}

// GetCertificateProfileTypeEnumValues Enumerates the set of values for CertificateProfileTypeEnum
func GetCertificateProfileTypeEnumValues() []CertificateProfileTypeEnum {
	values := make([]CertificateProfileTypeEnum, 0)
	for _, v := range mappingCertificateProfileTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCertificateProfileTypeEnumStringValues Enumerates the set of values in String for CertificateProfileTypeEnum
func GetCertificateProfileTypeEnumStringValues() []string {
	return []string{
		"TLS_SERVER_OR_CLIENT",
		"TLS_SERVER",
		"TLS_CLIENT",
		"TLS_CODE_SIGN",
	}
}

// GetMappingCertificateProfileTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertificateProfileTypeEnum(val string) (CertificateProfileTypeEnum, bool) {
	mappingCertificateProfileTypeEnumIgnoreCase := make(map[string]CertificateProfileTypeEnum)
	for k, v := range mappingCertificateProfileTypeEnum {
		mappingCertificateProfileTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCertificateProfileTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
