// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveredCloudClusterInstance The details of a cloud cluster instance discovered in a cloud DB system discovery run.
type DiscoveredCloudClusterInstance struct {

	// The identifier of the discovered DB system component.
	ComponentId *string `mandatory:"true" json:"componentId"`

	// The user-friendly name for the discovered DB system component. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the discovered DB system component.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The name of the host on which the cluster instance is running.
	HostName *string `mandatory:"true" json:"hostName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing OCI resource matching the discovered DB system component.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing Dbaas OCI resource matching the discovered DB system component.
	DbaasId *string `mandatory:"false" json:"dbaasId"`

	// Indicates whether the DB system component should be provisioned as an OCI resource or not.
	IsSelectedForMonitoring *bool `mandatory:"false" json:"isSelectedForMonitoring"`

	// The list of associated components.
	AssociatedComponents []AssociatedCloudComponent `mandatory:"false" json:"associatedComponents"`

	// The unique identifier of the Oracle cluster.
	ClusterId *string `mandatory:"false" json:"clusterId"`

	// The Oracle base location of Cluster Ready Services (CRS).
	CrsBaseDirectory *string `mandatory:"false" json:"crsBaseDirectory"`

	// The Automatic Diagnostic Repository (ADR) home directory for the cluster instance.
	AdrHomeDirectory *string `mandatory:"false" json:"adrHomeDirectory"`

	Connector CloudDbSystemDiscoveryConnector `mandatory:"false" json:"connector"`

	// The role of the cluster node.
	NodeRole DiscoveredCloudClusterInstanceNodeRoleEnum `mandatory:"false" json:"nodeRole,omitempty"`

	// The state of the discovered DB system component.
	Status DiscoveredCloudDbSystemComponentStatusEnum `mandatory:"false" json:"status,omitempty"`
}

// GetComponentId returns ComponentId
func (m DiscoveredCloudClusterInstance) GetComponentId() *string {
	return m.ComponentId
}

// GetDisplayName returns DisplayName
func (m DiscoveredCloudClusterInstance) GetDisplayName() *string {
	return m.DisplayName
}

// GetComponentName returns ComponentName
func (m DiscoveredCloudClusterInstance) GetComponentName() *string {
	return m.ComponentName
}

// GetResourceId returns ResourceId
func (m DiscoveredCloudClusterInstance) GetResourceId() *string {
	return m.ResourceId
}

// GetDbaasId returns DbaasId
func (m DiscoveredCloudClusterInstance) GetDbaasId() *string {
	return m.DbaasId
}

// GetIsSelectedForMonitoring returns IsSelectedForMonitoring
func (m DiscoveredCloudClusterInstance) GetIsSelectedForMonitoring() *bool {
	return m.IsSelectedForMonitoring
}

// GetStatus returns Status
func (m DiscoveredCloudClusterInstance) GetStatus() DiscoveredCloudDbSystemComponentStatusEnum {
	return m.Status
}

// GetAssociatedComponents returns AssociatedComponents
func (m DiscoveredCloudClusterInstance) GetAssociatedComponents() []AssociatedCloudComponent {
	return m.AssociatedComponents
}

func (m DiscoveredCloudClusterInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveredCloudClusterInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveredCloudClusterInstanceNodeRoleEnum(string(m.NodeRole)); !ok && m.NodeRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodeRole: %s. Supported values are: %s.", m.NodeRole, strings.Join(GetDiscoveredCloudClusterInstanceNodeRoleEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDiscoveredCloudDbSystemComponentStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDiscoveredCloudDbSystemComponentStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DiscoveredCloudClusterInstance) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDiscoveredCloudClusterInstance DiscoveredCloudClusterInstance
	s := struct {
		DiscriminatorParam string `json:"componentType"`
		MarshalTypeDiscoveredCloudClusterInstance
	}{
		"CLUSTER_INSTANCE",
		(MarshalTypeDiscoveredCloudClusterInstance)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DiscoveredCloudClusterInstance) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ResourceId              *string                                    `json:"resourceId"`
		DbaasId                 *string                                    `json:"dbaasId"`
		IsSelectedForMonitoring *bool                                      `json:"isSelectedForMonitoring"`
		Status                  DiscoveredCloudDbSystemComponentStatusEnum `json:"status"`
		AssociatedComponents    []AssociatedCloudComponent                 `json:"associatedComponents"`
		ClusterId               *string                                    `json:"clusterId"`
		NodeRole                DiscoveredCloudClusterInstanceNodeRoleEnum `json:"nodeRole"`
		CrsBaseDirectory        *string                                    `json:"crsBaseDirectory"`
		AdrHomeDirectory        *string                                    `json:"adrHomeDirectory"`
		Connector               clouddbsystemdiscoveryconnector            `json:"connector"`
		ComponentId             *string                                    `json:"componentId"`
		DisplayName             *string                                    `json:"displayName"`
		ComponentName           *string                                    `json:"componentName"`
		HostName                *string                                    `json:"hostName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ResourceId = model.ResourceId

	m.DbaasId = model.DbaasId

	m.IsSelectedForMonitoring = model.IsSelectedForMonitoring

	m.Status = model.Status

	m.AssociatedComponents = make([]AssociatedCloudComponent, len(model.AssociatedComponents))
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
		m.Connector = nn.(CloudDbSystemDiscoveryConnector)
	} else {
		m.Connector = nil
	}

	m.ComponentId = model.ComponentId

	m.DisplayName = model.DisplayName

	m.ComponentName = model.ComponentName

	m.HostName = model.HostName

	return
}

