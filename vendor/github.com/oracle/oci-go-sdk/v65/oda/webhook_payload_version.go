// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"strings"
)

// WebhookPayloadVersionEnum Enum with underlying type: string
type WebhookPayloadVersionEnum string

// Set of constants representing the allowable values for WebhookPayloadVersionEnum
const (
	WebhookPayloadVersion10 WebhookPayloadVersionEnum = "1.0"
	WebhookPayloadVersion11 WebhookPayloadVersionEnum = "1.1"
)

var mappingWebhookPayloadVersionEnum = map[string]WebhookPayloadVersionEnum{
	"1.0": WebhookPayloadVersion10,
	"1.1": WebhookPayloadVersion11,
}

var mappingWebhookPayloadVersionEnumLowerCase = map[string]WebhookPayloadVersionEnum{
	"1.0": WebhookPayloadVersion10,
	"1.1": WebhookPayloadVersion11,
}

// GetWebhookPayloadVersionEnumValues Enumerates the set of values for WebhookPayloadVersionEnum
func GetWebhookPayloadVersionEnumValues() []WebhookPayloadVersionEnum {
	values := make([]WebhookPayloadVersionEnum, 0)
	for _, v := range mappingWebhookPayloadVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetWebhookPayloadVersionEnumStringValues Enumerates the set of values in String for WebhookPayloadVersionEnum
func GetWebhookPayloadVersionEnumStringValues() []string {
	return []string{
		"1.0",
		"1.1",
	}
}

// GetMappingWebhookPayloadVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWebhookPayloadVersionEnum(val string) (WebhookPayloadVersionEnum, bool) {
	enum, ok := mappingWebhookPayloadVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
