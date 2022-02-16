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

// DeleteCommandDescriptor Command descriptor for querylanguage DELETE command.
type DeleteCommandDescriptor struct {

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

	// Value specified in DELETE command in queryString as to whether the delete is a dry-run (only report number of rows removed) rather than actually remove matching log records.
	IsDryRun *bool `mandatory:"false" json:"isDryRun"`
}

//GetDisplayQueryString returns DisplayQueryString
func (m DeleteCommandDescriptor) GetDisplayQueryString() *string {
	return m.DisplayQueryString
}

//GetInternalQueryString returns InternalQueryString
func (m DeleteCommandDescriptor) GetInternalQueryString() *string {
	return m.InternalQueryString
}

//GetCategory returns Category
func (m DeleteCommandDescriptor) GetCategory() *string {
	return m.Category
}

//GetReferencedFields returns ReferencedFields
func (m DeleteCommandDescriptor) GetReferencedFields() []AbstractField {
	return m.ReferencedFields
}

//GetDeclaredFields returns DeclaredFields
func (m DeleteCommandDescriptor) GetDeclaredFields() []AbstractField {
	return m.DeclaredFields
}

func (m DeleteCommandDescriptor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeleteCommandDescriptor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DeleteCommandDescriptor) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDeleteCommandDescriptor DeleteCommandDescriptor
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeDeleteCommandDescriptor
	}{
		"DELETE",
		(MarshalTypeDeleteCommandDescriptor)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DeleteCommandDescriptor) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Category            *string         `json:"category"`
		ReferencedFields    []abstractfield `json:"referencedFields"`
		DeclaredFields      []abstractfield `json:"declaredFields"`
		IsDryRun            *bool           `json:"isDryRun"`
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

	m.IsDryRun = model.IsDryRun

	m.DisplayQueryString = model.DisplayQueryString

	m.InternalQueryString = model.InternalQueryString

	return
}
