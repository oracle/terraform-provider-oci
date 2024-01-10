// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ImageFeature The type of image analysis.
type ImageFeature interface {
}

type imagefeature struct {
	JsonData    []byte
	FeatureType string `json:"featureType"`
}

// UnmarshalJSON unmarshals json
func (m *imagefeature) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerimagefeature imagefeature
	s := struct {
		Model Unmarshalerimagefeature
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FeatureType = s.Model.FeatureType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *imagefeature) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.FeatureType {
	case "TEXT_DETECTION":
		mm := ImageTextDetectionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FACE_DETECTION":
		mm := FaceDetectionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_DETECTION":
		mm := ImageObjectDetectionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IMAGE_CLASSIFICATION":
		mm := ImageClassificationFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ImageFeature: %s.", m.FeatureType)
		return *m, nil
	}
}

func (m imagefeature) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m imagefeature) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImageFeatureFeatureTypeEnum Enum with underlying type: string
type ImageFeatureFeatureTypeEnum string

// Set of constants representing the allowable values for ImageFeatureFeatureTypeEnum
const (
	ImageFeatureFeatureTypeImageClassification ImageFeatureFeatureTypeEnum = "IMAGE_CLASSIFICATION"
	ImageFeatureFeatureTypeObjectDetection     ImageFeatureFeatureTypeEnum = "OBJECT_DETECTION"
	ImageFeatureFeatureTypeTextDetection       ImageFeatureFeatureTypeEnum = "TEXT_DETECTION"
	ImageFeatureFeatureTypeFaceDetection       ImageFeatureFeatureTypeEnum = "FACE_DETECTION"
)

var mappingImageFeatureFeatureTypeEnum = map[string]ImageFeatureFeatureTypeEnum{
	"IMAGE_CLASSIFICATION": ImageFeatureFeatureTypeImageClassification,
	"OBJECT_DETECTION":     ImageFeatureFeatureTypeObjectDetection,
	"TEXT_DETECTION":       ImageFeatureFeatureTypeTextDetection,
	"FACE_DETECTION":       ImageFeatureFeatureTypeFaceDetection,
}

var mappingImageFeatureFeatureTypeEnumLowerCase = map[string]ImageFeatureFeatureTypeEnum{
	"image_classification": ImageFeatureFeatureTypeImageClassification,
	"object_detection":     ImageFeatureFeatureTypeObjectDetection,
	"text_detection":       ImageFeatureFeatureTypeTextDetection,
	"face_detection":       ImageFeatureFeatureTypeFaceDetection,
}

// GetImageFeatureFeatureTypeEnumValues Enumerates the set of values for ImageFeatureFeatureTypeEnum
func GetImageFeatureFeatureTypeEnumValues() []ImageFeatureFeatureTypeEnum {
	values := make([]ImageFeatureFeatureTypeEnum, 0)
	for _, v := range mappingImageFeatureFeatureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetImageFeatureFeatureTypeEnumStringValues Enumerates the set of values in String for ImageFeatureFeatureTypeEnum
func GetImageFeatureFeatureTypeEnumStringValues() []string {
	return []string{
		"IMAGE_CLASSIFICATION",
		"OBJECT_DETECTION",
		"TEXT_DETECTION",
		"FACE_DETECTION",
	}
}

// GetMappingImageFeatureFeatureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImageFeatureFeatureTypeEnum(val string) (ImageFeatureFeatureTypeEnum, bool) {
	enum, ok := mappingImageFeatureFeatureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
