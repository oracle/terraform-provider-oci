// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateOdaInstanceDetails Properties that are required to create a Digital Assistant instance.
type CreateOdaInstanceDetails struct {

	// Identifier of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Shape or size of the instance.
	ShapeName CreateOdaInstanceDetailsShapeNameEnum `mandatory:"true" json:"shapeName"`

	// User-friendly name for the instance. Avoid entering confidential information. You can change this value anytime.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the Digital Assistant instance.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for
	// cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateOdaInstanceDetails) String() string {
	return common.PointerString(m)
}

// CreateOdaInstanceDetailsShapeNameEnum Enum with underlying type: string
type CreateOdaInstanceDetailsShapeNameEnum string

// Set of constants representing the allowable values for CreateOdaInstanceDetailsShapeNameEnum
const (
	CreateOdaInstanceDetailsShapeNameDevelopment CreateOdaInstanceDetailsShapeNameEnum = "DEVELOPMENT"
	CreateOdaInstanceDetailsShapeNameProduction  CreateOdaInstanceDetailsShapeNameEnum = "PRODUCTION"
)

var mappingCreateOdaInstanceDetailsShapeName = map[string]CreateOdaInstanceDetailsShapeNameEnum{
	"DEVELOPMENT": CreateOdaInstanceDetailsShapeNameDevelopment,
	"PRODUCTION":  CreateOdaInstanceDetailsShapeNameProduction,
}

// GetCreateOdaInstanceDetailsShapeNameEnumValues Enumerates the set of values for CreateOdaInstanceDetailsShapeNameEnum
func GetCreateOdaInstanceDetailsShapeNameEnumValues() []CreateOdaInstanceDetailsShapeNameEnum {
	values := make([]CreateOdaInstanceDetailsShapeNameEnum, 0)
	for _, v := range mappingCreateOdaInstanceDetailsShapeName {
		values = append(values, v)
	}
	return values
}
