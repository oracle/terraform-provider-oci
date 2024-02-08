// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FieldValue The value of a form field.
type FieldValue interface {

	// The confidence score between 0 and 1.
	GetConfidence() *float32

	GetBoundingPolygon() *BoundingPolygon

	// The indexes of the words in the field value.
	GetWordIndexes() []int

	// The detected text of a field.
	GetText() *string
}

type fieldvalue struct {
	JsonData        []byte
	Text            *string          `mandatory:"false" json:"text"`
	Confidence      *float32         `mandatory:"true" json:"confidence"`
	BoundingPolygon *BoundingPolygon `mandatory:"true" json:"boundingPolygon"`
	WordIndexes     []int            `mandatory:"true" json:"wordIndexes"`
	ValueType       string           `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *fieldvalue) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfieldvalue fieldvalue
	s := struct {
		Model Unmarshalerfieldvalue
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Confidence = s.Model.Confidence
	m.BoundingPolygon = s.Model.BoundingPolygon
	m.WordIndexes = s.Model.WordIndexes
	m.Text = s.Model.Text
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *fieldvalue) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "TIME":
		mm := ValueTime{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTEGER":
		mm := ValueInteger{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATE":
		mm := ValueDate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NUMBER":
		mm := ValueNumber{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STRING":
		mm := ValueString{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PHONE_NUMBER":
		mm := ValuePhoneNumber{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ARRAY":
		mm := ValueArray{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for FieldValue: %s.", m.ValueType)
		return *m, nil
	}
}

// GetText returns Text
func (m fieldvalue) GetText() *string {
	return m.Text
}

// GetConfidence returns Confidence
func (m fieldvalue) GetConfidence() *float32 {
	return m.Confidence
}

// GetBoundingPolygon returns BoundingPolygon
func (m fieldvalue) GetBoundingPolygon() *BoundingPolygon {
	return m.BoundingPolygon
}

// GetWordIndexes returns WordIndexes
func (m fieldvalue) GetWordIndexes() []int {
	return m.WordIndexes
}

func (m fieldvalue) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m fieldvalue) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FieldValueValueTypeEnum Enum with underlying type: string
type FieldValueValueTypeEnum string

// Set of constants representing the allowable values for FieldValueValueTypeEnum
const (
	FieldValueValueTypeString      FieldValueValueTypeEnum = "STRING"
	FieldValueValueTypeDate        FieldValueValueTypeEnum = "DATE"
	FieldValueValueTypeTime        FieldValueValueTypeEnum = "TIME"
	FieldValueValueTypePhoneNumber FieldValueValueTypeEnum = "PHONE_NUMBER"
	FieldValueValueTypeNumber      FieldValueValueTypeEnum = "NUMBER"
	FieldValueValueTypeInteger     FieldValueValueTypeEnum = "INTEGER"
	FieldValueValueTypeArray       FieldValueValueTypeEnum = "ARRAY"
)

var mappingFieldValueValueTypeEnum = map[string]FieldValueValueTypeEnum{
	"STRING":       FieldValueValueTypeString,
	"DATE":         FieldValueValueTypeDate,
	"TIME":         FieldValueValueTypeTime,
	"PHONE_NUMBER": FieldValueValueTypePhoneNumber,
	"NUMBER":       FieldValueValueTypeNumber,
	"INTEGER":      FieldValueValueTypeInteger,
	"ARRAY":        FieldValueValueTypeArray,
}

var mappingFieldValueValueTypeEnumLowerCase = map[string]FieldValueValueTypeEnum{
	"string":       FieldValueValueTypeString,
	"date":         FieldValueValueTypeDate,
	"time":         FieldValueValueTypeTime,
	"phone_number": FieldValueValueTypePhoneNumber,
	"number":       FieldValueValueTypeNumber,
	"integer":      FieldValueValueTypeInteger,
	"array":        FieldValueValueTypeArray,
}

// GetFieldValueValueTypeEnumValues Enumerates the set of values for FieldValueValueTypeEnum
func GetFieldValueValueTypeEnumValues() []FieldValueValueTypeEnum {
	values := make([]FieldValueValueTypeEnum, 0)
	for _, v := range mappingFieldValueValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFieldValueValueTypeEnumStringValues Enumerates the set of values in String for FieldValueValueTypeEnum
func GetFieldValueValueTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"DATE",
		"TIME",
		"PHONE_NUMBER",
		"NUMBER",
		"INTEGER",
		"ARRAY",
	}
}

// GetMappingFieldValueValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFieldValueValueTypeEnum(val string) (FieldValueValueTypeEnum, bool) {
	enum, ok := mappingFieldValueValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
