// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle API Access Control
//
// This service is used to restrict the control plane service apis; so that everybody won't be
// able to access those apis.
// There are two main resouces defined as a part of this service
// 1. PrivilegedApiControl: This is created by the customer which defines which service apis are
//    controlled and who can access it.
// 2. PrivilegedApiRequest: This is a request object again created by the customer operators who           seek access to those privileged apis. After a request is obtained based on the                       PrivilegedAccessControl for which the api belongs to, either it can be approved so that the          requested person can execute the service apis or it will wait for the customer to approve it.
//

package apiaccesscontrol

import (
	"strings"
)

// PrivilegedApiRequestSeverityEnum Enum with underlying type: string
type PrivilegedApiRequestSeverityEnum string

// Set of constants representing the allowable values for PrivilegedApiRequestSeverityEnum
const (
	PrivilegedApiRequestSeveritySev1 PrivilegedApiRequestSeverityEnum = "SEV_1"
	PrivilegedApiRequestSeveritySev2 PrivilegedApiRequestSeverityEnum = "SEV_2"
	PrivilegedApiRequestSeveritySev3 PrivilegedApiRequestSeverityEnum = "SEV_3"
	PrivilegedApiRequestSeveritySev4 PrivilegedApiRequestSeverityEnum = "SEV_4"
)

var mappingPrivilegedApiRequestSeverityEnum = map[string]PrivilegedApiRequestSeverityEnum{
	"SEV_1": PrivilegedApiRequestSeveritySev1,
	"SEV_2": PrivilegedApiRequestSeveritySev2,
	"SEV_3": PrivilegedApiRequestSeveritySev3,
	"SEV_4": PrivilegedApiRequestSeveritySev4,
}

var mappingPrivilegedApiRequestSeverityEnumLowerCase = map[string]PrivilegedApiRequestSeverityEnum{
	"sev_1": PrivilegedApiRequestSeveritySev1,
	"sev_2": PrivilegedApiRequestSeveritySev2,
	"sev_3": PrivilegedApiRequestSeveritySev3,
	"sev_4": PrivilegedApiRequestSeveritySev4,
}

// GetPrivilegedApiRequestSeverityEnumValues Enumerates the set of values for PrivilegedApiRequestSeverityEnum
func GetPrivilegedApiRequestSeverityEnumValues() []PrivilegedApiRequestSeverityEnum {
	values := make([]PrivilegedApiRequestSeverityEnum, 0)
	for _, v := range mappingPrivilegedApiRequestSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivilegedApiRequestSeverityEnumStringValues Enumerates the set of values in String for PrivilegedApiRequestSeverityEnum
func GetPrivilegedApiRequestSeverityEnumStringValues() []string {
	return []string{
		"SEV_1",
		"SEV_2",
		"SEV_3",
		"SEV_4",
	}
}

// GetMappingPrivilegedApiRequestSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivilegedApiRequestSeverityEnum(val string) (PrivilegedApiRequestSeverityEnum, bool) {
	enum, ok := mappingPrivilegedApiRequestSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
