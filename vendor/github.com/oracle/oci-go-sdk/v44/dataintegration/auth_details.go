// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v44/common"
)

// AuthDetails Authentication type to be used for Generic REST invocation.
type AuthDetails struct {

	// Generated key that can be used in API calls to identify data flow. On scenarios where reference to the data flow is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// The authentication mode to be used for Generic REST invocation.
	ModelType AuthDetailsModelTypeEnum `mandatory:"false" json:"modelType,omitempty"`
}

func (m AuthDetails) String() string {
	return common.PointerString(m)
}

// AuthDetailsModelTypeEnum Enum with underlying type: string
type AuthDetailsModelTypeEnum string

// Set of constants representing the allowable values for AuthDetailsModelTypeEnum
const (
	AuthDetailsModelTypeNoAuthDetails                AuthDetailsModelTypeEnum = "NO_AUTH_DETAILS"
	AuthDetailsModelTypeResourcePrincipalAuthDetails AuthDetailsModelTypeEnum = "RESOURCE_PRINCIPAL_AUTH_DETAILS"
)

var mappingAuthDetailsModelType = map[string]AuthDetailsModelTypeEnum{
	"NO_AUTH_DETAILS":                 AuthDetailsModelTypeNoAuthDetails,
	"RESOURCE_PRINCIPAL_AUTH_DETAILS": AuthDetailsModelTypeResourcePrincipalAuthDetails,
}

// GetAuthDetailsModelTypeEnumValues Enumerates the set of values for AuthDetailsModelTypeEnum
func GetAuthDetailsModelTypeEnumValues() []AuthDetailsModelTypeEnum {
	values := make([]AuthDetailsModelTypeEnum, 0)
	for _, v := range mappingAuthDetailsModelType {
		values = append(values, v)
	}
	return values
}
