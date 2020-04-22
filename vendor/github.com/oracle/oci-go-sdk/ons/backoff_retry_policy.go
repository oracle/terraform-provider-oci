// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"github.com/oracle/oci-go-sdk/common"
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

// BackoffRetryPolicyPolicyTypeEnum Enum with underlying type: string
type BackoffRetryPolicyPolicyTypeEnum string

// Set of constants representing the allowable values for BackoffRetryPolicyPolicyTypeEnum
const (
	BackoffRetryPolicyPolicyTypeExponential BackoffRetryPolicyPolicyTypeEnum = "EXPONENTIAL"
)

var mappingBackoffRetryPolicyPolicyType = map[string]BackoffRetryPolicyPolicyTypeEnum{
	"EXPONENTIAL": BackoffRetryPolicyPolicyTypeExponential,
}

// GetBackoffRetryPolicyPolicyTypeEnumValues Enumerates the set of values for BackoffRetryPolicyPolicyTypeEnum
func GetBackoffRetryPolicyPolicyTypeEnumValues() []BackoffRetryPolicyPolicyTypeEnum {
	values := make([]BackoffRetryPolicyPolicyTypeEnum, 0)
	for _, v := range mappingBackoffRetryPolicyPolicyType {
		values = append(values, v)
	}
	return values
}
