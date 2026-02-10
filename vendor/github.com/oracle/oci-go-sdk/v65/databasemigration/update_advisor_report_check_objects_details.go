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

// UpdateAdvisorReportCheckObjectsDetails Optional additional properties for update advisor report check. Default kind is LIST_OBJECTS
type UpdateAdvisorReportCheckObjectsDetails interface {
}

type updateadvisorreportcheckobjectsdetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *updateadvisorreportcheckobjectsdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateadvisorreportcheckobjectsdetails updateadvisorreportcheckobjectsdetails
	s := struct {
		Model Unmarshalerupdateadvisorreportcheckobjectsdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateadvisorreportcheckobjectsdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "ALL_OBJECTS":
		mm := AllUpdateAdvisorReportCheckObjectsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LIST_OBJECTS":
		mm := ListUpdateAdvisorReportCheckObjectsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateAdvisorReportCheckObjectsDetails: %s.", m.Kind)
		return *m, nil
	}
}

func (m updateadvisorreportcheckobjectsdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateadvisorreportcheckobjectsdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateAdvisorReportCheckObjectsDetailsKindEnum Enum with underlying type: string
type UpdateAdvisorReportCheckObjectsDetailsKindEnum string

// Set of constants representing the allowable values for UpdateAdvisorReportCheckObjectsDetailsKindEnum
const (
	UpdateAdvisorReportCheckObjectsDetailsKindAllObjects  UpdateAdvisorReportCheckObjectsDetailsKindEnum = "ALL_OBJECTS"
	UpdateAdvisorReportCheckObjectsDetailsKindListObjects UpdateAdvisorReportCheckObjectsDetailsKindEnum = "LIST_OBJECTS"
)

var mappingUpdateAdvisorReportCheckObjectsDetailsKindEnum = map[string]UpdateAdvisorReportCheckObjectsDetailsKindEnum{
	"ALL_OBJECTS":  UpdateAdvisorReportCheckObjectsDetailsKindAllObjects,
	"LIST_OBJECTS": UpdateAdvisorReportCheckObjectsDetailsKindListObjects,
}

var mappingUpdateAdvisorReportCheckObjectsDetailsKindEnumLowerCase = map[string]UpdateAdvisorReportCheckObjectsDetailsKindEnum{
	"all_objects":  UpdateAdvisorReportCheckObjectsDetailsKindAllObjects,
	"list_objects": UpdateAdvisorReportCheckObjectsDetailsKindListObjects,
}

// GetUpdateAdvisorReportCheckObjectsDetailsKindEnumValues Enumerates the set of values for UpdateAdvisorReportCheckObjectsDetailsKindEnum
func GetUpdateAdvisorReportCheckObjectsDetailsKindEnumValues() []UpdateAdvisorReportCheckObjectsDetailsKindEnum {
	values := make([]UpdateAdvisorReportCheckObjectsDetailsKindEnum, 0)
	for _, v := range mappingUpdateAdvisorReportCheckObjectsDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAdvisorReportCheckObjectsDetailsKindEnumStringValues Enumerates the set of values in String for UpdateAdvisorReportCheckObjectsDetailsKindEnum
func GetUpdateAdvisorReportCheckObjectsDetailsKindEnumStringValues() []string {
	return []string{
		"ALL_OBJECTS",
		"LIST_OBJECTS",
	}
}

// GetMappingUpdateAdvisorReportCheckObjectsDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAdvisorReportCheckObjectsDetailsKindEnum(val string) (UpdateAdvisorReportCheckObjectsDetailsKindEnum, bool) {
	enum, ok := mappingUpdateAdvisorReportCheckObjectsDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
