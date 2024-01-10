// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DiscoveryJobResultSummary Summary of a discovery job result.
type DiscoveryJobResultSummary struct {

	// The unique key that identifies the discovery result.
	Key *string `mandatory:"true" json:"key"`

	// The type of the discovery result. It can be one of the following three types:
	// NEW: A new sensitive column in the target database that is not in the sensitive data model.
	// DELETED: A column that is present in the sensitive data model but has been deleted from the target database.
	// MODIFIED: A column that is present in the target database as well as the sensitive data model but some of its attributes have been modified.
	DiscoveryType DiscoveryJobResultDiscoveryTypeEnum `mandatory:"true" json:"discoveryType"`

	// The database schema that contains the sensitive column.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The database object that contains the sensitive column.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The name of the sensitive column.
	ColumnName *string `mandatory:"true" json:"columnName"`

	// The type of the database object that contains the sensitive column.
	ObjectType DiscoveryJobResultSummaryObjectTypeEnum `mandatory:"true" json:"objectType"`

	// The data type of the sensitive column.
	DataType *string `mandatory:"true" json:"dataType"`

	// The type of referential relationship the sensitive column has with its parent. NONE indicates that the sensitive
	// column does not have a parent. DB_DEFINED indicates that the relationship is defined in the database dictionary.
	// APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary.
	RelationType DiscoveryJobResultSummaryRelationTypeEnum `mandatory:"true" json:"relationType"`

	// The estimated number of data values the column has in the associated database.
	EstimatedDataValueCount *int64 `mandatory:"true" json:"estimatedDataValueCount"`

	// Specifies how to process the discovery result. It's set to NONE by default. Use the PatchDiscoveryJobResults operation to update this attribute. You can choose one of the following options:
	// ACCEPT: To accept the discovery result and update the sensitive data model to reflect the changes.
	// REJECT: To reject the discovery result so that it doesn't change the sensitive data model.
	// INVALIDATE: To invalidate a newly discovered column. It adds the column to the sensitive data model but marks it as invalid. It helps track false positives and ensure that they aren't reported by future discovery jobs.
	// After specifying the planned action, you can use the ApplyDiscoveryJobResults operation to automatically process the discovery results.
	PlannedAction DiscoveryJobResultPlannedActionEnum `mandatory:"true" json:"plannedAction"`

	// Indicates whether the discovery result has been processed. You can update this attribute using the PatchDiscoveryJobResults
	// operation to track whether the discovery result has already been processed and applied to the sensitive data model.
	IsResultApplied *bool `mandatory:"true" json:"isResultApplied"`

	// The OCID of the discovery job.
	DiscoveryJobId *string `mandatory:"true" json:"discoveryJobId"`

	// The unique key that identifies the sensitive column represented by the discovery result.
	SensitiveColumnkey *string `mandatory:"false" json:"sensitiveColumnkey"`

	// The OCID of the sensitive type associated with the sensitive column.
	SensitiveTypeId *string `mandatory:"false" json:"sensitiveTypeId"`

	// Unique keys identifying the columns that are parents of the sensitive column. At present, it tracks a single parent only.
	ParentColumnKeys []string `mandatory:"false" json:"parentColumnKeys"`

	// Original data values collected for the sensitive column from the associated database. Sample data helps review
	// the column and ensure that it actually contains sensitive data. Note that sample data is retrieved by a data
	// discovery job only if the isSampleDataCollectionEnabled attribute is set to true. At present, only one data
	// value is collected per sensitive column.
	SampleDataValues []string `mandatory:"false" json:"sampleDataValues"`
}

