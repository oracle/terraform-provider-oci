// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BaseOutput Base output.
type BaseOutput interface {

	// Output label shown to the user.
	GetTitle() *string

	// Extended help or summary for understanding output.
	GetDescription() *string

	// If true, marks this output as sensitive.
	GetIsSensitive() *bool

	// Hint about formatting or rendering the output value.
	GetFormat() *string

	// Expression to show/hide this output.
	GetVisible() *string
}

type baseoutput struct {
	JsonData    []byte
	Title       *string `mandatory:"false" json:"title"`
	Description *string `mandatory:"false" json:"description"`
	IsSensitive *bool   `mandatory:"false" json:"isSensitive"`
	Format      *string `mandatory:"false" json:"format"`
	Visible     *string `mandatory:"false" json:"visible"`
	Type        string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *baseoutput) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbaseoutput baseoutput
	s := struct {
		Model Unmarshalerbaseoutput
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Title = s.Model.Title
	m.Description = s.Model.Description
	m.IsSensitive = s.Model.IsSensitive
	m.Format = s.Model.Format
	m.Visible = s.Model.Visible
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *baseoutput) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "LIST":
		mm := ListOutput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCID":
		mm := OcidOutput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CSV":
		mm := CsvOutput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COPYABLESTRING":
		mm := CopyableStringOutput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JSON":
		mm := JsonOutput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NUMBER":
		mm := NumberOutput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STRING":
		mm := StringOutput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MAP":
		mm := MapOutput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LINK":
		mm := LinkOutput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "boolean":
		mm := BoolOutput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for BaseOutput: %s.", m.Type)
		return *m, nil
	}
}

// GetTitle returns Title
func (m baseoutput) GetTitle() *string {
	return m.Title
}

// GetDescription returns Description
func (m baseoutput) GetDescription() *string {
	return m.Description
}

// GetIsSensitive returns IsSensitive
func (m baseoutput) GetIsSensitive() *bool {
	return m.IsSensitive
}

// GetFormat returns Format
func (m baseoutput) GetFormat() *string {
	return m.Format
}

// GetVisible returns Visible
func (m baseoutput) GetVisible() *string {
	return m.Visible
}

func (m baseoutput) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m baseoutput) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseOutputTypeEnum Enum with underlying type: string
type BaseOutputTypeEnum string

// Set of constants representing the allowable values for BaseOutputTypeEnum
const (
	BaseOutputTypeArray          BaseOutputTypeEnum = "ARRAY"
	BaseOutputTypeBoolean        BaseOutputTypeEnum = "BOOLEAN"
	BaseOutputTypeNumber         BaseOutputTypeEnum = "NUMBER"
	BaseOutputTypeString         BaseOutputTypeEnum = "STRING"
	BaseOutputTypeCopyablestring BaseOutputTypeEnum = "COPYABLESTRING"
	BaseOutputTypeMap            BaseOutputTypeEnum = "MAP"
	BaseOutputTypeList           BaseOutputTypeEnum = "LIST"
	BaseOutputTypeJson           BaseOutputTypeEnum = "JSON"
	BaseOutputTypeCsv            BaseOutputTypeEnum = "CSV"
	BaseOutputTypeLink           BaseOutputTypeEnum = "LINK"
	BaseOutputTypeOcid           BaseOutputTypeEnum = "OCID"
)

var mappingBaseOutputTypeEnum = map[string]BaseOutputTypeEnum{
	"ARRAY":          BaseOutputTypeArray,
	"BOOLEAN":        BaseOutputTypeBoolean,
	"NUMBER":         BaseOutputTypeNumber,
	"STRING":         BaseOutputTypeString,
	"COPYABLESTRING": BaseOutputTypeCopyablestring,
	"MAP":            BaseOutputTypeMap,
	"LIST":           BaseOutputTypeList,
	"JSON":           BaseOutputTypeJson,
	"CSV":            BaseOutputTypeCsv,
	"LINK":           BaseOutputTypeLink,
	"OCID":           BaseOutputTypeOcid,
}

var mappingBaseOutputTypeEnumLowerCase = map[string]BaseOutputTypeEnum{
	"array":          BaseOutputTypeArray,
	"boolean":        BaseOutputTypeBoolean,
	"number":         BaseOutputTypeNumber,
	"string":         BaseOutputTypeString,
	"copyablestring": BaseOutputTypeCopyablestring,
	"map":            BaseOutputTypeMap,
	"list":           BaseOutputTypeList,
	"json":           BaseOutputTypeJson,
	"csv":            BaseOutputTypeCsv,
	"link":           BaseOutputTypeLink,
	"ocid":           BaseOutputTypeOcid,
}

// GetBaseOutputTypeEnumValues Enumerates the set of values for BaseOutputTypeEnum
func GetBaseOutputTypeEnumValues() []BaseOutputTypeEnum {
	values := make([]BaseOutputTypeEnum, 0)
	for _, v := range mappingBaseOutputTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseOutputTypeEnumStringValues Enumerates the set of values in String for BaseOutputTypeEnum
func GetBaseOutputTypeEnumStringValues() []string {
	return []string{
		"ARRAY",
		"BOOLEAN",
		"NUMBER",
		"STRING",
		"COPYABLESTRING",
		"MAP",
		"LIST",
		"JSON",
		"CSV",
		"LINK",
		"OCID",
	}
}

// GetMappingBaseOutputTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseOutputTypeEnum(val string) (BaseOutputTypeEnum, bool) {
	enum, ok := mappingBaseOutputTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
