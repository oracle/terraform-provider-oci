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

// Connection The connection object.
type Connection interface {

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	GetName() *string

	// Detailed description for the object.
	GetDescription() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	GetIdentifier() *string

	GetPrimarySchema() *Schema

	// The properties for the connection.
	GetConnectionProperties() []ConnectionProperty

	// The default property for the connection.
	GetIsDefault() *bool

	GetMetadata() *ObjectMetadata

	// A map, if provided key is replaced with generated key, this structure provides mapping between user provided key and generated key
	GetKeyMap() map[string]string
}

type connection struct {
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
func (m *connection) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconnection connection
	s := struct {
		Model Unmarshalerconnection
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
func (m *connection) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "ORACLE_OBJECT_STORAGE_CONNECTION":
		mm := ConnectionFromObjectStorage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_CONNECTION":
		mm := ConnectionFromAdwc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ATP_CONNECTION":
		mm := ConnectionFromAtp{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLEDB_CONNECTION":
		mm := ConnectionFromOracle{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m connection) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m connection) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m connection) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m connection) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m connection) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m connection) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m connection) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m connection) GetIdentifier() *string {
	return m.Identifier
}

//GetPrimarySchema returns PrimarySchema
func (m connection) GetPrimarySchema() *Schema {
	return m.PrimarySchema
}

//GetConnectionProperties returns ConnectionProperties
func (m connection) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

//GetIsDefault returns IsDefault
func (m connection) GetIsDefault() *bool {
	return m.IsDefault
}

//GetMetadata returns Metadata
func (m connection) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

//GetKeyMap returns KeyMap
func (m connection) GetKeyMap() map[string]string {
	return m.KeyMap
}

func (m connection) String() string {
	return common.PointerString(m)
}

// ConnectionModelTypeEnum Enum with underlying type: string
type ConnectionModelTypeEnum string

// Set of constants representing the allowable values for ConnectionModelTypeEnum
const (
	ConnectionModelTypeOracleAdwcConnection          ConnectionModelTypeEnum = "ORACLE_ADWC_CONNECTION"
	ConnectionModelTypeOracleAtpConnection           ConnectionModelTypeEnum = "ORACLE_ATP_CONNECTION"
	ConnectionModelTypeOracleObjectStorageConnection ConnectionModelTypeEnum = "ORACLE_OBJECT_STORAGE_CONNECTION"
	ConnectionModelTypeOracledbConnection            ConnectionModelTypeEnum = "ORACLEDB_CONNECTION"
)

var mappingConnectionModelType = map[string]ConnectionModelTypeEnum{
	"ORACLE_ADWC_CONNECTION":           ConnectionModelTypeOracleAdwcConnection,
	"ORACLE_ATP_CONNECTION":            ConnectionModelTypeOracleAtpConnection,
	"ORACLE_OBJECT_STORAGE_CONNECTION": ConnectionModelTypeOracleObjectStorageConnection,
	"ORACLEDB_CONNECTION":              ConnectionModelTypeOracledbConnection,
}

// GetConnectionModelTypeEnumValues Enumerates the set of values for ConnectionModelTypeEnum
func GetConnectionModelTypeEnumValues() []ConnectionModelTypeEnum {
	values := make([]ConnectionModelTypeEnum, 0)
	for _, v := range mappingConnectionModelType {
		values = append(values, v)
	}
	return values
}
