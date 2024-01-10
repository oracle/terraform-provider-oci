// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Dimensions The width and height of a page.
type Dimensions struct {

	// the width of a page.
	Width *float64 `mandatory:"true" json:"width"`

	// The height of a page.
	Height *float64 `mandatory:"true" json:"height"`

	// The unit of length.
	Unit DimensionsUnitEnum `mandatory:"true" json:"unit"`
}

func (m Dimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Dimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDimensionsUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetDimensionsUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DimensionsUnitEnum Enum with underlying type: string
type DimensionsUnitEnum string

// Set of constants representing the allowable values for DimensionsUnitEnum
const (
	DimensionsUnitPixel DimensionsUnitEnum = "PIXEL"
	DimensionsUnitInch  DimensionsUnitEnum = "INCH"
)

var mappingDimensionsUnitEnum = map[string]DimensionsUnitEnum{
	"PIXEL": DimensionsUnitPixel,
	"INCH":  DimensionsUnitInch,
}

var mappingDimensionsUnitEnumLowerCase = map[string]DimensionsUnitEnum{
	"pixel": DimensionsUnitPixel,
	"inch":  DimensionsUnitInch,
}

// GetDimensionsUnitEnumValues Enumerates the set of values for DimensionsUnitEnum
func GetDimensionsUnitEnumValues() []DimensionsUnitEnum {
	values := make([]DimensionsUnitEnum, 0)
	for _, v := range mappingDimensionsUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetDimensionsUnitEnumStringValues Enumerates the set of values in String for DimensionsUnitEnum
func GetDimensionsUnitEnumStringValues() []string {
	return []string{
		"PIXEL",
		"INCH",
	}
}

// GetMappingDimensionsUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDimensionsUnitEnum(val string) (DimensionsUnitEnum, bool) {
	enum, ok := mappingDimensionsUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
