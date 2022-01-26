// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CertificateSubjectAlternativeName A subject alternative name for the certificate that binds additional or alternate names to the subject of the certificate. In the certificate, the alternate subject name format is "type:name".
type CertificateSubjectAlternativeName struct {

	// The subject alternative name type. Currently only DNS domain or host names and IP addresses are supported.
	Type CertificateSubjectAlternativeNameTypeEnum `mandatory:"true" json:"type"`

	// The subject alternative name.
	Value *string `mandatory:"true" json:"value"`
}

func (m CertificateSubjectAlternativeName) String() string {
	return common.PointerString(m)
}

// CertificateSubjectAlternativeNameTypeEnum Enum with underlying type: string
type CertificateSubjectAlternativeNameTypeEnum string

// Set of constants representing the allowable values for CertificateSubjectAlternativeNameTypeEnum
const (
	CertificateSubjectAlternativeNameTypeDns CertificateSubjectAlternativeNameTypeEnum = "DNS"
	CertificateSubjectAlternativeNameTypeIp  CertificateSubjectAlternativeNameTypeEnum = "IP"
)

var mappingCertificateSubjectAlternativeNameType = map[string]CertificateSubjectAlternativeNameTypeEnum{
	"DNS": CertificateSubjectAlternativeNameTypeDns,
	"IP":  CertificateSubjectAlternativeNameTypeIp,
}

// GetCertificateSubjectAlternativeNameTypeEnumValues Enumerates the set of values for CertificateSubjectAlternativeNameTypeEnum
func GetCertificateSubjectAlternativeNameTypeEnumValues() []CertificateSubjectAlternativeNameTypeEnum {
	values := make([]CertificateSubjectAlternativeNameTypeEnum, 0)
	for _, v := range mappingCertificateSubjectAlternativeNameType {
		values = append(values, v)
	}
	return values
}
