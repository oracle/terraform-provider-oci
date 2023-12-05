// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"strings"
)

// CommitmentEnum Enum with underlying type: string
type CommitmentEnum string

// Set of constants representing the allowable values for CommitmentEnum
const (
	CommitmentHour       CommitmentEnum = "HOUR"
	CommitmentMonth      CommitmentEnum = "MONTH"
	CommitmentOneYear    CommitmentEnum = "ONE_YEAR"
	CommitmentThreeYears CommitmentEnum = "THREE_YEARS"
)

var mappingCommitmentEnum = map[string]CommitmentEnum{
	"HOUR":        CommitmentHour,
	"MONTH":       CommitmentMonth,
	"ONE_YEAR":    CommitmentOneYear,
	"THREE_YEARS": CommitmentThreeYears,
}

var mappingCommitmentEnumLowerCase = map[string]CommitmentEnum{
	"hour":        CommitmentHour,
	"month":       CommitmentMonth,
	"one_year":    CommitmentOneYear,
	"three_years": CommitmentThreeYears,
}

// GetCommitmentEnumValues Enumerates the set of values for CommitmentEnum
func GetCommitmentEnumValues() []CommitmentEnum {
	values := make([]CommitmentEnum, 0)
	for _, v := range mappingCommitmentEnum {
		values = append(values, v)
	}
	return values
}

// GetCommitmentEnumStringValues Enumerates the set of values in String for CommitmentEnum
func GetCommitmentEnumStringValues() []string {
	return []string{
		"HOUR",
		"MONTH",
		"ONE_YEAR",
		"THREE_YEARS",
	}
}

// GetMappingCommitmentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCommitmentEnum(val string) (CommitmentEnum, bool) {
	enum, ok := mappingCommitmentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
