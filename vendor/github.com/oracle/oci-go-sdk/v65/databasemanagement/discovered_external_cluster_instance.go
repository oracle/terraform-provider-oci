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

// DiscoveredExternalClusterInstance The details of an external cluster instance discovered in an external DB system discovery run.
type DiscoveredExternalClusterInstance struct {

	// The identifier of the discovered DB system component.
	ComponentId *string `mandatory:"true" json:"componentId"`

	// The user-friendly name for the discovered DB system component. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the discovered DB system component.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The name of the host on which the cluster instance is running.
	HostName *string `mandatory:"true" json:"hostName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the existing OCI resource matching the discovered DB system component.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Indicates whether the DB system component should be provisioned as an OCI resource or not.
	IsSelectedForMonitoring *bool `mandatory:"false" json:"isSelectedForMonitoring"`

	// The list of associated components.
	AssociatedComponents []AssociatedComponent `mandatory:"false" json:"associatedComponents"`

	// The unique identifier of the Oracle cluster.
	ClusterId *string `mandatory:"false" json:"clusterId"`

	// The Oracle base location of Cluster Ready Services (CRS).
	CrsBaseDirectory *string `mandatory:"false" json:"crsBaseDirectory"`

	// The Automatic Diagnostic Repository (ADR) home directory for the cluster instance.
	AdrHomeDirectory *string `mandatory:"false" json:"adrHomeDirectory"`

	Connector ExternalDbSystemDiscoveryConnector `mandatory:"false" json:"connector"`

	// The role of the cluster node.
	NodeRole DiscoveredExternalClusterInstanceNodeRoleEnum `mandatory:"false" json:"nodeRole,omitempty"`

	// The state of the discovered DB system component.
	Status DiscoveredExternalDbSystemComponentStatusEnum `mandatory:"false" json:"status,omitempty"`
}

// GetComponentId returns ComponentId
func (m DiscoveredExternalClusterInstance) GetComponentId() *string {
	return m.ComponentId
}

// GetDisplayName returns DisplayName
func (m DiscoveredExternalClusterInstance) GetDisplayName() *string {
	return m.DisplayName
}

// GetComponentName returns ComponentName
func (m DiscoveredExternalClusterInstance) GetComponentName() *string {
	return m.ComponentName
}

// GetResourceId returns ResourceId
func (m DiscoveredExternalClusterInstance) GetResourceId() *string {
	return m.ResourceId
}

// GetIsSelectedForMonitoring returns IsSelectedForMonitoring
func (m DiscoveredExternalClusterInstance) GetIsSelectedForMonitoring() *bool {
	return m.IsSelectedForMonitoring
}

// GetStatus returns Status
func (m DiscoveredExternalClusterInstance) GetStatus() DiscoveredExternalDbSystemComponentStatusEnum {
	return m.Status
}

// GetAssociatedComponents returns AssociatedComponents
func (m DiscoveredExternalClusterInstance) GetAssociatedComponents() []AssociatedComponent {
	return m.AssociatedComponents
}

func (m DiscoveredExternalClusterInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveredExternalClusterInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveredExternalClusterInstanceNodeRoleEnum(string(m.NodeRole)); !ok && m.NodeRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodeRole: %s. Supported values are: %s.", m.NodeRole, strings.Join(GetDiscoveredExternalClusterInstanceNodeRoleEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDiscoveredExternalDbSystemComponentStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDiscoveredExternalDbSystemComponentStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DiscoveredExternalClusterInstance) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDiscoveredExternalClusterInstance DiscoveredExternalClusterInstance
	s := struct {
		DiscriminatorParam string `json:"componentType"`
		MarshalTypeDiscoveredExternalClusterInstance
	}{
		"CLUSTER_INSTANCE",
		(MarshalTypeDiscoveredExternalClusterInstance)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DiscoveredExternalClusterInstance) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ResourceId              *string                                       `json:"resourceId"`
		IsSelectedForMonitoring *bool                                         `json:"isSelectedForMonitoring"`
		Status                  DiscoveredExternalDbSystemComponentStatusEnum `json:"status"`
		AssociatedComponents    []AssociatedComponent                         `json:"associatedComponents"`
		ClusterId               *string                                       `json:"clusterId"`
		NodeRole                DiscoveredExternalClusterInstanceNodeRoleEnum `json:"nodeRole"`
		CrsBaseDirectory        *string                                       `json:"crsBaseDirectory"`
		AdrHomeDirectory        *string                                       `json:"adrHomeDirectory"`
		Connector               externaldbsystemdiscoveryconnector            `json:"connector"`
		ComponentId             *string                                       `json:"componentId"`
		DisplayName             *string                                       `json:"displayName"`
		ComponentName           *string                                       `json:"componentName"`
		HostName                *string                                       `json:"hostName"`
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
	m.ClusterId = model.ClusterId

	m.NodeRole = model.NodeRole

	m.CrsBaseDirectory = model.CrsBaseDirectory

	m.AdrHomeDirectory = model.AdrHomeDirectory

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

	m.HostName = model.HostName

	return
}

// DiscoveredExternalClusterInstanceNodeRoleEnum Enum with underlying type: string
type DiscoveredExternalClusterInstanceNodeRoleEnum string

// Set of constants representing the allowable values for DiscoveredExternalClusterInstanceNodeRoleEnum
const (
	DiscoveredExternalClusterInstanceNodeRoleHub  DiscoveredExternalClusterInstanceNodeRoleEnum = "HUB"
	DiscoveredExternalClusterInstanceNodeRoleLeaf DiscoveredExternalClusterInstanceNodeRoleEnum = "LEAF"
)

var mappingDiscoveredExternalClusterInstanceNodeRoleEnum = map[string]DiscoveredExternalClusterInstanceNodeRoleEnum{
	"HUB":  DiscoveredExternalClusterInstanceNodeRoleHub,
	"LEAF": DiscoveredExternalClusterInstanceNodeRoleLeaf,
}

var mappingDiscoveredExternalClusterInstanceNodeRoleEnumLowerCase = map[string]DiscoveredExternalClusterInstanceNodeRoleEnum{
	"hub":  DiscoveredExternalClusterInstanceNodeRoleHub,
	"leaf": DiscoveredExternalClusterInstanceNodeRoleLeaf,
}

// GetDiscoveredExternalClusterInstanceNodeRoleEnumValues Enumerates the set of values for DiscoveredExternalClusterInstanceNodeRoleEnum
func GetDiscoveredExternalClusterInstanceNodeRoleEnumValues() []DiscoveredExternalClusterInstanceNodeRoleEnum {
	values := make([]DiscoveredExternalClusterInstanceNodeRoleEnum, 0)
	for _, v := range mappingDiscoveredExternalClusterInstanceNodeRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveredExternalClusterInstanceNodeRoleEnumStringValues Enumerates the set of values in String for DiscoveredExternalClusterInstanceNodeRoleEnum
func GetDiscoveredExternalClusterInstanceNodeRoleEnumStringValues() []string {
	return []string{
		"HUB",
		"LEAF",
	}
}

// GetMappingDiscoveredExternalClusterInstanceNodeRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveredExternalClusterInstanceNodeRoleEnum(val string) (DiscoveredExternalClusterInstanceNodeRoleEnum, bool) {
	enum, ok := mappingDiscoveredExternalClusterInstanceNodeRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
