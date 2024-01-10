// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// CryptoAnalysisResultSortByEnum Enum with underlying type: string
type CryptoAnalysisResultSortByEnum string

// Set of constants representing the allowable values for CryptoAnalysisResultSortByEnum
const (
	CryptoAnalysisResultSortByTimeCreated       CryptoAnalysisResultSortByEnum = "timeCreated"
	CryptoAnalysisResultSortByManagedInstanceId CryptoAnalysisResultSortByEnum = "managedInstanceId"
	CryptoAnalysisResultSortByWorkRequestId     CryptoAnalysisResultSortByEnum = "workRequestId"
)

var mappingCryptoAnalysisResultSortByEnum = map[string]CryptoAnalysisResultSortByEnum{
	"timeCreated":       CryptoAnalysisResultSortByTimeCreated,
	"managedInstanceId": CryptoAnalysisResultSortByManagedInstanceId,
	"workRequestId":     CryptoAnalysisResultSortByWorkRequestId,
}

var mappingCryptoAnalysisResultSortByEnumLowerCase = map[string]CryptoAnalysisResultSortByEnum{
	"timecreated":       CryptoAnalysisResultSortByTimeCreated,
	"managedinstanceid": CryptoAnalysisResultSortByManagedInstanceId,
	"workrequestid":     CryptoAnalysisResultSortByWorkRequestId,
}

// GetCryptoAnalysisResultSortByEnumValues Enumerates the set of values for CryptoAnalysisResultSortByEnum
func GetCryptoAnalysisResultSortByEnumValues() []CryptoAnalysisResultSortByEnum {
	values := make([]CryptoAnalysisResultSortByEnum, 0)
	for _, v := range mappingCryptoAnalysisResultSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAnalysisResultSortByEnumStringValues Enumerates the set of values in String for CryptoAnalysisResultSortByEnum
func GetCryptoAnalysisResultSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"managedInstanceId",
		"workRequestId",
	}
}

// GetMappingCryptoAnalysisResultSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAnalysisResultSortByEnum(val string) (CryptoAnalysisResultSortByEnum, bool) {
	enum, ok := mappingCryptoAnalysisResultSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
