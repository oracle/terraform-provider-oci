// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NetworkAddressListSummary Summary of NetworkAddressList.
type NetworkAddressListSummary interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the NetworkAddressList.
	GetId() *string

	// NetworkAddressList display name, can be renamed.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The time the NetworkAddressList was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The current state of the NetworkAddress List.
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

type networkaddresslistsummary struct {
	JsonData         []byte
	TimeUpdated      *common.SDKTime                      `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails *string                              `mandatory:"false" json:"lifecycleDetails"`
	Id               *string                              `mandatory:"true" json:"id"`
	DisplayName      *string                              `mandatory:"true" json:"displayName"`
	CompartmentId    *string                              `mandatory:"true" json:"compartmentId"`
	TimeCreated      *common.SDKTime                      `mandatory:"true" json:"timeCreated"`
	LifecycleState   NetworkAddressListLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	FreeformTags     map[string]string                    `mandatory:"true" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{}    `mandatory:"true" json:"definedTags"`
	SystemTags       map[string]map[string]interface{}    `mandatory:"true" json:"systemTags"`
	Type             string                               `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *networkaddresslistsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalernetworkaddresslistsummary networkaddresslistsummary
	s := struct {
		Model Unmarshalernetworkaddresslistsummary
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
func (m *networkaddresslistsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "VCN_ADDRESSES":
		mm := NetworkAddressListVcnAddressesSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ADDRESSES":
		mm := NetworkAddressListAddressesSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for NetworkAddressListSummary: %s.", m.Type)
		return *m, nil
	}
}

// GetTimeUpdated returns TimeUpdated
func (m networkaddresslistsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m networkaddresslistsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetId returns Id
func (m networkaddresslistsummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m networkaddresslistsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m networkaddresslistsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m networkaddresslistsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m networkaddresslistsummary) GetLifecycleState() NetworkAddressListLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m networkaddresslistsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m networkaddresslistsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m networkaddresslistsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m networkaddresslistsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m networkaddresslistsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNetworkAddressListLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNetworkAddressListLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NetworkAddressListSummaryTypeEnum Enum with underlying type: string
type NetworkAddressListSummaryTypeEnum string

// Set of constants representing the allowable values for NetworkAddressListSummaryTypeEnum
const (
	NetworkAddressListSummaryTypeAddresses    NetworkAddressListSummaryTypeEnum = "ADDRESSES"
	NetworkAddressListSummaryTypeVcnAddresses NetworkAddressListSummaryTypeEnum = "VCN_ADDRESSES"
)

var mappingNetworkAddressListSummaryTypeEnum = map[string]NetworkAddressListSummaryTypeEnum{
	"ADDRESSES":     NetworkAddressListSummaryTypeAddresses,
	"VCN_ADDRESSES": NetworkAddressListSummaryTypeVcnAddresses,
}

var mappingNetworkAddressListSummaryTypeEnumLowerCase = map[string]NetworkAddressListSummaryTypeEnum{
	"addresses":     NetworkAddressListSummaryTypeAddresses,
	"vcn_addresses": NetworkAddressListSummaryTypeVcnAddresses,
}

// GetNetworkAddressListSummaryTypeEnumValues Enumerates the set of values for NetworkAddressListSummaryTypeEnum
func GetNetworkAddressListSummaryTypeEnumValues() []NetworkAddressListSummaryTypeEnum {
	values := make([]NetworkAddressListSummaryTypeEnum, 0)
	for _, v := range mappingNetworkAddressListSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkAddressListSummaryTypeEnumStringValues Enumerates the set of values in String for NetworkAddressListSummaryTypeEnum
func GetNetworkAddressListSummaryTypeEnumStringValues() []string {
	return []string{
		"ADDRESSES",
		"VCN_ADDRESSES",
	}
}

// GetMappingNetworkAddressListSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkAddressListSummaryTypeEnum(val string) (NetworkAddressListSummaryTypeEnum, bool) {
	enum, ok := mappingNetworkAddressListSummaryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
