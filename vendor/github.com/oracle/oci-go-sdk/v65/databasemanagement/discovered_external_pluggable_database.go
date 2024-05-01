// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveredExternalPluggableDatabase The details of an external Pluggable Database (PDB) discovered in an external DB system discovery run.
type DiscoveredExternalPluggableDatabase struct {

	// The identifier of the discovered DB system component.
	ComponentId *string `mandatory:"true" json:"componentId"`

	// The user-friendly name for the discovered DB system component. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the discovered DB system component.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier of the parent Container Database (CDB).
	ContainerDatabaseId *string `mandatory:"true" json:"containerDatabaseId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the existing OCI resource matching the discovered DB system component.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Indicates whether the DB system component should be provisioned as an OCI resource or not.
	IsSelectedForMonitoring *bool `mandatory:"false" json:"isSelectedForMonitoring"`

	// The list of associated components.
	AssociatedComponents []AssociatedComponent `mandatory:"false" json:"associatedComponents"`

	// The unique identifier of the PDB.
	Guid *string `mandatory:"false" json:"guid"`

	Connector ExternalDbSystemDiscoveryConnector `mandatory:"false" json:"connector"`

	// The state of the discovered DB system component.
	Status DiscoveredExternalDbSystemComponentStatusEnum `mandatory:"false" json:"status,omitempty"`
}

// GetComponentId returns ComponentId
func (m DiscoveredExternalPluggableDatabase) GetComponentId() *string {
	return m.ComponentId
}

// GetDisplayName returns DisplayName
func (m DiscoveredExternalPluggableDatabase) GetDisplayName() *string {
	return m.DisplayName
}

// GetComponentName returns ComponentName
func (m DiscoveredExternalPluggableDatabase) GetComponentName() *string {
	return m.ComponentName
}

// GetResourceId returns ResourceId
func (m DiscoveredExternalPluggableDatabase) GetResourceId() *string {
	return m.ResourceId
}

// GetIsSelectedForMonitoring returns IsSelectedForMonitoring
func (m DiscoveredExternalPluggableDatabase) GetIsSelectedForMonitoring() *bool {
	return m.IsSelectedForMonitoring
}

// GetStatus returns Status
func (m DiscoveredExternalPluggableDatabase) GetStatus() DiscoveredExternalDbSystemComponentStatusEnum {
	return m.Status
}

// GetAssociatedComponents returns AssociatedComponents
func (m DiscoveredExternalPluggableDatabase) GetAssociatedComponents() []AssociatedComponent {
	return m.AssociatedComponents
}

func (m DiscoveredExternalPluggableDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveredExternalPluggableDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDiscoveredExternalDbSystemComponentStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDiscoveredExternalDbSystemComponentStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DiscoveredExternalPluggableDatabase) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDiscoveredExternalPluggableDatabase DiscoveredExternalPluggableDatabase
	s := struct {
		DiscriminatorParam string `json:"componentType"`
		MarshalTypeDiscoveredExternalPluggableDatabase
	}{
		"PLUGGABLE_DATABASE",
		(MarshalTypeDiscoveredExternalPluggableDatabase)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DiscoveredExternalPluggableDatabase) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ResourceId              *string                                       `json:"resourceId"`
		IsSelectedForMonitoring *bool                                         `json:"isSelectedForMonitoring"`
		Status                  DiscoveredExternalDbSystemComponentStatusEnum `json:"status"`
		AssociatedComponents    []AssociatedComponent                         `json:"associatedComponents"`
		Guid                    *string                                       `json:"guid"`
		Connector               externaldbsystemdiscoveryconnector            `json:"connector"`
		ComponentId             *string                                       `json:"componentId"`
		DisplayName             *string                                       `json:"displayName"`
		ComponentName           *string                                       `json:"componentName"`
		CompartmentId           *string                                       `json:"compartmentId"`
		ContainerDatabaseId     *string                                       `json:"containerDatabaseId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ResourceId = model.ResourceId

	m.IsSelectedForMonitoring = model.IsSelectedForMonitoring

	m.Status = model.Status

	m.AssociatedComponents = make([]AssociatedComponent, len(model.AssociatedComponents))
	copy(m.AssociatedComponents, model.AssociatedComponents)
	m.Guid = model.Guid

	nn, e = model.Connector.UnmarshalPolymorphicJSON(model.Connector.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Connector = nn.(ExternalDbSystemDiscoveryConnector)
	} else {
		m.Connector = nil
	}

	m.ComponentId = model.ComponentId

	m.DisplayName = model.DisplayName

	m.ComponentName = model.ComponentName

	m.CompartmentId = model.CompartmentId

	m.ContainerDatabaseId = model.ContainerDatabaseId

	return
}
