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

// BatchDetectLanguagePiiEntitiesDetails The documents details to detect personal identification information.
type BatchDetectLanguagePiiEntitiesDetails struct {

	// List of documents to detect personal identification information.
	Documents []TextDocument `mandatory:"true" json:"documents"`

	// The endpoint which have to be used for inferencing. If endpointId and compartmentId is provided, then inference will be served from custom model which is mapped to this Endpoint.
	EndpointId *string `mandatory:"false" json:"endpointId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that calls the API, inference will be served from pre trained model
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Mask recognized PII entities with different modes.
	Masking map[string]PiiEntityMasking `mandatory:"false" json:"masking"`

	Profile *Profile `mandatory:"false" json:"profile"`
}

func (m BatchDetectLanguagePiiEntitiesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BatchDetectLanguagePiiEntitiesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *BatchDetectLanguagePiiEntitiesDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		EndpointId    *string                     `json:"endpointId"`
		CompartmentId *string                     `json:"compartmentId"`
		Masking       map[string]piientitymasking `json:"masking"`
		Profile       *Profile                    `json:"profile"`
		Documents     []TextDocument              `json:"documents"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.EndpointId = model.EndpointId

	m.CompartmentId = model.CompartmentId

	m.Masking = make(map[string]PiiEntityMasking)
	for k, v := range model.Masking {
		nn, e = v.UnmarshalPolymorphicJSON(v.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Masking[k] = nn.(PiiEntityMasking)
		} else {
			m.Masking[k] = nil
		}
	}

	m.Profile = model.Profile

	m.Documents = make([]TextDocument, len(model.Documents))
	copy(m.Documents, model.Documents)
	return
}
