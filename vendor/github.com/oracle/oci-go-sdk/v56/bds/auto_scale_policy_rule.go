// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

// AutoScalePolicyRuleActionEnum Enum with underlying type: string
type AutoScalePolicyRuleActionEnum string

// Set of constants representing the allowable values for AutoScalePolicyRuleActionEnum
const (
	AutoScalePolicyRuleActionUp   AutoScalePolicyRuleActionEnum = "CHANGE_SHAPE_SCALE_UP"
	AutoScalePolicyRuleActionDown AutoScalePolicyRuleActionEnum = "CHANGE_SHAPE_SCALE_DOWN"
)

var mappingAutoScalePolicyRuleAction = map[string]AutoScalePolicyRuleActionEnum{
	"CHANGE_SHAPE_SCALE_UP":   AutoScalePolicyRuleActionUp,
	"CHANGE_SHAPE_SCALE_DOWN": AutoScalePolicyRuleActionDown,
}

// GetAutoScalePolicyRuleActionEnumValues Enumerates the set of values for AutoScalePolicyRuleActionEnum
func GetAutoScalePolicyRuleActionEnumValues() []AutoScalePolicyRuleActionEnum {
	values := make([]AutoScalePolicyRuleActionEnum, 0)
	for _, v := range mappingAutoScalePolicyRuleAction {
		values = append(values, v)
	}
	return values
}
