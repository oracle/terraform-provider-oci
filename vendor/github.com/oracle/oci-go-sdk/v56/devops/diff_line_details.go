// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DiffLineDetails Details about a line within the difference.
type DiffLineDetails struct {

	// The number of a line in the base version.
	BaseLine *int `mandatory:"false" json:"baseLine"`

	// The number of a line in the target version.
	TargetLine *int `mandatory:"false" json:"targetLine"`

	// The contents of a line.
	LineContent *string `mandatory:"false" json:"lineContent"`

	// Indicates whether a line in a conflicted section of the difference is from the base version, the target version, or if its just a marker indicating the beginning, middle, or end of a conflicted section.
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
