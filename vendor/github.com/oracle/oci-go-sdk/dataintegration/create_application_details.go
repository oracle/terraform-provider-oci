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

// CreateApplicationDetails Properties used in application create operations.
type CreateApplicationDetails struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"true" json:"identifier"`

	// Currently not used on application creation. Reserved for future.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// The type of the application.
	ModelType CreateApplicationDetailsModelTypeEnum `mandatory:"false" json:"modelType,omitempty"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
}

func (m CreateApplicationDetails) String() string {
	return common.PointerString(m)
}

// CreateApplicationDetailsModelTypeEnum Enum with underlying type: string
type CreateApplicationDetailsModelTypeEnum string

// Set of constants representing the allowable values for CreateApplicationDetailsModelTypeEnum
const (
	CreateApplicationDetailsModelTypeIntegrationApplication CreateApplicationDetailsModelTypeEnum = "INTEGRATION_APPLICATION"
)

var mappingCreateApplicationDetailsModelType = map[string]CreateApplicationDetailsModelTypeEnum{
	"INTEGRATION_APPLICATION": CreateApplicationDetailsModelTypeIntegrationApplication,
}

// GetCreateApplicationDetailsModelTypeEnumValues Enumerates the set of values for CreateApplicationDetailsModelTypeEnum
func GetCreateApplicationDetailsModelTypeEnumValues() []CreateApplicationDetailsModelTypeEnum {
	values := make([]CreateApplicationDetailsModelTypeEnum, 0)
	for _, v := range mappingCreateApplicationDetailsModelType {
		values = append(values, v)
	}
	return values
}
