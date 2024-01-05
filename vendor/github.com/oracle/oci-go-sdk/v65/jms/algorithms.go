// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// AlgorithmsEnum Enum with underlying type: string
type AlgorithmsEnum string

// Set of constants representing the allowable values for AlgorithmsEnum
const (
	AlgorithmsRsa AlgorithmsEnum = "RSA"
	AlgorithmsDsa AlgorithmsEnum = "DSA"
	AlgorithmsEc  AlgorithmsEnum = "EC"
	AlgorithmsDh  AlgorithmsEnum = "DH"
)

var mappingAlgorithmsEnum = map[string]AlgorithmsEnum{
	"RSA": AlgorithmsRsa,
	"DSA": AlgorithmsDsa,
	"EC":  AlgorithmsEc,
	"DH":  AlgorithmsDh,
}

var mappingAlgorithmsEnumLowerCase = map[string]AlgorithmsEnum{
	"rsa": AlgorithmsRsa,
	"dsa": AlgorithmsDsa,
	"ec":  AlgorithmsEc,
	"dh":  AlgorithmsDh,
}

// GetAlgorithmsEnumValues Enumerates the set of values for AlgorithmsEnum
func GetAlgorithmsEnumValues() []AlgorithmsEnum {
	values := make([]AlgorithmsEnum, 0)
	for _, v := range mappingAlgorithmsEnum {
		values = append(values, v)
	}
	return values
}

// GetAlgorithmsEnumStringValues Enumerates the set of values in String for AlgorithmsEnum
func GetAlgorithmsEnumStringValues() []string {
	return []string{
		"RSA",
		"DSA",
		"EC",
		"DH",
	}
}

// GetMappingAlgorithmsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlgorithmsEnum(val string) (AlgorithmsEnum, bool) {
	enum, ok := mappingAlgorithmsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
