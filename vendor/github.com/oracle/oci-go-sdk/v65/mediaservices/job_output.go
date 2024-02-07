// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.cloud.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.cloud.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JobOutput The output result of an executed MediaWorkflowJob.
type JobOutput struct {

	// Type of job output.
	AssetType JobOutputAssetTypeEnum `mandatory:"false" json:"assetType,omitempty"`

	// The namespace name of the job output.
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// The bucket name of the job output.
	BucketName *string `mandatory:"false" json:"bucketName"`

	// The object name of the job output.
	ObjectName *string `mandatory:"false" json:"objectName"`

	// The ID associated with the job output.
	Id *string `mandatory:"false" json:"id"`
}

func (m JobOutput) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobOutput) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJobOutputAssetTypeEnum(string(m.AssetType)); !ok && m.AssetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssetType: %s. Supported values are: %s.", m.AssetType, strings.Join(GetJobOutputAssetTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobOutputAssetTypeEnum Enum with underlying type: string
type JobOutputAssetTypeEnum string

// Set of constants representing the allowable values for JobOutputAssetTypeEnum
const (
	JobOutputAssetTypeAudio            JobOutputAssetTypeEnum = "AUDIO"
	JobOutputAssetTypeVideo            JobOutputAssetTypeEnum = "VIDEO"
	JobOutputAssetTypePlaylist         JobOutputAssetTypeEnum = "PLAYLIST"
	JobOutputAssetTypeImage            JobOutputAssetTypeEnum = "IMAGE"
	JobOutputAssetTypeCaptionFile      JobOutputAssetTypeEnum = "CAPTION_FILE"
	JobOutputAssetTypeTranscriptionJob JobOutputAssetTypeEnum = "TRANSCRIPTION_JOB"
	JobOutputAssetTypeVisionJob        JobOutputAssetTypeEnum = "VISION_JOB"
	JobOutputAssetTypeTextAnalysis     JobOutputAssetTypeEnum = "TEXT_ANALYSIS"
	JobOutputAssetTypeInputFile        JobOutputAssetTypeEnum = "INPUT_FILE"
	JobOutputAssetTypeOther            JobOutputAssetTypeEnum = "OTHER"
)

var mappingJobOutputAssetTypeEnum = map[string]JobOutputAssetTypeEnum{
	"AUDIO":             JobOutputAssetTypeAudio,
	"VIDEO":             JobOutputAssetTypeVideo,
	"PLAYLIST":          JobOutputAssetTypePlaylist,
	"IMAGE":             JobOutputAssetTypeImage,
	"CAPTION_FILE":      JobOutputAssetTypeCaptionFile,
	"TRANSCRIPTION_JOB": JobOutputAssetTypeTranscriptionJob,
	"VISION_JOB":        JobOutputAssetTypeVisionJob,
	"TEXT_ANALYSIS":     JobOutputAssetTypeTextAnalysis,
	"INPUT_FILE":        JobOutputAssetTypeInputFile,
	"OTHER":             JobOutputAssetTypeOther,
}

var mappingJobOutputAssetTypeEnumLowerCase = map[string]JobOutputAssetTypeEnum{
	"audio":             JobOutputAssetTypeAudio,
	"video":             JobOutputAssetTypeVideo,
	"playlist":          JobOutputAssetTypePlaylist,
	"image":             JobOutputAssetTypeImage,
	"caption_file":      JobOutputAssetTypeCaptionFile,
	"transcription_job": JobOutputAssetTypeTranscriptionJob,
	"vision_job":        JobOutputAssetTypeVisionJob,
	"text_analysis":     JobOutputAssetTypeTextAnalysis,
	"input_file":        JobOutputAssetTypeInputFile,
	"other":             JobOutputAssetTypeOther,
}

// GetJobOutputAssetTypeEnumValues Enumerates the set of values for JobOutputAssetTypeEnum
func GetJobOutputAssetTypeEnumValues() []JobOutputAssetTypeEnum {
	values := make([]JobOutputAssetTypeEnum, 0)
	for _, v := range mappingJobOutputAssetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobOutputAssetTypeEnumStringValues Enumerates the set of values in String for JobOutputAssetTypeEnum
func GetJobOutputAssetTypeEnumStringValues() []string {
	return []string{
		"AUDIO",
		"VIDEO",
		"PLAYLIST",
		"IMAGE",
		"CAPTION_FILE",
		"TRANSCRIPTION_JOB",
		"VISION_JOB",
		"TEXT_ANALYSIS",
		"INPUT_FILE",
		"OTHER",
	}
}

// GetMappingJobOutputAssetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobOutputAssetTypeEnum(val string) (JobOutputAssetTypeEnum, bool) {
	enum, ok := mappingJobOutputAssetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
