// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PathAnalyzerTest Defines the details saved in a `PathAnalyzerTest` resource. These configuration details are used to
// run a Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) analysis.
type PathAnalyzerTest struct {

	// A unique identifier established when the resource is created. The identifier can't be changed later.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `PathAnalyzerTest` resource's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The IP protocol to use for the `PathAnalyzerTest` resource.
	Protocol *int `mandatory:"true" json:"protocol"`

	SourceEndpoint Endpoint `mandatory:"true" json:"sourceEndpoint"`

	DestinationEndpoint Endpoint `mandatory:"true" json:"destinationEndpoint"`

	QueryOptions *QueryOptions `mandatory:"true" json:"queryOptions"`

	// The date and time the `PathAnalyzerTest` resource was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the `PathAnalyzerTest` resource was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the `PathAnalyzerTest` resource.
	LifecycleState PathAnalyzerTestLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	ProtocolParameters ProtocolParameters `mandatory:"false" json:"protocolParameters"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m PathAnalyzerTest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PathAnalyzerTest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPathAnalyzerTestLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPathAnalyzerTestLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *PathAnalyzerTest) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ProtocolParameters  protocolparameters                 `json:"protocolParameters"`
		FreeformTags        map[string]string                  `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{}  `json:"definedTags"`
		SystemTags          map[string]map[string]interface{}  `json:"systemTags"`
		Id                  *string                            `json:"id"`
		DisplayName         *string                            `json:"displayName"`
		CompartmentId       *string                            `json:"compartmentId"`
		Protocol            *int                               `json:"protocol"`
		SourceEndpoint      endpoint                           `json:"sourceEndpoint"`
		DestinationEndpoint endpoint                           `json:"destinationEndpoint"`
		QueryOptions        *QueryOptions                      `json:"queryOptions"`
		TimeCreated         *common.SDKTime                    `json:"timeCreated"`
		TimeUpdated         *common.SDKTime                    `json:"timeUpdated"`
		LifecycleState      PathAnalyzerTestLifecycleStateEnum `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ProtocolParameters.UnmarshalPolymorphicJSON(model.ProtocolParameters.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ProtocolParameters = nn.(ProtocolParameters)
	} else {
		m.ProtocolParameters = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.Protocol = model.Protocol

	nn, e = model.SourceEndpoint.UnmarshalPolymorphicJSON(model.SourceEndpoint.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceEndpoint = nn.(Endpoint)
	} else {
		m.SourceEndpoint = nil
	}

	nn, e = model.DestinationEndpoint.UnmarshalPolymorphicJSON(model.DestinationEndpoint.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DestinationEndpoint = nn.(Endpoint)
	} else {
		m.DestinationEndpoint = nil
	}

	m.QueryOptions = model.QueryOptions

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	return
}

// PathAnalyzerTestLifecycleStateEnum Enum with underlying type: string
type PathAnalyzerTestLifecycleStateEnum string

// Set of constants representing the allowable values for PathAnalyzerTestLifecycleStateEnum
const (
	PathAnalyzerTestLifecycleStateActive  PathAnalyzerTestLifecycleStateEnum = "ACTIVE"
	PathAnalyzerTestLifecycleStateDeleted PathAnalyzerTestLifecycleStateEnum = "DELETED"
)

var mappingPathAnalyzerTestLifecycleStateEnum = map[string]PathAnalyzerTestLifecycleStateEnum{
	"ACTIVE":  PathAnalyzerTestLifecycleStateActive,
	"DELETED": PathAnalyzerTestLifecycleStateDeleted,
}

var mappingPathAnalyzerTestLifecycleStateEnumLowerCase = map[string]PathAnalyzerTestLifecycleStateEnum{
	"active":  PathAnalyzerTestLifecycleStateActive,
	"deleted": PathAnalyzerTestLifecycleStateDeleted,
}

// GetPathAnalyzerTestLifecycleStateEnumValues Enumerates the set of values for PathAnalyzerTestLifecycleStateEnum
func GetPathAnalyzerTestLifecycleStateEnumValues() []PathAnalyzerTestLifecycleStateEnum {
	values := make([]PathAnalyzerTestLifecycleStateEnum, 0)
	for _, v := range mappingPathAnalyzerTestLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPathAnalyzerTestLifecycleStateEnumStringValues Enumerates the set of values in String for PathAnalyzerTestLifecycleStateEnum
func GetPathAnalyzerTestLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingPathAnalyzerTestLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPathAnalyzerTestLifecycleStateEnum(val string) (PathAnalyzerTestLifecycleStateEnum, bool) {
	enum, ok := mappingPathAnalyzerTestLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
