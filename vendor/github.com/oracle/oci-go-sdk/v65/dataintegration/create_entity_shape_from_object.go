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

// CreateEntityShapeFromObject The application object entity details.
type CreateEntityShapeFromObject struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

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

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The entity type.
	EntityType CreateEntityShapeFromObjectEntityTypeEnum `mandatory:"false" json:"entityType,omitempty"`
}

func (m CreateEntityShapeFromObject) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateEntityShapeFromObject) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateEntityShapeFromObjectEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetCreateEntityShapeFromObjectEntityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateEntityShapeFromObject) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateEntityShapeFromObject CreateEntityShapeFromObject
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCreateEntityShapeFromObject
	}{
		"OBJECT_ENTITY",
		(MarshalTypeCreateEntityShapeFromObject)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateEntityShapeFromObject) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key            *string                                   `json:"key"`
		ParentRef      *ParentReference                          `json:"parentRef"`
		Description    *string                                   `json:"description"`
		ExternalKey    *string                                   `json:"externalKey"`
		Shape          *Shape                                    `json:"shape"`
		ShapeId        *string                                   `json:"shapeId"`
		EntityType     CreateEntityShapeFromObjectEntityTypeEnum `json:"entityType"`
		OtherTypeLabel *string                                   `json:"otherTypeLabel"`
		UniqueKeys     []uniquekey                               `json:"uniqueKeys"`
		ForeignKeys    []ForeignKey                              `json:"foreignKeys"`
		ResourceName   *string                                   `json:"resourceName"`
		DataFormat     *DataFormat                               `json:"dataFormat"`
		Identifier     *string                                   `json:"identifier"`
		Name           *string                                   `json:"name"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ParentRef = model.ParentRef

	m.Description = model.Description

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

	m.Identifier = model.Identifier

	m.Name = model.Name

	return
}

// CreateEntityShapeFromObjectEntityTypeEnum Enum with underlying type: string
type CreateEntityShapeFromObjectEntityTypeEnum string

// Set of constants representing the allowable values for CreateEntityShapeFromObjectEntityTypeEnum
const (
	CreateEntityShapeFromObjectEntityTypeTable  CreateEntityShapeFromObjectEntityTypeEnum = "TABLE"
	CreateEntityShapeFromObjectEntityTypeView   CreateEntityShapeFromObjectEntityTypeEnum = "VIEW"
	CreateEntityShapeFromObjectEntityTypeFile   CreateEntityShapeFromObjectEntityTypeEnum = "FILE"
	CreateEntityShapeFromObjectEntityTypeSql    CreateEntityShapeFromObjectEntityTypeEnum = "SQL"
	CreateEntityShapeFromObjectEntityTypeObject CreateEntityShapeFromObjectEntityTypeEnum = "OBJECT"
)

var mappingCreateEntityShapeFromObjectEntityTypeEnum = map[string]CreateEntityShapeFromObjectEntityTypeEnum{
	"TABLE":  CreateEntityShapeFromObjectEntityTypeTable,
	"VIEW":   CreateEntityShapeFromObjectEntityTypeView,
	"FILE":   CreateEntityShapeFromObjectEntityTypeFile,
	"SQL":    CreateEntityShapeFromObjectEntityTypeSql,
	"OBJECT": CreateEntityShapeFromObjectEntityTypeObject,
}

var mappingCreateEntityShapeFromObjectEntityTypeEnumLowerCase = map[string]CreateEntityShapeFromObjectEntityTypeEnum{
	"table":  CreateEntityShapeFromObjectEntityTypeTable,
	"view":   CreateEntityShapeFromObjectEntityTypeView,
	"file":   CreateEntityShapeFromObjectEntityTypeFile,
	"sql":    CreateEntityShapeFromObjectEntityTypeSql,
	"object": CreateEntityShapeFromObjectEntityTypeObject,
}

// GetCreateEntityShapeFromObjectEntityTypeEnumValues Enumerates the set of values for CreateEntityShapeFromObjectEntityTypeEnum
func GetCreateEntityShapeFromObjectEntityTypeEnumValues() []CreateEntityShapeFromObjectEntityTypeEnum {
	values := make([]CreateEntityShapeFromObjectEntityTypeEnum, 0)
	for _, v := range mappingCreateEntityShapeFromObjectEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateEntityShapeFromObjectEntityTypeEnumStringValues Enumerates the set of values in String for CreateEntityShapeFromObjectEntityTypeEnum
func GetCreateEntityShapeFromObjectEntityTypeEnumStringValues() []string {
	return []string{
		"TABLE",
		"VIEW",
		"FILE",
		"SQL",
		"OBJECT",
	}
}

// GetMappingCreateEntityShapeFromObjectEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateEntityShapeFromObjectEntityTypeEnum(val string) (CreateEntityShapeFromObjectEntityTypeEnum, bool) {
	enum, ok := mappingCreateEntityShapeFromObjectEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
