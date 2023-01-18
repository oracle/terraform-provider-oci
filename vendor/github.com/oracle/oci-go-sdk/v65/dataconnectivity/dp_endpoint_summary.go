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

// DpEndpointSummary The endpoint summary object.
type DpEndpointSummary interface {

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

type dpendpointsummary struct {
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
func (m *dpendpointsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdpendpointsummary dpendpointsummary
	s := struct {
		Model Unmarshalerdpendpointsummary
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
func (m *dpendpointsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "PUBLIC_END_POINT":
		mm := DpEndpointSummaryFromPublic{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRIVATE_END_POINT":
		mm := DpEndpointSummaryFromPrivate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m dpendpointsummary) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m dpendpointsummary) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m dpendpointsummary) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m dpendpointsummary) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m dpendpointsummary) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m dpendpointsummary) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m dpendpointsummary) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m dpendpointsummary) GetIdentifier() *string {
	return m.Identifier
}

//GetDataAssets returns DataAssets
func (m dpendpointsummary) GetDataAssets() []DataAsset {
	return m.DataAssets
}

func (m dpendpointsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dpendpointsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DpEndpointSummaryModelTypeEnum Enum with underlying type: string
type DpEndpointSummaryModelTypeEnum string

// Set of constants representing the allowable values for DpEndpointSummaryModelTypeEnum
const (
	DpEndpointSummaryModelTypePrivateEndPoint DpEndpointSummaryModelTypeEnum = "PRIVATE_END_POINT"
	DpEndpointSummaryModelTypePublicEndPoint  DpEndpointSummaryModelTypeEnum = "PUBLIC_END_POINT"
)

var mappingDpEndpointSummaryModelTypeEnum = map[string]DpEndpointSummaryModelTypeEnum{
	"PRIVATE_END_POINT": DpEndpointSummaryModelTypePrivateEndPoint,
	"PUBLIC_END_POINT":  DpEndpointSummaryModelTypePublicEndPoint,
}

var mappingDpEndpointSummaryModelTypeEnumLowerCase = map[string]DpEndpointSummaryModelTypeEnum{
	"private_end_point": DpEndpointSummaryModelTypePrivateEndPoint,
	"public_end_point":  DpEndpointSummaryModelTypePublicEndPoint,
}

// GetDpEndpointSummaryModelTypeEnumValues Enumerates the set of values for DpEndpointSummaryModelTypeEnum
func GetDpEndpointSummaryModelTypeEnumValues() []DpEndpointSummaryModelTypeEnum {
	values := make([]DpEndpointSummaryModelTypeEnum, 0)
	for _, v := range mappingDpEndpointSummaryModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDpEndpointSummaryModelTypeEnumStringValues Enumerates the set of values in String for DpEndpointSummaryModelTypeEnum
func GetDpEndpointSummaryModelTypeEnumStringValues() []string {
	return []string{
		"PRIVATE_END_POINT",
		"PUBLIC_END_POINT",
	}
}

// GetMappingDpEndpointSummaryModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDpEndpointSummaryModelTypeEnum(val string) (DpEndpointSummaryModelTypeEnum, bool) {
	enum, ok := mappingDpEndpointSummaryModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
