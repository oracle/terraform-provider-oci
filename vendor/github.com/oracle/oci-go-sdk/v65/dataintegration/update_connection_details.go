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

// UpdateConnectionDetails Properties used in connection update operations.
type UpdateConnectionDetails interface {

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	GetKey() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// User-defined description for the connection.
	GetDescription() *string

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	// The properties for the connection.
	GetConnectionProperties() []ConnectionProperty

	GetRegistryMetadata() *RegistryMetadata
}

type updateconnectiondetails struct {
	JsonData             []byte
	ModelVersion         *string              `mandatory:"false" json:"modelVersion"`
	ParentRef            *ParentReference     `mandatory:"false" json:"parentRef"`
	Name                 *string              `mandatory:"false" json:"name"`
	Description          *string              `mandatory:"false" json:"description"`
	ObjectStatus         *int                 `mandatory:"false" json:"objectStatus"`
	Identifier           *string              `mandatory:"false" json:"identifier"`
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`
	RegistryMetadata     *RegistryMetadata    `mandatory:"false" json:"registryMetadata"`
	Key                  *string              `mandatory:"true" json:"key"`
	ObjectVersion        *int                 `mandatory:"true" json:"objectVersion"`
	ModelType            string               `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *updateconnectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateconnectiondetails updateconnectiondetails
	s := struct {
		Model Unmarshalerupdateconnectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ObjectVersion = s.Model.ObjectVersion
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ObjectStatus = s.Model.ObjectStatus
	m.Identifier = s.Model.Identifier
	m.ConnectionProperties = s.Model.ConnectionProperties
	m.RegistryMetadata = s.Model.RegistryMetadata
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateconnectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "LAKE_CONNECTION":
		mm := UpdateConnectionFromLake{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_EBS_CONNECTION":
		mm := UpdateConnectionFromOracleEbs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_OBJECT_STORAGE_CONNECTION":
		mm := UpdateConnectionFromObjectStorage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BICC_CONNECTION":
		mm := UpdateConnectionFromBicc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REST_NO_AUTH_CONNECTION":
		mm := UpdateConnectionFromRestNoAuth{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HDFS_CONNECTION":
		mm := UpdateConnectionFromHdfs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REST_BASIC_AUTH_CONNECTION":
		mm := UpdateConnectionFromRestBasicAuth{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_SIEBEL_CONNECTION":
		mm := UpdateConnectionFromOracleSiebel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL_HEATWAVE_CONNECTION":
		mm := UpdateConnectionFromMySqlHeatWave{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BIP_CONNECTION":
		mm := UpdateConnectionFromBip{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL_CONNECTION":
		mm := UpdateConnectionFromMySql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENERIC_JDBC_CONNECTION":
		mm := UpdateConnectionFromJdbc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMAZON_S3_CONNECTION":
		mm := UpdateConnectionFromAmazonS3{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ATP_CONNECTION":
		mm := UpdateConnectionFromAtp{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLEDB_CONNECTION":
		mm := UpdateConnectionFromOracle{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_CONNECTION":
		mm := UpdateConnectionFromAdwc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_PEOPLESOFT_CONNECTION":
		mm := UpdateConnectionFromOraclePeopleSoft{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateConnectionDetails: %s.", m.ModelType)
		return *m, nil
	}
}

// GetModelVersion returns ModelVersion
func (m updateconnectiondetails) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m updateconnectiondetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m updateconnectiondetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m updateconnectiondetails) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m updateconnectiondetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m updateconnectiondetails) GetIdentifier() *string {
	return m.Identifier
}

// GetConnectionProperties returns ConnectionProperties
func (m updateconnectiondetails) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

// GetRegistryMetadata returns RegistryMetadata
func (m updateconnectiondetails) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

// GetKey returns Key
func (m updateconnectiondetails) GetKey() *string {
	return m.Key
}

// GetObjectVersion returns ObjectVersion
func (m updateconnectiondetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

func (m updateconnectiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateconnectiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateConnectionDetailsModelTypeEnum Enum with underlying type: string
type UpdateConnectionDetailsModelTypeEnum string

// Set of constants representing the allowable values for UpdateConnectionDetailsModelTypeEnum
const (
	UpdateConnectionDetailsModelTypeOracleAdwcConnection          UpdateConnectionDetailsModelTypeEnum = "ORACLE_ADWC_CONNECTION"
	UpdateConnectionDetailsModelTypeOracleAtpConnection           UpdateConnectionDetailsModelTypeEnum = "ORACLE_ATP_CONNECTION"
	UpdateConnectionDetailsModelTypeOracleObjectStorageConnection UpdateConnectionDetailsModelTypeEnum = "ORACLE_OBJECT_STORAGE_CONNECTION"
	UpdateConnectionDetailsModelTypeOracledbConnection            UpdateConnectionDetailsModelTypeEnum = "ORACLEDB_CONNECTION"
	UpdateConnectionDetailsModelTypeMysqlConnection               UpdateConnectionDetailsModelTypeEnum = "MYSQL_CONNECTION"
	UpdateConnectionDetailsModelTypeGenericJdbcConnection         UpdateConnectionDetailsModelTypeEnum = "GENERIC_JDBC_CONNECTION"
	UpdateConnectionDetailsModelTypeBiccConnection                UpdateConnectionDetailsModelTypeEnum = "BICC_CONNECTION"
	UpdateConnectionDetailsModelTypeAmazonS3Connection            UpdateConnectionDetailsModelTypeEnum = "AMAZON_S3_CONNECTION"
	UpdateConnectionDetailsModelTypeBipConnection                 UpdateConnectionDetailsModelTypeEnum = "BIP_CONNECTION"
	UpdateConnectionDetailsModelTypeLakeConnection                UpdateConnectionDetailsModelTypeEnum = "LAKE_CONNECTION"
	UpdateConnectionDetailsModelTypeOraclePeoplesoftConnection    UpdateConnectionDetailsModelTypeEnum = "ORACLE_PEOPLESOFT_CONNECTION"
	UpdateConnectionDetailsModelTypeOracleEbsConnection           UpdateConnectionDetailsModelTypeEnum = "ORACLE_EBS_CONNECTION"
	UpdateConnectionDetailsModelTypeOracleSiebelConnection        UpdateConnectionDetailsModelTypeEnum = "ORACLE_SIEBEL_CONNECTION"
	UpdateConnectionDetailsModelTypeHdfsConnection                UpdateConnectionDetailsModelTypeEnum = "HDFS_CONNECTION"
	UpdateConnectionDetailsModelTypeMysqlHeatwaveConnection       UpdateConnectionDetailsModelTypeEnum = "MYSQL_HEATWAVE_CONNECTION"
	UpdateConnectionDetailsModelTypeRestNoAuthConnection          UpdateConnectionDetailsModelTypeEnum = "REST_NO_AUTH_CONNECTION"
	UpdateConnectionDetailsModelTypeRestBasicAuthConnection       UpdateConnectionDetailsModelTypeEnum = "REST_BASIC_AUTH_CONNECTION"
)

var mappingUpdateConnectionDetailsModelTypeEnum = map[string]UpdateConnectionDetailsModelTypeEnum{
	"ORACLE_ADWC_CONNECTION":           UpdateConnectionDetailsModelTypeOracleAdwcConnection,
	"ORACLE_ATP_CONNECTION":            UpdateConnectionDetailsModelTypeOracleAtpConnection,
	"ORACLE_OBJECT_STORAGE_CONNECTION": UpdateConnectionDetailsModelTypeOracleObjectStorageConnection,
	"ORACLEDB_CONNECTION":              UpdateConnectionDetailsModelTypeOracledbConnection,
	"MYSQL_CONNECTION":                 UpdateConnectionDetailsModelTypeMysqlConnection,
	"GENERIC_JDBC_CONNECTION":          UpdateConnectionDetailsModelTypeGenericJdbcConnection,
	"BICC_CONNECTION":                  UpdateConnectionDetailsModelTypeBiccConnection,
	"AMAZON_S3_CONNECTION":             UpdateConnectionDetailsModelTypeAmazonS3Connection,
	"BIP_CONNECTION":                   UpdateConnectionDetailsModelTypeBipConnection,
	"LAKE_CONNECTION":                  UpdateConnectionDetailsModelTypeLakeConnection,
	"ORACLE_PEOPLESOFT_CONNECTION":     UpdateConnectionDetailsModelTypeOraclePeoplesoftConnection,
	"ORACLE_EBS_CONNECTION":            UpdateConnectionDetailsModelTypeOracleEbsConnection,
	"ORACLE_SIEBEL_CONNECTION":         UpdateConnectionDetailsModelTypeOracleSiebelConnection,
	"HDFS_CONNECTION":                  UpdateConnectionDetailsModelTypeHdfsConnection,
	"MYSQL_HEATWAVE_CONNECTION":        UpdateConnectionDetailsModelTypeMysqlHeatwaveConnection,
	"REST_NO_AUTH_CONNECTION":          UpdateConnectionDetailsModelTypeRestNoAuthConnection,
	"REST_BASIC_AUTH_CONNECTION":       UpdateConnectionDetailsModelTypeRestBasicAuthConnection,
}

var mappingUpdateConnectionDetailsModelTypeEnumLowerCase = map[string]UpdateConnectionDetailsModelTypeEnum{
	"oracle_adwc_connection":           UpdateConnectionDetailsModelTypeOracleAdwcConnection,
	"oracle_atp_connection":            UpdateConnectionDetailsModelTypeOracleAtpConnection,
	"oracle_object_storage_connection": UpdateConnectionDetailsModelTypeOracleObjectStorageConnection,
	"oracledb_connection":              UpdateConnectionDetailsModelTypeOracledbConnection,
	"mysql_connection":                 UpdateConnectionDetailsModelTypeMysqlConnection,
	"generic_jdbc_connection":          UpdateConnectionDetailsModelTypeGenericJdbcConnection,
	"bicc_connection":                  UpdateConnectionDetailsModelTypeBiccConnection,
	"amazon_s3_connection":             UpdateConnectionDetailsModelTypeAmazonS3Connection,
	"bip_connection":                   UpdateConnectionDetailsModelTypeBipConnection,
	"lake_connection":                  UpdateConnectionDetailsModelTypeLakeConnection,
	"oracle_peoplesoft_connection":     UpdateConnectionDetailsModelTypeOraclePeoplesoftConnection,
	"oracle_ebs_connection":            UpdateConnectionDetailsModelTypeOracleEbsConnection,
	"oracle_siebel_connection":         UpdateConnectionDetailsModelTypeOracleSiebelConnection,
	"hdfs_connection":                  UpdateConnectionDetailsModelTypeHdfsConnection,
	"mysql_heatwave_connection":        UpdateConnectionDetailsModelTypeMysqlHeatwaveConnection,
	"rest_no_auth_connection":          UpdateConnectionDetailsModelTypeRestNoAuthConnection,
	"rest_basic_auth_connection":       UpdateConnectionDetailsModelTypeRestBasicAuthConnection,
}

// GetUpdateConnectionDetailsModelTypeEnumValues Enumerates the set of values for UpdateConnectionDetailsModelTypeEnum
func GetUpdateConnectionDetailsModelTypeEnumValues() []UpdateConnectionDetailsModelTypeEnum {
	values := make([]UpdateConnectionDetailsModelTypeEnum, 0)
	for _, v := range mappingUpdateConnectionDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateConnectionDetailsModelTypeEnumStringValues Enumerates the set of values in String for UpdateConnectionDetailsModelTypeEnum
func GetUpdateConnectionDetailsModelTypeEnumStringValues() []string {
	return []string{
		"ORACLE_ADWC_CONNECTION",
		"ORACLE_ATP_CONNECTION",
		"ORACLE_OBJECT_STORAGE_CONNECTION",
		"ORACLEDB_CONNECTION",
		"MYSQL_CONNECTION",
		"GENERIC_JDBC_CONNECTION",
		"BICC_CONNECTION",
		"AMAZON_S3_CONNECTION",
		"BIP_CONNECTION",
		"LAKE_CONNECTION",
		"ORACLE_PEOPLESOFT_CONNECTION",
		"ORACLE_EBS_CONNECTION",
		"ORACLE_SIEBEL_CONNECTION",
		"HDFS_CONNECTION",
		"MYSQL_HEATWAVE_CONNECTION",
		"REST_NO_AUTH_CONNECTION",
		"REST_BASIC_AUTH_CONNECTION",
	}
}

// GetMappingUpdateConnectionDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateConnectionDetailsModelTypeEnum(val string) (UpdateConnectionDetailsModelTypeEnum, bool) {
	enum, ok := mappingUpdateConnectionDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
