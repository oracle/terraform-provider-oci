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

// ConfigWorkRequestGenAiContentResult Result object for GenAI content related config work requests.
type ConfigWorkRequestGenAiContentResult struct {
	GenAIDetails *GenAiDetails `mandatory:"true" json:"genAIDetails"`

	ContentRequest BaseContentRequest `mandatory:"true" json:"contentRequest"`

	ContentResponse BaseContentResponse `mandatory:"true" json:"contentResponse"`
}

func (m ConfigWorkRequestGenAiContentResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConfigWorkRequestGenAiContentResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ConfigWorkRequestGenAiContentResult) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConfigWorkRequestGenAiContentResult ConfigWorkRequestGenAiContentResult
	s := struct {
		DiscriminatorParam string `json:"workRequestKind"`
		MarshalTypeConfigWorkRequestGenAiContentResult
	}{
		"GENERATE_AI_CONTENT",
		(MarshalTypeConfigWorkRequestGenAiContentResult)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ConfigWorkRequestGenAiContentResult) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		GenAIDetails    *GenAiDetails       `json:"genAIDetails"`
		ContentRequest  basecontentrequest  `json:"contentRequest"`
		ContentResponse basecontentresponse `json:"contentResponse"`
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

	nn, e = model.ContentResponse.UnmarshalPolymorphicJSON(model.ContentResponse.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ContentResponse = nn.(BaseContentResponse)
	} else {
		m.ContentResponse = nil
	}

	return
}
