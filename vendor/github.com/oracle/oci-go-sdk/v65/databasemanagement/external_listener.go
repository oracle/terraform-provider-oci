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

// ExternalListener The details of an external listener.
type ExternalListener struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external listener.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the external listener. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the external listener.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system that the listener is a part of.
	ExternalDbSystemId *string `mandatory:"true" json:"externalDbSystemId"`

	// The current lifecycle state of the external listener.
	LifecycleState ExternalListenerLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the external listener was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the external listener was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external connector.
	ExternalConnectorId *string `mandatory:"false" json:"externalConnectorId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB node.
	ExternalDbNodeId *string `mandatory:"false" json:"externalDbNodeId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB home.
	ExternalDbHomeId *string `mandatory:"false" json:"externalDbHomeId"`

	// The listener alias.
	ListenerAlias *string `mandatory:"false" json:"listenerAlias"`

	// The type of listener.
	ListenerType ExternalListenerListenerTypeEnum `mandatory:"false" json:"listenerType,omitempty"`

	// The additional details of the external listener defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The location of the listener configuration file listener.ora.
	ListenerOraLocation *string `mandatory:"false" json:"listenerOraLocation"`

	// The Oracle home location of the listener.
	OracleHome *string `mandatory:"false" json:"oracleHome"`

	// The name of the host on which the external listener is running.
	HostName *string `mandatory:"false" json:"hostName"`

	// The directory that stores tracing and logging incidents when Automatic Diagnostic Repository (ADR) is enabled.
	AdrHomeDirectory *string `mandatory:"false" json:"adrHomeDirectory"`

	// The destination directory of the listener log file.
	LogDirectory *string `mandatory:"false" json:"logDirectory"`

	// The destination directory of the listener trace file.
	TraceDirectory *string `mandatory:"false" json:"traceDirectory"`

	// The listener version.
	Version *string `mandatory:"false" json:"version"`

	// The list of protocol addresses the listener is configured to listen on.
	Endpoints []ExternalListenerEndpoint `mandatory:"false" json:"endpoints"`

	// The list of databases that are serviced by the listener.
	ServicedDatabases []ExternalListenerServicedDatabase `mandatory:"false" json:"servicedDatabases"`

	// The list of ASMs that are serviced by the listener.
	ServicedAsms []ExternalServicedAsm `mandatory:"false" json:"servicedAsms"`
}

