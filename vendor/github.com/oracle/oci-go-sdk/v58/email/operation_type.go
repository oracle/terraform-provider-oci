// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// API for the Email Delivery service. Use this API to send high-volume, application-generated
// emails. For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//
// **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
// If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateDkim        OperationTypeEnum = "CREATE_DKIM"
	OperationTypeDeleteDkim        OperationTypeEnum = "DELETE_DKIM"
	OperationTypeMoveDkim          OperationTypeEnum = "MOVE_DKIM"
	OperationTypeUpdateDkim        OperationTypeEnum = "UPDATE_DKIM"
	OperationTypeCreateEmailDomain OperationTypeEnum = "CREATE_EMAIL_DOMAIN"
	OperationTypeDeleteEmailDomain OperationTypeEnum = "DELETE_EMAIL_DOMAIN"
	OperationTypeMoveEmailDomain   OperationTypeEnum = "MOVE_EMAIL_DOMAIN"
	OperationTypeUpdateEmailDomain OperationTypeEnum = "UPDATE_EMAIL_DOMAIN"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_DKIM":         OperationTypeCreateDkim,
	"DELETE_DKIM":         OperationTypeDeleteDkim,
	"MOVE_DKIM":           OperationTypeMoveDkim,
	"UPDATE_DKIM":         OperationTypeUpdateDkim,
	"CREATE_EMAIL_DOMAIN": OperationTypeCreateEmailDomain,
	"DELETE_EMAIL_DOMAIN": OperationTypeDeleteEmailDomain,
	"MOVE_EMAIL_DOMAIN":   OperationTypeMoveEmailDomain,
	"UPDATE_EMAIL_DOMAIN": OperationTypeUpdateEmailDomain,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_DKIM",
		"DELETE_DKIM",
		"MOVE_DKIM",
		"UPDATE_DKIM",
		"CREATE_EMAIL_DOMAIN",
		"DELETE_EMAIL_DOMAIN",
		"MOVE_EMAIL_DOMAIN",
		"UPDATE_EMAIL_DOMAIN",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	mappingOperationTypeEnumIgnoreCase := make(map[string]OperationTypeEnum)
	for k, v := range mappingOperationTypeEnum {
		mappingOperationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOperationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
