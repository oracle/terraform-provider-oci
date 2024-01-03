// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BackoffRetryPolicy The backoff retry portion of the subscription delivery policy. For information about retry durations for subscriptions, see
// How Notifications Works (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm#how).
type BackoffRetryPolicy struct {

	// The maximum retry duration in milliseconds. Default value is `7200000` (2 hours).
	MaxRetryDuration *int `mandatory:"true" json:"maxRetryDuration"`

	// The type of delivery policy.
	PolicyType BackoffRetryPolicyPolicyTypeEnum `mandatory:"true" json:"policyType"`
}

func (m BackoffRetryPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackoffRetryPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBackoffRetryPolicyPolicyTypeEnum(string(m.PolicyType)); !ok && m.PolicyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PolicyType: %s. Supported values are: %s.", m.PolicyType, strings.Join(GetBackoffRetryPolicyPolicyTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackoffRetryPolicyPolicyTypeEnum Enum with underlying type: string
type BackoffRetryPolicyPolicyTypeEnum string

// Set of constants representing the allowable values for BackoffRetryPolicyPolicyTypeEnum
const (
	BackoffRetryPolicyPolicyTypeExponential BackoffRetryPolicyPolicyTypeEnum = "EXPONENTIAL"
)

var mappingBackoffRetryPolicyPolicyTypeEnum = map[string]BackoffRetryPolicyPolicyTypeEnum{
	"EXPONENTIAL": BackoffRetryPolicyPolicyTypeExponential,
}

var mappingBackoffRetryPolicyPolicyTypeEnumLowerCase = map[string]BackoffRetryPolicyPolicyTypeEnum{
	"exponential": BackoffRetryPolicyPolicyTypeExponential,
}

// GetBackoffRetryPolicyPolicyTypeEnumValues Enumerates the set of values for BackoffRetryPolicyPolicyTypeEnum
func GetBackoffRetryPolicyPolicyTypeEnumValues() []BackoffRetryPolicyPolicyTypeEnum {
	values := make([]BackoffRetryPolicyPolicyTypeEnum, 0)
	for _, v := range mappingBackoffRetryPolicyPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackoffRetryPolicyPolicyTypeEnumStringValues Enumerates the set of values in String for BackoffRetryPolicyPolicyTypeEnum
func GetBackoffRetryPolicyPolicyTypeEnumStringValues() []string {
	return []string{
		"EXPONENTIAL",
	}
}

// GetMappingBackoffRetryPolicyPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackoffRetryPolicyPolicyTypeEnum(val string) (BackoffRetryPolicyPolicyTypeEnum, bool) {
	enum, ok := mappingBackoffRetryPolicyPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
