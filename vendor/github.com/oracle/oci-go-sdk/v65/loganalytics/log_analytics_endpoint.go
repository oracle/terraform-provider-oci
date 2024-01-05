// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogAnalyticsEndpoint Endpoint configuration for REST API based log collection.
type LogAnalyticsEndpoint interface {
}

type loganalyticsendpoint struct {
	JsonData     []byte
	EndpointType string `json:"endpointType"`
}

// UnmarshalJSON unmarshals json
func (m *loganalyticsendpoint) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerloganalyticsendpoint loganalyticsendpoint
	s := struct {
		Model Unmarshalerloganalyticsendpoint
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.EndpointType = s.Model.EndpointType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *loganalyticsendpoint) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EndpointType {
	case "LOG_LIST":
		mm := LogListTypeEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOG":
		mm := LogTypeEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for LogAnalyticsEndpoint: %s.", m.EndpointType)
		return *m, nil
	}
}

func (m loganalyticsendpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m loganalyticsendpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LogAnalyticsEndpointEndpointTypeEnum Enum with underlying type: string
type LogAnalyticsEndpointEndpointTypeEnum string

// Set of constants representing the allowable values for LogAnalyticsEndpointEndpointTypeEnum
const (
	LogAnalyticsEndpointEndpointTypeLogList LogAnalyticsEndpointEndpointTypeEnum = "LOG_LIST"
	LogAnalyticsEndpointEndpointTypeLog     LogAnalyticsEndpointEndpointTypeEnum = "LOG"
)

var mappingLogAnalyticsEndpointEndpointTypeEnum = map[string]LogAnalyticsEndpointEndpointTypeEnum{
	"LOG_LIST": LogAnalyticsEndpointEndpointTypeLogList,
	"LOG":      LogAnalyticsEndpointEndpointTypeLog,
}

var mappingLogAnalyticsEndpointEndpointTypeEnumLowerCase = map[string]LogAnalyticsEndpointEndpointTypeEnum{
	"log_list": LogAnalyticsEndpointEndpointTypeLogList,
	"log":      LogAnalyticsEndpointEndpointTypeLog,
}

// GetLogAnalyticsEndpointEndpointTypeEnumValues Enumerates the set of values for LogAnalyticsEndpointEndpointTypeEnum
func GetLogAnalyticsEndpointEndpointTypeEnumValues() []LogAnalyticsEndpointEndpointTypeEnum {
	values := make([]LogAnalyticsEndpointEndpointTypeEnum, 0)
	for _, v := range mappingLogAnalyticsEndpointEndpointTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsEndpointEndpointTypeEnumStringValues Enumerates the set of values in String for LogAnalyticsEndpointEndpointTypeEnum
func GetLogAnalyticsEndpointEndpointTypeEnumStringValues() []string {
	return []string{
		"LOG_LIST",
		"LOG",
	}
}

// GetMappingLogAnalyticsEndpointEndpointTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsEndpointEndpointTypeEnum(val string) (LogAnalyticsEndpointEndpointTypeEnum, bool) {
	enum, ok := mappingLogAnalyticsEndpointEndpointTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
