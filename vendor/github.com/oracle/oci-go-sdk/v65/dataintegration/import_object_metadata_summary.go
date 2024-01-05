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

// ImportObjectMetadataSummary Details of the objects to imported.
type ImportObjectMetadataSummary struct {

	// Old key of the object
	OldKey *string `mandatory:"false" json:"oldKey"`

	// New key of the object
	NewKey *string `mandatory:"false" json:"newKey"`

	// Name of the object
	Name *string `mandatory:"false" json:"name"`

	// Object identifier
	Identifier *string `mandatory:"false" json:"identifier"`

	// Object type
	ObjectType *string `mandatory:"false" json:"objectType"`

	// Object version
	ObjectVersion *string `mandatory:"false" json:"objectVersion"`

	// Aggregator key
	AggregatorKey *string `mandatory:"false" json:"aggregatorKey"`

	// Object name path
	NamePath *string `mandatory:"false" json:"namePath"`

	// time at which this object was last updated.
	TimeUpdatedInMillis *int64 `mandatory:"false" json:"timeUpdatedInMillis"`

	// Object resolution action
	ResolutionAction ImportObjectMetadataSummaryResolutionActionEnum `mandatory:"false" json:"resolutionAction,omitempty"`
}

func (m ImportObjectMetadataSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportObjectMetadataSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingImportObjectMetadataSummaryResolutionActionEnum(string(m.ResolutionAction)); !ok && m.ResolutionAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResolutionAction: %s. Supported values are: %s.", m.ResolutionAction, strings.Join(GetImportObjectMetadataSummaryResolutionActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImportObjectMetadataSummaryResolutionActionEnum Enum with underlying type: string
type ImportObjectMetadataSummaryResolutionActionEnum string

// Set of constants representing the allowable values for ImportObjectMetadataSummaryResolutionActionEnum
const (
	ImportObjectMetadataSummaryResolutionActionCreated    ImportObjectMetadataSummaryResolutionActionEnum = "CREATED"
	ImportObjectMetadataSummaryResolutionActionRetained   ImportObjectMetadataSummaryResolutionActionEnum = "RETAINED"
	ImportObjectMetadataSummaryResolutionActionDuplicated ImportObjectMetadataSummaryResolutionActionEnum = "DUPLICATED"
	ImportObjectMetadataSummaryResolutionActionReplaced   ImportObjectMetadataSummaryResolutionActionEnum = "REPLACED"
)

var mappingImportObjectMetadataSummaryResolutionActionEnum = map[string]ImportObjectMetadataSummaryResolutionActionEnum{
	"CREATED":    ImportObjectMetadataSummaryResolutionActionCreated,
	"RETAINED":   ImportObjectMetadataSummaryResolutionActionRetained,
	"DUPLICATED": ImportObjectMetadataSummaryResolutionActionDuplicated,
	"REPLACED":   ImportObjectMetadataSummaryResolutionActionReplaced,
}

var mappingImportObjectMetadataSummaryResolutionActionEnumLowerCase = map[string]ImportObjectMetadataSummaryResolutionActionEnum{
	"created":    ImportObjectMetadataSummaryResolutionActionCreated,
	"retained":   ImportObjectMetadataSummaryResolutionActionRetained,
	"duplicated": ImportObjectMetadataSummaryResolutionActionDuplicated,
	"replaced":   ImportObjectMetadataSummaryResolutionActionReplaced,
}

// GetImportObjectMetadataSummaryResolutionActionEnumValues Enumerates the set of values for ImportObjectMetadataSummaryResolutionActionEnum
func GetImportObjectMetadataSummaryResolutionActionEnumValues() []ImportObjectMetadataSummaryResolutionActionEnum {
	values := make([]ImportObjectMetadataSummaryResolutionActionEnum, 0)
	for _, v := range mappingImportObjectMetadataSummaryResolutionActionEnum {
		values = append(values, v)
	}
	return values
}

// GetImportObjectMetadataSummaryResolutionActionEnumStringValues Enumerates the set of values in String for ImportObjectMetadataSummaryResolutionActionEnum
func GetImportObjectMetadataSummaryResolutionActionEnumStringValues() []string {
	return []string{
		"CREATED",
		"RETAINED",
		"DUPLICATED",
		"REPLACED",
	}
}

// GetMappingImportObjectMetadataSummaryResolutionActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImportObjectMetadataSummaryResolutionActionEnum(val string) (ImportObjectMetadataSummaryResolutionActionEnum, bool) {
	enum, ok := mappingImportObjectMetadataSummaryResolutionActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
