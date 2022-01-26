// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

// ApplicationTypeEnum Enum with underlying type: string
type ApplicationTypeEnum string

// Set of constants representing the allowable values for ApplicationTypeEnum
const (
	ApplicationTypeBatch     ApplicationTypeEnum = "BATCH"
	ApplicationTypeStreaming ApplicationTypeEnum = "STREAMING"
)

var mappingApplicationType = map[string]ApplicationTypeEnum{
	"BATCH":     ApplicationTypeBatch,
	"STREAMING": ApplicationTypeStreaming,
}

// GetApplicationTypeEnumValues Enumerates the set of values for ApplicationTypeEnum
func GetApplicationTypeEnumValues() []ApplicationTypeEnum {
	values := make([]ApplicationTypeEnum, 0)
	for _, v := range mappingApplicationType {
		values = append(values, v)
	}
	return values
}
