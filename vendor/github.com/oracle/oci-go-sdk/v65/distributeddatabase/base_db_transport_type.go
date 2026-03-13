// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"strings"
)

// BaseDbTransportTypeEnum Enum with underlying type: string
type BaseDbTransportTypeEnum string

// Set of constants representing the allowable values for BaseDbTransportTypeEnum
const (
	BaseDbTransportTypeSync     BaseDbTransportTypeEnum = "SYNC"
	BaseDbTransportTypeAsync    BaseDbTransportTypeEnum = "ASYNC"
	BaseDbTransportTypeFastsync BaseDbTransportTypeEnum = "FASTSYNC"
)

var mappingBaseDbTransportTypeEnum = map[string]BaseDbTransportTypeEnum{
	"SYNC":     BaseDbTransportTypeSync,
	"ASYNC":    BaseDbTransportTypeAsync,
	"FASTSYNC": BaseDbTransportTypeFastsync,
}

var mappingBaseDbTransportTypeEnumLowerCase = map[string]BaseDbTransportTypeEnum{
	"sync":     BaseDbTransportTypeSync,
	"async":    BaseDbTransportTypeAsync,
	"fastsync": BaseDbTransportTypeFastsync,
}

// GetBaseDbTransportTypeEnumValues Enumerates the set of values for BaseDbTransportTypeEnum
func GetBaseDbTransportTypeEnumValues() []BaseDbTransportTypeEnum {
	values := make([]BaseDbTransportTypeEnum, 0)
	for _, v := range mappingBaseDbTransportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseDbTransportTypeEnumStringValues Enumerates the set of values in String for BaseDbTransportTypeEnum
func GetBaseDbTransportTypeEnumStringValues() []string {
	return []string{
		"SYNC",
		"ASYNC",
		"FASTSYNC",
	}
}

// GetMappingBaseDbTransportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseDbTransportTypeEnum(val string) (BaseDbTransportTypeEnum, bool) {
	enum, ok := mappingBaseDbTransportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
