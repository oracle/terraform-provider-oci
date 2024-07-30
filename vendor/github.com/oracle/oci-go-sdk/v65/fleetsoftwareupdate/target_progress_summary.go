// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TargetProgressSummary Progress details of the executing job for a Database target.
type TargetProgressSummary struct {

	// Type of operations being executed.
	OperationType TargetProgressSummaryOperationTypeEnum `mandatory:"false" json:"operationType,omitempty"`

	// Percentage of progress of the operation in execution.
	ProgressOfOperation *int `mandatory:"false" json:"progressOfOperation"`
}

func (m TargetProgressSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetProgressSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTargetProgressSummaryOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetTargetProgressSummaryOperationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TargetProgressSummaryOperationTypeEnum Enum with underlying type: string
type TargetProgressSummaryOperationTypeEnum string

// Set of constants representing the allowable values for TargetProgressSummaryOperationTypeEnum
const (
	TargetProgressSummaryOperationTypeStage    TargetProgressSummaryOperationTypeEnum = "STAGE"
	TargetProgressSummaryOperationTypePrecheck TargetProgressSummaryOperationTypeEnum = "PRECHECK"
	TargetProgressSummaryOperationTypeApply    TargetProgressSummaryOperationTypeEnum = "APPLY"
	TargetProgressSummaryOperationTypeRollback TargetProgressSummaryOperationTypeEnum = "ROLLBACK"
)

var mappingTargetProgressSummaryOperationTypeEnum = map[string]TargetProgressSummaryOperationTypeEnum{
	"STAGE":    TargetProgressSummaryOperationTypeStage,
	"PRECHECK": TargetProgressSummaryOperationTypePrecheck,
	"APPLY":    TargetProgressSummaryOperationTypeApply,
	"ROLLBACK": TargetProgressSummaryOperationTypeRollback,
}

var mappingTargetProgressSummaryOperationTypeEnumLowerCase = map[string]TargetProgressSummaryOperationTypeEnum{
	"stage":    TargetProgressSummaryOperationTypeStage,
	"precheck": TargetProgressSummaryOperationTypePrecheck,
	"apply":    TargetProgressSummaryOperationTypeApply,
	"rollback": TargetProgressSummaryOperationTypeRollback,
}

// GetTargetProgressSummaryOperationTypeEnumValues Enumerates the set of values for TargetProgressSummaryOperationTypeEnum
func GetTargetProgressSummaryOperationTypeEnumValues() []TargetProgressSummaryOperationTypeEnum {
	values := make([]TargetProgressSummaryOperationTypeEnum, 0)
	for _, v := range mappingTargetProgressSummaryOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetProgressSummaryOperationTypeEnumStringValues Enumerates the set of values in String for TargetProgressSummaryOperationTypeEnum
func GetTargetProgressSummaryOperationTypeEnumStringValues() []string {
	return []string{
		"STAGE",
		"PRECHECK",
		"APPLY",
		"ROLLBACK",
	}
}

// GetMappingTargetProgressSummaryOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetProgressSummaryOperationTypeEnum(val string) (TargetProgressSummaryOperationTypeEnum, bool) {
	enum, ok := mappingTargetProgressSummaryOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
