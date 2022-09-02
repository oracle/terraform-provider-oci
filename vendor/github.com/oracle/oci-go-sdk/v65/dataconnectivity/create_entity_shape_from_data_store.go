// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateEntityShapeFromDataStore The data store entity details.
type CreateEntityShapeFromDataStore struct {

	// Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The model version of the object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The external key of the object.
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

	// Value can only contain upper case letters, underscore, and numbers. It should begin with an upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	Types *TypeLibrary `mandatory:"false" json:"types"`

	// Map<String, String> for entity properties
	EntityProperties map[string]string `mandatory:"false" json:"entityProperties"`

	// The entity type.
	EntityType CreateEntityShapeDetailsEntityTypeEnum `mandatory:"false" json:"entityType,omitempty"`
}

//GetKey returns Key
func (m CreateEntityShapeFromDataStore) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m CreateEntityShapeFromDataStore) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m CreateEntityShapeFromDataStore) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m CreateEntityShapeFromDataStore) GetName() *string {
	return m.Name
}

//GetObjectVersion returns ObjectVersion
func (m CreateEntityShapeFromDataStore) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetExternalKey returns ExternalKey
func (m CreateEntityShapeFromDataStore) GetExternalKey() *string {
	return m.ExternalKey
}

//GetShape returns Shape
func (m CreateEntityShapeFromDataStore) GetShape() *Shape {
	return m.Shape
}

//GetShapeId returns ShapeId
func (m CreateEntityShapeFromDataStore) GetShapeId() *string {
	return m.ShapeId
}

//GetEntityType returns EntityType
func (m CreateEntityShapeFromDataStore) GetEntityType() CreateEntityShapeDetailsEntityTypeEnum {
	return m.EntityType
}

//GetOtherTypeLabel returns OtherTypeLabel
func (m CreateEntityShapeFromDataStore) GetOtherTypeLabel() *string {
	return m.OtherTypeLabel
}

//GetUniqueKeys returns UniqueKeys
func (m CreateEntityShapeFromDataStore) GetUniqueKeys() []UniqueKey {
	return m.UniqueKeys
}

//GetForeignKeys returns ForeignKeys
func (m CreateEntityShapeFromDataStore) GetForeignKeys() []ForeignKey {
	return m.ForeignKeys
}

//GetResourceName returns ResourceName
func (m CreateEntityShapeFromDataStore) GetResourceName() *string {
	return m.ResourceName
}

//GetObjectStatus returns ObjectStatus
func (m CreateEntityShapeFromDataStore) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m CreateEntityShapeFromDataStore) GetIdentifier() *string {
	return m.Identifier
}

//GetTypes returns Types
func (m CreateEntityShapeFromDataStore) GetTypes() *TypeLibrary {
	return m.Types
}

//GetEntityProperties returns EntityProperties
func (m CreateEntityShapeFromDataStore) GetEntityProperties() map[string]string {
	return m.EntityProperties
}

func (m CreateEntityShapeFromDataStore) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateEntityShapeFromDataStore) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateEntityShapeDetailsEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetCreateEntityShapeDetailsEntityTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateEntityShapeFromDataStore) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateEntityShapeFromDataStore CreateEntityShapeFromDataStore
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCreateEntityShapeFromDataStore
	}{
		"DATA_STORE_ENTITY",
		(MarshalTypeCreateEntityShapeFromDataStore)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateEntityShapeFromDataStore) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key              *string                                `json:"key"`
		ModelVersion     *string                                `json:"modelVersion"`
		ParentRef        *ParentReference                       `json:"parentRef"`
		ObjectVersion    *int                                   `json:"objectVersion"`
		ExternalKey      *string                                `json:"externalKey"`
		Shape            *Shape                                 `json:"shape"`
		ShapeId          *string                                `json:"shapeId"`
		EntityType       CreateEntityShapeDetailsEntityTypeEnum `json:"entityType"`
		OtherTypeLabel   *string                                `json:"otherTypeLabel"`
		UniqueKeys       []uniquekey                            `json:"uniqueKeys"`
		ForeignKeys      []ForeignKey                           `json:"foreignKeys"`
		ResourceName     *string                                `json:"resourceName"`
		ObjectStatus     *int                                   `json:"objectStatus"`
		Identifier       *string                                `json:"identifier"`
		Types            *TypeLibrary                           `json:"types"`
		EntityProperties map[string]string                      `json:"entityProperties"`
		Name             *string                                `json:"name"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

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
	for i, n := range model.ForeignKeys {
		m.ForeignKeys[i] = n
	}

	m.ResourceName = model.ResourceName

	m.ObjectStatus = model.ObjectStatus

	m.Identifier = model.Identifier

	m.Types = model.Types

	m.EntityProperties = model.EntityProperties

	m.Name = model.Name

	return
}
