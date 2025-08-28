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

// CloudListener The details of a cloud listener.
type CloudListener struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud listener.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the cloud listener. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the cloud listener.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the listener is a part of.
	CloudDbSystemId *string `mandatory:"true" json:"cloudDbSystemId"`

	// The current lifecycle state of the cloud listener.
	LifecycleState CloudListenerLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the cloud listener was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the cloud listener was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) in DBaas service.
	DbaasId *string `mandatory:"false" json:"dbaasId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud connector.
	CloudConnectorId *string `mandatory:"false" json:"cloudConnectorId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB node.
	CloudDbNodeId *string `mandatory:"false" json:"cloudDbNodeId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB home.
	CloudDbHomeId *string `mandatory:"false" json:"cloudDbHomeId"`

	// The listener alias.
	ListenerAlias *string `mandatory:"false" json:"listenerAlias"`

	// The type of listener.
	ListenerType CloudListenerListenerTypeEnum `mandatory:"false" json:"listenerType,omitempty"`

	// The additional details of the cloud listener defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The location of the listener configuration file listener.ora.
	ListenerOraLocation *string `mandatory:"false" json:"listenerOraLocation"`

	// The Oracle home location of the listener.
	OracleHome *string `mandatory:"false" json:"oracleHome"`

	// The name of the host on which the cloud listener is running.
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
	Endpoints []CloudListenerEndpoint `mandatory:"false" json:"endpoints"`

	// The list of databases that are serviced by the listener.
	ServicedDatabases []CloudListenerServicedDatabase `mandatory:"false" json:"servicedDatabases"`

	// The list of ASMs that are serviced by the listener.
	ServicedAsms []CloudServicedAsm `mandatory:"false" json:"servicedAsms"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CloudListener) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudListener) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudListenerLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCloudListenerLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCloudListenerListenerTypeEnum(string(m.ListenerType)); !ok && m.ListenerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListenerType: %s. Supported values are: %s.", m.ListenerType, strings.Join(GetCloudListenerListenerTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CloudListener) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DbaasId             *string                           `json:"dbaasId"`
		CloudConnectorId    *string                           `json:"cloudConnectorId"`
		CloudDbNodeId       *string                           `json:"cloudDbNodeId"`
		CloudDbHomeId       *string                           `json:"cloudDbHomeId"`
		ListenerAlias       *string                           `json:"listenerAlias"`
		ListenerType        CloudListenerListenerTypeEnum     `json:"listenerType"`
		AdditionalDetails   map[string]string                 `json:"additionalDetails"`
		LifecycleDetails    *string                           `json:"lifecycleDetails"`
		ListenerOraLocation *string                           `json:"listenerOraLocation"`
		OracleHome          *string                           `json:"oracleHome"`
		HostName            *string                           `json:"hostName"`
		AdrHomeDirectory    *string                           `json:"adrHomeDirectory"`
		LogDirectory        *string                           `json:"logDirectory"`
		TraceDirectory      *string                           `json:"traceDirectory"`
		Version             *string                           `json:"version"`
		Endpoints           []cloudlistenerendpoint           `json:"endpoints"`
		ServicedDatabases   []CloudListenerServicedDatabase   `json:"servicedDatabases"`
		ServicedAsms        []CloudServicedAsm                `json:"servicedAsms"`
		FreeformTags        map[string]string                 `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{} `json:"definedTags"`
		SystemTags          map[string]map[string]interface{} `json:"systemTags"`
		Id                  *string                           `json:"id"`
		DisplayName         *string                           `json:"displayName"`
		ComponentName       *string                           `json:"componentName"`
		CompartmentId       *string                           `json:"compartmentId"`
		CloudDbSystemId     *string                           `json:"cloudDbSystemId"`
		LifecycleState      CloudListenerLifecycleStateEnum   `json:"lifecycleState"`
		TimeCreated         *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated         *common.SDKTime                   `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DbaasId = model.DbaasId

	m.CloudConnectorId = model.CloudConnectorId

	m.CloudDbNodeId = model.CloudDbNodeId

	m.CloudDbHomeId = model.CloudDbHomeId

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

	m.Endpoints = make([]CloudListenerEndpoint, len(model.Endpoints))
	for i, n := range model.Endpoints {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Endpoints[i] = nn.(CloudListenerEndpoint)
		} else {
			m.Endpoints[i] = nil
		}
	}
	m.ServicedDatabases = make([]CloudListenerServicedDatabase, len(model.ServicedDatabases))
	copy(m.ServicedDatabases, model.ServicedDatabases)
	m.ServicedAsms = make([]CloudServicedAsm, len(model.ServicedAsms))
	copy(m.ServicedAsms, model.ServicedAsms)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.ComponentName = model.ComponentName

	m.CompartmentId = model.CompartmentId

	m.CloudDbSystemId = model.CloudDbSystemId

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}

