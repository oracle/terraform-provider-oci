// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ValueExtractionDetails Value extraction details
type ValueExtractionDetails interface {
}

type valueextractiondetails struct {
	JsonData []byte
	Source   string `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *valueextractiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalervalueextractiondetails valueextractiondetails
	s := struct {
		Model Unmarshalervalueextractiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *valueextractiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "HEADER":
		mm := ValueExtractionFromResponseHeader{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BODY":
		mm := ValueExtractionFromResponseBody{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ValueExtractionDetails: %s.", m.Source)
		return *m, nil
	}
}

func (m valueextractiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m valueextractiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ValueExtractionDetailsSourceEnum Enum with underlying type: string
type ValueExtractionDetailsSourceEnum string

// Set of constants representing the allowable values for ValueExtractionDetailsSourceEnum
const (
	ValueExtractionDetailsSourceHeader ValueExtractionDetailsSourceEnum = "HEADER"
	ValueExtractionDetailsSourceBody   ValueExtractionDetailsSourceEnum = "BODY"
)

var mappingValueExtractionDetailsSourceEnum = map[string]ValueExtractionDetailsSourceEnum{
	"HEADER": ValueExtractionDetailsSourceHeader,
	"BODY":   ValueExtractionDetailsSourceBody,
}

var mappingValueExtractionDetailsSourceEnumLowerCase = map[string]ValueExtractionDetailsSourceEnum{
	"header": ValueExtractionDetailsSourceHeader,
	"body":   ValueExtractionDetailsSourceBody,
}

// GetValueExtractionDetailsSourceEnumValues Enumerates the set of values for ValueExtractionDetailsSourceEnum
func GetValueExtractionDetailsSourceEnumValues() []ValueExtractionDetailsSourceEnum {
	values := make([]ValueExtractionDetailsSourceEnum, 0)
	for _, v := range mappingValueExtractionDetailsSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetValueExtractionDetailsSourceEnumStringValues Enumerates the set of values in String for ValueExtractionDetailsSourceEnum
func GetValueExtractionDetailsSourceEnumStringValues() []string {
	return []string{
		"HEADER",
		"BODY",
	}
}

// GetMappingValueExtractionDetailsSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingValueExtractionDetailsSourceEnum(val string) (ValueExtractionDetailsSourceEnum, bool) {
	enum, ok := mappingValueExtractionDetailsSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
