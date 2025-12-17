// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LoggingConfiguration Logging configuration for batch context.
type LoggingConfiguration interface {
}

type loggingconfiguration struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *loggingconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerloggingconfiguration loggingconfiguration
	s := struct {
		Model Unmarshalerloggingconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *loggingconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OCI_LOGGING":
		mm := OciLoggingConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for LoggingConfiguration: %s.", m.Type)
		return *m, nil
	}
}

func (m loggingconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m loggingconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LoggingConfigurationTypeEnum Enum with underlying type: string
type LoggingConfigurationTypeEnum string

// Set of constants representing the allowable values for LoggingConfigurationTypeEnum
const (
	LoggingConfigurationTypeOciLogging LoggingConfigurationTypeEnum = "OCI_LOGGING"
)

var mappingLoggingConfigurationTypeEnum = map[string]LoggingConfigurationTypeEnum{
	"OCI_LOGGING": LoggingConfigurationTypeOciLogging,
}

var mappingLoggingConfigurationTypeEnumLowerCase = map[string]LoggingConfigurationTypeEnum{
	"oci_logging": LoggingConfigurationTypeOciLogging,
}

// GetLoggingConfigurationTypeEnumValues Enumerates the set of values for LoggingConfigurationTypeEnum
func GetLoggingConfigurationTypeEnumValues() []LoggingConfigurationTypeEnum {
	values := make([]LoggingConfigurationTypeEnum, 0)
	for _, v := range mappingLoggingConfigurationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLoggingConfigurationTypeEnumStringValues Enumerates the set of values in String for LoggingConfigurationTypeEnum
func GetLoggingConfigurationTypeEnumStringValues() []string {
	return []string{
		"OCI_LOGGING",
	}
}

// GetMappingLoggingConfigurationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoggingConfigurationTypeEnum(val string) (LoggingConfigurationTypeEnum, bool) {
	enum, ok := mappingLoggingConfigurationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
