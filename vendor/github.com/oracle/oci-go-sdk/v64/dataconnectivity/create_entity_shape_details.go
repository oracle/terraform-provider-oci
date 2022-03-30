// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// CreateEntityShapeDetails The data entity shape object.
type CreateEntityShapeDetails interface {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// The object key.
	GetKey() *string

	// The object's model version.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// The external key for the object.
	GetExternalKey() *string

	GetShape() *Shape

	// The shape ID.
	GetShapeId() *string

	// The entity type.
	GetEntityType() CreateEntityShapeDetailsEntityTypeEnum

	// Specifies other type label.
	GetOtherTypeLabel() *string

	// An array of unique keys.
	GetUniqueKeys() []UniqueKey

	// An array of foreign keys.
	GetForeignKeys() []ForeignKey

	// The resource name.
	GetResourceName() *string

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	GetTypes() *TypeLibrary

	// Map<String, String> for entity properties
	GetEntityProperties() map[string]string
}

type createentityshapedetails struct {
	JsonData         []byte
	Name             *string                                `mandatory:"true" json:"name"`
	Key              *string                                `mandatory:"false" json:"key"`
	ModelVersion     *string                                `mandatory:"false" json:"modelVersion"`
	ParentRef        *ParentReference                       `mandatory:"false" json:"parentRef"`
	ObjectVersion    *int                                   `mandatory:"false" json:"objectVersion"`
	ExternalKey      *string                                `mandatory:"false" json:"externalKey"`
	Shape            *Shape                                 `mandatory:"false" json:"shape"`
	ShapeId          *string                                `mandatory:"false" json:"shapeId"`
	EntityType       CreateEntityShapeDetailsEntityTypeEnum `mandatory:"false" json:"entityType,omitempty"`
	OtherTypeLabel   *string                                `mandatory:"false" json:"otherTypeLabel"`
	UniqueKeys       json.RawMessage                        `mandatory:"false" json:"uniqueKeys"`
	ForeignKeys      []ForeignKey                           `mandatory:"false" json:"foreignKeys"`
	ResourceName     *string                                `mandatory:"false" json:"resourceName"`
	ObjectStatus     *int                                   `mandatory:"false" json:"objectStatus"`
	Identifier       *string                                `mandatory:"false" json:"identifier"`
	Types            *TypeLibrary                           `mandatory:"false" json:"types"`
	EntityProperties map[string]string                      `mandatory:"false" json:"entityProperties"`
	ModelType        string                                 `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *createentityshapedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateentityshapedetails createentityshapedetails
	s := struct {
		Model Unmarshalercreateentityshapedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.ObjectVersion = s.Model.ObjectVersion
	m.ExternalKey = s.Model.ExternalKey
	m.Shape = s.Model.Shape
	m.ShapeId = s.Model.ShapeId
	m.EntityType = s.Model.EntityType
	m.OtherTypeLabel = s.Model.OtherTypeLabel
	m.UniqueKeys = s.Model.UniqueKeys
	m.ForeignKeys = s.Model.ForeignKeys
	m.ResourceName = s.Model.ResourceName
	m.ObjectStatus = s.Model.ObjectStatus
	m.Identifier = s.Model.Identifier
	m.Types = s.Model.Types
	m.EntityProperties = s.Model.EntityProperties
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createentityshapedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "DATA_STORE_ENTITY":
		mm := CreateEntityShapeFromDataStore{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TABLE_ENTITY":
		mm := CreateEntityShapeFromTable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQL_ENTITY":
		mm := CreateEntityShapeFromSql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FILE_ENTITY":
		mm := CreateEntityShapeFromFile{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VIEW_ENTITY":
		mm := CreateEntityShapeFromView{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m createentityshapedetails) GetName() *string {
	return m.Name
}

//GetKey returns Key
func (m createentityshapedetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m createentityshapedetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m createentityshapedetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetObjectVersion returns ObjectVersion
func (m createentityshapedetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetExternalKey returns ExternalKey
func (m createentityshapedetails) GetExternalKey() *string {
	return m.ExternalKey
}

//GetShape returns Shape
func (m createentityshapedetails) GetShape() *Shape {
	return m.Shape
}

//GetShapeId returns ShapeId
func (m createentityshapedetails) GetShapeId() *string {
	return m.ShapeId
}

//GetEntityType returns EntityType
func (m createentityshapedetails) GetEntityType() CreateEntityShapeDetailsEntityTypeEnum {
	return m.EntityType
}

//GetOtherTypeLabel returns OtherTypeLabel
func (m createentityshapedetails) GetOtherTypeLabel() *string {
	return m.OtherTypeLabel
}

//GetUniqueKeys returns UniqueKeys
func (m createentityshapedetails) GetUniqueKeys() json.RawMessage {
	return m.UniqueKeys
}

//GetForeignKeys returns ForeignKeys
func (m createentityshapedetails) GetForeignKeys() []ForeignKey {
	return m.ForeignKeys
}

//GetResourceName returns ResourceName
func (m createentityshapedetails) GetResourceName() *string {
	return m.ResourceName
}

//GetObjectStatus returns ObjectStatus
func (m createentityshapedetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m createentityshapedetails) GetIdentifier() *string {
	return m.Identifier
}

//GetTypes returns Types
func (m createentityshapedetails) GetTypes() *TypeLibrary {
	return m.Types
}

//GetEntityProperties returns EntityProperties
func (m createentityshapedetails) GetEntityProperties() map[string]string {
	return m.EntityProperties
}

func (m createentityshapedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createentityshapedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateEntityShapeDetailsEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetCreateEntityShapeDetailsEntityTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateEntityShapeDetailsEntityTypeEnum Enum with underlying type: string
type CreateEntityShapeDetailsEntityTypeEnum string

// Set of constants representing the allowable values for CreateEntityShapeDetailsEntityTypeEnum
const (
	CreateEntityShapeDetailsEntityTypeTable CreateEntityShapeDetailsEntityTypeEnum = "TABLE"
	CreateEntityShapeDetailsEntityTypeView  CreateEntityShapeDetailsEntityTypeEnum = "VIEW"
	CreateEntityShapeDetailsEntityTypeFile  CreateEntityShapeDetailsEntityTypeEnum = "FILE"
	CreateEntityShapeDetailsEntityTypeSql   CreateEntityShapeDetailsEntityTypeEnum = "SQL"
)

var mappingCreateEntityShapeDetailsEntityTypeEnum = map[string]CreateEntityShapeDetailsEntityTypeEnum{
	"TABLE": CreateEntityShapeDetailsEntityTypeTable,
	"VIEW":  CreateEntityShapeDetailsEntityTypeView,
	"FILE":  CreateEntityShapeDetailsEntityTypeFile,
	"SQL":   CreateEntityShapeDetailsEntityTypeSql,
}

var mappingCreateEntityShapeDetailsEntityTypeEnumLowerCase = map[string]CreateEntityShapeDetailsEntityTypeEnum{
	"table": CreateEntityShapeDetailsEntityTypeTable,
	"view":  CreateEntityShapeDetailsEntityTypeView,
	"file":  CreateEntityShapeDetailsEntityTypeFile,
	"sql":   CreateEntityShapeDetailsEntityTypeSql,
}

// GetCreateEntityShapeDetailsEntityTypeEnumValues Enumerates the set of values for CreateEntityShapeDetailsEntityTypeEnum
func GetCreateEntityShapeDetailsEntityTypeEnumValues() []CreateEntityShapeDetailsEntityTypeEnum {
	values := make([]CreateEntityShapeDetailsEntityTypeEnum, 0)
	for _, v := range mappingCreateEntityShapeDetailsEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateEntityShapeDetailsEntityTypeEnumStringValues Enumerates the set of values in String for CreateEntityShapeDetailsEntityTypeEnum
func GetCreateEntityShapeDetailsEntityTypeEnumStringValues() []string {
	return []string{
		"TABLE",
		"VIEW",
		"FILE",
		"SQL",
	}
}

// GetMappingCreateEntityShapeDetailsEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateEntityShapeDetailsEntityTypeEnum(val string) (CreateEntityShapeDetailsEntityTypeEnum, bool) {
	enum, ok := mappingCreateEntityShapeDetailsEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateEntityShapeDetailsModelTypeEnum Enum with underlying type: string
type CreateEntityShapeDetailsModelTypeEnum string

// Set of constants representing the allowable values for CreateEntityShapeDetailsModelTypeEnum
const (
	CreateEntityShapeDetailsModelTypeViewEntity      CreateEntityShapeDetailsModelTypeEnum = "VIEW_ENTITY"
	CreateEntityShapeDetailsModelTypeTableEntity     CreateEntityShapeDetailsModelTypeEnum = "TABLE_ENTITY"
	CreateEntityShapeDetailsModelTypeFileEntity      CreateEntityShapeDetailsModelTypeEnum = "FILE_ENTITY"
	CreateEntityShapeDetailsModelTypeDataStoreEntity CreateEntityShapeDetailsModelTypeEnum = "DATA_STORE_ENTITY"
	CreateEntityShapeDetailsModelTypeSqlEntity       CreateEntityShapeDetailsModelTypeEnum = "SQL_ENTITY"
)

var mappingCreateEntityShapeDetailsModelTypeEnum = map[string]CreateEntityShapeDetailsModelTypeEnum{
	"VIEW_ENTITY":       CreateEntityShapeDetailsModelTypeViewEntity,
	"TABLE_ENTITY":      CreateEntityShapeDetailsModelTypeTableEntity,
	"FILE_ENTITY":       CreateEntityShapeDetailsModelTypeFileEntity,
	"DATA_STORE_ENTITY": CreateEntityShapeDetailsModelTypeDataStoreEntity,
	"SQL_ENTITY":        CreateEntityShapeDetailsModelTypeSqlEntity,
}

var mappingCreateEntityShapeDetailsModelTypeEnumLowerCase = map[string]CreateEntityShapeDetailsModelTypeEnum{
	"view_entity":       CreateEntityShapeDetailsModelTypeViewEntity,
	"table_entity":      CreateEntityShapeDetailsModelTypeTableEntity,
	"file_entity":       CreateEntityShapeDetailsModelTypeFileEntity,
	"data_store_entity": CreateEntityShapeDetailsModelTypeDataStoreEntity,
	"sql_entity":        CreateEntityShapeDetailsModelTypeSqlEntity,
}

// GetCreateEntityShapeDetailsModelTypeEnumValues Enumerates the set of values for CreateEntityShapeDetailsModelTypeEnum
func GetCreateEntityShapeDetailsModelTypeEnumValues() []CreateEntityShapeDetailsModelTypeEnum {
	values := make([]CreateEntityShapeDetailsModelTypeEnum, 0)
	for _, v := range mappingCreateEntityShapeDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateEntityShapeDetailsModelTypeEnumStringValues Enumerates the set of values in String for CreateEntityShapeDetailsModelTypeEnum
func GetCreateEntityShapeDetailsModelTypeEnumStringValues() []string {
	return []string{
		"VIEW_ENTITY",
		"TABLE_ENTITY",
		"FILE_ENTITY",
		"DATA_STORE_ENTITY",
		"SQL_ENTITY",
	}
}

// GetMappingCreateEntityShapeDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateEntityShapeDetailsModelTypeEnum(val string) (CreateEntityShapeDetailsModelTypeEnum, bool) {
	enum, ok := mappingCreateEntityShapeDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
