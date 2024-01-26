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

// ConnectionDetails The connection details for a data asset.
type ConnectionDetails interface {

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// User-defined description for the connection.
	GetDescription() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	GetPrimarySchema() *Schema

	// The properties for the connection.
	GetConnectionProperties() []ConnectionProperty

	// The default property for the connection.
	GetIsDefault() *bool

	GetMetadata() *ObjectMetadata
}

type connectiondetails struct {
	JsonData             []byte
	Key                  *string              `mandatory:"false" json:"key"`
	ModelVersion         *string              `mandatory:"false" json:"modelVersion"`
	ParentRef            *ParentReference     `mandatory:"false" json:"parentRef"`
	Name                 *string              `mandatory:"false" json:"name"`
	Description          *string              `mandatory:"false" json:"description"`
	ObjectVersion        *int                 `mandatory:"false" json:"objectVersion"`
	ObjectStatus         *int                 `mandatory:"false" json:"objectStatus"`
	Identifier           *string              `mandatory:"false" json:"identifier"`
	PrimarySchema        *Schema              `mandatory:"false" json:"primarySchema"`
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`
	IsDefault            *bool                `mandatory:"false" json:"isDefault"`
	Metadata             *ObjectMetadata      `mandatory:"false" json:"metadata"`
	ModelType            string               `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *connectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconnectiondetails connectiondetails
	s := struct {
		Model Unmarshalerconnectiondetails
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
	m.PrimarySchema = s.Model.PrimarySchema
	m.ConnectionProperties = s.Model.ConnectionProperties
	m.IsDefault = s.Model.IsDefault
	m.Metadata = s.Model.Metadata
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *connectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "GENERIC_JDBC_CONNECTION":
		mm := ConnectionFromJdbcDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REST_NO_AUTH_CONNECTION":
		mm := ConnectionFromRestNoAuthDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMAZON_S3_CONNECTION":
		mm := ConnectionFromAmazonS3Details{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_OBJECT_STORAGE_CONNECTION":
		mm := ConnectionFromObjectStorageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_SIEBEL_CONNECTION":
		mm := ConnectionFromOracleSiebelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HDFS_CONNECTION":
		mm := ConnectionFromHdfsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL_CONNECTION":
		mm := ConnectionFromMySqlDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL_HEATWAVE_CONNECTION":
		mm := ConnectionFromMySqlHeatWaveDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ATP_CONNECTION":
		mm := ConnectionFromAtpDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_PEOPLESOFT_CONNECTION":
		mm := ConnectionFromOraclePeopleSoftDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REST_BASIC_AUTH_CONNECTION":
		mm := ConnectionFromRestBasicAuthDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OAUTH2_CONNECTION":
		mm := ConnectionFromOAuth2Details{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BICC_CONNECTION":
		mm := ConnectionFromBiccDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_CONNECTION":
		mm := ConnectionFromAdwcDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLEDB_CONNECTION":
		mm := ConnectionFromOracleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_EBS_CONNECTION":
		mm := ConnectionFromOracleEbsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LAKE_CONNECTION":
		mm := ConnectionFromLakeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BIP_CONNECTION":
		mm := ConnectionFromBipDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ConnectionDetails: %s.", m.ModelType)
		return *m, nil
	}
}

// GetKey returns Key
func (m connectiondetails) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m connectiondetails) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m connectiondetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m connectiondetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m connectiondetails) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m connectiondetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetObjectStatus returns ObjectStatus
func (m connectiondetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m connectiondetails) GetIdentifier() *string {
	return m.Identifier
}

// GetPrimarySchema returns PrimarySchema
func (m connectiondetails) GetPrimarySchema() *Schema {
	return m.PrimarySchema
}

// GetConnectionProperties returns ConnectionProperties
func (m connectiondetails) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

// GetIsDefault returns IsDefault
func (m connectiondetails) GetIsDefault() *bool {
	return m.IsDefault
}

