// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutoScalePolicyRule A rule that defines a specific autoscale action to take and the metric that triggers that action.
type AutoScalePolicyRule struct {

	// The valid value are CHANGE_SHAPE_SCALE_UP or CHANGE_SHAPE_SCALE_DOWN.
	Action AutoScalePolicyRuleActionEnum `mandatory:"true" json:"action"`

	Metric *AutoScalePolicyMetricRule `mandatory:"true" json:"metric"`
}

func (m AutoScalePolicyRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoScalePolicyRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutoScalePolicyRuleActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetAutoScalePolicyRuleActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutoScalePolicyRuleActionEnum Enum with underlying type: string
type AutoScalePolicyRuleActionEnum string

// Set of constants representing the allowable values for AutoScalePolicyRuleActionEnum
const (
	AutoScalePolicyRuleActionUp   AutoScalePolicyRuleActionEnum = "CHANGE_SHAPE_SCALE_UP"
	AutoScalePolicyRuleActionDown AutoScalePolicyRuleActionEnum = "CHANGE_SHAPE_SCALE_DOWN"
)

var mappingAutoScalePolicyRuleActionEnum = map[string]AutoScalePolicyRuleActionEnum{
	"CHANGE_SHAPE_SCALE_UP":   AutoScalePolicyRuleActionUp,
	"CHANGE_SHAPE_SCALE_DOWN": AutoScalePolicyRuleActionDown,
}

var mappingAutoScalePolicyRuleActionEnumLowerCase = map[string]AutoScalePolicyRuleActionEnum{
	"change_shape_scale_up":   AutoScalePolicyRuleActionUp,
	"change_shape_scale_down": AutoScalePolicyRuleActionDown,
}

// GetAutoScalePolicyRuleActionEnumValues Enumerates the set of values for AutoScalePolicyRuleActionEnum
func GetAutoScalePolicyRuleActionEnumValues() []AutoScalePolicyRuleActionEnum {
	values := make([]AutoScalePolicyRuleActionEnum, 0)
	for _, v := range mappingAutoScalePolicyRuleActionEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoScalePolicyRuleActionEnumStringValues Enumerates the set of values in String for AutoScalePolicyRuleActionEnum
func GetAutoScalePolicyRuleActionEnumStringValues() []string {
	return []string{
		"CHANGE_SHAPE_SCALE_UP",
		"CHANGE_SHAPE_SCALE_DOWN",
	}
}

// GetMappingAutoScalePolicyRuleActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoScalePolicyRuleActionEnum(val string) (AutoScalePolicyRuleActionEnum, bool) {
	enum, ok := mappingAutoScalePolicyRuleActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
