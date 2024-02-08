// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeClusterCommandOutputData An object representing execution output of a command.
type ComputeClusterCommandOutputData interface {
}

type computeclustercommandoutputdata struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *computeclustercommandoutputdata) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercomputeclustercommandoutputdata computeclustercommandoutputdata
	s := struct {
		Model Unmarshalercomputeclustercommandoutputdata
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *computeclustercommandoutputdata) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "TEXT_HTML":
		mm := TextHtmlComputeClusterCommandOutputData{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT_PLAIN":
		mm := TextPlainComputeClusterCommandOutputData{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IMAGE_PNG":
		mm := ImagePngComputeClusterCommandOutputData{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ComputeClusterCommandOutputData: %s.", m.Type)
		return *m, nil
	}
}

func (m computeclustercommandoutputdata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m computeclustercommandoutputdata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputeClusterCommandOutputDataTypeEnum Enum with underlying type: string
type ComputeClusterCommandOutputDataTypeEnum string

// Set of constants representing the allowable values for ComputeClusterCommandOutputDataTypeEnum
const (
	ComputeClusterCommandOutputDataTypeTextPlain ComputeClusterCommandOutputDataTypeEnum = "TEXT_PLAIN"
	ComputeClusterCommandOutputDataTypeTextHtml  ComputeClusterCommandOutputDataTypeEnum = "TEXT_HTML"
	ComputeClusterCommandOutputDataTypeImagePng  ComputeClusterCommandOutputDataTypeEnum = "IMAGE_PNG"
)

var mappingComputeClusterCommandOutputDataTypeEnum = map[string]ComputeClusterCommandOutputDataTypeEnum{
	"TEXT_PLAIN": ComputeClusterCommandOutputDataTypeTextPlain,
	"TEXT_HTML":  ComputeClusterCommandOutputDataTypeTextHtml,
	"IMAGE_PNG":  ComputeClusterCommandOutputDataTypeImagePng,
}

var mappingComputeClusterCommandOutputDataTypeEnumLowerCase = map[string]ComputeClusterCommandOutputDataTypeEnum{
	"text_plain": ComputeClusterCommandOutputDataTypeTextPlain,
	"text_html":  ComputeClusterCommandOutputDataTypeTextHtml,
	"image_png":  ComputeClusterCommandOutputDataTypeImagePng,
}

// GetComputeClusterCommandOutputDataTypeEnumValues Enumerates the set of values for ComputeClusterCommandOutputDataTypeEnum
func GetComputeClusterCommandOutputDataTypeEnumValues() []ComputeClusterCommandOutputDataTypeEnum {
	values := make([]ComputeClusterCommandOutputDataTypeEnum, 0)
	for _, v := range mappingComputeClusterCommandOutputDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeClusterCommandOutputDataTypeEnumStringValues Enumerates the set of values in String for ComputeClusterCommandOutputDataTypeEnum
func GetComputeClusterCommandOutputDataTypeEnumStringValues() []string {
	return []string{
		"TEXT_PLAIN",
		"TEXT_HTML",
		"IMAGE_PNG",
	}
}

// GetMappingComputeClusterCommandOutputDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeClusterCommandOutputDataTypeEnum(val string) (ComputeClusterCommandOutputDataTypeEnum, bool) {
	enum, ok := mappingComputeClusterCommandOutputDataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
