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

// ImportPreAnnotatedDataDetails Allows user to import dataset labels, records and annotations from dataset files
type ImportPreAnnotatedDataDetails struct {
	ImportFormat *ImportFormat `mandatory:"false" json:"importFormat"`

	ImportMetadataPath ImportMetadataPath `mandatory:"false" json:"importMetadataPath"`
}

func (m ImportPreAnnotatedDataDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportPreAnnotatedDataDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ImportPreAnnotatedDataDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ImportFormat       *ImportFormat      `json:"importFormat"`
		ImportMetadataPath importmetadatapath `json:"importMetadataPath"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ImportFormat = model.ImportFormat

	nn, e = model.ImportMetadataPath.UnmarshalPolymorphicJSON(model.ImportMetadataPath.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ImportMetadataPath = nn.(ImportMetadataPath)
	} else {
		m.ImportMetadataPath = nil
	}

	return
}
