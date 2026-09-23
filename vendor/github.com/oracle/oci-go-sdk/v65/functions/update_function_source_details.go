// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateFunctionSourceDetails The source details for updating the Function.
type UpdateFunctionSourceDetails interface {
}

type updatefunctionsourcedetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *updatefunctionsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatefunctionsourcedetails updatefunctionsourcedetails
	s := struct {
		Model Unmarshalerupdatefunctionsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatefunctionsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "CONTAINER_IMAGE":
		mm := UpdateContainerImageFunctionSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ARCHIVE":
		mm := UpdateArchiveFunctionSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateFunctionSourceDetails: %s.", m.SourceType)
		return *m, nil
	}
}

func (m updatefunctionsourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatefunctionsourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateFunctionSourceDetailsSourceTypeEnum Enum with underlying type: string
type UpdateFunctionSourceDetailsSourceTypeEnum string

// Set of constants representing the allowable values for UpdateFunctionSourceDetailsSourceTypeEnum
const (
	UpdateFunctionSourceDetailsSourceTypeArchive        UpdateFunctionSourceDetailsSourceTypeEnum = "ARCHIVE"
	UpdateFunctionSourceDetailsSourceTypeContainerImage UpdateFunctionSourceDetailsSourceTypeEnum = "CONTAINER_IMAGE"
)

var mappingUpdateFunctionSourceDetailsSourceTypeEnum = map[string]UpdateFunctionSourceDetailsSourceTypeEnum{
	"ARCHIVE":         UpdateFunctionSourceDetailsSourceTypeArchive,
	"CONTAINER_IMAGE": UpdateFunctionSourceDetailsSourceTypeContainerImage,
}

var mappingUpdateFunctionSourceDetailsSourceTypeEnumLowerCase = map[string]UpdateFunctionSourceDetailsSourceTypeEnum{
	"archive":         UpdateFunctionSourceDetailsSourceTypeArchive,
	"container_image": UpdateFunctionSourceDetailsSourceTypeContainerImage,
}

// GetUpdateFunctionSourceDetailsSourceTypeEnumValues Enumerates the set of values for UpdateFunctionSourceDetailsSourceTypeEnum
func GetUpdateFunctionSourceDetailsSourceTypeEnumValues() []UpdateFunctionSourceDetailsSourceTypeEnum {
	values := make([]UpdateFunctionSourceDetailsSourceTypeEnum, 0)
	for _, v := range mappingUpdateFunctionSourceDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateFunctionSourceDetailsSourceTypeEnumStringValues Enumerates the set of values in String for UpdateFunctionSourceDetailsSourceTypeEnum
func GetUpdateFunctionSourceDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"ARCHIVE",
		"CONTAINER_IMAGE",
	}
}

// GetMappingUpdateFunctionSourceDetailsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateFunctionSourceDetailsSourceTypeEnum(val string) (UpdateFunctionSourceDetailsSourceTypeEnum, bool) {
	enum, ok := mappingUpdateFunctionSourceDetailsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
