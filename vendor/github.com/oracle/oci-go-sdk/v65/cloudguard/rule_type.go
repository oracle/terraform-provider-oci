// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RuleType The detector rule types.
type RuleType struct {

	// The unique identifier of the detector rule type
	Key RuleTypeKeyEnum `mandatory:"true" json:"key"`

	// Detector rule type value
	Value []string `mandatory:"true" json:"value"`
}

func (m RuleType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RuleType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRuleTypeKeyEnum(string(m.Key)); !ok && m.Key != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Key: %s. Supported values are: %s.", m.Key, strings.Join(GetRuleTypeKeyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RuleTypeKeyEnum Enum with underlying type: string
type RuleTypeKeyEnum string

// Set of constants representing the allowable values for RuleTypeKeyEnum
const (
	RuleTypeKeyCategory RuleTypeKeyEnum = "CATEGORY"
)

var mappingRuleTypeKeyEnum = map[string]RuleTypeKeyEnum{
	"CATEGORY": RuleTypeKeyCategory,
}

var mappingRuleTypeKeyEnumLowerCase = map[string]RuleTypeKeyEnum{
	"category": RuleTypeKeyCategory,
}

// GetRuleTypeKeyEnumValues Enumerates the set of values for RuleTypeKeyEnum
func GetRuleTypeKeyEnumValues() []RuleTypeKeyEnum {
	values := make([]RuleTypeKeyEnum, 0)
	for _, v := range mappingRuleTypeKeyEnum {
		values = append(values, v)
	}
	return values
}

// GetRuleTypeKeyEnumStringValues Enumerates the set of values in String for RuleTypeKeyEnum
func GetRuleTypeKeyEnumStringValues() []string {
	return []string{
		"CATEGORY",
	}
}

// GetMappingRuleTypeKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuleTypeKeyEnum(val string) (RuleTypeKeyEnum, bool) {
	enum, ok := mappingRuleTypeKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
