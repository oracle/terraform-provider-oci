// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ClusterCompareCommandDescriptor Command descriptor for querylanguage CLUSTERCOMPARE command.
type ClusterCompareCommandDescriptor struct {

	// Command fragment display string from user specified query string formatted by query builder.
	DisplayQueryString *string `mandatory:"true" json:"displayQueryString"`

	// Command fragment internal string from user specified query string formatted by query builder.
	InternalQueryString *string `mandatory:"true" json:"internalQueryString"`

	// querylanguage command designation for example; reporting vs filtering
	Category *string `mandatory:"false" json:"category"`

	// Fields referenced in command fragment from user specified query string.
	ReferencedFields []AbstractField `mandatory:"false" json:"referencedFields"`

	// Fields declared in command fragment from user specified query string.
	DeclaredFields []AbstractField `mandatory:"false" json:"declaredFields"`

	// To shift time range of main query backwards using a relative time expression e.g -24hrs. E.g compare against the previous 24 hrs.
	TimeShift *string `mandatory:"false" json:"timeShift"`

	// Start time to apply to base line query if specified.
	TimeStart *int64 `mandatory:"false" json:"timeStart"`

	// End time to apply to base line query if specified.
	TimeEnd *int64 `mandatory:"false" json:"timeEnd"`

	// Option to calculate trends of each cluster if specified.
	ShouldIncludeTrends *bool `mandatory:"false" json:"shouldIncludeTrends"`

	// Option to control the size of buckets in the histogram e.g 8hrs  - each bar other than first and last should represent 8hr time span. Will be adjusted to a larger span if time range is very large.
	Span *string `mandatory:"false" json:"span"`

	// Query to use to compute base line to compare top level query results against to identify differences if specified.
	BaselineQuery *string `mandatory:"false" json:"baselineQuery"`
}

//GetDisplayQueryString returns DisplayQueryString
func (m ClusterCompareCommandDescriptor) GetDisplayQueryString() *string {
	return m.DisplayQueryString
}

//GetInternalQueryString returns InternalQueryString
func (m ClusterCompareCommandDescriptor) GetInternalQueryString() *string {
	return m.InternalQueryString
}

//GetCategory returns Category
func (m ClusterCompareCommandDescriptor) GetCategory() *string {
	return m.Category
}

//GetReferencedFields returns ReferencedFields
func (m ClusterCompareCommandDescriptor) GetReferencedFields() []AbstractField {
	return m.ReferencedFields
}

//GetDeclaredFields returns DeclaredFields
func (m ClusterCompareCommandDescriptor) GetDeclaredFields() []AbstractField {
	return m.DeclaredFields
}

func (m ClusterCompareCommandDescriptor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClusterCompareCommandDescriptor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ClusterCompareCommandDescriptor) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeClusterCompareCommandDescriptor ClusterCompareCommandDescriptor
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeClusterCompareCommandDescriptor
	}{
		"CLUSTER_COMPARE",
		(MarshalTypeClusterCompareCommandDescriptor)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ClusterCompareCommandDescriptor) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Category            *string         `json:"category"`
		ReferencedFields    []abstractfield `json:"referencedFields"`
		DeclaredFields      []abstractfield `json:"declaredFields"`
		TimeShift           *string         `json:"timeShift"`
		TimeStart           *int64          `json:"timeStart"`
		TimeEnd             *int64          `json:"timeEnd"`
		ShouldIncludeTrends *bool           `json:"shouldIncludeTrends"`
		Span                *string         `json:"span"`
		BaselineQuery       *string         `json:"baselineQuery"`
		DisplayQueryString  *string         `json:"displayQueryString"`
		InternalQueryString *string         `json:"internalQueryString"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Category = model.Category

	m.ReferencedFields = make([]AbstractField, len(model.ReferencedFields))
	for i, n := range model.ReferencedFields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ReferencedFields[i] = nn.(AbstractField)
		} else {
			m.ReferencedFields[i] = nil
		}
	}

	m.DeclaredFields = make([]AbstractField, len(model.DeclaredFields))
	for i, n := range model.DeclaredFields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.DeclaredFields[i] = nn.(AbstractField)
		} else {
			m.DeclaredFields[i] = nil
		}
	}

	m.TimeShift = model.TimeShift

	m.TimeStart = model.TimeStart

	m.TimeEnd = model.TimeEnd

	m.ShouldIncludeTrends = model.ShouldIncludeTrends

	m.Span = model.Span

	m.BaselineQuery = model.BaselineQuery

	m.DisplayQueryString = model.DisplayQueryString

	m.InternalQueryString = model.InternalQueryString

	return
}
