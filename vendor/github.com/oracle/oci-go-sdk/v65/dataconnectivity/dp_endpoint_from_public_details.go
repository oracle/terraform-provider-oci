// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DpEndpointFromPublicDetails The endpoint details of a public endpoint.
type DpEndpointFromPublicDetails struct {

	// Generated key that can be used in API calls to identify the endpoint. In scenarios where reference to the endpoint is required, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// User-defined description of the endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with an upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The list of data assets that belong to the endpoint.
	DataAssets []DataAsset `mandatory:"false" json:"dataAssets"`
}

//GetKey returns Key
func (m DpEndpointFromPublicDetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m DpEndpointFromPublicDetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m DpEndpointFromPublicDetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m DpEndpointFromPublicDetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m DpEndpointFromPublicDetails) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m DpEndpointFromPublicDetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m DpEndpointFromPublicDetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m DpEndpointFromPublicDetails) GetIdentifier() *string {
	return m.Identifier
}

//GetDataAssets returns DataAssets
func (m DpEndpointFromPublicDetails) GetDataAssets() []DataAsset {
	return m.DataAssets
}

func (m DpEndpointFromPublicDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DpEndpointFromPublicDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DpEndpointFromPublicDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDpEndpointFromPublicDetails DpEndpointFromPublicDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDpEndpointFromPublicDetails
	}{
		"PUBLIC_END_POINT",
		(MarshalTypeDpEndpointFromPublicDetails)(m),
	}

	return json.Marshal(&s)
}
