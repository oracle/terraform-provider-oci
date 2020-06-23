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

// DataAssetSummary The summary object for data asset.
type DataAssetSummary interface {

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
}

type dataassetsummary struct {
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
	ModelType        string            `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *dataassetsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataassetsummary dataassetsummary
	s := struct {
		Model Unmarshalerdataassetsummary
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
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dataassetsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "ORACLE_ATP_DATA_ASSET":
		mm := DataAssetSummaryFromAtp{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_DATA_ASSET":
		mm := DataAssetSummaryFromAdwc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_OBJECT_STORAGE_DATA_ASSET":
		mm := DataAssetSummaryFromObjectStorage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_DATA_ASSET":
		mm := DataAssetSummaryFromOracle{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m dataassetsummary) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m dataassetsummary) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m dataassetsummary) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m dataassetsummary) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m dataassetsummary) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m dataassetsummary) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m dataassetsummary) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m dataassetsummary) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetNativeTypeSystem returns NativeTypeSystem
func (m dataassetsummary) GetNativeTypeSystem() *TypeSystem {
	return m.NativeTypeSystem
}

//GetObjectVersion returns ObjectVersion
func (m dataassetsummary) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetParentRef returns ParentRef
func (m dataassetsummary) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetMetadata returns Metadata
func (m dataassetsummary) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m dataassetsummary) String() string {
	return common.PointerString(m)
}

// DataAssetSummaryModelTypeEnum Enum with underlying type: string
type DataAssetSummaryModelTypeEnum string

// Set of constants representing the allowable values for DataAssetSummaryModelTypeEnum
const (
	DataAssetSummaryModelTypeDataAsset              DataAssetSummaryModelTypeEnum = "ORACLE_DATA_ASSET"
	DataAssetSummaryModelTypeObjectStorageDataAsset DataAssetSummaryModelTypeEnum = "ORACLE_OBJECT_STORAGE_DATA_ASSET"
	DataAssetSummaryModelTypeAtpDataAsset           DataAssetSummaryModelTypeEnum = "ORACLE_ATP_DATA_ASSET"
	DataAssetSummaryModelTypeAdwcDataAsset          DataAssetSummaryModelTypeEnum = "ORACLE_ADWC_DATA_ASSET"
)

var mappingDataAssetSummaryModelType = map[string]DataAssetSummaryModelTypeEnum{
	"ORACLE_DATA_ASSET":                DataAssetSummaryModelTypeDataAsset,
	"ORACLE_OBJECT_STORAGE_DATA_ASSET": DataAssetSummaryModelTypeObjectStorageDataAsset,
	"ORACLE_ATP_DATA_ASSET":            DataAssetSummaryModelTypeAtpDataAsset,
	"ORACLE_ADWC_DATA_ASSET":           DataAssetSummaryModelTypeAdwcDataAsset,
}

// GetDataAssetSummaryModelTypeEnumValues Enumerates the set of values for DataAssetSummaryModelTypeEnum
func GetDataAssetSummaryModelTypeEnumValues() []DataAssetSummaryModelTypeEnum {
	values := make([]DataAssetSummaryModelTypeEnum, 0)
	for _, v := range mappingDataAssetSummaryModelType {
		values = append(values, v)
	}
	return values
}
