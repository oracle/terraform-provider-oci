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

// UpdateCertificateAuthorityActionTypeEnum Enum with underlying type: string
type UpdateCertificateAuthorityActionTypeEnum string

// Set of constants representing the allowable values for UpdateCertificateAuthorityActionTypeEnum
const (
	UpdateCertificateAuthorityActionTypeUpdateCertificate UpdateCertificateAuthorityActionTypeEnum = "UPDATE_CERTIFICATE"
	UpdateCertificateAuthorityActionTypeGenerateCsr       UpdateCertificateAuthorityActionTypeEnum = "GENERATE_CSR"
)

var mappingUpdateCertificateAuthorityActionTypeEnum = map[string]UpdateCertificateAuthorityActionTypeEnum{
	"UPDATE_CERTIFICATE": UpdateCertificateAuthorityActionTypeUpdateCertificate,
	"GENERATE_CSR":       UpdateCertificateAuthorityActionTypeGenerateCsr,
}

var mappingUpdateCertificateAuthorityActionTypeEnumLowerCase = map[string]UpdateCertificateAuthorityActionTypeEnum{
	"update_certificate": UpdateCertificateAuthorityActionTypeUpdateCertificate,
	"generate_csr":       UpdateCertificateAuthorityActionTypeGenerateCsr,
}

// GetUpdateCertificateAuthorityActionTypeEnumValues Enumerates the set of values for UpdateCertificateAuthorityActionTypeEnum
func GetUpdateCertificateAuthorityActionTypeEnumValues() []UpdateCertificateAuthorityActionTypeEnum {
	values := make([]UpdateCertificateAuthorityActionTypeEnum, 0)
	for _, v := range mappingUpdateCertificateAuthorityActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateCertificateAuthorityActionTypeEnumStringValues Enumerates the set of values in String for UpdateCertificateAuthorityActionTypeEnum
func GetUpdateCertificateAuthorityActionTypeEnumStringValues() []string {
	return []string{
		"UPDATE_CERTIFICATE",
		"GENERATE_CSR",
	}
}

// GetMappingUpdateCertificateAuthorityActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateCertificateAuthorityActionTypeEnum(val string) (UpdateCertificateAuthorityActionTypeEnum, bool) {
	enum, ok := mappingUpdateCertificateAuthorityActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
