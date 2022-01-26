// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

var mappingExportFormatName = map[string]ExportFormatNameEnum{
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
	for _, v := range mappingExportFormatName {
		values = append(values, v)
	}
	return values
}

// ExportFormatVersionEnum Enum with underlying type: string
type ExportFormatVersionEnum string

// Set of constants representing the allowable values for ExportFormatVersionEnum
const (
	ExportFormatVersionV2003 ExportFormatVersionEnum = "V2003"
	ExportFormatVersionV5    ExportFormatVersionEnum = "V5"
)

var mappingExportFormatVersion = map[string]ExportFormatVersionEnum{
	"V2003": ExportFormatVersionV2003,
	"V5":    ExportFormatVersionV5,
}

// GetExportFormatVersionEnumValues Enumerates the set of values for ExportFormatVersionEnum
func GetExportFormatVersionEnumValues() []ExportFormatVersionEnum {
	values := make([]ExportFormatVersionEnum, 0)
	for _, v := range mappingExportFormatVersion {
		values = append(values, v)
	}
	return values
}
