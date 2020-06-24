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

// PublishedObjectSummary The published obect summary.
type PublishedObjectSummary interface {

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
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

	GetMetadata() *ObjectMetadata
}

type publishedobjectsummary struct {
	JsonData      []byte
	Key           *string          `mandatory:"false" json:"key"`
	ModelVersion  *string          `mandatory:"false" json:"modelVersion"`
	ParentRef     *ParentReference `mandatory:"false" json:"parentRef"`
	Name          *string          `mandatory:"false" json:"name"`
	Description   *string          `mandatory:"false" json:"description"`
	ObjectVersion *int             `mandatory:"false" json:"objectVersion"`
	ObjectStatus  *int             `mandatory:"false" json:"objectStatus"`
	Identifier    *string          `mandatory:"false" json:"identifier"`
	Metadata      *ObjectMetadata  `mandatory:"false" json:"metadata"`
	ModelType     string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *publishedobjectsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpublishedobjectsummary publishedobjectsummary
	s := struct {
		Model Unmarshalerpublishedobjectsummary
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
	m.Metadata = s.Model.Metadata
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *publishedobjectsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "INTEGRATION_TASK":
		mm := PublishedObjectSummaryFromIntegrationTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATA_LOADER_TASK":
		mm := PublishedObjectSummaryFromDataLoaderTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m publishedobjectsummary) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m publishedobjectsummary) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m publishedobjectsummary) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m publishedobjectsummary) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m publishedobjectsummary) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m publishedobjectsummary) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m publishedobjectsummary) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m publishedobjectsummary) GetIdentifier() *string {
	return m.Identifier
}

//GetMetadata returns Metadata
func (m publishedobjectsummary) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m publishedobjectsummary) String() string {
	return common.PointerString(m)
}

// PublishedObjectSummaryModelTypeEnum Enum with underlying type: string
type PublishedObjectSummaryModelTypeEnum string

// Set of constants representing the allowable values for PublishedObjectSummaryModelTypeEnum
const (
	PublishedObjectSummaryModelTypeIntegrationTask PublishedObjectSummaryModelTypeEnum = "INTEGRATION_TASK"
	PublishedObjectSummaryModelTypeDataLoaderTask  PublishedObjectSummaryModelTypeEnum = "DATA_LOADER_TASK"
)

var mappingPublishedObjectSummaryModelType = map[string]PublishedObjectSummaryModelTypeEnum{
	"INTEGRATION_TASK": PublishedObjectSummaryModelTypeIntegrationTask,
	"DATA_LOADER_TASK": PublishedObjectSummaryModelTypeDataLoaderTask,
}

// GetPublishedObjectSummaryModelTypeEnumValues Enumerates the set of values for PublishedObjectSummaryModelTypeEnum
func GetPublishedObjectSummaryModelTypeEnumValues() []PublishedObjectSummaryModelTypeEnum {
	values := make([]PublishedObjectSummaryModelTypeEnum, 0)
	for _, v := range mappingPublishedObjectSummaryModelType {
		values = append(values, v)
	}
	return values
}
