// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// SignatureAlgorithmEnum Enum with underlying type: string
type SignatureAlgorithmEnum string

// Set of constants representing the allowable values for SignatureAlgorithmEnum
const (
	SignatureAlgorithmSha256WithRsa   SignatureAlgorithmEnum = "SHA256_WITH_RSA"
	SignatureAlgorithmSha384WithRsa   SignatureAlgorithmEnum = "SHA384_WITH_RSA"
	SignatureAlgorithmSha512WithRsa   SignatureAlgorithmEnum = "SHA512_WITH_RSA"
	SignatureAlgorithmSha256WithEcdsa SignatureAlgorithmEnum = "SHA256_WITH_ECDSA"
	SignatureAlgorithmSha384WithEcdsa SignatureAlgorithmEnum = "SHA384_WITH_ECDSA"
	SignatureAlgorithmSha512WithEcdsa SignatureAlgorithmEnum = "SHA512_WITH_ECDSA"
)

var mappingSignatureAlgorithmEnum = map[string]SignatureAlgorithmEnum{
	"SHA256_WITH_RSA":   SignatureAlgorithmSha256WithRsa,
	"SHA384_WITH_RSA":   SignatureAlgorithmSha384WithRsa,
	"SHA512_WITH_RSA":   SignatureAlgorithmSha512WithRsa,
	"SHA256_WITH_ECDSA": SignatureAlgorithmSha256WithEcdsa,
	"SHA384_WITH_ECDSA": SignatureAlgorithmSha384WithEcdsa,
	"SHA512_WITH_ECDSA": SignatureAlgorithmSha512WithEcdsa,
}

var mappingSignatureAlgorithmEnumLowerCase = map[string]SignatureAlgorithmEnum{
	"sha256_with_rsa":   SignatureAlgorithmSha256WithRsa,
	"sha384_with_rsa":   SignatureAlgorithmSha384WithRsa,
	"sha512_with_rsa":   SignatureAlgorithmSha512WithRsa,
	"sha256_with_ecdsa": SignatureAlgorithmSha256WithEcdsa,
	"sha384_with_ecdsa": SignatureAlgorithmSha384WithEcdsa,
	"sha512_with_ecdsa": SignatureAlgorithmSha512WithEcdsa,
}

// GetSignatureAlgorithmEnumValues Enumerates the set of values for SignatureAlgorithmEnum
func GetSignatureAlgorithmEnumValues() []SignatureAlgorithmEnum {
	values := make([]SignatureAlgorithmEnum, 0)
	for _, v := range mappingSignatureAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetSignatureAlgorithmEnumStringValues Enumerates the set of values in String for SignatureAlgorithmEnum
func GetSignatureAlgorithmEnumStringValues() []string {
	return []string{
		"SHA256_WITH_RSA",
		"SHA384_WITH_RSA",
		"SHA512_WITH_RSA",
		"SHA256_WITH_ECDSA",
		"SHA384_WITH_ECDSA",
		"SHA512_WITH_ECDSA",
	}
}

// GetMappingSignatureAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSignatureAlgorithmEnum(val string) (SignatureAlgorithmEnum, bool) {
	enum, ok := mappingSignatureAlgorithmEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
