// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// DataMaskRuleStatusEnum Enum with underlying type: string
type DataMaskRuleStatusEnum string

// Set of constants representing the allowable values for DataMaskRuleStatusEnum
const (
	DataMaskRuleStatusEnabled  DataMaskRuleStatusEnum = "ENABLED"
	DataMaskRuleStatusDisabled DataMaskRuleStatusEnum = "DISABLED"
)

var mappingDataMaskRuleStatus = map[string]DataMaskRuleStatusEnum{
	"ENABLED":  DataMaskRuleStatusEnabled,
	"DISABLED": DataMaskRuleStatusDisabled,
}

// GetDataMaskRuleStatusEnumValues Enumerates the set of values for DataMaskRuleStatusEnum
func GetDataMaskRuleStatusEnumValues() []DataMaskRuleStatusEnum {
	values := make([]DataMaskRuleStatusEnum, 0)
	for _, v := range mappingDataMaskRuleStatus {
		values = append(values, v)
	}
	return values
}
