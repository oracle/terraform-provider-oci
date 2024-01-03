// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateProcessorJobDetails The details used to create a processor job.
type CreateProcessorJobDetails struct {
	InputLocation InputLocation `mandatory:"true" json:"inputLocation"`

	OutputLocation *OutputLocation `mandatory:"true" json:"outputLocation"`

	// The compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ProcessorConfig ProcessorConfig `mandatory:"true" json:"processorConfig"`

	// The display name of the processor job.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CreateProcessorJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateProcessorJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateProcessorJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName     *string         `json:"displayName"`
		InputLocation   inputlocation   `json:"inputLocation"`
		OutputLocation  *OutputLocation `json:"outputLocation"`
		CompartmentId   *string         `json:"compartmentId"`
		ProcessorConfig processorconfig `json:"processorConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	nn, e = model.InputLocation.UnmarshalPolymorphicJSON(model.InputLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InputLocation = nn.(InputLocation)
	} else {
		m.InputLocation = nil
	}

	m.OutputLocation = model.OutputLocation

	m.CompartmentId = model.CompartmentId

	nn, e = model.ProcessorConfig.UnmarshalPolymorphicJSON(model.ProcessorConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ProcessorConfig = nn.(ProcessorConfig)
	} else {
		m.ProcessorConfig = nil
	}

	return
}
