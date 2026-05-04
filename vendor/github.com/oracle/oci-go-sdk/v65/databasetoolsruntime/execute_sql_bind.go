// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecuteSqlBind Bind information.
type ExecuteSqlBind struct {

	// The data type.
	DataType *string `mandatory:"true" json:"dataType"`

	// Name of the bind.
	Name *string `mandatory:"false" json:"name"`

	// Index of the bind. Index of the first bind is 1.
	Index *int `mandatory:"false" json:"index"`

	// The mode in which the bind is used.
	Mode ExecuteSqlBindModeEnum `mandatory:"false" json:"mode,omitempty"`

	// values
	Values []interface{} `mandatory:"false" json:"values"`

	// results
	Results []interface{} `mandatory:"false" json:"results"`

	PlsqlTableTypeDetails *ExecuteSqlBindPlsqlTable `mandatory:"false" json:"plsqlTableTypeDetails"`
}

func (m ExecuteSqlBind) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlBind) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExecuteSqlBindModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetExecuteSqlBindModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExecuteSqlBindModeEnum Enum with underlying type: string
type ExecuteSqlBindModeEnum string

// Set of constants representing the allowable values for ExecuteSqlBindModeEnum
const (
	ExecuteSqlBindModeIn    ExecuteSqlBindModeEnum = "IN"
	ExecuteSqlBindModeInout ExecuteSqlBindModeEnum = "INOUT"
	ExecuteSqlBindModeOut   ExecuteSqlBindModeEnum = "OUT"
)

var mappingExecuteSqlBindModeEnum = map[string]ExecuteSqlBindModeEnum{
	"IN":    ExecuteSqlBindModeIn,
	"INOUT": ExecuteSqlBindModeInout,
	"OUT":   ExecuteSqlBindModeOut,
}

var mappingExecuteSqlBindModeEnumLowerCase = map[string]ExecuteSqlBindModeEnum{
	"in":    ExecuteSqlBindModeIn,
	"inout": ExecuteSqlBindModeInout,
	"out":   ExecuteSqlBindModeOut,
}

// GetExecuteSqlBindModeEnumValues Enumerates the set of values for ExecuteSqlBindModeEnum
func GetExecuteSqlBindModeEnumValues() []ExecuteSqlBindModeEnum {
	values := make([]ExecuteSqlBindModeEnum, 0)
	for _, v := range mappingExecuteSqlBindModeEnum {
		values = append(values, v)
	}
	return values
}

// GetExecuteSqlBindModeEnumStringValues Enumerates the set of values in String for ExecuteSqlBindModeEnum
func GetExecuteSqlBindModeEnumStringValues() []string {
	return []string{
		"IN",
		"INOUT",
		"OUT",
	}
}

// GetMappingExecuteSqlBindModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecuteSqlBindModeEnum(val string) (ExecuteSqlBindModeEnum, bool) {
	enum, ok := mappingExecuteSqlBindModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
