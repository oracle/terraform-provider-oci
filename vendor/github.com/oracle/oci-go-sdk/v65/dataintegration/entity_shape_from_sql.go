// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// EntityShapeFromSql The SQL entity details.
type EntityShapeFromSql struct {
	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The external key for the object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	Shape *Shape `mandatory:"false" json:"shape"`

	// The shape ID.
	ShapeId *string `mandatory:"false" json:"shapeId"`

	// Specifies other type label.
	OtherTypeLabel *string `mandatory:"false" json:"otherTypeLabel"`

	// An array of unique keys.
	UniqueKeys []UniqueKey `mandatory:"false" json:"uniqueKeys"`

	// An array of foreign keys.
	ForeignKeys []ForeignKey `mandatory:"false" json:"foreignKeys"`

	// The resource name.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	DataFormat *DataFormat `mandatory:"false" json:"dataFormat"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// sqlQuery
	SqlQuery *string `mandatory:"false" json:"sqlQuery"`

	// The entity type.
	EntityType EntityShapeFromSqlEntityTypeEnum `mandatory:"false" json:"entityType,omitempty"`
}

// GetMetadata returns Metadata
func (m EntityShapeFromSql) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m EntityShapeFromSql) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EntityShapeFromSql) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEntityShapeFromSqlEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetEntityShapeFromSqlEntityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EntityShapeFromSql) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEntityShapeFromSql EntityShapeFromSql
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeEntityShapeFromSql
	}{
		"SQL_ENTITY",
		(MarshalTypeEntityShapeFromSql)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *EntityShapeFromSql) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Metadata       *ObjectMetadata                  `json:"metadata"`
		Key            *string                          `json:"key"`
		ModelVersion   *string                          `json:"modelVersion"`
		ParentRef      *ParentReference                 `json:"parentRef"`
		Name           *string                          `json:"name"`
		Description    *string                          `json:"description"`
		ObjectVersion  *int                             `json:"objectVersion"`
		ExternalKey    *string                          `json:"externalKey"`
		Shape          *Shape                           `json:"shape"`
		ShapeId        *string                          `json:"shapeId"`
		EntityType     EntityShapeFromSqlEntityTypeEnum `json:"entityType"`
		OtherTypeLabel *string                          `json:"otherTypeLabel"`
		UniqueKeys     []uniquekey                      `json:"uniqueKeys"`
		ForeignKeys    []ForeignKey                     `json:"foreignKeys"`
		ResourceName   *string                          `json:"resourceName"`
		DataFormat     *DataFormat                      `json:"dataFormat"`
		ObjectStatus   *int                             `json:"objectStatus"`
		Identifier     *string                          `json:"identifier"`
		SqlQuery       *string                          `json:"sqlQuery"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Metadata = model.Metadata

	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.Description = model.Description

	m.ObjectVersion = model.ObjectVersion

	m.ExternalKey = model.ExternalKey

	m.Shape = model.Shape

	m.ShapeId = model.ShapeId

	m.EntityType = model.EntityType

	m.OtherTypeLabel = model.OtherTypeLabel

	m.UniqueKeys = make([]UniqueKey, len(model.UniqueKeys))
	for i, n := range model.UniqueKeys {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.UniqueKeys[i] = nn.(UniqueKey)
		} else {
			m.UniqueKeys[i] = nil
		}
	}
	m.ForeignKeys = make([]ForeignKey, len(model.ForeignKeys))
	copy(m.ForeignKeys, model.ForeignKeys)
	m.ResourceName = model.ResourceName

	m.DataFormat = model.DataFormat

	m.ObjectStatus = model.ObjectStatus

	m.Identifier = model.Identifier

	m.SqlQuery = model.SqlQuery

	return
}

// EntityShapeFromSqlEntityTypeEnum Enum with underlying type: string
type EntityShapeFromSqlEntityTypeEnum string

// Set of constants representing the allowable values for EntityShapeFromSqlEntityTypeEnum
const (
	EntityShapeFromSqlEntityTypeTable EntityShapeFromSqlEntityTypeEnum = "TABLE"
	EntityShapeFromSqlEntityTypeView  EntityShapeFromSqlEntityTypeEnum = "VIEW"
	EntityShapeFromSqlEntityTypeFile  EntityShapeFromSqlEntityTypeEnum = "FILE"
	EntityShapeFromSqlEntityTypeSql   EntityShapeFromSqlEntityTypeEnum = "SQL"
)

var mappingEntityShapeFromSqlEntityTypeEnum = map[string]EntityShapeFromSqlEntityTypeEnum{
	"TABLE": EntityShapeFromSqlEntityTypeTable,
	"VIEW":  EntityShapeFromSqlEntityTypeView,
	"FILE":  EntityShapeFromSqlEntityTypeFile,
	"SQL":   EntityShapeFromSqlEntityTypeSql,
}

var mappingEntityShapeFromSqlEntityTypeEnumLowerCase = map[string]EntityShapeFromSqlEntityTypeEnum{
	"table": EntityShapeFromSqlEntityTypeTable,
	"view":  EntityShapeFromSqlEntityTypeView,
	"file":  EntityShapeFromSqlEntityTypeFile,
	"sql":   EntityShapeFromSqlEntityTypeSql,
}

// GetEntityShapeFromSqlEntityTypeEnumValues Enumerates the set of values for EntityShapeFromSqlEntityTypeEnum
func GetEntityShapeFromSqlEntityTypeEnumValues() []EntityShapeFromSqlEntityTypeEnum {
	values := make([]EntityShapeFromSqlEntityTypeEnum, 0)
	for _, v := range mappingEntityShapeFromSqlEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEntityShapeFromSqlEntityTypeEnumStringValues Enumerates the set of values in String for EntityShapeFromSqlEntityTypeEnum
func GetEntityShapeFromSqlEntityTypeEnumStringValues() []string {
	return []string{
		"TABLE",
		"VIEW",
		"FILE",
		"SQL",
	}
}

// GetMappingEntityShapeFromSqlEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEntityShapeFromSqlEntityTypeEnum(val string) (EntityShapeFromSqlEntityTypeEnum, bool) {
	enum, ok := mappingEntityShapeFromSqlEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
