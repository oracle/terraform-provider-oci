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

// CopyObjectMetadataSummary Details of copied objects.
type CopyObjectMetadataSummary struct {

	// Old key of the object from where the object was copied. For example a dataflow key within the project being copied.
	OldKey *string `mandatory:"false" json:"oldKey"`

	// New key of the object to identify the copied object. For example the new dataflow key.
	NewKey *string `mandatory:"false" json:"newKey"`

	// Name of the object.
	Name *string `mandatory:"false" json:"name"`

	// Object identifier.
	Identifier *string `mandatory:"false" json:"identifier"`

	// Object type.
	ObjectType *string `mandatory:"false" json:"objectType"`

	// Object version.
	ObjectVersion *string `mandatory:"false" json:"objectVersion"`

	// Aggregator key
	AggregatorKey *string `mandatory:"false" json:"aggregatorKey"`

	// Object name path.
	NamePath *string `mandatory:"false" json:"namePath"`

	// time at which this object was last updated.
	TimeUpdatedInMillis *int64 `mandatory:"false" json:"timeUpdatedInMillis"`

	// Object resolution action.
	ResolutionAction CopyObjectMetadataSummaryResolutionActionEnum `mandatory:"false" json:"resolutionAction,omitempty"`
}

func (m CopyObjectMetadataSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CopyObjectMetadataSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCopyObjectMetadataSummaryResolutionActionEnum(string(m.ResolutionAction)); !ok && m.ResolutionAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResolutionAction: %s. Supported values are: %s.", m.ResolutionAction, strings.Join(GetCopyObjectMetadataSummaryResolutionActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CopyObjectMetadataSummaryResolutionActionEnum Enum with underlying type: string
type CopyObjectMetadataSummaryResolutionActionEnum string

// Set of constants representing the allowable values for CopyObjectMetadataSummaryResolutionActionEnum
const (
	CopyObjectMetadataSummaryResolutionActionCreated    CopyObjectMetadataSummaryResolutionActionEnum = "CREATED"
	CopyObjectMetadataSummaryResolutionActionRetained   CopyObjectMetadataSummaryResolutionActionEnum = "RETAINED"
	CopyObjectMetadataSummaryResolutionActionDuplicated CopyObjectMetadataSummaryResolutionActionEnum = "DUPLICATED"
	CopyObjectMetadataSummaryResolutionActionReplaced   CopyObjectMetadataSummaryResolutionActionEnum = "REPLACED"
)

var mappingCopyObjectMetadataSummaryResolutionActionEnum = map[string]CopyObjectMetadataSummaryResolutionActionEnum{
	"CREATED":    CopyObjectMetadataSummaryResolutionActionCreated,
	"RETAINED":   CopyObjectMetadataSummaryResolutionActionRetained,
	"DUPLICATED": CopyObjectMetadataSummaryResolutionActionDuplicated,
	"REPLACED":   CopyObjectMetadataSummaryResolutionActionReplaced,
}

var mappingCopyObjectMetadataSummaryResolutionActionEnumLowerCase = map[string]CopyObjectMetadataSummaryResolutionActionEnum{
	"created":    CopyObjectMetadataSummaryResolutionActionCreated,
	"retained":   CopyObjectMetadataSummaryResolutionActionRetained,
	"duplicated": CopyObjectMetadataSummaryResolutionActionDuplicated,
	"replaced":   CopyObjectMetadataSummaryResolutionActionReplaced,
}

// GetCopyObjectMetadataSummaryResolutionActionEnumValues Enumerates the set of values for CopyObjectMetadataSummaryResolutionActionEnum
func GetCopyObjectMetadataSummaryResolutionActionEnumValues() []CopyObjectMetadataSummaryResolutionActionEnum {
	values := make([]CopyObjectMetadataSummaryResolutionActionEnum, 0)
	for _, v := range mappingCopyObjectMetadataSummaryResolutionActionEnum {
		values = append(values, v)
	}
	return values
}

// GetCopyObjectMetadataSummaryResolutionActionEnumStringValues Enumerates the set of values in String for CopyObjectMetadataSummaryResolutionActionEnum
func GetCopyObjectMetadataSummaryResolutionActionEnumStringValues() []string {
	return []string{
		"CREATED",
		"RETAINED",
		"DUPLICATED",
		"REPLACED",
	}
}

// GetMappingCopyObjectMetadataSummaryResolutionActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCopyObjectMetadataSummaryResolutionActionEnum(val string) (CopyObjectMetadataSummaryResolutionActionEnum, bool) {
	enum, ok := mappingCopyObjectMetadataSummaryResolutionActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
