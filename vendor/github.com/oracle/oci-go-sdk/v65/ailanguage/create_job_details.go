// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateJobDetails Job creation detail which will have documents on which language services need to run prediction along with output folder
type CreateJobDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	InputLocation InputLocation `mandatory:"true" json:"inputLocation"`

	// training model details
	// For this release only one model is allowed to be input here.
	// One of the three modelType, ModelId, endpointId should be given other wise error will be thrown from API
	ModelMetadataDetails []ModelMetadataDetails `mandatory:"true" json:"modelMetadataDetails"`

	OutputLocation *ObjectPrefixOutputLocation `mandatory:"true" json:"outputLocation"`

	// A user-friendly display name for the job.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the job.
	Description *string `mandatory:"false" json:"description"`

	InputConfiguration *InputConfiguration `mandatory:"false" json:"inputConfiguration"`
}

func (m CreateJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                     `json:"displayName"`
		Description          *string                     `json:"description"`
		InputConfiguration   *InputConfiguration         `json:"inputConfiguration"`
		CompartmentId        *string                     `json:"compartmentId"`
		InputLocation        inputlocation               `json:"inputLocation"`
		ModelMetadataDetails []ModelMetadataDetails      `json:"modelMetadataDetails"`
		OutputLocation       *ObjectPrefixOutputLocation `json:"outputLocation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.InputConfiguration = model.InputConfiguration

	m.CompartmentId = model.CompartmentId

	nn, e = model.InputLocation.UnmarshalPolymorphicJSON(model.InputLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InputLocation = nn.(InputLocation)
	} else {
		m.InputLocation = nil
	}

	m.ModelMetadataDetails = make([]ModelMetadataDetails, len(model.ModelMetadataDetails))
	copy(m.ModelMetadataDetails, model.ModelMetadataDetails)
	m.OutputLocation = model.OutputLocation

	return
}
