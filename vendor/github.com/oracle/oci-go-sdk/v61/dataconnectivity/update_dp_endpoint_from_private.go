// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// UpdateDpEndpointFromPrivate The details to update a private endpoint.
type UpdateDpEndpointFromPrivate struct {

	// Generated key that can be used in API calls to identify endpoint. On scenarios where reference to the endpoint is needed, a value can be passed in create.
	Key *string `mandatory:"true" json:"key"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// The endpoint ID provided by control plane.
	DcmsEndpointId *string `mandatory:"true" json:"dcmsEndpointId"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// User-defined description for the endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// List of data assets which belongs to this endpoint
	DataAssets []DataAsset `mandatory:"false" json:"dataAssets"`

	// The ocid of private endpoint resource.
	PeId *string `mandatory:"false" json:"peId"`

	// The compartmentId of private endpoint resource.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The IP address of dns proxy.
	DnsProxyIp *string `mandatory:"false" json:"dnsProxyIp"`

	// The ocid of private endpoint resource.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// Array of dns zones to be use during private endpoint resolution.
	DnsZones []string `mandatory:"false" json:"dnsZones"`

	// Specifies the private endpoint state.
	State UpdateDpEndpointFromPrivateStateEnum `mandatory:"false" json:"state,omitempty"`
}

//GetKey returns Key
func (m UpdateDpEndpointFromPrivate) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m UpdateDpEndpointFromPrivate) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m UpdateDpEndpointFromPrivate) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m UpdateDpEndpointFromPrivate) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m UpdateDpEndpointFromPrivate) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m UpdateDpEndpointFromPrivate) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetObjectVersion returns ObjectVersion
func (m UpdateDpEndpointFromPrivate) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetIdentifier returns Identifier
func (m UpdateDpEndpointFromPrivate) GetIdentifier() *string {
	return m.Identifier
}

//GetDataAssets returns DataAssets
func (m UpdateDpEndpointFromPrivate) GetDataAssets() []DataAsset {
	return m.DataAssets
}

func (m UpdateDpEndpointFromPrivate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDpEndpointFromPrivate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateDpEndpointFromPrivateStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetUpdateDpEndpointFromPrivateStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateDpEndpointFromPrivate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDpEndpointFromPrivate UpdateDpEndpointFromPrivate
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUpdateDpEndpointFromPrivate
	}{
		"PRIVATE_END_POINT",
		(MarshalTypeUpdateDpEndpointFromPrivate)(m),
	}

	return json.Marshal(&s)
}

// UpdateDpEndpointFromPrivateStateEnum Enum with underlying type: string
type UpdateDpEndpointFromPrivateStateEnum string

// Set of constants representing the allowable values for UpdateDpEndpointFromPrivateStateEnum
const (
	UpdateDpEndpointFromPrivateStateActive   UpdateDpEndpointFromPrivateStateEnum = "ACTIVE"
	UpdateDpEndpointFromPrivateStateInactive UpdateDpEndpointFromPrivateStateEnum = "INACTIVE"
)

var mappingUpdateDpEndpointFromPrivateStateEnum = map[string]UpdateDpEndpointFromPrivateStateEnum{
	"ACTIVE":   UpdateDpEndpointFromPrivateStateActive,
	"INACTIVE": UpdateDpEndpointFromPrivateStateInactive,
}

var mappingUpdateDpEndpointFromPrivateStateEnumLowerCase = map[string]UpdateDpEndpointFromPrivateStateEnum{
	"active":   UpdateDpEndpointFromPrivateStateActive,
	"inactive": UpdateDpEndpointFromPrivateStateInactive,
}

// GetUpdateDpEndpointFromPrivateStateEnumValues Enumerates the set of values for UpdateDpEndpointFromPrivateStateEnum
func GetUpdateDpEndpointFromPrivateStateEnumValues() []UpdateDpEndpointFromPrivateStateEnum {
	values := make([]UpdateDpEndpointFromPrivateStateEnum, 0)
	for _, v := range mappingUpdateDpEndpointFromPrivateStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDpEndpointFromPrivateStateEnumStringValues Enumerates the set of values in String for UpdateDpEndpointFromPrivateStateEnum
func GetUpdateDpEndpointFromPrivateStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingUpdateDpEndpointFromPrivateStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDpEndpointFromPrivateStateEnum(val string) (UpdateDpEndpointFromPrivateStateEnum, bool) {
	enum, ok := mappingUpdateDpEndpointFromPrivateStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
