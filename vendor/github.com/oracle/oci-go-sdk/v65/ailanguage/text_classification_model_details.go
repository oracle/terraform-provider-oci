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

// TextClassificationModelDetails Possible TextClassificationModelDetails
type TextClassificationModelDetails struct {

	// supported language default value is en
	LanguageCode *string `mandatory:"false" json:"languageCode"`

	ClassificationMode ClassificationType `mandatory:"false" json:"classificationMode"`
}

// GetLanguageCode returns LanguageCode
func (m TextClassificationModelDetails) GetLanguageCode() *string {
	return m.LanguageCode
}

func (m TextClassificationModelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TextClassificationModelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TextClassificationModelDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTextClassificationModelDetails TextClassificationModelDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTextClassificationModelDetails
	}{
		"TEXT_CLASSIFICATION",
		(MarshalTypeTextClassificationModelDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TextClassificationModelDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LanguageCode       *string            `json:"languageCode"`
		ClassificationMode classificationtype `json:"classificationMode"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.LanguageCode = model.LanguageCode

	nn, e = model.ClassificationMode.UnmarshalPolymorphicJSON(model.ClassificationMode.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ClassificationMode = nn.(ClassificationType)
	} else {
		m.ClassificationMode = nil
	}

	return
}
