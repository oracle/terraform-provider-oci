// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// DataSourceFeedProviderEnum Enum with underlying type: string
type DataSourceFeedProviderEnum string

// Set of constants representing the allowable values for DataSourceFeedProviderEnum
const (
	DataSourceFeedProviderLoggingquery   DataSourceFeedProviderEnum = "LOGGINGQUERY"
	DataSourceFeedProviderScheduledquery DataSourceFeedProviderEnum = "SCHEDULEDQUERY"
)

var mappingDataSourceFeedProviderEnum = map[string]DataSourceFeedProviderEnum{
	"LOGGINGQUERY":   DataSourceFeedProviderLoggingquery,
	"SCHEDULEDQUERY": DataSourceFeedProviderScheduledquery,
}

var mappingDataSourceFeedProviderEnumLowerCase = map[string]DataSourceFeedProviderEnum{
	"loggingquery":   DataSourceFeedProviderLoggingquery,
	"scheduledquery": DataSourceFeedProviderScheduledquery,
}

// GetDataSourceFeedProviderEnumValues Enumerates the set of values for DataSourceFeedProviderEnum
func GetDataSourceFeedProviderEnumValues() []DataSourceFeedProviderEnum {
	values := make([]DataSourceFeedProviderEnum, 0)
	for _, v := range mappingDataSourceFeedProviderEnum {
		values = append(values, v)
	}
	return values
}

// GetDataSourceFeedProviderEnumStringValues Enumerates the set of values in String for DataSourceFeedProviderEnum
func GetDataSourceFeedProviderEnumStringValues() []string {
	return []string{
		"LOGGINGQUERY",
		"SCHEDULEDQUERY",
	}
}

// GetMappingDataSourceFeedProviderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataSourceFeedProviderEnum(val string) (DataSourceFeedProviderEnum, bool) {
	enum, ok := mappingDataSourceFeedProviderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