// GetMetadata returns Metadata
func (m connectiondetails) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m connectiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m connectiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectionDetailsModelTypeEnum Enum with underlying type: string
type ConnectionDetailsModelTypeEnum string

// Set of constants representing the allowable values for ConnectionDetailsModelTypeEnum
const (
	ConnectionDetailsModelTypeOracleAdwcConnection          ConnectionDetailsModelTypeEnum = "ORACLE_ADWC_CONNECTION"
	ConnectionDetailsModelTypeOracleAtpConnection           ConnectionDetailsModelTypeEnum = "ORACLE_ATP_CONNECTION"
	ConnectionDetailsModelTypeOracleObjectStorageConnection ConnectionDetailsModelTypeEnum = "ORACLE_OBJECT_STORAGE_CONNECTION"
	ConnectionDetailsModelTypeOracledbConnection            ConnectionDetailsModelTypeEnum = "ORACLEDB_CONNECTION"
	ConnectionDetailsModelTypeMysqlConnection               ConnectionDetailsModelTypeEnum = "MYSQL_CONNECTION"
	ConnectionDetailsModelTypeGenericJdbcConnection         ConnectionDetailsModelTypeEnum = "GENERIC_JDBC_CONNECTION"
	ConnectionDetailsModelTypeBiccConnection                ConnectionDetailsModelTypeEnum = "BICC_CONNECTION"
	ConnectionDetailsModelTypeAmazonS3Connection            ConnectionDetailsModelTypeEnum = "AMAZON_S3_CONNECTION"
	ConnectionDetailsModelTypeBipConnection                 ConnectionDetailsModelTypeEnum = "BIP_CONNECTION"
	ConnectionDetailsModelTypeLakeConnection                ConnectionDetailsModelTypeEnum = "LAKE_CONNECTION"
	ConnectionDetailsModelTypeOraclePeoplesoftConnection    ConnectionDetailsModelTypeEnum = "ORACLE_PEOPLESOFT_CONNECTION"
	ConnectionDetailsModelTypeOracleEbsConnection           ConnectionDetailsModelTypeEnum = "ORACLE_EBS_CONNECTION"
	ConnectionDetailsModelTypeOracleSiebelConnection        ConnectionDetailsModelTypeEnum = "ORACLE_SIEBEL_CONNECTION"
	ConnectionDetailsModelTypeHdfsConnection                ConnectionDetailsModelTypeEnum = "HDFS_CONNECTION"
	ConnectionDetailsModelTypeMysqlHeatwaveConnection       ConnectionDetailsModelTypeEnum = "MYSQL_HEATWAVE_CONNECTION"
	ConnectionDetailsModelTypeRestNoAuthConnection          ConnectionDetailsModelTypeEnum = "REST_NO_AUTH_CONNECTION"
	ConnectionDetailsModelTypeRestBasicAuthConnection       ConnectionDetailsModelTypeEnum = "REST_BASIC_AUTH_CONNECTION"
	ConnectionDetailsModelTypeOauth2Connection              ConnectionDetailsModelTypeEnum = "OAUTH2_CONNECTION"
)

