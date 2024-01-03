// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.cloud.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"strings"
)

// ConnectorPluginLifecycleStateEnum Enum with underlying type: string
type ConnectorPluginLifecycleStateEnum string

// Set of constants representing the allowable values for ConnectorPluginLifecycleStateEnum
const (
	ConnectorPluginLifecycleStateActive  ConnectorPluginLifecycleStateEnum = "ACTIVE"
	ConnectorPluginLifecycleStateDeleted ConnectorPluginLifecycleStateEnum = "DELETED"
)

var mappingConnectorPluginLifecycleStateEnum = map[string]ConnectorPluginLifecycleStateEnum{
	"ACTIVE":  ConnectorPluginLifecycleStateActive,
	"DELETED": ConnectorPluginLifecycleStateDeleted,
}

var mappingConnectorPluginLifecycleStateEnumLowerCase = map[string]ConnectorPluginLifecycleStateEnum{
	"active":  ConnectorPluginLifecycleStateActive,
	"deleted": ConnectorPluginLifecycleStateDeleted,
}

// GetConnectorPluginLifecycleStateEnumValues Enumerates the set of values for ConnectorPluginLifecycleStateEnum
func GetConnectorPluginLifecycleStateEnumValues() []ConnectorPluginLifecycleStateEnum {
	values := make([]ConnectorPluginLifecycleStateEnum, 0)
	for _, v := range mappingConnectorPluginLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectorPluginLifecycleStateEnumStringValues Enumerates the set of values in String for ConnectorPluginLifecycleStateEnum
func GetConnectorPluginLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingConnectorPluginLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectorPluginLifecycleStateEnum(val string) (ConnectorPluginLifecycleStateEnum, bool) {
	enum, ok := mappingConnectorPluginLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
