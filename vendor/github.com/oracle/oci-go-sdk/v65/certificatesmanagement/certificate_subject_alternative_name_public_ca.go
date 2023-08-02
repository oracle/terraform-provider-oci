// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CertificateSubjectAlternativeNamePublicCa A subject alternative name for the certificate that binds additional or alternate names to the subject of the certificate. In the certificate, the alternate subject name format is "type:name".
type CertificateSubjectAlternativeNamePublicCa struct {

	// The subject alternative name type. Currently only DNS domain or host names and IP addresses are supported.
	Type CertificateSubjectAlternativeNamePublicCaTypeEnum `mandatory:"true" json:"type"`

	// The subject alternative name.
	Value *string `mandatory:"true" json:"value"`

	// DNS Zone Name or OCID, required only when CertificateConfigType is ISSUED_BY_PUBLIC_CA.
	DnsZoneNameOrId *string `mandatory:"false" json:"dnsZoneNameOrId"`
}

func (m CertificateSubjectAlternativeNamePublicCa) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateSubjectAlternativeNamePublicCa) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCertificateSubjectAlternativeNamePublicCaTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCertificateSubjectAlternativeNamePublicCaTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CertificateSubjectAlternativeNamePublicCaTypeEnum Enum with underlying type: string
type CertificateSubjectAlternativeNamePublicCaTypeEnum string

// Set of constants representing the allowable values for CertificateSubjectAlternativeNamePublicCaTypeEnum
const (
	CertificateSubjectAlternativeNamePublicCaTypeDns CertificateSubjectAlternativeNamePublicCaTypeEnum = "DNS"
	CertificateSubjectAlternativeNamePublicCaTypeIp  CertificateSubjectAlternativeNamePublicCaTypeEnum = "IP"
)

var mappingCertificateSubjectAlternativeNamePublicCaTypeEnum = map[string]CertificateSubjectAlternativeNamePublicCaTypeEnum{
	"DNS": CertificateSubjectAlternativeNamePublicCaTypeDns,
	"IP":  CertificateSubjectAlternativeNamePublicCaTypeIp,
}

var mappingCertificateSubjectAlternativeNamePublicCaTypeEnumLowerCase = map[string]CertificateSubjectAlternativeNamePublicCaTypeEnum{
	"dns": CertificateSubjectAlternativeNamePublicCaTypeDns,
	"ip":  CertificateSubjectAlternativeNamePublicCaTypeIp,
}

// GetCertificateSubjectAlternativeNamePublicCaTypeEnumValues Enumerates the set of values for CertificateSubjectAlternativeNamePublicCaTypeEnum
func GetCertificateSubjectAlternativeNamePublicCaTypeEnumValues() []CertificateSubjectAlternativeNamePublicCaTypeEnum {
	values := make([]CertificateSubjectAlternativeNamePublicCaTypeEnum, 0)
	for _, v := range mappingCertificateSubjectAlternativeNamePublicCaTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCertificateSubjectAlternativeNamePublicCaTypeEnumStringValues Enumerates the set of values in String for CertificateSubjectAlternativeNamePublicCaTypeEnum
func GetCertificateSubjectAlternativeNamePublicCaTypeEnumStringValues() []string {
	return []string{
		"DNS",
		"IP",
	}
}

// GetMappingCertificateSubjectAlternativeNamePublicCaTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertificateSubjectAlternativeNamePublicCaTypeEnum(val string) (CertificateSubjectAlternativeNamePublicCaTypeEnum, bool) {
	enum, ok := mappingCertificateSubjectAlternativeNamePublicCaTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
