// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"strings"
)

// JmsPluginSortByEnum Enum with underlying type: string
type JmsPluginSortByEnum string

// Set of constants representing the allowable values for JmsPluginSortByEnum
const (
	JmsPluginSortById                 JmsPluginSortByEnum = "id"
	JmsPluginSortByTimeLastSeen       JmsPluginSortByEnum = "timeLastSeen"
	JmsPluginSortByTimeRegistered     JmsPluginSortByEnum = "timeRegistered"
	JmsPluginSortByHostname           JmsPluginSortByEnum = "hostname"
	JmsPluginSortByAgentId            JmsPluginSortByEnum = "agentId"
	JmsPluginSortByAgentType          JmsPluginSortByEnum = "agentType"
	JmsPluginSortByLifecycleState     JmsPluginSortByEnum = "lifecycleState"
	JmsPluginSortByAvailabilityStatus JmsPluginSortByEnum = "availabilityStatus"
	JmsPluginSortByFleetId            JmsPluginSortByEnum = "fleetId"
	JmsPluginSortByCompartmentId      JmsPluginSortByEnum = "compartmentId"
	JmsPluginSortByOsFamily           JmsPluginSortByEnum = "osFamily"
	JmsPluginSortByOsArchitecture     JmsPluginSortByEnum = "osArchitecture"
	JmsPluginSortByOsDistribution     JmsPluginSortByEnum = "osDistribution"
	JmsPluginSortByPluginVersion      JmsPluginSortByEnum = "pluginVersion"
)

var mappingJmsPluginSortByEnum = map[string]JmsPluginSortByEnum{
	"id":                 JmsPluginSortById,
	"timeLastSeen":       JmsPluginSortByTimeLastSeen,
	"timeRegistered":     JmsPluginSortByTimeRegistered,
	"hostname":           JmsPluginSortByHostname,
	"agentId":            JmsPluginSortByAgentId,
	"agentType":          JmsPluginSortByAgentType,
	"lifecycleState":     JmsPluginSortByLifecycleState,
	"availabilityStatus": JmsPluginSortByAvailabilityStatus,
	"fleetId":            JmsPluginSortByFleetId,
	"compartmentId":      JmsPluginSortByCompartmentId,
	"osFamily":           JmsPluginSortByOsFamily,
	"osArchitecture":     JmsPluginSortByOsArchitecture,
	"osDistribution":     JmsPluginSortByOsDistribution,
	"pluginVersion":      JmsPluginSortByPluginVersion,
}

var mappingJmsPluginSortByEnumLowerCase = map[string]JmsPluginSortByEnum{
	"id":                 JmsPluginSortById,
	"timelastseen":       JmsPluginSortByTimeLastSeen,
	"timeregistered":     JmsPluginSortByTimeRegistered,
	"hostname":           JmsPluginSortByHostname,
	"agentid":            JmsPluginSortByAgentId,
	"agenttype":          JmsPluginSortByAgentType,
	"lifecyclestate":     JmsPluginSortByLifecycleState,
	"availabilitystatus": JmsPluginSortByAvailabilityStatus,
	"fleetid":            JmsPluginSortByFleetId,
	"compartmentid":      JmsPluginSortByCompartmentId,
	"osfamily":           JmsPluginSortByOsFamily,
	"osarchitecture":     JmsPluginSortByOsArchitecture,
	"osdistribution":     JmsPluginSortByOsDistribution,
	"pluginversion":      JmsPluginSortByPluginVersion,
}

// GetJmsPluginSortByEnumValues Enumerates the set of values for JmsPluginSortByEnum
func GetJmsPluginSortByEnumValues() []JmsPluginSortByEnum {
	values := make([]JmsPluginSortByEnum, 0)
	for _, v := range mappingJmsPluginSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetJmsPluginSortByEnumStringValues Enumerates the set of values in String for JmsPluginSortByEnum
func GetJmsPluginSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeLastSeen",
		"timeRegistered",
		"hostname",
		"agentId",
		"agentType",
		"lifecycleState",
		"availabilityStatus",
		"fleetId",
		"compartmentId",
		"osFamily",
		"osArchitecture",
		"osDistribution",
		"pluginVersion",
	}
}

// GetMappingJmsPluginSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJmsPluginSortByEnum(val string) (JmsPluginSortByEnum, bool) {
	enum, ok := mappingJmsPluginSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
