// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// RestCallConfig The REST API configuration.
type RestCallConfig struct {

	// The REST method to use.
	MethodType RestCallConfigMethodTypeEnum `mandatory:"false" json:"methodType,omitempty"`

	// The headers for the REST call.
	RequestHeaders map[string]string `mandatory:"false" json:"requestHeaders"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`
}

func (m RestCallConfig) String() string {
	return common.PointerString(m)
}

// RestCallConfigMethodTypeEnum Enum with underlying type: string
type RestCallConfigMethodTypeEnum string

// Set of constants representing the allowable values for RestCallConfigMethodTypeEnum
const (
	RestCallConfigMethodTypeGet    RestCallConfigMethodTypeEnum = "GET"
	RestCallConfigMethodTypePost   RestCallConfigMethodTypeEnum = "POST"
	RestCallConfigMethodTypePatch  RestCallConfigMethodTypeEnum = "PATCH"
	RestCallConfigMethodTypeDelete RestCallConfigMethodTypeEnum = "DELETE"
	RestCallConfigMethodTypePut    RestCallConfigMethodTypeEnum = "PUT"
)

var mappingRestCallConfigMethodType = map[string]RestCallConfigMethodTypeEnum{
	"GET":    RestCallConfigMethodTypeGet,
	"POST":   RestCallConfigMethodTypePost,
	"PATCH":  RestCallConfigMethodTypePatch,
	"DELETE": RestCallConfigMethodTypeDelete,
	"PUT":    RestCallConfigMethodTypePut,
}

// GetRestCallConfigMethodTypeEnumValues Enumerates the set of values for RestCallConfigMethodTypeEnum
func GetRestCallConfigMethodTypeEnumValues() []RestCallConfigMethodTypeEnum {
	values := make([]RestCallConfigMethodTypeEnum, 0)
	for _, v := range mappingRestCallConfigMethodType {
		values = append(values, v)
	}
	return values
}
