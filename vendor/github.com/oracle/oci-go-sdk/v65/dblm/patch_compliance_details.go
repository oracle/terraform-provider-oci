// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchComplianceDetails Patch Compliance Status
type PatchComplianceDetails struct {

	// Patch compliance status.
	PatchComplianceStatus PatchComplianceDetailsPatchComplianceStatusEnum `mandatory:"false" json:"patchComplianceStatus,omitempty"`

	// Resource patch compliance version name.
	PatchComplianceVersion *string `mandatory:"false" json:"patchComplianceVersion"`
}

func (m PatchComplianceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchComplianceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchComplianceDetailsPatchComplianceStatusEnum(string(m.PatchComplianceStatus)); !ok && m.PatchComplianceStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchComplianceStatus: %s. Supported values are: %s.", m.PatchComplianceStatus, strings.Join(GetPatchComplianceDetailsPatchComplianceStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchComplianceDetailsPatchComplianceStatusEnum Enum with underlying type: string
type PatchComplianceDetailsPatchComplianceStatusEnum string

// Set of constants representing the allowable values for PatchComplianceDetailsPatchComplianceStatusEnum
const (
	PatchComplianceDetailsPatchComplianceStatusGreen  PatchComplianceDetailsPatchComplianceStatusEnum = "GREEN"
	PatchComplianceDetailsPatchComplianceStatusYellow PatchComplianceDetailsPatchComplianceStatusEnum = "YELLOW"
	PatchComplianceDetailsPatchComplianceStatusRed    PatchComplianceDetailsPatchComplianceStatusEnum = "RED"
)

var mappingPatchComplianceDetailsPatchComplianceStatusEnum = map[string]PatchComplianceDetailsPatchComplianceStatusEnum{
	"GREEN":  PatchComplianceDetailsPatchComplianceStatusGreen,
	"YELLOW": PatchComplianceDetailsPatchComplianceStatusYellow,
	"RED":    PatchComplianceDetailsPatchComplianceStatusRed,
}

var mappingPatchComplianceDetailsPatchComplianceStatusEnumLowerCase = map[string]PatchComplianceDetailsPatchComplianceStatusEnum{
	"green":  PatchComplianceDetailsPatchComplianceStatusGreen,
	"yellow": PatchComplianceDetailsPatchComplianceStatusYellow,
	"red":    PatchComplianceDetailsPatchComplianceStatusRed,
}

// GetPatchComplianceDetailsPatchComplianceStatusEnumValues Enumerates the set of values for PatchComplianceDetailsPatchComplianceStatusEnum
func GetPatchComplianceDetailsPatchComplianceStatusEnumValues() []PatchComplianceDetailsPatchComplianceStatusEnum {
	values := make([]PatchComplianceDetailsPatchComplianceStatusEnum, 0)
	for _, v := range mappingPatchComplianceDetailsPatchComplianceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchComplianceDetailsPatchComplianceStatusEnumStringValues Enumerates the set of values in String for PatchComplianceDetailsPatchComplianceStatusEnum
func GetPatchComplianceDetailsPatchComplianceStatusEnumStringValues() []string {
	return []string{
		"GREEN",
		"YELLOW",
		"RED",
	}
}

// GetMappingPatchComplianceDetailsPatchComplianceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchComplianceDetailsPatchComplianceStatusEnum(val string) (PatchComplianceDetailsPatchComplianceStatusEnum, bool) {
	enum, ok := mappingPatchComplianceDetailsPatchComplianceStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
