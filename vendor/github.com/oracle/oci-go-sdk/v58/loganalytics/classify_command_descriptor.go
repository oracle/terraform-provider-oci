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

// ClassifyCommandDescriptor Command descriptor for querylanguage CLASSIFY command.
type ClassifyCommandDescriptor struct {

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

	// Value specified in CLASSIFY command in queryString if set limits the results returned to top N.
	TopCount *int `mandatory:"false" json:"topCount"`

	// Value specified in CLASSIFY command in queryString if set limits the results returned to bottom N.
	BottomCount *int `mandatory:"false" json:"bottomCount"`

	// Fields specified in CLASSIFY command in queryString if set include / exclude fields in correlate results.
	Correlate []FieldsAddRemoveField `mandatory:"false" json:"correlate"`
}

//GetDisplayQueryString returns DisplayQueryString
func (m ClassifyCommandDescriptor) GetDisplayQueryString() *string {
	return m.DisplayQueryString
}

//GetInternalQueryString returns InternalQueryString
func (m ClassifyCommandDescriptor) GetInternalQueryString() *string {
	return m.InternalQueryString
}

//GetCategory returns Category
func (m ClassifyCommandDescriptor) GetCategory() *string {
	return m.Category
}

//GetReferencedFields returns ReferencedFields
func (m ClassifyCommandDescriptor) GetReferencedFields() []AbstractField {
	return m.ReferencedFields
}

//GetDeclaredFields returns DeclaredFields
func (m ClassifyCommandDescriptor) GetDeclaredFields() []AbstractField {
	return m.DeclaredFields
}

func (m ClassifyCommandDescriptor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClassifyCommandDescriptor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ClassifyCommandDescriptor) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeClassifyCommandDescriptor ClassifyCommandDescriptor
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeClassifyCommandDescriptor
	}{
		"CLASSIFY",
		(MarshalTypeClassifyCommandDescriptor)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ClassifyCommandDescriptor) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Category            *string                `json:"category"`
		ReferencedFields    []abstractfield        `json:"referencedFields"`
		DeclaredFields      []abstractfield        `json:"declaredFields"`
		TopCount            *int                   `json:"topCount"`
		BottomCount         *int                   `json:"bottomCount"`
		Correlate           []FieldsAddRemoveField `json:"correlate"`
		DisplayQueryString  *string                `json:"displayQueryString"`
		InternalQueryString *string                `json:"internalQueryString"`
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

	m.TopCount = model.TopCount

	m.BottomCount = model.BottomCount

	m.Correlate = make([]FieldsAddRemoveField, len(model.Correlate))
	for i, n := range model.Correlate {
		m.Correlate[i] = n
	}

	m.DisplayQueryString = model.DisplayQueryString

	m.InternalQueryString = model.InternalQueryString

	return
}