func (m ExternalListener) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalListener) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalListenerLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalListenerLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExternalListenerListenerTypeEnum(string(m.ListenerType)); !ok && m.ListenerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListenerType: %s. Supported values are: %s.", m.ListenerType, strings.Join(GetExternalListenerListenerTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ExternalListener) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ExternalConnectorId *string                            `json:"externalConnectorId"`
		ExternalDbNodeId    *string                            `json:"externalDbNodeId"`
		ExternalDbHomeId    *string                            `json:"externalDbHomeId"`
		ListenerAlias       *string                            `json:"listenerAlias"`
		ListenerType        ExternalListenerListenerTypeEnum   `json:"listenerType"`
		AdditionalDetails   map[string]string                  `json:"additionalDetails"`
		LifecycleDetails    *string                            `json:"lifecycleDetails"`
		ListenerOraLocation *string                            `json:"listenerOraLocation"`
		OracleHome          *string                            `json:"oracleHome"`
		HostName            *string                            `json:"hostName"`
		AdrHomeDirectory    *string                            `json:"adrHomeDirectory"`
		LogDirectory        *string                            `json:"logDirectory"`
		TraceDirectory      *string                            `json:"traceDirectory"`
		Version             *string                            `json:"version"`
		Endpoints           []externallistenerendpoint         `json:"endpoints"`
		ServicedDatabases   []ExternalListenerServicedDatabase `json:"servicedDatabases"`
		ServicedAsms        []ExternalServicedAsm              `json:"servicedAsms"`
		Id                  *string                            `json:"id"`
		DisplayName         *string                            `json:"displayName"`
		ComponentName       *string                            `json:"componentName"`
		CompartmentId       *string                            `json:"compartmentId"`
		ExternalDbSystemId  *string                            `json:"externalDbSystemId"`
		LifecycleState      ExternalListenerLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated         *common.SDKTime                    `json:"timeCreated"`
		TimeUpdated         *common.SDKTime                    `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ExternalConnectorId = model.ExternalConnectorId

	m.ExternalDbNodeId = model.ExternalDbNodeId

	m.ExternalDbHomeId = model.ExternalDbHomeId

	m.ListenerAlias = model.ListenerAlias

	m.ListenerType = model.ListenerType

	m.AdditionalDetails = model.AdditionalDetails

	m.LifecycleDetails = model.LifecycleDetails

	m.ListenerOraLocation = model.ListenerOraLocation

	m.OracleHome = model.OracleHome

	m.HostName = model.HostName

	m.AdrHomeDirectory = model.AdrHomeDirectory

	m.LogDirectory = model.LogDirectory

	m.TraceDirectory = model.TraceDirectory

	m.Version = model.Version

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
	m.ServicedDatabases = make([]ExternalListenerServicedDatabase, len(model.ServicedDatabases))
	copy(m.ServicedDatabases, model.ServicedDatabases)
	m.ServicedAsms = make([]ExternalServicedAsm, len(model.ServicedAsms))
	copy(m.ServicedAsms, model.ServicedAsms)
	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.ComponentName = model.ComponentName

	m.CompartmentId = model.CompartmentId

	m.ExternalDbSystemId = model.ExternalDbSystemId

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}

// ExternalListenerListenerTypeEnum Enum with underlying type: string
type ExternalListenerListenerTypeEnum string

// Set of constants representing the allowable values for ExternalListenerListenerTypeEnum
const (
	ExternalListenerListenerTypeAsm   ExternalListenerListenerTypeEnum = "ASM"
	ExternalListenerListenerTypeLocal ExternalListenerListenerTypeEnum = "LOCAL"
	ExternalListenerListenerTypeScan  ExternalListenerListenerTypeEnum = "SCAN"
)

var mappingExternalListenerListenerTypeEnum = map[string]ExternalListenerListenerTypeEnum{
	"ASM":   ExternalListenerListenerTypeAsm,
	"LOCAL": ExternalListenerListenerTypeLocal,
	"SCAN":  ExternalListenerListenerTypeScan,
}

var mappingExternalListenerListenerTypeEnumLowerCase = map[string]ExternalListenerListenerTypeEnum{
	"asm":   ExternalListenerListenerTypeAsm,
	"local": ExternalListenerListenerTypeLocal,
	"scan":  ExternalListenerListenerTypeScan,
}

// GetExternalListenerListenerTypeEnumValues Enumerates the set of values for ExternalListenerListenerTypeEnum
func GetExternalListenerListenerTypeEnumValues() []ExternalListenerListenerTypeEnum {
	values := make([]ExternalListenerListenerTypeEnum, 0)
	for _, v := range mappingExternalListenerListenerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalListenerListenerTypeEnumStringValues Enumerates the set of values in String for ExternalListenerListenerTypeEnum
func GetExternalListenerListenerTypeEnumStringValues() []string {
	return []string{
		"ASM",
		"LOCAL",
		"SCAN",
	}
}

// GetMappingExternalListenerListenerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalListenerListenerTypeEnum(val string) (ExternalListenerListenerTypeEnum, bool) {
	enum, ok := mappingExternalListenerListenerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalListenerLifecycleStateEnum Enum with underlying type: string
type ExternalListenerLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalListenerLifecycleStateEnum
const (
	ExternalListenerLifecycleStateCreating     ExternalListenerLifecycleStateEnum = "CREATING"
	ExternalListenerLifecycleStateNotConnected ExternalListenerLifecycleStateEnum = "NOT_CONNECTED"
	ExternalListenerLifecycleStateActive       ExternalListenerLifecycleStateEnum = "ACTIVE"
	ExternalListenerLifecycleStateInactive     ExternalListenerLifecycleStateEnum = "INACTIVE"
	ExternalListenerLifecycleStateUpdating     ExternalListenerLifecycleStateEnum = "UPDATING"
	ExternalListenerLifecycleStateDeleting     ExternalListenerLifecycleStateEnum = "DELETING"
	ExternalListenerLifecycleStateDeleted      ExternalListenerLifecycleStateEnum = "DELETED"
	ExternalListenerLifecycleStateFailed       ExternalListenerLifecycleStateEnum = "FAILED"
)

var mappingExternalListenerLifecycleStateEnum = map[string]ExternalListenerLifecycleStateEnum{
	"CREATING":      ExternalListenerLifecycleStateCreating,
	"NOT_CONNECTED": ExternalListenerLifecycleStateNotConnected,
	"ACTIVE":        ExternalListenerLifecycleStateActive,
	"INACTIVE":      ExternalListenerLifecycleStateInactive,
	"UPDATING":      ExternalListenerLifecycleStateUpdating,
	"DELETING":      ExternalListenerLifecycleStateDeleting,
	"DELETED":       ExternalListenerLifecycleStateDeleted,
	"FAILED":        ExternalListenerLifecycleStateFailed,
}

var mappingExternalListenerLifecycleStateEnumLowerCase = map[string]ExternalListenerLifecycleStateEnum{
	"creating":      ExternalListenerLifecycleStateCreating,
	"not_connected": ExternalListenerLifecycleStateNotConnected,
	"active":        ExternalListenerLifecycleStateActive,
	"inactive":      ExternalListenerLifecycleStateInactive,
	"updating":      ExternalListenerLifecycleStateUpdating,
	"deleting":      ExternalListenerLifecycleStateDeleting,
	"deleted":       ExternalListenerLifecycleStateDeleted,
	"failed":        ExternalListenerLifecycleStateFailed,
}

// GetExternalListenerLifecycleStateEnumValues Enumerates the set of values for ExternalListenerLifecycleStateEnum
func GetExternalListenerLifecycleStateEnumValues() []ExternalListenerLifecycleStateEnum {
	values := make([]ExternalListenerLifecycleStateEnum, 0)
	for _, v := range mappingExternalListenerLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalListenerLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalListenerLifecycleStateEnum
func GetExternalListenerLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"NOT_CONNECTED",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingExternalListenerLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalListenerLifecycleStateEnum(val string) (ExternalListenerLifecycleStateEnum, bool) {
	enum, ok := mappingExternalListenerLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
