// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// NetworkAddressList IP addresses that can be used between different WebAppFirewallPolicies.
type NetworkAddressList interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the NetworkAddressList.
	GetId() *string

	// NetworkAddressList display name, can be renamed.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The time the NetworkAddressList was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The current state of the NetworkAddressList.
	GetLifecycleState() NetworkAddressListLifecycleStateEnum

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// The time the NetworkAddressList was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in FAILED state.
	GetLifecycleDetails() *string
}

type networkaddresslist struct {
	JsonData         []byte
	Id               *string                              `mandatory:"true" json:"id"`
	DisplayName      *string                              `mandatory:"true" json:"displayName"`
	CompartmentId    *string                              `mandatory:"true" json:"compartmentId"`
	TimeCreated      *common.SDKTime                      `mandatory:"true" json:"timeCreated"`
	LifecycleState   NetworkAddressListLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	FreeformTags     map[string]string                    `mandatory:"true" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{}    `mandatory:"true" json:"definedTags"`
	SystemTags       map[string]map[string]interface{}    `mandatory:"true" json:"systemTags"`
	TimeUpdated      *common.SDKTime                      `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails *string                              `mandatory:"false" json:"lifecycleDetails"`
	Type             string                               `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *networkaddresslist) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalernetworkaddresslist networkaddresslist
	s := struct {
		Model Unmarshalernetworkaddresslist
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *networkaddresslist) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ADDRESSES":
		mm := NetworkAddressListAddresses{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VCN_ADDRESSES":
		mm := NetworkAddressListVcnAddresses{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m networkaddresslist) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m networkaddresslist) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m networkaddresslist) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m networkaddresslist) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetLifecycleState returns LifecycleState
func (m networkaddresslist) GetLifecycleState() NetworkAddressListLifecycleStateEnum {
	return m.LifecycleState
}

//GetFreeformTags returns FreeformTags
func (m networkaddresslist) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m networkaddresslist) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m networkaddresslist) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

//GetTimeUpdated returns TimeUpdated
func (m networkaddresslist) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleDetails returns LifecycleDetails
func (m networkaddresslist) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m networkaddresslist) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m networkaddresslist) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNetworkAddressListLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNetworkAddressListLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NetworkAddressListLifecycleStateEnum Enum with underlying type: string
type NetworkAddressListLifecycleStateEnum string

// Set of constants representing the allowable values for NetworkAddressListLifecycleStateEnum
const (
	NetworkAddressListLifecycleStateCreating NetworkAddressListLifecycleStateEnum = "CREATING"
	NetworkAddressListLifecycleStateUpdating NetworkAddressListLifecycleStateEnum = "UPDATING"
	NetworkAddressListLifecycleStateActive   NetworkAddressListLifecycleStateEnum = "ACTIVE"
	NetworkAddressListLifecycleStateDeleting NetworkAddressListLifecycleStateEnum = "DELETING"
	NetworkAddressListLifecycleStateDeleted  NetworkAddressListLifecycleStateEnum = "DELETED"
	NetworkAddressListLifecycleStateFailed   NetworkAddressListLifecycleStateEnum = "FAILED"
)

var mappingNetworkAddressListLifecycleStateEnum = map[string]NetworkAddressListLifecycleStateEnum{
	"CREATING": NetworkAddressListLifecycleStateCreating,
	"UPDATING": NetworkAddressListLifecycleStateUpdating,
	"ACTIVE":   NetworkAddressListLifecycleStateActive,
	"DELETING": NetworkAddressListLifecycleStateDeleting,
	"DELETED":  NetworkAddressListLifecycleStateDeleted,
	"FAILED":   NetworkAddressListLifecycleStateFailed,
}

// GetNetworkAddressListLifecycleStateEnumValues Enumerates the set of values for NetworkAddressListLifecycleStateEnum
func GetNetworkAddressListLifecycleStateEnumValues() []NetworkAddressListLifecycleStateEnum {
	values := make([]NetworkAddressListLifecycleStateEnum, 0)
	for _, v := range mappingNetworkAddressListLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkAddressListLifecycleStateEnumStringValues Enumerates the set of values in String for NetworkAddressListLifecycleStateEnum
func GetNetworkAddressListLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingNetworkAddressListLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkAddressListLifecycleStateEnum(val string) (NetworkAddressListLifecycleStateEnum, bool) {
	mappingNetworkAddressListLifecycleStateEnumIgnoreCase := make(map[string]NetworkAddressListLifecycleStateEnum)
	for k, v := range mappingNetworkAddressListLifecycleStateEnum {
		mappingNetworkAddressListLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingNetworkAddressListLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// NetworkAddressListTypeEnum Enum with underlying type: string
type NetworkAddressListTypeEnum string

// Set of constants representing the allowable values for NetworkAddressListTypeEnum
const (
	NetworkAddressListTypeAddresses    NetworkAddressListTypeEnum = "ADDRESSES"
	NetworkAddressListTypeVcnAddresses NetworkAddressListTypeEnum = "VCN_ADDRESSES"
)

var mappingNetworkAddressListTypeEnum = map[string]NetworkAddressListTypeEnum{
	"ADDRESSES":     NetworkAddressListTypeAddresses,
	"VCN_ADDRESSES": NetworkAddressListTypeVcnAddresses,
}

// GetNetworkAddressListTypeEnumValues Enumerates the set of values for NetworkAddressListTypeEnum
func GetNetworkAddressListTypeEnumValues() []NetworkAddressListTypeEnum {
	values := make([]NetworkAddressListTypeEnum, 0)
	for _, v := range mappingNetworkAddressListTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkAddressListTypeEnumStringValues Enumerates the set of values in String for NetworkAddressListTypeEnum
func GetNetworkAddressListTypeEnumStringValues() []string {
	return []string{
		"ADDRESSES",
		"VCN_ADDRESSES",
	}
}

// GetMappingNetworkAddressListTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkAddressListTypeEnum(val string) (NetworkAddressListTypeEnum, bool) {
	mappingNetworkAddressListTypeEnumIgnoreCase := make(map[string]NetworkAddressListTypeEnum)
	for k, v := range mappingNetworkAddressListTypeEnum {
		mappingNetworkAddressListTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingNetworkAddressListTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
