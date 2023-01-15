// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// AttributeProfileResult A summary of profiling results of a specific attribute.
type AttributeProfileResult interface {

	// Name of the attribute
	GetName() *string

	GetMin() *ProfileStat

	GetMax() *ProfileStat

	GetNullCount() *ProfileStat

	GetDistinctCount() *ProfileStat

	GetUniqueCount() *ProfileStat

	GetDuplicateCount() *ProfileStat

	// Top N value frequencies for the column as described already in the topNValueFrequency profile config property.
	GetValueFrequencies() []ObjectFreqStat
}

type attributeprofileresult struct {
	JsonData         []byte
	Name             *string          `mandatory:"false" json:"name"`
	Min              *ProfileStat     `mandatory:"false" json:"min"`
	Max              *ProfileStat     `mandatory:"false" json:"max"`
	NullCount        *ProfileStat     `mandatory:"false" json:"nullCount"`
	DistinctCount    *ProfileStat     `mandatory:"false" json:"distinctCount"`
	UniqueCount      *ProfileStat     `mandatory:"false" json:"uniqueCount"`
	DuplicateCount   *ProfileStat     `mandatory:"false" json:"duplicateCount"`
	ValueFrequencies []ObjectFreqStat `mandatory:"false" json:"valueFrequencies"`
	Type             string           `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *attributeprofileresult) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerattributeprofileresult attributeprofileresult
	s := struct {
		Model Unmarshalerattributeprofileresult
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Min = s.Model.Min
	m.Max = s.Model.Max
	m.NullCount = s.Model.NullCount
	m.DistinctCount = s.Model.DistinctCount
	m.UniqueCount = s.Model.UniqueCount
	m.DuplicateCount = s.Model.DuplicateCount
	m.ValueFrequencies = s.Model.ValueFrequencies
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *attributeprofileresult) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "STRING":
		mm := StringAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NUMERIC":
		mm := NumericAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATE_TIME":
		mm := DateAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m attributeprofileresult) GetName() *string {
	return m.Name
}

//GetMin returns Min
func (m attributeprofileresult) GetMin() *ProfileStat {
	return m.Min
}

//GetMax returns Max
func (m attributeprofileresult) GetMax() *ProfileStat {
	return m.Max
}

//GetNullCount returns NullCount
func (m attributeprofileresult) GetNullCount() *ProfileStat {
	return m.NullCount
}

//GetDistinctCount returns DistinctCount
func (m attributeprofileresult) GetDistinctCount() *ProfileStat {
	return m.DistinctCount
}

//GetUniqueCount returns UniqueCount
func (m attributeprofileresult) GetUniqueCount() *ProfileStat {
	return m.UniqueCount
}

//GetDuplicateCount returns DuplicateCount
func (m attributeprofileresult) GetDuplicateCount() *ProfileStat {
	return m.DuplicateCount
}

//GetValueFrequencies returns ValueFrequencies
func (m attributeprofileresult) GetValueFrequencies() []ObjectFreqStat {
	return m.ValueFrequencies
}

func (m attributeprofileresult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m attributeprofileresult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