// DiscoveredCloudClusterInstanceNodeRoleEnum Enum with underlying type: string
type DiscoveredCloudClusterInstanceNodeRoleEnum string

// Set of constants representing the allowable values for DiscoveredCloudClusterInstanceNodeRoleEnum
const (
	DiscoveredCloudClusterInstanceNodeRoleHub  DiscoveredCloudClusterInstanceNodeRoleEnum = "HUB"
	DiscoveredCloudClusterInstanceNodeRoleLeaf DiscoveredCloudClusterInstanceNodeRoleEnum = "LEAF"
)

var mappingDiscoveredCloudClusterInstanceNodeRoleEnum = map[string]DiscoveredCloudClusterInstanceNodeRoleEnum{
	"HUB":  DiscoveredCloudClusterInstanceNodeRoleHub,
	"LEAF": DiscoveredCloudClusterInstanceNodeRoleLeaf,
}

var mappingDiscoveredCloudClusterInstanceNodeRoleEnumLowerCase = map[string]DiscoveredCloudClusterInstanceNodeRoleEnum{
	"hub":  DiscoveredCloudClusterInstanceNodeRoleHub,
	"leaf": DiscoveredCloudClusterInstanceNodeRoleLeaf,
}

// GetDiscoveredCloudClusterInstanceNodeRoleEnumValues Enumerates the set of values for DiscoveredCloudClusterInstanceNodeRoleEnum
func GetDiscoveredCloudClusterInstanceNodeRoleEnumValues() []DiscoveredCloudClusterInstanceNodeRoleEnum {
	values := make([]DiscoveredCloudClusterInstanceNodeRoleEnum, 0)
	for _, v := range mappingDiscoveredCloudClusterInstanceNodeRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveredCloudClusterInstanceNodeRoleEnumStringValues Enumerates the set of values in String for DiscoveredCloudClusterInstanceNodeRoleEnum
func GetDiscoveredCloudClusterInstanceNodeRoleEnumStringValues() []string {
	return []string{
		"HUB",
		"LEAF",
	}
}

// GetMappingDiscoveredCloudClusterInstanceNodeRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveredCloudClusterInstanceNodeRoleEnum(val string) (DiscoveredCloudClusterInstanceNodeRoleEnum, bool) {
	enum, ok := mappingDiscoveredCloudClusterInstanceNodeRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
