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

// ConnectionSummary The connection summary object.
type ConnectionSummary interface {

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

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	GetPrimarySchema() *Schema

	// The properties for the connection.
	GetConnectionProperties() []ConnectionProperty

	// The default property for the connection.
	GetIsDefault() *bool

	GetMetadata() *ObjectMetadata

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	GetKeyMap() map[string]string
}

type connectionsummary struct {
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
	KeyMap               map[string]string    `mandatory:"false" json:"keyMap"`
	ModelType            string               `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *connectionsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconnectionsummary connectionsummary
	s := struct {
		Model Unmarshalerconnectionsummary
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
	m.KeyMap = s.Model.KeyMap
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *connectionsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "GENERIC_JDBC_CONNECTION":
		mm := ConnectionSummaryFromJdbc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REST_NO_AUTH_CONNECTION":
		mm := ConnectionSummaryFromRestNoAuth{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_SIEBEL_CONNECTION":
		mm := ConnectionSummaryFromOracleSiebel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLEDB_CONNECTION":
		mm := ConnectionSummaryFromOracle{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL_HEATWAVE_CONNECTION":
		mm := ConnectionSummaryFromMySqlHeatWave{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMAZON_S3_CONNECTION":
		mm := ConnectionSummaryFromAmazonS3{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OAUTH2_CONNECTION":
		mm := ConnectionSummaryFromOAuth2{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_PEOPLESOFT_CONNECTION":
		mm := ConnectionSummaryFromOraclePeopleSoft{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_EBS_CONNECTION":
		mm := ConnectionSummaryFromOracleEbs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_CONNECTION":
		mm := ConnectionSummaryFromAdwc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL_CONNECTION":
		mm := ConnectionSummaryFromMySql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LAKE_CONNECTION":
		mm := ConnectionSummaryFromLake{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BIP_CONNECTION":
		mm := ConnectionSummaryFromBip{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HDFS_CONNECTION":
		mm := ConnectionSummaryFromHdfs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BICC_CONNECTION":
		mm := ConnectionSummaryFromBicc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ATP_CONNECTION":
		mm := ConnectionSummaryFromAtp{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REST_BASIC_AUTH_CONNECTION":
		mm := ConnectionSummaryFromRestBasicAuth{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_OBJECT_STORAGE_CONNECTION":
		mm := ConnectionSummaryFromObjectStorage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ConnectionSummary: %s.", m.ModelType)
		return *m, nil
	}
}

// GetKey returns Key
func (m connectionsummary) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m connectionsummary) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m connectionsummary) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m connectionsummary) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m connectionsummary) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m connectionsummary) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetObjectStatus returns ObjectStatus
func (m connectionsummary) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m connectionsummary) GetIdentifier() *string {
	return m.Identifier
}

// GetPrimarySchema returns PrimarySchema
func (m connectionsummary) GetPrimarySchema() *Schema {
	return m.PrimarySchema
}

// GetConnectionProperties returns ConnectionProperties
func (m connectionsummary) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

// GetIsDefault returns IsDefault
func (m connectionsummary) GetIsDefault() *bool {
	return m.IsDefault
}

// GetMetadata returns Metadata
func (m connectionsummary) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

// GetKeyMap returns KeyMap
func (m connectionsummary) GetKeyMap() map[string]string {
	return m.KeyMap
}

func (m connectionsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m connectionsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectionSummaryModelTypeEnum Enum with underlying type: string
type ConnectionSummaryModelTypeEnum string

// Set of constants representing the allowable values for ConnectionSummaryModelTypeEnum
const (
	ConnectionSummaryModelTypeOracleAdwcConnection          ConnectionSummaryModelTypeEnum = "ORACLE_ADWC_CONNECTION"
	ConnectionSummaryModelTypeOracleAtpConnection           ConnectionSummaryModelTypeEnum = "ORACLE_ATP_CONNECTION"
	ConnectionSummaryModelTypeOracleObjectStorageConnection ConnectionSummaryModelTypeEnum = "ORACLE_OBJECT_STORAGE_CONNECTION"
	ConnectionSummaryModelTypeOracledbConnection            ConnectionSummaryModelTypeEnum = "ORACLEDB_CONNECTION"
	ConnectionSummaryModelTypeMysqlConnection               ConnectionSummaryModelTypeEnum = "MYSQL_CONNECTION"
	ConnectionSummaryModelTypeGenericJdbcConnection         ConnectionSummaryModelTypeEnum = "GENERIC_JDBC_CONNECTION"
	ConnectionSummaryModelTypeBiccConnection                ConnectionSummaryModelTypeEnum = "BICC_CONNECTION"
	ConnectionSummaryModelTypeAmazonS3Connection            ConnectionSummaryModelTypeEnum = "AMAZON_S3_CONNECTION"
	ConnectionSummaryModelTypeBipConnection                 ConnectionSummaryModelTypeEnum = "BIP_CONNECTION"
	ConnectionSummaryModelTypeLakeConnection                ConnectionSummaryModelTypeEnum = "LAKE_CONNECTION"
	ConnectionSummaryModelTypeOraclePeoplesoftConnection    ConnectionSummaryModelTypeEnum = "ORACLE_PEOPLESOFT_CONNECTION"
	ConnectionSummaryModelTypeOracleEbsConnection           ConnectionSummaryModelTypeEnum = "ORACLE_EBS_CONNECTION"
	ConnectionSummaryModelTypeOracleSiebelConnection        ConnectionSummaryModelTypeEnum = "ORACLE_SIEBEL_CONNECTION"
	ConnectionSummaryModelTypeHdfsConnection                ConnectionSummaryModelTypeEnum = "HDFS_CONNECTION"
	ConnectionSummaryModelTypeMysqlHeatwaveConnection       ConnectionSummaryModelTypeEnum = "MYSQL_HEATWAVE_CONNECTION"
	ConnectionSummaryModelTypeRestNoAuthConnection          ConnectionSummaryModelTypeEnum = "REST_NO_AUTH_CONNECTION"
	ConnectionSummaryModelTypeRestBasicAuthConnection       ConnectionSummaryModelTypeEnum = "REST_BASIC_AUTH_CONNECTION"
	ConnectionSummaryModelTypeOauth2Connection              ConnectionSummaryModelTypeEnum = "OAUTH2_CONNECTION"
)

var mappingConnectionSummaryModelTypeEnum = map[string]ConnectionSummaryModelTypeEnum{
	"ORACLE_ADWC_CONNECTION":           ConnectionSummaryModelTypeOracleAdwcConnection,
	"ORACLE_ATP_CONNECTION":            ConnectionSummaryModelTypeOracleAtpConnection,
	"ORACLE_OBJECT_STORAGE_CONNECTION": ConnectionSummaryModelTypeOracleObjectStorageConnection,
	"ORACLEDB_CONNECTION":              ConnectionSummaryModelTypeOracledbConnection,
	"MYSQL_CONNECTION":                 ConnectionSummaryModelTypeMysqlConnection,
	"GENERIC_JDBC_CONNECTION":          ConnectionSummaryModelTypeGenericJdbcConnection,
	"BICC_CONNECTION":                  ConnectionSummaryModelTypeBiccConnection,
	"AMAZON_S3_CONNECTION":             ConnectionSummaryModelTypeAmazonS3Connection,
	"BIP_CONNECTION":                   ConnectionSummaryModelTypeBipConnection,
	"LAKE_CONNECTION":                  ConnectionSummaryModelTypeLakeConnection,
	"ORACLE_PEOPLESOFT_CONNECTION":     ConnectionSummaryModelTypeOraclePeoplesoftConnection,
	"ORACLE_EBS_CONNECTION":            ConnectionSummaryModelTypeOracleEbsConnection,
	"ORACLE_SIEBEL_CONNECTION":         ConnectionSummaryModelTypeOracleSiebelConnection,
	"HDFS_CONNECTION":                  ConnectionSummaryModelTypeHdfsConnection,
	"MYSQL_HEATWAVE_CONNECTION":        ConnectionSummaryModelTypeMysqlHeatwaveConnection,
	"REST_NO_AUTH_CONNECTION":          ConnectionSummaryModelTypeRestNoAuthConnection,
	"REST_BASIC_AUTH_CONNECTION":       ConnectionSummaryModelTypeRestBasicAuthConnection,
	"OAUTH2_CONNECTION":                ConnectionSummaryModelTypeOauth2Connection,
}

var mappingConnectionSummaryModelTypeEnumLowerCase = map[string]ConnectionSummaryModelTypeEnum{
	"oracle_adwc_connection":           ConnectionSummaryModelTypeOracleAdwcConnection,
	"oracle_atp_connection":            ConnectionSummaryModelTypeOracleAtpConnection,
	"oracle_object_storage_connection": ConnectionSummaryModelTypeOracleObjectStorageConnection,
	"oracledb_connection":              ConnectionSummaryModelTypeOracledbConnection,
	"mysql_connection":                 ConnectionSummaryModelTypeMysqlConnection,
	"generic_jdbc_connection":          ConnectionSummaryModelTypeGenericJdbcConnection,
	"bicc_connection":                  ConnectionSummaryModelTypeBiccConnection,
	"amazon_s3_connection":             ConnectionSummaryModelTypeAmazonS3Connection,
	"bip_connection":                   ConnectionSummaryModelTypeBipConnection,
	"lake_connection":                  ConnectionSummaryModelTypeLakeConnection,
	"oracle_peoplesoft_connection":     ConnectionSummaryModelTypeOraclePeoplesoftConnection,
	"oracle_ebs_connection":            ConnectionSummaryModelTypeOracleEbsConnection,
	"oracle_siebel_connection":         ConnectionSummaryModelTypeOracleSiebelConnection,
	"hdfs_connection":                  ConnectionSummaryModelTypeHdfsConnection,
	"mysql_heatwave_connection":        ConnectionSummaryModelTypeMysqlHeatwaveConnection,
	"rest_no_auth_connection":          ConnectionSummaryModelTypeRestNoAuthConnection,
	"rest_basic_auth_connection":       ConnectionSummaryModelTypeRestBasicAuthConnection,
	"oauth2_connection":                ConnectionSummaryModelTypeOauth2Connection,
}

// GetConnectionSummaryModelTypeEnumValues Enumerates the set of values for ConnectionSummaryModelTypeEnum
func GetConnectionSummaryModelTypeEnumValues() []ConnectionSummaryModelTypeEnum {
	values := make([]ConnectionSummaryModelTypeEnum, 0)
	for _, v := range mappingConnectionSummaryModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionSummaryModelTypeEnumStringValues Enumerates the set of values in String for ConnectionSummaryModelTypeEnum
func GetConnectionSummaryModelTypeEnumStringValues() []string {
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

// GetMappingConnectionSummaryModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionSummaryModelTypeEnum(val string) (ConnectionSummaryModelTypeEnum, bool) {
	enum, ok := mappingConnectionSummaryModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
