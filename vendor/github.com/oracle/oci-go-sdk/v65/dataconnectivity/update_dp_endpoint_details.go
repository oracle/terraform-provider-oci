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

// UpdateDpEndpointDetails Properties used in the update operations of the endpoint.
type UpdateDpEndpointDetails interface {

	// Generated key that can be used in API calls to identify the endpoint. In scenarios where reference to the endpoint is required, a value can be passed in create.
	GetKey() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// User-defined description of the endpoint.
	GetDescription() *string

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore, and numbers. It should begin with an upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	// The list of data assets that belong to the endpoint.
	GetDataAssets() []DataAsset
}

type updatedpendpointdetails struct {
	JsonData      []byte
	Key           *string          `mandatory:"true" json:"key"`
	ObjectVersion *int             `mandatory:"true" json:"objectVersion"`
	ModelVersion  *string          `mandatory:"false" json:"modelVersion"`
	ParentRef     *ParentReference `mandatory:"false" json:"parentRef"`
	Name          *string          `mandatory:"false" json:"name"`
	Description   *string          `mandatory:"false" json:"description"`
	ObjectStatus  *int             `mandatory:"false" json:"objectStatus"`
	Identifier    *string          `mandatory:"false" json:"identifier"`
	DataAssets    []DataAsset      `mandatory:"false" json:"dataAssets"`
	ModelType     string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *updatedpendpointdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatedpendpointdetails updatedpendpointdetails
	s := struct {
		Model Unmarshalerupdatedpendpointdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ObjectVersion = s.Model.ObjectVersion
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ObjectStatus = s.Model.ObjectStatus
	m.Identifier = s.Model.Identifier
	m.DataAssets = s.Model.DataAssets
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatedpendpointdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "PUBLIC_END_POINT":
		mm := UpdateDpEndpointFromPublic{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRIVATE_END_POINT":
		mm := UpdateDpEndpointFromPrivate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m updatedpendpointdetails) GetKey() *string {
	return m.Key
}

//GetObjectVersion returns ObjectVersion
func (m updatedpendpointdetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetModelVersion returns ModelVersion
func (m updatedpendpointdetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m updatedpendpointdetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m updatedpendpointdetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m updatedpendpointdetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m updatedpendpointdetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m updatedpendpointdetails) GetIdentifier() *string {
	return m.Identifier
}

//GetDataAssets returns DataAssets
func (m updatedpendpointdetails) GetDataAssets() []DataAsset {
	return m.DataAssets
}

func (m updatedpendpointdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatedpendpointdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDpEndpointDetailsModelTypeEnum Enum with underlying type: string
type UpdateDpEndpointDetailsModelTypeEnum string

// Set of constants representing the allowable values for UpdateDpEndpointDetailsModelTypeEnum
const (
	UpdateDpEndpointDetailsModelTypePrivateEndPoint UpdateDpEndpointDetailsModelTypeEnum = "PRIVATE_END_POINT"
	UpdateDpEndpointDetailsModelTypePublicEndPoint  UpdateDpEndpointDetailsModelTypeEnum = "PUBLIC_END_POINT"
)

var mappingUpdateDpEndpointDetailsModelTypeEnum = map[string]UpdateDpEndpointDetailsModelTypeEnum{
	"PRIVATE_END_POINT": UpdateDpEndpointDetailsModelTypePrivateEndPoint,
	"PUBLIC_END_POINT":  UpdateDpEndpointDetailsModelTypePublicEndPoint,
}

var mappingUpdateDpEndpointDetailsModelTypeEnumLowerCase = map[string]UpdateDpEndpointDetailsModelTypeEnum{
	"private_end_point": UpdateDpEndpointDetailsModelTypePrivateEndPoint,
	"public_end_point":  UpdateDpEndpointDetailsModelTypePublicEndPoint,
}

// GetUpdateDpEndpointDetailsModelTypeEnumValues Enumerates the set of values for UpdateDpEndpointDetailsModelTypeEnum
func GetUpdateDpEndpointDetailsModelTypeEnumValues() []UpdateDpEndpointDetailsModelTypeEnum {
	values := make([]UpdateDpEndpointDetailsModelTypeEnum, 0)
	for _, v := range mappingUpdateDpEndpointDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDpEndpointDetailsModelTypeEnumStringValues Enumerates the set of values in String for UpdateDpEndpointDetailsModelTypeEnum
func GetUpdateDpEndpointDetailsModelTypeEnumStringValues() []string {
	return []string{
		"PRIVATE_END_POINT",
		"PUBLIC_END_POINT",
	}
}

// GetMappingUpdateDpEndpointDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDpEndpointDetailsModelTypeEnum(val string) (UpdateDpEndpointDetailsModelTypeEnum, bool) {
	enum, ok := mappingUpdateDpEndpointDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
