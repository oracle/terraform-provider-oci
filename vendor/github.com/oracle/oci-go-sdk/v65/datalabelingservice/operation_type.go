// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateDataset          OperationTypeEnum = "CREATE_DATASET"
	OperationTypeDeleteDataset          OperationTypeEnum = "DELETE_DATASET"
	OperationTypeMoveDataset            OperationTypeEnum = "MOVE_DATASET"
	OperationTypeGenerateDatasetRecords OperationTypeEnum = "GENERATE_DATASET_RECORDS"
	OperationTypeSnapshotDataset        OperationTypeEnum = "SNAPSHOT_DATASET"
	OperationTypeAddDatasetLabels       OperationTypeEnum = "ADD_DATASET_LABELS"
	OperationTypeRemoveDatasetLabels    OperationTypeEnum = "REMOVE_DATASET_LABELS"
	OperationTypeRenameDatasetLabels    OperationTypeEnum = "RENAME_DATASET_LABELS"
	OperationTypeImportDataset          OperationTypeEnum = "IMPORT_DATASET"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_DATASET":           OperationTypeCreateDataset,
	"DELETE_DATASET":           OperationTypeDeleteDataset,
	"MOVE_DATASET":             OperationTypeMoveDataset,
	"GENERATE_DATASET_RECORDS": OperationTypeGenerateDatasetRecords,
	"SNAPSHOT_DATASET":         OperationTypeSnapshotDataset,
	"ADD_DATASET_LABELS":       OperationTypeAddDatasetLabels,
	"REMOVE_DATASET_LABELS":    OperationTypeRemoveDatasetLabels,
	"RENAME_DATASET_LABELS":    OperationTypeRenameDatasetLabels,
	"IMPORT_DATASET":           OperationTypeImportDataset,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_dataset":           OperationTypeCreateDataset,
	"delete_dataset":           OperationTypeDeleteDataset,
	"move_dataset":             OperationTypeMoveDataset,
	"generate_dataset_records": OperationTypeGenerateDatasetRecords,
	"snapshot_dataset":         OperationTypeSnapshotDataset,
	"add_dataset_labels":       OperationTypeAddDatasetLabels,
	"remove_dataset_labels":    OperationTypeRemoveDatasetLabels,
	"rename_dataset_labels":    OperationTypeRenameDatasetLabels,
	"import_dataset":           OperationTypeImportDataset,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_DATASET",
		"DELETE_DATASET",
		"MOVE_DATASET",
		"GENERATE_DATASET_RECORDS",
		"SNAPSHOT_DATASET",
		"ADD_DATASET_LABELS",
		"REMOVE_DATASET_LABELS",
		"RENAME_DATASET_LABELS",
		"IMPORT_DATASET",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
