// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecuteRestCallConfig The REST API configuration for execution.
type ExecuteRestCallConfig struct {

	// The REST method to use.
	MethodType ExecuteRestCallConfigMethodTypeEnum `mandatory:"false" json:"methodType,omitempty"`

	// The headers for the REST call.
	RequestHeaders map[string]string `mandatory:"false" json:"requestHeaders"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`
}

func (m ExecuteRestCallConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteRestCallConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExecuteRestCallConfigMethodTypeEnum(string(m.MethodType)); !ok && m.MethodType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MethodType: %s. Supported values are: %s.", m.MethodType, strings.Join(GetExecuteRestCallConfigMethodTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExecuteRestCallConfigMethodTypeEnum Enum with underlying type: string
type ExecuteRestCallConfigMethodTypeEnum string

// Set of constants representing the allowable values for ExecuteRestCallConfigMethodTypeEnum
const (
	ExecuteRestCallConfigMethodTypeGet    ExecuteRestCallConfigMethodTypeEnum = "GET"
	ExecuteRestCallConfigMethodTypePost   ExecuteRestCallConfigMethodTypeEnum = "POST"
	ExecuteRestCallConfigMethodTypePatch  ExecuteRestCallConfigMethodTypeEnum = "PATCH"
	ExecuteRestCallConfigMethodTypeDelete ExecuteRestCallConfigMethodTypeEnum = "DELETE"
	ExecuteRestCallConfigMethodTypePut    ExecuteRestCallConfigMethodTypeEnum = "PUT"
)

var mappingExecuteRestCallConfigMethodTypeEnum = map[string]ExecuteRestCallConfigMethodTypeEnum{
	"GET":    ExecuteRestCallConfigMethodTypeGet,
	"POST":   ExecuteRestCallConfigMethodTypePost,
	"PATCH":  ExecuteRestCallConfigMethodTypePatch,
	"DELETE": ExecuteRestCallConfigMethodTypeDelete,
	"PUT":    ExecuteRestCallConfigMethodTypePut,
}

var mappingExecuteRestCallConfigMethodTypeEnumLowerCase = map[string]ExecuteRestCallConfigMethodTypeEnum{
	"get":    ExecuteRestCallConfigMethodTypeGet,
	"post":   ExecuteRestCallConfigMethodTypePost,
	"patch":  ExecuteRestCallConfigMethodTypePatch,
	"delete": ExecuteRestCallConfigMethodTypeDelete,
	"put":    ExecuteRestCallConfigMethodTypePut,
}

// GetExecuteRestCallConfigMethodTypeEnumValues Enumerates the set of values for ExecuteRestCallConfigMethodTypeEnum
func GetExecuteRestCallConfigMethodTypeEnumValues() []ExecuteRestCallConfigMethodTypeEnum {
	values := make([]ExecuteRestCallConfigMethodTypeEnum, 0)
	for _, v := range mappingExecuteRestCallConfigMethodTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExecuteRestCallConfigMethodTypeEnumStringValues Enumerates the set of values in String for ExecuteRestCallConfigMethodTypeEnum
func GetExecuteRestCallConfigMethodTypeEnumStringValues() []string {
	return []string{
		"GET",
		"POST",
		"PATCH",
		"DELETE",
		"PUT",
	}
}

// GetMappingExecuteRestCallConfigMethodTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecuteRestCallConfigMethodTypeEnum(val string) (ExecuteRestCallConfigMethodTypeEnum, bool) {
	enum, ok := mappingExecuteRestCallConfigMethodTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
