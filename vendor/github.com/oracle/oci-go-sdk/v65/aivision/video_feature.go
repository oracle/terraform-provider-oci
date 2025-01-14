// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VideoFeature Details about a video feature request.
type VideoFeature interface {
}

type videofeature struct {
	JsonData    []byte
	FeatureType string `json:"featureType"`
}

// UnmarshalJSON unmarshals json
func (m *videofeature) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalervideofeature videofeature
	s := struct {
		Model Unmarshalervideofeature
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FeatureType = s.Model.FeatureType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *videofeature) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.FeatureType {
	case "OBJECT_DETECTION":
		mm := VideoObjectDetectionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FACE_DETECTION":
		mm := VideoFaceDetectionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT_DETECTION":
		mm := VideoTextDetectionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_TRACKING":
		mm := VideoObjectTrackingFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LABEL_DETECTION":
		mm := VideoLabelDetectionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for VideoFeature: %s.", m.FeatureType)
		return *m, nil
	}
}

func (m videofeature) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m videofeature) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VideoFeatureFeatureTypeEnum Enum with underlying type: string
type VideoFeatureFeatureTypeEnum string

// Set of constants representing the allowable values for VideoFeatureFeatureTypeEnum
const (
	VideoFeatureFeatureTypeLabelDetection  VideoFeatureFeatureTypeEnum = "LABEL_DETECTION"
	VideoFeatureFeatureTypeObjectDetection VideoFeatureFeatureTypeEnum = "OBJECT_DETECTION"
	VideoFeatureFeatureTypeTextDetection   VideoFeatureFeatureTypeEnum = "TEXT_DETECTION"
	VideoFeatureFeatureTypeFaceDetection   VideoFeatureFeatureTypeEnum = "FACE_DETECTION"
	VideoFeatureFeatureTypeObjectTracking  VideoFeatureFeatureTypeEnum = "OBJECT_TRACKING"
)

var mappingVideoFeatureFeatureTypeEnum = map[string]VideoFeatureFeatureTypeEnum{
	"LABEL_DETECTION":  VideoFeatureFeatureTypeLabelDetection,
	"OBJECT_DETECTION": VideoFeatureFeatureTypeObjectDetection,
	"TEXT_DETECTION":   VideoFeatureFeatureTypeTextDetection,
	"FACE_DETECTION":   VideoFeatureFeatureTypeFaceDetection,
	"OBJECT_TRACKING":  VideoFeatureFeatureTypeObjectTracking,
}

var mappingVideoFeatureFeatureTypeEnumLowerCase = map[string]VideoFeatureFeatureTypeEnum{
	"label_detection":  VideoFeatureFeatureTypeLabelDetection,
	"object_detection": VideoFeatureFeatureTypeObjectDetection,
	"text_detection":   VideoFeatureFeatureTypeTextDetection,
	"face_detection":   VideoFeatureFeatureTypeFaceDetection,
	"object_tracking":  VideoFeatureFeatureTypeObjectTracking,
}

// GetVideoFeatureFeatureTypeEnumValues Enumerates the set of values for VideoFeatureFeatureTypeEnum
func GetVideoFeatureFeatureTypeEnumValues() []VideoFeatureFeatureTypeEnum {
	values := make([]VideoFeatureFeatureTypeEnum, 0)
	for _, v := range mappingVideoFeatureFeatureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVideoFeatureFeatureTypeEnumStringValues Enumerates the set of values in String for VideoFeatureFeatureTypeEnum
func GetVideoFeatureFeatureTypeEnumStringValues() []string {
	return []string{
		"LABEL_DETECTION",
		"OBJECT_DETECTION",
		"TEXT_DETECTION",
		"FACE_DETECTION",
		"OBJECT_TRACKING",
	}
}

// GetMappingVideoFeatureFeatureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVideoFeatureFeatureTypeEnum(val string) (VideoFeatureFeatureTypeEnum, bool) {
	enum, ok := mappingVideoFeatureFeatureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
