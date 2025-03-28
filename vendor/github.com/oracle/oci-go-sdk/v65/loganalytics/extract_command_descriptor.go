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

// ExtractCommandDescriptor Command descriptor for querylanguage EXTRACT command.
type ExtractCommandDescriptor struct {

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

	// Field denoting if this is a hidden command that is not shown in the query string.
	IsHidden *bool `mandatory:"false" json:"isHidden"`
}

// GetDisplayQueryString returns DisplayQueryString
func (m ExtractCommandDescriptor) GetDisplayQueryString() *string {
	return m.DisplayQueryString
}

// GetInternalQueryString returns InternalQueryString
func (m ExtractCommandDescriptor) GetInternalQueryString() *string {
	return m.InternalQueryString
}

// GetCategory returns Category
func (m ExtractCommandDescriptor) GetCategory() *string {
	return m.Category
}

// GetReferencedFields returns ReferencedFields
func (m ExtractCommandDescriptor) GetReferencedFields() []AbstractField {
	return m.ReferencedFields
}

// GetDeclaredFields returns DeclaredFields
func (m ExtractCommandDescriptor) GetDeclaredFields() []AbstractField {
	return m.DeclaredFields
}

// GetIsHidden returns IsHidden
func (m ExtractCommandDescriptor) GetIsHidden() *bool {
	return m.IsHidden
}

func (m ExtractCommandDescriptor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtractCommandDescriptor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExtractCommandDescriptor) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExtractCommandDescriptor ExtractCommandDescriptor
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeExtractCommandDescriptor
	}{
		"EXTRACT",
		(MarshalTypeExtractCommandDescriptor)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ExtractCommandDescriptor) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Category            *string         `json:"category"`
		ReferencedFields    []abstractfield `json:"referencedFields"`
		DeclaredFields      []abstractfield `json:"declaredFields"`
		IsHidden            *bool           `json:"isHidden"`
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
	m.IsHidden = model.IsHidden

	m.DisplayQueryString = model.DisplayQueryString

	m.InternalQueryString = model.InternalQueryString

	return
}
