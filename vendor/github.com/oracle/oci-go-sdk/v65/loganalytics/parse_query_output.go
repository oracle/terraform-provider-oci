// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ParseQueryOutput Returns a parser agnostic breakdown of a query string for client query string introspection.
type ParseQueryOutput struct {

	// Display string formatted by query builder of user specified query string.
	DisplayQueryString *string `mandatory:"true" json:"displayQueryString"`

	// Internal string formatted by query builder of user specified query string.
	InternalQueryString *string `mandatory:"true" json:"internalQueryString"`

	// List of columns returned by the specified query string as result output.
	Columns []AbstractColumn `mandatory:"false" json:"columns"`

	// Operation response time.
	ResponseTimeInMs *int64 `mandatory:"false" json:"responseTimeInMs"`

	TimeFilter *TimeRange `mandatory:"false" json:"timeFilter"`

	// List of querylanguage command descriptors, describing the specfied query string.
	Commands []AbstractCommandDescriptor `mandatory:"false" json:"commands"`

	// Optional list of variables for the query.
	Variables []VariableDefinition `mandatory:"false" json:"variables"`

	// Optional list of properties for the query.
	Properties []PropertyDefinition `mandatory:"false" json:"properties"`
}

func (m ParseQueryOutput) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ParseQueryOutput) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ParseQueryOutput) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Columns             []abstractcolumn            `json:"columns"`
		ResponseTimeInMs    *int64                      `json:"responseTimeInMs"`
		TimeFilter          *TimeRange                  `json:"timeFilter"`
		Commands            []abstractcommanddescriptor `json:"commands"`
		Variables           []VariableDefinition        `json:"variables"`
		Properties          []PropertyDefinition        `json:"properties"`
		DisplayQueryString  *string                     `json:"displayQueryString"`
		InternalQueryString *string                     `json:"internalQueryString"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Columns = make([]AbstractColumn, len(model.Columns))
	for i, n := range model.Columns {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Columns[i] = nn.(AbstractColumn)
		} else {
			m.Columns[i] = nil
		}
	}
	m.ResponseTimeInMs = model.ResponseTimeInMs

	m.TimeFilter = model.TimeFilter

	m.Commands = make([]AbstractCommandDescriptor, len(model.Commands))
	for i, n := range model.Commands {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Commands[i] = nn.(AbstractCommandDescriptor)
		} else {
			m.Commands[i] = nil
		}
	}
	m.Variables = make([]VariableDefinition, len(model.Variables))
	copy(m.Variables, model.Variables)
	m.Properties = make([]PropertyDefinition, len(model.Properties))
	copy(m.Properties, model.Properties)
	m.DisplayQueryString = model.DisplayQueryString

	m.InternalQueryString = model.InternalQueryString

	return
}
