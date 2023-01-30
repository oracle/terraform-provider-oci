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

// CreateDpEndpointDetails Properties used in the create operations of the endpoint.
type CreateDpEndpointDetails interface {

	// Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// Value can only contain upper case letters, underscore, and numbers. It should begin with an upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	// Generated key that can be used in API calls to identify the endpoint. In scenarios where reference to the endpoint is required, a value can be passed in create.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// User-defined description of the endpoint.
	GetDescription() *string

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// The list of data assets that belong to the endpoint.
	GetDataAssets() []DataAsset
}

type createdpendpointdetails struct {
	JsonData      []byte
	Name          *string          `mandatory:"true" json:"name"`
	Identifier    *string          `mandatory:"true" json:"identifier"`
	Key           *string          `mandatory:"false" json:"key"`
	ModelVersion  *string          `mandatory:"false" json:"modelVersion"`
	ParentRef     *ParentReference `mandatory:"false" json:"parentRef"`
	Description   *string          `mandatory:"false" json:"description"`
	ObjectStatus  *int             `mandatory:"false" json:"objectStatus"`
	ObjectVersion *int             `mandatory:"false" json:"objectVersion"`
	DataAssets    []DataAsset      `mandatory:"false" json:"dataAssets"`
	ModelType     string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *createdpendpointdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedpendpointdetails createdpendpointdetails
	s := struct {
		Model Unmarshalercreatedpendpointdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Identifier = s.Model.Identifier
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Description = s.Model.Description
	m.ObjectStatus = s.Model.ObjectStatus
	m.ObjectVersion = s.Model.ObjectVersion
	m.DataAssets = s.Model.DataAssets
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdpendpointdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "PUBLIC_END_POINT":
		mm := CreateDpEndpointFromPublic{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRIVATE_END_POINT":
		mm := CreateDpEndpointFromPrivate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m createdpendpointdetails) GetName() *string {
	return m.Name
}

//GetIdentifier returns Identifier
func (m createdpendpointdetails) GetIdentifier() *string {
	return m.Identifier
}

//GetKey returns Key
func (m createdpendpointdetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m createdpendpointdetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m createdpendpointdetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetDescription returns Description
func (m createdpendpointdetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m createdpendpointdetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetObjectVersion returns ObjectVersion
func (m createdpendpointdetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetDataAssets returns DataAssets
func (m createdpendpointdetails) GetDataAssets() []DataAsset {
	return m.DataAssets
}

func (m createdpendpointdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdpendpointdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDpEndpointDetailsModelTypeEnum Enum with underlying type: string
type CreateDpEndpointDetailsModelTypeEnum string

// Set of constants representing the allowable values for CreateDpEndpointDetailsModelTypeEnum
const (
	CreateDpEndpointDetailsModelTypePrivateEndPoint CreateDpEndpointDetailsModelTypeEnum = "PRIVATE_END_POINT"
	CreateDpEndpointDetailsModelTypePublicEndPoint  CreateDpEndpointDetailsModelTypeEnum = "PUBLIC_END_POINT"
)

var mappingCreateDpEndpointDetailsModelTypeEnum = map[string]CreateDpEndpointDetailsModelTypeEnum{
	"PRIVATE_END_POINT": CreateDpEndpointDetailsModelTypePrivateEndPoint,
	"PUBLIC_END_POINT":  CreateDpEndpointDetailsModelTypePublicEndPoint,
}

var mappingCreateDpEndpointDetailsModelTypeEnumLowerCase = map[string]CreateDpEndpointDetailsModelTypeEnum{
	"private_end_point": CreateDpEndpointDetailsModelTypePrivateEndPoint,
	"public_end_point":  CreateDpEndpointDetailsModelTypePublicEndPoint,
}

// GetCreateDpEndpointDetailsModelTypeEnumValues Enumerates the set of values for CreateDpEndpointDetailsModelTypeEnum
func GetCreateDpEndpointDetailsModelTypeEnumValues() []CreateDpEndpointDetailsModelTypeEnum {
	values := make([]CreateDpEndpointDetailsModelTypeEnum, 0)
	for _, v := range mappingCreateDpEndpointDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDpEndpointDetailsModelTypeEnumStringValues Enumerates the set of values in String for CreateDpEndpointDetailsModelTypeEnum
func GetCreateDpEndpointDetailsModelTypeEnumStringValues() []string {
	return []string{
		"PRIVATE_END_POINT",
		"PUBLIC_END_POINT",
	}
}

// GetMappingCreateDpEndpointDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDpEndpointDetailsModelTypeEnum(val string) (CreateDpEndpointDetailsModelTypeEnum, bool) {
	enum, ok := mappingCreateDpEndpointDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
