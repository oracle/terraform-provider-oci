// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"strings"
)

// RecipientInvitationStatusEnum Enum with underlying type: string
type RecipientInvitationStatusEnum string

// Set of constants representing the allowable values for RecipientInvitationStatusEnum
const (
	RecipientInvitationStatusPending  RecipientInvitationStatusEnum = "PENDING"
	RecipientInvitationStatusCanceled RecipientInvitationStatusEnum = "CANCELED"
	RecipientInvitationStatusAccepted RecipientInvitationStatusEnum = "ACCEPTED"
	RecipientInvitationStatusIgnored  RecipientInvitationStatusEnum = "IGNORED"
	RecipientInvitationStatusExpired  RecipientInvitationStatusEnum = "EXPIRED"
	RecipientInvitationStatusFailed   RecipientInvitationStatusEnum = "FAILED"
)

var mappingRecipientInvitationStatusEnum = map[string]RecipientInvitationStatusEnum{
	"PENDING":  RecipientInvitationStatusPending,
	"CANCELED": RecipientInvitationStatusCanceled,
	"ACCEPTED": RecipientInvitationStatusAccepted,
	"IGNORED":  RecipientInvitationStatusIgnored,
	"EXPIRED":  RecipientInvitationStatusExpired,
	"FAILED":   RecipientInvitationStatusFailed,
}

var mappingRecipientInvitationStatusEnumLowerCase = map[string]RecipientInvitationStatusEnum{
	"pending":  RecipientInvitationStatusPending,
	"canceled": RecipientInvitationStatusCanceled,
	"accepted": RecipientInvitationStatusAccepted,
	"ignored":  RecipientInvitationStatusIgnored,
	"expired":  RecipientInvitationStatusExpired,
	"failed":   RecipientInvitationStatusFailed,
}

// GetRecipientInvitationStatusEnumValues Enumerates the set of values for RecipientInvitationStatusEnum
func GetRecipientInvitationStatusEnumValues() []RecipientInvitationStatusEnum {
	values := make([]RecipientInvitationStatusEnum, 0)
	for _, v := range mappingRecipientInvitationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRecipientInvitationStatusEnumStringValues Enumerates the set of values in String for RecipientInvitationStatusEnum
func GetRecipientInvitationStatusEnumStringValues() []string {
	return []string{
		"PENDING",
		"CANCELED",
		"ACCEPTED",
		"IGNORED",
		"EXPIRED",
		"FAILED",
	}
}

// GetMappingRecipientInvitationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecipientInvitationStatusEnum(val string) (RecipientInvitationStatusEnum, bool) {
	enum, ok := mappingRecipientInvitationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
