// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ArrayVariable Array variable.
type ArrayVariable struct {

	// The display name for the variable as shown in the UI.
	Title *string `mandatory:"false" json:"title"`

	// Detailed information about this variable's purpose and usage.
	Description *string `mandatory:"false" json:"description"`

	// Indicates if this input variable is required for stack execution.
	IsRequired *bool `mandatory:"false" json:"isRequired"`

	// Hint to control whether this variable is visible.
	Visible *string `mandatory:"false" json:"visible"`

	Items BaseVariable `mandatory:"false" json:"items"`

	// Maximum allowed items in the array.
	MaxItems *int `mandatory:"false" json:"maxItems"`

	// Minimum allowed items in the array.
	MinItems *int `mandatory:"false" json:"minItems"`

	// If true, array entries will be unique.
	AreUniqueItems *bool `mandatory:"false" json:"areUniqueItems"`

	// The default value for this variable.
	DefaultValue *interface{} `mandatory:"false" json:"defaultValue"`

	Contains BaseVariable `mandatory:"false" json:"contains"`
}

// GetTitle returns Title
func (m ArrayVariable) GetTitle() *string {
	return m.Title
}

// GetDescription returns Description
func (m ArrayVariable) GetDescription() *string {
	return m.Description
}

// GetIsRequired returns IsRequired
func (m ArrayVariable) GetIsRequired() *bool {
	return m.IsRequired
}

// GetVisible returns Visible
func (m ArrayVariable) GetVisible() *string {
	return m.Visible
}

func (m ArrayVariable) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ArrayVariable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ArrayVariable) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeArrayVariable ArrayVariable
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeArrayVariable
	}{
		"ARRAY",
		(MarshalTypeArrayVariable)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ArrayVariable) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Title          *string      `json:"title"`
		Description    *string      `json:"description"`
		IsRequired     *bool        `json:"isRequired"`
		Visible        *string      `json:"visible"`
		Items          basevariable `json:"items"`
		MaxItems       *int         `json:"maxItems"`
		MinItems       *int         `json:"minItems"`
		AreUniqueItems *bool        `json:"areUniqueItems"`
		DefaultValue   *interface{} `json:"defaultValue"`
		Contains       basevariable `json:"contains"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Title = model.Title

	m.Description = model.Description

	m.IsRequired = model.IsRequired

	m.Visible = model.Visible

	nn, e = model.Items.UnmarshalPolymorphicJSON(model.Items.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Items = nn.(BaseVariable)
	} else {
		m.Items = nil
	}

	m.MaxItems = model.MaxItems

	m.MinItems = model.MinItems

	m.AreUniqueItems = model.AreUniqueItems

	m.DefaultValue = model.DefaultValue

	nn, e = model.Contains.UnmarshalPolymorphicJSON(model.Contains.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Contains = nn.(BaseVariable)
	} else {
		m.Contains = nil
	}

	return
}
