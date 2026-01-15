// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateProject               OperationTypeEnum = "CREATE_PROJECT"
	OperationTypeUpdateProject               OperationTypeEnum = "UPDATE_PROJECT"
	OperationTypeDeleteProject               OperationTypeEnum = "DELETE_PROJECT"
	OperationTypeMoveProject                 OperationTypeEnum = "MOVE_PROJECT"
	OperationTypeCreateModel                 OperationTypeEnum = "CREATE_MODEL"
	OperationTypeUpdateModel                 OperationTypeEnum = "UPDATE_MODEL"
	OperationTypeDeleteModel                 OperationTypeEnum = "DELETE_MODEL"
	OperationTypeMoveModel                   OperationTypeEnum = "MOVE_MODEL"
	OperationTypeAddStreamSource             OperationTypeEnum = "ADD_STREAM_SOURCE"
	OperationTypeUpdateStreamSource          OperationTypeEnum = "UPDATE_STREAM_SOURCE"
	OperationTypeDeleteStreamSource          OperationTypeEnum = "DELETE_STREAM_SOURCE"
	OperationTypeMoveStreamSource            OperationTypeEnum = "MOVE_STREAM_SOURCE"
	OperationTypeCreateStreamJob             OperationTypeEnum = "CREATE_STREAM_JOB"
	OperationTypeDeleteStreamJob             OperationTypeEnum = "DELETE_STREAM_JOB"
	OperationTypeUpdateStreamJob             OperationTypeEnum = "UPDATE_STREAM_JOB"
	OperationTypeStartStreamJob              OperationTypeEnum = "START_STREAM_JOB"
	OperationTypeStopStreamJob               OperationTypeEnum = "STOP_STREAM_JOB"
	OperationTypeMoveStreamJob               OperationTypeEnum = "MOVE_STREAM_JOB"
	OperationTypeAddStreamGroup              OperationTypeEnum = "ADD_STREAM_GROUP"
	OperationTypeUpdateStreamGroup           OperationTypeEnum = "UPDATE_STREAM_GROUP"
	OperationTypeDeleteStreamGroup           OperationTypeEnum = "DELETE_STREAM_GROUP"
	OperationTypeCreateVisionPrivateEndpoint OperationTypeEnum = "CREATE_VISION_PRIVATE_ENDPOINT"
	OperationTypeUpdateVisionPrivateEndpoint OperationTypeEnum = "UPDATE_VISION_PRIVATE_ENDPOINT"
	OperationTypeDeleteVisionPrivateEndpoint OperationTypeEnum = "DELETE_VISION_PRIVATE_ENDPOINT"
	OperationTypeMoveVisionPrivateEndpoint   OperationTypeEnum = "MOVE_VISION_PRIVATE_ENDPOINT"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_PROJECT":                 OperationTypeCreateProject,
	"UPDATE_PROJECT":                 OperationTypeUpdateProject,
	"DELETE_PROJECT":                 OperationTypeDeleteProject,
	"MOVE_PROJECT":                   OperationTypeMoveProject,
	"CREATE_MODEL":                   OperationTypeCreateModel,
	"UPDATE_MODEL":                   OperationTypeUpdateModel,
	"DELETE_MODEL":                   OperationTypeDeleteModel,
	"MOVE_MODEL":                     OperationTypeMoveModel,
	"ADD_STREAM_SOURCE":              OperationTypeAddStreamSource,
	"UPDATE_STREAM_SOURCE":           OperationTypeUpdateStreamSource,
	"DELETE_STREAM_SOURCE":           OperationTypeDeleteStreamSource,
	"MOVE_STREAM_SOURCE":             OperationTypeMoveStreamSource,
	"CREATE_STREAM_JOB":              OperationTypeCreateStreamJob,
	"DELETE_STREAM_JOB":              OperationTypeDeleteStreamJob,
	"UPDATE_STREAM_JOB":              OperationTypeUpdateStreamJob,
	"START_STREAM_JOB":               OperationTypeStartStreamJob,
	"STOP_STREAM_JOB":                OperationTypeStopStreamJob,
	"MOVE_STREAM_JOB":                OperationTypeMoveStreamJob,
	"ADD_STREAM_GROUP":               OperationTypeAddStreamGroup,
	"UPDATE_STREAM_GROUP":            OperationTypeUpdateStreamGroup,
	"DELETE_STREAM_GROUP":            OperationTypeDeleteStreamGroup,
	"CREATE_VISION_PRIVATE_ENDPOINT": OperationTypeCreateVisionPrivateEndpoint,
	"UPDATE_VISION_PRIVATE_ENDPOINT": OperationTypeUpdateVisionPrivateEndpoint,
	"DELETE_VISION_PRIVATE_ENDPOINT": OperationTypeDeleteVisionPrivateEndpoint,
	"MOVE_VISION_PRIVATE_ENDPOINT":   OperationTypeMoveVisionPrivateEndpoint,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_project":                 OperationTypeCreateProject,
	"update_project":                 OperationTypeUpdateProject,
	"delete_project":                 OperationTypeDeleteProject,
	"move_project":                   OperationTypeMoveProject,
	"create_model":                   OperationTypeCreateModel,
	"update_model":                   OperationTypeUpdateModel,
	"delete_model":                   OperationTypeDeleteModel,
	"move_model":                     OperationTypeMoveModel,
	"add_stream_source":              OperationTypeAddStreamSource,
	"update_stream_source":           OperationTypeUpdateStreamSource,
	"delete_stream_source":           OperationTypeDeleteStreamSource,
	"move_stream_source":             OperationTypeMoveStreamSource,
	"create_stream_job":              OperationTypeCreateStreamJob,
	"delete_stream_job":              OperationTypeDeleteStreamJob,
	"update_stream_job":              OperationTypeUpdateStreamJob,
	"start_stream_job":               OperationTypeStartStreamJob,
	"stop_stream_job":                OperationTypeStopStreamJob,
	"move_stream_job":                OperationTypeMoveStreamJob,
	"add_stream_group":               OperationTypeAddStreamGroup,
	"update_stream_group":            OperationTypeUpdateStreamGroup,
	"delete_stream_group":            OperationTypeDeleteStreamGroup,
	"create_vision_private_endpoint": OperationTypeCreateVisionPrivateEndpoint,
	"update_vision_private_endpoint": OperationTypeUpdateVisionPrivateEndpoint,
	"delete_vision_private_endpoint": OperationTypeDeleteVisionPrivateEndpoint,
	"move_vision_private_endpoint":   OperationTypeMoveVisionPrivateEndpoint,
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
		"CREATE_PROJECT",
		"UPDATE_PROJECT",
		"DELETE_PROJECT",
		"MOVE_PROJECT",
		"CREATE_MODEL",
		"UPDATE_MODEL",
		"DELETE_MODEL",
		"MOVE_MODEL",
		"ADD_STREAM_SOURCE",
		"UPDATE_STREAM_SOURCE",
		"DELETE_STREAM_SOURCE",
		"MOVE_STREAM_SOURCE",
		"CREATE_STREAM_JOB",
		"DELETE_STREAM_JOB",
		"UPDATE_STREAM_JOB",
		"START_STREAM_JOB",
		"STOP_STREAM_JOB",
		"MOVE_STREAM_JOB",
		"ADD_STREAM_GROUP",
		"UPDATE_STREAM_GROUP",
		"DELETE_STREAM_GROUP",
		"CREATE_VISION_PRIVATE_ENDPOINT",
		"UPDATE_VISION_PRIVATE_ENDPOINT",
		"DELETE_VISION_PRIVATE_ENDPOINT",
		"MOVE_VISION_PRIVATE_ENDPOINT",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
