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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateNetworkAddressListDetails The information to be updated.
type UpdateNetworkAddressListDetails interface {

	// NetworkAddressList display name, can be renamed.
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type updatenetworkaddresslistdetails struct {
	JsonData     []byte
	DisplayName  *string                           `mandatory:"false" json:"displayName"`
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags   map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Type         string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updatenetworkaddresslistdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatenetworkaddresslistdetails updatenetworkaddresslistdetails
	s := struct {
		Model Unmarshalerupdatenetworkaddresslistdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatenetworkaddresslistdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ADDRESSES":
		mm := UpdateNetworkAddressListAddressesDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VCN_ADDRESSES":
		mm := UpdateNetworkAddressListVcnAddressesDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m updatenetworkaddresslistdetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m updatenetworkaddresslistdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m updatenetworkaddresslistdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m updatenetworkaddresslistdetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m updatenetworkaddresslistdetails) String() string {
	return common.PointerString(m)
}

// UpdateNetworkAddressListDetailsTypeEnum Enum with underlying type: string
type UpdateNetworkAddressListDetailsTypeEnum string

// Set of constants representing the allowable values for UpdateNetworkAddressListDetailsTypeEnum
const (
	UpdateNetworkAddressListDetailsTypeAddresses    UpdateNetworkAddressListDetailsTypeEnum = "ADDRESSES"
	UpdateNetworkAddressListDetailsTypeVcnAddresses UpdateNetworkAddressListDetailsTypeEnum = "VCN_ADDRESSES"
)

var mappingUpdateNetworkAddressListDetailsType = map[string]UpdateNetworkAddressListDetailsTypeEnum{
	"ADDRESSES":     UpdateNetworkAddressListDetailsTypeAddresses,
	"VCN_ADDRESSES": UpdateNetworkAddressListDetailsTypeVcnAddresses,
}

// GetUpdateNetworkAddressListDetailsTypeEnumValues Enumerates the set of values for UpdateNetworkAddressListDetailsTypeEnum
func GetUpdateNetworkAddressListDetailsTypeEnumValues() []UpdateNetworkAddressListDetailsTypeEnum {
	values := make([]UpdateNetworkAddressListDetailsTypeEnum, 0)
	for _, v := range mappingUpdateNetworkAddressListDetailsType {
		values = append(values, v)
	}
	return values
}
