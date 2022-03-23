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

// CreateDpEndpointFromPublic The details to create a public endpoint.
type CreateDpEndpointFromPublic struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"true" json:"identifier"`

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
}

//GetKey returns Key
func (m CreateDpEndpointFromPublic) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m CreateDpEndpointFromPublic) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m CreateDpEndpointFromPublic) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m CreateDpEndpointFromPublic) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m CreateDpEndpointFromPublic) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m CreateDpEndpointFromPublic) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetObjectVersion returns ObjectVersion
func (m CreateDpEndpointFromPublic) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetIdentifier returns Identifier
func (m CreateDpEndpointFromPublic) GetIdentifier() *string {
	return m.Identifier
}

//GetDataAssets returns DataAssets
func (m CreateDpEndpointFromPublic) GetDataAssets() []DataAsset {
	return m.DataAssets
}

func (m CreateDpEndpointFromPublic) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDpEndpointFromPublic) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDpEndpointFromPublic) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDpEndpointFromPublic CreateDpEndpointFromPublic
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCreateDpEndpointFromPublic
	}{
		"PUBLIC_END_POINT",
		(MarshalTypeCreateDpEndpointFromPublic)(m),
	}

	return json.Marshal(&s)
}
