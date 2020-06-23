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

// UpdateConnectionDetails Properties used in connection update operations.
type UpdateConnectionDetails interface {

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	GetKey() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	GetName() *string

	// Detailed description for the object.
	GetDescription() *string

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	GetIdentifier() *string

	// The properties for the connection.
	GetConnectionProperties() []ConnectionProperty

	GetRegistryMetadata() *RegistryMetadata
}

type updateconnectiondetails struct {
	JsonData             []byte
	Key                  *string              `mandatory:"true" json:"key"`
	ObjectVersion        *int                 `mandatory:"true" json:"objectVersion"`
	ModelVersion         *string              `mandatory:"false" json:"modelVersion"`
	ParentRef            *ParentReference     `mandatory:"false" json:"parentRef"`
	Name                 *string              `mandatory:"false" json:"name"`
	Description          *string              `mandatory:"false" json:"description"`
	ObjectStatus         *int                 `mandatory:"false" json:"objectStatus"`
	Identifier           *string              `mandatory:"false" json:"identifier"`
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`
	RegistryMetadata     *RegistryMetadata    `mandatory:"false" json:"registryMetadata"`
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
	case "ORACLE_OBJECT_STORAGE_CONNECTION":
		mm := UpdateConnectionFromObjectStorage{}
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
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m updateconnectiondetails) GetKey() *string {
	return m.Key
}

//GetObjectVersion returns ObjectVersion
func (m updateconnectiondetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetModelVersion returns ModelVersion
func (m updateconnectiondetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m updateconnectiondetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m updateconnectiondetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m updateconnectiondetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m updateconnectiondetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m updateconnectiondetails) GetIdentifier() *string {
	return m.Identifier
}

//GetConnectionProperties returns ConnectionProperties
func (m updateconnectiondetails) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

//GetRegistryMetadata returns RegistryMetadata
func (m updateconnectiondetails) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m updateconnectiondetails) String() string {
	return common.PointerString(m)
}

// UpdateConnectionDetailsModelTypeEnum Enum with underlying type: string
type UpdateConnectionDetailsModelTypeEnum string

// Set of constants representing the allowable values for UpdateConnectionDetailsModelTypeEnum
const (
	UpdateConnectionDetailsModelTypeOracleAdwcConnection          UpdateConnectionDetailsModelTypeEnum = "ORACLE_ADWC_CONNECTION"
	UpdateConnectionDetailsModelTypeOracleAtpConnection           UpdateConnectionDetailsModelTypeEnum = "ORACLE_ATP_CONNECTION"
	UpdateConnectionDetailsModelTypeOracleObjectStorageConnection UpdateConnectionDetailsModelTypeEnum = "ORACLE_OBJECT_STORAGE_CONNECTION"
	UpdateConnectionDetailsModelTypeOracledbConnection            UpdateConnectionDetailsModelTypeEnum = "ORACLEDB_CONNECTION"
)

var mappingUpdateConnectionDetailsModelType = map[string]UpdateConnectionDetailsModelTypeEnum{
	"ORACLE_ADWC_CONNECTION":           UpdateConnectionDetailsModelTypeOracleAdwcConnection,
	"ORACLE_ATP_CONNECTION":            UpdateConnectionDetailsModelTypeOracleAtpConnection,
	"ORACLE_OBJECT_STORAGE_CONNECTION": UpdateConnectionDetailsModelTypeOracleObjectStorageConnection,
	"ORACLEDB_CONNECTION":              UpdateConnectionDetailsModelTypeOracledbConnection,
}

// GetUpdateConnectionDetailsModelTypeEnumValues Enumerates the set of values for UpdateConnectionDetailsModelTypeEnum
func GetUpdateConnectionDetailsModelTypeEnumValues() []UpdateConnectionDetailsModelTypeEnum {
	values := make([]UpdateConnectionDetailsModelTypeEnum, 0)
	for _, v := range mappingUpdateConnectionDetailsModelType {
		values = append(values, v)
	}
	return values
}
