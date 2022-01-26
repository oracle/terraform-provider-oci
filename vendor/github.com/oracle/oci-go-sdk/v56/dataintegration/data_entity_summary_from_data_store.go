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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DataEntitySummaryFromDataStore The view entity data entity details.
type DataEntitySummaryFromDataStore struct {
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

	Types *TypeLibrary `mandatory:"false" json:"types"`

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

	// Query filter for the extract. It can be Null.
	Filters *string `mandatory:"false" json:"filters"`

	// It shows whether or not effective date is disabled
	IsEffectiveDateDisabled *bool `mandatory:"false" json:"isEffectiveDateDisabled"`

	// Is Flex data store. Metadata csv will be generated for flex data store
	IsFlexDataStore *bool `mandatory:"false" json:"isFlexDataStore"`

	// Should the VO failure fail the whole batch?
	IsSilentError *bool `mandatory:"false" json:"isSilentError"`

	// It shows whether the datastore supports Incremental Extract or not.
	SupportsIncremental *bool `mandatory:"false" json:"supportsIncremental"`

	// The entity type.
	EntityType DataEntitySummaryFromDataStoreEntityTypeEnum `mandatory:"false" json:"entityType,omitempty"`
}

//GetMetadata returns Metadata
func (m DataEntitySummaryFromDataStore) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m DataEntitySummaryFromDataStore) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DataEntitySummaryFromDataStore) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataEntitySummaryFromDataStore DataEntitySummaryFromDataStore
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataEntitySummaryFromDataStore
	}{
		"DATA_STORE_ENTITY",
		(MarshalTypeDataEntitySummaryFromDataStore)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DataEntitySummaryFromDataStore) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Metadata                *ObjectMetadata                              `json:"metadata"`
		Key                     *string                                      `json:"key"`
		ModelVersion            *string                                      `json:"modelVersion"`
		ParentRef               *ParentReference                             `json:"parentRef"`
		Name                    *string                                      `json:"name"`
		Description             *string                                      `json:"description"`
		ObjectVersion           *int                                         `json:"objectVersion"`
		ExternalKey             *string                                      `json:"externalKey"`
		Shape                   *Shape                                       `json:"shape"`
		ShapeId                 *string                                      `json:"shapeId"`
		Types                   *TypeLibrary                                 `json:"types"`
		EntityType              DataEntitySummaryFromDataStoreEntityTypeEnum `json:"entityType"`
		OtherTypeLabel          *string                                      `json:"otherTypeLabel"`
		UniqueKeys              []uniquekey                                  `json:"uniqueKeys"`
		ForeignKeys             []ForeignKey                                 `json:"foreignKeys"`
		ResourceName            *string                                      `json:"resourceName"`
		ObjectStatus            *int                                         `json:"objectStatus"`
		Identifier              *string                                      `json:"identifier"`
		Filters                 *string                                      `json:"filters"`
		IsEffectiveDateDisabled *bool                                        `json:"isEffectiveDateDisabled"`
		IsFlexDataStore         *bool                                        `json:"isFlexDataStore"`
		IsSilentError           *bool                                        `json:"isSilentError"`
		SupportsIncremental     *bool                                        `json:"supportsIncremental"`
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
	for i, n := range model.ForeignKeys {
		m.ForeignKeys[i] = n
	}

	m.ResourceName = model.ResourceName

	m.ObjectStatus = model.ObjectStatus

	m.Identifier = model.Identifier

	m.Filters = model.Filters

	m.IsEffectiveDateDisabled = model.IsEffectiveDateDisabled

	m.IsFlexDataStore = model.IsFlexDataStore

	m.IsSilentError = model.IsSilentError

	m.SupportsIncremental = model.SupportsIncremental

	return
}

// DataEntitySummaryFromDataStoreEntityTypeEnum Enum with underlying type: string
type DataEntitySummaryFromDataStoreEntityTypeEnum string

// Set of constants representing the allowable values for DataEntitySummaryFromDataStoreEntityTypeEnum
const (
	DataEntitySummaryFromDataStoreEntityTypeTable     DataEntitySummaryFromDataStoreEntityTypeEnum = "TABLE"
	DataEntitySummaryFromDataStoreEntityTypeView      DataEntitySummaryFromDataStoreEntityTypeEnum = "VIEW"
	DataEntitySummaryFromDataStoreEntityTypeFile      DataEntitySummaryFromDataStoreEntityTypeEnum = "FILE"
	DataEntitySummaryFromDataStoreEntityTypeQueue     DataEntitySummaryFromDataStoreEntityTypeEnum = "QUEUE"
	DataEntitySummaryFromDataStoreEntityTypeStream    DataEntitySummaryFromDataStoreEntityTypeEnum = "STREAM"
	DataEntitySummaryFromDataStoreEntityTypeOther     DataEntitySummaryFromDataStoreEntityTypeEnum = "OTHER"
	DataEntitySummaryFromDataStoreEntityTypeDataStore DataEntitySummaryFromDataStoreEntityTypeEnum = "DATA_STORE"
)

var mappingDataEntitySummaryFromDataStoreEntityType = map[string]DataEntitySummaryFromDataStoreEntityTypeEnum{
	"TABLE":      DataEntitySummaryFromDataStoreEntityTypeTable,
	"VIEW":       DataEntitySummaryFromDataStoreEntityTypeView,
	"FILE":       DataEntitySummaryFromDataStoreEntityTypeFile,
	"QUEUE":      DataEntitySummaryFromDataStoreEntityTypeQueue,
	"STREAM":     DataEntitySummaryFromDataStoreEntityTypeStream,
	"OTHER":      DataEntitySummaryFromDataStoreEntityTypeOther,
	"DATA_STORE": DataEntitySummaryFromDataStoreEntityTypeDataStore,
}

// GetDataEntitySummaryFromDataStoreEntityTypeEnumValues Enumerates the set of values for DataEntitySummaryFromDataStoreEntityTypeEnum
func GetDataEntitySummaryFromDataStoreEntityTypeEnumValues() []DataEntitySummaryFromDataStoreEntityTypeEnum {
	values := make([]DataEntitySummaryFromDataStoreEntityTypeEnum, 0)
	for _, v := range mappingDataEntitySummaryFromDataStoreEntityType {
		values = append(values, v)
	}
	return values
}
