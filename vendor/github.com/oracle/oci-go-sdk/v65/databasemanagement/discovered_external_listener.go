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

// DiscoveredExternalListener The details of an Oracle listener discovered in an external DB system discovery run.
type DiscoveredExternalListener struct {

	// The identifier of the discovered DB system component.
	ComponentId *string `mandatory:"true" json:"componentId"`

	// The user-friendly name for the discovered DB system component. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the discovered DB system component.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the existing OCI resource matching the discovered DB system component.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Indicates whether the DB system component should be provisioned as an OCI resource or not.
	IsSelectedForMonitoring *bool `mandatory:"false" json:"isSelectedForMonitoring"`

	// The list of associated components.
	AssociatedComponents []AssociatedComponent `mandatory:"false" json:"associatedComponents"`

	// The name of the DB node.
	DbNodeName *string `mandatory:"false" json:"dbNodeName"`

	// The Oracle home location of the listener.
	OracleHome *string `mandatory:"false" json:"oracleHome"`

	// The listener alias.
	ListenerAlias *string `mandatory:"false" json:"listenerAlias"`

	// The directory that stores tracing and logging incidents when Automatic Diagnostic Repository (ADR) is enabled.
	AdrHomeDirectory *string `mandatory:"false" json:"adrHomeDirectory"`

	// The destination directory of the listener log file.
	LogDirectory *string `mandatory:"false" json:"logDirectory"`

	// The destination directory of the listener trace file.
	TraceDirectory *string `mandatory:"false" json:"traceDirectory"`

	// The listener version.
	Version *string `mandatory:"false" json:"version"`

	// The name of the host on which the external listener is running.
	HostName *string `mandatory:"false" json:"hostName"`

	// The list of protocol addresses the listener is configured to listen on.
	Endpoints []ExternalListenerEndpoint `mandatory:"false" json:"endpoints"`

	Connector ExternalDbSystemDiscoveryConnector `mandatory:"false" json:"connector"`

	// The type of listener.
	ListenerType DiscoveredExternalListenerListenerTypeEnum `mandatory:"false" json:"listenerType,omitempty"`

	// The state of the discovered DB system component.
	Status DiscoveredExternalDbSystemComponentStatusEnum `mandatory:"false" json:"status,omitempty"`
}

// GetComponentId returns ComponentId
func (m DiscoveredExternalListener) GetComponentId() *string {
	return m.ComponentId
}

// GetDisplayName returns DisplayName
func (m DiscoveredExternalListener) GetDisplayName() *string {
	return m.DisplayName
}

// GetComponentName returns ComponentName
func (m DiscoveredExternalListener) GetComponentName() *string {
	return m.ComponentName
}

// GetResourceId returns ResourceId
func (m DiscoveredExternalListener) GetResourceId() *string {
	return m.ResourceId
}

// GetIsSelectedForMonitoring returns IsSelectedForMonitoring
func (m DiscoveredExternalListener) GetIsSelectedForMonitoring() *bool {
	return m.IsSelectedForMonitoring
}

// GetStatus returns Status
func (m DiscoveredExternalListener) GetStatus() DiscoveredExternalDbSystemComponentStatusEnum {
	return m.Status
}

// GetAssociatedComponents returns AssociatedComponents
func (m DiscoveredExternalListener) GetAssociatedComponents() []AssociatedComponent {
	return m.AssociatedComponents
}

