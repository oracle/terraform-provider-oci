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

// IncrementalDataEntityClause Data Entity clause for Incremental Read operation.
type IncrementalDataEntityClause struct {

	// Name of incremental data entity filter.
	IncrementalDataEntityName *string `mandatory:"true" json:"incrementalDataEntityName"`

	// Value of incremental data entity filter.
	IncrementalDataEntityValue map[string]string `mandatory:"true" json:"incrementalDataEntityValue"`

	// Incremental comparator symbol.
	IncrementalComparator IncrementalDataEntityClauseIncrementalComparatorEnum `mandatory:"true" json:"incrementalComparator"`
}

func (m IncrementalDataEntityClause) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IncrementalDataEntityClause) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIncrementalDataEntityClauseIncrementalComparatorEnum(string(m.IncrementalComparator)); !ok && m.IncrementalComparator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IncrementalComparator: %s. Supported values are: %s.", m.IncrementalComparator, strings.Join(GetIncrementalDataEntityClauseIncrementalComparatorEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IncrementalDataEntityClauseIncrementalComparatorEnum Enum with underlying type: string
type IncrementalDataEntityClauseIncrementalComparatorEnum string

// Set of constants representing the allowable values for IncrementalDataEntityClauseIncrementalComparatorEnum
const (
	IncrementalDataEntityClauseIncrementalComparatorLessthan          IncrementalDataEntityClauseIncrementalComparatorEnum = "LESSTHAN"
	IncrementalDataEntityClauseIncrementalComparatorGreaterthan       IncrementalDataEntityClauseIncrementalComparatorEnum = "GREATERTHAN"
	IncrementalDataEntityClauseIncrementalComparatorEquals            IncrementalDataEntityClauseIncrementalComparatorEnum = "EQUALS"
	IncrementalDataEntityClauseIncrementalComparatorLessthanequals    IncrementalDataEntityClauseIncrementalComparatorEnum = "LESSTHANEQUALS"
	IncrementalDataEntityClauseIncrementalComparatorGreaterthanequals IncrementalDataEntityClauseIncrementalComparatorEnum = "GREATERTHANEQUALS"
	IncrementalDataEntityClauseIncrementalComparatorStartswith        IncrementalDataEntityClauseIncrementalComparatorEnum = "STARTSWITH"
	IncrementalDataEntityClauseIncrementalComparatorContains          IncrementalDataEntityClauseIncrementalComparatorEnum = "CONTAINS"
)

var mappingIncrementalDataEntityClauseIncrementalComparatorEnum = map[string]IncrementalDataEntityClauseIncrementalComparatorEnum{
	"LESSTHAN":          IncrementalDataEntityClauseIncrementalComparatorLessthan,
	"GREATERTHAN":       IncrementalDataEntityClauseIncrementalComparatorGreaterthan,
	"EQUALS":            IncrementalDataEntityClauseIncrementalComparatorEquals,
	"LESSTHANEQUALS":    IncrementalDataEntityClauseIncrementalComparatorLessthanequals,
	"GREATERTHANEQUALS": IncrementalDataEntityClauseIncrementalComparatorGreaterthanequals,
	"STARTSWITH":        IncrementalDataEntityClauseIncrementalComparatorStartswith,
	"CONTAINS":          IncrementalDataEntityClauseIncrementalComparatorContains,
}

var mappingIncrementalDataEntityClauseIncrementalComparatorEnumLowerCase = map[string]IncrementalDataEntityClauseIncrementalComparatorEnum{
	"lessthan":          IncrementalDataEntityClauseIncrementalComparatorLessthan,
	"greaterthan":       IncrementalDataEntityClauseIncrementalComparatorGreaterthan,
	"equals":            IncrementalDataEntityClauseIncrementalComparatorEquals,
	"lessthanequals":    IncrementalDataEntityClauseIncrementalComparatorLessthanequals,
	"greaterthanequals": IncrementalDataEntityClauseIncrementalComparatorGreaterthanequals,
	"startswith":        IncrementalDataEntityClauseIncrementalComparatorStartswith,
	"contains":          IncrementalDataEntityClauseIncrementalComparatorContains,
}

// GetIncrementalDataEntityClauseIncrementalComparatorEnumValues Enumerates the set of values for IncrementalDataEntityClauseIncrementalComparatorEnum
func GetIncrementalDataEntityClauseIncrementalComparatorEnumValues() []IncrementalDataEntityClauseIncrementalComparatorEnum {
	values := make([]IncrementalDataEntityClauseIncrementalComparatorEnum, 0)
	for _, v := range mappingIncrementalDataEntityClauseIncrementalComparatorEnum {
		values = append(values, v)
	}
	return values
}

// GetIncrementalDataEntityClauseIncrementalComparatorEnumStringValues Enumerates the set of values in String for IncrementalDataEntityClauseIncrementalComparatorEnum
func GetIncrementalDataEntityClauseIncrementalComparatorEnumStringValues() []string {
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

// GetMappingIncrementalDataEntityClauseIncrementalComparatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIncrementalDataEntityClauseIncrementalComparatorEnum(val string) (IncrementalDataEntityClauseIncrementalComparatorEnum, bool) {
	enum, ok := mappingIncrementalDataEntityClauseIncrementalComparatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
