// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDataAssetDetails Properties used in data asset update operations.
type CreateDataAssetDetails interface {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	GetName() *string

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	GetIdentifier() *string

	// Currently not used on data asset creation. Reserved for future.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	// Detailed description for the object.
	GetDescription() *string

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// The external key for the object
	GetExternalKey() *string

	// assetProperties
	GetAssetProperties() map[string]string

	GetRegistryMetadata() *RegistryMetadata
}

type createdataassetdetails struct {
	JsonData         []byte
	Name             *string           `mandatory:"true" json:"name"`
	Identifier       *string           `mandatory:"true" json:"identifier"`
	Key              *string           `mandatory:"false" json:"key"`
	ModelVersion     *string           `mandatory:"false" json:"modelVersion"`
	Description      *string           `mandatory:"false" json:"description"`
	ObjectStatus     *int              `mandatory:"false" json:"objectStatus"`
	ExternalKey      *string           `mandatory:"false" json:"externalKey"`
	AssetProperties  map[string]string `mandatory:"false" json:"assetProperties"`
	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
	ModelType        string            `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *createdataassetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedataassetdetails createdataassetdetails
	s := struct {
		Model Unmarshalercreatedataassetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Identifier = s.Model.Identifier
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.Description = s.Model.Description
	m.ObjectStatus = s.Model.ObjectStatus
	m.ExternalKey = s.Model.ExternalKey
	m.AssetProperties = s.Model.AssetProperties
	m.RegistryMetadata = s.Model.RegistryMetadata
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdataassetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "ORACLE_DATA_ASSET":
		mm := CreateDataAssetFromOracle{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_DATA_ASSET":
		mm := CreateDataAssetFromAdwc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ATP_DATA_ASSET":
		mm := CreateDataAssetFromAtp{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_OBJECT_STORAGE_DATA_ASSET":
		mm := CreateDataAssetFromObjectStorage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m createdataassetdetails) GetName() *string {
	return m.Name
}

//GetIdentifier returns Identifier
func (m createdataassetdetails) GetIdentifier() *string {
	return m.Identifier
}

//GetKey returns Key
func (m createdataassetdetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m createdataassetdetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetDescription returns Description
func (m createdataassetdetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m createdataassetdetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetExternalKey returns ExternalKey
func (m createdataassetdetails) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m createdataassetdetails) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetRegistryMetadata returns RegistryMetadata
func (m createdataassetdetails) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m createdataassetdetails) String() string {
	return common.PointerString(m)
}

// CreateDataAssetDetailsModelTypeEnum Enum with underlying type: string
type CreateDataAssetDetailsModelTypeEnum string

// Set of constants representing the allowable values for CreateDataAssetDetailsModelTypeEnum
const (
	CreateDataAssetDetailsModelTypeDataAsset              CreateDataAssetDetailsModelTypeEnum = "ORACLE_DATA_ASSET"
	CreateDataAssetDetailsModelTypeObjectStorageDataAsset CreateDataAssetDetailsModelTypeEnum = "ORACLE_OBJECT_STORAGE_DATA_ASSET"
	CreateDataAssetDetailsModelTypeAtpDataAsset           CreateDataAssetDetailsModelTypeEnum = "ORACLE_ATP_DATA_ASSET"
	CreateDataAssetDetailsModelTypeAdwcDataAsset          CreateDataAssetDetailsModelTypeEnum = "ORACLE_ADWC_DATA_ASSET"
)

var mappingCreateDataAssetDetailsModelType = map[string]CreateDataAssetDetailsModelTypeEnum{
	"ORACLE_DATA_ASSET":                CreateDataAssetDetailsModelTypeDataAsset,
	"ORACLE_OBJECT_STORAGE_DATA_ASSET": CreateDataAssetDetailsModelTypeObjectStorageDataAsset,
	"ORACLE_ATP_DATA_ASSET":            CreateDataAssetDetailsModelTypeAtpDataAsset,
	"ORACLE_ADWC_DATA_ASSET":           CreateDataAssetDetailsModelTypeAdwcDataAsset,
}

// GetCreateDataAssetDetailsModelTypeEnumValues Enumerates the set of values for CreateDataAssetDetailsModelTypeEnum
func GetCreateDataAssetDetailsModelTypeEnumValues() []CreateDataAssetDetailsModelTypeEnum {
	values := make([]CreateDataAssetDetailsModelTypeEnum, 0)
	for _, v := range mappingCreateDataAssetDetailsModelType {
		values = append(values, v)
	}
	return values
}
