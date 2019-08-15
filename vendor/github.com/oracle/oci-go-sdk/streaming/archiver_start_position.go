// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Streaming Service API
//
// The API for the Streaming Service.
//

package streaming

// ArchiverStartPositionEnum Enum with underlying type: string
type ArchiverStartPositionEnum string

// Set of constants representing the allowable values for ArchiverStartPositionEnum
const (
	ArchiverStartPositionLatest      ArchiverStartPositionEnum = "LATEST"
	ArchiverStartPositionTrimHorizon ArchiverStartPositionEnum = "TRIM_HORIZON"
)

var mappingArchiverStartPosition = map[string]ArchiverStartPositionEnum{
	"LATEST":       ArchiverStartPositionLatest,
	"TRIM_HORIZON": ArchiverStartPositionTrimHorizon,
}

// GetArchiverStartPositionEnumValues Enumerates the set of values for ArchiverStartPositionEnum
func GetArchiverStartPositionEnumValues() []ArchiverStartPositionEnum {
	values := make([]ArchiverStartPositionEnum, 0)
	for _, v := range mappingArchiverStartPosition {
		values = append(values, v)
	}
	return values
}
