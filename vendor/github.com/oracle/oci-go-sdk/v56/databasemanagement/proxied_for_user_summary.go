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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ProxiedForUserSummary A collection of Users for which the current user acts as proxy.
type ProxiedForUserSummary struct {

	// The name of a proxy user or name of the user who the proxy user can act as
	Name *string `mandatory:"false" json:"name"`

	// Indicates whether the proxy is required to supply the client credentials (YES) or not (NO)
	Authentication ProxiedForUserSummaryAuthenticationEnum `mandatory:"false" json:"authentication,omitempty"`

	// Flags associated with the proxy/client pair
	Flags ProxiedForUserSummaryFlagsEnum `mandatory:"false" json:"flags,omitempty"`
}

func (m ProxiedForUserSummary) String() string {
	return common.PointerString(m)
}

// ProxiedForUserSummaryAuthenticationEnum Enum with underlying type: string
type ProxiedForUserSummaryAuthenticationEnum string

// Set of constants representing the allowable values for ProxiedForUserSummaryAuthenticationEnum
const (
	ProxiedForUserSummaryAuthenticationYes ProxiedForUserSummaryAuthenticationEnum = "YES"
	ProxiedForUserSummaryAuthenticationNo  ProxiedForUserSummaryAuthenticationEnum = "NO"
)

var mappingProxiedForUserSummaryAuthentication = map[string]ProxiedForUserSummaryAuthenticationEnum{
	"YES": ProxiedForUserSummaryAuthenticationYes,
	"NO":  ProxiedForUserSummaryAuthenticationNo,
}

// GetProxiedForUserSummaryAuthenticationEnumValues Enumerates the set of values for ProxiedForUserSummaryAuthenticationEnum
func GetProxiedForUserSummaryAuthenticationEnumValues() []ProxiedForUserSummaryAuthenticationEnum {
	values := make([]ProxiedForUserSummaryAuthenticationEnum, 0)
	for _, v := range mappingProxiedForUserSummaryAuthentication {
		values = append(values, v)
	}
	return values
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

var mappingProxiedForUserSummaryFlags = map[string]ProxiedForUserSummaryFlagsEnum{
	"PROXY_MAY_ACTIVATE_ALL_CLIENT_ROLES": ProxiedForUserSummaryFlagsProxyMayActivateAllClientRoles,
	"NO_CLIENT_ROLES_MAY_BE_ACTIVATED":    ProxiedForUserSummaryFlagsNoClientRolesMayBeActivated,
	"PROXY_MAY_ACTIVATE_ROLE":             ProxiedForUserSummaryFlagsProxyMayActivateRole,
	"PROXY_MAY_NOT_ACTIVATE_ROLE":         ProxiedForUserSummaryFlagsProxyMayNotActivateRole,
}

// GetProxiedForUserSummaryFlagsEnumValues Enumerates the set of values for ProxiedForUserSummaryFlagsEnum
func GetProxiedForUserSummaryFlagsEnumValues() []ProxiedForUserSummaryFlagsEnum {
	values := make([]ProxiedForUserSummaryFlagsEnum, 0)
	for _, v := range mappingProxiedForUserSummaryFlags {
		values = append(values, v)
	}
	return values
}
