// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Internet of Things API
//
// Use the Internet of Things (IoT) API to manage IoT domain groups, domains, and digital twin resources including models, adapters, instances, and relationships.
// For more information, see Internet of Things (https://docs.oracle.com/iaas/Content/internet-of-things/home.htm).
//

package iot

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DigitalTwinAdapterPayload Reference payload structure template received from IoT device. This payload
// must specify its content type using the `dataFormat` property.
type DigitalTwinAdapterPayload interface {
}

type digitaltwinadapterpayload struct {
	JsonData   []byte
	DataFormat string `json:"dataFormat"`
}

// UnmarshalJSON unmarshals json
func (m *digitaltwinadapterpayload) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdigitaltwinadapterpayload digitaltwinadapterpayload
	s := struct {
		Model Unmarshalerdigitaltwinadapterpayload
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DataFormat = s.Model.DataFormat

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *digitaltwinadapterpayload) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DataFormat {
	case "JSON":
		mm := DigitalTwinAdapterJsonPayload{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DigitalTwinAdapterPayload: %s.", m.DataFormat)
		return *m, nil
	}
}

func (m digitaltwinadapterpayload) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m digitaltwinadapterpayload) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DigitalTwinAdapterPayloadDataFormatEnum Enum with underlying type: string
type DigitalTwinAdapterPayloadDataFormatEnum string

// Set of constants representing the allowable values for DigitalTwinAdapterPayloadDataFormatEnum
const (
	DigitalTwinAdapterPayloadDataFormatJson DigitalTwinAdapterPayloadDataFormatEnum = "JSON"
)

var mappingDigitalTwinAdapterPayloadDataFormatEnum = map[string]DigitalTwinAdapterPayloadDataFormatEnum{
	"JSON": DigitalTwinAdapterPayloadDataFormatJson,
}

var mappingDigitalTwinAdapterPayloadDataFormatEnumLowerCase = map[string]DigitalTwinAdapterPayloadDataFormatEnum{
	"json": DigitalTwinAdapterPayloadDataFormatJson,
}

// GetDigitalTwinAdapterPayloadDataFormatEnumValues Enumerates the set of values for DigitalTwinAdapterPayloadDataFormatEnum
func GetDigitalTwinAdapterPayloadDataFormatEnumValues() []DigitalTwinAdapterPayloadDataFormatEnum {
	values := make([]DigitalTwinAdapterPayloadDataFormatEnum, 0)
	for _, v := range mappingDigitalTwinAdapterPayloadDataFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetDigitalTwinAdapterPayloadDataFormatEnumStringValues Enumerates the set of values in String for DigitalTwinAdapterPayloadDataFormatEnum
func GetDigitalTwinAdapterPayloadDataFormatEnumStringValues() []string {
	return []string{
		"JSON",
	}
}

// GetMappingDigitalTwinAdapterPayloadDataFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDigitalTwinAdapterPayloadDataFormatEnum(val string) (DigitalTwinAdapterPayloadDataFormatEnum, bool) {
	enum, ok := mappingDigitalTwinAdapterPayloadDataFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
