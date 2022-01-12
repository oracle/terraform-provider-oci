// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

// CertificateAuthorityConfigTypeEnum Enum with underlying type: string
type CertificateAuthorityConfigTypeEnum string

// Set of constants representing the allowable values for CertificateAuthorityConfigTypeEnum
const (
	CertificateAuthorityConfigTypeRootCaGeneratedInternally       CertificateAuthorityConfigTypeEnum = "ROOT_CA_GENERATED_INTERNALLY"
	CertificateAuthorityConfigTypeSubordinateCaIssuedByInternalCa CertificateAuthorityConfigTypeEnum = "SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA"
)

var mappingCertificateAuthorityConfigType = map[string]CertificateAuthorityConfigTypeEnum{
	"ROOT_CA_GENERATED_INTERNALLY":         CertificateAuthorityConfigTypeRootCaGeneratedInternally,
	"SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA": CertificateAuthorityConfigTypeSubordinateCaIssuedByInternalCa,
}

// GetCertificateAuthorityConfigTypeEnumValues Enumerates the set of values for CertificateAuthorityConfigTypeEnum
func GetCertificateAuthorityConfigTypeEnumValues() []CertificateAuthorityConfigTypeEnum {
	values := make([]CertificateAuthorityConfigTypeEnum, 0)
	for _, v := range mappingCertificateAuthorityConfigType {
		values = append(values, v)
	}
	return values
}
