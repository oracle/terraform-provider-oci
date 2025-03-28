// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

import (
	"strings"
)

// WorkRequestResourceMetadataKeyEnum Enum with underlying type: string
type WorkRequestResourceMetadataKeyEnum string

// Set of constants representing the allowable values for WorkRequestResourceMetadataKeyEnum
const (
	WorkRequestResourceMetadataKeyNamespace       WorkRequestResourceMetadataKeyEnum = "NAMESPACE"
	WorkRequestResourceMetadataKeyBucket          WorkRequestResourceMetadataKeyEnum = "BUCKET"
	WorkRequestResourceMetadataKeyObject          WorkRequestResourceMetadataKeyEnum = "OBJECT"
	WorkRequestResourceMetadataKeySnapshotVersion WorkRequestResourceMetadataKeyEnum = "SNAPSHOT_VERSION"
)

var mappingWorkRequestResourceMetadataKeyEnum = map[string]WorkRequestResourceMetadataKeyEnum{
	"NAMESPACE":        WorkRequestResourceMetadataKeyNamespace,
	"BUCKET":           WorkRequestResourceMetadataKeyBucket,
	"OBJECT":           WorkRequestResourceMetadataKeyObject,
	"SNAPSHOT_VERSION": WorkRequestResourceMetadataKeySnapshotVersion,
}

var mappingWorkRequestResourceMetadataKeyEnumLowerCase = map[string]WorkRequestResourceMetadataKeyEnum{
	"namespace":        WorkRequestResourceMetadataKeyNamespace,
	"bucket":           WorkRequestResourceMetadataKeyBucket,
	"object":           WorkRequestResourceMetadataKeyObject,
	"snapshot_version": WorkRequestResourceMetadataKeySnapshotVersion,
}

// GetWorkRequestResourceMetadataKeyEnumValues Enumerates the set of values for WorkRequestResourceMetadataKeyEnum
func GetWorkRequestResourceMetadataKeyEnumValues() []WorkRequestResourceMetadataKeyEnum {
	values := make([]WorkRequestResourceMetadataKeyEnum, 0)
	for _, v := range mappingWorkRequestResourceMetadataKeyEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResourceMetadataKeyEnumStringValues Enumerates the set of values in String for WorkRequestResourceMetadataKeyEnum
func GetWorkRequestResourceMetadataKeyEnumStringValues() []string {
	return []string{
		"NAMESPACE",
		"BUCKET",
		"OBJECT",
		"SNAPSHOT_VERSION",
	}
}

// GetMappingWorkRequestResourceMetadataKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResourceMetadataKeyEnum(val string) (WorkRequestResourceMetadataKeyEnum, bool) {
	enum, ok := mappingWorkRequestResourceMetadataKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
