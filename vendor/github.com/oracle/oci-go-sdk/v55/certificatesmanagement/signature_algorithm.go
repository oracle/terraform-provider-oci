// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

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

var mappingSignatureAlgorithm = map[string]SignatureAlgorithmEnum{
	"SHA256_WITH_RSA":   SignatureAlgorithmSha256WithRsa,
	"SHA384_WITH_RSA":   SignatureAlgorithmSha384WithRsa,
	"SHA512_WITH_RSA":   SignatureAlgorithmSha512WithRsa,
	"SHA256_WITH_ECDSA": SignatureAlgorithmSha256WithEcdsa,
	"SHA384_WITH_ECDSA": SignatureAlgorithmSha384WithEcdsa,
	"SHA512_WITH_ECDSA": SignatureAlgorithmSha512WithEcdsa,
}

// GetSignatureAlgorithmEnumValues Enumerates the set of values for SignatureAlgorithmEnum
func GetSignatureAlgorithmEnumValues() []SignatureAlgorithmEnum {
	values := make([]SignatureAlgorithmEnum, 0)
	for _, v := range mappingSignatureAlgorithm {
		values = append(values, v)
	}
	return values
}
