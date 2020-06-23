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

// UpdateDataAssetDetails Properties used in data asset update operations.
type UpdateDataAssetDetails interface {

	// Generated key that can be used in API calls to identify data asset.
	GetKey() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// The model version of an object.
	GetModelVersion() *string

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	GetName() *string

	// Detailed description for the object.
	GetDescription() *string

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	GetIdentifier() *string

	// The external key for the object
	GetExternalKey() *string

	// assetProperties
	GetAssetProperties() map[string]string

	GetRegistryMetadata() *RegistryMetadata
}

type updatedataassetdetails struct {
	JsonData         []byte
	Key              *string           `mandatory:"true" json:"key"`
	ObjectVersion    *int              `mandatory:"true" json:"objectVersion"`
	ModelVersion     *string           `mandatory:"false" json:"modelVersion"`
	Name             *string           `mandatory:"false" json:"name"`
	Description      *string           `mandatory:"false" json:"description"`
	ObjectStatus     *int              `mandatory:"false" json:"objectStatus"`
	Identifier       *string           `mandatory:"false" json:"identifier"`
	ExternalKey      *string           `mandatory:"false" json:"externalKey"`
	AssetProperties  map[string]string `mandatory:"false" json:"assetProperties"`
	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
	ModelType        string            `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *updatedataassetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatedataassetdetails updatedataassetdetails
	s := struct {
		Model Unmarshalerupdatedataassetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ObjectVersion = s.Model.ObjectVersion
	m.ModelVersion = s.Model.ModelVersion
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ObjectStatus = s.Model.ObjectStatus
	m.Identifier = s.Model.Identifier
	m.ExternalKey = s.Model.ExternalKey
	m.AssetProperties = s.Model.AssetProperties
	m.RegistryMetadata = s.Model.RegistryMetadata
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatedataassetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "ORACLE_ATP_DATA_ASSET":
		mm := UpdateDataAssetFromAtp{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_DATA_ASSET":
		mm := UpdateDataAssetFromAdwc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_OBJECT_STORAGE_DATA_ASSET":
		mm := UpdateDataAssetFromObjectStorage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_DATA_ASSET":
		mm := UpdateDataAssetFromOracle{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m updatedataassetdetails) GetKey() *string {
	return m.Key
}

//GetObjectVersion returns ObjectVersion
func (m updatedataassetdetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetModelVersion returns ModelVersion
func (m updatedataassetdetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m updatedataassetdetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m updatedataassetdetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m updatedataassetdetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m updatedataassetdetails) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m updatedataassetdetails) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m updatedataassetdetails) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetRegistryMetadata returns RegistryMetadata
func (m updatedataassetdetails) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m updatedataassetdetails) String() string {
	return common.PointerString(m)
}

// UpdateDataAssetDetailsModelTypeEnum Enum with underlying type: string
type UpdateDataAssetDetailsModelTypeEnum string

// Set of constants representing the allowable values for UpdateDataAssetDetailsModelTypeEnum
const (
	UpdateDataAssetDetailsModelTypeDataAsset              UpdateDataAssetDetailsModelTypeEnum = "ORACLE_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeObjectStorageDataAsset UpdateDataAssetDetailsModelTypeEnum = "ORACLE_OBJECT_STORAGE_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeAtpDataAsset           UpdateDataAssetDetailsModelTypeEnum = "ORACLE_ATP_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeAdwcDataAsset          UpdateDataAssetDetailsModelTypeEnum = "ORACLE_ADWC_DATA_ASSET"
)

var mappingUpdateDataAssetDetailsModelType = map[string]UpdateDataAssetDetailsModelTypeEnum{
	"ORACLE_DATA_ASSET":                UpdateDataAssetDetailsModelTypeDataAsset,
	"ORACLE_OBJECT_STORAGE_DATA_ASSET": UpdateDataAssetDetailsModelTypeObjectStorageDataAsset,
	"ORACLE_ATP_DATA_ASSET":            UpdateDataAssetDetailsModelTypeAtpDataAsset,
	"ORACLE_ADWC_DATA_ASSET":           UpdateDataAssetDetailsModelTypeAdwcDataAsset,
}

// GetUpdateDataAssetDetailsModelTypeEnumValues Enumerates the set of values for UpdateDataAssetDetailsModelTypeEnum
func GetUpdateDataAssetDetailsModelTypeEnumValues() []UpdateDataAssetDetailsModelTypeEnum {
	values := make([]UpdateDataAssetDetailsModelTypeEnum, 0)
	for _, v := range mappingUpdateDataAssetDetailsModelType {
		values = append(values, v)
	}
	return values
}
