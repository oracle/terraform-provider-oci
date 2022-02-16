// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ExportFormat Specifies the export format to be used for exporting snapshot.
type ExportFormat struct {

	// Name of export format.
	Name ExportFormatNameEnum `mandatory:"false" json:"name,omitempty"`

	// Version of export format.
	Version ExportFormatVersionEnum `mandatory:"false" json:"version,omitempty"`
}

func (m ExportFormat) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportFormat) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExportFormatNameEnum(string(m.Name)); !ok && m.Name != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Name: %s. Supported values are: %s.", m.Name, strings.Join(GetExportFormatNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExportFormatVersionEnum(string(m.Version)); !ok && m.Version != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Version: %s. Supported values are: %s.", m.Version, strings.Join(GetExportFormatVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExportFormatNameEnum Enum with underlying type: string
type ExportFormatNameEnum string

// Set of constants representing the allowable values for ExportFormatNameEnum
const (
	ExportFormatNameJsonl             ExportFormatNameEnum = "JSONL"
	ExportFormatNameJsonlConsolidated ExportFormatNameEnum = "JSONL_CONSOLIDATED"
	ExportFormatNameConll             ExportFormatNameEnum = "CONLL"
	ExportFormatNameSpacy             ExportFormatNameEnum = "SPACY"
	ExportFormatNameCoco              ExportFormatNameEnum = "COCO"
	ExportFormatNameYolo              ExportFormatNameEnum = "YOLO"
	ExportFormatNamePascalVoc         ExportFormatNameEnum = "PASCAL_VOC"
)

var mappingExportFormatNameEnum = map[string]ExportFormatNameEnum{
	"JSONL":              ExportFormatNameJsonl,
	"JSONL_CONSOLIDATED": ExportFormatNameJsonlConsolidated,
	"CONLL":              ExportFormatNameConll,
	"SPACY":              ExportFormatNameSpacy,
	"COCO":               ExportFormatNameCoco,
	"YOLO":               ExportFormatNameYolo,
	"PASCAL_VOC":         ExportFormatNamePascalVoc,
}

// GetExportFormatNameEnumValues Enumerates the set of values for ExportFormatNameEnum
func GetExportFormatNameEnumValues() []ExportFormatNameEnum {
	values := make([]ExportFormatNameEnum, 0)
	for _, v := range mappingExportFormatNameEnum {
		values = append(values, v)
	}
	return values
}

// GetExportFormatNameEnumStringValues Enumerates the set of values in String for ExportFormatNameEnum
func GetExportFormatNameEnumStringValues() []string {
	return []string{
		"JSONL",
		"JSONL_CONSOLIDATED",
		"CONLL",
		"SPACY",
		"COCO",
		"YOLO",
		"PASCAL_VOC",
	}
}

// GetMappingExportFormatNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportFormatNameEnum(val string) (ExportFormatNameEnum, bool) {
	mappingExportFormatNameEnumIgnoreCase := make(map[string]ExportFormatNameEnum)
	for k, v := range mappingExportFormatNameEnum {
		mappingExportFormatNameEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingExportFormatNameEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ExportFormatVersionEnum Enum with underlying type: string
type ExportFormatVersionEnum string

// Set of constants representing the allowable values for ExportFormatVersionEnum
const (
	ExportFormatVersionV2003 ExportFormatVersionEnum = "V2003"
	ExportFormatVersionV5    ExportFormatVersionEnum = "V5"
)

var mappingExportFormatVersionEnum = map[string]ExportFormatVersionEnum{
	"V2003": ExportFormatVersionV2003,
	"V5":    ExportFormatVersionV5,
}

// GetExportFormatVersionEnumValues Enumerates the set of values for ExportFormatVersionEnum
func GetExportFormatVersionEnumValues() []ExportFormatVersionEnum {
	values := make([]ExportFormatVersionEnum, 0)
	for _, v := range mappingExportFormatVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetExportFormatVersionEnumStringValues Enumerates the set of values in String for ExportFormatVersionEnum
func GetExportFormatVersionEnumStringValues() []string {
	return []string{
		"V2003",
		"V5",
	}
}

// GetMappingExportFormatVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportFormatVersionEnum(val string) (ExportFormatVersionEnum, bool) {
	mappingExportFormatVersionEnumIgnoreCase := make(map[string]ExportFormatVersionEnum)
	for k, v := range mappingExportFormatVersionEnum {
		mappingExportFormatVersionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingExportFormatVersionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
