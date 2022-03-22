// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v63/common"
	"strings"
)

// CsvFormatAttribute The CSV format attribute.
type CsvFormatAttribute struct {

	// The encoding for the file.
	Encoding *string `mandatory:"false" json:"encoding"`

	// The escape character for the CSV format.
	EscapeCharacter *string `mandatory:"false" json:"escapeCharacter"`

	// The delimiter for the CSV format.
	Delimiter *string `mandatory:"false" json:"delimiter"`

	// The quote character for the CSV format.
	QuoteCharacter *string `mandatory:"false" json:"quoteCharacter"`

	// Defines whether the file has a header row.
	HasHeader *bool `mandatory:"false" json:"hasHeader"`

	// Defines whether a file pattern is supported.
	IsFilePattern *bool `mandatory:"false" json:"isFilePattern"`

	// Format for timestamp information.
	TimestampFormat *string `mandatory:"false" json:"timestampFormat"`
}

func (m CsvFormatAttribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CsvFormatAttribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CsvFormatAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCsvFormatAttribute CsvFormatAttribute
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCsvFormatAttribute
	}{
		"CSV_FORMAT",
		(MarshalTypeCsvFormatAttribute)(m),
	}

	return json.Marshal(&s)
}
