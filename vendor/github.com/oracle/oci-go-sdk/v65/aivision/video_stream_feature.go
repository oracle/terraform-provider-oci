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

// VideoStreamFeature Details about a stream video feature request.
type VideoStreamFeature interface {
}

type videostreamfeature struct {
	JsonData    []byte
	FeatureType string `json:"featureType"`
}

// UnmarshalJSON unmarshals json
func (m *videostreamfeature) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalervideostreamfeature videostreamfeature
	s := struct {
		Model Unmarshalervideostreamfeature
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FeatureType = s.Model.FeatureType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *videostreamfeature) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.FeatureType {
	case "OBJECT_TRACKING":
		mm := VideoStreamObjectTrackingFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FACE_DETECTION":
		mm := VideoStreamFaceDetectionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_DETECTION":
		mm := VideoStreamObjectDetectionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for VideoStreamFeature: %s.", m.FeatureType)
		return *m, nil
	}
}

func (m videostreamfeature) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m videostreamfeature) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VideoStreamFeatureFeatureTypeEnum Enum with underlying type: string
type VideoStreamFeatureFeatureTypeEnum string

// Set of constants representing the allowable values for VideoStreamFeatureFeatureTypeEnum
const (
	VideoStreamFeatureFeatureTypeObjectTracking  VideoStreamFeatureFeatureTypeEnum = "OBJECT_TRACKING"
	VideoStreamFeatureFeatureTypeFaceDetection   VideoStreamFeatureFeatureTypeEnum = "FACE_DETECTION"
	VideoStreamFeatureFeatureTypeObjectDetection VideoStreamFeatureFeatureTypeEnum = "OBJECT_DETECTION"
)

var mappingVideoStreamFeatureFeatureTypeEnum = map[string]VideoStreamFeatureFeatureTypeEnum{
	"OBJECT_TRACKING":  VideoStreamFeatureFeatureTypeObjectTracking,
	"FACE_DETECTION":   VideoStreamFeatureFeatureTypeFaceDetection,
	"OBJECT_DETECTION": VideoStreamFeatureFeatureTypeObjectDetection,
}

var mappingVideoStreamFeatureFeatureTypeEnumLowerCase = map[string]VideoStreamFeatureFeatureTypeEnum{
	"object_tracking":  VideoStreamFeatureFeatureTypeObjectTracking,
	"face_detection":   VideoStreamFeatureFeatureTypeFaceDetection,
	"object_detection": VideoStreamFeatureFeatureTypeObjectDetection,
}

// GetVideoStreamFeatureFeatureTypeEnumValues Enumerates the set of values for VideoStreamFeatureFeatureTypeEnum
func GetVideoStreamFeatureFeatureTypeEnumValues() []VideoStreamFeatureFeatureTypeEnum {
	values := make([]VideoStreamFeatureFeatureTypeEnum, 0)
	for _, v := range mappingVideoStreamFeatureFeatureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVideoStreamFeatureFeatureTypeEnumStringValues Enumerates the set of values in String for VideoStreamFeatureFeatureTypeEnum
func GetVideoStreamFeatureFeatureTypeEnumStringValues() []string {
	return []string{
		"OBJECT_TRACKING",
		"FACE_DETECTION",
		"OBJECT_DETECTION",
	}
}

// GetMappingVideoStreamFeatureFeatureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVideoStreamFeatureFeatureTypeEnum(val string) (VideoStreamFeatureFeatureTypeEnum, bool) {
	enum, ok := mappingVideoStreamFeatureFeatureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
