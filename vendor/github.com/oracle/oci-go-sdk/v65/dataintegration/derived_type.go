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

// DerivedType A `DerivedType` object represents a more complex type that is derived from a set of simple types, for example an `Address` or `SSN` data type.
type DerivedType struct {

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
}

// GetKey returns Key
func (m DerivedType) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m DerivedType) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m DerivedType) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m DerivedType) GetName() *string {
	return m.Name
}

// GetObjectStatus returns ObjectStatus
func (m DerivedType) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetDescription returns Description
func (m DerivedType) GetDescription() *string {
	return m.Description
}

func (m DerivedType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DerivedType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DerivedType) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDerivedType DerivedType
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDerivedType
	}{
		"DERIVED_TYPE",
		(MarshalTypeDerivedType)(m),
	}

	return json.Marshal(&s)
}
