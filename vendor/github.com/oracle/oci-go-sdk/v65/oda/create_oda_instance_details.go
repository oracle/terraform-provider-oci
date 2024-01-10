// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Should this Digital Assistant instance use role-based authorization via an identity domain (true) or use the default policy-based authorization via IAM policies (false)
	IsRoleBasedAccess *bool `mandatory:"false" json:"isRoleBasedAccess"`

	// If isRoleBasedAccess is set to true, this property specifies the identity domain that is to be used to implement this type of authorzation. Digital Assistant will create an Identity Application instance and Application Roles within this identity domain. The caller may then perform and user roll mappings they like to grant access to users within the identity domain.
	IdentityDomain *string `mandatory:"false" json:"identityDomain"`
}

func (m CreateOdaInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOdaInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateOdaInstanceDetailsShapeNameEnum(string(m.ShapeName)); !ok && m.ShapeName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShapeName: %s. Supported values are: %s.", m.ShapeName, strings.Join(GetCreateOdaInstanceDetailsShapeNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateOdaInstanceDetailsShapeNameEnum Enum with underlying type: string
type CreateOdaInstanceDetailsShapeNameEnum string

// Set of constants representing the allowable values for CreateOdaInstanceDetailsShapeNameEnum
const (
	CreateOdaInstanceDetailsShapeNameDevelopment CreateOdaInstanceDetailsShapeNameEnum = "DEVELOPMENT"
	CreateOdaInstanceDetailsShapeNameProduction  CreateOdaInstanceDetailsShapeNameEnum = "PRODUCTION"
)

var mappingCreateOdaInstanceDetailsShapeNameEnum = map[string]CreateOdaInstanceDetailsShapeNameEnum{
	"DEVELOPMENT": CreateOdaInstanceDetailsShapeNameDevelopment,
	"PRODUCTION":  CreateOdaInstanceDetailsShapeNameProduction,
}

var mappingCreateOdaInstanceDetailsShapeNameEnumLowerCase = map[string]CreateOdaInstanceDetailsShapeNameEnum{
	"development": CreateOdaInstanceDetailsShapeNameDevelopment,
	"production":  CreateOdaInstanceDetailsShapeNameProduction,
}

// GetCreateOdaInstanceDetailsShapeNameEnumValues Enumerates the set of values for CreateOdaInstanceDetailsShapeNameEnum
func GetCreateOdaInstanceDetailsShapeNameEnumValues() []CreateOdaInstanceDetailsShapeNameEnum {
	values := make([]CreateOdaInstanceDetailsShapeNameEnum, 0)
	for _, v := range mappingCreateOdaInstanceDetailsShapeNameEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOdaInstanceDetailsShapeNameEnumStringValues Enumerates the set of values in String for CreateOdaInstanceDetailsShapeNameEnum
func GetCreateOdaInstanceDetailsShapeNameEnumStringValues() []string {
	return []string{
		"DEVELOPMENT",
		"PRODUCTION",
	}
}

// GetMappingCreateOdaInstanceDetailsShapeNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOdaInstanceDetailsShapeNameEnum(val string) (CreateOdaInstanceDetailsShapeNameEnum, bool) {
	enum, ok := mappingCreateOdaInstanceDetailsShapeNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
