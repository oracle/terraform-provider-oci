// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// ArtifactSigningAlgorithmEnum Enum with underlying type: string
type ArtifactSigningAlgorithmEnum string

// Set of constants representing the allowable values for ArtifactSigningAlgorithmEnum
const (
	ArtifactSigningAlgorithmSha224RsaPkcsPss ArtifactSigningAlgorithmEnum = "SHA_224_RSA_PKCS_PSS"
	ArtifactSigningAlgorithmSha256RsaPkcsPss ArtifactSigningAlgorithmEnum = "SHA_256_RSA_PKCS_PSS"
	ArtifactSigningAlgorithmSha384RsaPkcsPss ArtifactSigningAlgorithmEnum = "SHA_384_RSA_PKCS_PSS"
	ArtifactSigningAlgorithmSha512RsaPkcsPss ArtifactSigningAlgorithmEnum = "SHA_512_RSA_PKCS_PSS"
)

var mappingArtifactSigningAlgorithmEnum = map[string]ArtifactSigningAlgorithmEnum{
	"SHA_224_RSA_PKCS_PSS": ArtifactSigningAlgorithmSha224RsaPkcsPss,
	"SHA_256_RSA_PKCS_PSS": ArtifactSigningAlgorithmSha256RsaPkcsPss,
	"SHA_384_RSA_PKCS_PSS": ArtifactSigningAlgorithmSha384RsaPkcsPss,
	"SHA_512_RSA_PKCS_PSS": ArtifactSigningAlgorithmSha512RsaPkcsPss,
}

var mappingArtifactSigningAlgorithmEnumLowerCase = map[string]ArtifactSigningAlgorithmEnum{
	"sha_224_rsa_pkcs_pss": ArtifactSigningAlgorithmSha224RsaPkcsPss,
	"sha_256_rsa_pkcs_pss": ArtifactSigningAlgorithmSha256RsaPkcsPss,
	"sha_384_rsa_pkcs_pss": ArtifactSigningAlgorithmSha384RsaPkcsPss,
	"sha_512_rsa_pkcs_pss": ArtifactSigningAlgorithmSha512RsaPkcsPss,
}

// GetArtifactSigningAlgorithmEnumValues Enumerates the set of values for ArtifactSigningAlgorithmEnum
func GetArtifactSigningAlgorithmEnumValues() []ArtifactSigningAlgorithmEnum {
	values := make([]ArtifactSigningAlgorithmEnum, 0)
	for _, v := range mappingArtifactSigningAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetArtifactSigningAlgorithmEnumStringValues Enumerates the set of values in String for ArtifactSigningAlgorithmEnum
func GetArtifactSigningAlgorithmEnumStringValues() []string {
	return []string{
		"SHA_224_RSA_PKCS_PSS",
		"SHA_256_RSA_PKCS_PSS",
		"SHA_384_RSA_PKCS_PSS",
		"SHA_512_RSA_PKCS_PSS",
	}
}

// GetMappingArtifactSigningAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArtifactSigningAlgorithmEnum(val string) (ArtifactSigningAlgorithmEnum, bool) {
	enum, ok := mappingArtifactSigningAlgorithmEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
