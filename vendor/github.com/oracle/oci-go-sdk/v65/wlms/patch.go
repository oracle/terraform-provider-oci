// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Patch A WebLogic patch.
type Patch struct {

	// The ID of the WebLogic patch.
	Id *string `mandatory:"true" json:"id"`

	// The name of the WebLogic patch.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The description of the WebLogic patch.
	Description *string `mandatory:"true" json:"description"`

	// The WebLogic version for this patch. The patch can be installed to domains with this version.
	WeblogicVersion *string `mandatory:"true" json:"weblogicVersion"`

	// The type of middleware for which this patch is applicable. A patch can be applicable to more than one type of middleware.
	MiddlewareType []PatchMiddlewareTypeEnum `mandatory:"true" json:"middlewareType"`

	// The operating system architecture for which the patch can be applied.
	OsArch *string `mandatory:"true" json:"osArch"`
}

func (m Patch) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Patch) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.MiddlewareType {
		if _, ok := GetMappingPatchMiddlewareTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MiddlewareType: %s. Supported values are: %s.", val, strings.Join(GetPatchMiddlewareTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchMiddlewareTypeEnum Enum with underlying type: string
type PatchMiddlewareTypeEnum string

// Set of constants representing the allowable values for PatchMiddlewareTypeEnum
const (
	PatchMiddlewareTypeFmw    PatchMiddlewareTypeEnum = "FMW"
	PatchMiddlewareTypeWls    PatchMiddlewareTypeEnum = "WLS"
	PatchMiddlewareTypeOpatch PatchMiddlewareTypeEnum = "OPATCH"
)

var mappingPatchMiddlewareTypeEnum = map[string]PatchMiddlewareTypeEnum{
	"FMW":    PatchMiddlewareTypeFmw,
	"WLS":    PatchMiddlewareTypeWls,
	"OPATCH": PatchMiddlewareTypeOpatch,
}

var mappingPatchMiddlewareTypeEnumLowerCase = map[string]PatchMiddlewareTypeEnum{
	"fmw":    PatchMiddlewareTypeFmw,
	"wls":    PatchMiddlewareTypeWls,
	"opatch": PatchMiddlewareTypeOpatch,
}

// GetPatchMiddlewareTypeEnumValues Enumerates the set of values for PatchMiddlewareTypeEnum
func GetPatchMiddlewareTypeEnumValues() []PatchMiddlewareTypeEnum {
	values := make([]PatchMiddlewareTypeEnum, 0)
	for _, v := range mappingPatchMiddlewareTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchMiddlewareTypeEnumStringValues Enumerates the set of values in String for PatchMiddlewareTypeEnum
func GetPatchMiddlewareTypeEnumStringValues() []string {
	return []string{
		"FMW",
		"WLS",
		"OPATCH",
	}
}

// GetMappingPatchMiddlewareTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchMiddlewareTypeEnum(val string) (PatchMiddlewareTypeEnum, bool) {
	enum, ok := mappingPatchMiddlewareTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
