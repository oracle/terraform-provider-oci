// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// InitialLoadActionEnum Enum with underlying type: string
type InitialLoadActionEnum string

// Set of constants representing the allowable values for InitialLoadActionEnum
const (
	InitialLoadActionTruncate InitialLoadActionEnum = "TRUNCATE"
	InitialLoadActionReplace  InitialLoadActionEnum = "REPLACE"
	InitialLoadActionAppend   InitialLoadActionEnum = "APPEND"
	InitialLoadActionSkip     InitialLoadActionEnum = "SKIP"
)

var mappingInitialLoadActionEnum = map[string]InitialLoadActionEnum{
	"TRUNCATE": InitialLoadActionTruncate,
	"REPLACE":  InitialLoadActionReplace,
	"APPEND":   InitialLoadActionAppend,
	"SKIP":     InitialLoadActionSkip,
}

var mappingInitialLoadActionEnumLowerCase = map[string]InitialLoadActionEnum{
	"truncate": InitialLoadActionTruncate,
	"replace":  InitialLoadActionReplace,
	"append":   InitialLoadActionAppend,
	"skip":     InitialLoadActionSkip,
}

// GetInitialLoadActionEnumValues Enumerates the set of values for InitialLoadActionEnum
func GetInitialLoadActionEnumValues() []InitialLoadActionEnum {
	values := make([]InitialLoadActionEnum, 0)
	for _, v := range mappingInitialLoadActionEnum {
		values = append(values, v)
	}
	return values
}

// GetInitialLoadActionEnumStringValues Enumerates the set of values in String for InitialLoadActionEnum
func GetInitialLoadActionEnumStringValues() []string {
	return []string{
		"TRUNCATE",
		"REPLACE",
		"APPEND",
		"SKIP",
	}
}

// GetMappingInitialLoadActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInitialLoadActionEnum(val string) (InitialLoadActionEnum, bool) {
	enum, ok := mappingInitialLoadActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
