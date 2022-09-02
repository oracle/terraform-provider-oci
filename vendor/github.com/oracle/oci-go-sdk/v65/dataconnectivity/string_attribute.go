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

// StringAttribute A summary of profiling results of a specefic attribute.
type StringAttribute struct {

	// Name of the attribute
	Name *string `mandatory:"false" json:"name"`

	Min *ProfileStat `mandatory:"false" json:"min"`

	Max *ProfileStat `mandatory:"false" json:"max"`

	NullCount *ProfileStat `mandatory:"false" json:"nullCount"`

	DistinctCount *ProfileStat `mandatory:"false" json:"distinctCount"`

	UniqueCount *ProfileStat `mandatory:"false" json:"uniqueCount"`

	DuplicateCount *ProfileStat `mandatory:"false" json:"duplicateCount"`

	// Top N value frequencies for the column as described already in the topNValueFrequency profile config property.
	ValueFrequencies []ObjectFreqStat `mandatory:"false" json:"valueFrequencies"`

	MinLength *ProfileStat `mandatory:"false" json:"minLength"`

	MaxLength *ProfileStat `mandatory:"false" json:"maxLength"`

	MeanLength *ProfileStat `mandatory:"false" json:"meanLength"`

	// Pattern frequencies for the column as described in the profile config.
	PatternFrequencies []ObjectFreqStat `mandatory:"false" json:"patternFrequencies"`

	// Inferred DataType for the column.
	InferredDataTypes []DataTypeStat `mandatory:"false" json:"inferredDataTypes"`
}

//GetName returns Name
func (m StringAttribute) GetName() *string {
	return m.Name
}

//GetMin returns Min
func (m StringAttribute) GetMin() *ProfileStat {
	return m.Min
}

//GetMax returns Max
func (m StringAttribute) GetMax() *ProfileStat {
	return m.Max
}

//GetNullCount returns NullCount
func (m StringAttribute) GetNullCount() *ProfileStat {
	return m.NullCount
}

//GetDistinctCount returns DistinctCount
func (m StringAttribute) GetDistinctCount() *ProfileStat {
	return m.DistinctCount
}

//GetUniqueCount returns UniqueCount
func (m StringAttribute) GetUniqueCount() *ProfileStat {
	return m.UniqueCount
}

//GetDuplicateCount returns DuplicateCount
func (m StringAttribute) GetDuplicateCount() *ProfileStat {
	return m.DuplicateCount
}

//GetValueFrequencies returns ValueFrequencies
func (m StringAttribute) GetValueFrequencies() []ObjectFreqStat {
	return m.ValueFrequencies
}

func (m StringAttribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StringAttribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StringAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStringAttribute StringAttribute
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeStringAttribute
	}{
		"STRING",
		(MarshalTypeStringAttribute)(m),
	}

	return json.Marshal(&s)
}
