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

// MaterializedCompositeType A `MaterializedCompositeType` represents a type that is composed of a list of sub-types, for example an `Address` type.   The sub-types can be simple `DataType` or other `CompositeType` objects. Typically, a `CompositeType` may represent an arbitrarily deep hierarchy of types.
type MaterializedCompositeType struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// A user defined description for the object.
	Description *string `mandatory:"false" json:"description"`

	// An array of elements.
	Elements []string `mandatory:"false" json:"elements"`

	// An array of path names corresponding to the elements.  The path names are used when referring to the field in an expression.
	PathNames []string `mandatory:"false" json:"pathNames"`

	ConfigDefinition *ConfigDefinition `mandatory:"false" json:"configDefinition"`
}

// GetKey returns Key
func (m MaterializedCompositeType) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m MaterializedCompositeType) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m MaterializedCompositeType) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m MaterializedCompositeType) GetName() *string {
	return m.Name
}

// GetObjectStatus returns ObjectStatus
func (m MaterializedCompositeType) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetDescription returns Description
func (m MaterializedCompositeType) GetDescription() *string {
	return m.Description
}

func (m MaterializedCompositeType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaterializedCompositeType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MaterializedCompositeType) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMaterializedCompositeType MaterializedCompositeType
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeMaterializedCompositeType
	}{
		"MATERIALIZED_COMPOSITE_TYPE",
		(MarshalTypeMaterializedCompositeType)(m),
	}

	return json.Marshal(&s)
}
