// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"strings"
)

// CccModeEnum Enum with underlying type: string
type CccModeEnum string

// Set of constants representing the allowable values for CccModeEnum
const (
	CccModeSearchOnly           CccModeEnum = "SEARCH_ONLY"
	CccModeReplicationOnly      CccModeEnum = "REPLICATION_ONLY"
	CccModeSearchAndReplication CccModeEnum = "SEARCH_AND_REPLICATION"
)

var mappingCccModeEnum = map[string]CccModeEnum{
	"SEARCH_ONLY":            CccModeSearchOnly,
	"REPLICATION_ONLY":       CccModeReplicationOnly,
	"SEARCH_AND_REPLICATION": CccModeSearchAndReplication,
}

var mappingCccModeEnumLowerCase = map[string]CccModeEnum{
	"search_only":            CccModeSearchOnly,
	"replication_only":       CccModeReplicationOnly,
	"search_and_replication": CccModeSearchAndReplication,
}

// GetCccModeEnumValues Enumerates the set of values for CccModeEnum
func GetCccModeEnumValues() []CccModeEnum {
	values := make([]CccModeEnum, 0)
	for _, v := range mappingCccModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCccModeEnumStringValues Enumerates the set of values in String for CccModeEnum
func GetCccModeEnumStringValues() []string {
	return []string{
		"SEARCH_ONLY",
		"REPLICATION_ONLY",
		"SEARCH_AND_REPLICATION",
	}
}

// GetMappingCccModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccModeEnum(val string) (CccModeEnum, bool) {
	enum, ok := mappingCccModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
