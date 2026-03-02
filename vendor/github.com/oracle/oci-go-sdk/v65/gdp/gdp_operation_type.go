// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Guarded Data Pipelines API
//
// Use Guarded Data Pipelines to facilitate data transfer between different security domains. The service provides physical, network, and logistical isolation between security domains, malware and vulnerability scanning, auditing, and logging, with deep content inspection capabilities.
//

package gdp

import (
	"strings"
)

// GdpOperationTypeEnum Enum with underlying type: string
type GdpOperationTypeEnum string

// Set of constants representing the allowable values for GdpOperationTypeEnum
const (
	GdpOperationTypeCreateGdpPipeline     GdpOperationTypeEnum = "CREATE_GDP_PIPELINE"
	GdpOperationTypeStartGdpPipeline      GdpOperationTypeEnum = "START_GDP_PIPELINE"
	GdpOperationTypeStopGdpPipeline       GdpOperationTypeEnum = "STOP_GDP_PIPELINE"
	GdpOperationTypeUpdateGdpPipeline     GdpOperationTypeEnum = "UPDATE_GDP_PIPELINE"
	GdpOperationTypeDeleteGdpPipeline     GdpOperationTypeEnum = "DELETE_GDP_PIPELINE"
	GdpOperationTypePeerGdpPipeline       GdpOperationTypeEnum = "PEER_GDP_PIPELINE"
	GdpOperationTypeMoveGdpPipeline       GdpOperationTypeEnum = "MOVE_GDP_PIPELINE"
	GdpOperationTypeRotateGdpPipelineKeys GdpOperationTypeEnum = "ROTATE_GDP_PIPELINE_KEYS"
)

var mappingGdpOperationTypeEnum = map[string]GdpOperationTypeEnum{
	"CREATE_GDP_PIPELINE":      GdpOperationTypeCreateGdpPipeline,
	"START_GDP_PIPELINE":       GdpOperationTypeStartGdpPipeline,
	"STOP_GDP_PIPELINE":        GdpOperationTypeStopGdpPipeline,
	"UPDATE_GDP_PIPELINE":      GdpOperationTypeUpdateGdpPipeline,
	"DELETE_GDP_PIPELINE":      GdpOperationTypeDeleteGdpPipeline,
	"PEER_GDP_PIPELINE":        GdpOperationTypePeerGdpPipeline,
	"MOVE_GDP_PIPELINE":        GdpOperationTypeMoveGdpPipeline,
	"ROTATE_GDP_PIPELINE_KEYS": GdpOperationTypeRotateGdpPipelineKeys,
}

var mappingGdpOperationTypeEnumLowerCase = map[string]GdpOperationTypeEnum{
	"create_gdp_pipeline":      GdpOperationTypeCreateGdpPipeline,
	"start_gdp_pipeline":       GdpOperationTypeStartGdpPipeline,
	"stop_gdp_pipeline":        GdpOperationTypeStopGdpPipeline,
	"update_gdp_pipeline":      GdpOperationTypeUpdateGdpPipeline,
	"delete_gdp_pipeline":      GdpOperationTypeDeleteGdpPipeline,
	"peer_gdp_pipeline":        GdpOperationTypePeerGdpPipeline,
	"move_gdp_pipeline":        GdpOperationTypeMoveGdpPipeline,
	"rotate_gdp_pipeline_keys": GdpOperationTypeRotateGdpPipelineKeys,
}

// GetGdpOperationTypeEnumValues Enumerates the set of values for GdpOperationTypeEnum
func GetGdpOperationTypeEnumValues() []GdpOperationTypeEnum {
	values := make([]GdpOperationTypeEnum, 0)
	for _, v := range mappingGdpOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGdpOperationTypeEnumStringValues Enumerates the set of values in String for GdpOperationTypeEnum
func GetGdpOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_GDP_PIPELINE",
		"START_GDP_PIPELINE",
		"STOP_GDP_PIPELINE",
		"UPDATE_GDP_PIPELINE",
		"DELETE_GDP_PIPELINE",
		"PEER_GDP_PIPELINE",
		"MOVE_GDP_PIPELINE",
		"ROTATE_GDP_PIPELINE_KEYS",
	}
}

// GetMappingGdpOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGdpOperationTypeEnum(val string) (GdpOperationTypeEnum, bool) {
	enum, ok := mappingGdpOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
