// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// SourceTypeEnum Enum with underlying type: string
type SourceTypeEnum string

// Set of constants representing the allowable values for SourceTypeEnum
const (
	SourceTypeSmMgmtAgentMonitored SourceTypeEnum = "SM_MGMT_AGENT_MONITORED"
	SourceTypeSmRepoOnly           SourceTypeEnum = "SM_REPO_ONLY"
	SourceTypeOciNative            SourceTypeEnum = "OCI_NATIVE"
	SourceTypePrometheus           SourceTypeEnum = "PROMETHEUS"
	SourceTypeTelegraf             SourceTypeEnum = "TELEGRAF"
	SourceTypeCollectd             SourceTypeEnum = "COLLECTD"
)

var mappingSourceTypeEnum = map[string]SourceTypeEnum{
	"SM_MGMT_AGENT_MONITORED": SourceTypeSmMgmtAgentMonitored,
	"SM_REPO_ONLY":            SourceTypeSmRepoOnly,
	"OCI_NATIVE":              SourceTypeOciNative,
	"PROMETHEUS":              SourceTypePrometheus,
	"TELEGRAF":                SourceTypeTelegraf,
	"COLLECTD":                SourceTypeCollectd,
}

var mappingSourceTypeEnumLowerCase = map[string]SourceTypeEnum{
	"sm_mgmt_agent_monitored": SourceTypeSmMgmtAgentMonitored,
	"sm_repo_only":            SourceTypeSmRepoOnly,
	"oci_native":              SourceTypeOciNative,
	"prometheus":              SourceTypePrometheus,
	"telegraf":                SourceTypeTelegraf,
	"collectd":                SourceTypeCollectd,
}

// GetSourceTypeEnumValues Enumerates the set of values for SourceTypeEnum
func GetSourceTypeEnumValues() []SourceTypeEnum {
	values := make([]SourceTypeEnum, 0)
	for _, v := range mappingSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSourceTypeEnumStringValues Enumerates the set of values in String for SourceTypeEnum
func GetSourceTypeEnumStringValues() []string {
	return []string{
		"SM_MGMT_AGENT_MONITORED",
		"SM_REPO_ONLY",
		"OCI_NATIVE",
		"PROMETHEUS",
		"TELEGRAF",
		"COLLECTD",
	}
}

// GetMappingSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSourceTypeEnum(val string) (SourceTypeEnum, bool) {
	enum, ok := mappingSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
