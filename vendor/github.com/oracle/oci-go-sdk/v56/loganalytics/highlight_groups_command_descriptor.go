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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// HighlightGroupsCommandDescriptor Command descriptor for querylanguage HIGHLIGHTGROUPS command.
type HighlightGroupsCommandDescriptor struct {

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

	// User specified color to highlight matches with if found.
	Color *string `mandatory:"false" json:"color"`

	// User specified priority assigned to highlighted matches if found.
	Priority *string `mandatory:"false" json:"priority"`

	// List of fields to search for terms or phrases to highlight.  If not specified all string fields are scanned.
	MatchOnly []string `mandatory:"false" json:"matchOnly"`

	// List of fields to search for terms or phrases to highlight.
	Fields []string `mandatory:"false" json:"fields"`

	// List of terms or phrases to highlight if found.
	Keywords []string `mandatory:"false" json:"keywords"`

	// List of subQueries specified as highlightgroups command arguments
	SubQueries []ParseQueryOutput `mandatory:"false" json:"subQueries"`
}

//GetDisplayQueryString returns DisplayQueryString
func (m HighlightGroupsCommandDescriptor) GetDisplayQueryString() *string {
	return m.DisplayQueryString
}

//GetInternalQueryString returns InternalQueryString
func (m HighlightGroupsCommandDescriptor) GetInternalQueryString() *string {
	return m.InternalQueryString
}

//GetCategory returns Category
func (m HighlightGroupsCommandDescriptor) GetCategory() *string {
	return m.Category
}

//GetReferencedFields returns ReferencedFields
func (m HighlightGroupsCommandDescriptor) GetReferencedFields() []AbstractField {
	return m.ReferencedFields
}

//GetDeclaredFields returns DeclaredFields
func (m HighlightGroupsCommandDescriptor) GetDeclaredFields() []AbstractField {
	return m.DeclaredFields
}

func (m HighlightGroupsCommandDescriptor) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m HighlightGroupsCommandDescriptor) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHighlightGroupsCommandDescriptor HighlightGroupsCommandDescriptor
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeHighlightGroupsCommandDescriptor
	}{
		"HIGHLIGHT_GROUPS",
		(MarshalTypeHighlightGroupsCommandDescriptor)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *HighlightGroupsCommandDescriptor) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Category            *string            `json:"category"`
		ReferencedFields    []abstractfield    `json:"referencedFields"`
		DeclaredFields      []abstractfield    `json:"declaredFields"`
		Color               *string            `json:"color"`
		Priority            *string            `json:"priority"`
		MatchOnly           []string           `json:"matchOnly"`
		Fields              []string           `json:"fields"`
		Keywords            []string           `json:"keywords"`
		SubQueries          []ParseQueryOutput `json:"subQueries"`
		DisplayQueryString  *string            `json:"displayQueryString"`
		InternalQueryString *string            `json:"internalQueryString"`
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

	m.Color = model.Color

	m.Priority = model.Priority

	m.MatchOnly = make([]string, len(model.MatchOnly))
	for i, n := range model.MatchOnly {
		m.MatchOnly[i] = n
	}

	m.Fields = make([]string, len(model.Fields))
	for i, n := range model.Fields {
		m.Fields[i] = n
	}

	m.Keywords = make([]string, len(model.Keywords))
	for i, n := range model.Keywords {
		m.Keywords[i] = n
	}

	m.SubQueries = make([]ParseQueryOutput, len(model.SubQueries))
	for i, n := range model.SubQueries {
		m.SubQueries[i] = n
	}

	m.DisplayQueryString = model.DisplayQueryString

	m.InternalQueryString = model.InternalQueryString

	return
}
