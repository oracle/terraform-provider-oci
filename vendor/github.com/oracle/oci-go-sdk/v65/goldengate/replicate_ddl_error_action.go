// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ReplicateDdlErrorActionEnum Enum with underlying type: string
type ReplicateDdlErrorActionEnum string

// Set of constants representing the allowable values for ReplicateDdlErrorActionEnum
const (
	ReplicateDdlErrorActionTerminate ReplicateDdlErrorActionEnum = "TERMINATE"
	ReplicateDdlErrorActionDiscard   ReplicateDdlErrorActionEnum = "DISCARD"
	ReplicateDdlErrorActionIgnore    ReplicateDdlErrorActionEnum = "IGNORE"
)

var mappingReplicateDdlErrorActionEnum = map[string]ReplicateDdlErrorActionEnum{
	"TERMINATE": ReplicateDdlErrorActionTerminate,
	"DISCARD":   ReplicateDdlErrorActionDiscard,
	"IGNORE":    ReplicateDdlErrorActionIgnore,
}

var mappingReplicateDdlErrorActionEnumLowerCase = map[string]ReplicateDdlErrorActionEnum{
	"terminate": ReplicateDdlErrorActionTerminate,
	"discard":   ReplicateDdlErrorActionDiscard,
	"ignore":    ReplicateDdlErrorActionIgnore,
}

// GetReplicateDdlErrorActionEnumValues Enumerates the set of values for ReplicateDdlErrorActionEnum
func GetReplicateDdlErrorActionEnumValues() []ReplicateDdlErrorActionEnum {
	values := make([]ReplicateDdlErrorActionEnum, 0)
	for _, v := range mappingReplicateDdlErrorActionEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicateDdlErrorActionEnumStringValues Enumerates the set of values in String for ReplicateDdlErrorActionEnum
func GetReplicateDdlErrorActionEnumStringValues() []string {
	return []string{
		"TERMINATE",
		"DISCARD",
		"IGNORE",
	}
}

// GetMappingReplicateDdlErrorActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicateDdlErrorActionEnum(val string) (ReplicateDdlErrorActionEnum, bool) {
	enum, ok := mappingReplicateDdlErrorActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
