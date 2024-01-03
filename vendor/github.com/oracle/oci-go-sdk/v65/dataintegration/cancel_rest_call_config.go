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

// CancelRestCallConfig The REST API configuration for cancelling the task.
type CancelRestCallConfig struct {

	// The REST method to use.
	MethodType CancelRestCallConfigMethodTypeEnum `mandatory:"false" json:"methodType,omitempty"`

	// The headers for the REST call.
	RequestHeaders map[string]string `mandatory:"false" json:"requestHeaders"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`
}

func (m CancelRestCallConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CancelRestCallConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCancelRestCallConfigMethodTypeEnum(string(m.MethodType)); !ok && m.MethodType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MethodType: %s. Supported values are: %s.", m.MethodType, strings.Join(GetCancelRestCallConfigMethodTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CancelRestCallConfigMethodTypeEnum Enum with underlying type: string
type CancelRestCallConfigMethodTypeEnum string

// Set of constants representing the allowable values for CancelRestCallConfigMethodTypeEnum
const (
	CancelRestCallConfigMethodTypeGet    CancelRestCallConfigMethodTypeEnum = "GET"
	CancelRestCallConfigMethodTypePost   CancelRestCallConfigMethodTypeEnum = "POST"
	CancelRestCallConfigMethodTypePatch  CancelRestCallConfigMethodTypeEnum = "PATCH"
	CancelRestCallConfigMethodTypeDelete CancelRestCallConfigMethodTypeEnum = "DELETE"
	CancelRestCallConfigMethodTypePut    CancelRestCallConfigMethodTypeEnum = "PUT"
)

var mappingCancelRestCallConfigMethodTypeEnum = map[string]CancelRestCallConfigMethodTypeEnum{
	"GET":    CancelRestCallConfigMethodTypeGet,
	"POST":   CancelRestCallConfigMethodTypePost,
	"PATCH":  CancelRestCallConfigMethodTypePatch,
	"DELETE": CancelRestCallConfigMethodTypeDelete,
	"PUT":    CancelRestCallConfigMethodTypePut,
}

var mappingCancelRestCallConfigMethodTypeEnumLowerCase = map[string]CancelRestCallConfigMethodTypeEnum{
	"get":    CancelRestCallConfigMethodTypeGet,
	"post":   CancelRestCallConfigMethodTypePost,
	"patch":  CancelRestCallConfigMethodTypePatch,
	"delete": CancelRestCallConfigMethodTypeDelete,
	"put":    CancelRestCallConfigMethodTypePut,
}

// GetCancelRestCallConfigMethodTypeEnumValues Enumerates the set of values for CancelRestCallConfigMethodTypeEnum
func GetCancelRestCallConfigMethodTypeEnumValues() []CancelRestCallConfigMethodTypeEnum {
	values := make([]CancelRestCallConfigMethodTypeEnum, 0)
	for _, v := range mappingCancelRestCallConfigMethodTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCancelRestCallConfigMethodTypeEnumStringValues Enumerates the set of values in String for CancelRestCallConfigMethodTypeEnum
func GetCancelRestCallConfigMethodTypeEnumStringValues() []string {
	return []string{
		"GET",
		"POST",
		"PATCH",
		"DELETE",
		"PUT",
	}
}

// GetMappingCancelRestCallConfigMethodTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCancelRestCallConfigMethodTypeEnum(val string) (CancelRestCallConfigMethodTypeEnum, bool) {
	enum, ok := mappingCancelRestCallConfigMethodTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
