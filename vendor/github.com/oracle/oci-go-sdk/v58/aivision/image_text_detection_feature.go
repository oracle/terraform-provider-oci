// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// VisionService API
//
// A description of the VisionService API.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ImageTextDetectionFeature Text detection parameters.
type ImageTextDetectionFeature struct {

	// Language of the document image, abbreviated according to ISO 639-2.
	Language DocumentLanguageEnum `mandatory:"false" json:"language,omitempty"`
}

func (m ImageTextDetectionFeature) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImageTextDetectionFeature) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDocumentLanguageEnum(string(m.Language)); !ok && m.Language != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Language: %s. Supported values are: %s.", m.Language, strings.Join(GetDocumentLanguageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ImageTextDetectionFeature) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImageTextDetectionFeature ImageTextDetectionFeature
	s := struct {
		DiscriminatorParam string `json:"featureType"`
		MarshalTypeImageTextDetectionFeature
	}{
		"TEXT_DETECTION",
		(MarshalTypeImageTextDetectionFeature)(m),
	}

	return json.Marshal(&s)
}
