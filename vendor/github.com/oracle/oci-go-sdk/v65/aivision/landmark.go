// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Landmark The landmark on the face.
type Landmark struct {

	// The face landmark type
	Type LandmarkTypeEnum `mandatory:"true" json:"type"`

	// The X-axis normalized coordinate.
	X *float32 `mandatory:"true" json:"x"`

	// The Y-axis normalized coordinate.
	Y *float32 `mandatory:"true" json:"y"`
}

func (m Landmark) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Landmark) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLandmarkTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetLandmarkTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LandmarkTypeEnum Enum with underlying type: string
type LandmarkTypeEnum string

// Set of constants representing the allowable values for LandmarkTypeEnum
const (
	LandmarkTypeLeftEye          LandmarkTypeEnum = "LEFT_EYE"
	LandmarkTypeRightEye         LandmarkTypeEnum = "RIGHT_EYE"
	LandmarkTypeNoseTip          LandmarkTypeEnum = "NOSE_TIP"
	LandmarkTypeLeftEdgeOfMouth  LandmarkTypeEnum = "LEFT_EDGE_OF_MOUTH"
	LandmarkTypeRightEdgeOfMouth LandmarkTypeEnum = "RIGHT_EDGE_OF_MOUTH"
)

var mappingLandmarkTypeEnum = map[string]LandmarkTypeEnum{
	"LEFT_EYE":            LandmarkTypeLeftEye,
	"RIGHT_EYE":           LandmarkTypeRightEye,
	"NOSE_TIP":            LandmarkTypeNoseTip,
	"LEFT_EDGE_OF_MOUTH":  LandmarkTypeLeftEdgeOfMouth,
	"RIGHT_EDGE_OF_MOUTH": LandmarkTypeRightEdgeOfMouth,
}

var mappingLandmarkTypeEnumLowerCase = map[string]LandmarkTypeEnum{
	"left_eye":            LandmarkTypeLeftEye,
	"right_eye":           LandmarkTypeRightEye,
	"nose_tip":            LandmarkTypeNoseTip,
	"left_edge_of_mouth":  LandmarkTypeLeftEdgeOfMouth,
	"right_edge_of_mouth": LandmarkTypeRightEdgeOfMouth,
}

// GetLandmarkTypeEnumValues Enumerates the set of values for LandmarkTypeEnum
func GetLandmarkTypeEnumValues() []LandmarkTypeEnum {
	values := make([]LandmarkTypeEnum, 0)
	for _, v := range mappingLandmarkTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLandmarkTypeEnumStringValues Enumerates the set of values in String for LandmarkTypeEnum
func GetLandmarkTypeEnumStringValues() []string {
	return []string{
		"LEFT_EYE",
		"RIGHT_EYE",
		"NOSE_TIP",
		"LEFT_EDGE_OF_MOUTH",
		"RIGHT_EDGE_OF_MOUTH",
	}
}

// GetMappingLandmarkTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLandmarkTypeEnum(val string) (LandmarkTypeEnum, bool) {
	enum, ok := mappingLandmarkTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
