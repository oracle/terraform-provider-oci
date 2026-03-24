// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FileSyncIngestionLogs A Log object that gives the ingestion status of a File from a datasource read by a VectorStoreConnector
type FileSyncIngestionLogs struct {

	// An identifier that identifies a File ingested to a VectorStore.
	FileId *string `mandatory:"true" json:"fileId"`

	// The path where the file was located in the datasource.
	FilePath *string `mandatory:"true" json:"filePath"`

	// The current status of ingestion for the File
	Status FileSyncIngestionLogsStatusEnum `mandatory:"true" json:"status"`

	// The size of the file.
	FileSizeInBytes *int `mandatory:"false" json:"fileSizeInBytes"`

	// The total number of chunks created from the File.
	TotalChunksCreated *int `mandatory:"false" json:"totalChunksCreated"`

	// The duration taken(in seconds) to ingest the File.
	DurationInSeconds *int `mandatory:"false" json:"durationInSeconds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m FileSyncIngestionLogs) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FileSyncIngestionLogs) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFileSyncIngestionLogsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetFileSyncIngestionLogsStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FileSyncIngestionLogsStatusEnum Enum with underlying type: string
type FileSyncIngestionLogsStatusEnum string

// Set of constants representing the allowable values for FileSyncIngestionLogsStatusEnum
const (
	FileSyncIngestionLogsStatusCompleted  FileSyncIngestionLogsStatusEnum = "COMPLETED"
	FileSyncIngestionLogsStatusCancelled  FileSyncIngestionLogsStatusEnum = "CANCELLED"
	FileSyncIngestionLogsStatusFailed     FileSyncIngestionLogsStatusEnum = "FAILED"
	FileSyncIngestionLogsStatusQueued     FileSyncIngestionLogsStatusEnum = "QUEUED"
	FileSyncIngestionLogsStatusInProgress FileSyncIngestionLogsStatusEnum = "IN_PROGRESS"
)

var mappingFileSyncIngestionLogsStatusEnum = map[string]FileSyncIngestionLogsStatusEnum{
	"COMPLETED":   FileSyncIngestionLogsStatusCompleted,
	"CANCELLED":   FileSyncIngestionLogsStatusCancelled,
	"FAILED":      FileSyncIngestionLogsStatusFailed,
	"QUEUED":      FileSyncIngestionLogsStatusQueued,
	"IN_PROGRESS": FileSyncIngestionLogsStatusInProgress,
}

var mappingFileSyncIngestionLogsStatusEnumLowerCase = map[string]FileSyncIngestionLogsStatusEnum{
	"completed":   FileSyncIngestionLogsStatusCompleted,
	"cancelled":   FileSyncIngestionLogsStatusCancelled,
	"failed":      FileSyncIngestionLogsStatusFailed,
	"queued":      FileSyncIngestionLogsStatusQueued,
	"in_progress": FileSyncIngestionLogsStatusInProgress,
}

// GetFileSyncIngestionLogsStatusEnumValues Enumerates the set of values for FileSyncIngestionLogsStatusEnum
func GetFileSyncIngestionLogsStatusEnumValues() []FileSyncIngestionLogsStatusEnum {
	values := make([]FileSyncIngestionLogsStatusEnum, 0)
	for _, v := range mappingFileSyncIngestionLogsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetFileSyncIngestionLogsStatusEnumStringValues Enumerates the set of values in String for FileSyncIngestionLogsStatusEnum
func GetFileSyncIngestionLogsStatusEnumStringValues() []string {
	return []string{
		"COMPLETED",
		"CANCELLED",
		"FAILED",
		"QUEUED",
		"IN_PROGRESS",
	}
}

// GetMappingFileSyncIngestionLogsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFileSyncIngestionLogsStatusEnum(val string) (FileSyncIngestionLogsStatusEnum, bool) {
	enum, ok := mappingFileSyncIngestionLogsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
