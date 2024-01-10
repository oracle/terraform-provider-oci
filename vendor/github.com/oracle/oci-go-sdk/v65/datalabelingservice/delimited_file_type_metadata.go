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

// DelimitedFileTypeMetadata Metadata of delimited files.
type DelimitedFileTypeMetadata struct {

	// The index of a selected column. This is a zero-based index.
	ColumnIndex *int `mandatory:"true" json:"columnIndex"`

	// The name of a selected column.
	ColumnName *string `mandatory:"false" json:"columnName"`

	// A column delimiter
	ColumnDelimiter *string `mandatory:"false" json:"columnDelimiter"`

	// A line delimiter.
	LineDelimiter *string `mandatory:"false" json:"lineDelimiter"`

	// An escape character.
	EscapeCharacter *string `mandatory:"false" json:"escapeCharacter"`
}

func (m DelimitedFileTypeMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DelimitedFileTypeMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DelimitedFileTypeMetadata) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDelimitedFileTypeMetadata DelimitedFileTypeMetadata
	s := struct {
		DiscriminatorParam string `json:"formatType"`
		MarshalTypeDelimitedFileTypeMetadata
	}{
		"DELIMITED",
		(MarshalTypeDelimitedFileTypeMetadata)(m),
	}

	return json.Marshal(&s)
}
