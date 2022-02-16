// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SnapshotExportDetails Specifies where to output the export.
type SnapshotExportDetails interface {
}

type snapshotexportdetails struct {
	JsonData   []byte
	ExportType string `json:"exportType"`
}

// UnmarshalJSON unmarshals json
func (m *snapshotexportdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersnapshotexportdetails snapshotexportdetails
	s := struct {
		Model Unmarshalersnapshotexportdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ExportType = s.Model.ExportType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *snapshotexportdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ExportType {
	case "OBJECT_STORAGE":
		mm := ObjectStorageSnapshotExportDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m snapshotexportdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m snapshotexportdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SnapshotExportDetailsExportTypeEnum Enum with underlying type: string
type SnapshotExportDetailsExportTypeEnum string

// Set of constants representing the allowable values for SnapshotExportDetailsExportTypeEnum
const (
	SnapshotExportDetailsExportTypeObjectStorage SnapshotExportDetailsExportTypeEnum = "OBJECT_STORAGE"
)

var mappingSnapshotExportDetailsExportTypeEnum = map[string]SnapshotExportDetailsExportTypeEnum{
	"OBJECT_STORAGE": SnapshotExportDetailsExportTypeObjectStorage,
}

// GetSnapshotExportDetailsExportTypeEnumValues Enumerates the set of values for SnapshotExportDetailsExportTypeEnum
func GetSnapshotExportDetailsExportTypeEnumValues() []SnapshotExportDetailsExportTypeEnum {
	values := make([]SnapshotExportDetailsExportTypeEnum, 0)
	for _, v := range mappingSnapshotExportDetailsExportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSnapshotExportDetailsExportTypeEnumStringValues Enumerates the set of values in String for SnapshotExportDetailsExportTypeEnum
func GetSnapshotExportDetailsExportTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingSnapshotExportDetailsExportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSnapshotExportDetailsExportTypeEnum(val string) (SnapshotExportDetailsExportTypeEnum, bool) {
	mappingSnapshotExportDetailsExportTypeEnumIgnoreCase := make(map[string]SnapshotExportDetailsExportTypeEnum)
	for k, v := range mappingSnapshotExportDetailsExportTypeEnum {
		mappingSnapshotExportDetailsExportTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSnapshotExportDetailsExportTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
