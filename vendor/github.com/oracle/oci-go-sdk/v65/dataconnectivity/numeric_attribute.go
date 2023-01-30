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

// NumericAttribute A summary of profiling results of a specific attribute.
type NumericAttribute struct {

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

	Mean *ProfileStat `mandatory:"false" json:"mean"`

	Median *ProfileStat `mandatory:"false" json:"median"`

	StandardDeviation *ProfileStat `mandatory:"false" json:"standardDeviation"`

	Variance *ProfileStat `mandatory:"false" json:"variance"`

	Outlier *Outlier `mandatory:"false" json:"outlier"`

	Histogram *Histogram `mandatory:"false" json:"histogram"`

	// Pattern frequencies for the column as described already in profile config.
	PatternFrequencies []ObjectFreqStat `mandatory:"false" json:"patternFrequencies"`
}

//GetName returns Name
func (m NumericAttribute) GetName() *string {
	return m.Name
}

//GetMin returns Min
func (m NumericAttribute) GetMin() *ProfileStat {
	return m.Min
}

//GetMax returns Max
func (m NumericAttribute) GetMax() *ProfileStat {
	return m.Max
}

//GetNullCount returns NullCount
func (m NumericAttribute) GetNullCount() *ProfileStat {
	return m.NullCount
}

//GetDistinctCount returns DistinctCount
func (m NumericAttribute) GetDistinctCount() *ProfileStat {
	return m.DistinctCount
}

//GetUniqueCount returns UniqueCount
func (m NumericAttribute) GetUniqueCount() *ProfileStat {
	return m.UniqueCount
}

//GetDuplicateCount returns DuplicateCount
func (m NumericAttribute) GetDuplicateCount() *ProfileStat {
	return m.DuplicateCount
}

//GetValueFrequencies returns ValueFrequencies
func (m NumericAttribute) GetValueFrequencies() []ObjectFreqStat {
	return m.ValueFrequencies
}

func (m NumericAttribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NumericAttribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NumericAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNumericAttribute NumericAttribute
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeNumericAttribute
	}{
		"NUMERIC",
		(MarshalTypeNumericAttribute)(m),
	}

	return json.Marshal(&s)
}
