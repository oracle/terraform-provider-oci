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

// TypedExpression The expression that can be created, using the execute stage output in REST Task.
type TypedExpression struct {

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

	// The expression string for the object.
	Expression *string `mandatory:"false" json:"expression"`

	// The object type.
	Type *string `mandatory:"false" json:"type"`
}

// GetKey returns Key
func (m TypedExpression) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m TypedExpression) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m TypedExpression) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetConfigValues returns ConfigValues
func (m TypedExpression) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

// GetObjectStatus returns ObjectStatus
func (m TypedExpression) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetName returns Name
func (m TypedExpression) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m TypedExpression) GetDescription() *string {
	return m.Description
}

func (m TypedExpression) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TypedExpression) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TypedExpression) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTypedExpression TypedExpression
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTypedExpression
	}{
		"TYPED_EXPRESSION",
		(MarshalTypeTypedExpression)(m),
	}

	return json.Marshal(&s)
}
