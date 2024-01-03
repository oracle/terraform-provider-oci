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

// OnPremiseVantagePointWorkerStatusEnum Enum with underlying type: string
type OnPremiseVantagePointWorkerStatusEnum string

// Set of constants representing the allowable values for OnPremiseVantagePointWorkerStatusEnum
const (
	OnPremiseVantagePointWorkerStatusEnabled  OnPremiseVantagePointWorkerStatusEnum = "ENABLED"
	OnPremiseVantagePointWorkerStatusDisabled OnPremiseVantagePointWorkerStatusEnum = "DISABLED"
)

var mappingOnPremiseVantagePointWorkerStatusEnum = map[string]OnPremiseVantagePointWorkerStatusEnum{
	"ENABLED":  OnPremiseVantagePointWorkerStatusEnabled,
	"DISABLED": OnPremiseVantagePointWorkerStatusDisabled,
}

var mappingOnPremiseVantagePointWorkerStatusEnumLowerCase = map[string]OnPremiseVantagePointWorkerStatusEnum{
	"enabled":  OnPremiseVantagePointWorkerStatusEnabled,
	"disabled": OnPremiseVantagePointWorkerStatusDisabled,
}

// GetOnPremiseVantagePointWorkerStatusEnumValues Enumerates the set of values for OnPremiseVantagePointWorkerStatusEnum
func GetOnPremiseVantagePointWorkerStatusEnumValues() []OnPremiseVantagePointWorkerStatusEnum {
	values := make([]OnPremiseVantagePointWorkerStatusEnum, 0)
	for _, v := range mappingOnPremiseVantagePointWorkerStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOnPremiseVantagePointWorkerStatusEnumStringValues Enumerates the set of values in String for OnPremiseVantagePointWorkerStatusEnum
func GetOnPremiseVantagePointWorkerStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingOnPremiseVantagePointWorkerStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOnPremiseVantagePointWorkerStatusEnum(val string) (OnPremiseVantagePointWorkerStatusEnum, bool) {
	enum, ok := mappingOnPremiseVantagePointWorkerStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
