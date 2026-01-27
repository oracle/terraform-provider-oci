// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CompressionOptions The additional compression options used while exporting the DB system backup.
type CompressionOptions struct {

	// The compression status of the exported data.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The level of compression to use when creating the dump files.
	Level *int `mandatory:"false" json:"level"`

	// The compression type to use when creating the dump files.
	Type CompressionOptionsTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m CompressionOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CompressionOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCompressionOptionsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCompressionOptionsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CompressionOptionsTypeEnum Enum with underlying type: string
type CompressionOptionsTypeEnum string

// Set of constants representing the allowable values for CompressionOptionsTypeEnum
const (
	CompressionOptionsTypeGzip CompressionOptionsTypeEnum = "GZIP"
	CompressionOptionsTypeZstd CompressionOptionsTypeEnum = "ZSTD"
)

var mappingCompressionOptionsTypeEnum = map[string]CompressionOptionsTypeEnum{
	"GZIP": CompressionOptionsTypeGzip,
	"ZSTD": CompressionOptionsTypeZstd,
}

var mappingCompressionOptionsTypeEnumLowerCase = map[string]CompressionOptionsTypeEnum{
	"gzip": CompressionOptionsTypeGzip,
	"zstd": CompressionOptionsTypeZstd,
}

// GetCompressionOptionsTypeEnumValues Enumerates the set of values for CompressionOptionsTypeEnum
func GetCompressionOptionsTypeEnumValues() []CompressionOptionsTypeEnum {
	values := make([]CompressionOptionsTypeEnum, 0)
	for _, v := range mappingCompressionOptionsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCompressionOptionsTypeEnumStringValues Enumerates the set of values in String for CompressionOptionsTypeEnum
func GetCompressionOptionsTypeEnumStringValues() []string {
	return []string{
		"GZIP",
		"ZSTD",
	}
}

// GetMappingCompressionOptionsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCompressionOptionsTypeEnum(val string) (CompressionOptionsTypeEnum, bool) {
	enum, ok := mappingCompressionOptionsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
