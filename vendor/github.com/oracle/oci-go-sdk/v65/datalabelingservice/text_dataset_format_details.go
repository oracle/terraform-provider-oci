// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TextDatasetFormatDetails It indicates the dataset is comprised of TXT files.
type TextDatasetFormatDetails struct {
	TextFileTypeMetadata TextFileTypeMetadata `mandatory:"false" json:"textFileTypeMetadata"`
}

func (m TextDatasetFormatDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TextDatasetFormatDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TextDatasetFormatDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTextDatasetFormatDetails TextDatasetFormatDetails
	s := struct {
		DiscriminatorParam string `json:"formatType"`
		MarshalTypeTextDatasetFormatDetails
	}{
		"TEXT",
		(MarshalTypeTextDatasetFormatDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TextDatasetFormatDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TextFileTypeMetadata textfiletypemetadata `json:"textFileTypeMetadata"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.TextFileTypeMetadata.UnmarshalPolymorphicJSON(model.TextFileTypeMetadata.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TextFileTypeMetadata = nn.(TextFileTypeMetadata)
	} else {
		m.TextFileTypeMetadata = nil
	}

	return
}
