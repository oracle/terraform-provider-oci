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

// RecommendedPatchSummary Summary information about an recommended patches.
type RecommendedPatchSummary struct {

	// The OCID of the patch.
	PatchId *string `mandatory:"true" json:"patchId"`

	// Name of the patch
	PatchName *string `mandatory:"true" json:"patchName"`

	// Description of the patch
	PatchDescription *string `mandatory:"true" json:"patchDescription"`

	// Date on which the patch was released
	TimeReleased *common.SDKTime `mandatory:"false" json:"timeReleased"`

	// Description of the patch
	PatchType *string `mandatory:"false" json:"patchType"`

	// Patch Severity.
	Severity PatchSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// Patch Level.
	PatchLevel RecommendedPatchSummaryPatchLevelEnum `mandatory:"false" json:"patchLevel,omitempty"`
}

func (m RecommendedPatchSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecommendedPatchSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetPatchSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRecommendedPatchSummaryPatchLevelEnum(string(m.PatchLevel)); !ok && m.PatchLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchLevel: %s. Supported values are: %s.", m.PatchLevel, strings.Join(GetRecommendedPatchSummaryPatchLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RecommendedPatchSummaryPatchLevelEnum Enum with underlying type: string
type RecommendedPatchSummaryPatchLevelEnum string

// Set of constants representing the allowable values for RecommendedPatchSummaryPatchLevelEnum
const (
	RecommendedPatchSummaryPatchLevelLatest         RecommendedPatchSummaryPatchLevelEnum = "LATEST"
	RecommendedPatchSummaryPatchLevelLatestMinusOne RecommendedPatchSummaryPatchLevelEnum = "LATEST_MINUS_ONE"
	RecommendedPatchSummaryPatchLevelLatestMinusTwo RecommendedPatchSummaryPatchLevelEnum = "LATEST_MINUS_TWO"
)

var mappingRecommendedPatchSummaryPatchLevelEnum = map[string]RecommendedPatchSummaryPatchLevelEnum{
	"LATEST":           RecommendedPatchSummaryPatchLevelLatest,
	"LATEST_MINUS_ONE": RecommendedPatchSummaryPatchLevelLatestMinusOne,
	"LATEST_MINUS_TWO": RecommendedPatchSummaryPatchLevelLatestMinusTwo,
}

var mappingRecommendedPatchSummaryPatchLevelEnumLowerCase = map[string]RecommendedPatchSummaryPatchLevelEnum{
	"latest":           RecommendedPatchSummaryPatchLevelLatest,
	"latest_minus_one": RecommendedPatchSummaryPatchLevelLatestMinusOne,
	"latest_minus_two": RecommendedPatchSummaryPatchLevelLatestMinusTwo,
}

// GetRecommendedPatchSummaryPatchLevelEnumValues Enumerates the set of values for RecommendedPatchSummaryPatchLevelEnum
func GetRecommendedPatchSummaryPatchLevelEnumValues() []RecommendedPatchSummaryPatchLevelEnum {
	values := make([]RecommendedPatchSummaryPatchLevelEnum, 0)
	for _, v := range mappingRecommendedPatchSummaryPatchLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetRecommendedPatchSummaryPatchLevelEnumStringValues Enumerates the set of values in String for RecommendedPatchSummaryPatchLevelEnum
func GetRecommendedPatchSummaryPatchLevelEnumStringValues() []string {
	return []string{
		"LATEST",
		"LATEST_MINUS_ONE",
		"LATEST_MINUS_TWO",
	}
}

// GetMappingRecommendedPatchSummaryPatchLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecommendedPatchSummaryPatchLevelEnum(val string) (RecommendedPatchSummaryPatchLevelEnum, bool) {
	enum, ok := mappingRecommendedPatchSummaryPatchLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
