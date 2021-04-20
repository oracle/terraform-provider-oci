// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// API for the Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service
// build on Hadoop, Spark and Data Science distribution, which can be fully integrated with existing enterprise
// data in Oracle Database and Oracle Applications..
//

package bds

import (
	"github.com/oracle/oci-go-sdk/v40/common"
)

// AutoScalePolicy Policy definitions for the autoscaling configuration
type AutoScalePolicy struct {

	// Types of autoscaling policies. SCHEDULE-BASED or  THRESHOLD-BASED, current only supported THRESHOLD-BASED.
	PolicyType AutoScalePolicyPolicyTypeEnum `mandatory:"true" json:"policyType"`

	// The list of rules for autoscaling. If an action have multiple rules, last rule in the array will be applied.
	Rules []AutoScalePolicyRule `mandatory:"true" json:"rules"`
}

func (m AutoScalePolicy) String() string {
	return common.PointerString(m)
}

// AutoScalePolicyPolicyTypeEnum Enum with underlying type: string
type AutoScalePolicyPolicyTypeEnum string

// Set of constants representing the allowable values for AutoScalePolicyPolicyTypeEnum
const (
	AutoScalePolicyPolicyTypeThresholdBased AutoScalePolicyPolicyTypeEnum = "THRESHOLD_BASED"
	AutoScalePolicyPolicyTypeScheduleBased  AutoScalePolicyPolicyTypeEnum = "SCHEDULE_BASED"
)

var mappingAutoScalePolicyPolicyType = map[string]AutoScalePolicyPolicyTypeEnum{
	"THRESHOLD_BASED": AutoScalePolicyPolicyTypeThresholdBased,
	"SCHEDULE_BASED":  AutoScalePolicyPolicyTypeScheduleBased,
}

// GetAutoScalePolicyPolicyTypeEnumValues Enumerates the set of values for AutoScalePolicyPolicyTypeEnum
func GetAutoScalePolicyPolicyTypeEnumValues() []AutoScalePolicyPolicyTypeEnum {
	values := make([]AutoScalePolicyPolicyTypeEnum, 0)
	for _, v := range mappingAutoScalePolicyPolicyType {
		values = append(values, v)
	}
	return values
}
