// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v46/common"
)

// ParquetFormatAttribute The PARQUET format attribute.
type ParquetFormatAttribute struct {

	// Defines whether a file pattern is supported.
	IsFilePattern *bool `mandatory:"false" json:"isFilePattern"`

	// The compression for the file.
	Compression *string `mandatory:"false" json:"compression"`
}

//GetIsFilePattern returns IsFilePattern
func (m ParquetFormatAttribute) GetIsFilePattern() *bool {
	return m.IsFilePattern
}

func (m ParquetFormatAttribute) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ParquetFormatAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeParquetFormatAttribute ParquetFormatAttribute
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeParquetFormatAttribute
	}{
		"PARQUET_FORMAT",
		(MarshalTypeParquetFormatAttribute)(m),
	}

	return json.Marshal(&s)
}
