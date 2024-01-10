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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RestCallConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRestCallConfigMethodTypeEnum(string(m.MethodType)); !ok && m.MethodType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MethodType: %s. Supported values are: %s.", m.MethodType, strings.Join(GetRestCallConfigMethodTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingRestCallConfigMethodTypeEnum = map[string]RestCallConfigMethodTypeEnum{
	"GET":    RestCallConfigMethodTypeGet,
	"POST":   RestCallConfigMethodTypePost,
	"PATCH":  RestCallConfigMethodTypePatch,
	"DELETE": RestCallConfigMethodTypeDelete,
	"PUT":    RestCallConfigMethodTypePut,
}

var mappingRestCallConfigMethodTypeEnumLowerCase = map[string]RestCallConfigMethodTypeEnum{
	"get":    RestCallConfigMethodTypeGet,
	"post":   RestCallConfigMethodTypePost,
	"patch":  RestCallConfigMethodTypePatch,
	"delete": RestCallConfigMethodTypeDelete,
	"put":    RestCallConfigMethodTypePut,
}

// GetRestCallConfigMethodTypeEnumValues Enumerates the set of values for RestCallConfigMethodTypeEnum
func GetRestCallConfigMethodTypeEnumValues() []RestCallConfigMethodTypeEnum {
	values := make([]RestCallConfigMethodTypeEnum, 0)
	for _, v := range mappingRestCallConfigMethodTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRestCallConfigMethodTypeEnumStringValues Enumerates the set of values in String for RestCallConfigMethodTypeEnum
func GetRestCallConfigMethodTypeEnumStringValues() []string {
	return []string{
		"GET",
		"POST",
		"PATCH",
		"DELETE",
		"PUT",
	}
}

// GetMappingRestCallConfigMethodTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRestCallConfigMethodTypeEnum(val string) (RestCallConfigMethodTypeEnum, bool) {
	enum, ok := mappingRestCallConfigMethodTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
