// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// PostInstallationActionsEnum Enum with underlying type: string
type PostInstallationActionsEnum string

// Set of constants representing the allowable values for PostInstallationActionsEnum
const (
	PostInstallationActionsChangeMinimumKeyLengthForEc            PostInstallationActionsEnum = "CHANGE_MINIMUM_KEY_LENGTH_FOR_EC"
	PostInstallationActionsChangeMinimumKeyLengthForDsa           PostInstallationActionsEnum = "CHANGE_MINIMUM_KEY_LENGTH_FOR_DSA"
	PostInstallationActionsChangeMinimumKeyLengthForDiffieHellman PostInstallationActionsEnum = "CHANGE_MINIMUM_KEY_LENGTH_FOR_DIFFIE_HELLMAN"
	PostInstallationActionsChangeMinimumKeyLengthForRsaSignedJars PostInstallationActionsEnum = "CHANGE_MINIMUM_KEY_LENGTH_FOR_RSA_SIGNED_JARS"
	PostInstallationActionsDisableTls                             PostInstallationActionsEnum = "DISABLE_TLS"
	PostInstallationActionsUseOsCacerts                           PostInstallationActionsEnum = "USE_OS_CACERTS"
)

var mappingPostInstallationActionsEnum = map[string]PostInstallationActionsEnum{
	"CHANGE_MINIMUM_KEY_LENGTH_FOR_EC":              PostInstallationActionsChangeMinimumKeyLengthForEc,
	"CHANGE_MINIMUM_KEY_LENGTH_FOR_DSA":             PostInstallationActionsChangeMinimumKeyLengthForDsa,
	"CHANGE_MINIMUM_KEY_LENGTH_FOR_DIFFIE_HELLMAN":  PostInstallationActionsChangeMinimumKeyLengthForDiffieHellman,
	"CHANGE_MINIMUM_KEY_LENGTH_FOR_RSA_SIGNED_JARS": PostInstallationActionsChangeMinimumKeyLengthForRsaSignedJars,
	"DISABLE_TLS":    PostInstallationActionsDisableTls,
	"USE_OS_CACERTS": PostInstallationActionsUseOsCacerts,
}

var mappingPostInstallationActionsEnumLowerCase = map[string]PostInstallationActionsEnum{
	"change_minimum_key_length_for_ec":              PostInstallationActionsChangeMinimumKeyLengthForEc,
	"change_minimum_key_length_for_dsa":             PostInstallationActionsChangeMinimumKeyLengthForDsa,
	"change_minimum_key_length_for_diffie_hellman":  PostInstallationActionsChangeMinimumKeyLengthForDiffieHellman,
	"change_minimum_key_length_for_rsa_signed_jars": PostInstallationActionsChangeMinimumKeyLengthForRsaSignedJars,
	"disable_tls":    PostInstallationActionsDisableTls,
	"use_os_cacerts": PostInstallationActionsUseOsCacerts,
}

// GetPostInstallationActionsEnumValues Enumerates the set of values for PostInstallationActionsEnum
func GetPostInstallationActionsEnumValues() []PostInstallationActionsEnum {
	values := make([]PostInstallationActionsEnum, 0)
	for _, v := range mappingPostInstallationActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetPostInstallationActionsEnumStringValues Enumerates the set of values in String for PostInstallationActionsEnum
func GetPostInstallationActionsEnumStringValues() []string {
	return []string{
		"CHANGE_MINIMUM_KEY_LENGTH_FOR_EC",
		"CHANGE_MINIMUM_KEY_LENGTH_FOR_DSA",
		"CHANGE_MINIMUM_KEY_LENGTH_FOR_DIFFIE_HELLMAN",
		"CHANGE_MINIMUM_KEY_LENGTH_FOR_RSA_SIGNED_JARS",
		"DISABLE_TLS",
		"USE_OS_CACERTS",
	}
}

// GetMappingPostInstallationActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPostInstallationActionsEnum(val string) (PostInstallationActionsEnum, bool) {
	enum, ok := mappingPostInstallationActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
