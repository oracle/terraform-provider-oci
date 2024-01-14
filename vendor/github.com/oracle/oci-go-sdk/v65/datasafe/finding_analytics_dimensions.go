// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FindingAnalyticsDimensions The scope of analytics data.
type FindingAnalyticsDimensions struct {

	// Each finding in security assessment has an associated key (think of key as a finding's name).
	// For a given finding, the key will be the same across targets. The user can use these keys to filter the findings.
	Key *string `mandatory:"false" json:"key"`

	// The category of the top finding.
	TopFindingCategory *string `mandatory:"false" json:"topFindingCategory"`

	// The short title of the finding.
	Title *string `mandatory:"false" json:"title"`

	// The status of the top finding.
	// All findings will have "severity" to indicate the risk level, but only top findings will have "status".
	// Possible status: Pass / Risk (Low, Medium, High)/ Evaluate / Advisory / Deferred
	// Instead of having "Low, Medium, High" in severity, "Risk" will include these three situations in status.
	TopFindingStatus FindingAnalyticsDimensionsTopFindingStatusEnum `mandatory:"false" json:"topFindingStatus,omitempty"`

	// The severity (risk level) of the finding.
	Severity FindingAnalyticsDimensionsSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// The OCID of the target database.
	TargetId *string `mandatory:"false" json:"targetId"`
}