func (m DiscoveryJobResultSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveryJobResultSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveryJobResultDiscoveryTypeEnum(string(m.DiscoveryType)); !ok && m.DiscoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoveryType: %s. Supported values are: %s.", m.DiscoveryType, strings.Join(GetDiscoveryJobResultDiscoveryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDiscoveryJobResultSummaryObjectTypeEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetDiscoveryJobResultSummaryObjectTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDiscoveryJobResultSummaryRelationTypeEnum(string(m.RelationType)); !ok && m.RelationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", m.RelationType, strings.Join(GetDiscoveryJobResultSummaryRelationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDiscoveryJobResultPlannedActionEnum(string(m.PlannedAction)); !ok && m.PlannedAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlannedAction: %s. Supported values are: %s.", m.PlannedAction, strings.Join(GetDiscoveryJobResultPlannedActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiscoveryJobResultSummaryObjectTypeEnum Enum with underlying type: string
type DiscoveryJobResultSummaryObjectTypeEnum string

// Set of constants representing the allowable values for DiscoveryJobResultSummaryObjectTypeEnum
const (
	DiscoveryJobResultSummaryObjectTypeTable          DiscoveryJobResultSummaryObjectTypeEnum = "TABLE"
	DiscoveryJobResultSummaryObjectTypeEditioningView DiscoveryJobResultSummaryObjectTypeEnum = "EDITIONING_VIEW"
)

var mappingDiscoveryJobResultSummaryObjectTypeEnum = map[string]DiscoveryJobResultSummaryObjectTypeEnum{
	"TABLE":           DiscoveryJobResultSummaryObjectTypeTable,
	"EDITIONING_VIEW": DiscoveryJobResultSummaryObjectTypeEditioningView,
}

var mappingDiscoveryJobResultSummaryObjectTypeEnumLowerCase = map[string]DiscoveryJobResultSummaryObjectTypeEnum{
	"table":           DiscoveryJobResultSummaryObjectTypeTable,
	"editioning_view": DiscoveryJobResultSummaryObjectTypeEditioningView,
}

// GetDiscoveryJobResultSummaryObjectTypeEnumValues Enumerates the set of values for DiscoveryJobResultSummaryObjectTypeEnum
func GetDiscoveryJobResultSummaryObjectTypeEnumValues() []DiscoveryJobResultSummaryObjectTypeEnum {
	values := make([]DiscoveryJobResultSummaryObjectTypeEnum, 0)
	for _, v := range mappingDiscoveryJobResultSummaryObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryJobResultSummaryObjectTypeEnumStringValues Enumerates the set of values in String for DiscoveryJobResultSummaryObjectTypeEnum
func GetDiscoveryJobResultSummaryObjectTypeEnumStringValues() []string {
	return []string{
		"TABLE",
		"EDITIONING_VIEW",
	}
}

// GetMappingDiscoveryJobResultSummaryObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryJobResultSummaryObjectTypeEnum(val string) (DiscoveryJobResultSummaryObjectTypeEnum, bool) {
	enum, ok := mappingDiscoveryJobResultSummaryObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DiscoveryJobResultSummaryRelationTypeEnum Enum with underlying type: string
type DiscoveryJobResultSummaryRelationTypeEnum string

// Set of constants representing the allowable values for DiscoveryJobResultSummaryRelationTypeEnum
const (
	DiscoveryJobResultSummaryRelationTypeNone       DiscoveryJobResultSummaryRelationTypeEnum = "NONE"
	DiscoveryJobResultSummaryRelationTypeAppDefined DiscoveryJobResultSummaryRelationTypeEnum = "APP_DEFINED"
	DiscoveryJobResultSummaryRelationTypeDbDefined  DiscoveryJobResultSummaryRelationTypeEnum = "DB_DEFINED"
)

var mappingDiscoveryJobResultSummaryRelationTypeEnum = map[string]DiscoveryJobResultSummaryRelationTypeEnum{
	"NONE":        DiscoveryJobResultSummaryRelationTypeNone,
	"APP_DEFINED": DiscoveryJobResultSummaryRelationTypeAppDefined,
	"DB_DEFINED":  DiscoveryJobResultSummaryRelationTypeDbDefined,
}

var mappingDiscoveryJobResultSummaryRelationTypeEnumLowerCase = map[string]DiscoveryJobResultSummaryRelationTypeEnum{
	"none":        DiscoveryJobResultSummaryRelationTypeNone,
	"app_defined": DiscoveryJobResultSummaryRelationTypeAppDefined,
	"db_defined":  DiscoveryJobResultSummaryRelationTypeDbDefined,
}

// GetDiscoveryJobResultSummaryRelationTypeEnumValues Enumerates the set of values for DiscoveryJobResultSummaryRelationTypeEnum
func GetDiscoveryJobResultSummaryRelationTypeEnumValues() []DiscoveryJobResultSummaryRelationTypeEnum {
	values := make([]DiscoveryJobResultSummaryRelationTypeEnum, 0)
	for _, v := range mappingDiscoveryJobResultSummaryRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryJobResultSummaryRelationTypeEnumStringValues Enumerates the set of values in String for DiscoveryJobResultSummaryRelationTypeEnum
func GetDiscoveryJobResultSummaryRelationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingDiscoveryJobResultSummaryRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryJobResultSummaryRelationTypeEnum(val string) (DiscoveryJobResultSummaryRelationTypeEnum, bool) {
	enum, ok := mappingDiscoveryJobResultSummaryRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