// CloudListenerListenerTypeEnum Enum with underlying type: string
type CloudListenerListenerTypeEnum string

// Set of constants representing the allowable values for CloudListenerListenerTypeEnum
const (
	CloudListenerListenerTypeAsm   CloudListenerListenerTypeEnum = "ASM"
	CloudListenerListenerTypeLocal CloudListenerListenerTypeEnum = "LOCAL"
	CloudListenerListenerTypeScan  CloudListenerListenerTypeEnum = "SCAN"
)

var mappingCloudListenerListenerTypeEnum = map[string]CloudListenerListenerTypeEnum{
	"ASM":   CloudListenerListenerTypeAsm,
	"LOCAL": CloudListenerListenerTypeLocal,
	"SCAN":  CloudListenerListenerTypeScan,
}

var mappingCloudListenerListenerTypeEnumLowerCase = map[string]CloudListenerListenerTypeEnum{
	"asm":   CloudListenerListenerTypeAsm,
	"local": CloudListenerListenerTypeLocal,
	"scan":  CloudListenerListenerTypeScan,
}

// GetCloudListenerListenerTypeEnumValues Enumerates the set of values for CloudListenerListenerTypeEnum
func GetCloudListenerListenerTypeEnumValues() []CloudListenerListenerTypeEnum {
	values := make([]CloudListenerListenerTypeEnum, 0)
	for _, v := range mappingCloudListenerListenerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudListenerListenerTypeEnumStringValues Enumerates the set of values in String for CloudListenerListenerTypeEnum
func GetCloudListenerListenerTypeEnumStringValues() []string {
	return []string{
		"ASM",
		"LOCAL",
		"SCAN",
	}
}

// GetMappingCloudListenerListenerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudListenerListenerTypeEnum(val string) (CloudListenerListenerTypeEnum, bool) {
	enum, ok := mappingCloudListenerListenerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CloudListenerLifecycleStateEnum Enum with underlying type: string
type CloudListenerLifecycleStateEnum string

// Set of constants representing the allowable values for CloudListenerLifecycleStateEnum
const (
	CloudListenerLifecycleStateCreating     CloudListenerLifecycleStateEnum = "CREATING"
	CloudListenerLifecycleStateNotConnected CloudListenerLifecycleStateEnum = "NOT_CONNECTED"
	CloudListenerLifecycleStateActive       CloudListenerLifecycleStateEnum = "ACTIVE"
	CloudListenerLifecycleStateInactive     CloudListenerLifecycleStateEnum = "INACTIVE"
	CloudListenerLifecycleStateUpdating     CloudListenerLifecycleStateEnum = "UPDATING"
	CloudListenerLifecycleStateDeleting     CloudListenerLifecycleStateEnum = "DELETING"
	CloudListenerLifecycleStateDeleted      CloudListenerLifecycleStateEnum = "DELETED"
	CloudListenerLifecycleStateFailed       CloudListenerLifecycleStateEnum = "FAILED"
)

var mappingCloudListenerLifecycleStateEnum = map[string]CloudListenerLifecycleStateEnum{
	"CREATING":      CloudListenerLifecycleStateCreating,
	"NOT_CONNECTED": CloudListenerLifecycleStateNotConnected,
	"ACTIVE":        CloudListenerLifecycleStateActive,
	"INACTIVE":      CloudListenerLifecycleStateInactive,
	"UPDATING":      CloudListenerLifecycleStateUpdating,
	"DELETING":      CloudListenerLifecycleStateDeleting,
	"DELETED":       CloudListenerLifecycleStateDeleted,
	"FAILED":        CloudListenerLifecycleStateFailed,
}

var mappingCloudListenerLifecycleStateEnumLowerCase = map[string]CloudListenerLifecycleStateEnum{
	"creating":      CloudListenerLifecycleStateCreating,
	"not_connected": CloudListenerLifecycleStateNotConnected,
	"active":        CloudListenerLifecycleStateActive,
	"inactive":      CloudListenerLifecycleStateInactive,
	"updating":      CloudListenerLifecycleStateUpdating,
	"deleting":      CloudListenerLifecycleStateDeleting,
	"deleted":       CloudListenerLifecycleStateDeleted,
	"failed":        CloudListenerLifecycleStateFailed,
}

// GetCloudListenerLifecycleStateEnumValues Enumerates the set of values for CloudListenerLifecycleStateEnum
func GetCloudListenerLifecycleStateEnumValues() []CloudListenerLifecycleStateEnum {
	values := make([]CloudListenerLifecycleStateEnum, 0)
	for _, v := range mappingCloudListenerLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudListenerLifecycleStateEnumStringValues Enumerates the set of values in String for CloudListenerLifecycleStateEnum
func GetCloudListenerLifecycleStateEnumStringValues() []string {
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

// GetMappingCloudListenerLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudListenerLifecycleStateEnum(val string) (CloudListenerLifecycleStateEnum, bool) {
	enum, ok := mappingCloudListenerLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
