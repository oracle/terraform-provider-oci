// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"strings"
)

// OnPremiseVantagePointWorkerTypeEnum Enum with underlying type: string
type OnPremiseVantagePointWorkerTypeEnum string

// Set of constants representing the allowable values for OnPremiseVantagePointWorkerTypeEnum
const (
	OnPremiseVantagePointWorkerTypeDocker OnPremiseVantagePointWorkerTypeEnum = "DOCKER"
)

var mappingOnPremiseVantagePointWorkerTypeEnum = map[string]OnPremiseVantagePointWorkerTypeEnum{
	"DOCKER": OnPremiseVantagePointWorkerTypeDocker,
}

var mappingOnPremiseVantagePointWorkerTypeEnumLowerCase = map[string]OnPremiseVantagePointWorkerTypeEnum{
	"docker": OnPremiseVantagePointWorkerTypeDocker,
}

// GetOnPremiseVantagePointWorkerTypeEnumValues Enumerates the set of values for OnPremiseVantagePointWorkerTypeEnum
func GetOnPremiseVantagePointWorkerTypeEnumValues() []OnPremiseVantagePointWorkerTypeEnum {
	values := make([]OnPremiseVantagePointWorkerTypeEnum, 0)
	for _, v := range mappingOnPremiseVantagePointWorkerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOnPremiseVantagePointWorkerTypeEnumStringValues Enumerates the set of values in String for OnPremiseVantagePointWorkerTypeEnum
func GetOnPremiseVantagePointWorkerTypeEnumStringValues() []string {
	return []string{
		"DOCKER",
	}
}

// GetMappingOnPremiseVantagePointWorkerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOnPremiseVantagePointWorkerTypeEnum(val string) (OnPremiseVantagePointWorkerTypeEnum, bool) {
	enum, ok := mappingOnPremiseVantagePointWorkerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
