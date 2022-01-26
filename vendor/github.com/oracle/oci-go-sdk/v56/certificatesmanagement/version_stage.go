// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

// VersionStageEnum Enum with underlying type: string
type VersionStageEnum string

// Set of constants representing the allowable values for VersionStageEnum
const (
	VersionStageCurrent    VersionStageEnum = "CURRENT"
	VersionStagePending    VersionStageEnum = "PENDING"
	VersionStageLatest     VersionStageEnum = "LATEST"
	VersionStagePrevious   VersionStageEnum = "PREVIOUS"
	VersionStageDeprecated VersionStageEnum = "DEPRECATED"
	VersionStageFailed     VersionStageEnum = "FAILED"
)

var mappingVersionStage = map[string]VersionStageEnum{
	"CURRENT":    VersionStageCurrent,
	"PENDING":    VersionStagePending,
	"LATEST":     VersionStageLatest,
	"PREVIOUS":   VersionStagePrevious,
	"DEPRECATED": VersionStageDeprecated,
	"FAILED":     VersionStageFailed,
}

// GetVersionStageEnumValues Enumerates the set of values for VersionStageEnum
func GetVersionStageEnumValues() []VersionStageEnum {
	values := make([]VersionStageEnum, 0)
	for _, v := range mappingVersionStage {
		values = append(values, v)
	}
	return values
}
