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

// DataAsset The data asset type.
type DataAsset interface {

	// Generated key that can be used in API calls to identify data asset.
	GetKey() *string

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

	GetNativeTypeSystem() *TypeSystem

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	GetParentRef() *ParentReference

	GetMetadata() *ObjectMetadata

	// A map, if provided key is replaced with generated key, this structure provides mapping between user provided key and generated key
	GetKeyMap() map[string]string
}

type dataasset struct {
	JsonData         []byte
	Key              *string           `mandatory:"false" json:"key"`
	ModelVersion     *string           `mandatory:"false" json:"modelVersion"`
	Name             *string           `mandatory:"false" json:"name"`
	Description      *string           `mandatory:"false" json:"description"`
	ObjectStatus     *int              `mandatory:"false" json:"objectStatus"`
	Identifier       *string           `mandatory:"false" json:"identifier"`
	ExternalKey      *string           `mandatory:"false" json:"externalKey"`
	AssetProperties  map[string]string `mandatory:"false" json:"assetProperties"`
	NativeTypeSystem *TypeSystem       `mandatory:"false" json:"nativeTypeSystem"`
	ObjectVersion    *int              `mandatory:"false" json:"objectVersion"`
	ParentRef        *ParentReference  `mandatory:"false" json:"parentRef"`
	Metadata         *ObjectMetadata   `mandatory:"false" json:"metadata"`
	KeyMap           map[string]string `mandatory:"false" json:"keyMap"`
	ModelType        string            `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *dataasset) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataasset dataasset
	s := struct {
		Model Unmarshalerdataasset
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ObjectStatus = s.Model.ObjectStatus
	m.Identifier = s.Model.Identifier
	m.ExternalKey = s.Model.ExternalKey
	m.AssetProperties = s.Model.AssetProperties
	m.NativeTypeSystem = s.Model.NativeTypeSystem
	m.ObjectVersion = s.Model.ObjectVersion
	m.ParentRef = s.Model.ParentRef
	m.Metadata = s.Model.Metadata
	m.KeyMap = s.Model.KeyMap
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dataasset) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "ORACLE_DATA_ASSET":
		mm := DataAssetFromOracleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_DATA_ASSET":
		mm := DataAssetFromAdwcDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_OBJECT_STORAGE_DATA_ASSET":
		mm := DataAssetFromObjectStorageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ATP_DATA_ASSET":
		mm := DataAssetFromAtpDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m dataasset) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m dataasset) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m dataasset) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m dataasset) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m dataasset) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m dataasset) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m dataasset) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m dataasset) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetNativeTypeSystem returns NativeTypeSystem
func (m dataasset) GetNativeTypeSystem() *TypeSystem {
	return m.NativeTypeSystem
}

//GetObjectVersion returns ObjectVersion
func (m dataasset) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetParentRef returns ParentRef
func (m dataasset) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetMetadata returns Metadata
func (m dataasset) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

//GetKeyMap returns KeyMap
func (m dataasset) GetKeyMap() map[string]string {
	return m.KeyMap
}

func (m dataasset) String() string {
	return common.PointerString(m)
}

// DataAssetModelTypeEnum Enum with underlying type: string
type DataAssetModelTypeEnum string

// Set of constants representing the allowable values for DataAssetModelTypeEnum
const (
	DataAssetModelTypeDataAsset              DataAssetModelTypeEnum = "ORACLE_DATA_ASSET"
	DataAssetModelTypeObjectStorageDataAsset DataAssetModelTypeEnum = "ORACLE_OBJECT_STORAGE_DATA_ASSET"
	DataAssetModelTypeAtpDataAsset           DataAssetModelTypeEnum = "ORACLE_ATP_DATA_ASSET"
	DataAssetModelTypeAdwcDataAsset          DataAssetModelTypeEnum = "ORACLE_ADWC_DATA_ASSET"
)

var mappingDataAssetModelType = map[string]DataAssetModelTypeEnum{
	"ORACLE_DATA_ASSET":                DataAssetModelTypeDataAsset,
	"ORACLE_OBJECT_STORAGE_DATA_ASSET": DataAssetModelTypeObjectStorageDataAsset,
	"ORACLE_ATP_DATA_ASSET":            DataAssetModelTypeAtpDataAsset,
	"ORACLE_ADWC_DATA_ASSET":           DataAssetModelTypeAdwcDataAsset,
}

// GetDataAssetModelTypeEnumValues Enumerates the set of values for DataAssetModelTypeEnum
func GetDataAssetModelTypeEnumValues() []DataAssetModelTypeEnum {
	values := make([]DataAssetModelTypeEnum, 0)
	for _, v := range mappingDataAssetModelType {
		values = append(values, v)
	}
	return values
}
