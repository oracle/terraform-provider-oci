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

// AnalyzeImageDetails The details of how to analyze an image.
type AnalyzeImageDetails struct {

	// The types of image analysis.
	Features []ImageFeature `mandatory:"true" json:"features"`

	Image ImageDetails `mandatory:"true" json:"image"`

	// The OCID of the compartment that calls the API.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m AnalyzeImageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnalyzeImageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AnalyzeImageDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId *string        `json:"compartmentId"`
		Features      []imagefeature `json:"features"`
		Image         imagedetails   `json:"image"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.Features = make([]ImageFeature, len(model.Features))
	for i, n := range model.Features {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Features[i] = nn.(ImageFeature)
		} else {
			m.Features[i] = nil
		}
	}
	nn, e = model.Image.UnmarshalPolymorphicJSON(model.Image.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Image = nn.(ImageDetails)
	} else {
		m.Image = nil
	}

	return
}
