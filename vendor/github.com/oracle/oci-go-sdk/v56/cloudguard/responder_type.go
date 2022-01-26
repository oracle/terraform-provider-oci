// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// ResponderTypeEnum Enum with underlying type: string
type ResponderTypeEnum string

// Set of constants representing the allowable values for ResponderTypeEnum
const (
	ResponderTypeRemediation  ResponderTypeEnum = "REMEDIATION"
	ResponderTypeNotification ResponderTypeEnum = "NOTIFICATION"
)

var mappingResponderType = map[string]ResponderTypeEnum{
	"REMEDIATION":  ResponderTypeRemediation,
	"NOTIFICATION": ResponderTypeNotification,
}

// GetResponderTypeEnumValues Enumerates the set of values for ResponderTypeEnum
func GetResponderTypeEnumValues() []ResponderTypeEnum {
	values := make([]ResponderTypeEnum, 0)
	for _, v := range mappingResponderType {
		values = append(values, v)
	}
	return values
}
