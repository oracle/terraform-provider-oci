// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FlexTunnel Flex tunnel will set up a network tunnel between your SD-WAN appliance and DRG, allowing you to establish BGP and advertise routes.
type FlexTunnel struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the flex tunnel.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the flex tunnel.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the flex tunnel was created, in the format defined
	// by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the flex tunnel.
	LifecycleState FlexTunnelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the drg for loopback attachment.
	DrgId *string `mandatory:"true" json:"drgId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the transport attachment.
	TransportAttachmentId *string `mandatory:"true" json:"transportAttachmentId"`

	FlexTunnelConfiguration FlexTunnelConfiguration `mandatory:"true" json:"flexTunnelConfiguration"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m FlexTunnel) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FlexTunnel) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFlexTunnelLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFlexTunnelLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *FlexTunnel) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName             *string                           `json:"displayName"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		Id                      *string                           `json:"id"`
		CompartmentId           *string                           `json:"compartmentId"`
		TimeCreated             *common.SDKTime                   `json:"timeCreated"`
		LifecycleState          FlexTunnelLifecycleStateEnum      `json:"lifecycleState"`
		DrgId                   *string                           `json:"drgId"`
		TransportAttachmentId   *string                           `json:"transportAttachmentId"`
		FlexTunnelConfiguration flextunnelconfiguration           `json:"flexTunnelConfiguration"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.DrgId = model.DrgId

	m.TransportAttachmentId = model.TransportAttachmentId

	nn, e = model.FlexTunnelConfiguration.UnmarshalPolymorphicJSON(model.FlexTunnelConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FlexTunnelConfiguration = nn.(FlexTunnelConfiguration)
	} else {
		m.FlexTunnelConfiguration = nil
	}

	return
}

// FlexTunnelLifecycleStateEnum Enum with underlying type: string
type FlexTunnelLifecycleStateEnum string

// Set of constants representing the allowable values for FlexTunnelLifecycleStateEnum
const (
	FlexTunnelLifecycleStateCreating FlexTunnelLifecycleStateEnum = "CREATING"
	FlexTunnelLifecycleStateActive   FlexTunnelLifecycleStateEnum = "ACTIVE"
	FlexTunnelLifecycleStateUpdating FlexTunnelLifecycleStateEnum = "UPDATING"
	FlexTunnelLifecycleStateDeleting FlexTunnelLifecycleStateEnum = "DELETING"
	FlexTunnelLifecycleStateDeleted  FlexTunnelLifecycleStateEnum = "DELETED"
	FlexTunnelLifecycleStateFailed   FlexTunnelLifecycleStateEnum = "FAILED"
)

var mappingFlexTunnelLifecycleStateEnum = map[string]FlexTunnelLifecycleStateEnum{
	"CREATING": FlexTunnelLifecycleStateCreating,
	"ACTIVE":   FlexTunnelLifecycleStateActive,
	"UPDATING": FlexTunnelLifecycleStateUpdating,
	"DELETING": FlexTunnelLifecycleStateDeleting,
	"DELETED":  FlexTunnelLifecycleStateDeleted,
	"FAILED":   FlexTunnelLifecycleStateFailed,
}

var mappingFlexTunnelLifecycleStateEnumLowerCase = map[string]FlexTunnelLifecycleStateEnum{
	"creating": FlexTunnelLifecycleStateCreating,
	"active":   FlexTunnelLifecycleStateActive,
	"updating": FlexTunnelLifecycleStateUpdating,
	"deleting": FlexTunnelLifecycleStateDeleting,
	"deleted":  FlexTunnelLifecycleStateDeleted,
	"failed":   FlexTunnelLifecycleStateFailed,
}

// GetFlexTunnelLifecycleStateEnumValues Enumerates the set of values for FlexTunnelLifecycleStateEnum
func GetFlexTunnelLifecycleStateEnumValues() []FlexTunnelLifecycleStateEnum {
	values := make([]FlexTunnelLifecycleStateEnum, 0)
	for _, v := range mappingFlexTunnelLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFlexTunnelLifecycleStateEnumStringValues Enumerates the set of values in String for FlexTunnelLifecycleStateEnum
func GetFlexTunnelLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingFlexTunnelLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFlexTunnelLifecycleStateEnum(val string) (FlexTunnelLifecycleStateEnum, bool) {
	enum, ok := mappingFlexTunnelLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
