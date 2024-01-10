// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AbstractFrequencyDetails The model that holds the frequency details.
type AbstractFrequencyDetails interface {

	// the frequency of the schedule.
	GetFrequency() AbstractFrequencyDetailsFrequencyEnum
}

type abstractfrequencydetails struct {
	JsonData  []byte
	Frequency AbstractFrequencyDetailsFrequencyEnum `mandatory:"false" json:"frequency,omitempty"`
	ModelType string                                `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *abstractfrequencydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractfrequencydetails abstractfrequencydetails
	s := struct {
		Model Unmarshalerabstractfrequencydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Frequency = s.Model.Frequency
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractfrequencydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "MONTHLY":
		mm := MonthlyFrequencyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM":
		mm := CustomFrequencyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DAILY":
		mm := DailyFrequencyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WEEKLY":
		mm := WeeklyFrequencyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MONTHLY_RULE":
		mm := MonthlyRuleFrequencyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOURLY":
		mm := HourlyFrequencyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AbstractFrequencyDetails: %s.", m.ModelType)
		return *m, nil
	}
}

// GetFrequency returns Frequency
func (m abstractfrequencydetails) GetFrequency() AbstractFrequencyDetailsFrequencyEnum {
	return m.Frequency
}

func (m abstractfrequencydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m abstractfrequencydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAbstractFrequencyDetailsFrequencyEnum(string(m.Frequency)); !ok && m.Frequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Frequency: %s. Supported values are: %s.", m.Frequency, strings.Join(GetAbstractFrequencyDetailsFrequencyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AbstractFrequencyDetailsFrequencyEnum Enum with underlying type: string
type AbstractFrequencyDetailsFrequencyEnum string

// Set of constants representing the allowable values for AbstractFrequencyDetailsFrequencyEnum
const (
	AbstractFrequencyDetailsFrequencyHourly  AbstractFrequencyDetailsFrequencyEnum = "HOURLY"
	AbstractFrequencyDetailsFrequencyDaily   AbstractFrequencyDetailsFrequencyEnum = "DAILY"
	AbstractFrequencyDetailsFrequencyMonthly AbstractFrequencyDetailsFrequencyEnum = "MONTHLY"
	AbstractFrequencyDetailsFrequencyWeekly  AbstractFrequencyDetailsFrequencyEnum = "WEEKLY"
	AbstractFrequencyDetailsFrequencyCustom  AbstractFrequencyDetailsFrequencyEnum = "CUSTOM"
)

var mappingAbstractFrequencyDetailsFrequencyEnum = map[string]AbstractFrequencyDetailsFrequencyEnum{
	"HOURLY":  AbstractFrequencyDetailsFrequencyHourly,
	"DAILY":   AbstractFrequencyDetailsFrequencyDaily,
	"MONTHLY": AbstractFrequencyDetailsFrequencyMonthly,
	"WEEKLY":  AbstractFrequencyDetailsFrequencyWeekly,
	"CUSTOM":  AbstractFrequencyDetailsFrequencyCustom,
}

var mappingAbstractFrequencyDetailsFrequencyEnumLowerCase = map[string]AbstractFrequencyDetailsFrequencyEnum{
	"hourly":  AbstractFrequencyDetailsFrequencyHourly,
	"daily":   AbstractFrequencyDetailsFrequencyDaily,
	"monthly": AbstractFrequencyDetailsFrequencyMonthly,
	"weekly":  AbstractFrequencyDetailsFrequencyWeekly,
	"custom":  AbstractFrequencyDetailsFrequencyCustom,
}

// GetAbstractFrequencyDetailsFrequencyEnumValues Enumerates the set of values for AbstractFrequencyDetailsFrequencyEnum
func GetAbstractFrequencyDetailsFrequencyEnumValues() []AbstractFrequencyDetailsFrequencyEnum {
	values := make([]AbstractFrequencyDetailsFrequencyEnum, 0)
	for _, v := range mappingAbstractFrequencyDetailsFrequencyEnum {
		values = append(values, v)
	}
	return values
}

// GetAbstractFrequencyDetailsFrequencyEnumStringValues Enumerates the set of values in String for AbstractFrequencyDetailsFrequencyEnum
func GetAbstractFrequencyDetailsFrequencyEnumStringValues() []string {
	return []string{
		"HOURLY",
		"DAILY",
		"MONTHLY",
		"WEEKLY",
		"CUSTOM",
	}
}

// GetMappingAbstractFrequencyDetailsFrequencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAbstractFrequencyDetailsFrequencyEnum(val string) (AbstractFrequencyDetailsFrequencyEnum, bool) {
	enum, ok := mappingAbstractFrequencyDetailsFrequencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AbstractFrequencyDetailsModelTypeEnum Enum with underlying type: string
type AbstractFrequencyDetailsModelTypeEnum string

// Set of constants representing the allowable values for AbstractFrequencyDetailsModelTypeEnum
const (
	AbstractFrequencyDetailsModelTypeHourly      AbstractFrequencyDetailsModelTypeEnum = "HOURLY"
	AbstractFrequencyDetailsModelTypeDaily       AbstractFrequencyDetailsModelTypeEnum = "DAILY"
	AbstractFrequencyDetailsModelTypeMonthly     AbstractFrequencyDetailsModelTypeEnum = "MONTHLY"
	AbstractFrequencyDetailsModelTypeWeekly      AbstractFrequencyDetailsModelTypeEnum = "WEEKLY"
	AbstractFrequencyDetailsModelTypeMonthlyRule AbstractFrequencyDetailsModelTypeEnum = "MONTHLY_RULE"
	AbstractFrequencyDetailsModelTypeCustom      AbstractFrequencyDetailsModelTypeEnum = "CUSTOM"
)

var mappingAbstractFrequencyDetailsModelTypeEnum = map[string]AbstractFrequencyDetailsModelTypeEnum{
	"HOURLY":       AbstractFrequencyDetailsModelTypeHourly,
	"DAILY":        AbstractFrequencyDetailsModelTypeDaily,
	"MONTHLY":      AbstractFrequencyDetailsModelTypeMonthly,
	"WEEKLY":       AbstractFrequencyDetailsModelTypeWeekly,
	"MONTHLY_RULE": AbstractFrequencyDetailsModelTypeMonthlyRule,
	"CUSTOM":       AbstractFrequencyDetailsModelTypeCustom,
}

var mappingAbstractFrequencyDetailsModelTypeEnumLowerCase = map[string]AbstractFrequencyDetailsModelTypeEnum{
	"hourly":       AbstractFrequencyDetailsModelTypeHourly,
	"daily":        AbstractFrequencyDetailsModelTypeDaily,
	"monthly":      AbstractFrequencyDetailsModelTypeMonthly,
	"weekly":       AbstractFrequencyDetailsModelTypeWeekly,
	"monthly_rule": AbstractFrequencyDetailsModelTypeMonthlyRule,
	"custom":       AbstractFrequencyDetailsModelTypeCustom,
}

// GetAbstractFrequencyDetailsModelTypeEnumValues Enumerates the set of values for AbstractFrequencyDetailsModelTypeEnum
func GetAbstractFrequencyDetailsModelTypeEnumValues() []AbstractFrequencyDetailsModelTypeEnum {
	values := make([]AbstractFrequencyDetailsModelTypeEnum, 0)
	for _, v := range mappingAbstractFrequencyDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAbstractFrequencyDetailsModelTypeEnumStringValues Enumerates the set of values in String for AbstractFrequencyDetailsModelTypeEnum
func GetAbstractFrequencyDetailsModelTypeEnumStringValues() []string {
	return []string{
		"HOURLY",
		"DAILY",
		"MONTHLY",
		"WEEKLY",
		"MONTHLY_RULE",
		"CUSTOM",
	}
}

// GetMappingAbstractFrequencyDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAbstractFrequencyDetailsModelTypeEnum(val string) (AbstractFrequencyDetailsModelTypeEnum, bool) {
	enum, ok := mappingAbstractFrequencyDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
