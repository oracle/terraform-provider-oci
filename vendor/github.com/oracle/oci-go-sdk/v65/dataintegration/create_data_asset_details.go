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

// CreateDataAssetDetails Properties used in data asset update operations.
type CreateDataAssetDetails interface {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	// Currently not used on data asset creation. Reserved for future.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	// User-defined description of the data asset.
	GetDescription() *string

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// The external key for the object.
	GetExternalKey() *string

	// Additional properties for the data asset.
	GetAssetProperties() map[string]string

	GetRegistryMetadata() *RegistryMetadata
}

type createdataassetdetails struct {
	JsonData         []byte
	Key              *string           `mandatory:"false" json:"key"`
	ModelVersion     *string           `mandatory:"false" json:"modelVersion"`
	Description      *string           `mandatory:"false" json:"description"`
	ObjectStatus     *int              `mandatory:"false" json:"objectStatus"`
	ExternalKey      *string           `mandatory:"false" json:"externalKey"`
	AssetProperties  map[string]string `mandatory:"false" json:"assetProperties"`
	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
	Name             *string           `mandatory:"true" json:"name"`
	Identifier       *string           `mandatory:"true" json:"identifier"`
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
	case "HDFS_DATA_ASSET":
		mm := CreateDataAssetFromHdfs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LAKE_DATA_ASSET":
		mm := CreateDataAssetFromLake{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL_HEATWAVE_DATA_ASSET":
		mm := CreateDataAssetFromMySqlHeatWave{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMAZON_S3_DATA_ASSET":
		mm := CreateDataAssetFromAmazonS3{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REST_DATA_ASSET":
		mm := CreateDataAssetFromRest{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_SIEBEL_DATA_ASSET":
		mm := CreateDataAssetFromOracleSiebel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENERIC_JDBC_DATA_ASSET":
		mm := CreateDataAssetFromJdbc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL_DATA_ASSET":
		mm := CreateDataAssetFromMySql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_DATA_ASSET":
		mm := CreateDataAssetFromOracle{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_DATA_ASSET":
		mm := CreateDataAssetFromAdwc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_EBS_DATA_ASSET":
		mm := CreateDataAssetFromOracleEbs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_PEOPLESOFT_DATA_ASSET":
		mm := CreateDataAssetFromOraclePeopleSoft{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FUSION_APP_DATA_ASSET":
		mm := CreateDataAssetFromFusionApp{}
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
		common.Logf("Recieved unsupported enum value for CreateDataAssetDetails: %s.", m.ModelType)
		return *m, nil
	}
}

// GetKey returns Key
func (m createdataassetdetails) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m createdataassetdetails) GetModelVersion() *string {
	return m.ModelVersion
}

// GetDescription returns Description
func (m createdataassetdetails) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m createdataassetdetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetExternalKey returns ExternalKey
func (m createdataassetdetails) GetExternalKey() *string {
	return m.ExternalKey
}

// GetAssetProperties returns AssetProperties
func (m createdataassetdetails) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

// GetRegistryMetadata returns RegistryMetadata
func (m createdataassetdetails) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

// GetName returns Name
func (m createdataassetdetails) GetName() *string {
	return m.Name
}

// GetIdentifier returns Identifier
func (m createdataassetdetails) GetIdentifier() *string {
	return m.Identifier
}

func (m createdataassetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdataassetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDataAssetDetailsModelTypeEnum Enum with underlying type: string
type CreateDataAssetDetailsModelTypeEnum string

// Set of constants representing the allowable values for CreateDataAssetDetailsModelTypeEnum
const (
	CreateDataAssetDetailsModelTypeOracleDataAsset              CreateDataAssetDetailsModelTypeEnum = "ORACLE_DATA_ASSET"
	CreateDataAssetDetailsModelTypeOracleObjectStorageDataAsset CreateDataAssetDetailsModelTypeEnum = "ORACLE_OBJECT_STORAGE_DATA_ASSET"
	CreateDataAssetDetailsModelTypeOracleAtpDataAsset           CreateDataAssetDetailsModelTypeEnum = "ORACLE_ATP_DATA_ASSET"
	CreateDataAssetDetailsModelTypeOracleAdwcDataAsset          CreateDataAssetDetailsModelTypeEnum = "ORACLE_ADWC_DATA_ASSET"
	CreateDataAssetDetailsModelTypeMysqlDataAsset               CreateDataAssetDetailsModelTypeEnum = "MYSQL_DATA_ASSET"
	CreateDataAssetDetailsModelTypeGenericJdbcDataAsset         CreateDataAssetDetailsModelTypeEnum = "GENERIC_JDBC_DATA_ASSET"
	CreateDataAssetDetailsModelTypeFusionAppDataAsset           CreateDataAssetDetailsModelTypeEnum = "FUSION_APP_DATA_ASSET"
	CreateDataAssetDetailsModelTypeAmazonS3DataAsset            CreateDataAssetDetailsModelTypeEnum = "AMAZON_S3_DATA_ASSET"
	CreateDataAssetDetailsModelTypeLakeDataAsset                CreateDataAssetDetailsModelTypeEnum = "LAKE_DATA_ASSET"
	CreateDataAssetDetailsModelTypeOraclePeoplesoftDataAsset    CreateDataAssetDetailsModelTypeEnum = "ORACLE_PEOPLESOFT_DATA_ASSET"
	CreateDataAssetDetailsModelTypeOracleSiebelDataAsset        CreateDataAssetDetailsModelTypeEnum = "ORACLE_SIEBEL_DATA_ASSET"
	CreateDataAssetDetailsModelTypeOracleEbsDataAsset           CreateDataAssetDetailsModelTypeEnum = "ORACLE_EBS_DATA_ASSET"
	CreateDataAssetDetailsModelTypeHdfsDataAsset                CreateDataAssetDetailsModelTypeEnum = "HDFS_DATA_ASSET"
	CreateDataAssetDetailsModelTypeMysqlHeatwaveDataAsset       CreateDataAssetDetailsModelTypeEnum = "MYSQL_HEATWAVE_DATA_ASSET"
	CreateDataAssetDetailsModelTypeRestDataAsset                CreateDataAssetDetailsModelTypeEnum = "REST_DATA_ASSET"
)

var mappingCreateDataAssetDetailsModelTypeEnum = map[string]CreateDataAssetDetailsModelTypeEnum{
	"ORACLE_DATA_ASSET":                CreateDataAssetDetailsModelTypeOracleDataAsset,
	"ORACLE_OBJECT_STORAGE_DATA_ASSET": CreateDataAssetDetailsModelTypeOracleObjectStorageDataAsset,
	"ORACLE_ATP_DATA_ASSET":            CreateDataAssetDetailsModelTypeOracleAtpDataAsset,
	"ORACLE_ADWC_DATA_ASSET":           CreateDataAssetDetailsModelTypeOracleAdwcDataAsset,
	"MYSQL_DATA_ASSET":                 CreateDataAssetDetailsModelTypeMysqlDataAsset,
	"GENERIC_JDBC_DATA_ASSET":          CreateDataAssetDetailsModelTypeGenericJdbcDataAsset,
	"FUSION_APP_DATA_ASSET":            CreateDataAssetDetailsModelTypeFusionAppDataAsset,
	"AMAZON_S3_DATA_ASSET":             CreateDataAssetDetailsModelTypeAmazonS3DataAsset,
	"LAKE_DATA_ASSET":                  CreateDataAssetDetailsModelTypeLakeDataAsset,
	"ORACLE_PEOPLESOFT_DATA_ASSET":     CreateDataAssetDetailsModelTypeOraclePeoplesoftDataAsset,
	"ORACLE_SIEBEL_DATA_ASSET":         CreateDataAssetDetailsModelTypeOracleSiebelDataAsset,
	"ORACLE_EBS_DATA_ASSET":            CreateDataAssetDetailsModelTypeOracleEbsDataAsset,
	"HDFS_DATA_ASSET":                  CreateDataAssetDetailsModelTypeHdfsDataAsset,
	"MYSQL_HEATWAVE_DATA_ASSET":        CreateDataAssetDetailsModelTypeMysqlHeatwaveDataAsset,
	"REST_DATA_ASSET":                  CreateDataAssetDetailsModelTypeRestDataAsset,
}

var mappingCreateDataAssetDetailsModelTypeEnumLowerCase = map[string]CreateDataAssetDetailsModelTypeEnum{
	"oracle_data_asset":                CreateDataAssetDetailsModelTypeOracleDataAsset,
	"oracle_object_storage_data_asset": CreateDataAssetDetailsModelTypeOracleObjectStorageDataAsset,
	"oracle_atp_data_asset":            CreateDataAssetDetailsModelTypeOracleAtpDataAsset,
	"oracle_adwc_data_asset":           CreateDataAssetDetailsModelTypeOracleAdwcDataAsset,
	"mysql_data_asset":                 CreateDataAssetDetailsModelTypeMysqlDataAsset,
	"generic_jdbc_data_asset":          CreateDataAssetDetailsModelTypeGenericJdbcDataAsset,
	"fusion_app_data_asset":            CreateDataAssetDetailsModelTypeFusionAppDataAsset,
	"amazon_s3_data_asset":             CreateDataAssetDetailsModelTypeAmazonS3DataAsset,
	"lake_data_asset":                  CreateDataAssetDetailsModelTypeLakeDataAsset,
	"oracle_peoplesoft_data_asset":     CreateDataAssetDetailsModelTypeOraclePeoplesoftDataAsset,
	"oracle_siebel_data_asset":         CreateDataAssetDetailsModelTypeOracleSiebelDataAsset,
	"oracle_ebs_data_asset":            CreateDataAssetDetailsModelTypeOracleEbsDataAsset,
	"hdfs_data_asset":                  CreateDataAssetDetailsModelTypeHdfsDataAsset,
	"mysql_heatwave_data_asset":        CreateDataAssetDetailsModelTypeMysqlHeatwaveDataAsset,
	"rest_data_asset":                  CreateDataAssetDetailsModelTypeRestDataAsset,
}

// GetCreateDataAssetDetailsModelTypeEnumValues Enumerates the set of values for CreateDataAssetDetailsModelTypeEnum
func GetCreateDataAssetDetailsModelTypeEnumValues() []CreateDataAssetDetailsModelTypeEnum {
	values := make([]CreateDataAssetDetailsModelTypeEnum, 0)
	for _, v := range mappingCreateDataAssetDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDataAssetDetailsModelTypeEnumStringValues Enumerates the set of values in String for CreateDataAssetDetailsModelTypeEnum
func GetCreateDataAssetDetailsModelTypeEnumStringValues() []string {
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

// GetMappingCreateDataAssetDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDataAssetDetailsModelTypeEnum(val string) (CreateDataAssetDetailsModelTypeEnum, bool) {
	enum, ok := mappingCreateDataAssetDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
