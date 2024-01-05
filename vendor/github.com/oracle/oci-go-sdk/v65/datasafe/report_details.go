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

// ReportDetails The details of the report schedule.
type ReportDetails interface {
}

type reportdetails struct {
	JsonData   []byte
	ReportType string `json:"reportType"`
}

// UnmarshalJSON unmarshals json
func (m *reportdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerreportdetails reportdetails
	s := struct {
		Model Unmarshalerreportdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ReportType = s.Model.ReportType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *reportdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ReportType {
	case "AUDIT":
		mm := ScheduleAuditReportDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ReportDetails: %s.", m.ReportType)
		return *m, nil
	}
}

func (m reportdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m reportdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReportDetailsReportTypeEnum Enum with underlying type: string
type ReportDetailsReportTypeEnum string

// Set of constants representing the allowable values for ReportDetailsReportTypeEnum
const (
	ReportDetailsReportTypeAudit ReportDetailsReportTypeEnum = "AUDIT"
)

var mappingReportDetailsReportTypeEnum = map[string]ReportDetailsReportTypeEnum{
	"AUDIT": ReportDetailsReportTypeAudit,
}

var mappingReportDetailsReportTypeEnumLowerCase = map[string]ReportDetailsReportTypeEnum{
	"audit": ReportDetailsReportTypeAudit,
}

// GetReportDetailsReportTypeEnumValues Enumerates the set of values for ReportDetailsReportTypeEnum
func GetReportDetailsReportTypeEnumValues() []ReportDetailsReportTypeEnum {
	values := make([]ReportDetailsReportTypeEnum, 0)
	for _, v := range mappingReportDetailsReportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReportDetailsReportTypeEnumStringValues Enumerates the set of values in String for ReportDetailsReportTypeEnum
func GetReportDetailsReportTypeEnumStringValues() []string {
	return []string{
		"AUDIT",
	}
}

// GetMappingReportDetailsReportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportDetailsReportTypeEnum(val string) (ReportDetailsReportTypeEnum, bool) {
	enum, ok := mappingReportDetailsReportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
