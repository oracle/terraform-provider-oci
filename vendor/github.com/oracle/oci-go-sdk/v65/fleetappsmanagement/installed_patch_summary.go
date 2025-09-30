// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstalledPatchSummary Summary information about an installed patches.
type InstalledPatchSummary struct {

	// The OCID of the patch.
	PatchId *string `mandatory:"true" json:"patchId"`

	// Name of the patch
	PatchName *string `mandatory:"true" json:"patchName"`

	// Description of the patch
	PatchDescription *string `mandatory:"true" json:"patchDescription"`

	// Date on which the patch was applied to the target
	TimeApplied *common.SDKTime `mandatory:"false" json:"timeApplied"`

	// Date on which the patch was released
	TimeReleased *common.SDKTime `mandatory:"false" json:"timeReleased"`

	// Description of the patch
	PatchType *string `mandatory:"false" json:"patchType"`

	// Patch Name.
	PatchLevel InstalledPatchSummaryPatchLevelEnum `mandatory:"false" json:"patchLevel,omitempty"`

	// Patch Severity.
	Severity PatchSeverityEnum `mandatory:"false" json:"severity,omitempty"`
}

func (m InstalledPatchSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstalledPatchSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInstalledPatchSummaryPatchLevelEnum(string(m.PatchLevel)); !ok && m.PatchLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchLevel: %s. Supported values are: %s.", m.PatchLevel, strings.Join(GetInstalledPatchSummaryPatchLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetPatchSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstalledPatchSummaryPatchLevelEnum Enum with underlying type: string
type InstalledPatchSummaryPatchLevelEnum string

// Set of constants representing the allowable values for InstalledPatchSummaryPatchLevelEnum
const (
	InstalledPatchSummaryPatchLevelLatest         InstalledPatchSummaryPatchLevelEnum = "LATEST"
	InstalledPatchSummaryPatchLevelLatestMinusOne InstalledPatchSummaryPatchLevelEnum = "LATEST_MINUS_ONE"
	InstalledPatchSummaryPatchLevelLatestMinusTwo InstalledPatchSummaryPatchLevelEnum = "LATEST_MINUS_TWO"
)

var mappingInstalledPatchSummaryPatchLevelEnum = map[string]InstalledPatchSummaryPatchLevelEnum{
	"LATEST":           InstalledPatchSummaryPatchLevelLatest,
	"LATEST_MINUS_ONE": InstalledPatchSummaryPatchLevelLatestMinusOne,
	"LATEST_MINUS_TWO": InstalledPatchSummaryPatchLevelLatestMinusTwo,
}

var mappingInstalledPatchSummaryPatchLevelEnumLowerCase = map[string]InstalledPatchSummaryPatchLevelEnum{
	"latest":           InstalledPatchSummaryPatchLevelLatest,
	"latest_minus_one": InstalledPatchSummaryPatchLevelLatestMinusOne,
	"latest_minus_two": InstalledPatchSummaryPatchLevelLatestMinusTwo,
}

// GetInstalledPatchSummaryPatchLevelEnumValues Enumerates the set of values for InstalledPatchSummaryPatchLevelEnum
func GetInstalledPatchSummaryPatchLevelEnumValues() []InstalledPatchSummaryPatchLevelEnum {
	values := make([]InstalledPatchSummaryPatchLevelEnum, 0)
	for _, v := range mappingInstalledPatchSummaryPatchLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetInstalledPatchSummaryPatchLevelEnumStringValues Enumerates the set of values in String for InstalledPatchSummaryPatchLevelEnum
func GetInstalledPatchSummaryPatchLevelEnumStringValues() []string {
	return []string{
		"LATEST",
		"LATEST_MINUS_ONE",
		"LATEST_MINUS_TWO",
	}
}

// GetMappingInstalledPatchSummaryPatchLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstalledPatchSummaryPatchLevelEnum(val string) (InstalledPatchSummaryPatchLevelEnum, bool) {
	enum, ok := mappingInstalledPatchSummaryPatchLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
