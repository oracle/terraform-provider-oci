// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DeriveEntitiesItem The request object for single derived Entity.
type DeriveEntitiesItem struct {

	// The model type of DeriveEntitiesRequestItem
	ModelType *string `mandatory:"true" json:"modelType"`

	// Determines whether derived entity is treated as source or target
	Mode DeriveEntitiesItemModeEnum `mandatory:"true" json:"mode"`

	ReferencedDataObject ReferencedDataObject `mandatory:"true" json:"referencedDataObject"`
}

func (m DeriveEntitiesItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeriveEntitiesItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeriveEntitiesItemModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetDeriveEntitiesItemModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DeriveEntitiesItem) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ModelType            *string                    `json:"modelType"`
		Mode                 DeriveEntitiesItemModeEnum `json:"mode"`
		ReferencedDataObject referenceddataobject       `json:"referencedDataObject"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ModelType = model.ModelType

	m.Mode = model.Mode

	nn, e = model.ReferencedDataObject.UnmarshalPolymorphicJSON(model.ReferencedDataObject.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ReferencedDataObject = nn.(ReferencedDataObject)
	} else {
		m.ReferencedDataObject = nil
	}

	return
}

// DeriveEntitiesItemModeEnum Enum with underlying type: string
type DeriveEntitiesItemModeEnum string

// Set of constants representing the allowable values for DeriveEntitiesItemModeEnum
const (
	DeriveEntitiesItemModeIn  DeriveEntitiesItemModeEnum = "IN"
	DeriveEntitiesItemModeOut DeriveEntitiesItemModeEnum = "OUT"
)

var mappingDeriveEntitiesItemModeEnum = map[string]DeriveEntitiesItemModeEnum{
	"IN":  DeriveEntitiesItemModeIn,
	"OUT": DeriveEntitiesItemModeOut,
}

var mappingDeriveEntitiesItemModeEnumLowerCase = map[string]DeriveEntitiesItemModeEnum{
	"in":  DeriveEntitiesItemModeIn,
	"out": DeriveEntitiesItemModeOut,
}

// GetDeriveEntitiesItemModeEnumValues Enumerates the set of values for DeriveEntitiesItemModeEnum
func GetDeriveEntitiesItemModeEnumValues() []DeriveEntitiesItemModeEnum {
	values := make([]DeriveEntitiesItemModeEnum, 0)
	for _, v := range mappingDeriveEntitiesItemModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeriveEntitiesItemModeEnumStringValues Enumerates the set of values in String for DeriveEntitiesItemModeEnum
func GetDeriveEntitiesItemModeEnumStringValues() []string {
	return []string{
		"IN",
		"OUT",
	}
}

// GetMappingDeriveEntitiesItemModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeriveEntitiesItemModeEnum(val string) (DeriveEntitiesItemModeEnum, bool) {
	enum, ok := mappingDeriveEntitiesItemModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
