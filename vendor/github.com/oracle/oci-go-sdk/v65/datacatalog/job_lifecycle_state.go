// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"strings"
)

// JobLifecycleStateEnum Enum with underlying type: string
type JobLifecycleStateEnum string

// Set of constants representing the allowable values for JobLifecycleStateEnum
const (
	JobLifecycleStateActive   JobLifecycleStateEnum = "ACTIVE"
	JobLifecycleStateInactive JobLifecycleStateEnum = "INACTIVE"
	JobLifecycleStateExpired  JobLifecycleStateEnum = "EXPIRED"
)

var mappingJobLifecycleStateEnum = map[string]JobLifecycleStateEnum{
	"ACTIVE":   JobLifecycleStateActive,
	"INACTIVE": JobLifecycleStateInactive,
	"EXPIRED":  JobLifecycleStateExpired,
}

var mappingJobLifecycleStateEnumLowerCase = map[string]JobLifecycleStateEnum{
	"active":   JobLifecycleStateActive,
	"inactive": JobLifecycleStateInactive,
	"expired":  JobLifecycleStateExpired,
}

// GetJobLifecycleStateEnumValues Enumerates the set of values for JobLifecycleStateEnum
func GetJobLifecycleStateEnumValues() []JobLifecycleStateEnum {
	values := make([]JobLifecycleStateEnum, 0)
	for _, v := range mappingJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetJobLifecycleStateEnumStringValues Enumerates the set of values in String for JobLifecycleStateEnum
func GetJobLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"EXPIRED",
	}
}

// GetMappingJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobLifecycleStateEnum(val string) (JobLifecycleStateEnum, bool) {
	enum, ok := mappingJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
