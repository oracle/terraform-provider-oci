// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Compression The optional compression configuration.
type Compression struct {

	// Compression algorithm
	Codec CompressionCodecEnum `mandatory:"false" json:"codec,omitempty"`
}

func (m Compression) String() string {
	return common.PointerString(m)
}

// CompressionCodecEnum Enum with underlying type: string
type CompressionCodecEnum string

// Set of constants representing the allowable values for CompressionCodecEnum
const (
	CompressionCodecNone    CompressionCodecEnum = "NONE"
	CompressionCodecAuto    CompressionCodecEnum = "AUTO"
	CompressionCodecGzip    CompressionCodecEnum = "GZIP"
	CompressionCodecBzip2   CompressionCodecEnum = "BZIP2"
	CompressionCodecDeflate CompressionCodecEnum = "DEFLATE"
	CompressionCodecLz4     CompressionCodecEnum = "LZ4"
	CompressionCodecSnappy  CompressionCodecEnum = "SNAPPY"
)

var mappingCompressionCodec = map[string]CompressionCodecEnum{
	"NONE":    CompressionCodecNone,
	"AUTO":    CompressionCodecAuto,
	"GZIP":    CompressionCodecGzip,
	"BZIP2":   CompressionCodecBzip2,
	"DEFLATE": CompressionCodecDeflate,
	"LZ4":     CompressionCodecLz4,
	"SNAPPY":  CompressionCodecSnappy,
}

// GetCompressionCodecEnumValues Enumerates the set of values for CompressionCodecEnum
func GetCompressionCodecEnumValues() []CompressionCodecEnum {
	values := make([]CompressionCodecEnum, 0)
	for _, v := range mappingCompressionCodec {
		values = append(values, v)
	}
	return values
}
