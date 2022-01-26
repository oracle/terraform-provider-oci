// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UploadFileSummary Details of Upload File.
type UploadFileSummary struct {

	// Unique internal identifier to refer upload file.
	Reference *string `mandatory:"true" json:"reference"`

	// Name of the file
	Name *string `mandatory:"true" json:"name"`

	// Processing status of the file.
	Status UploadFileSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Number of estimated chunks for this file. A chunk is a portion of the log file used for the processing.
	TotalChunks *float32 `mandatory:"false" json:"totalChunks"`

	// Number of chunks processed.
	ChunksConsumed *float32 `mandatory:"false" json:"chunksConsumed"`

	// Number of chunks processed successfully.
	ChunksSuccess *float32 `mandatory:"false" json:"chunksSuccess"`

	// Number of chunks failed processing.
	ChunksFail *float32 `mandatory:"false" json:"chunksFail"`

	// The time when this file processing started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Name of the log source used for processing this file.
	SourceName *string `mandatory:"false" json:"sourceName"`

	// Name of the entity type.
	EntityType *string `mandatory:"false" json:"entityType"`

	// Name of the entity associated with the file.
	EntityName *string `mandatory:"false" json:"entityName"`

	// (Deprecated) Name of the log namespace associated with the file.
	LogNamespace *string `mandatory:"false" json:"logNamespace"`

	// Log group OCID associated with the file.
	LogGroupId *string `mandatory:"false" json:"logGroupId"`

	// Name of the log group associated with the file.
	LogGroupName *string `mandatory:"false" json:"logGroupName"`

	// The details about upload processing failure.
	FailureDetails *string `mandatory:"false" json:"failureDetails"`
}

func (m UploadFileSummary) String() string {
	return common.PointerString(m)
}

// UploadFileSummaryStatusEnum Enum with underlying type: string
type UploadFileSummaryStatusEnum string

// Set of constants representing the allowable values for UploadFileSummaryStatusEnum
const (
	UploadFileSummaryStatusInProgress UploadFileSummaryStatusEnum = "IN_PROGRESS"
	UploadFileSummaryStatusSuccessful UploadFileSummaryStatusEnum = "SUCCESSFUL"
	UploadFileSummaryStatusFailed     UploadFileSummaryStatusEnum = "FAILED"
)

var mappingUploadFileSummaryStatus = map[string]UploadFileSummaryStatusEnum{
	"IN_PROGRESS": UploadFileSummaryStatusInProgress,
	"SUCCESSFUL":  UploadFileSummaryStatusSuccessful,
	"FAILED":      UploadFileSummaryStatusFailed,
}

// GetUploadFileSummaryStatusEnumValues Enumerates the set of values for UploadFileSummaryStatusEnum
func GetUploadFileSummaryStatusEnumValues() []UploadFileSummaryStatusEnum {
	values := make([]UploadFileSummaryStatusEnum, 0)
	for _, v := range mappingUploadFileSummaryStatus {
		values = append(values, v)
	}
	return values
}
