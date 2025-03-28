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

// CreateImageJobDetails The details of the batch image analysis.
type CreateImageJobDetails struct {
	InputLocation InputLocation `mandatory:"true" json:"inputLocation"`

	// The list of requested image analysis types.
	Features []ImageFeature `mandatory:"true" json:"features"`

	OutputLocation *OutputLocation `mandatory:"true" json:"outputLocation"`

	// The compartment identifier from the requester.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The image job display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether or not to generate a ZIP file containing the results.
	IsZipOutputEnabled *bool `mandatory:"false" json:"isZipOutputEnabled"`
}

func (m CreateImageJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateImageJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateImageJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId      *string         `json:"compartmentId"`
		DisplayName        *string         `json:"displayName"`
		IsZipOutputEnabled *bool           `json:"isZipOutputEnabled"`
		InputLocation      inputlocation   `json:"inputLocation"`
		Features           []imagefeature  `json:"features"`
		OutputLocation     *OutputLocation `json:"outputLocation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.IsZipOutputEnabled = model.IsZipOutputEnabled

	nn, e = model.InputLocation.UnmarshalPolymorphicJSON(model.InputLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InputLocation = nn.(InputLocation)
	} else {
		m.InputLocation = nil
	}

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
	m.OutputLocation = model.OutputLocation

	return
}
