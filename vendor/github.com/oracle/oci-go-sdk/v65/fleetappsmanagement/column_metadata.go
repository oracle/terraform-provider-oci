// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ColumnMetadata Column metadata.
type ColumnMetadata struct {

	// Column name.
	Name *string `mandatory:"true" json:"name"`

	// Column description
	Description *string `mandatory:"true" json:"description"`

	// Column value type.
	Type ColumnMetadataTypeEnum `mandatory:"true" json:"type"`
}

func (m ColumnMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ColumnMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingColumnMetadataTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetColumnMetadataTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ColumnMetadataTypeEnum Enum with underlying type: string
type ColumnMetadataTypeEnum string

// Set of constants representing the allowable values for ColumnMetadataTypeEnum
const (
	ColumnMetadataTypeString   ColumnMetadataTypeEnum = "STRING"
	ColumnMetadataTypeNumber   ColumnMetadataTypeEnum = "NUMBER"
	ColumnMetadataTypeDateTime ColumnMetadataTypeEnum = "DATE_TIME"
)

var mappingColumnMetadataTypeEnum = map[string]ColumnMetadataTypeEnum{
	"STRING":    ColumnMetadataTypeString,
	"NUMBER":    ColumnMetadataTypeNumber,
	"DATE_TIME": ColumnMetadataTypeDateTime,
}

var mappingColumnMetadataTypeEnumLowerCase = map[string]ColumnMetadataTypeEnum{
	"string":    ColumnMetadataTypeString,
	"number":    ColumnMetadataTypeNumber,
	"date_time": ColumnMetadataTypeDateTime,
}

// GetColumnMetadataTypeEnumValues Enumerates the set of values for ColumnMetadataTypeEnum
func GetColumnMetadataTypeEnumValues() []ColumnMetadataTypeEnum {
	values := make([]ColumnMetadataTypeEnum, 0)
	for _, v := range mappingColumnMetadataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetColumnMetadataTypeEnumStringValues Enumerates the set of values in String for ColumnMetadataTypeEnum
func GetColumnMetadataTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"NUMBER",
		"DATE_TIME",
	}
}

// GetMappingColumnMetadataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingColumnMetadataTypeEnum(val string) (ColumnMetadataTypeEnum, bool) {
	enum, ok := mappingColumnMetadataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
