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

// InvitationSubjectEnum Enum with underlying type: string
type InvitationSubjectEnum string

// Set of constants representing the allowable values for InvitationSubjectEnum
const (
	InvitationSubjectLink       InvitationSubjectEnum = "LINK"
	InvitationSubjectGovernance InvitationSubjectEnum = "GOVERNANCE"
)

var mappingInvitationSubjectEnum = map[string]InvitationSubjectEnum{
	"LINK":       InvitationSubjectLink,
	"GOVERNANCE": InvitationSubjectGovernance,
}

var mappingInvitationSubjectEnumLowerCase = map[string]InvitationSubjectEnum{
	"link":       InvitationSubjectLink,
	"governance": InvitationSubjectGovernance,
}

// GetInvitationSubjectEnumValues Enumerates the set of values for InvitationSubjectEnum
func GetInvitationSubjectEnumValues() []InvitationSubjectEnum {
	values := make([]InvitationSubjectEnum, 0)
	for _, v := range mappingInvitationSubjectEnum {
		values = append(values, v)
	}
	return values
}

// GetInvitationSubjectEnumStringValues Enumerates the set of values in String for InvitationSubjectEnum
func GetInvitationSubjectEnumStringValues() []string {
	return []string{
		"LINK",
		"GOVERNANCE",
	}
}

// GetMappingInvitationSubjectEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInvitationSubjectEnum(val string) (InvitationSubjectEnum, bool) {
	enum, ok := mappingInvitationSubjectEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
