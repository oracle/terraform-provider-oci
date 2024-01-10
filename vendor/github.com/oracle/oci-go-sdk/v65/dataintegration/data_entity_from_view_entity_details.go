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

// DataEntityFromViewEntityDetails The view entity data entity details.
type DataEntityFromViewEntityDetails struct {

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

	Types *TypeLibrary `mandatory:"false" json:"types"`

	// Specifies other type labels.
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

	// The entity type.
	EntityType DataEntityFromViewEntityDetailsEntityTypeEnum `mandatory:"false" json:"entityType,omitempty"`
}

func (m DataEntityFromViewEntityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataEntityFromViewEntityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataEntityFromViewEntityDetailsEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetDataEntityFromViewEntityDetailsEntityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataEntityFromViewEntityDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataEntityFromViewEntityDetails DataEntityFromViewEntityDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataEntityFromViewEntityDetails
	}{
		"VIEW_ENTITY",
		(MarshalTypeDataEntityFromViewEntityDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DataEntityFromViewEntityDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key            *string                                       `json:"key"`
		ModelVersion   *string                                       `json:"modelVersion"`
		ParentRef      *ParentReference                              `json:"parentRef"`
		Name           *string                                       `json:"name"`
		Description    *string                                       `json:"description"`
		ObjectVersion  *int                                          `json:"objectVersion"`
		ExternalKey    *string                                       `json:"externalKey"`
		Shape          *Shape                                        `json:"shape"`
		ShapeId        *string                                       `json:"shapeId"`
		Types          *TypeLibrary                                  `json:"types"`
		EntityType     DataEntityFromViewEntityDetailsEntityTypeEnum `json:"entityType"`
		OtherTypeLabel *string                                       `json:"otherTypeLabel"`
		UniqueKeys     []uniquekey                                   `json:"uniqueKeys"`
		ForeignKeys    []ForeignKey                                  `json:"foreignKeys"`
		ResourceName   *string                                       `json:"resourceName"`
		ObjectStatus   *int                                          `json:"objectStatus"`
		Identifier     *string                                       `json:"identifier"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.Description = model.Description

	m.ObjectVersion = model.ObjectVersion

	m.ExternalKey = model.ExternalKey

	m.Shape = model.Shape

	m.ShapeId = model.ShapeId

	m.Types = model.Types

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

	return
}

// DataEntityFromViewEntityDetailsEntityTypeEnum Enum with underlying type: string
type DataEntityFromViewEntityDetailsEntityTypeEnum string

// Set of constants representing the allowable values for DataEntityFromViewEntityDetailsEntityTypeEnum
const (
	DataEntityFromViewEntityDetailsEntityTypeTable  DataEntityFromViewEntityDetailsEntityTypeEnum = "TABLE"
	DataEntityFromViewEntityDetailsEntityTypeView   DataEntityFromViewEntityDetailsEntityTypeEnum = "VIEW"
	DataEntityFromViewEntityDetailsEntityTypeFile   DataEntityFromViewEntityDetailsEntityTypeEnum = "FILE"
	DataEntityFromViewEntityDetailsEntityTypeQueue  DataEntityFromViewEntityDetailsEntityTypeEnum = "QUEUE"
	DataEntityFromViewEntityDetailsEntityTypeStream DataEntityFromViewEntityDetailsEntityTypeEnum = "STREAM"
	DataEntityFromViewEntityDetailsEntityTypeOther  DataEntityFromViewEntityDetailsEntityTypeEnum = "OTHER"
)

var mappingDataEntityFromViewEntityDetailsEntityTypeEnum = map[string]DataEntityFromViewEntityDetailsEntityTypeEnum{
	"TABLE":  DataEntityFromViewEntityDetailsEntityTypeTable,
	"VIEW":   DataEntityFromViewEntityDetailsEntityTypeView,
	"FILE":   DataEntityFromViewEntityDetailsEntityTypeFile,
	"QUEUE":  DataEntityFromViewEntityDetailsEntityTypeQueue,
	"STREAM": DataEntityFromViewEntityDetailsEntityTypeStream,
	"OTHER":  DataEntityFromViewEntityDetailsEntityTypeOther,
}

var mappingDataEntityFromViewEntityDetailsEntityTypeEnumLowerCase = map[string]DataEntityFromViewEntityDetailsEntityTypeEnum{
	"table":  DataEntityFromViewEntityDetailsEntityTypeTable,
	"view":   DataEntityFromViewEntityDetailsEntityTypeView,
	"file":   DataEntityFromViewEntityDetailsEntityTypeFile,
	"queue":  DataEntityFromViewEntityDetailsEntityTypeQueue,
	"stream": DataEntityFromViewEntityDetailsEntityTypeStream,
	"other":  DataEntityFromViewEntityDetailsEntityTypeOther,
}

// GetDataEntityFromViewEntityDetailsEntityTypeEnumValues Enumerates the set of values for DataEntityFromViewEntityDetailsEntityTypeEnum
func GetDataEntityFromViewEntityDetailsEntityTypeEnumValues() []DataEntityFromViewEntityDetailsEntityTypeEnum {
	values := make([]DataEntityFromViewEntityDetailsEntityTypeEnum, 0)
	for _, v := range mappingDataEntityFromViewEntityDetailsEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataEntityFromViewEntityDetailsEntityTypeEnumStringValues Enumerates the set of values in String for DataEntityFromViewEntityDetailsEntityTypeEnum
func GetDataEntityFromViewEntityDetailsEntityTypeEnumStringValues() []string {
	return []string{
		"TABLE",
		"VIEW",
		"FILE",
		"QUEUE",
		"STREAM",
		"OTHER",
	}
}

// GetMappingDataEntityFromViewEntityDetailsEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataEntityFromViewEntityDetailsEntityTypeEnum(val string) (DataEntityFromViewEntityDetailsEntityTypeEnum, bool) {
	enum, ok := mappingDataEntityFromViewEntityDetailsEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
