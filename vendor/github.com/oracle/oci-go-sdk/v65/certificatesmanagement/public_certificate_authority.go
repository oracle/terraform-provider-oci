// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// PublicCertificateAuthorityEnum Enum with underlying type: string
type PublicCertificateAuthorityEnum string

// Set of constants representing the allowable values for PublicCertificateAuthorityEnum
const (
	PublicCertificateAuthorityDigicert    PublicCertificateAuthorityEnum = "DIGICERT"
	PublicCertificateAuthorityLetsEncrypt PublicCertificateAuthorityEnum = "LETS_ENCRYPT"
)

var mappingPublicCertificateAuthorityEnum = map[string]PublicCertificateAuthorityEnum{
	"DIGICERT":     PublicCertificateAuthorityDigicert,
	"LETS_ENCRYPT": PublicCertificateAuthorityLetsEncrypt,
}

var mappingPublicCertificateAuthorityEnumLowerCase = map[string]PublicCertificateAuthorityEnum{
	"digicert":     PublicCertificateAuthorityDigicert,
	"lets_encrypt": PublicCertificateAuthorityLetsEncrypt,
}

// GetPublicCertificateAuthorityEnumValues Enumerates the set of values for PublicCertificateAuthorityEnum
func GetPublicCertificateAuthorityEnumValues() []PublicCertificateAuthorityEnum {
	values := make([]PublicCertificateAuthorityEnum, 0)
	for _, v := range mappingPublicCertificateAuthorityEnum {
		values = append(values, v)
	}
	return values
}

// GetPublicCertificateAuthorityEnumStringValues Enumerates the set of values in String for PublicCertificateAuthorityEnum
func GetPublicCertificateAuthorityEnumStringValues() []string {
	return []string{
		"DIGICERT",
		"LETS_ENCRYPT",
	}
}

// GetMappingPublicCertificateAuthorityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPublicCertificateAuthorityEnum(val string) (PublicCertificateAuthorityEnum, bool) {
	enum, ok := mappingPublicCertificateAuthorityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
