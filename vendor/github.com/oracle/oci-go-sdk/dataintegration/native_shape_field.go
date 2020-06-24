// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// NativeShapeField The native shape field object.
type NativeShapeField struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// The model type reference.
	ModelType *string `mandatory:"false" json:"modelType"`

	// The type reference.
	Type *string `mandatory:"false" json:"type"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The position of the attribute.
	Position *int `mandatory:"false" json:"position"`

	// The default value.
	DefaultValueString *string `mandatory:"false" json:"defaultValueString"`

	// Specifies whether the field is mandatory.
	IsMandatory *bool `mandatory:"false" json:"isMandatory"`
}

func (m NativeShapeField) String() string {
	return common.PointerString(m)
}
