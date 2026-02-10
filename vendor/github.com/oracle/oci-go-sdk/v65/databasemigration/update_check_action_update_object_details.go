// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateCheckActionUpdateObjectDetails Details for UpdateCheckActionUpdateObject.
type UpdateCheckActionUpdateObjectDetails interface {
}

type updatecheckactionupdateobjectdetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *updatecheckactionupdateobjectdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatecheckactionupdateobjectdetails updatecheckactionupdateobjectdetails
	s := struct {
		Model Unmarshalerupdatecheckactionupdateobjectdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatecheckactionupdateobjectdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "LIST_OBJECTS":
		mm := ListUpdateCheckActionUpdateObjectDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ALL_OBJECTS":
		mm := AllUpdateCheckActionUpdateObjectDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateCheckActionUpdateObjectDetails: %s.", m.Kind)
		return *m, nil
	}
}

func (m updatecheckactionupdateobjectdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatecheckactionupdateobjectdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateCheckActionUpdateObjectDetailsKindEnum Enum with underlying type: string
type UpdateCheckActionUpdateObjectDetailsKindEnum string

// Set of constants representing the allowable values for UpdateCheckActionUpdateObjectDetailsKindEnum
const (
	UpdateCheckActionUpdateObjectDetailsKindAllObjects  UpdateCheckActionUpdateObjectDetailsKindEnum = "ALL_OBJECTS"
	UpdateCheckActionUpdateObjectDetailsKindListObjects UpdateCheckActionUpdateObjectDetailsKindEnum = "LIST_OBJECTS"
)

var mappingUpdateCheckActionUpdateObjectDetailsKindEnum = map[string]UpdateCheckActionUpdateObjectDetailsKindEnum{
	"ALL_OBJECTS":  UpdateCheckActionUpdateObjectDetailsKindAllObjects,
	"LIST_OBJECTS": UpdateCheckActionUpdateObjectDetailsKindListObjects,
}

var mappingUpdateCheckActionUpdateObjectDetailsKindEnumLowerCase = map[string]UpdateCheckActionUpdateObjectDetailsKindEnum{
	"all_objects":  UpdateCheckActionUpdateObjectDetailsKindAllObjects,
	"list_objects": UpdateCheckActionUpdateObjectDetailsKindListObjects,
}

// GetUpdateCheckActionUpdateObjectDetailsKindEnumValues Enumerates the set of values for UpdateCheckActionUpdateObjectDetailsKindEnum
func GetUpdateCheckActionUpdateObjectDetailsKindEnumValues() []UpdateCheckActionUpdateObjectDetailsKindEnum {
	values := make([]UpdateCheckActionUpdateObjectDetailsKindEnum, 0)
	for _, v := range mappingUpdateCheckActionUpdateObjectDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateCheckActionUpdateObjectDetailsKindEnumStringValues Enumerates the set of values in String for UpdateCheckActionUpdateObjectDetailsKindEnum
func GetUpdateCheckActionUpdateObjectDetailsKindEnumStringValues() []string {
	return []string{
		"ALL_OBJECTS",
		"LIST_OBJECTS",
	}
}

// GetMappingUpdateCheckActionUpdateObjectDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateCheckActionUpdateObjectDetailsKindEnum(val string) (UpdateCheckActionUpdateObjectDetailsKindEnum, bool) {
	enum, ok := mappingUpdateCheckActionUpdateObjectDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
