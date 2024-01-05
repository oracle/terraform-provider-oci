// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveredExternalDbSystemComponent The details of an external DB system component.
type DiscoveredExternalDbSystemComponent interface {

	// The identifier of the discovered DB system component.
	GetComponentId() *string

	// The user-friendly name for the discovered DB system component. The name does not have to be unique.
	GetDisplayName() *string

	// The name of the discovered DB system component.
	GetComponentName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the existing OCI resource matching the discovered DB system component.
	GetResourceId() *string

	// Indicates whether the DB system component should be provisioned as an OCI resource or not.
	GetIsSelectedForMonitoring() *bool

	// The state of the discovered DB system component.
	GetStatus() DiscoveredExternalDbSystemComponentStatusEnum

	// The list of associated components.
	GetAssociatedComponents() []AssociatedComponent
}

type discoveredexternaldbsystemcomponent struct {
	JsonData                []byte
	ResourceId              *string                                       `mandatory:"false" json:"resourceId"`
	IsSelectedForMonitoring *bool                                         `mandatory:"false" json:"isSelectedForMonitoring"`
	Status                  DiscoveredExternalDbSystemComponentStatusEnum `mandatory:"false" json:"status,omitempty"`
	AssociatedComponents    []AssociatedComponent                         `mandatory:"false" json:"associatedComponents"`
	ComponentId             *string                                       `mandatory:"true" json:"componentId"`
	DisplayName             *string                                       `mandatory:"true" json:"displayName"`
	ComponentName           *string                                       `mandatory:"true" json:"componentName"`
	ComponentType           string                                        `json:"componentType"`
}

// UnmarshalJSON unmarshals json
func (m *discoveredexternaldbsystemcomponent) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdiscoveredexternaldbsystemcomponent discoveredexternaldbsystemcomponent
	s := struct {
		Model Unmarshalerdiscoveredexternaldbsystemcomponent
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ComponentId = s.Model.ComponentId
	m.DisplayName = s.Model.DisplayName
	m.ComponentName = s.Model.ComponentName
	m.ResourceId = s.Model.ResourceId
	m.IsSelectedForMonitoring = s.Model.IsSelectedForMonitoring
	m.Status = s.Model.Status
	m.AssociatedComponents = s.Model.AssociatedComponents
	m.ComponentType = s.Model.ComponentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *discoveredexternaldbsystemcomponent) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ComponentType {
	case "CLUSTER":
		mm := DiscoveredExternalCluster{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE_HOME":
		mm := DiscoveredExternalDbHome{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE":
		mm := DiscoveredExternalDatabase{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PLUGGABLE_DATABASE":
		mm := DiscoveredExternalPluggableDatabase{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLUSTER_INSTANCE":
		mm := DiscoveredExternalClusterInstance{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LISTENER":
		mm := DiscoveredExternalListener{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE_NODE":
		mm := DiscoveredExternalDbNode{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ASM":
		mm := DiscoveredExternalAsm{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ASM_INSTANCE":
		mm := DiscoveredExternalAsmInstance{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DiscoveredExternalDbSystemComponent: %s.", m.ComponentType)
		return *m, nil
	}
}

// GetResourceId returns ResourceId
func (m discoveredexternaldbsystemcomponent) GetResourceId() *string {
	return m.ResourceId
}

// GetIsSelectedForMonitoring returns IsSelectedForMonitoring
func (m discoveredexternaldbsystemcomponent) GetIsSelectedForMonitoring() *bool {
	return m.IsSelectedForMonitoring
}

// GetStatus returns Status
func (m discoveredexternaldbsystemcomponent) GetStatus() DiscoveredExternalDbSystemComponentStatusEnum {
	return m.Status
}

// GetAssociatedComponents returns AssociatedComponents
func (m discoveredexternaldbsystemcomponent) GetAssociatedComponents() []AssociatedComponent {
	return m.AssociatedComponents
}

// GetComponentId returns ComponentId
func (m discoveredexternaldbsystemcomponent) GetComponentId() *string {
	return m.ComponentId
}

// GetDisplayName returns DisplayName
func (m discoveredexternaldbsystemcomponent) GetDisplayName() *string {
	return m.DisplayName
}

// GetComponentName returns ComponentName
func (m discoveredexternaldbsystemcomponent) GetComponentName() *string {
	return m.ComponentName
}

func (m discoveredexternaldbsystemcomponent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m discoveredexternaldbsystemcomponent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDiscoveredExternalDbSystemComponentStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDiscoveredExternalDbSystemComponentStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiscoveredExternalDbSystemComponentStatusEnum Enum with underlying type: string
type DiscoveredExternalDbSystemComponentStatusEnum string

// Set of constants representing the allowable values for DiscoveredExternalDbSystemComponentStatusEnum
const (
	DiscoveredExternalDbSystemComponentStatusNew               DiscoveredExternalDbSystemComponentStatusEnum = "NEW"
	DiscoveredExternalDbSystemComponentStatusExisting          DiscoveredExternalDbSystemComponentStatusEnum = "EXISTING"
	DiscoveredExternalDbSystemComponentStatusMarkedForDeletion DiscoveredExternalDbSystemComponentStatusEnum = "MARKED_FOR_DELETION"
	DiscoveredExternalDbSystemComponentStatusUnknown           DiscoveredExternalDbSystemComponentStatusEnum = "UNKNOWN"
)

var mappingDiscoveredExternalDbSystemComponentStatusEnum = map[string]DiscoveredExternalDbSystemComponentStatusEnum{
	"NEW":                 DiscoveredExternalDbSystemComponentStatusNew,
	"EXISTING":            DiscoveredExternalDbSystemComponentStatusExisting,
	"MARKED_FOR_DELETION": DiscoveredExternalDbSystemComponentStatusMarkedForDeletion,
	"UNKNOWN":             DiscoveredExternalDbSystemComponentStatusUnknown,
}

var mappingDiscoveredExternalDbSystemComponentStatusEnumLowerCase = map[string]DiscoveredExternalDbSystemComponentStatusEnum{
	"new":                 DiscoveredExternalDbSystemComponentStatusNew,
	"existing":            DiscoveredExternalDbSystemComponentStatusExisting,
	"marked_for_deletion": DiscoveredExternalDbSystemComponentStatusMarkedForDeletion,
	"unknown":             DiscoveredExternalDbSystemComponentStatusUnknown,
}

// GetDiscoveredExternalDbSystemComponentStatusEnumValues Enumerates the set of values for DiscoveredExternalDbSystemComponentStatusEnum
func GetDiscoveredExternalDbSystemComponentStatusEnumValues() []DiscoveredExternalDbSystemComponentStatusEnum {
	values := make([]DiscoveredExternalDbSystemComponentStatusEnum, 0)
	for _, v := range mappingDiscoveredExternalDbSystemComponentStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveredExternalDbSystemComponentStatusEnumStringValues Enumerates the set of values in String for DiscoveredExternalDbSystemComponentStatusEnum
func GetDiscoveredExternalDbSystemComponentStatusEnumStringValues() []string {
	return []string{
		"NEW",
		"EXISTING",
		"MARKED_FOR_DELETION",
		"UNKNOWN",
	}
}

// GetMappingDiscoveredExternalDbSystemComponentStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveredExternalDbSystemComponentStatusEnum(val string) (DiscoveredExternalDbSystemComponentStatusEnum, bool) {
	enum, ok := mappingDiscoveredExternalDbSystemComponentStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
