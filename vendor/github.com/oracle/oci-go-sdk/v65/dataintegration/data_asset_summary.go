// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataAssetSummary The summary object for data asset.
type DataAssetSummary interface {

	// Generated key that can be used in API calls to identify data asset.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// The user-defined description of the data asset.
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
	case "ORACLE_SIEBEL_DATA_ASSET":
		mm := DataAssetSummaryFromOracleSiebel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL_HEATWAVE_DATA_ASSET":
		mm := DataAssetSummaryFromMySqlHeatWave{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL_DATA_ASSET":
		mm := DataAssetSummaryFromMySql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HDFS_DATA_ASSET":
		mm := DataAssetSummaryFromHdfs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMAZON_S3_DATA_ASSET":
		mm := DataAssetSummaryFromAmazonS3{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_PEOPLESOFT_DATA_ASSET":
		mm := DataAssetSummaryFromOraclePeopleSoft{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_OBJECT_STORAGE_DATA_ASSET":
		mm := DataAssetSummaryFromObjectStorage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LAKE_DATA_ASSET":
		mm := DataAssetSummaryFromLake{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REST_DATA_ASSET":
		mm := DataAssetSummaryFromRest{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_DATA_ASSET":
		mm := DataAssetSummaryFromOracle{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ATP_DATA_ASSET":
		mm := DataAssetSummaryFromAtp{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_DATA_ASSET":
		mm := DataAssetSummaryFromAdwc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENERIC_JDBC_DATA_ASSET":
		mm := DataAssetSummaryFromJdbc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_EBS_DATA_ASSET":
		mm := DataAssetSummaryFromOracleEbs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FUSION_APP_DATA_ASSET":
		mm := DataAssetSummaryFromFusionApp{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DataAssetSummary: %s.", m.ModelType)
		return *m, nil
	}
}

// GetKey returns Key
func (m dataassetsummary) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m dataassetsummary) GetModelVersion() *string {
	return m.ModelVersion
}

// GetName returns Name
func (m dataassetsummary) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m dataassetsummary) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m dataassetsummary) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m dataassetsummary) GetIdentifier() *string {
	return m.Identifier
}

// GetExternalKey returns ExternalKey
func (m dataassetsummary) GetExternalKey() *string {
	return m.ExternalKey
}

// GetAssetProperties returns AssetProperties
func (m dataassetsummary) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

// GetNativeTypeSystem returns NativeTypeSystem
func (m dataassetsummary) GetNativeTypeSystem() *TypeSystem {
	return m.NativeTypeSystem
}

// GetObjectVersion returns ObjectVersion
func (m dataassetsummary) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetParentRef returns ParentRef
func (m dataassetsummary) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetMetadata returns Metadata
func (m dataassetsummary) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m dataassetsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dataassetsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataAssetSummaryModelTypeEnum Enum with underlying type: string
type DataAssetSummaryModelTypeEnum string

// Set of constants representing the allowable values for DataAssetSummaryModelTypeEnum
const (
	DataAssetSummaryModelTypeOracleDataAsset              DataAssetSummaryModelTypeEnum = "ORACLE_DATA_ASSET"
	DataAssetSummaryModelTypeOracleObjectStorageDataAsset DataAssetSummaryModelTypeEnum = "ORACLE_OBJECT_STORAGE_DATA_ASSET"
	DataAssetSummaryModelTypeOracleAtpDataAsset           DataAssetSummaryModelTypeEnum = "ORACLE_ATP_DATA_ASSET"
	DataAssetSummaryModelTypeOracleAdwcDataAsset          DataAssetSummaryModelTypeEnum = "ORACLE_ADWC_DATA_ASSET"
	DataAssetSummaryModelTypeMysqlDataAsset               DataAssetSummaryModelTypeEnum = "MYSQL_DATA_ASSET"
	DataAssetSummaryModelTypeGenericJdbcDataAsset         DataAssetSummaryModelTypeEnum = "GENERIC_JDBC_DATA_ASSET"
	DataAssetSummaryModelTypeFusionAppDataAsset           DataAssetSummaryModelTypeEnum = "FUSION_APP_DATA_ASSET"
	DataAssetSummaryModelTypeAmazonS3DataAsset            DataAssetSummaryModelTypeEnum = "AMAZON_S3_DATA_ASSET"
	DataAssetSummaryModelTypeLakeDataAsset                DataAssetSummaryModelTypeEnum = "LAKE_DATA_ASSET"
	DataAssetSummaryModelTypeOraclePeoplesoftDataAsset    DataAssetSummaryModelTypeEnum = "ORACLE_PEOPLESOFT_DATA_ASSET"
	DataAssetSummaryModelTypeOracleSiebelDataAsset        DataAssetSummaryModelTypeEnum = "ORACLE_SIEBEL_DATA_ASSET"
	DataAssetSummaryModelTypeOracleEbsDataAsset           DataAssetSummaryModelTypeEnum = "ORACLE_EBS_DATA_ASSET"
	DataAssetSummaryModelTypeHdfsDataAsset                DataAssetSummaryModelTypeEnum = "HDFS_DATA_ASSET"
	DataAssetSummaryModelTypeMysqlHeatwaveDataAsset       DataAssetSummaryModelTypeEnum = "MYSQL_HEATWAVE_DATA_ASSET"
	DataAssetSummaryModelTypeRestDataAsset                DataAssetSummaryModelTypeEnum = "REST_DATA_ASSET"
)

var mappingDataAssetSummaryModelTypeEnum = map[string]DataAssetSummaryModelTypeEnum{
	"ORACLE_DATA_ASSET":                DataAssetSummaryModelTypeOracleDataAsset,
	"ORACLE_OBJECT_STORAGE_DATA_ASSET": DataAssetSummaryModelTypeOracleObjectStorageDataAsset,
	"ORACLE_ATP_DATA_ASSET":            DataAssetSummaryModelTypeOracleAtpDataAsset,
	"ORACLE_ADWC_DATA_ASSET":           DataAssetSummaryModelTypeOracleAdwcDataAsset,
	"MYSQL_DATA_ASSET":                 DataAssetSummaryModelTypeMysqlDataAsset,
	"GENERIC_JDBC_DATA_ASSET":          DataAssetSummaryModelTypeGenericJdbcDataAsset,
	"FUSION_APP_DATA_ASSET":            DataAssetSummaryModelTypeFusionAppDataAsset,
	"AMAZON_S3_DATA_ASSET":             DataAssetSummaryModelTypeAmazonS3DataAsset,
	"LAKE_DATA_ASSET":                  DataAssetSummaryModelTypeLakeDataAsset,
	"ORACLE_PEOPLESOFT_DATA_ASSET":     DataAssetSummaryModelTypeOraclePeoplesoftDataAsset,
	"ORACLE_SIEBEL_DATA_ASSET":         DataAssetSummaryModelTypeOracleSiebelDataAsset,
	"ORACLE_EBS_DATA_ASSET":            DataAssetSummaryModelTypeOracleEbsDataAsset,
	"HDFS_DATA_ASSET":                  DataAssetSummaryModelTypeHdfsDataAsset,
	"MYSQL_HEATWAVE_DATA_ASSET":        DataAssetSummaryModelTypeMysqlHeatwaveDataAsset,
	"REST_DATA_ASSET":                  DataAssetSummaryModelTypeRestDataAsset,
}

var mappingDataAssetSummaryModelTypeEnumLowerCase = map[string]DataAssetSummaryModelTypeEnum{
	"oracle_data_asset":                DataAssetSummaryModelTypeOracleDataAsset,
	"oracle_object_storage_data_asset": DataAssetSummaryModelTypeOracleObjectStorageDataAsset,
	"oracle_atp_data_asset":            DataAssetSummaryModelTypeOracleAtpDataAsset,
	"oracle_adwc_data_asset":           DataAssetSummaryModelTypeOracleAdwcDataAsset,
	"mysql_data_asset":                 DataAssetSummaryModelTypeMysqlDataAsset,
	"generic_jdbc_data_asset":          DataAssetSummaryModelTypeGenericJdbcDataAsset,
	"fusion_app_data_asset":            DataAssetSummaryModelTypeFusionAppDataAsset,
	"amazon_s3_data_asset":             DataAssetSummaryModelTypeAmazonS3DataAsset,
	"lake_data_asset":                  DataAssetSummaryModelTypeLakeDataAsset,
	"oracle_peoplesoft_data_asset":     DataAssetSummaryModelTypeOraclePeoplesoftDataAsset,
	"oracle_siebel_data_asset":         DataAssetSummaryModelTypeOracleSiebelDataAsset,
	"oracle_ebs_data_asset":            DataAssetSummaryModelTypeOracleEbsDataAsset,
	"hdfs_data_asset":                  DataAssetSummaryModelTypeHdfsDataAsset,
	"mysql_heatwave_data_asset":        DataAssetSummaryModelTypeMysqlHeatwaveDataAsset,
	"rest_data_asset":                  DataAssetSummaryModelTypeRestDataAsset,
}

// GetDataAssetSummaryModelTypeEnumValues Enumerates the set of values for DataAssetSummaryModelTypeEnum
func GetDataAssetSummaryModelTypeEnumValues() []DataAssetSummaryModelTypeEnum {
	values := make([]DataAssetSummaryModelTypeEnum, 0)
	for _, v := range mappingDataAssetSummaryModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataAssetSummaryModelTypeEnumStringValues Enumerates the set of values in String for DataAssetSummaryModelTypeEnum
func GetDataAssetSummaryModelTypeEnumStringValues() []string {
	return []string{
		"ORACLE_DATA_ASSET",
		"ORACLE_OBJECT_STORAGE_DATA_ASSET",
		"ORACLE_ATP_DATA_ASSET",
		"ORACLE_ADWC_DATA_ASSET",
		"MYSQL_DATA_ASSET",
		"GENERIC_JDBC_DATA_ASSET",
		"FUSION_APP_DATA_ASSET",
		"AMAZON_S3_DATA_ASSET",
		"LAKE_DATA_ASSET",
		"ORACLE_PEOPLESOFT_DATA_ASSET",
		"ORACLE_SIEBEL_DATA_ASSET",
		"ORACLE_EBS_DATA_ASSET",
		"HDFS_DATA_ASSET",
		"MYSQL_HEATWAVE_DATA_ASSET",
		"REST_DATA_ASSET",
	}
}

// GetMappingDataAssetSummaryModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataAssetSummaryModelTypeEnum(val string) (DataAssetSummaryModelTypeEnum, bool) {
	enum, ok := mappingDataAssetSummaryModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
