// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ImportFormat File format details used for importing dataset
type ImportFormat struct {

	// Name of import format
	Name ImportFormatNameEnum `mandatory:"true" json:"name"`

	// Version of import format
	Version ImportFormatVersionEnum `mandatory:"false" json:"version,omitempty"`
}

func (m ImportFormat) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportFormat) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingImportFormatNameEnum(string(m.Name)); !ok && m.Name != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Name: %s. Supported values are: %s.", m.Name, strings.Join(GetImportFormatNameEnumStringValues(), ",")))
	}

	if _, ok := GetMappingImportFormatVersionEnum(string(m.Version)); !ok && m.Version != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Version: %s. Supported values are: %s.", m.Version, strings.Join(GetImportFormatVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImportFormatNameEnum Enum with underlying type: string
type ImportFormatNameEnum string

// Set of constants representing the allowable values for ImportFormatNameEnum
const (
	ImportFormatNameJsonlConsolidated       ImportFormatNameEnum = "JSONL_CONSOLIDATED"
	ImportFormatNameJsonlCompactPlusContent ImportFormatNameEnum = "JSONL_COMPACT_PLUS_CONTENT"
	ImportFormatNameConll                   ImportFormatNameEnum = "CONLL"
	ImportFormatNameSpacy                   ImportFormatNameEnum = "SPACY"
	ImportFormatNameCoco                    ImportFormatNameEnum = "COCO"
	ImportFormatNameYolo                    ImportFormatNameEnum = "YOLO"
	ImportFormatNamePascalVoc               ImportFormatNameEnum = "PASCAL_VOC"
)

var mappingImportFormatNameEnum = map[string]ImportFormatNameEnum{
	"JSONL_CONSOLIDATED":         ImportFormatNameJsonlConsolidated,
	"JSONL_COMPACT_PLUS_CONTENT": ImportFormatNameJsonlCompactPlusContent,
	"CONLL":                      ImportFormatNameConll,
	"SPACY":                      ImportFormatNameSpacy,
	"COCO":                       ImportFormatNameCoco,
	"YOLO":                       ImportFormatNameYolo,
	"PASCAL_VOC":                 ImportFormatNamePascalVoc,
}

var mappingImportFormatNameEnumLowerCase = map[string]ImportFormatNameEnum{
	"jsonl_consolidated":         ImportFormatNameJsonlConsolidated,
	"jsonl_compact_plus_content": ImportFormatNameJsonlCompactPlusContent,
	"conll":                      ImportFormatNameConll,
	"spacy":                      ImportFormatNameSpacy,
	"coco":                       ImportFormatNameCoco,
	"yolo":                       ImportFormatNameYolo,
	"pascal_voc":                 ImportFormatNamePascalVoc,
}

// GetImportFormatNameEnumValues Enumerates the set of values for ImportFormatNameEnum
func GetImportFormatNameEnumValues() []ImportFormatNameEnum {
	values := make([]ImportFormatNameEnum, 0)
	for _, v := range mappingImportFormatNameEnum {
		values = append(values, v)
	}
	return values
}

// GetImportFormatNameEnumStringValues Enumerates the set of values in String for ImportFormatNameEnum
func GetImportFormatNameEnumStringValues() []string {
	return []string{
		"JSONL_CONSOLIDATED",
		"JSONL_COMPACT_PLUS_CONTENT",
		"CONLL",
		"SPACY",
		"COCO",
		"YOLO",
		"PASCAL_VOC",
	}
}

// GetMappingImportFormatNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImportFormatNameEnum(val string) (ImportFormatNameEnum, bool) {
	enum, ok := mappingImportFormatNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ImportFormatVersionEnum Enum with underlying type: string
type ImportFormatVersionEnum string

// Set of constants representing the allowable values for ImportFormatVersionEnum
const (
	ImportFormatVersionV2003 ImportFormatVersionEnum = "V2003"
	ImportFormatVersionV5    ImportFormatVersionEnum = "V5"
)

var mappingImportFormatVersionEnum = map[string]ImportFormatVersionEnum{
	"V2003": ImportFormatVersionV2003,
	"V5":    ImportFormatVersionV5,
}

var mappingImportFormatVersionEnumLowerCase = map[string]ImportFormatVersionEnum{
	"v2003": ImportFormatVersionV2003,
	"v5":    ImportFormatVersionV5,
}

// GetImportFormatVersionEnumValues Enumerates the set of values for ImportFormatVersionEnum
func GetImportFormatVersionEnumValues() []ImportFormatVersionEnum {
	values := make([]ImportFormatVersionEnum, 0)
	for _, v := range mappingImportFormatVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetImportFormatVersionEnumStringValues Enumerates the set of values in String for ImportFormatVersionEnum
func GetImportFormatVersionEnumStringValues() []string {
	return []string{
		"V2003",
		"V5",
	}
}

// GetMappingImportFormatVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImportFormatVersionEnum(val string) (ImportFormatVersionEnum, bool) {
	enum, ok := mappingImportFormatVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
