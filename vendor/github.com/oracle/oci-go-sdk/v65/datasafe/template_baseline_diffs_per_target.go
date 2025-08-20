// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// TemplateBaselineDiffsPerTarget The results of the comparison between two security assessment resources, and one of them is TEMPLATE_BASELINE type.
type TemplateBaselineDiffsPerTarget struct {

	// The OCID of the target database.
	TargetId *string `mandatory:"false" json:"targetId"`

	// A unique identifier for the finding. This is common for the finding across targets.
	Key *string `mandatory:"false" json:"key"`

	// The short title for the finding.
	Title *string `mandatory:"false" json:"title"`

	// The severity of this diff.
	Severity TemplateBaselineDiffsPerTargetSeverityEnum `mandatory:"false" json:"severity,omitempty"`
}

func (m TemplateBaselineDiffsPerTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TemplateBaselineDiffsPerTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTemplateBaselineDiffsPerTargetSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetTemplateBaselineDiffsPerTargetSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TemplateBaselineDiffsPerTargetSeverityEnum Enum with underlying type: string
type TemplateBaselineDiffsPerTargetSeverityEnum string

// Set of constants representing the allowable values for TemplateBaselineDiffsPerTargetSeverityEnum
const (
	TemplateBaselineDiffsPerTargetSeverityHigh     TemplateBaselineDiffsPerTargetSeverityEnum = "HIGH"
	TemplateBaselineDiffsPerTargetSeverityMedium   TemplateBaselineDiffsPerTargetSeverityEnum = "MEDIUM"
	TemplateBaselineDiffsPerTargetSeverityLow      TemplateBaselineDiffsPerTargetSeverityEnum = "LOW"
	TemplateBaselineDiffsPerTargetSeverityEvaluate TemplateBaselineDiffsPerTargetSeverityEnum = "EVALUATE"
	TemplateBaselineDiffsPerTargetSeverityAdvisory TemplateBaselineDiffsPerTargetSeverityEnum = "ADVISORY"
	TemplateBaselineDiffsPerTargetSeverityPass     TemplateBaselineDiffsPerTargetSeverityEnum = "PASS"
	TemplateBaselineDiffsPerTargetSeverityDeferred TemplateBaselineDiffsPerTargetSeverityEnum = "DEFERRED"
)

var mappingTemplateBaselineDiffsPerTargetSeverityEnum = map[string]TemplateBaselineDiffsPerTargetSeverityEnum{
	"HIGH":     TemplateBaselineDiffsPerTargetSeverityHigh,
	"MEDIUM":   TemplateBaselineDiffsPerTargetSeverityMedium,
	"LOW":      TemplateBaselineDiffsPerTargetSeverityLow,
	"EVALUATE": TemplateBaselineDiffsPerTargetSeverityEvaluate,
	"ADVISORY": TemplateBaselineDiffsPerTargetSeverityAdvisory,
	"PASS":     TemplateBaselineDiffsPerTargetSeverityPass,
	"DEFERRED": TemplateBaselineDiffsPerTargetSeverityDeferred,
}

var mappingTemplateBaselineDiffsPerTargetSeverityEnumLowerCase = map[string]TemplateBaselineDiffsPerTargetSeverityEnum{
	"high":     TemplateBaselineDiffsPerTargetSeverityHigh,
	"medium":   TemplateBaselineDiffsPerTargetSeverityMedium,
	"low":      TemplateBaselineDiffsPerTargetSeverityLow,
	"evaluate": TemplateBaselineDiffsPerTargetSeverityEvaluate,
	"advisory": TemplateBaselineDiffsPerTargetSeverityAdvisory,
	"pass":     TemplateBaselineDiffsPerTargetSeverityPass,
	"deferred": TemplateBaselineDiffsPerTargetSeverityDeferred,
}

// GetTemplateBaselineDiffsPerTargetSeverityEnumValues Enumerates the set of values for TemplateBaselineDiffsPerTargetSeverityEnum
func GetTemplateBaselineDiffsPerTargetSeverityEnumValues() []TemplateBaselineDiffsPerTargetSeverityEnum {
	values := make([]TemplateBaselineDiffsPerTargetSeverityEnum, 0)
	for _, v := range mappingTemplateBaselineDiffsPerTargetSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetTemplateBaselineDiffsPerTargetSeverityEnumStringValues Enumerates the set of values in String for TemplateBaselineDiffsPerTargetSeverityEnum
func GetTemplateBaselineDiffsPerTargetSeverityEnumStringValues() []string {
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

// GetMappingTemplateBaselineDiffsPerTargetSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTemplateBaselineDiffsPerTargetSeverityEnum(val string) (TemplateBaselineDiffsPerTargetSeverityEnum, bool) {
	enum, ok := mappingTemplateBaselineDiffsPerTargetSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
