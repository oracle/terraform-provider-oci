// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GenerateContentDetails Represents the top-level request to generate new log analytics content, specifying the required details for the type of generation.
type GenerateContentDetails struct {
	GenAIDetails *GenAiDetails `mandatory:"true" json:"genAIDetails"`

	ContentRequest BaseContentRequest `mandatory:"true" json:"contentRequest"`
}

func (m GenerateContentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateContentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *GenerateContentDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		GenAIDetails   *GenAiDetails      `json:"genAIDetails"`
		ContentRequest basecontentrequest `json:"contentRequest"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.GenAIDetails = model.GenAIDetails

	nn, e = model.ContentRequest.UnmarshalPolymorphicJSON(model.ContentRequest.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ContentRequest = nn.(BaseContentRequest)
	} else {
		m.ContentRequest = nil
	}

	return
}
