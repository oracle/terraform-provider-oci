// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// ResponderDimensionEnum Enum with underlying type: string
type ResponderDimensionEnum string

// Set of constants representing the allowable values for ResponderDimensionEnum
const (
	ResponderDimensionResponderRuleType        ResponderDimensionEnum = "RESPONDER_RULE_TYPE"
	ResponderDimensionResponderExecutionStatus ResponderDimensionEnum = "RESPONDER_EXECUTION_STATUS"
)

var mappingResponderDimension = map[string]ResponderDimensionEnum{
	"RESPONDER_RULE_TYPE":        ResponderDimensionResponderRuleType,
	"RESPONDER_EXECUTION_STATUS": ResponderDimensionResponderExecutionStatus,
}

// GetResponderDimensionEnumValues Enumerates the set of values for ResponderDimensionEnum
func GetResponderDimensionEnumValues() []ResponderDimensionEnum {
	values := make([]ResponderDimensionEnum, 0)
	for _, v := range mappingResponderDimension {
		values = append(values, v)
	}
	return values
}
