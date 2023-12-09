// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// IncrementalFieldClause Field clause for incremental read operation.
type IncrementalFieldClause struct {

	// Name of incremental field filter.
	IncrementalFieldName *string `mandatory:"true" json:"incrementalFieldName"`

	// Value of incremental field filter.
	IncrementalFieldValue map[string]string `mandatory:"true" json:"incrementalFieldValue"`

	// Incremental comparator symbol.
	IncrementalComparator IncrementalFieldClauseIncrementalComparatorEnum `mandatory:"true" json:"incrementalComparator"`
}

func (m IncrementalFieldClause) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IncrementalFieldClause) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIncrementalFieldClauseIncrementalComparatorEnum(string(m.IncrementalComparator)); !ok && m.IncrementalComparator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IncrementalComparator: %s. Supported values are: %s.", m.IncrementalComparator, strings.Join(GetIncrementalFieldClauseIncrementalComparatorEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IncrementalFieldClauseIncrementalComparatorEnum Enum with underlying type: string
type IncrementalFieldClauseIncrementalComparatorEnum string

// Set of constants representing the allowable values for IncrementalFieldClauseIncrementalComparatorEnum
const (
	IncrementalFieldClauseIncrementalComparatorLessthan          IncrementalFieldClauseIncrementalComparatorEnum = "LESSTHAN"
	IncrementalFieldClauseIncrementalComparatorGreaterthan       IncrementalFieldClauseIncrementalComparatorEnum = "GREATERTHAN"
	IncrementalFieldClauseIncrementalComparatorEquals            IncrementalFieldClauseIncrementalComparatorEnum = "EQUALS"
	IncrementalFieldClauseIncrementalComparatorLessthanequals    IncrementalFieldClauseIncrementalComparatorEnum = "LESSTHANEQUALS"
	IncrementalFieldClauseIncrementalComparatorGreaterthanequals IncrementalFieldClauseIncrementalComparatorEnum = "GREATERTHANEQUALS"
	IncrementalFieldClauseIncrementalComparatorStartswith        IncrementalFieldClauseIncrementalComparatorEnum = "STARTSWITH"
	IncrementalFieldClauseIncrementalComparatorContains          IncrementalFieldClauseIncrementalComparatorEnum = "CONTAINS"
)

var mappingIncrementalFieldClauseIncrementalComparatorEnum = map[string]IncrementalFieldClauseIncrementalComparatorEnum{
	"LESSTHAN":          IncrementalFieldClauseIncrementalComparatorLessthan,
	"GREATERTHAN":       IncrementalFieldClauseIncrementalComparatorGreaterthan,
	"EQUALS":            IncrementalFieldClauseIncrementalComparatorEquals,
	"LESSTHANEQUALS":    IncrementalFieldClauseIncrementalComparatorLessthanequals,
	"GREATERTHANEQUALS": IncrementalFieldClauseIncrementalComparatorGreaterthanequals,
	"STARTSWITH":        IncrementalFieldClauseIncrementalComparatorStartswith,
	"CONTAINS":          IncrementalFieldClauseIncrementalComparatorContains,
}

var mappingIncrementalFieldClauseIncrementalComparatorEnumLowerCase = map[string]IncrementalFieldClauseIncrementalComparatorEnum{
	"lessthan":          IncrementalFieldClauseIncrementalComparatorLessthan,
	"greaterthan":       IncrementalFieldClauseIncrementalComparatorGreaterthan,
	"equals":            IncrementalFieldClauseIncrementalComparatorEquals,
	"lessthanequals":    IncrementalFieldClauseIncrementalComparatorLessthanequals,
	"greaterthanequals": IncrementalFieldClauseIncrementalComparatorGreaterthanequals,
	"startswith":        IncrementalFieldClauseIncrementalComparatorStartswith,
	"contains":          IncrementalFieldClauseIncrementalComparatorContains,
}

// GetIncrementalFieldClauseIncrementalComparatorEnumValues Enumerates the set of values for IncrementalFieldClauseIncrementalComparatorEnum
func GetIncrementalFieldClauseIncrementalComparatorEnumValues() []IncrementalFieldClauseIncrementalComparatorEnum {
	values := make([]IncrementalFieldClauseIncrementalComparatorEnum, 0)
	for _, v := range mappingIncrementalFieldClauseIncrementalComparatorEnum {
		values = append(values, v)
	}
	return values
}

// GetIncrementalFieldClauseIncrementalComparatorEnumStringValues Enumerates the set of values in String for IncrementalFieldClauseIncrementalComparatorEnum
func GetIncrementalFieldClauseIncrementalComparatorEnumStringValues() []string {
	return []string{
		"LESSTHAN",
		"GREATERTHAN",
		"EQUALS",
		"LESSTHANEQUALS",
		"GREATERTHANEQUALS",
		"STARTSWITH",
		"CONTAINS",
	}
}

// GetMappingIncrementalFieldClauseIncrementalComparatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIncrementalFieldClauseIncrementalComparatorEnum(val string) (IncrementalFieldClauseIncrementalComparatorEnum, bool) {
	enum, ok := mappingIncrementalFieldClauseIncrementalComparatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
