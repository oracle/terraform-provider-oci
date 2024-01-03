// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"strings"
)

// SuppressionReasonEnum Enum with underlying type: string
type SuppressionReasonEnum string

// Set of constants representing the allowable values for SuppressionReasonEnum
const (
	SuppressionReasonUnknown     SuppressionReasonEnum = "UNKNOWN"
	SuppressionReasonHardbounce  SuppressionReasonEnum = "HARDBOUNCE"
	SuppressionReasonComplaint   SuppressionReasonEnum = "COMPLAINT"
	SuppressionReasonManual      SuppressionReasonEnum = "MANUAL"
	SuppressionReasonSoftbounce  SuppressionReasonEnum = "SOFTBOUNCE"
	SuppressionReasonUnsubscribe SuppressionReasonEnum = "UNSUBSCRIBE"
)

var mappingSuppressionReasonEnum = map[string]SuppressionReasonEnum{
	"UNKNOWN":     SuppressionReasonUnknown,
	"HARDBOUNCE":  SuppressionReasonHardbounce,
	"COMPLAINT":   SuppressionReasonComplaint,
	"MANUAL":      SuppressionReasonManual,
	"SOFTBOUNCE":  SuppressionReasonSoftbounce,
	"UNSUBSCRIBE": SuppressionReasonUnsubscribe,
}

var mappingSuppressionReasonEnumLowerCase = map[string]SuppressionReasonEnum{
	"unknown":     SuppressionReasonUnknown,
	"hardbounce":  SuppressionReasonHardbounce,
	"complaint":   SuppressionReasonComplaint,
	"manual":      SuppressionReasonManual,
	"softbounce":  SuppressionReasonSoftbounce,
	"unsubscribe": SuppressionReasonUnsubscribe,
}

// GetSuppressionReasonEnumValues Enumerates the set of values for SuppressionReasonEnum
func GetSuppressionReasonEnumValues() []SuppressionReasonEnum {
	values := make([]SuppressionReasonEnum, 0)
	for _, v := range mappingSuppressionReasonEnum {
		values = append(values, v)
	}
	return values
}

// GetSuppressionReasonEnumStringValues Enumerates the set of values in String for SuppressionReasonEnum
func GetSuppressionReasonEnumStringValues() []string {
	return []string{
		"UNKNOWN",
		"HARDBOUNCE",
		"COMPLAINT",
		"MANUAL",
		"SOFTBOUNCE",
		"UNSUBSCRIBE",
	}
}

// GetMappingSuppressionReasonEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSuppressionReasonEnum(val string) (SuppressionReasonEnum, bool) {
	enum, ok := mappingSuppressionReasonEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
