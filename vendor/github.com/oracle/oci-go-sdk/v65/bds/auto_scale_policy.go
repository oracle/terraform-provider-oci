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

// AutoScalePolicy This model for autoscaling policy is deprecated and not supported for ODH clusters. Use the `AutoScalePolicyDetails` model to manage autoscale policy details for ODH clusters.
type AutoScalePolicy struct {

	// Types of autoscale policies. Options are SCHEDULE-BASED or THRESHOLD-BASED. (Only THRESHOLD-BASED is supported in this release.)
	PolicyType AutoScalePolicyPolicyTypeEnum `mandatory:"true" json:"policyType"`

	// The list of rules for autoscaling. If an action has multiple rules, the last rule in the array will be applied.
	Rules []AutoScalePolicyRule `mandatory:"true" json:"rules"`
}

func (m AutoScalePolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoScalePolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutoScalePolicyPolicyTypeEnum(string(m.PolicyType)); !ok && m.PolicyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PolicyType: %s. Supported values are: %s.", m.PolicyType, strings.Join(GetAutoScalePolicyPolicyTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutoScalePolicyPolicyTypeEnum Enum with underlying type: string
type AutoScalePolicyPolicyTypeEnum string

// Set of constants representing the allowable values for AutoScalePolicyPolicyTypeEnum
const (
	AutoScalePolicyPolicyTypeThresholdBased AutoScalePolicyPolicyTypeEnum = "THRESHOLD_BASED"
	AutoScalePolicyPolicyTypeScheduleBased  AutoScalePolicyPolicyTypeEnum = "SCHEDULE_BASED"
	AutoScalePolicyPolicyTypeNone           AutoScalePolicyPolicyTypeEnum = "NONE"
)

var mappingAutoScalePolicyPolicyTypeEnum = map[string]AutoScalePolicyPolicyTypeEnum{
	"THRESHOLD_BASED": AutoScalePolicyPolicyTypeThresholdBased,
	"SCHEDULE_BASED":  AutoScalePolicyPolicyTypeScheduleBased,
	"NONE":            AutoScalePolicyPolicyTypeNone,
}

var mappingAutoScalePolicyPolicyTypeEnumLowerCase = map[string]AutoScalePolicyPolicyTypeEnum{
	"threshold_based": AutoScalePolicyPolicyTypeThresholdBased,
	"schedule_based":  AutoScalePolicyPolicyTypeScheduleBased,
	"none":            AutoScalePolicyPolicyTypeNone,
}

// GetAutoScalePolicyPolicyTypeEnumValues Enumerates the set of values for AutoScalePolicyPolicyTypeEnum
func GetAutoScalePolicyPolicyTypeEnumValues() []AutoScalePolicyPolicyTypeEnum {
	values := make([]AutoScalePolicyPolicyTypeEnum, 0)
	for _, v := range mappingAutoScalePolicyPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoScalePolicyPolicyTypeEnumStringValues Enumerates the set of values in String for AutoScalePolicyPolicyTypeEnum
func GetAutoScalePolicyPolicyTypeEnumStringValues() []string {
	return []string{
		"THRESHOLD_BASED",
		"SCHEDULE_BASED",
		"NONE",
	}
}

// GetMappingAutoScalePolicyPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoScalePolicyPolicyTypeEnum(val string) (AutoScalePolicyPolicyTypeEnum, bool) {
	enum, ok := mappingAutoScalePolicyPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
