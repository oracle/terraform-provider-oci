// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// ProxyAuthenticationTypeEnum Enum with underlying type: string
type ProxyAuthenticationTypeEnum string

// Set of constants representing the allowable values for ProxyAuthenticationTypeEnum
const (
	ProxyAuthenticationTypeUserName ProxyAuthenticationTypeEnum = "USER_NAME"
	ProxyAuthenticationTypeNoProxy  ProxyAuthenticationTypeEnum = "NO_PROXY"
)

var mappingProxyAuthenticationTypeEnum = map[string]ProxyAuthenticationTypeEnum{
	"USER_NAME": ProxyAuthenticationTypeUserName,
	"NO_PROXY":  ProxyAuthenticationTypeNoProxy,
}

var mappingProxyAuthenticationTypeEnumLowerCase = map[string]ProxyAuthenticationTypeEnum{
	"user_name": ProxyAuthenticationTypeUserName,
	"no_proxy":  ProxyAuthenticationTypeNoProxy,
}

// GetProxyAuthenticationTypeEnumValues Enumerates the set of values for ProxyAuthenticationTypeEnum
func GetProxyAuthenticationTypeEnumValues() []ProxyAuthenticationTypeEnum {
	values := make([]ProxyAuthenticationTypeEnum, 0)
	for _, v := range mappingProxyAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetProxyAuthenticationTypeEnumStringValues Enumerates the set of values in String for ProxyAuthenticationTypeEnum
func GetProxyAuthenticationTypeEnumStringValues() []string {
	return []string{
		"USER_NAME",
		"NO_PROXY",
	}
}

// GetMappingProxyAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProxyAuthenticationTypeEnum(val string) (ProxyAuthenticationTypeEnum, bool) {
	enum, ok := mappingProxyAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
