// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SelectionDetails Details of the items to be selected and the mode of selection.
type SelectionDetails interface {
}

type selectiondetails struct {
	JsonData      []byte
	SelectionMode string `json:"selectionMode"`
}

// UnmarshalJSON unmarshals json
func (m *selectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerselectiondetails selectiondetails
	s := struct {
		Model Unmarshalerselectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SelectionMode = s.Model.SelectionMode

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *selectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SelectionMode {
	case "LIST":
		mm := ListSelectionMode{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCIM_QUERY":
		mm := ScimQuerySelectionMode{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SelectionDetails: %s.", m.SelectionMode)
		return *m, nil
	}
}

func (m selectiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m selectiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SelectionDetailsSelectionModeEnum Enum with underlying type: string
type SelectionDetailsSelectionModeEnum string

// Set of constants representing the allowable values for SelectionDetailsSelectionModeEnum
const (
	SelectionDetailsSelectionModeList      SelectionDetailsSelectionModeEnum = "LIST"
	SelectionDetailsSelectionModeScimQuery SelectionDetailsSelectionModeEnum = "SCIM_QUERY"
)

var mappingSelectionDetailsSelectionModeEnum = map[string]SelectionDetailsSelectionModeEnum{
	"LIST":       SelectionDetailsSelectionModeList,
	"SCIM_QUERY": SelectionDetailsSelectionModeScimQuery,
}

var mappingSelectionDetailsSelectionModeEnumLowerCase = map[string]SelectionDetailsSelectionModeEnum{
	"list":       SelectionDetailsSelectionModeList,
	"scim_query": SelectionDetailsSelectionModeScimQuery,
}

// GetSelectionDetailsSelectionModeEnumValues Enumerates the set of values for SelectionDetailsSelectionModeEnum
func GetSelectionDetailsSelectionModeEnumValues() []SelectionDetailsSelectionModeEnum {
	values := make([]SelectionDetailsSelectionModeEnum, 0)
	for _, v := range mappingSelectionDetailsSelectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetSelectionDetailsSelectionModeEnumStringValues Enumerates the set of values in String for SelectionDetailsSelectionModeEnum
func GetSelectionDetailsSelectionModeEnumStringValues() []string {
	return []string{
		"LIST",
		"SCIM_QUERY",
	}
}

// GetMappingSelectionDetailsSelectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSelectionDetailsSelectionModeEnum(val string) (SelectionDetailsSelectionModeEnum, bool) {
	enum, ok := mappingSelectionDetailsSelectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
