// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v42/common"
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
	case "DAILY":
		mm := DailyFrequencyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOURLY":
		mm := HourlyFrequencyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetFrequency returns Frequency
func (m abstractfrequencydetails) GetFrequency() AbstractFrequencyDetailsFrequencyEnum {
	return m.Frequency
}

func (m abstractfrequencydetails) String() string {
	return common.PointerString(m)
}

// AbstractFrequencyDetailsFrequencyEnum Enum with underlying type: string
type AbstractFrequencyDetailsFrequencyEnum string

// Set of constants representing the allowable values for AbstractFrequencyDetailsFrequencyEnum
const (
	AbstractFrequencyDetailsFrequencyHourly  AbstractFrequencyDetailsFrequencyEnum = "HOURLY"
	AbstractFrequencyDetailsFrequencyDaily   AbstractFrequencyDetailsFrequencyEnum = "DAILY"
	AbstractFrequencyDetailsFrequencyMonthly AbstractFrequencyDetailsFrequencyEnum = "MONTHLY"
)

var mappingAbstractFrequencyDetailsFrequency = map[string]AbstractFrequencyDetailsFrequencyEnum{
	"HOURLY":  AbstractFrequencyDetailsFrequencyHourly,
	"DAILY":   AbstractFrequencyDetailsFrequencyDaily,
	"MONTHLY": AbstractFrequencyDetailsFrequencyMonthly,
}

// GetAbstractFrequencyDetailsFrequencyEnumValues Enumerates the set of values for AbstractFrequencyDetailsFrequencyEnum
func GetAbstractFrequencyDetailsFrequencyEnumValues() []AbstractFrequencyDetailsFrequencyEnum {
	values := make([]AbstractFrequencyDetailsFrequencyEnum, 0)
	for _, v := range mappingAbstractFrequencyDetailsFrequency {
		values = append(values, v)
	}
	return values
}

// AbstractFrequencyDetailsModelTypeEnum Enum with underlying type: string
type AbstractFrequencyDetailsModelTypeEnum string

// Set of constants representing the allowable values for AbstractFrequencyDetailsModelTypeEnum
const (
	AbstractFrequencyDetailsModelTypeHourly  AbstractFrequencyDetailsModelTypeEnum = "HOURLY"
	AbstractFrequencyDetailsModelTypeDaily   AbstractFrequencyDetailsModelTypeEnum = "DAILY"
	AbstractFrequencyDetailsModelTypeMonthly AbstractFrequencyDetailsModelTypeEnum = "MONTHLY"
)

var mappingAbstractFrequencyDetailsModelType = map[string]AbstractFrequencyDetailsModelTypeEnum{
	"HOURLY":  AbstractFrequencyDetailsModelTypeHourly,
	"DAILY":   AbstractFrequencyDetailsModelTypeDaily,
	"MONTHLY": AbstractFrequencyDetailsModelTypeMonthly,
}

// GetAbstractFrequencyDetailsModelTypeEnumValues Enumerates the set of values for AbstractFrequencyDetailsModelTypeEnum
func GetAbstractFrequencyDetailsModelTypeEnumValues() []AbstractFrequencyDetailsModelTypeEnum {
	values := make([]AbstractFrequencyDetailsModelTypeEnum, 0)
	for _, v := range mappingAbstractFrequencyDetailsModelType {
		values = append(values, v)
	}
	return values
}
