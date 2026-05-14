// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsDatabaseApiGatewayConfigPoolSummary Summary of the Database Tools database API gateway config pool resource.
type DatabaseToolsDatabaseApiGatewayConfigPoolSummary interface {

	// A system generated string that uniquely identifies a pool sub resource.
	GetKey() *string

	// A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection.
	GetDatabaseToolsConnectionId() *string

	// The time the resource was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the resource was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime
}

type databasetoolsdatabaseapigatewayconfigpoolsummary struct {
	JsonData                  []byte
	Key                       *string         `mandatory:"true" json:"key"`
	DisplayName               *string         `mandatory:"true" json:"displayName"`
	DatabaseToolsConnectionId *string         `mandatory:"true" json:"databaseToolsConnectionId"`
	TimeCreated               *common.SDKTime `mandatory:"true" json:"timeCreated"`
	TimeUpdated               *common.SDKTime `mandatory:"true" json:"timeUpdated"`
	Type                      string          `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsdatabaseapigatewayconfigpoolsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsdatabaseapigatewayconfigpoolsummary databasetoolsdatabaseapigatewayconfigpoolsummary
	s := struct {
		Model Unmarshalerdatabasetoolsdatabaseapigatewayconfigpoolsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.DisplayName = s.Model.DisplayName
	m.DatabaseToolsConnectionId = s.Model.DatabaseToolsConnectionId
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolsdatabaseapigatewayconfigpoolsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := DatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsDatabaseApiGatewayConfigPoolSummary: %s.", m.Type)
		return *m, nil
	}
}

// GetKey returns Key
func (m databasetoolsdatabaseapigatewayconfigpoolsummary) GetKey() *string {
	return m.Key
}

// GetDisplayName returns DisplayName
func (m databasetoolsdatabaseapigatewayconfigpoolsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m databasetoolsdatabaseapigatewayconfigpoolsummary) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

// GetTimeCreated returns TimeCreated
func (m databasetoolsdatabaseapigatewayconfigpoolsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m databasetoolsdatabaseapigatewayconfigpoolsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m databasetoolsdatabaseapigatewayconfigpoolsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsdatabaseapigatewayconfigpoolsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
