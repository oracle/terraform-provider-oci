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

// StatementOutputData An object representing execution output of a statement.
type StatementOutputData interface {
}

type statementoutputdata struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *statementoutputdata) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstatementoutputdata statementoutputdata
	s := struct {
		Model Unmarshalerstatementoutputdata
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *statementoutputdata) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "IMAGE_PNG":
		mm := ImagePngStatementOutputData{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT_HTML":
		mm := TextHtmlStatementOutputData{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT_PLAIN":
		mm := TextPlainStatementOutputData{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for StatementOutputData: %s.", m.Type)
		return *m, nil
	}
}

func (m statementoutputdata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m statementoutputdata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StatementOutputDataTypeEnum Enum with underlying type: string
type StatementOutputDataTypeEnum string

// Set of constants representing the allowable values for StatementOutputDataTypeEnum
const (
	StatementOutputDataTypeTextPlain StatementOutputDataTypeEnum = "TEXT_PLAIN"
	StatementOutputDataTypeTextHtml  StatementOutputDataTypeEnum = "TEXT_HTML"
	StatementOutputDataTypeImagePng  StatementOutputDataTypeEnum = "IMAGE_PNG"
)

var mappingStatementOutputDataTypeEnum = map[string]StatementOutputDataTypeEnum{
	"TEXT_PLAIN": StatementOutputDataTypeTextPlain,
	"TEXT_HTML":  StatementOutputDataTypeTextHtml,
	"IMAGE_PNG":  StatementOutputDataTypeImagePng,
}

var mappingStatementOutputDataTypeEnumLowerCase = map[string]StatementOutputDataTypeEnum{
	"text_plain": StatementOutputDataTypeTextPlain,
	"text_html":  StatementOutputDataTypeTextHtml,
	"image_png":  StatementOutputDataTypeImagePng,
}

// GetStatementOutputDataTypeEnumValues Enumerates the set of values for StatementOutputDataTypeEnum
func GetStatementOutputDataTypeEnumValues() []StatementOutputDataTypeEnum {
	values := make([]StatementOutputDataTypeEnum, 0)
	for _, v := range mappingStatementOutputDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStatementOutputDataTypeEnumStringValues Enumerates the set of values in String for StatementOutputDataTypeEnum
func GetStatementOutputDataTypeEnumStringValues() []string {
	return []string{
		"TEXT_PLAIN",
		"TEXT_HTML",
		"IMAGE_PNG",
	}
}

// GetMappingStatementOutputDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStatementOutputDataTypeEnum(val string) (StatementOutputDataTypeEnum, bool) {
	enum, ok := mappingStatementOutputDataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
