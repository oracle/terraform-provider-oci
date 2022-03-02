// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v60/common"
	"strings"
)

// DateAttribute A summary of profiling results of a specefic attribute.
type DateAttribute struct {

	// Name of attribute
	Name *string `mandatory:"false" json:"name"`

	Min *ProfileStat `mandatory:"false" json:"min"`

	Max *ProfileStat `mandatory:"false" json:"max"`

	NullCount *ProfileStat `mandatory:"false" json:"nullCount"`

	DistinctCount *ProfileStat `mandatory:"false" json:"distinctCount"`

	UniqueCount *ProfileStat `mandatory:"false" json:"uniqueCount"`

	DuplicateCount *ProfileStat `mandatory:"false" json:"duplicateCount"`

	// Top N value frequencies for the column as described already in profile config topNValueFrequency property.
	ValueFrequencies []ObjectFreqStat `mandatory:"false" json:"valueFrequencies"`
}

//GetName returns Name
func (m DateAttribute) GetName() *string {
	return m.Name
}

//GetMin returns Min
func (m DateAttribute) GetMin() *ProfileStat {
	return m.Min
}

//GetMax returns Max
func (m DateAttribute) GetMax() *ProfileStat {
	return m.Max
}

//GetNullCount returns NullCount
func (m DateAttribute) GetNullCount() *ProfileStat {
	return m.NullCount
}

//GetDistinctCount returns DistinctCount
func (m DateAttribute) GetDistinctCount() *ProfileStat {
	return m.DistinctCount
}

//GetUniqueCount returns UniqueCount
func (m DateAttribute) GetUniqueCount() *ProfileStat {
	return m.UniqueCount
}

//GetDuplicateCount returns DuplicateCount
func (m DateAttribute) GetDuplicateCount() *ProfileStat {
	return m.DuplicateCount
}

//GetValueFrequencies returns ValueFrequencies
func (m DateAttribute) GetValueFrequencies() []ObjectFreqStat {
	return m.ValueFrequencies
}

func (m DateAttribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DateAttribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DateAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDateAttribute DateAttribute
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDateAttribute
	}{
		"DATE_TIME",
		(MarshalTypeDateAttribute)(m),
	}

	return json.Marshal(&s)
}
