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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DpEndpointSummaryFromPrivate The endpoint summary details for a private endpoint.
type DpEndpointSummaryFromPrivate struct {

	// The endpoint ID provided by control plane.
	DcmsEndpointId *string `mandatory:"true" json:"dcmsEndpointId"`

	// Generated key that can be used in API calls to identify endpoint. On scenarios where reference to the endpoint is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// User-defined description for the endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
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
	State DpEndpointSummaryFromPrivateStateEnum `mandatory:"false" json:"state,omitempty"`
}

//GetKey returns Key
func (m DpEndpointSummaryFromPrivate) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m DpEndpointSummaryFromPrivate) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m DpEndpointSummaryFromPrivate) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m DpEndpointSummaryFromPrivate) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m DpEndpointSummaryFromPrivate) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m DpEndpointSummaryFromPrivate) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m DpEndpointSummaryFromPrivate) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m DpEndpointSummaryFromPrivate) GetIdentifier() *string {
	return m.Identifier
}

//GetDataAssets returns DataAssets
func (m DpEndpointSummaryFromPrivate) GetDataAssets() []DataAsset {
	return m.DataAssets
}

func (m DpEndpointSummaryFromPrivate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DpEndpointSummaryFromPrivate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingDpEndpointSummaryFromPrivateStateEnum[string(m.State)]; !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetDpEndpointSummaryFromPrivateStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DpEndpointSummaryFromPrivate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDpEndpointSummaryFromPrivate DpEndpointSummaryFromPrivate
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDpEndpointSummaryFromPrivate
	}{
		"PRIVATE_END_POINT",
		(MarshalTypeDpEndpointSummaryFromPrivate)(m),
	}

	return json.Marshal(&s)
}

// DpEndpointSummaryFromPrivateStateEnum Enum with underlying type: string
type DpEndpointSummaryFromPrivateStateEnum string

// Set of constants representing the allowable values for DpEndpointSummaryFromPrivateStateEnum
const (
	DpEndpointSummaryFromPrivateStateActive   DpEndpointSummaryFromPrivateStateEnum = "ACTIVE"
	DpEndpointSummaryFromPrivateStateInactive DpEndpointSummaryFromPrivateStateEnum = "INACTIVE"
)

var mappingDpEndpointSummaryFromPrivateStateEnum = map[string]DpEndpointSummaryFromPrivateStateEnum{
	"ACTIVE":   DpEndpointSummaryFromPrivateStateActive,
	"INACTIVE": DpEndpointSummaryFromPrivateStateInactive,
}

// GetDpEndpointSummaryFromPrivateStateEnumValues Enumerates the set of values for DpEndpointSummaryFromPrivateStateEnum
func GetDpEndpointSummaryFromPrivateStateEnumValues() []DpEndpointSummaryFromPrivateStateEnum {
	values := make([]DpEndpointSummaryFromPrivateStateEnum, 0)
	for _, v := range mappingDpEndpointSummaryFromPrivateStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDpEndpointSummaryFromPrivateStateEnumStringValues Enumerates the set of values in String for DpEndpointSummaryFromPrivateStateEnum
func GetDpEndpointSummaryFromPrivateStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}
