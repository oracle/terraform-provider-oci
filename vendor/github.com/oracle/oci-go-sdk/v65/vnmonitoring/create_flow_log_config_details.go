// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// CreateFlowLogConfigDetails The representation of CreateFlowLogConfigDetails
type CreateFlowLogConfigDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the
	// flow log configuration.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type or types of flow logs to store. `ALL` includes records for both accepted traffic and
	// rejected traffic.
	FlowLogType CreateFlowLogConfigDetailsFlowLogTypeEnum `mandatory:"true" json:"flowLogType"`

	Destination FlowLogDestination `mandatory:"true" json:"destination"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateFlowLogConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateFlowLogConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateFlowLogConfigDetailsFlowLogTypeEnum(string(m.FlowLogType)); !ok && m.FlowLogType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FlowLogType: %s. Supported values are: %s.", m.FlowLogType, strings.Join(GetCreateFlowLogConfigDetailsFlowLogTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateFlowLogConfigDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags   map[string]map[string]interface{}         `json:"definedTags"`
		DisplayName   *string                                   `json:"displayName"`
		FreeformTags  map[string]string                         `json:"freeformTags"`
		CompartmentId *string                                   `json:"compartmentId"`
		FlowLogType   CreateFlowLogConfigDetailsFlowLogTypeEnum `json:"flowLogType"`
		Destination   flowlogdestination                        `json:"destination"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.CompartmentId = model.CompartmentId

	m.FlowLogType = model.FlowLogType

	nn, e = model.Destination.UnmarshalPolymorphicJSON(model.Destination.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Destination = nn.(FlowLogDestination)
	} else {
		m.Destination = nil
	}

	return
}

// CreateFlowLogConfigDetailsFlowLogTypeEnum Enum with underlying type: string
type CreateFlowLogConfigDetailsFlowLogTypeEnum string

// Set of constants representing the allowable values for CreateFlowLogConfigDetailsFlowLogTypeEnum
const (
	CreateFlowLogConfigDetailsFlowLogTypeAll CreateFlowLogConfigDetailsFlowLogTypeEnum = "ALL"
)

var mappingCreateFlowLogConfigDetailsFlowLogTypeEnum = map[string]CreateFlowLogConfigDetailsFlowLogTypeEnum{
	"ALL": CreateFlowLogConfigDetailsFlowLogTypeAll,
}

var mappingCreateFlowLogConfigDetailsFlowLogTypeEnumLowerCase = map[string]CreateFlowLogConfigDetailsFlowLogTypeEnum{
	"all": CreateFlowLogConfigDetailsFlowLogTypeAll,
}

// GetCreateFlowLogConfigDetailsFlowLogTypeEnumValues Enumerates the set of values for CreateFlowLogConfigDetailsFlowLogTypeEnum
func GetCreateFlowLogConfigDetailsFlowLogTypeEnumValues() []CreateFlowLogConfigDetailsFlowLogTypeEnum {
	values := make([]CreateFlowLogConfigDetailsFlowLogTypeEnum, 0)
	for _, v := range mappingCreateFlowLogConfigDetailsFlowLogTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateFlowLogConfigDetailsFlowLogTypeEnumStringValues Enumerates the set of values in String for CreateFlowLogConfigDetailsFlowLogTypeEnum
func GetCreateFlowLogConfigDetailsFlowLogTypeEnumStringValues() []string {
	return []string{
		"ALL",
	}
}

// GetMappingCreateFlowLogConfigDetailsFlowLogTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateFlowLogConfigDetailsFlowLogTypeEnum(val string) (CreateFlowLogConfigDetailsFlowLogTypeEnum, bool) {
	enum, ok := mappingCreateFlowLogConfigDetailsFlowLogTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
