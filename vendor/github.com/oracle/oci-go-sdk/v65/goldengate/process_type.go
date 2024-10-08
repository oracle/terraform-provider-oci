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

// ProcessTypeEnum Enum with underlying type: string
type ProcessTypeEnum string

// Set of constants representing the allowable values for ProcessTypeEnum
const (
	ProcessTypeExtract  ProcessTypeEnum = "EXTRACT"
	ProcessTypeReplicat ProcessTypeEnum = "REPLICAT"
)

var mappingProcessTypeEnum = map[string]ProcessTypeEnum{
	"EXTRACT":  ProcessTypeExtract,
	"REPLICAT": ProcessTypeReplicat,
}

var mappingProcessTypeEnumLowerCase = map[string]ProcessTypeEnum{
	"extract":  ProcessTypeExtract,
	"replicat": ProcessTypeReplicat,
}

// GetProcessTypeEnumValues Enumerates the set of values for ProcessTypeEnum
func GetProcessTypeEnumValues() []ProcessTypeEnum {
	values := make([]ProcessTypeEnum, 0)
	for _, v := range mappingProcessTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetProcessTypeEnumStringValues Enumerates the set of values in String for ProcessTypeEnum
func GetProcessTypeEnumStringValues() []string {
	return []string{
		"EXTRACT",
		"REPLICAT",
	}
}

// GetMappingProcessTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProcessTypeEnum(val string) (ProcessTypeEnum, bool) {
	enum, ok := mappingProcessTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
