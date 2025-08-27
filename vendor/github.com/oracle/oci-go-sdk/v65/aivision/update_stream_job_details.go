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

// UpdateStreamJobDetails The information needed to update streamjob
type UpdateStreamJobDetails struct {

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of streamSource
	StreamSourceId *string `mandatory:"false" json:"streamSourceId"`

	// List of stream analysis features.
	Features []VideoStreamFeature `mandatory:"false" json:"features"`

	StreamOutputLocation StreamOutputLocation `mandatory:"false" json:"streamOutputLocation"`

	// Stream job display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only.
	// For example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateStreamJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateStreamJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateStreamJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		StreamSourceId       *string                           `json:"streamSourceId"`
		Features             []videostreamfeature              `json:"features"`
		StreamOutputLocation streamoutputlocation              `json:"streamOutputLocation"`
		DisplayName          *string                           `json:"displayName"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.StreamSourceId = model.StreamSourceId

	m.Features = make([]VideoStreamFeature, len(model.Features))
	for i, n := range model.Features {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Features[i] = nn.(VideoStreamFeature)
		} else {
			m.Features[i] = nil
		}
	}
	nn, e = model.StreamOutputLocation.UnmarshalPolymorphicJSON(model.StreamOutputLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StreamOutputLocation = nn.(StreamOutputLocation)
	} else {
		m.StreamOutputLocation = nil
	}

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
