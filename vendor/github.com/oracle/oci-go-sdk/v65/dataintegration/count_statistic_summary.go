// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CountStatisticSummary Details of the count statistic summary object.
type CountStatisticSummary struct {

	// The type of object for the count statistic object.
	ObjectType CountStatisticSummaryObjectTypeEnum `mandatory:"false" json:"objectType,omitempty"`

	// The value for the count statistic object.
	ObjectCount *int64 `mandatory:"false" json:"objectCount"`
}

func (m CountStatisticSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CountStatisticSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCountStatisticSummaryObjectTypeEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetCountStatisticSummaryObjectTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CountStatisticSummaryObjectTypeEnum Enum with underlying type: string
type CountStatisticSummaryObjectTypeEnum string

// Set of constants representing the allowable values for CountStatisticSummaryObjectTypeEnum
const (
	CountStatisticSummaryObjectTypeProject             CountStatisticSummaryObjectTypeEnum = "PROJECT"
	CountStatisticSummaryObjectTypeFolder              CountStatisticSummaryObjectTypeEnum = "FOLDER"
	CountStatisticSummaryObjectTypeDataFlow            CountStatisticSummaryObjectTypeEnum = "DATA_FLOW"
	CountStatisticSummaryObjectTypeDataAsset           CountStatisticSummaryObjectTypeEnum = "DATA_ASSET"
	CountStatisticSummaryObjectTypeConnection          CountStatisticSummaryObjectTypeEnum = "CONNECTION"
	CountStatisticSummaryObjectTypeTask                CountStatisticSummaryObjectTypeEnum = "TASK"
	CountStatisticSummaryObjectTypeApplication         CountStatisticSummaryObjectTypeEnum = "APPLICATION"
	CountStatisticSummaryObjectTypeFunctionLibrary     CountStatisticSummaryObjectTypeEnum = "FUNCTION_LIBRARY"
	CountStatisticSummaryObjectTypeUserDefinedFunction CountStatisticSummaryObjectTypeEnum = "USER_DEFINED_FUNCTION"
)

var mappingCountStatisticSummaryObjectTypeEnum = map[string]CountStatisticSummaryObjectTypeEnum{
	"PROJECT":               CountStatisticSummaryObjectTypeProject,
	"FOLDER":                CountStatisticSummaryObjectTypeFolder,
	"DATA_FLOW":             CountStatisticSummaryObjectTypeDataFlow,
	"DATA_ASSET":            CountStatisticSummaryObjectTypeDataAsset,
	"CONNECTION":            CountStatisticSummaryObjectTypeConnection,
	"TASK":                  CountStatisticSummaryObjectTypeTask,
	"APPLICATION":           CountStatisticSummaryObjectTypeApplication,
	"FUNCTION_LIBRARY":      CountStatisticSummaryObjectTypeFunctionLibrary,
	"USER_DEFINED_FUNCTION": CountStatisticSummaryObjectTypeUserDefinedFunction,
}

var mappingCountStatisticSummaryObjectTypeEnumLowerCase = map[string]CountStatisticSummaryObjectTypeEnum{
	"project":               CountStatisticSummaryObjectTypeProject,
	"folder":                CountStatisticSummaryObjectTypeFolder,
	"data_flow":             CountStatisticSummaryObjectTypeDataFlow,
	"data_asset":            CountStatisticSummaryObjectTypeDataAsset,
	"connection":            CountStatisticSummaryObjectTypeConnection,
	"task":                  CountStatisticSummaryObjectTypeTask,
	"application":           CountStatisticSummaryObjectTypeApplication,
	"function_library":      CountStatisticSummaryObjectTypeFunctionLibrary,
	"user_defined_function": CountStatisticSummaryObjectTypeUserDefinedFunction,
}

// GetCountStatisticSummaryObjectTypeEnumValues Enumerates the set of values for CountStatisticSummaryObjectTypeEnum
func GetCountStatisticSummaryObjectTypeEnumValues() []CountStatisticSummaryObjectTypeEnum {
	values := make([]CountStatisticSummaryObjectTypeEnum, 0)
	for _, v := range mappingCountStatisticSummaryObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCountStatisticSummaryObjectTypeEnumStringValues Enumerates the set of values in String for CountStatisticSummaryObjectTypeEnum
func GetCountStatisticSummaryObjectTypeEnumStringValues() []string {
	return []string{
		"PROJECT",
		"FOLDER",
		"DATA_FLOW",
		"DATA_ASSET",
		"CONNECTION",
		"TASK",
		"APPLICATION",
		"FUNCTION_LIBRARY",
		"USER_DEFINED_FUNCTION",
	}
}

// GetMappingCountStatisticSummaryObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCountStatisticSummaryObjectTypeEnum(val string) (CountStatisticSummaryObjectTypeEnum, bool) {
	enum, ok := mappingCountStatisticSummaryObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
