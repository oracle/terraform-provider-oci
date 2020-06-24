// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
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

	// Format for timestamp data.
	TimestampFormat *string `mandatory:"false" json:"timestampFormat"`
}

func (m CsvFormatAttribute) String() string {
	return common.PointerString(m)
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
