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

// DistributedDbTransportTypeEnum Enum with underlying type: string
type DistributedDbTransportTypeEnum string

// Set of constants representing the allowable values for DistributedDbTransportTypeEnum
const (
	DistributedDbTransportTypeSync     DistributedDbTransportTypeEnum = "SYNC"
	DistributedDbTransportTypeAsync    DistributedDbTransportTypeEnum = "ASYNC"
	DistributedDbTransportTypeFastsync DistributedDbTransportTypeEnum = "FASTSYNC"
)

var mappingDistributedDbTransportTypeEnum = map[string]DistributedDbTransportTypeEnum{
	"SYNC":     DistributedDbTransportTypeSync,
	"ASYNC":    DistributedDbTransportTypeAsync,
	"FASTSYNC": DistributedDbTransportTypeFastsync,
}

var mappingDistributedDbTransportTypeEnumLowerCase = map[string]DistributedDbTransportTypeEnum{
	"sync":     DistributedDbTransportTypeSync,
	"async":    DistributedDbTransportTypeAsync,
	"fastsync": DistributedDbTransportTypeFastsync,
}

// GetDistributedDbTransportTypeEnumValues Enumerates the set of values for DistributedDbTransportTypeEnum
func GetDistributedDbTransportTypeEnumValues() []DistributedDbTransportTypeEnum {
	values := make([]DistributedDbTransportTypeEnum, 0)
	for _, v := range mappingDistributedDbTransportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDbTransportTypeEnumStringValues Enumerates the set of values in String for DistributedDbTransportTypeEnum
func GetDistributedDbTransportTypeEnumStringValues() []string {
	return []string{
		"SYNC",
		"ASYNC",
		"FASTSYNC",
	}
}

// GetMappingDistributedDbTransportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDbTransportTypeEnum(val string) (DistributedDbTransportTypeEnum, bool) {
	enum, ok := mappingDistributedDbTransportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
