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

// ProxyUserSummary Summary of proxy user
type ProxyUserSummary struct {

	// The name of a proxy user or name of the user who the proxy user can act as
	Name *string `mandatory:"false" json:"name"`

	// Indicates whether the proxy is required to supply the client credentials (YES) or not (NO)
	Authentication ProxyUserSummaryAuthenticationEnum `mandatory:"false" json:"authentication,omitempty"`

	// Flags associated with the proxy/client pair
	Flags ProxyUserSummaryFlagsEnum `mandatory:"false" json:"flags,omitempty"`
}

func (m ProxyUserSummary) String() string {
	return common.PointerString(m)
}

// ProxyUserSummaryAuthenticationEnum Enum with underlying type: string
type ProxyUserSummaryAuthenticationEnum string

// Set of constants representing the allowable values for ProxyUserSummaryAuthenticationEnum
const (
	ProxyUserSummaryAuthenticationYes ProxyUserSummaryAuthenticationEnum = "YES"
	ProxyUserSummaryAuthenticationNo  ProxyUserSummaryAuthenticationEnum = "NO"
)

var mappingProxyUserSummaryAuthentication = map[string]ProxyUserSummaryAuthenticationEnum{
	"YES": ProxyUserSummaryAuthenticationYes,
	"NO":  ProxyUserSummaryAuthenticationNo,
}

// GetProxyUserSummaryAuthenticationEnumValues Enumerates the set of values for ProxyUserSummaryAuthenticationEnum
func GetProxyUserSummaryAuthenticationEnumValues() []ProxyUserSummaryAuthenticationEnum {
	values := make([]ProxyUserSummaryAuthenticationEnum, 0)
	for _, v := range mappingProxyUserSummaryAuthentication {
		values = append(values, v)
	}
	return values
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

var mappingProxyUserSummaryFlags = map[string]ProxyUserSummaryFlagsEnum{
	"PROXY_MAY_ACTIVATE_ALL_CLIENT_ROLES": ProxyUserSummaryFlagsProxyMayActivateAllClientRoles,
	"NO_CLIENT_ROLES_MAY_BE_ACTIVATED":    ProxyUserSummaryFlagsNoClientRolesMayBeActivated,
	"PROXY_MAY_ACTIVATE_ROLE":             ProxyUserSummaryFlagsProxyMayActivateRole,
	"PROXY_MAY_NOT_ACTIVATE_ROLE":         ProxyUserSummaryFlagsProxyMayNotActivateRole,
}

// GetProxyUserSummaryFlagsEnumValues Enumerates the set of values for ProxyUserSummaryFlagsEnum
func GetProxyUserSummaryFlagsEnumValues() []ProxyUserSummaryFlagsEnum {
	values := make([]ProxyUserSummaryFlagsEnum, 0)
	for _, v := range mappingProxyUserSummaryFlags {
		values = append(values, v)
	}
	return values
}
