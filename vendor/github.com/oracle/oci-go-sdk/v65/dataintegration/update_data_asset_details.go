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

// UpdateDataAssetDetails Properties used in data asset update operations.
type UpdateDataAssetDetails interface {

	// Generated key that can be used in API calls to identify data asset.
	GetKey() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

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

	GetRegistryMetadata() *RegistryMetadata
}

type updatedataassetdetails struct {
	JsonData         []byte
	ModelVersion     *string           `mandatory:"false" json:"modelVersion"`
	Name             *string           `mandatory:"false" json:"name"`
	Description      *string           `mandatory:"false" json:"description"`
	ObjectStatus     *int              `mandatory:"false" json:"objectStatus"`
	Identifier       *string           `mandatory:"false" json:"identifier"`
	ExternalKey      *string           `mandatory:"false" json:"externalKey"`
	AssetProperties  map[string]string `mandatory:"false" json:"assetProperties"`
	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
	Key              *string           `mandatory:"true" json:"key"`
	ObjectVersion    *int              `mandatory:"true" json:"objectVersion"`
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
	case "REST_DATA_ASSET":
		mm := UpdateDataAssetFromRest{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL_HEATWAVE_DATA_ASSET":
		mm := UpdateDataAssetFromMySqlHeatWave{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_EBS_DATA_ASSET":
		mm := UpdateDataAssetFromOracleEbs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_DATA_ASSET":
		mm := UpdateDataAssetFromAdwc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HDFS_DATA_ASSET":
		mm := UpdateDataAssetFromHdfs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_PEOPLESOFT_DATA_ASSET":
		mm := UpdateDataAssetFromOraclePeopleSoft{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_OBJECT_STORAGE_DATA_ASSET":
		mm := UpdateDataAssetFromObjectStorage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMAZON_S3_DATA_ASSET":
		mm := UpdateDataAssetFromAmazonS3{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL_DATA_ASSET":
		mm := UpdateDataAssetFromMySql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_SIEBEL_DATA_ASSET":
		mm := UpdateDataAssetFromOracleSiebel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LAKE_DATA_ASSET":
		mm := UpdateDataAssetFromLake{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ATP_DATA_ASSET":
		mm := UpdateDataAssetFromAtp{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENERIC_JDBC_DATA_ASSET":
		mm := UpdateDataAssetFromJdbc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FUSION_APP_DATA_ASSET":
		mm := UpdateDataAssetFromFusionApp{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_DATA_ASSET":
		mm := UpdateDataAssetFromOracle{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateDataAssetDetails: %s.", m.ModelType)
		return *m, nil
	}
}

// GetModelVersion returns ModelVersion
func (m updatedataassetdetails) GetModelVersion() *string {
	return m.ModelVersion
}

// GetName returns Name
func (m updatedataassetdetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m updatedataassetdetails) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m updatedataassetdetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m updatedataassetdetails) GetIdentifier() *string {
	return m.Identifier
}

// GetExternalKey returns ExternalKey
func (m updatedataassetdetails) GetExternalKey() *string {
	return m.ExternalKey
}

// GetAssetProperties returns AssetProperties
func (m updatedataassetdetails) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

// GetRegistryMetadata returns RegistryMetadata
func (m updatedataassetdetails) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

// GetKey returns Key
func (m updatedataassetdetails) GetKey() *string {
	return m.Key
}

// GetObjectVersion returns ObjectVersion
func (m updatedataassetdetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

func (m updatedataassetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatedataassetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDataAssetDetailsModelTypeEnum Enum with underlying type: string
type UpdateDataAssetDetailsModelTypeEnum string

// Set of constants representing the allowable values for UpdateDataAssetDetailsModelTypeEnum
const (
	UpdateDataAssetDetailsModelTypeOracleDataAsset              UpdateDataAssetDetailsModelTypeEnum = "ORACLE_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeOracleObjectStorageDataAsset UpdateDataAssetDetailsModelTypeEnum = "ORACLE_OBJECT_STORAGE_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeOracleAtpDataAsset           UpdateDataAssetDetailsModelTypeEnum = "ORACLE_ATP_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeOracleAdwcDataAsset          UpdateDataAssetDetailsModelTypeEnum = "ORACLE_ADWC_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeMysqlDataAsset               UpdateDataAssetDetailsModelTypeEnum = "MYSQL_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeGenericJdbcDataAsset         UpdateDataAssetDetailsModelTypeEnum = "GENERIC_JDBC_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeFusionAppDataAsset           UpdateDataAssetDetailsModelTypeEnum = "FUSION_APP_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeAmazonS3DataAsset            UpdateDataAssetDetailsModelTypeEnum = "AMAZON_S3_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeLakeDataAsset                UpdateDataAssetDetailsModelTypeEnum = "LAKE_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeOraclePeoplesoftDataAsset    UpdateDataAssetDetailsModelTypeEnum = "ORACLE_PEOPLESOFT_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeOracleSiebelDataAsset        UpdateDataAssetDetailsModelTypeEnum = "ORACLE_SIEBEL_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeOracleEbsDataAsset           UpdateDataAssetDetailsModelTypeEnum = "ORACLE_EBS_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeHdfsDataAsset                UpdateDataAssetDetailsModelTypeEnum = "HDFS_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeMysqlHeatwaveDataAsset       UpdateDataAssetDetailsModelTypeEnum = "MYSQL_HEATWAVE_DATA_ASSET"
	UpdateDataAssetDetailsModelTypeRestDataAsset                UpdateDataAssetDetailsModelTypeEnum = "REST_DATA_ASSET"
)

var mappingUpdateDataAssetDetailsModelTypeEnum = map[string]UpdateDataAssetDetailsModelTypeEnum{
	"ORACLE_DATA_ASSET":                UpdateDataAssetDetailsModelTypeOracleDataAsset,
	"ORACLE_OBJECT_STORAGE_DATA_ASSET": UpdateDataAssetDetailsModelTypeOracleObjectStorageDataAsset,
	"ORACLE_ATP_DATA_ASSET":            UpdateDataAssetDetailsModelTypeOracleAtpDataAsset,
	"ORACLE_ADWC_DATA_ASSET":           UpdateDataAssetDetailsModelTypeOracleAdwcDataAsset,
	"MYSQL_DATA_ASSET":                 UpdateDataAssetDetailsModelTypeMysqlDataAsset,
	"GENERIC_JDBC_DATA_ASSET":          UpdateDataAssetDetailsModelTypeGenericJdbcDataAsset,
	"FUSION_APP_DATA_ASSET":            UpdateDataAssetDetailsModelTypeFusionAppDataAsset,
	"AMAZON_S3_DATA_ASSET":             UpdateDataAssetDetailsModelTypeAmazonS3DataAsset,
	"LAKE_DATA_ASSET":                  UpdateDataAssetDetailsModelTypeLakeDataAsset,
	"ORACLE_PEOPLESOFT_DATA_ASSET":     UpdateDataAssetDetailsModelTypeOraclePeoplesoftDataAsset,
	"ORACLE_SIEBEL_DATA_ASSET":         UpdateDataAssetDetailsModelTypeOracleSiebelDataAsset,
	"ORACLE_EBS_DATA_ASSET":            UpdateDataAssetDetailsModelTypeOracleEbsDataAsset,
	"HDFS_DATA_ASSET":                  UpdateDataAssetDetailsModelTypeHdfsDataAsset,
	"MYSQL_HEATWAVE_DATA_ASSET":        UpdateDataAssetDetailsModelTypeMysqlHeatwaveDataAsset,
	"REST_DATA_ASSET":                  UpdateDataAssetDetailsModelTypeRestDataAsset,
}

var mappingUpdateDataAssetDetailsModelTypeEnumLowerCase = map[string]UpdateDataAssetDetailsModelTypeEnum{
	"oracle_data_asset":                UpdateDataAssetDetailsModelTypeOracleDataAsset,
	"oracle_object_storage_data_asset": UpdateDataAssetDetailsModelTypeOracleObjectStorageDataAsset,
	"oracle_atp_data_asset":            UpdateDataAssetDetailsModelTypeOracleAtpDataAsset,
	"oracle_adwc_data_asset":           UpdateDataAssetDetailsModelTypeOracleAdwcDataAsset,
	"mysql_data_asset":                 UpdateDataAssetDetailsModelTypeMysqlDataAsset,
	"generic_jdbc_data_asset":          UpdateDataAssetDetailsModelTypeGenericJdbcDataAsset,
	"fusion_app_data_asset":            UpdateDataAssetDetailsModelTypeFusionAppDataAsset,
	"amazon_s3_data_asset":             UpdateDataAssetDetailsModelTypeAmazonS3DataAsset,
	"lake_data_asset":                  UpdateDataAssetDetailsModelTypeLakeDataAsset,
	"oracle_peoplesoft_data_asset":     UpdateDataAssetDetailsModelTypeOraclePeoplesoftDataAsset,
	"oracle_siebel_data_asset":         UpdateDataAssetDetailsModelTypeOracleSiebelDataAsset,
	"oracle_ebs_data_asset":            UpdateDataAssetDetailsModelTypeOracleEbsDataAsset,
	"hdfs_data_asset":                  UpdateDataAssetDetailsModelTypeHdfsDataAsset,
	"mysql_heatwave_data_asset":        UpdateDataAssetDetailsModelTypeMysqlHeatwaveDataAsset,
	"rest_data_asset":                  UpdateDataAssetDetailsModelTypeRestDataAsset,
}

// GetUpdateDataAssetDetailsModelTypeEnumValues Enumerates the set of values for UpdateDataAssetDetailsModelTypeEnum
func GetUpdateDataAssetDetailsModelTypeEnumValues() []UpdateDataAssetDetailsModelTypeEnum {
	values := make([]UpdateDataAssetDetailsModelTypeEnum, 0)
	for _, v := range mappingUpdateDataAssetDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDataAssetDetailsModelTypeEnumStringValues Enumerates the set of values in String for UpdateDataAssetDetailsModelTypeEnum
func GetUpdateDataAssetDetailsModelTypeEnumStringValues() []string {
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

// GetMappingUpdateDataAssetDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDataAssetDetailsModelTypeEnum(val string) (UpdateDataAssetDetailsModelTypeEnum, bool) {
	enum, ok := mappingUpdateDataAssetDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
