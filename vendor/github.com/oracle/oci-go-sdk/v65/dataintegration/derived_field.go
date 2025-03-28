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

// DerivedField The type representing the derived field concept. Derived fields have an expression to define how to derive the field.
type DerivedField struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	Expr *Expression `mandatory:"false" json:"expr"`

	// The type of the field.
	Type *string `mandatory:"false" json:"type"`

	// Specifies whether to use inferred expression output type as output type of the derived field. Default value of this flag is false.
	IsUseInferredType *bool `mandatory:"false" json:"isUseInferredType"`

	// Labels are keywords or labels that you can add to data assets, dataflows and so on. You can define your own labels and use them to categorize content.
	Labels []string `mandatory:"false" json:"labels"`
}

// GetKey returns Key
func (m DerivedField) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m DerivedField) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m DerivedField) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetConfigValues returns ConfigValues
func (m DerivedField) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

// GetObjectStatus returns ObjectStatus
func (m DerivedField) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetName returns Name
func (m DerivedField) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m DerivedField) GetDescription() *string {
	return m.Description
}

func (m DerivedField) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DerivedField) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DerivedField) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDerivedField DerivedField
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDerivedField
	}{
		"DERIVED_FIELD",
		(MarshalTypeDerivedField)(m),
	}

	return json.Marshal(&s)
}
