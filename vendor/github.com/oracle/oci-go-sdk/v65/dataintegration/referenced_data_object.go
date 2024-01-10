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

// ReferencedDataObject The input Operation for which derived entity is to be formed.
type ReferencedDataObject interface {

	// The object's model version.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// The resource name.
	GetResourceName() *string

	// The status of an object that can be set to value 1 for shallow reference across objects, other values reserved.
	GetObjectStatus() *int

	// The external key for the object.
	GetExternalKey() *string
}

type referenceddataobject struct {
	JsonData      []byte
	ModelVersion  *string          `mandatory:"false" json:"modelVersion"`
	ParentRef     *ParentReference `mandatory:"false" json:"parentRef"`
	Name          *string          `mandatory:"false" json:"name"`
	ObjectVersion *int             `mandatory:"false" json:"objectVersion"`
	ResourceName  *string          `mandatory:"false" json:"resourceName"`
	ObjectStatus  *int             `mandatory:"false" json:"objectStatus"`
	ExternalKey   *string          `mandatory:"false" json:"externalKey"`
	ModelType     string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *referenceddataobject) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerreferenceddataobject referenceddataobject
	s := struct {
		Model Unmarshalerreferenceddataobject
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Name = s.Model.Name
	m.ObjectVersion = s.Model.ObjectVersion
	m.ResourceName = s.Model.ResourceName
	m.ObjectStatus = s.Model.ObjectStatus
	m.ExternalKey = s.Model.ExternalKey
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *referenceddataobject) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "API":
		mm := ReferencedDataObjectFromApi{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PROCEDURE":
		mm := ReferencedDataObjectFromProcedure{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ReferencedDataObject: %s.", m.ModelType)
		return *m, nil
	}
}

// GetModelVersion returns ModelVersion
func (m referenceddataobject) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m referenceddataobject) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m referenceddataobject) GetName() *string {
	return m.Name
}

// GetObjectVersion returns ObjectVersion
func (m referenceddataobject) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetResourceName returns ResourceName
func (m referenceddataobject) GetResourceName() *string {
	return m.ResourceName
}

// GetObjectStatus returns ObjectStatus
func (m referenceddataobject) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetExternalKey returns ExternalKey
func (m referenceddataobject) GetExternalKey() *string {
	return m.ExternalKey
}

func (m referenceddataobject) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m referenceddataobject) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReferencedDataObjectModelTypeEnum Enum with underlying type: string
type ReferencedDataObjectModelTypeEnum string

// Set of constants representing the allowable values for ReferencedDataObjectModelTypeEnum
const (
	ReferencedDataObjectModelTypeProcedure ReferencedDataObjectModelTypeEnum = "PROCEDURE"
	ReferencedDataObjectModelTypeApi       ReferencedDataObjectModelTypeEnum = "API"
)

var mappingReferencedDataObjectModelTypeEnum = map[string]ReferencedDataObjectModelTypeEnum{
	"PROCEDURE": ReferencedDataObjectModelTypeProcedure,
	"API":       ReferencedDataObjectModelTypeApi,
}

var mappingReferencedDataObjectModelTypeEnumLowerCase = map[string]ReferencedDataObjectModelTypeEnum{
	"procedure": ReferencedDataObjectModelTypeProcedure,
	"api":       ReferencedDataObjectModelTypeApi,
}

// GetReferencedDataObjectModelTypeEnumValues Enumerates the set of values for ReferencedDataObjectModelTypeEnum
func GetReferencedDataObjectModelTypeEnumValues() []ReferencedDataObjectModelTypeEnum {
	values := make([]ReferencedDataObjectModelTypeEnum, 0)
	for _, v := range mappingReferencedDataObjectModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReferencedDataObjectModelTypeEnumStringValues Enumerates the set of values in String for ReferencedDataObjectModelTypeEnum
func GetReferencedDataObjectModelTypeEnumStringValues() []string {
	return []string{
		"PROCEDURE",
		"API",
	}
}

// GetMappingReferencedDataObjectModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReferencedDataObjectModelTypeEnum(val string) (ReferencedDataObjectModelTypeEnum, bool) {
	enum, ok := mappingReferencedDataObjectModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
