// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"strings"
)

// ActionTypeEnum Enum with underlying type: string
type ActionTypeEnum string

// Set of constants representing the allowable values for ActionTypeEnum
const (
	ActionTypeCreated        ActionTypeEnum = "CREATED"
	ActionTypeUpdated        ActionTypeEnum = "UPDATED"
	ActionTypeDeleted        ActionTypeEnum = "DELETED"
	ActionTypeInProgress     ActionTypeEnum = "IN_PROGRESS"
	ActionTypeFailed         ActionTypeEnum = "FAILED"
	ActionTypeRelated        ActionTypeEnum = "RELATED"
	ActionTypeCanceledCreate ActionTypeEnum = "CANCELED_CREATE"
	ActionTypeCanceledUpdate ActionTypeEnum = "CANCELED_UPDATE"
	ActionTypeCanceledDelete ActionTypeEnum = "CANCELED_DELETE"
)

var mappingActionTypeEnum = map[string]ActionTypeEnum{
	"CREATED":         ActionTypeCreated,
	"UPDATED":         ActionTypeUpdated,
	"DELETED":         ActionTypeDeleted,
	"IN_PROGRESS":     ActionTypeInProgress,
	"FAILED":          ActionTypeFailed,
	"RELATED":         ActionTypeRelated,
	"CANCELED_CREATE": ActionTypeCanceledCreate,
	"CANCELED_UPDATE": ActionTypeCanceledUpdate,
	"CANCELED_DELETE": ActionTypeCanceledDelete,
}

var mappingActionTypeEnumLowerCase = map[string]ActionTypeEnum{
	"created":         ActionTypeCreated,
	"updated":         ActionTypeUpdated,
	"deleted":         ActionTypeDeleted,
	"in_progress":     ActionTypeInProgress,
	"failed":          ActionTypeFailed,
	"related":         ActionTypeRelated,
	"canceled_create": ActionTypeCanceledCreate,
	"canceled_update": ActionTypeCanceledUpdate,
	"canceled_delete": ActionTypeCanceledDelete,
}

// GetActionTypeEnumValues Enumerates the set of values for ActionTypeEnum
func GetActionTypeEnumValues() []ActionTypeEnum {
	values := make([]ActionTypeEnum, 0)
	for _, v := range mappingActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetActionTypeEnumStringValues Enumerates the set of values in String for ActionTypeEnum
func GetActionTypeEnumStringValues() []string {
	return []string{
		"CREATED",
		"UPDATED",
		"DELETED",
		"IN_PROGRESS",
		"FAILED",
		"RELATED",
		"CANCELED_CREATE",
		"CANCELED_UPDATE",
		"CANCELED_DELETE",
	}
}

// GetMappingActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionTypeEnum(val string) (ActionTypeEnum, bool) {
	enum, ok := mappingActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
