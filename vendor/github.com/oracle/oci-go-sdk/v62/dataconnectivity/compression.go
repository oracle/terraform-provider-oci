// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"strings"
)

// Compression The optional compression configuration.
type Compression struct {

	// Compression algorithm
	Codec CompressionCodecEnum `mandatory:"true" json:"codec"`
}

func (m Compression) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Compression) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCompressionCodecEnum(string(m.Codec)); !ok && m.Codec != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Codec: %s. Supported values are: %s.", m.Codec, strings.Join(GetCompressionCodecEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingCompressionCodecEnum = map[string]CompressionCodecEnum{
	"NONE":    CompressionCodecNone,
	"AUTO":    CompressionCodecAuto,
	"GZIP":    CompressionCodecGzip,
	"BZIP2":   CompressionCodecBzip2,
	"DEFLATE": CompressionCodecDeflate,
	"LZ4":     CompressionCodecLz4,
	"SNAPPY":  CompressionCodecSnappy,
}

var mappingCompressionCodecEnumLowerCase = map[string]CompressionCodecEnum{
	"none":    CompressionCodecNone,
	"auto":    CompressionCodecAuto,
	"gzip":    CompressionCodecGzip,
	"bzip2":   CompressionCodecBzip2,
	"deflate": CompressionCodecDeflate,
	"lz4":     CompressionCodecLz4,
	"snappy":  CompressionCodecSnappy,
}

// GetCompressionCodecEnumValues Enumerates the set of values for CompressionCodecEnum
func GetCompressionCodecEnumValues() []CompressionCodecEnum {
	values := make([]CompressionCodecEnum, 0)
	for _, v := range mappingCompressionCodecEnum {
		values = append(values, v)
	}
	return values
}

// GetCompressionCodecEnumStringValues Enumerates the set of values in String for CompressionCodecEnum
func GetCompressionCodecEnumStringValues() []string {
	return []string{
		"NONE",
		"AUTO",
		"GZIP",
		"BZIP2",
		"DEFLATE",
		"LZ4",
		"SNAPPY",
	}
}

// GetMappingCompressionCodecEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCompressionCodecEnum(val string) (CompressionCodecEnum, bool) {
	enum, ok := mappingCompressionCodecEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
