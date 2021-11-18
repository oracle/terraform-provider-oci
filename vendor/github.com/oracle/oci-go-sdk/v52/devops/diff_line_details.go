// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// DiffLineDetails Details about a line within the diff.
type DiffLineDetails struct {

	// The number of a line in the base version.
	BaseLine *int `mandatory:"false" json:"baseLine"`

	// The number of a line in the target version.
	TargetLine *int `mandatory:"false" json:"targetLine"`

	// The contents of a line.
	LineContent *string `mandatory:"false" json:"lineContent"`

	// Indicates whether a line in a conflicted section of the diff is from the base version, the target version, or if its just a marker indicating the beginning, middle, or end of a conflicted section.
	ConflictMarker DiffLineDetailsConflictMarkerEnum `mandatory:"false" json:"conflictMarker,omitempty"`
}

func (m DiffLineDetails) String() string {
	return common.PointerString(m)
}

// DiffLineDetailsConflictMarkerEnum Enum with underlying type: string
type DiffLineDetailsConflictMarkerEnum string

// Set of constants representing the allowable values for DiffLineDetailsConflictMarkerEnum
const (
	DiffLineDetailsConflictMarkerBase   DiffLineDetailsConflictMarkerEnum = "BASE"
	DiffLineDetailsConflictMarkerTarget DiffLineDetailsConflictMarkerEnum = "TARGET"
	DiffLineDetailsConflictMarkerMarker DiffLineDetailsConflictMarkerEnum = "MARKER"
	DiffLineDetailsConflictMarkerNone   DiffLineDetailsConflictMarkerEnum = "NONE"
)

var mappingDiffLineDetailsConflictMarker = map[string]DiffLineDetailsConflictMarkerEnum{
	"BASE":   DiffLineDetailsConflictMarkerBase,
	"TARGET": DiffLineDetailsConflictMarkerTarget,
	"MARKER": DiffLineDetailsConflictMarkerMarker,
	"NONE":   DiffLineDetailsConflictMarkerNone,
}

// GetDiffLineDetailsConflictMarkerEnumValues Enumerates the set of values for DiffLineDetailsConflictMarkerEnum
func GetDiffLineDetailsConflictMarkerEnumValues() []DiffLineDetailsConflictMarkerEnum {
	values := make([]DiffLineDetailsConflictMarkerEnum, 0)
	for _, v := range mappingDiffLineDetailsConflictMarker {
		values = append(values, v)
	}
	return values
}
