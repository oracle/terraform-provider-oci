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
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
)

// DpEndpointDetails The endpoint details.
type DpEndpointDetails interface {

	// Generated key that can be used in API calls to identify endpoint. On scenarios where reference to the endpoint is needed, a value can be passed in create.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// User-defined description for the endpoint.
	GetDescription() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	// List of data assets which belongs to this endpoint
	GetDataAssets() []DataAsset
}

type dpendpointdetails struct {
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
func (m *dpendpointdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdpendpointdetails dpendpointdetails
	s := struct {
		Model Unmarshalerdpendpointdetails
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
func (m *dpendpointdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "PRIVATE_END_POINT":
		mm := DpEndpointFromPrivateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PUBLIC_END_POINT":
		mm := DpEndpointFromPublicDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m dpendpointdetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m dpendpointdetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m dpendpointdetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m dpendpointdetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m dpendpointdetails) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m dpendpointdetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m dpendpointdetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m dpendpointdetails) GetIdentifier() *string {
	return m.Identifier
}

//GetDataAssets returns DataAssets
func (m dpendpointdetails) GetDataAssets() []DataAsset {
	return m.DataAssets
}

func (m dpendpointdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dpendpointdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DpEndpointDetailsModelTypeEnum Enum with underlying type: string
type DpEndpointDetailsModelTypeEnum string

// Set of constants representing the allowable values for DpEndpointDetailsModelTypeEnum
const (
	DpEndpointDetailsModelTypePrivateEndPoint DpEndpointDetailsModelTypeEnum = "PRIVATE_END_POINT"
	DpEndpointDetailsModelTypePublicEndPoint  DpEndpointDetailsModelTypeEnum = "PUBLIC_END_POINT"
)

var mappingDpEndpointDetailsModelTypeEnum = map[string]DpEndpointDetailsModelTypeEnum{
	"PRIVATE_END_POINT": DpEndpointDetailsModelTypePrivateEndPoint,
	"PUBLIC_END_POINT":  DpEndpointDetailsModelTypePublicEndPoint,
}

var mappingDpEndpointDetailsModelTypeEnumLowerCase = map[string]DpEndpointDetailsModelTypeEnum{
	"private_end_point": DpEndpointDetailsModelTypePrivateEndPoint,
	"public_end_point":  DpEndpointDetailsModelTypePublicEndPoint,
}

// GetDpEndpointDetailsModelTypeEnumValues Enumerates the set of values for DpEndpointDetailsModelTypeEnum
func GetDpEndpointDetailsModelTypeEnumValues() []DpEndpointDetailsModelTypeEnum {
	values := make([]DpEndpointDetailsModelTypeEnum, 0)
	for _, v := range mappingDpEndpointDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDpEndpointDetailsModelTypeEnumStringValues Enumerates the set of values in String for DpEndpointDetailsModelTypeEnum
func GetDpEndpointDetailsModelTypeEnumStringValues() []string {
	return []string{
		"PRIVATE_END_POINT",
		"PUBLIC_END_POINT",
	}
}

// GetMappingDpEndpointDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDpEndpointDetailsModelTypeEnum(val string) (DpEndpointDetailsModelTypeEnum, bool) {
	enum, ok := mappingDpEndpointDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
