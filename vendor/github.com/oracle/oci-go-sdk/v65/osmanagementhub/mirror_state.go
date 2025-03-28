// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// MirrorStateEnum Enum with underlying type: string
type MirrorStateEnum string

// Set of constants representing the allowable values for MirrorStateEnum
const (
	MirrorStateUnsynced MirrorStateEnum = "UNSYNCED"
	MirrorStateQueued   MirrorStateEnum = "QUEUED"
	MirrorStateSyncing  MirrorStateEnum = "SYNCING"
	MirrorStateSynced   MirrorStateEnum = "SYNCED"
	MirrorStateFailed   MirrorStateEnum = "FAILED"
)

var mappingMirrorStateEnum = map[string]MirrorStateEnum{
	"UNSYNCED": MirrorStateUnsynced,
	"QUEUED":   MirrorStateQueued,
	"SYNCING":  MirrorStateSyncing,
	"SYNCED":   MirrorStateSynced,
	"FAILED":   MirrorStateFailed,
}

var mappingMirrorStateEnumLowerCase = map[string]MirrorStateEnum{
	"unsynced": MirrorStateUnsynced,
	"queued":   MirrorStateQueued,
	"syncing":  MirrorStateSyncing,
	"synced":   MirrorStateSynced,
	"failed":   MirrorStateFailed,
}

// GetMirrorStateEnumValues Enumerates the set of values for MirrorStateEnum
func GetMirrorStateEnumValues() []MirrorStateEnum {
	values := make([]MirrorStateEnum, 0)
	for _, v := range mappingMirrorStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMirrorStateEnumStringValues Enumerates the set of values in String for MirrorStateEnum
func GetMirrorStateEnumStringValues() []string {
	return []string{
		"UNSYNCED",
		"QUEUED",
		"SYNCING",
		"SYNCED",
		"FAILED",
	}
}

// GetMappingMirrorStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMirrorStateEnum(val string) (MirrorStateEnum, bool) {
	enum, ok := mappingMirrorStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
