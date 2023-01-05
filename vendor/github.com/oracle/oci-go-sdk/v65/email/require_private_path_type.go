// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to send high-volume and application-generated emails.
// For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"strings"
)

// RequirePrivatePathTypeEnum Enum with underlying type: string
type RequirePrivatePathTypeEnum string

// Set of constants representing the allowable values for RequirePrivatePathTypeEnum
const (
	RequirePrivatePathTypeNone    RequirePrivatePathTypeEnum = "NONE"
	RequirePrivatePathTypeSend    RequirePrivatePathTypeEnum = "SEND"
	RequirePrivatePathTypeReceive RequirePrivatePathTypeEnum = "RECEIVE"
	RequirePrivatePathTypeBoth    RequirePrivatePathTypeEnum = "BOTH"
)

var mappingRequirePrivatePathTypeEnum = map[string]RequirePrivatePathTypeEnum{
	"NONE":    RequirePrivatePathTypeNone,
	"SEND":    RequirePrivatePathTypeSend,
	"RECEIVE": RequirePrivatePathTypeReceive,
	"BOTH":    RequirePrivatePathTypeBoth,
}

var mappingRequirePrivatePathTypeEnumLowerCase = map[string]RequirePrivatePathTypeEnum{
	"none":    RequirePrivatePathTypeNone,
	"send":    RequirePrivatePathTypeSend,
	"receive": RequirePrivatePathTypeReceive,
	"both":    RequirePrivatePathTypeBoth,
}

// GetRequirePrivatePathTypeEnumValues Enumerates the set of values for RequirePrivatePathTypeEnum
func GetRequirePrivatePathTypeEnumValues() []RequirePrivatePathTypeEnum {
	values := make([]RequirePrivatePathTypeEnum, 0)
	for _, v := range mappingRequirePrivatePathTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRequirePrivatePathTypeEnumStringValues Enumerates the set of values in String for RequirePrivatePathTypeEnum
func GetRequirePrivatePathTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"SEND",
		"RECEIVE",
		"BOTH",
	}
}

// GetMappingRequirePrivatePathTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequirePrivatePathTypeEnum(val string) (RequirePrivatePathTypeEnum, bool) {
	enum, ok := mappingRequirePrivatePathTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
