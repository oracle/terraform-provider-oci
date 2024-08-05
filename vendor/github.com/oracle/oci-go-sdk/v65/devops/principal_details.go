// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrincipalDetails The principal details
type PrincipalDetails struct {

	// the OCID of the principal
	PrincipalId *string `mandatory:"true" json:"principalId"`

	// the name of the principal
	PrincipalName *string `mandatory:"false" json:"principalName"`

	// the type of principal
	PrincipalType PrincipalDetailsPrincipalTypeEnum `mandatory:"false" json:"principalType,omitempty"`

	// The state of the principal, it can be active or inactive or suppressed for emails
	PrincipalState PrincipalDetailsPrincipalStateEnum `mandatory:"false" json:"principalState,omitempty"`
}

func (m PrincipalDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrincipalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPrincipalDetailsPrincipalTypeEnum(string(m.PrincipalType)); !ok && m.PrincipalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrincipalType: %s. Supported values are: %s.", m.PrincipalType, strings.Join(GetPrincipalDetailsPrincipalTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPrincipalDetailsPrincipalStateEnum(string(m.PrincipalState)); !ok && m.PrincipalState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrincipalState: %s. Supported values are: %s.", m.PrincipalState, strings.Join(GetPrincipalDetailsPrincipalStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PrincipalDetailsPrincipalTypeEnum Enum with underlying type: string
type PrincipalDetailsPrincipalTypeEnum string

// Set of constants representing the allowable values for PrincipalDetailsPrincipalTypeEnum
const (
	PrincipalDetailsPrincipalTypeService  PrincipalDetailsPrincipalTypeEnum = "SERVICE"
	PrincipalDetailsPrincipalTypeUser     PrincipalDetailsPrincipalTypeEnum = "USER"
	PrincipalDetailsPrincipalTypeInstance PrincipalDetailsPrincipalTypeEnum = "INSTANCE"
	PrincipalDetailsPrincipalTypeResource PrincipalDetailsPrincipalTypeEnum = "RESOURCE"
)

var mappingPrincipalDetailsPrincipalTypeEnum = map[string]PrincipalDetailsPrincipalTypeEnum{
	"SERVICE":  PrincipalDetailsPrincipalTypeService,
	"USER":     PrincipalDetailsPrincipalTypeUser,
	"INSTANCE": PrincipalDetailsPrincipalTypeInstance,
	"RESOURCE": PrincipalDetailsPrincipalTypeResource,
}

var mappingPrincipalDetailsPrincipalTypeEnumLowerCase = map[string]PrincipalDetailsPrincipalTypeEnum{
	"service":  PrincipalDetailsPrincipalTypeService,
	"user":     PrincipalDetailsPrincipalTypeUser,
	"instance": PrincipalDetailsPrincipalTypeInstance,
	"resource": PrincipalDetailsPrincipalTypeResource,
}

// GetPrincipalDetailsPrincipalTypeEnumValues Enumerates the set of values for PrincipalDetailsPrincipalTypeEnum
func GetPrincipalDetailsPrincipalTypeEnumValues() []PrincipalDetailsPrincipalTypeEnum {
	values := make([]PrincipalDetailsPrincipalTypeEnum, 0)
	for _, v := range mappingPrincipalDetailsPrincipalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPrincipalDetailsPrincipalTypeEnumStringValues Enumerates the set of values in String for PrincipalDetailsPrincipalTypeEnum
func GetPrincipalDetailsPrincipalTypeEnumStringValues() []string {
	return []string{
		"SERVICE",
		"USER",
		"INSTANCE",
		"RESOURCE",
	}
}

// GetMappingPrincipalDetailsPrincipalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrincipalDetailsPrincipalTypeEnum(val string) (PrincipalDetailsPrincipalTypeEnum, bool) {
	enum, ok := mappingPrincipalDetailsPrincipalTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PrincipalDetailsPrincipalStateEnum Enum with underlying type: string
type PrincipalDetailsPrincipalStateEnum string

// Set of constants representing the allowable values for PrincipalDetailsPrincipalStateEnum
const (
	PrincipalDetailsPrincipalStateActive     PrincipalDetailsPrincipalStateEnum = "ACTIVE"
	PrincipalDetailsPrincipalStateInactive   PrincipalDetailsPrincipalStateEnum = "INACTIVE"
	PrincipalDetailsPrincipalStateSuppressed PrincipalDetailsPrincipalStateEnum = "SUPPRESSED"
)

var mappingPrincipalDetailsPrincipalStateEnum = map[string]PrincipalDetailsPrincipalStateEnum{
	"ACTIVE":     PrincipalDetailsPrincipalStateActive,
	"INACTIVE":   PrincipalDetailsPrincipalStateInactive,
	"SUPPRESSED": PrincipalDetailsPrincipalStateSuppressed,
}

var mappingPrincipalDetailsPrincipalStateEnumLowerCase = map[string]PrincipalDetailsPrincipalStateEnum{
	"active":     PrincipalDetailsPrincipalStateActive,
	"inactive":   PrincipalDetailsPrincipalStateInactive,
	"suppressed": PrincipalDetailsPrincipalStateSuppressed,
}

// GetPrincipalDetailsPrincipalStateEnumValues Enumerates the set of values for PrincipalDetailsPrincipalStateEnum
func GetPrincipalDetailsPrincipalStateEnumValues() []PrincipalDetailsPrincipalStateEnum {
	values := make([]PrincipalDetailsPrincipalStateEnum, 0)
	for _, v := range mappingPrincipalDetailsPrincipalStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPrincipalDetailsPrincipalStateEnumStringValues Enumerates the set of values in String for PrincipalDetailsPrincipalStateEnum
func GetPrincipalDetailsPrincipalStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"SUPPRESSED",
	}
}

// GetMappingPrincipalDetailsPrincipalStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrincipalDetailsPrincipalStateEnum(val string) (PrincipalDetailsPrincipalStateEnum, bool) {
	enum, ok := mappingPrincipalDetailsPrincipalStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
