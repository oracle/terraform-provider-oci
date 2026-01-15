// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"strings"
)

// OkeClusterImageReplicationEnum Enum with underlying type: string
type OkeClusterImageReplicationEnum string

// Set of constants representing the allowable values for OkeClusterImageReplicationEnum
const (
	OkeClusterImageReplicationEnable  OkeClusterImageReplicationEnum = "ENABLE"
	OkeClusterImageReplicationDisable OkeClusterImageReplicationEnum = "DISABLE"
)

var mappingOkeClusterImageReplicationEnum = map[string]OkeClusterImageReplicationEnum{
	"ENABLE":  OkeClusterImageReplicationEnable,
	"DISABLE": OkeClusterImageReplicationDisable,
}

var mappingOkeClusterImageReplicationEnumLowerCase = map[string]OkeClusterImageReplicationEnum{
	"enable":  OkeClusterImageReplicationEnable,
	"disable": OkeClusterImageReplicationDisable,
}

// GetOkeClusterImageReplicationEnumValues Enumerates the set of values for OkeClusterImageReplicationEnum
func GetOkeClusterImageReplicationEnumValues() []OkeClusterImageReplicationEnum {
	values := make([]OkeClusterImageReplicationEnum, 0)
	for _, v := range mappingOkeClusterImageReplicationEnum {
		values = append(values, v)
	}
	return values
}

// GetOkeClusterImageReplicationEnumStringValues Enumerates the set of values in String for OkeClusterImageReplicationEnum
func GetOkeClusterImageReplicationEnumStringValues() []string {
	return []string{
		"ENABLE",
		"DISABLE",
	}
}

// GetMappingOkeClusterImageReplicationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOkeClusterImageReplicationEnum(val string) (OkeClusterImageReplicationEnum, bool) {
	enum, ok := mappingOkeClusterImageReplicationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
