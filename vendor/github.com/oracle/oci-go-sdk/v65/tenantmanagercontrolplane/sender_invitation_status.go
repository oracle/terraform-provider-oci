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

// SenderInvitationStatusEnum Enum with underlying type: string
type SenderInvitationStatusEnum string

// Set of constants representing the allowable values for SenderInvitationStatusEnum
const (
	SenderInvitationStatusPending  SenderInvitationStatusEnum = "PENDING"
	SenderInvitationStatusCanceled SenderInvitationStatusEnum = "CANCELED"
	SenderInvitationStatusAccepted SenderInvitationStatusEnum = "ACCEPTED"
	SenderInvitationStatusExpired  SenderInvitationStatusEnum = "EXPIRED"
	SenderInvitationStatusFailed   SenderInvitationStatusEnum = "FAILED"
)

var mappingSenderInvitationStatusEnum = map[string]SenderInvitationStatusEnum{
	"PENDING":  SenderInvitationStatusPending,
	"CANCELED": SenderInvitationStatusCanceled,
	"ACCEPTED": SenderInvitationStatusAccepted,
	"EXPIRED":  SenderInvitationStatusExpired,
	"FAILED":   SenderInvitationStatusFailed,
}

var mappingSenderInvitationStatusEnumLowerCase = map[string]SenderInvitationStatusEnum{
	"pending":  SenderInvitationStatusPending,
	"canceled": SenderInvitationStatusCanceled,
	"accepted": SenderInvitationStatusAccepted,
	"expired":  SenderInvitationStatusExpired,
	"failed":   SenderInvitationStatusFailed,
}

// GetSenderInvitationStatusEnumValues Enumerates the set of values for SenderInvitationStatusEnum
func GetSenderInvitationStatusEnumValues() []SenderInvitationStatusEnum {
	values := make([]SenderInvitationStatusEnum, 0)
	for _, v := range mappingSenderInvitationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSenderInvitationStatusEnumStringValues Enumerates the set of values in String for SenderInvitationStatusEnum
func GetSenderInvitationStatusEnumStringValues() []string {
	return []string{
		"PENDING",
		"CANCELED",
		"ACCEPTED",
		"EXPIRED",
		"FAILED",
	}
}

// GetMappingSenderInvitationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSenderInvitationStatusEnum(val string) (SenderInvitationStatusEnum, bool) {
	enum, ok := mappingSenderInvitationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
