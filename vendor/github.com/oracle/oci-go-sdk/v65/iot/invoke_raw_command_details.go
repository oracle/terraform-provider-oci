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

// InvokeRawCommandDetails Definition of unstructured command invocation payload
type InvokeRawCommandDetails interface {

	// Device endpoint where request should be forwarded to.
	GetRequestEndpoint() *string

	// Specified duration by which to send the request by.
	GetRequestDuration() *string

	// Specified duration by which to receive the response by.
	GetResponseDuration() *string

	// Device endpoint from which response is expected to come.
	GetResponseEndpoint() *string
}

type invokerawcommanddetails struct {
	JsonData          []byte
	RequestDuration   *string `mandatory:"false" json:"requestDuration"`
	ResponseDuration  *string `mandatory:"false" json:"responseDuration"`
	ResponseEndpoint  *string `mandatory:"false" json:"responseEndpoint"`
	RequestEndpoint   *string `mandatory:"true" json:"requestEndpoint"`
	RequestDataFormat string  `json:"requestDataFormat"`
}

// UnmarshalJSON unmarshals json
func (m *invokerawcommanddetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinvokerawcommanddetails invokerawcommanddetails
	s := struct {
		Model Unmarshalerinvokerawcommanddetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RequestEndpoint = s.Model.RequestEndpoint
	m.RequestDuration = s.Model.RequestDuration
	m.ResponseDuration = s.Model.ResponseDuration
	m.ResponseEndpoint = s.Model.ResponseEndpoint
	m.RequestDataFormat = s.Model.RequestDataFormat

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *invokerawcommanddetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RequestDataFormat {
	case "BINARY":
		mm := InvokeRawBinaryCommandDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT":
		mm := InvokeRawTextCommandDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JSON":
		mm := InvokeRawJsonCommandDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for InvokeRawCommandDetails: %s.", m.RequestDataFormat)
		return *m, nil
	}
}

// GetRequestDuration returns RequestDuration
func (m invokerawcommanddetails) GetRequestDuration() *string {
	return m.RequestDuration
}

// GetResponseDuration returns ResponseDuration
func (m invokerawcommanddetails) GetResponseDuration() *string {
	return m.ResponseDuration
}

// GetResponseEndpoint returns ResponseEndpoint
func (m invokerawcommanddetails) GetResponseEndpoint() *string {
	return m.ResponseEndpoint
}

// GetRequestEndpoint returns RequestEndpoint
func (m invokerawcommanddetails) GetRequestEndpoint() *string {
	return m.RequestEndpoint
}

func (m invokerawcommanddetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m invokerawcommanddetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InvokeRawCommandDetailsRequestDataFormatEnum Enum with underlying type: string
type InvokeRawCommandDetailsRequestDataFormatEnum string

// Set of constants representing the allowable values for InvokeRawCommandDetailsRequestDataFormatEnum
const (
	InvokeRawCommandDetailsRequestDataFormatText   InvokeRawCommandDetailsRequestDataFormatEnum = "TEXT"
	InvokeRawCommandDetailsRequestDataFormatJson   InvokeRawCommandDetailsRequestDataFormatEnum = "JSON"
	InvokeRawCommandDetailsRequestDataFormatBinary InvokeRawCommandDetailsRequestDataFormatEnum = "BINARY"
)

var mappingInvokeRawCommandDetailsRequestDataFormatEnum = map[string]InvokeRawCommandDetailsRequestDataFormatEnum{
	"TEXT":   InvokeRawCommandDetailsRequestDataFormatText,
	"JSON":   InvokeRawCommandDetailsRequestDataFormatJson,
	"BINARY": InvokeRawCommandDetailsRequestDataFormatBinary,
}

var mappingInvokeRawCommandDetailsRequestDataFormatEnumLowerCase = map[string]InvokeRawCommandDetailsRequestDataFormatEnum{
	"text":   InvokeRawCommandDetailsRequestDataFormatText,
	"json":   InvokeRawCommandDetailsRequestDataFormatJson,
	"binary": InvokeRawCommandDetailsRequestDataFormatBinary,
}

// GetInvokeRawCommandDetailsRequestDataFormatEnumValues Enumerates the set of values for InvokeRawCommandDetailsRequestDataFormatEnum
func GetInvokeRawCommandDetailsRequestDataFormatEnumValues() []InvokeRawCommandDetailsRequestDataFormatEnum {
	values := make([]InvokeRawCommandDetailsRequestDataFormatEnum, 0)
	for _, v := range mappingInvokeRawCommandDetailsRequestDataFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetInvokeRawCommandDetailsRequestDataFormatEnumStringValues Enumerates the set of values in String for InvokeRawCommandDetailsRequestDataFormatEnum
func GetInvokeRawCommandDetailsRequestDataFormatEnumStringValues() []string {
	return []string{
		"TEXT",
		"JSON",
		"BINARY",
	}
}

// GetMappingInvokeRawCommandDetailsRequestDataFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInvokeRawCommandDetailsRequestDataFormatEnum(val string) (InvokeRawCommandDetailsRequestDataFormatEnum, bool) {
	enum, ok := mappingInvokeRawCommandDetailsRequestDataFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
