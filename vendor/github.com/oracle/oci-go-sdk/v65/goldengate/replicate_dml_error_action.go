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

// ReplicateDmlErrorActionEnum Enum with underlying type: string
type ReplicateDmlErrorActionEnum string

// Set of constants representing the allowable values for ReplicateDmlErrorActionEnum
const (
	ReplicateDmlErrorActionTerminate ReplicateDmlErrorActionEnum = "TERMINATE"
	ReplicateDmlErrorActionDiscard   ReplicateDmlErrorActionEnum = "DISCARD"
	ReplicateDmlErrorActionIgnore    ReplicateDmlErrorActionEnum = "IGNORE"
)

var mappingReplicateDmlErrorActionEnum = map[string]ReplicateDmlErrorActionEnum{
	"TERMINATE": ReplicateDmlErrorActionTerminate,
	"DISCARD":   ReplicateDmlErrorActionDiscard,
	"IGNORE":    ReplicateDmlErrorActionIgnore,
}

var mappingReplicateDmlErrorActionEnumLowerCase = map[string]ReplicateDmlErrorActionEnum{
	"terminate": ReplicateDmlErrorActionTerminate,
	"discard":   ReplicateDmlErrorActionDiscard,
	"ignore":    ReplicateDmlErrorActionIgnore,
}

// GetReplicateDmlErrorActionEnumValues Enumerates the set of values for ReplicateDmlErrorActionEnum
func GetReplicateDmlErrorActionEnumValues() []ReplicateDmlErrorActionEnum {
	values := make([]ReplicateDmlErrorActionEnum, 0)
	for _, v := range mappingReplicateDmlErrorActionEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicateDmlErrorActionEnumStringValues Enumerates the set of values in String for ReplicateDmlErrorActionEnum
func GetReplicateDmlErrorActionEnumStringValues() []string {
	return []string{
		"TERMINATE",
		"DISCARD",
		"IGNORE",
	}
}

// GetMappingReplicateDmlErrorActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicateDmlErrorActionEnum(val string) (ReplicateDmlErrorActionEnum, bool) {
	enum, ok := mappingReplicateDmlErrorActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
