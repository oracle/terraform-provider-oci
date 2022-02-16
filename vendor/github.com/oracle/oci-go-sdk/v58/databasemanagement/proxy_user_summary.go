// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ProxyUserSummary A summary of the proxy user.
type ProxyUserSummary struct {

	// The name of a proxy user or the name of the client user.
	Name *string `mandatory:"false" json:"name"`

	// Indicates whether the proxy is required to supply the client credentials (YES) or not (NO).
	Authentication ProxyUserSummaryAuthenticationEnum `mandatory:"false" json:"authentication,omitempty"`

	// The flags associated with the proxy/client pair.
	Flags ProxyUserSummaryFlagsEnum `mandatory:"false" json:"flags,omitempty"`
}

func (m ProxyUserSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProxyUserSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingProxyUserSummaryAuthenticationEnum(string(m.Authentication)); !ok && m.Authentication != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Authentication: %s. Supported values are: %s.", m.Authentication, strings.Join(GetProxyUserSummaryAuthenticationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingProxyUserSummaryFlagsEnum(string(m.Flags)); !ok && m.Flags != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Flags: %s. Supported values are: %s.", m.Flags, strings.Join(GetProxyUserSummaryFlagsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProxyUserSummaryAuthenticationEnum Enum with underlying type: string
type ProxyUserSummaryAuthenticationEnum string

// Set of constants representing the allowable values for ProxyUserSummaryAuthenticationEnum
const (
	ProxyUserSummaryAuthenticationYes ProxyUserSummaryAuthenticationEnum = "YES"
	ProxyUserSummaryAuthenticationNo  ProxyUserSummaryAuthenticationEnum = "NO"
)

var mappingProxyUserSummaryAuthenticationEnum = map[string]ProxyUserSummaryAuthenticationEnum{
	"YES": ProxyUserSummaryAuthenticationYes,
	"NO":  ProxyUserSummaryAuthenticationNo,
}

// GetProxyUserSummaryAuthenticationEnumValues Enumerates the set of values for ProxyUserSummaryAuthenticationEnum
func GetProxyUserSummaryAuthenticationEnumValues() []ProxyUserSummaryAuthenticationEnum {
	values := make([]ProxyUserSummaryAuthenticationEnum, 0)
	for _, v := range mappingProxyUserSummaryAuthenticationEnum {
		values = append(values, v)
	}
	return values
}

// GetProxyUserSummaryAuthenticationEnumStringValues Enumerates the set of values in String for ProxyUserSummaryAuthenticationEnum
func GetProxyUserSummaryAuthenticationEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingProxyUserSummaryAuthenticationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProxyUserSummaryAuthenticationEnum(val string) (ProxyUserSummaryAuthenticationEnum, bool) {
	mappingProxyUserSummaryAuthenticationEnumIgnoreCase := make(map[string]ProxyUserSummaryAuthenticationEnum)
	for k, v := range mappingProxyUserSummaryAuthenticationEnum {
		mappingProxyUserSummaryAuthenticationEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingProxyUserSummaryAuthenticationEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ProxyUserSummaryFlagsEnum Enum with underlying type: string
type ProxyUserSummaryFlagsEnum string

// Set of constants representing the allowable values for ProxyUserSummaryFlagsEnum
const (
	ProxyUserSummaryFlagsProxyMayActivateAllClientRoles ProxyUserSummaryFlagsEnum = "PROXY_MAY_ACTIVATE_ALL_CLIENT_ROLES"
	ProxyUserSummaryFlagsNoClientRolesMayBeActivated    ProxyUserSummaryFlagsEnum = "NO_CLIENT_ROLES_MAY_BE_ACTIVATED"
	ProxyUserSummaryFlagsProxyMayActivateRole           ProxyUserSummaryFlagsEnum = "PROXY_MAY_ACTIVATE_ROLE"
	ProxyUserSummaryFlagsProxyMayNotActivateRole        ProxyUserSummaryFlagsEnum = "PROXY_MAY_NOT_ACTIVATE_ROLE"
)

var mappingProxyUserSummaryFlagsEnum = map[string]ProxyUserSummaryFlagsEnum{
	"PROXY_MAY_ACTIVATE_ALL_CLIENT_ROLES": ProxyUserSummaryFlagsProxyMayActivateAllClientRoles,
	"NO_CLIENT_ROLES_MAY_BE_ACTIVATED":    ProxyUserSummaryFlagsNoClientRolesMayBeActivated,
	"PROXY_MAY_ACTIVATE_ROLE":             ProxyUserSummaryFlagsProxyMayActivateRole,
	"PROXY_MAY_NOT_ACTIVATE_ROLE":         ProxyUserSummaryFlagsProxyMayNotActivateRole,
}

// GetProxyUserSummaryFlagsEnumValues Enumerates the set of values for ProxyUserSummaryFlagsEnum
func GetProxyUserSummaryFlagsEnumValues() []ProxyUserSummaryFlagsEnum {
	values := make([]ProxyUserSummaryFlagsEnum, 0)
	for _, v := range mappingProxyUserSummaryFlagsEnum {
		values = append(values, v)
	}
	return values
}

// GetProxyUserSummaryFlagsEnumStringValues Enumerates the set of values in String for ProxyUserSummaryFlagsEnum
func GetProxyUserSummaryFlagsEnumStringValues() []string {
	return []string{
		"PROXY_MAY_ACTIVATE_ALL_CLIENT_ROLES",
		"NO_CLIENT_ROLES_MAY_BE_ACTIVATED",
		"PROXY_MAY_ACTIVATE_ROLE",
		"PROXY_MAY_NOT_ACTIVATE_ROLE",
	}
}

// GetMappingProxyUserSummaryFlagsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProxyUserSummaryFlagsEnum(val string) (ProxyUserSummaryFlagsEnum, bool) {
	mappingProxyUserSummaryFlagsEnumIgnoreCase := make(map[string]ProxyUserSummaryFlagsEnum)
	for k, v := range mappingProxyUserSummaryFlagsEnum {
		mappingProxyUserSummaryFlagsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingProxyUserSummaryFlagsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
