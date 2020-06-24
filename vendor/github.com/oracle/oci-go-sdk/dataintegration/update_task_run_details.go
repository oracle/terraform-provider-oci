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

// UpdateTaskRunDetails Properties used in task run update operations.
type UpdateTaskRunDetails struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// status
	Status UpdateTaskRunDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
}

func (m UpdateTaskRunDetails) String() string {
	return common.PointerString(m)
}

// UpdateTaskRunDetailsStatusEnum Enum with underlying type: string
type UpdateTaskRunDetailsStatusEnum string

// Set of constants representing the allowable values for UpdateTaskRunDetailsStatusEnum
const (
	UpdateTaskRunDetailsStatusTerminating UpdateTaskRunDetailsStatusEnum = "TERMINATING"
)

var mappingUpdateTaskRunDetailsStatus = map[string]UpdateTaskRunDetailsStatusEnum{
	"TERMINATING": UpdateTaskRunDetailsStatusTerminating,
}

// GetUpdateTaskRunDetailsStatusEnumValues Enumerates the set of values for UpdateTaskRunDetailsStatusEnum
func GetUpdateTaskRunDetailsStatusEnumValues() []UpdateTaskRunDetailsStatusEnum {
	values := make([]UpdateTaskRunDetailsStatusEnum, 0)
	for _, v := range mappingUpdateTaskRunDetailsStatus {
		values = append(values, v)
	}
	return values
}
