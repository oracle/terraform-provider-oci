// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiffLineDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDiffLineDetailsConflictMarkerEnum(string(m.ConflictMarker)); !ok && m.ConflictMarker != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConflictMarker: %s. Supported values are: %s.", m.ConflictMarker, strings.Join(GetDiffLineDetailsConflictMarkerEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingDiffLineDetailsConflictMarkerEnum = map[string]DiffLineDetailsConflictMarkerEnum{
	"BASE":   DiffLineDetailsConflictMarkerBase,
	"TARGET": DiffLineDetailsConflictMarkerTarget,
	"MARKER": DiffLineDetailsConflictMarkerMarker,
	"NONE":   DiffLineDetailsConflictMarkerNone,
}

// GetDiffLineDetailsConflictMarkerEnumValues Enumerates the set of values for DiffLineDetailsConflictMarkerEnum
func GetDiffLineDetailsConflictMarkerEnumValues() []DiffLineDetailsConflictMarkerEnum {
	values := make([]DiffLineDetailsConflictMarkerEnum, 0)
	for _, v := range mappingDiffLineDetailsConflictMarkerEnum {
		values = append(values, v)
	}
	return values
}

// GetDiffLineDetailsConflictMarkerEnumStringValues Enumerates the set of values in String for DiffLineDetailsConflictMarkerEnum
func GetDiffLineDetailsConflictMarkerEnumStringValues() []string {
	return []string{
		"BASE",
		"TARGET",
		"MARKER",
		"NONE",
	}
}

// GetMappingDiffLineDetailsConflictMarkerEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiffLineDetailsConflictMarkerEnum(val string) (DiffLineDetailsConflictMarkerEnum, bool) {
	mappingDiffLineDetailsConflictMarkerEnumIgnoreCase := make(map[string]DiffLineDetailsConflictMarkerEnum)
	for k, v := range mappingDiffLineDetailsConflictMarkerEnum {
		mappingDiffLineDetailsConflictMarkerEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDiffLineDetailsConflictMarkerEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