func (m FindingAnalyticsDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FindingAnalyticsDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFindingAnalyticsDimensionsTopFindingStatusEnum(string(m.TopFindingStatus)); !ok && m.TopFindingStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TopFindingStatus: %s. Supported values are: %s.", m.TopFindingStatus, strings.Join(GetFindingAnalyticsDimensionsTopFindingStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFindingAnalyticsDimensionsSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetFindingAnalyticsDimensionsSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FindingAnalyticsDimensionsTopFindingStatusEnum Enum with underlying type: string
type FindingAnalyticsDimensionsTopFindingStatusEnum string

// Set of constants representing the allowable values for FindingAnalyticsDimensionsTopFindingStatusEnum
const (
	FindingAnalyticsDimensionsTopFindingStatusRisk     FindingAnalyticsDimensionsTopFindingStatusEnum = "RISK"
	FindingAnalyticsDimensionsTopFindingStatusEvaluate FindingAnalyticsDimensionsTopFindingStatusEnum = "EVALUATE"
	FindingAnalyticsDimensionsTopFindingStatusAdvisory FindingAnalyticsDimensionsTopFindingStatusEnum = "ADVISORY"
	FindingAnalyticsDimensionsTopFindingStatusPass     FindingAnalyticsDimensionsTopFindingStatusEnum = "PASS"
	FindingAnalyticsDimensionsTopFindingStatusDeferred FindingAnalyticsDimensionsTopFindingStatusEnum = "DEFERRED"
)

var mappingFindingAnalyticsDimensionsTopFindingStatusEnum = map[string]FindingAnalyticsDimensionsTopFindingStatusEnum{
	"RISK":     FindingAnalyticsDimensionsTopFindingStatusRisk,
	"EVALUATE": FindingAnalyticsDimensionsTopFindingStatusEvaluate,
	"ADVISORY": FindingAnalyticsDimensionsTopFindingStatusAdvisory,
	"PASS":     FindingAnalyticsDimensionsTopFindingStatusPass,
	"DEFERRED": FindingAnalyticsDimensionsTopFindingStatusDeferred,
}

var mappingFindingAnalyticsDimensionsTopFindingStatusEnumLowerCase = map[string]FindingAnalyticsDimensionsTopFindingStatusEnum{
	"risk":     FindingAnalyticsDimensionsTopFindingStatusRisk,
	"evaluate": FindingAnalyticsDimensionsTopFindingStatusEvaluate,
	"advisory": FindingAnalyticsDimensionsTopFindingStatusAdvisory,
	"pass":     FindingAnalyticsDimensionsTopFindingStatusPass,
	"deferred": FindingAnalyticsDimensionsTopFindingStatusDeferred,
}

// GetFindingAnalyticsDimensionsTopFindingStatusEnumValues Enumerates the set of values for FindingAnalyticsDimensionsTopFindingStatusEnum
func GetFindingAnalyticsDimensionsTopFindingStatusEnumValues() []FindingAnalyticsDimensionsTopFindingStatusEnum {
	values := make([]FindingAnalyticsDimensionsTopFindingStatusEnum, 0)
	for _, v := range mappingFindingAnalyticsDimensionsTopFindingStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetFindingAnalyticsDimensionsTopFindingStatusEnumStringValues Enumerates the set of values in String for FindingAnalyticsDimensionsTopFindingStatusEnum
func GetFindingAnalyticsDimensionsTopFindingStatusEnumStringValues() []string {
	return []string{
		"RISK",
		"EVALUATE",
		"ADVISORY",
		"PASS",
		"DEFERRED",
	}
}

// GetMappingFindingAnalyticsDimensionsTopFindingStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFindingAnalyticsDimensionsTopFindingStatusEnum(val string) (FindingAnalyticsDimensionsTopFindingStatusEnum, bool) {
	enum, ok := mappingFindingAnalyticsDimensionsTopFindingStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FindingAnalyticsDimensionsSeverityEnum Enum with underlying type: string
type FindingAnalyticsDimensionsSeverityEnum string

// Set of constants representing the allowable values for FindingAnalyticsDimensionsSeverityEnum
const (
	FindingAnalyticsDimensionsSeverityHigh     FindingAnalyticsDimensionsSeverityEnum = "HIGH"
	FindingAnalyticsDimensionsSeverityMedium   FindingAnalyticsDimensionsSeverityEnum = "MEDIUM"
	FindingAnalyticsDimensionsSeverityLow      FindingAnalyticsDimensionsSeverityEnum = "LOW"
	FindingAnalyticsDimensionsSeverityEvaluate FindingAnalyticsDimensionsSeverityEnum = "EVALUATE"
	FindingAnalyticsDimensionsSeverityAdvisory FindingAnalyticsDimensionsSeverityEnum = "ADVISORY"
	FindingAnalyticsDimensionsSeverityPass     FindingAnalyticsDimensionsSeverityEnum = "PASS"
	FindingAnalyticsDimensionsSeverityDeferred FindingAnalyticsDimensionsSeverityEnum = "DEFERRED"
)

var mappingFindingAnalyticsDimensionsSeverityEnum = map[string]FindingAnalyticsDimensionsSeverityEnum{
	"HIGH":     FindingAnalyticsDimensionsSeverityHigh,
	"MEDIUM":   FindingAnalyticsDimensionsSeverityMedium,
	"LOW":      FindingAnalyticsDimensionsSeverityLow,
	"EVALUATE": FindingAnalyticsDimensionsSeverityEvaluate,
	"ADVISORY": FindingAnalyticsDimensionsSeverityAdvisory,
	"PASS":     FindingAnalyticsDimensionsSeverityPass,
	"DEFERRED": FindingAnalyticsDimensionsSeverityDeferred,
}

var mappingFindingAnalyticsDimensionsSeverityEnumLowerCase = map[string]FindingAnalyticsDimensionsSeverityEnum{
	"high":     FindingAnalyticsDimensionsSeverityHigh,
	"medium":   FindingAnalyticsDimensionsSeverityMedium,
	"low":      FindingAnalyticsDimensionsSeverityLow,
	"evaluate": FindingAnalyticsDimensionsSeverityEvaluate,
	"advisory": FindingAnalyticsDimensionsSeverityAdvisory,
	"pass":     FindingAnalyticsDimensionsSeverityPass,
	"deferred": FindingAnalyticsDimensionsSeverityDeferred,
}

// GetFindingAnalyticsDimensionsSeverityEnumValues Enumerates the set of values for FindingAnalyticsDimensionsSeverityEnum
func GetFindingAnalyticsDimensionsSeverityEnumValues() []FindingAnalyticsDimensionsSeverityEnum {
	values := make([]FindingAnalyticsDimensionsSeverityEnum, 0)
	for _, v := range mappingFindingAnalyticsDimensionsSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetFindingAnalyticsDimensionsSeverityEnumStringValues Enumerates the set of values in String for FindingAnalyticsDimensionsSeverityEnum
func GetFindingAnalyticsDimensionsSeverityEnumStringValues() []string {
	return []string{
		"HIGH",
		"MEDIUM",
		"LOW",
		"EVALUATE",
		"ADVISORY",
		"PASS",
		"DEFERRED",
	}
}

// GetMappingFindingAnalyticsDimensionsSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFindingAnalyticsDimensionsSeverityEnum(val string) (FindingAnalyticsDimensionsSeverityEnum, bool) {
	enum, ok := mappingFindingAnalyticsDimensionsSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
