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

// CreateNetworkAddressListDetails The information about new NetworkAddressList.
type CreateNetworkAddressListDetails interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

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

type createnetworkaddresslistdetails struct {
	JsonData      []byte
	CompartmentId *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName   *string                           `mandatory:"false" json:"displayName"`
	FreeformTags  map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags    map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Type          string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createnetworkaddresslistdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatenetworkaddresslistdetails createnetworkaddresslistdetails
	s := struct {
		Model Unmarshalercreatenetworkaddresslistdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createnetworkaddresslistdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "VCN_ADDRESSES":
		mm := CreateNetworkAddressListVcnAddressesDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ADDRESSES":
		mm := CreateNetworkAddressListAddressesDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetCompartmentId returns CompartmentId
func (m createnetworkaddresslistdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m createnetworkaddresslistdetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m createnetworkaddresslistdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createnetworkaddresslistdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m createnetworkaddresslistdetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m createnetworkaddresslistdetails) String() string {
	return common.PointerString(m)
}

// CreateNetworkAddressListDetailsTypeEnum Enum with underlying type: string
type CreateNetworkAddressListDetailsTypeEnum string

// Set of constants representing the allowable values for CreateNetworkAddressListDetailsTypeEnum
const (
	CreateNetworkAddressListDetailsTypeAddresses    CreateNetworkAddressListDetailsTypeEnum = "ADDRESSES"
	CreateNetworkAddressListDetailsTypeVcnAddresses CreateNetworkAddressListDetailsTypeEnum = "VCN_ADDRESSES"
)

var mappingCreateNetworkAddressListDetailsType = map[string]CreateNetworkAddressListDetailsTypeEnum{
	"ADDRESSES":     CreateNetworkAddressListDetailsTypeAddresses,
	"VCN_ADDRESSES": CreateNetworkAddressListDetailsTypeVcnAddresses,
}

// GetCreateNetworkAddressListDetailsTypeEnumValues Enumerates the set of values for CreateNetworkAddressListDetailsTypeEnum
func GetCreateNetworkAddressListDetailsTypeEnumValues() []CreateNetworkAddressListDetailsTypeEnum {
	values := make([]CreateNetworkAddressListDetailsTypeEnum, 0)
	for _, v := range mappingCreateNetworkAddressListDetailsType {
		values = append(values, v)
	}
	return values
}
