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

// PollRestCallConfig The REST API configuration for polling.
type PollRestCallConfig struct {

	// The REST method to use.
	MethodType PollRestCallConfigMethodTypeEnum `mandatory:"false" json:"methodType,omitempty"`

	// The headers for the REST call.
	RequestHeaders map[string]string `mandatory:"false" json:"requestHeaders"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`
}

func (m PollRestCallConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PollRestCallConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPollRestCallConfigMethodTypeEnum(string(m.MethodType)); !ok && m.MethodType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MethodType: %s. Supported values are: %s.", m.MethodType, strings.Join(GetPollRestCallConfigMethodTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PollRestCallConfigMethodTypeEnum Enum with underlying type: string
type PollRestCallConfigMethodTypeEnum string

// Set of constants representing the allowable values for PollRestCallConfigMethodTypeEnum
const (
	PollRestCallConfigMethodTypeGet    PollRestCallConfigMethodTypeEnum = "GET"
	PollRestCallConfigMethodTypePost   PollRestCallConfigMethodTypeEnum = "POST"
	PollRestCallConfigMethodTypePatch  PollRestCallConfigMethodTypeEnum = "PATCH"
	PollRestCallConfigMethodTypeDelete PollRestCallConfigMethodTypeEnum = "DELETE"
	PollRestCallConfigMethodTypePut    PollRestCallConfigMethodTypeEnum = "PUT"
)

var mappingPollRestCallConfigMethodTypeEnum = map[string]PollRestCallConfigMethodTypeEnum{
	"GET":    PollRestCallConfigMethodTypeGet,
	"POST":   PollRestCallConfigMethodTypePost,
	"PATCH":  PollRestCallConfigMethodTypePatch,
	"DELETE": PollRestCallConfigMethodTypeDelete,
	"PUT":    PollRestCallConfigMethodTypePut,
}

var mappingPollRestCallConfigMethodTypeEnumLowerCase = map[string]PollRestCallConfigMethodTypeEnum{
	"get":    PollRestCallConfigMethodTypeGet,
	"post":   PollRestCallConfigMethodTypePost,
	"patch":  PollRestCallConfigMethodTypePatch,
	"delete": PollRestCallConfigMethodTypeDelete,
	"put":    PollRestCallConfigMethodTypePut,
}

// GetPollRestCallConfigMethodTypeEnumValues Enumerates the set of values for PollRestCallConfigMethodTypeEnum
func GetPollRestCallConfigMethodTypeEnumValues() []PollRestCallConfigMethodTypeEnum {
	values := make([]PollRestCallConfigMethodTypeEnum, 0)
	for _, v := range mappingPollRestCallConfigMethodTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPollRestCallConfigMethodTypeEnumStringValues Enumerates the set of values in String for PollRestCallConfigMethodTypeEnum
func GetPollRestCallConfigMethodTypeEnumStringValues() []string {
	return []string{
		"GET",
		"POST",
		"PATCH",
		"DELETE",
		"PUT",
	}
}

// GetMappingPollRestCallConfigMethodTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPollRestCallConfigMethodTypeEnum(val string) (PollRestCallConfigMethodTypeEnum, bool) {
	enum, ok := mappingPollRestCallConfigMethodTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
