// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DataAsset Represents a data source in the Data Integration service.
type DataAsset interface {

	// Generated key that can be used in API calls to identify data asset.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// User-defined description of the data asset.
	GetDescription() *string

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	// The external key for the object.
	GetExternalKey() *string

	// Additional properties for the data asset.
	GetAssetProperties() map[string]string

	GetNativeTypeSystem() *TypeSystem

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	GetParentRef() *ParentReference

	GetMetadata() *ObjectMetadata

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
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
	case "GENERIC_JDBC_DATA_ASSET":
		mm := DataAssetFromJdbc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_DATA_ASSET":
		mm := DataAssetFromOracleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_DATA_ASSET":
		mm := DataAssetFromAdwcDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMAZON_S3_DATA_ASSET":
		mm := DataAssetFromAmazonS3{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_OBJECT_STORAGE_DATA_ASSET":
		mm := DataAssetFromObjectStorageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FUSION_APP_DATA_ASSET":
		mm := DataAssetFromFusionApp{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ATP_DATA_ASSET":
		mm := DataAssetFromAtpDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL_DATA_ASSET":
		mm := DataAssetFromMySql{}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dataasset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataAssetModelTypeEnum Enum with underlying type: string
type DataAssetModelTypeEnum string

// Set of constants representing the allowable values for DataAssetModelTypeEnum
const (
	DataAssetModelTypeOracleDataAsset              DataAssetModelTypeEnum = "ORACLE_DATA_ASSET"
	DataAssetModelTypeOracleObjectStorageDataAsset DataAssetModelTypeEnum = "ORACLE_OBJECT_STORAGE_DATA_ASSET"
	DataAssetModelTypeOracleAtpDataAsset           DataAssetModelTypeEnum = "ORACLE_ATP_DATA_ASSET"
	DataAssetModelTypeOracleAdwcDataAsset          DataAssetModelTypeEnum = "ORACLE_ADWC_DATA_ASSET"
	DataAssetModelTypeMysqlDataAsset               DataAssetModelTypeEnum = "MYSQL_DATA_ASSET"
	DataAssetModelTypeGenericJdbcDataAsset         DataAssetModelTypeEnum = "GENERIC_JDBC_DATA_ASSET"
	DataAssetModelTypeFusionAppDataAsset           DataAssetModelTypeEnum = "FUSION_APP_DATA_ASSET"
	DataAssetModelTypeAmazonS3DataAsset            DataAssetModelTypeEnum = "AMAZON_S3_DATA_ASSET"
)

var mappingDataAssetModelTypeEnum = map[string]DataAssetModelTypeEnum{
	"ORACLE_DATA_ASSET":                DataAssetModelTypeOracleDataAsset,
	"ORACLE_OBJECT_STORAGE_DATA_ASSET": DataAssetModelTypeOracleObjectStorageDataAsset,
	"ORACLE_ATP_DATA_ASSET":            DataAssetModelTypeOracleAtpDataAsset,
	"ORACLE_ADWC_DATA_ASSET":           DataAssetModelTypeOracleAdwcDataAsset,
	"MYSQL_DATA_ASSET":                 DataAssetModelTypeMysqlDataAsset,
	"GENERIC_JDBC_DATA_ASSET":          DataAssetModelTypeGenericJdbcDataAsset,
	"FUSION_APP_DATA_ASSET":            DataAssetModelTypeFusionAppDataAsset,
	"AMAZON_S3_DATA_ASSET":             DataAssetModelTypeAmazonS3DataAsset,
}

// GetDataAssetModelTypeEnumValues Enumerates the set of values for DataAssetModelTypeEnum
func GetDataAssetModelTypeEnumValues() []DataAssetModelTypeEnum {
	values := make([]DataAssetModelTypeEnum, 0)
	for _, v := range mappingDataAssetModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataAssetModelTypeEnumStringValues Enumerates the set of values in String for DataAssetModelTypeEnum
func GetDataAssetModelTypeEnumStringValues() []string {
	return []string{
		"ORACLE_DATA_ASSET",
		"ORACLE_OBJECT_STORAGE_DATA_ASSET",
		"ORACLE_ATP_DATA_ASSET",
		"ORACLE_ADWC_DATA_ASSET",
		"MYSQL_DATA_ASSET",
		"GENERIC_JDBC_DATA_ASSET",
		"FUSION_APP_DATA_ASSET",
		"AMAZON_S3_DATA_ASSET",
	}
}

// GetMappingDataAssetModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataAssetModelTypeEnum(val string) (DataAssetModelTypeEnum, bool) {
	mappingDataAssetModelTypeEnumIgnoreCase := make(map[string]DataAssetModelTypeEnum)
	for k, v := range mappingDataAssetModelTypeEnum {
		mappingDataAssetModelTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDataAssetModelTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
