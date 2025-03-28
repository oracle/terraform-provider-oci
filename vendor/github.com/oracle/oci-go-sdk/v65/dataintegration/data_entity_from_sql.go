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

// DataEntityFromSql The sql entity data entity details.
type DataEntityFromSql struct {
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

	// The external key for the object
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

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// sqlQuery
	SqlQuery *string `mandatory:"false" json:"sqlQuery"`

	// The entity type.
	EntityType DataEntityFromSqlEntityTypeEnum `mandatory:"false" json:"entityType,omitempty"`
}

// GetMetadata returns Metadata
func (m DataEntityFromSql) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m DataEntityFromSql) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataEntityFromSql) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataEntityFromSqlEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetDataEntityFromSqlEntityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataEntityFromSql) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataEntityFromSql DataEntityFromSql
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataEntityFromSql
	}{
		"SQL_ENTITY",
		(MarshalTypeDataEntityFromSql)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DataEntityFromSql) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Metadata       *ObjectMetadata                 `json:"metadata"`
		Key            *string                         `json:"key"`
		ModelVersion   *string                         `json:"modelVersion"`
		ParentRef      *ParentReference                `json:"parentRef"`
		Name           *string                         `json:"name"`
		Description    *string                         `json:"description"`
		ObjectVersion  *int                            `json:"objectVersion"`
		ExternalKey    *string                         `json:"externalKey"`
		Shape          *Shape                          `json:"shape"`
		ShapeId        *string                         `json:"shapeId"`
		EntityType     DataEntityFromSqlEntityTypeEnum `json:"entityType"`
		OtherTypeLabel *string                         `json:"otherTypeLabel"`
		UniqueKeys     []uniquekey                     `json:"uniqueKeys"`
		ForeignKeys    []ForeignKey                    `json:"foreignKeys"`
		ResourceName   *string                         `json:"resourceName"`
		ObjectStatus   *int                            `json:"objectStatus"`
		Identifier     *string                         `json:"identifier"`
		SqlQuery       *string                         `json:"sqlQuery"`
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

	m.ObjectStatus = model.ObjectStatus

	m.Identifier = model.Identifier

	m.SqlQuery = model.SqlQuery

	return
}

// DataEntityFromSqlEntityTypeEnum Enum with underlying type: string
type DataEntityFromSqlEntityTypeEnum string

// Set of constants representing the allowable values for DataEntityFromSqlEntityTypeEnum
const (
	DataEntityFromSqlEntityTypeTable DataEntityFromSqlEntityTypeEnum = "TABLE"
	DataEntityFromSqlEntityTypeView  DataEntityFromSqlEntityTypeEnum = "VIEW"
	DataEntityFromSqlEntityTypeFile  DataEntityFromSqlEntityTypeEnum = "FILE"
	DataEntityFromSqlEntityTypeSql   DataEntityFromSqlEntityTypeEnum = "SQL"
)

var mappingDataEntityFromSqlEntityTypeEnum = map[string]DataEntityFromSqlEntityTypeEnum{
	"TABLE": DataEntityFromSqlEntityTypeTable,
	"VIEW":  DataEntityFromSqlEntityTypeView,
	"FILE":  DataEntityFromSqlEntityTypeFile,
	"SQL":   DataEntityFromSqlEntityTypeSql,
}

var mappingDataEntityFromSqlEntityTypeEnumLowerCase = map[string]DataEntityFromSqlEntityTypeEnum{
	"table": DataEntityFromSqlEntityTypeTable,
	"view":  DataEntityFromSqlEntityTypeView,
	"file":  DataEntityFromSqlEntityTypeFile,
	"sql":   DataEntityFromSqlEntityTypeSql,
}

// GetDataEntityFromSqlEntityTypeEnumValues Enumerates the set of values for DataEntityFromSqlEntityTypeEnum
func GetDataEntityFromSqlEntityTypeEnumValues() []DataEntityFromSqlEntityTypeEnum {
	values := make([]DataEntityFromSqlEntityTypeEnum, 0)
	for _, v := range mappingDataEntityFromSqlEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataEntityFromSqlEntityTypeEnumStringValues Enumerates the set of values in String for DataEntityFromSqlEntityTypeEnum
func GetDataEntityFromSqlEntityTypeEnumStringValues() []string {
	return []string{
		"TABLE",
		"VIEW",
		"FILE",
		"SQL",
	}
}

// GetMappingDataEntityFromSqlEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataEntityFromSqlEntityTypeEnum(val string) (DataEntityFromSqlEntityTypeEnum, bool) {
	enum, ok := mappingDataEntityFromSqlEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
