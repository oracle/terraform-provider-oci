// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// DpEndpoint The endpoint for a data asset.
type DpEndpoint interface {

	// Generated key that can be used in API calls to identify the endpoint. In scenarios where reference to the endpoint is required, a value can be passed in create.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// User-defined description of the endpoint.
	GetDescription() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore, and numbers. It should begin with an upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	// The list of data assets that belong to the endpoint.
	GetDataAssets() []DataAsset
}

type dpendpoint struct {
	JsonData      []byte
	Key           *string          `mandatory:"false" json:"key"`
	ModelVersion  *string          `mandatory:"false" json:"modelVersion"`
	ParentRef     *ParentReference `mandatory:"false" json:"parentRef"`
	Name          *string          `mandatory:"false" json:"name"`
	Description   *string          `mandatory:"false" json:"description"`
	ObjectVersion *int             `mandatory:"false" json:"objectVersion"`
	ObjectStatus  *int             `mandatory:"false" json:"objectStatus"`
	Identifier    *string          `mandatory:"false" json:"identifier"`
	DataAssets    []DataAsset      `mandatory:"false" json:"dataAssets"`
	ModelType     string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *dpendpoint) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdpendpoint dpendpoint
	s := struct {
		Model Unmarshalerdpendpoint
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ObjectVersion = s.Model.ObjectVersion
	m.ObjectStatus = s.Model.ObjectStatus
	m.Identifier = s.Model.Identifier
	m.DataAssets = s.Model.DataAssets
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dpendpoint) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "PRIVATE_END_POINT":
		mm := DpEndpointFromPrivate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PUBLIC_END_POINT":
		mm := DpEndpointFromPublic{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DpEndpoint: %s.", m.ModelType)
		return *m, nil
	}
}

//GetKey returns Key
func (m dpendpoint) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m dpendpoint) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m dpendpoint) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m dpendpoint) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m dpendpoint) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m dpendpoint) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m dpendpoint) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m dpendpoint) GetIdentifier() *string {
	return m.Identifier
}

//GetDataAssets returns DataAssets
func (m dpendpoint) GetDataAssets() []DataAsset {
	return m.DataAssets
}

func (m dpendpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dpendpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DpEndpointModelTypeEnum Enum with underlying type: string
type DpEndpointModelTypeEnum string

// Set of constants representing the allowable values for DpEndpointModelTypeEnum
const (
	DpEndpointModelTypePrivateEndPoint DpEndpointModelTypeEnum = "PRIVATE_END_POINT"
	DpEndpointModelTypePublicEndPoint  DpEndpointModelTypeEnum = "PUBLIC_END_POINT"
)

var mappingDpEndpointModelTypeEnum = map[string]DpEndpointModelTypeEnum{
	"PRIVATE_END_POINT": DpEndpointModelTypePrivateEndPoint,
	"PUBLIC_END_POINT":  DpEndpointModelTypePublicEndPoint,
}

var mappingDpEndpointModelTypeEnumLowerCase = map[string]DpEndpointModelTypeEnum{
	"private_end_point": DpEndpointModelTypePrivateEndPoint,
	"public_end_point":  DpEndpointModelTypePublicEndPoint,
}

// GetDpEndpointModelTypeEnumValues Enumerates the set of values for DpEndpointModelTypeEnum
func GetDpEndpointModelTypeEnumValues() []DpEndpointModelTypeEnum {
	values := make([]DpEndpointModelTypeEnum, 0)
	for _, v := range mappingDpEndpointModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDpEndpointModelTypeEnumStringValues Enumerates the set of values in String for DpEndpointModelTypeEnum
func GetDpEndpointModelTypeEnumStringValues() []string {
	return []string{
		"PRIVATE_END_POINT",
		"PUBLIC_END_POINT",
	}
}

// GetMappingDpEndpointModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDpEndpointModelTypeEnum(val string) (DpEndpointModelTypeEnum, bool) {
	enum, ok := mappingDpEndpointModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