func (m DiscoveredExternalListener) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveredExternalListener) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveredExternalListenerListenerTypeEnum(string(m.ListenerType)); !ok && m.ListenerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListenerType: %s. Supported values are: %s.", m.ListenerType, strings.Join(GetDiscoveredExternalListenerListenerTypeEnumStringValues(), ",")))
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
func (m DiscoveredExternalListener) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDiscoveredExternalListener DiscoveredExternalListener
	s := struct {
		DiscriminatorParam string `json:"componentType"`
		MarshalTypeDiscoveredExternalListener
	}{
		"LISTENER",
		(MarshalTypeDiscoveredExternalListener)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DiscoveredExternalListener) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ResourceId              *string                                       `json:"resourceId"`
		IsSelectedForMonitoring *bool                                         `json:"isSelectedForMonitoring"`
		Status                  DiscoveredExternalDbSystemComponentStatusEnum `json:"status"`
		AssociatedComponents    []AssociatedComponent                         `json:"associatedComponents"`
		DbNodeName              *string                                       `json:"dbNodeName"`
		OracleHome              *string                                       `json:"oracleHome"`
		ListenerAlias           *string                                       `json:"listenerAlias"`
		AdrHomeDirectory        *string                                       `json:"adrHomeDirectory"`
		LogDirectory            *string                                       `json:"logDirectory"`
		TraceDirectory          *string                                       `json:"traceDirectory"`
		Version                 *string                                       `json:"version"`
		ListenerType            DiscoveredExternalListenerListenerTypeEnum    `json:"listenerType"`
		HostName                *string                                       `json:"hostName"`
		Endpoints               []externallistenerendpoint                    `json:"endpoints"`
		Connector               externaldbsystemdiscoveryconnector            `json:"connector"`
		ComponentId             *string                                       `json:"componentId"`
		DisplayName             *string                                       `json:"displayName"`
		ComponentName           *string                                       `json:"componentName"`
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
	m.DbNodeName = model.DbNodeName

	m.OracleHome = model.OracleHome

	m.ListenerAlias = model.ListenerAlias

	m.AdrHomeDirectory = model.AdrHomeDirectory

	m.LogDirectory = model.LogDirectory

	m.TraceDirectory = model.TraceDirectory

	m.Version = model.Version

	m.ListenerType = model.ListenerType

	m.HostName = model.HostName

	m.Endpoints = make([]ExternalListenerEndpoint, len(model.Endpoints))
	for i, n := range model.Endpoints {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Endpoints[i] = nn.(ExternalListenerEndpoint)
		} else {
			m.Endpoints[i] = nil
		}
	}
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

	return
}

// DiscoveredExternalListenerListenerTypeEnum Enum with underlying type: string
type DiscoveredExternalListenerListenerTypeEnum string

// Set of constants representing the allowable values for DiscoveredExternalListenerListenerTypeEnum
const (
	DiscoveredExternalListenerListenerTypeAsm   DiscoveredExternalListenerListenerTypeEnum = "ASM"
	DiscoveredExternalListenerListenerTypeLocal DiscoveredExternalListenerListenerTypeEnum = "LOCAL"
	DiscoveredExternalListenerListenerTypeScan  DiscoveredExternalListenerListenerTypeEnum = "SCAN"
)

var mappingDiscoveredExternalListenerListenerTypeEnum = map[string]DiscoveredExternalListenerListenerTypeEnum{
	"ASM":   DiscoveredExternalListenerListenerTypeAsm,
	"LOCAL": DiscoveredExternalListenerListenerTypeLocal,
	"SCAN":  DiscoveredExternalListenerListenerTypeScan,
}

var mappingDiscoveredExternalListenerListenerTypeEnumLowerCase = map[string]DiscoveredExternalListenerListenerTypeEnum{
	"asm":   DiscoveredExternalListenerListenerTypeAsm,
	"local": DiscoveredExternalListenerListenerTypeLocal,
	"scan":  DiscoveredExternalListenerListenerTypeScan,
}

// GetDiscoveredExternalListenerListenerTypeEnumValues Enumerates the set of values for DiscoveredExternalListenerListenerTypeEnum
func GetDiscoveredExternalListenerListenerTypeEnumValues() []DiscoveredExternalListenerListenerTypeEnum {
	values := make([]DiscoveredExternalListenerListenerTypeEnum, 0)
	for _, v := range mappingDiscoveredExternalListenerListenerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveredExternalListenerListenerTypeEnumStringValues Enumerates the set of values in String for DiscoveredExternalListenerListenerTypeEnum
func GetDiscoveredExternalListenerListenerTypeEnumStringValues() []string {
	return []string{
		"ASM",
		"LOCAL",
		"SCAN",
	}
}

// GetMappingDiscoveredExternalListenerListenerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveredExternalListenerListenerTypeEnum(val string) (DiscoveredExternalListenerListenerTypeEnum, bool) {
	enum, ok := mappingDiscoveredExternalListenerListenerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
