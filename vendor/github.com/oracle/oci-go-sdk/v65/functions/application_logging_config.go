// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApplicationLoggingConfig Set logging configuration for an application. This is only used if Service Logs for the application are enabled in the OCI Logging service.
type ApplicationLoggingConfig struct {

	// Specify the format of log lines emitted by functions in this application.
	LineFormat ApplicationLoggingConfigLineFormatEnum `mandatory:"false" json:"lineFormat,omitempty"`
}

func (m ApplicationLoggingConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplicationLoggingConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingApplicationLoggingConfigLineFormatEnum(string(m.LineFormat)); !ok && m.LineFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LineFormat: %s. Supported values are: %s.", m.LineFormat, strings.Join(GetApplicationLoggingConfigLineFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApplicationLoggingConfigLineFormatEnum Enum with underlying type: string
type ApplicationLoggingConfigLineFormatEnum string

// Set of constants representing the allowable values for ApplicationLoggingConfigLineFormatEnum
const (
	ApplicationLoggingConfigLineFormatJson      ApplicationLoggingConfigLineFormatEnum = "JSON"
	ApplicationLoggingConfigLineFormatPlainText ApplicationLoggingConfigLineFormatEnum = "PLAIN_TEXT"
)

var mappingApplicationLoggingConfigLineFormatEnum = map[string]ApplicationLoggingConfigLineFormatEnum{
	"JSON":       ApplicationLoggingConfigLineFormatJson,
	"PLAIN_TEXT": ApplicationLoggingConfigLineFormatPlainText,
}

var mappingApplicationLoggingConfigLineFormatEnumLowerCase = map[string]ApplicationLoggingConfigLineFormatEnum{
	"json":       ApplicationLoggingConfigLineFormatJson,
	"plain_text": ApplicationLoggingConfigLineFormatPlainText,
}

// GetApplicationLoggingConfigLineFormatEnumValues Enumerates the set of values for ApplicationLoggingConfigLineFormatEnum
func GetApplicationLoggingConfigLineFormatEnumValues() []ApplicationLoggingConfigLineFormatEnum {
	values := make([]ApplicationLoggingConfigLineFormatEnum, 0)
	for _, v := range mappingApplicationLoggingConfigLineFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetApplicationLoggingConfigLineFormatEnumStringValues Enumerates the set of values in String for ApplicationLoggingConfigLineFormatEnum
func GetApplicationLoggingConfigLineFormatEnumStringValues() []string {
	return []string{
		"JSON",
		"PLAIN_TEXT",
	}
}

// GetMappingApplicationLoggingConfigLineFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApplicationLoggingConfigLineFormatEnum(val string) (ApplicationLoggingConfigLineFormatEnum, bool) {
	enum, ok := mappingApplicationLoggingConfigLineFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
