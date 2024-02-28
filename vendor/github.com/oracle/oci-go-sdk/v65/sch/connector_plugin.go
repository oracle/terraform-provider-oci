// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.cloud.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConnectorPlugin A service source or service target used to create a connector.
// Example connector plugins include the Queue source and the Notifications target.
// For more information about flows defined by connectors, see
// Overview of Connector Hub (https://docs.cloud.oracle.com/iaas/Content/connector-hub/overview.htm).
// For configuration instructions, see
// Creating a Connector (https://docs.cloud.oracle.com/iaas/Content/connector-hub/create-service-connector.htm).
type ConnectorPlugin interface {

	// The service to be called by the connector plugin.
	// Example: `QueueSource`
	GetName() *string

	// The date and time when this plugin became available.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2023-09-09T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	// The current state of the service connector.
	GetLifecycleState() ConnectorPluginLifecycleStateEnum

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// The estimated throughput range (LOW, MEDIUM, HIGH).
	GetEstimatedThroughput() EstimatedThroughputEnum

	// Gets the specified connector plugin configuration information in OpenAPI specification format.
	GetSchema() *string
}

type connectorplugin struct {
	JsonData            []byte
	EstimatedThroughput EstimatedThroughputEnum           `mandatory:"false" json:"estimatedThroughput,omitempty"`
	Schema              *string                           `mandatory:"false" json:"schema"`
	Name                *string                           `mandatory:"true" json:"name"`
	TimeCreated         *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState      ConnectorPluginLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	DisplayName         *string                           `mandatory:"true" json:"displayName"`
	Kind                string                            `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *connectorplugin) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconnectorplugin connectorplugin
	s := struct {
		Model Unmarshalerconnectorplugin
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.DisplayName = s.Model.DisplayName
	m.EstimatedThroughput = s.Model.EstimatedThroughput
	m.Schema = s.Model.Schema
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *connectorplugin) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "TARGET":
		mm := TargetConnectorPlugin{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SOURCE":
		mm := SourceConnectorPlugin{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ConnectorPlugin: %s.", m.Kind)
		return *m, nil
	}
}

// GetEstimatedThroughput returns EstimatedThroughput
func (m connectorplugin) GetEstimatedThroughput() EstimatedThroughputEnum {
	return m.EstimatedThroughput
}

// GetSchema returns Schema
func (m connectorplugin) GetSchema() *string {
	return m.Schema
}

// GetName returns Name
func (m connectorplugin) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m connectorplugin) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m connectorplugin) GetLifecycleState() ConnectorPluginLifecycleStateEnum {
	return m.LifecycleState
}

// GetDisplayName returns DisplayName
func (m connectorplugin) GetDisplayName() *string {
	return m.DisplayName
}

func (m connectorplugin) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m connectorplugin) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConnectorPluginLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectorPluginLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingEstimatedThroughputEnum(string(m.EstimatedThroughput)); !ok && m.EstimatedThroughput != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EstimatedThroughput: %s. Supported values are: %s.", m.EstimatedThroughput, strings.Join(GetEstimatedThroughputEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectorPluginKindEnum Enum with underlying type: string
type ConnectorPluginKindEnum string

// Set of constants representing the allowable values for ConnectorPluginKindEnum
const (
	ConnectorPluginKindSource ConnectorPluginKindEnum = "SOURCE"
	ConnectorPluginKindTarget ConnectorPluginKindEnum = "TARGET"
)

var mappingConnectorPluginKindEnum = map[string]ConnectorPluginKindEnum{
	"SOURCE": ConnectorPluginKindSource,
	"TARGET": ConnectorPluginKindTarget,
}

var mappingConnectorPluginKindEnumLowerCase = map[string]ConnectorPluginKindEnum{
	"source": ConnectorPluginKindSource,
	"target": ConnectorPluginKindTarget,
}

// GetConnectorPluginKindEnumValues Enumerates the set of values for ConnectorPluginKindEnum
func GetConnectorPluginKindEnumValues() []ConnectorPluginKindEnum {
	values := make([]ConnectorPluginKindEnum, 0)
	for _, v := range mappingConnectorPluginKindEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectorPluginKindEnumStringValues Enumerates the set of values in String for ConnectorPluginKindEnum
func GetConnectorPluginKindEnumStringValues() []string {
	return []string{
		"SOURCE",
		"TARGET",
	}
}

// GetMappingConnectorPluginKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectorPluginKindEnum(val string) (ConnectorPluginKindEnum, bool) {
	enum, ok := mappingConnectorPluginKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
