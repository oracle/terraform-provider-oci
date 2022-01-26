// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateDatabaseToolsConnectionDetails The information about new DatabaseToolsConnection.
type CreateDatabaseToolsConnectionDetails interface {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the containing Compartment.
	GetCompartmentId() *string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string
}

type createdatabasetoolsconnectiondetails struct {
	JsonData      []byte
	DisplayName   *string                           `mandatory:"true" json:"displayName"`
	CompartmentId *string                           `mandatory:"true" json:"compartmentId"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	FreeformTags  map[string]string                 `mandatory:"false" json:"freeformTags"`
	Type          string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createdatabasetoolsconnectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedatabasetoolsconnectiondetails createdatabasetoolsconnectiondetails
	s := struct {
		Model Unmarshalercreatedatabasetoolsconnectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdatabasetoolsconnectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ORACLE_DATABASE":
		mm := CreateDatabaseToolsConnectionOracleDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m createdatabasetoolsconnectiondetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m createdatabasetoolsconnectiondetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDefinedTags returns DefinedTags
func (m createdatabasetoolsconnectiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetFreeformTags returns FreeformTags
func (m createdatabasetoolsconnectiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

func (m createdatabasetoolsconnectiondetails) String() string {
	return common.PointerString(m)
}
