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
	"github.com/oracle/oci-go-sdk/v63/common"
	"strings"
)

// CreateDpEndpointFromPrivate The details to create a private endpoint.
type CreateDpEndpointFromPrivate struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The endpoint ID provided by control plane.
	DcmsEndpointId *string `mandatory:"true" json:"dcmsEndpointId"`

	// Generated key that can be used in API calls to identify endpoint. On scenarios where reference to the endpoint is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// User-defined description for the endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

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
	State CreateDpEndpointFromPrivateStateEnum `mandatory:"false" json:"state,omitempty"`
}

//GetKey returns Key
func (m CreateDpEndpointFromPrivate) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m CreateDpEndpointFromPrivate) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m CreateDpEndpointFromPrivate) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m CreateDpEndpointFromPrivate) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m CreateDpEndpointFromPrivate) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m CreateDpEndpointFromPrivate) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetObjectVersion returns ObjectVersion
func (m CreateDpEndpointFromPrivate) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetIdentifier returns Identifier
func (m CreateDpEndpointFromPrivate) GetIdentifier() *string {
	return m.Identifier
}

//GetDataAssets returns DataAssets
func (m CreateDpEndpointFromPrivate) GetDataAssets() []DataAsset {
	return m.DataAssets
}

func (m CreateDpEndpointFromPrivate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDpEndpointFromPrivate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateDpEndpointFromPrivateStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetCreateDpEndpointFromPrivateStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDpEndpointFromPrivate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDpEndpointFromPrivate CreateDpEndpointFromPrivate
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCreateDpEndpointFromPrivate
	}{
		"PRIVATE_END_POINT",
		(MarshalTypeCreateDpEndpointFromPrivate)(m),
	}

	return json.Marshal(&s)
}

// CreateDpEndpointFromPrivateStateEnum Enum with underlying type: string
type CreateDpEndpointFromPrivateStateEnum string

// Set of constants representing the allowable values for CreateDpEndpointFromPrivateStateEnum
const (
	CreateDpEndpointFromPrivateStateActive   CreateDpEndpointFromPrivateStateEnum = "ACTIVE"
	CreateDpEndpointFromPrivateStateInactive CreateDpEndpointFromPrivateStateEnum = "INACTIVE"
)

var mappingCreateDpEndpointFromPrivateStateEnum = map[string]CreateDpEndpointFromPrivateStateEnum{
	"ACTIVE":   CreateDpEndpointFromPrivateStateActive,
	"INACTIVE": CreateDpEndpointFromPrivateStateInactive,
}

var mappingCreateDpEndpointFromPrivateStateEnumLowerCase = map[string]CreateDpEndpointFromPrivateStateEnum{
	"active":   CreateDpEndpointFromPrivateStateActive,
	"inactive": CreateDpEndpointFromPrivateStateInactive,
}

// GetCreateDpEndpointFromPrivateStateEnumValues Enumerates the set of values for CreateDpEndpointFromPrivateStateEnum
func GetCreateDpEndpointFromPrivateStateEnumValues() []CreateDpEndpointFromPrivateStateEnum {
	values := make([]CreateDpEndpointFromPrivateStateEnum, 0)
	for _, v := range mappingCreateDpEndpointFromPrivateStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDpEndpointFromPrivateStateEnumStringValues Enumerates the set of values in String for CreateDpEndpointFromPrivateStateEnum
func GetCreateDpEndpointFromPrivateStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingCreateDpEndpointFromPrivateStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDpEndpointFromPrivateStateEnum(val string) (CreateDpEndpointFromPrivateStateEnum, bool) {
	enum, ok := mappingCreateDpEndpointFromPrivateStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
