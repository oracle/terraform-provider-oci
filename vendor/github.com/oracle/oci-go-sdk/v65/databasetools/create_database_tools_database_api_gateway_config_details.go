// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDatabaseToolsDatabaseApiGatewayConfigDetails Details for the new Database Tools database API gateway config.
type CreateDatabaseToolsDatabaseApiGatewayConfigDetails interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools database API gateway config.
	GetCompartmentId() *string

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The RESTful service definition location.
	GetMetadataSource() DatabaseApiGatewayConfigMetadataSourceEnum

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Locks associated with this resource.
	GetLocks() []ResourceLock
}

type createdatabasetoolsdatabaseapigatewayconfigdetails struct {
	JsonData       []byte
	DefinedTags    map[string]map[string]interface{}          `mandatory:"false" json:"definedTags"`
	FreeformTags   map[string]string                          `mandatory:"false" json:"freeformTags"`
	Locks          []ResourceLock                             `mandatory:"false" json:"locks"`
	CompartmentId  *string                                    `mandatory:"true" json:"compartmentId"`
	DisplayName    *string                                    `mandatory:"true" json:"displayName"`
	MetadataSource DatabaseApiGatewayConfigMetadataSourceEnum `mandatory:"true" json:"metadataSource"`
	Type           string                                     `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createdatabasetoolsdatabaseapigatewayconfigdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedatabasetoolsdatabaseapigatewayconfigdetails createdatabasetoolsdatabaseapigatewayconfigdetails
	s := struct {
		Model Unmarshalercreatedatabasetoolsdatabaseapigatewayconfigdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.MetadataSource = s.Model.MetadataSource
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.Locks = s.Model.Locks
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdatabasetoolsdatabaseapigatewayconfigdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := CreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDatabaseToolsDatabaseApiGatewayConfigDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDefinedTags returns DefinedTags
func (m createdatabasetoolsdatabaseapigatewayconfigdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m createdatabasetoolsdatabaseapigatewayconfigdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetLocks returns Locks
func (m createdatabasetoolsdatabaseapigatewayconfigdetails) GetLocks() []ResourceLock {
	return m.Locks
}

// GetCompartmentId returns CompartmentId
func (m createdatabasetoolsdatabaseapigatewayconfigdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m createdatabasetoolsdatabaseapigatewayconfigdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetMetadataSource returns MetadataSource
func (m createdatabasetoolsdatabaseapigatewayconfigdetails) GetMetadataSource() DatabaseApiGatewayConfigMetadataSourceEnum {
	return m.MetadataSource
}

func (m createdatabasetoolsdatabaseapigatewayconfigdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdatabasetoolsdatabaseapigatewayconfigdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseApiGatewayConfigMetadataSourceEnum(string(m.MetadataSource)); !ok && m.MetadataSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetadataSource: %s. Supported values are: %s.", m.MetadataSource, strings.Join(GetDatabaseApiGatewayConfigMetadataSourceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
