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

// CreateConnectionDetails Properties used in connection create operations.
type CreateConnectionDetails interface {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// User-defined description for the connection.
	GetDescription() *string

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// The properties for the connection.
	GetConnectionProperties() []ConnectionProperty

	GetRegistryMetadata() *RegistryMetadata
}

type createconnectiondetails struct {
	JsonData             []byte
	Name                 *string              `mandatory:"true" json:"name"`
	Identifier           *string              `mandatory:"true" json:"identifier"`
	Key                  *string              `mandatory:"false" json:"key"`
	ModelVersion         *string              `mandatory:"false" json:"modelVersion"`
	ParentRef            *ParentReference     `mandatory:"false" json:"parentRef"`
	Description          *string              `mandatory:"false" json:"description"`
	ObjectStatus         *int                 `mandatory:"false" json:"objectStatus"`
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`
	RegistryMetadata     *RegistryMetadata    `mandatory:"false" json:"registryMetadata"`
	ModelType            string               `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *createconnectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateconnectiondetails createconnectiondetails
	s := struct {
		Model Unmarshalercreateconnectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Identifier = s.Model.Identifier
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Description = s.Model.Description
	m.ObjectStatus = s.Model.ObjectStatus
	m.ConnectionProperties = s.Model.ConnectionProperties
	m.RegistryMetadata = s.Model.RegistryMetadata
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createconnectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "MYSQL_CONNECTION":
		mm := CreateConnectionFromMySql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMAZON_S3_CONNECTION":
		mm := CreateConnectionFromAmazonS3{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENERIC_JDBC_CONNECTION":
		mm := CreateConnectionFromJdbc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BICC_CONNECTION":
		mm := CreateConnectionFromBicc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ATP_CONNECTION":
		mm := CreateConnectionFromAtp{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_CONNECTION":
		mm := CreateConnectionFromAdwc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLEDB_CONNECTION":
		mm := CreateConnectionFromOracle{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_OBJECT_STORAGE_CONNECTION":
		mm := CreateConnectionFromObjectStorage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m createconnectiondetails) GetName() *string {
	return m.Name
}

//GetIdentifier returns Identifier
func (m createconnectiondetails) GetIdentifier() *string {
	return m.Identifier
}

//GetKey returns Key
func (m createconnectiondetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m createconnectiondetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m createconnectiondetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetDescription returns Description
func (m createconnectiondetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m createconnectiondetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetConnectionProperties returns ConnectionProperties
func (m createconnectiondetails) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

//GetRegistryMetadata returns RegistryMetadata
func (m createconnectiondetails) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m createconnectiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createconnectiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateConnectionDetailsModelTypeEnum Enum with underlying type: string
type CreateConnectionDetailsModelTypeEnum string

// Set of constants representing the allowable values for CreateConnectionDetailsModelTypeEnum
const (
	CreateConnectionDetailsModelTypeOracleAdwcConnection          CreateConnectionDetailsModelTypeEnum = "ORACLE_ADWC_CONNECTION"
	CreateConnectionDetailsModelTypeOracleAtpConnection           CreateConnectionDetailsModelTypeEnum = "ORACLE_ATP_CONNECTION"
	CreateConnectionDetailsModelTypeOracleObjectStorageConnection CreateConnectionDetailsModelTypeEnum = "ORACLE_OBJECT_STORAGE_CONNECTION"
	CreateConnectionDetailsModelTypeOracledbConnection            CreateConnectionDetailsModelTypeEnum = "ORACLEDB_CONNECTION"
	CreateConnectionDetailsModelTypeMysqlConnection               CreateConnectionDetailsModelTypeEnum = "MYSQL_CONNECTION"
	CreateConnectionDetailsModelTypeGenericJdbcConnection         CreateConnectionDetailsModelTypeEnum = "GENERIC_JDBC_CONNECTION"
	CreateConnectionDetailsModelTypeBiccConnection                CreateConnectionDetailsModelTypeEnum = "BICC_CONNECTION"
	CreateConnectionDetailsModelTypeAmazonS3Connection            CreateConnectionDetailsModelTypeEnum = "AMAZON_S3_CONNECTION"
)

var mappingCreateConnectionDetailsModelTypeEnum = map[string]CreateConnectionDetailsModelTypeEnum{
	"ORACLE_ADWC_CONNECTION":           CreateConnectionDetailsModelTypeOracleAdwcConnection,
	"ORACLE_ATP_CONNECTION":            CreateConnectionDetailsModelTypeOracleAtpConnection,
	"ORACLE_OBJECT_STORAGE_CONNECTION": CreateConnectionDetailsModelTypeOracleObjectStorageConnection,
	"ORACLEDB_CONNECTION":              CreateConnectionDetailsModelTypeOracledbConnection,
	"MYSQL_CONNECTION":                 CreateConnectionDetailsModelTypeMysqlConnection,
	"GENERIC_JDBC_CONNECTION":          CreateConnectionDetailsModelTypeGenericJdbcConnection,
	"BICC_CONNECTION":                  CreateConnectionDetailsModelTypeBiccConnection,
	"AMAZON_S3_CONNECTION":             CreateConnectionDetailsModelTypeAmazonS3Connection,
}

// GetCreateConnectionDetailsModelTypeEnumValues Enumerates the set of values for CreateConnectionDetailsModelTypeEnum
func GetCreateConnectionDetailsModelTypeEnumValues() []CreateConnectionDetailsModelTypeEnum {
	values := make([]CreateConnectionDetailsModelTypeEnum, 0)
	for _, v := range mappingCreateConnectionDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateConnectionDetailsModelTypeEnumStringValues Enumerates the set of values in String for CreateConnectionDetailsModelTypeEnum
func GetCreateConnectionDetailsModelTypeEnumStringValues() []string {
	return []string{
		"ORACLE_ADWC_CONNECTION",
		"ORACLE_ATP_CONNECTION",
		"ORACLE_OBJECT_STORAGE_CONNECTION",
		"ORACLEDB_CONNECTION",
		"MYSQL_CONNECTION",
		"GENERIC_JDBC_CONNECTION",
		"BICC_CONNECTION",
		"AMAZON_S3_CONNECTION",
	}
}

// GetMappingCreateConnectionDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateConnectionDetailsModelTypeEnum(val string) (CreateConnectionDetailsModelTypeEnum, bool) {
	mappingCreateConnectionDetailsModelTypeEnumIgnoreCase := make(map[string]CreateConnectionDetailsModelTypeEnum)
	for k, v := range mappingCreateConnectionDetailsModelTypeEnum {
		mappingCreateConnectionDetailsModelTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCreateConnectionDetailsModelTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
