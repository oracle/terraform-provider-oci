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

// PublicCertificateValidityEnum Enum with underlying type: string
type PublicCertificateValidityEnum string

// Set of constants representing the allowable values for PublicCertificateValidityEnum
const (
	PublicCertificateValidityDays90  PublicCertificateValidityEnum = "DAYS_90"
	PublicCertificateValidityDays180 PublicCertificateValidityEnum = "DAYS_180"
	PublicCertificateValidityDays397 PublicCertificateValidityEnum = "DAYS_397"
)

var mappingPublicCertificateValidityEnum = map[string]PublicCertificateValidityEnum{
	"DAYS_90":  PublicCertificateValidityDays90,
	"DAYS_180": PublicCertificateValidityDays180,
	"DAYS_397": PublicCertificateValidityDays397,
}

var mappingPublicCertificateValidityEnumLowerCase = map[string]PublicCertificateValidityEnum{
	"days_90":  PublicCertificateValidityDays90,
	"days_180": PublicCertificateValidityDays180,
	"days_397": PublicCertificateValidityDays397,
}

// GetPublicCertificateValidityEnumValues Enumerates the set of values for PublicCertificateValidityEnum
func GetPublicCertificateValidityEnumValues() []PublicCertificateValidityEnum {
	values := make([]PublicCertificateValidityEnum, 0)
	for _, v := range mappingPublicCertificateValidityEnum {
		values = append(values, v)
	}
	return values
}

// GetPublicCertificateValidityEnumStringValues Enumerates the set of values in String for PublicCertificateValidityEnum
func GetPublicCertificateValidityEnumStringValues() []string {
	return []string{
		"DAYS_90",
		"DAYS_180",
		"DAYS_397",
	}
}

// GetMappingPublicCertificateValidityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPublicCertificateValidityEnum(val string) (PublicCertificateValidityEnum, bool) {
	enum, ok := mappingPublicCertificateValidityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