var mappingConnectionDetailsModelTypeEnum = map[string]ConnectionDetailsModelTypeEnum{
	"ORACLE_ADWC_CONNECTION":           ConnectionDetailsModelTypeOracleAdwcConnection,
	"ORACLE_ATP_CONNECTION":            ConnectionDetailsModelTypeOracleAtpConnection,
	"ORACLE_OBJECT_STORAGE_CONNECTION": ConnectionDetailsModelTypeOracleObjectStorageConnection,
	"ORACLEDB_CONNECTION":              ConnectionDetailsModelTypeOracledbConnection,
	"MYSQL_CONNECTION":                 ConnectionDetailsModelTypeMysqlConnection,
	"GENERIC_JDBC_CONNECTION":          ConnectionDetailsModelTypeGenericJdbcConnection,
	"BICC_CONNECTION":                  ConnectionDetailsModelTypeBiccConnection,
	"AMAZON_S3_CONNECTION":             ConnectionDetailsModelTypeAmazonS3Connection,
	"BIP_CONNECTION":                   ConnectionDetailsModelTypeBipConnection,
	"LAKE_CONNECTION":                  ConnectionDetailsModelTypeLakeConnection,
	"ORACLE_PEOPLESOFT_CONNECTION":     ConnectionDetailsModelTypeOraclePeoplesoftConnection,
	"ORACLE_EBS_CONNECTION":            ConnectionDetailsModelTypeOracleEbsConnection,
	"ORACLE_SIEBEL_CONNECTION":         ConnectionDetailsModelTypeOracleSiebelConnection,
	"HDFS_CONNECTION":                  ConnectionDetailsModelTypeHdfsConnection,
	"MYSQL_HEATWAVE_CONNECTION":        ConnectionDetailsModelTypeMysqlHeatwaveConnection,
	"REST_NO_AUTH_CONNECTION":          ConnectionDetailsModelTypeRestNoAuthConnection,
	"REST_BASIC_AUTH_CONNECTION":       ConnectionDetailsModelTypeRestBasicAuthConnection,
	"OAUTH2_CONNECTION":                ConnectionDetailsModelTypeOauth2Connection,
}

var mappingConnectionDetailsModelTypeEnumLowerCase = map[string]ConnectionDetailsModelTypeEnum{
	"oracle_adwc_connection":           ConnectionDetailsModelTypeOracleAdwcConnection,
	"oracle_atp_connection":            ConnectionDetailsModelTypeOracleAtpConnection,
	"oracle_object_storage_connection": ConnectionDetailsModelTypeOracleObjectStorageConnection,
	"oracledb_connection":              ConnectionDetailsModelTypeOracledbConnection,
	"mysql_connection":                 ConnectionDetailsModelTypeMysqlConnection,
	"generic_jdbc_connection":          ConnectionDetailsModelTypeGenericJdbcConnection,
	"bicc_connection":                  ConnectionDetailsModelTypeBiccConnection,
	"amazon_s3_connection":             ConnectionDetailsModelTypeAmazonS3Connection,
	"bip_connection":                   ConnectionDetailsModelTypeBipConnection,
	"lake_connection":                  ConnectionDetailsModelTypeLakeConnection,
	"oracle_peoplesoft_connection":     ConnectionDetailsModelTypeOraclePeoplesoftConnection,
	"oracle_ebs_connection":            ConnectionDetailsModelTypeOracleEbsConnection,
	"oracle_siebel_connection":         ConnectionDetailsModelTypeOracleSiebelConnection,
	"hdfs_connection":                  ConnectionDetailsModelTypeHdfsConnection,
	"mysql_heatwave_connection":        ConnectionDetailsModelTypeMysqlHeatwaveConnection,
	"rest_no_auth_connection":          ConnectionDetailsModelTypeRestNoAuthConnection,
	"rest_basic_auth_connection":       ConnectionDetailsModelTypeRestBasicAuthConnection,
	"oauth2_connection":                ConnectionDetailsModelTypeOauth2Connection,
}

// GetConnectionDetailsModelTypeEnumValues Enumerates the set of values for ConnectionDetailsModelTypeEnum
func GetConnectionDetailsModelTypeEnumValues() []ConnectionDetailsModelTypeEnum {
	values := make([]ConnectionDetailsModelTypeEnum, 0)
	for _, v := range mappingConnectionDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionDetailsModelTypeEnumStringValues Enumerates the set of values in String for ConnectionDetailsModelTypeEnum
func GetConnectionDetailsModelTypeEnumStringValues() []string {
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
		"OAUTH2_CONNECTION",
	}
}

// GetMappingConnectionDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionDetailsModelTypeEnum(val string) (ConnectionDetailsModelTypeEnum, bool) {
	enum, ok := mappingConnectionDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
