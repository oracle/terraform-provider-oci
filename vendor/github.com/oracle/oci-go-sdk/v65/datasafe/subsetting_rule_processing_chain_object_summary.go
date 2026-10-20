// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubsettingRuleProcessingChainObjectSummary Summary of a subsetting schema relation processed while extracting rows for processing a subsetting rule.
type SubsettingRuleProcessingChainObjectSummary struct {

	// The unique key that identifies a subsetting relation processed. The key is numeric and unique within a processing order
	Key *string `mandatory:"true" json:"key"`

	// The impact on the related table due to the processing of subsetting rule
	PropagationImpact SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum `mandatory:"true" json:"propagationImpact"`

	// The unique key that identifies a subsetting relation.
	SubsettingSchemaRelationKey *string `mandatory:"false" json:"subsettingSchemaRelationKey"`

	// The database schema that contains the parent subsetting table
	ParentSchemaName *string `mandatory:"false" json:"parentSchemaName"`

	// The name of the parent subsetting table
	ParentObjectName *string `mandatory:"false" json:"parentObjectName"`

	// Unique identifiers identifying the parents columns in the relation.
	ParentColumns []string `mandatory:"false" json:"parentColumns"`

	// The database schema that contains the child subsetting table
	ChildSchemaName *string `mandatory:"false" json:"childSchemaName"`

	// The name of the child subsetting table
	ChildObjectName *string `mandatory:"false" json:"childObjectName"`

	// Unique identifiers identifying the child columns in the relation.
	ChildColumns []string `mandatory:"false" json:"childColumns"`

	// The approximate count of rows in the subsetting table before subsetting
	ApproximateRowCountBeforeSubsetting *int64 `mandatory:"false" json:"approximateRowCountBeforeSubsetting"`

	// The estimated count of rows in the subsetting table after subsetting
	EstimatedRowCountAfterSubsetting *int64 `mandatory:"false" json:"estimatedRowCountAfterSubsetting"`

	// Indicates if this object/edge is enabled for processing
	IsEnabledForProcessing *bool `mandatory:"false" json:"isEnabledForProcessing"`
}

func (m SubsettingRuleProcessingChainObjectSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubsettingRuleProcessingChainObjectSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum(string(m.PropagationImpact)); !ok && m.PropagationImpact != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PropagationImpact: %s. Supported values are: %s.", m.PropagationImpact, strings.Join(GetSubsettingRuleProcessingChainObjectSummaryPropagationImpactEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum Enum with underlying type: string
type SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum string

// Set of constants representing the allowable values for SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum
const (
	SubsettingRuleProcessingChainObjectSummaryPropagationImpactSubsetTable       SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum = "SUBSET_TABLE"
	SubsettingRuleProcessingChainObjectSummaryPropagationImpactParentChildSubset SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum = "PARENT_CHILD_SUBSET"
	SubsettingRuleProcessingChainObjectSummaryPropagationImpactChildParentSubset SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum = "CHILD_PARENT_SUBSET"
	SubsettingRuleProcessingChainObjectSummaryPropagationImpactTruncate          SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum = "TRUNCATE"
	SubsettingRuleProcessingChainObjectSummaryPropagationImpactKeep              SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum = "KEEP"
)

var mappingSubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum = map[string]SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum{
	"SUBSET_TABLE":        SubsettingRuleProcessingChainObjectSummaryPropagationImpactSubsetTable,
	"PARENT_CHILD_SUBSET": SubsettingRuleProcessingChainObjectSummaryPropagationImpactParentChildSubset,
	"CHILD_PARENT_SUBSET": SubsettingRuleProcessingChainObjectSummaryPropagationImpactChildParentSubset,
	"TRUNCATE":            SubsettingRuleProcessingChainObjectSummaryPropagationImpactTruncate,
	"KEEP":                SubsettingRuleProcessingChainObjectSummaryPropagationImpactKeep,
}

var mappingSubsettingRuleProcessingChainObjectSummaryPropagationImpactEnumLowerCase = map[string]SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum{
	"subset_table":        SubsettingRuleProcessingChainObjectSummaryPropagationImpactSubsetTable,
	"parent_child_subset": SubsettingRuleProcessingChainObjectSummaryPropagationImpactParentChildSubset,
	"child_parent_subset": SubsettingRuleProcessingChainObjectSummaryPropagationImpactChildParentSubset,
	"truncate":            SubsettingRuleProcessingChainObjectSummaryPropagationImpactTruncate,
	"keep":                SubsettingRuleProcessingChainObjectSummaryPropagationImpactKeep,
}

// GetSubsettingRuleProcessingChainObjectSummaryPropagationImpactEnumValues Enumerates the set of values for SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum
func GetSubsettingRuleProcessingChainObjectSummaryPropagationImpactEnumValues() []SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum {
	values := make([]SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum, 0)
	for _, v := range mappingSubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingRuleProcessingChainObjectSummaryPropagationImpactEnumStringValues Enumerates the set of values in String for SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum
func GetSubsettingRuleProcessingChainObjectSummaryPropagationImpactEnumStringValues() []string {
	return []string{
		"SUBSET_TABLE",
		"PARENT_CHILD_SUBSET",
		"CHILD_PARENT_SUBSET",
		"TRUNCATE",
		"KEEP",
	}
}

// GetMappingSubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum(val string) (SubsettingRuleProcessingChainObjectSummaryPropagationImpactEnum, bool) {
	enum, ok := mappingSubsettingRuleProcessingChainObjectSummaryPropagationImpactEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
