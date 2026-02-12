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

// ValueExtractionFromResponseBody Value extraction details from response body.
type ValueExtractionFromResponseBody struct {

	// Expression to get value from response body.
	Expression *string `mandatory:"true" json:"expression"`

	// extraction method
	ExtractionMethod ValueExtractionFromResponseBodyExtractionMethodEnum `mandatory:"true" json:"extractionMethod"`
}

func (m ValueExtractionFromResponseBody) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValueExtractionFromResponseBody) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingValueExtractionFromResponseBodyExtractionMethodEnum(string(m.ExtractionMethod)); !ok && m.ExtractionMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExtractionMethod: %s. Supported values are: %s.", m.ExtractionMethod, strings.Join(GetValueExtractionFromResponseBodyExtractionMethodEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ValueExtractionFromResponseBody) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeValueExtractionFromResponseBody ValueExtractionFromResponseBody
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeValueExtractionFromResponseBody
	}{
		"BODY",
		(MarshalTypeValueExtractionFromResponseBody)(m),
	}

	return json.Marshal(&s)
}

// ValueExtractionFromResponseBodyExtractionMethodEnum Enum with underlying type: string
type ValueExtractionFromResponseBodyExtractionMethodEnum string

// Set of constants representing the allowable values for ValueExtractionFromResponseBodyExtractionMethodEnum
const (
	ValueExtractionFromResponseBodyExtractionMethodJsonPath ValueExtractionFromResponseBodyExtractionMethodEnum = "JSON_PATH"
	ValueExtractionFromResponseBodyExtractionMethodRegex    ValueExtractionFromResponseBodyExtractionMethodEnum = "REGEX"
)

var mappingValueExtractionFromResponseBodyExtractionMethodEnum = map[string]ValueExtractionFromResponseBodyExtractionMethodEnum{
	"JSON_PATH": ValueExtractionFromResponseBodyExtractionMethodJsonPath,
	"REGEX":     ValueExtractionFromResponseBodyExtractionMethodRegex,
}

var mappingValueExtractionFromResponseBodyExtractionMethodEnumLowerCase = map[string]ValueExtractionFromResponseBodyExtractionMethodEnum{
	"json_path": ValueExtractionFromResponseBodyExtractionMethodJsonPath,
	"regex":     ValueExtractionFromResponseBodyExtractionMethodRegex,
}

// GetValueExtractionFromResponseBodyExtractionMethodEnumValues Enumerates the set of values for ValueExtractionFromResponseBodyExtractionMethodEnum
func GetValueExtractionFromResponseBodyExtractionMethodEnumValues() []ValueExtractionFromResponseBodyExtractionMethodEnum {
	values := make([]ValueExtractionFromResponseBodyExtractionMethodEnum, 0)
	for _, v := range mappingValueExtractionFromResponseBodyExtractionMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetValueExtractionFromResponseBodyExtractionMethodEnumStringValues Enumerates the set of values in String for ValueExtractionFromResponseBodyExtractionMethodEnum
func GetValueExtractionFromResponseBodyExtractionMethodEnumStringValues() []string {
	return []string{
		"JSON_PATH",
		"REGEX",
	}
}

// GetMappingValueExtractionFromResponseBodyExtractionMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingValueExtractionFromResponseBodyExtractionMethodEnum(val string) (ValueExtractionFromResponseBodyExtractionMethodEnum, bool) {
	enum, ok := mappingValueExtractionFromResponseBodyExtractionMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
