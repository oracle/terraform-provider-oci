// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CryptoSqlnetParameterValue SQLNET parameter value payload.
type CryptoSqlnetParameterValue struct {

	// Value type of the SQLNET parameter.
	Type CryptoSqlnetParameterValueTypeEnum `mandatory:"true" json:"type"`

	// Parsed SQLNET parameter value. TEXT returns a string, BOOLEAN returns a boolean, and LIST returns an array of strings.
	Value *interface{} `mandatory:"true" json:"value"`
}

func (m CryptoSqlnetParameterValue) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoSqlnetParameterValue) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCryptoSqlnetParameterValueTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCryptoSqlnetParameterValueTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CryptoSqlnetParameterValueTypeEnum Enum with underlying type: string
type CryptoSqlnetParameterValueTypeEnum string

// Set of constants representing the allowable values for CryptoSqlnetParameterValueTypeEnum
const (
	CryptoSqlnetParameterValueTypeText    CryptoSqlnetParameterValueTypeEnum = "TEXT"
	CryptoSqlnetParameterValueTypeBoolean CryptoSqlnetParameterValueTypeEnum = "BOOLEAN"
	CryptoSqlnetParameterValueTypeList    CryptoSqlnetParameterValueTypeEnum = "LIST"
)

var mappingCryptoSqlnetParameterValueTypeEnum = map[string]CryptoSqlnetParameterValueTypeEnum{
	"TEXT":    CryptoSqlnetParameterValueTypeText,
	"BOOLEAN": CryptoSqlnetParameterValueTypeBoolean,
	"LIST":    CryptoSqlnetParameterValueTypeList,
}

var mappingCryptoSqlnetParameterValueTypeEnumLowerCase = map[string]CryptoSqlnetParameterValueTypeEnum{
	"text":    CryptoSqlnetParameterValueTypeText,
	"boolean": CryptoSqlnetParameterValueTypeBoolean,
	"list":    CryptoSqlnetParameterValueTypeList,
}

// GetCryptoSqlnetParameterValueTypeEnumValues Enumerates the set of values for CryptoSqlnetParameterValueTypeEnum
func GetCryptoSqlnetParameterValueTypeEnumValues() []CryptoSqlnetParameterValueTypeEnum {
	values := make([]CryptoSqlnetParameterValueTypeEnum, 0)
	for _, v := range mappingCryptoSqlnetParameterValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoSqlnetParameterValueTypeEnumStringValues Enumerates the set of values in String for CryptoSqlnetParameterValueTypeEnum
func GetCryptoSqlnetParameterValueTypeEnumStringValues() []string {
	return []string{
		"TEXT",
		"BOOLEAN",
		"LIST",
	}
}

// GetMappingCryptoSqlnetParameterValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoSqlnetParameterValueTypeEnum(val string) (CryptoSqlnetParameterValueTypeEnum, bool) {
	enum, ok := mappingCryptoSqlnetParameterValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
