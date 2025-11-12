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

// Filter An array of Managed Entity objects
type Filter interface {

	// Name of report.
	GetFieldName() *string
}

type filter struct {
	JsonData   []byte
	FieldName  *string `mandatory:"true" json:"fieldName"`
	FilterType string  `json:"filterType"`
}

// UnmarshalJSON unmarshals json
func (m *filter) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfilter filter
	s := struct {
		Model Unmarshalerfilter
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FieldName = s.Model.FieldName
	m.FilterType = s.Model.FilterType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *filter) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.FilterType {
	case "MULTI_VALUE":
		mm := MultiValueFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SINGLE_VALUE":
		mm := SingleValueFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Filter: %s.", m.FilterType)
		return *m, nil
	}
}

// GetFieldName returns FieldName
func (m filter) GetFieldName() *string {
	return m.FieldName
}

func (m filter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m filter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FilterFilterTypeEnum Enum with underlying type: string
type FilterFilterTypeEnum string

// Set of constants representing the allowable values for FilterFilterTypeEnum
const (
	FilterFilterTypeSingleValue FilterFilterTypeEnum = "SINGLE_VALUE"
	FilterFilterTypeMultiValue  FilterFilterTypeEnum = "MULTI_VALUE"
)

var mappingFilterFilterTypeEnum = map[string]FilterFilterTypeEnum{
	"SINGLE_VALUE": FilterFilterTypeSingleValue,
	"MULTI_VALUE":  FilterFilterTypeMultiValue,
}

var mappingFilterFilterTypeEnumLowerCase = map[string]FilterFilterTypeEnum{
	"single_value": FilterFilterTypeSingleValue,
	"multi_value":  FilterFilterTypeMultiValue,
}

// GetFilterFilterTypeEnumValues Enumerates the set of values for FilterFilterTypeEnum
func GetFilterFilterTypeEnumValues() []FilterFilterTypeEnum {
	values := make([]FilterFilterTypeEnum, 0)
	for _, v := range mappingFilterFilterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFilterFilterTypeEnumStringValues Enumerates the set of values in String for FilterFilterTypeEnum
func GetFilterFilterTypeEnumStringValues() []string {
	return []string{
		"SINGLE_VALUE",
		"MULTI_VALUE",
	}
}

// GetMappingFilterFilterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFilterFilterTypeEnum(val string) (FilterFilterTypeEnum, bool) {
	enum, ok := mappingFilterFilterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
