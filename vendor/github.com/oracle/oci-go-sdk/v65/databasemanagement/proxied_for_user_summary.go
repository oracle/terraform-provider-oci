// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProxiedForUserSummary A summary of users on whose behalf the current user acts as proxy.
type ProxiedForUserSummary struct {

	// The name of a proxy user or the name of the client user.
	Name *string `mandatory:"false" json:"name"`

	// Indicates whether the proxy is required to supply the client credentials (YES) or not (NO).
	Authentication ProxiedForUserSummaryAuthenticationEnum `mandatory:"false" json:"authentication,omitempty"`

	// The flags associated with the proxy/client pair.
	Flags ProxiedForUserSummaryFlagsEnum `mandatory:"false" json:"flags,omitempty"`
}

func (m ProxiedForUserSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProxiedForUserSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingProxiedForUserSummaryAuthenticationEnum(string(m.Authentication)); !ok && m.Authentication != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Authentication: %s. Supported values are: %s.", m.Authentication, strings.Join(GetProxiedForUserSummaryAuthenticationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingProxiedForUserSummaryFlagsEnum(string(m.Flags)); !ok && m.Flags != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Flags: %s. Supported values are: %s.", m.Flags, strings.Join(GetProxiedForUserSummaryFlagsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProxiedForUserSummaryAuthenticationEnum Enum with underlying type: string
type ProxiedForUserSummaryAuthenticationEnum string

// Set of constants representing the allowable values for ProxiedForUserSummaryAuthenticationEnum
const (
	ProxiedForUserSummaryAuthenticationYes ProxiedForUserSummaryAuthenticationEnum = "YES"
	ProxiedForUserSummaryAuthenticationNo  ProxiedForUserSummaryAuthenticationEnum = "NO"
)

var mappingProxiedForUserSummaryAuthenticationEnum = map[string]ProxiedForUserSummaryAuthenticationEnum{
	"YES": ProxiedForUserSummaryAuthenticationYes,
	"NO":  ProxiedForUserSummaryAuthenticationNo,
}

var mappingProxiedForUserSummaryAuthenticationEnumLowerCase = map[string]ProxiedForUserSummaryAuthenticationEnum{
	"yes": ProxiedForUserSummaryAuthenticationYes,
	"no":  ProxiedForUserSummaryAuthenticationNo,
}

// GetProxiedForUserSummaryAuthenticationEnumValues Enumerates the set of values for ProxiedForUserSummaryAuthenticationEnum
func GetProxiedForUserSummaryAuthenticationEnumValues() []ProxiedForUserSummaryAuthenticationEnum {
	values := make([]ProxiedForUserSummaryAuthenticationEnum, 0)
	for _, v := range mappingProxiedForUserSummaryAuthenticationEnum {
		values = append(values, v)
	}
	return values
}

// GetProxiedForUserSummaryAuthenticationEnumStringValues Enumerates the set of values in String for ProxiedForUserSummaryAuthenticationEnum
func GetProxiedForUserSummaryAuthenticationEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingProxiedForUserSummaryAuthenticationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProxiedForUserSummaryAuthenticationEnum(val string) (ProxiedForUserSummaryAuthenticationEnum, bool) {
	enum, ok := mappingProxiedForUserSummaryAuthenticationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ProxiedForUserSummaryFlagsEnum Enum with underlying type: string
type ProxiedForUserSummaryFlagsEnum string

// Set of constants representing the allowable values for ProxiedForUserSummaryFlagsEnum
const (
	ProxiedForUserSummaryFlagsProxyMayActivateAllClientRoles ProxiedForUserSummaryFlagsEnum = "PROXY_MAY_ACTIVATE_ALL_CLIENT_ROLES"
	ProxiedForUserSummaryFlagsNoClientRolesMayBeActivated    ProxiedForUserSummaryFlagsEnum = "NO_CLIENT_ROLES_MAY_BE_ACTIVATED"
	ProxiedForUserSummaryFlagsProxyMayActivateRole           ProxiedForUserSummaryFlagsEnum = "PROXY_MAY_ACTIVATE_ROLE"
	ProxiedForUserSummaryFlagsProxyMayNotActivateRole        ProxiedForUserSummaryFlagsEnum = "PROXY_MAY_NOT_ACTIVATE_ROLE"
)

var mappingProxiedForUserSummaryFlagsEnum = map[string]ProxiedForUserSummaryFlagsEnum{
	"PROXY_MAY_ACTIVATE_ALL_CLIENT_ROLES": ProxiedForUserSummaryFlagsProxyMayActivateAllClientRoles,
	"NO_CLIENT_ROLES_MAY_BE_ACTIVATED":    ProxiedForUserSummaryFlagsNoClientRolesMayBeActivated,
	"PROXY_MAY_ACTIVATE_ROLE":             ProxiedForUserSummaryFlagsProxyMayActivateRole,
	"PROXY_MAY_NOT_ACTIVATE_ROLE":         ProxiedForUserSummaryFlagsProxyMayNotActivateRole,
}

var mappingProxiedForUserSummaryFlagsEnumLowerCase = map[string]ProxiedForUserSummaryFlagsEnum{
	"proxy_may_activate_all_client_roles": ProxiedForUserSummaryFlagsProxyMayActivateAllClientRoles,
	"no_client_roles_may_be_activated":    ProxiedForUserSummaryFlagsNoClientRolesMayBeActivated,
	"proxy_may_activate_role":             ProxiedForUserSummaryFlagsProxyMayActivateRole,
	"proxy_may_not_activate_role":         ProxiedForUserSummaryFlagsProxyMayNotActivateRole,
}

// GetProxiedForUserSummaryFlagsEnumValues Enumerates the set of values for ProxiedForUserSummaryFlagsEnum
func GetProxiedForUserSummaryFlagsEnumValues() []ProxiedForUserSummaryFlagsEnum {
	values := make([]ProxiedForUserSummaryFlagsEnum, 0)
	for _, v := range mappingProxiedForUserSummaryFlagsEnum {
		values = append(values, v)
	}
	return values
}

// GetProxiedForUserSummaryFlagsEnumStringValues Enumerates the set of values in String for ProxiedForUserSummaryFlagsEnum
func GetProxiedForUserSummaryFlagsEnumStringValues() []string {
	return []string{
		"PROXY_MAY_ACTIVATE_ALL_CLIENT_ROLES",
		"NO_CLIENT_ROLES_MAY_BE_ACTIVATED",
		"PROXY_MAY_ACTIVATE_ROLE",
		"PROXY_MAY_NOT_ACTIVATE_ROLE",
	}
}

// GetMappingProxiedForUserSummaryFlagsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProxiedForUserSummaryFlagsEnum(val string) (ProxiedForUserSummaryFlagsEnum, bool) {
	enum, ok := mappingProxiedForUserSummaryFlagsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
