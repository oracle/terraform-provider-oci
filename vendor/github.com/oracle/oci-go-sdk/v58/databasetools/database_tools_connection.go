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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DatabaseToolsConnection Description of DatabaseToolsConnection.
type DatabaseToolsConnection interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DatabaseToolsConnection.
	GetId() *string

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the containing Compartment.
	GetCompartmentId() *string

	// The current state of the DatabaseToolsConnection.
	GetLifecycleState() LifecycleStateEnum

	// The time the DatabaseToolsConnection was created. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time the DatabaseToolsConnection was updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type databasetoolsconnection struct {
	JsonData         []byte
	Id               *string                           `mandatory:"true" json:"id"`
	DisplayName      *string                           `mandatory:"true" json:"displayName"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	LifecycleState   LifecycleStateEnum                `mandatory:"true" json:"lifecycleState"`
	TimeCreated      *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated      *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Type             string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsconnection) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsconnection databasetoolsconnection
	s := struct {
		Model Unmarshalerdatabasetoolsconnection
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.SystemTags = s.Model.SystemTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolsconnection) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ORACLE_DATABASE":
		mm := DatabaseToolsConnectionOracleDatabase{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m databasetoolsconnection) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m databasetoolsconnection) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m databasetoolsconnection) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetLifecycleState returns LifecycleState
func (m databasetoolsconnection) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

//GetTimeCreated returns TimeCreated
func (m databasetoolsconnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m databasetoolsconnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleDetails returns LifecycleDetails
func (m databasetoolsconnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetDefinedTags returns DefinedTags
func (m databasetoolsconnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetFreeformTags returns FreeformTags
func (m databasetoolsconnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetSystemTags returns SystemTags
func (m databasetoolsconnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m databasetoolsconnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsconnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
