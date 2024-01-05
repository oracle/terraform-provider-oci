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

// DiscoveryJobResult A discovery job result representing a sensitive column. It can be one of the following three types:
// NEW: A new sensitive column in the target database that is not in the sensitive data model.
// DELETED: A column that is present in the sensitive data model but has been deleted from the target database.
// MODIFIED: A column that is present in the target database as well as the sensitive data model but some of its attributes have been modified.
type DiscoveryJobResult struct {

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
	ObjectType DiscoveryJobResultObjectTypeEnum `mandatory:"true" json:"objectType"`

	// The data type of the sensitive column.
	DataType *string `mandatory:"true" json:"dataType"`

	// The type of referential relationship the sensitive column has with its parent. NONE indicates that the sensitive
	// column does not have a parent. DB_DEFINED indicates that the relationship is defined in the database dictionary.
	// APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary.
	RelationType DiscoveryJobResultRelationTypeEnum `mandatory:"true" json:"relationType"`

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

	// The name of the application. An application is an entity that is identified by a schema and stores sensitive information for that schema. Its value will be same as schemaName, if no value is passed.
	AppName *string `mandatory:"false" json:"appName"`

	// The OCID of the sensitive type associated with the sensitive column.
	SensitiveTypeId *string `mandatory:"false" json:"sensitiveTypeId"`

	// Unique keys identifying the columns that are parents of the sensitive column. At present, it tracks a single parent only.
	ParentColumnKeys []string `mandatory:"false" json:"parentColumnKeys"`

	// Original data values collected for the sensitive column from the associated database. Sample data helps review
	// the column and ensure that it actually contains sensitive data. Note that sample data is retrieved by a data
	// discovery job only if the isSampleDataCollectionEnabled attribute is set to true. At present, only one data
	// value is collected per sensitive column.
	SampleDataValues []string `mandatory:"false" json:"sampleDataValues"`

	// Unique keys identifying the columns that are application-level (non-dictionary) children of the sensitive column.
	AppDefinedChildColumnKeys []string `mandatory:"false" json:"appDefinedChildColumnKeys"`

	// Unique keys identifying the columns that are database-level (dictionary-defined) children of the sensitive column.
	DbDefinedChildColumnKeys []string `mandatory:"false" json:"dbDefinedChildColumnKeys"`

	ModifiedAttributes *ModifiedAttributes `mandatory:"false" json:"modifiedAttributes"`
}

func (m DiscoveryJobResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveryJobResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveryJobResultDiscoveryTypeEnum(string(m.DiscoveryType)); !ok && m.DiscoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoveryType: %s. Supported values are: %s.", m.DiscoveryType, strings.Join(GetDiscoveryJobResultDiscoveryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDiscoveryJobResultObjectTypeEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetDiscoveryJobResultObjectTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDiscoveryJobResultRelationTypeEnum(string(m.RelationType)); !ok && m.RelationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", m.RelationType, strings.Join(GetDiscoveryJobResultRelationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDiscoveryJobResultPlannedActionEnum(string(m.PlannedAction)); !ok && m.PlannedAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlannedAction: %s. Supported values are: %s.", m.PlannedAction, strings.Join(GetDiscoveryJobResultPlannedActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiscoveryJobResultDiscoveryTypeEnum Enum with underlying type: string
type DiscoveryJobResultDiscoveryTypeEnum string

// Set of constants representing the allowable values for DiscoveryJobResultDiscoveryTypeEnum
const (
	DiscoveryJobResultDiscoveryTypeNew      DiscoveryJobResultDiscoveryTypeEnum = "NEW"
	DiscoveryJobResultDiscoveryTypeModified DiscoveryJobResultDiscoveryTypeEnum = "MODIFIED"
	DiscoveryJobResultDiscoveryTypeDeleted  DiscoveryJobResultDiscoveryTypeEnum = "DELETED"
)

var mappingDiscoveryJobResultDiscoveryTypeEnum = map[string]DiscoveryJobResultDiscoveryTypeEnum{
	"NEW":      DiscoveryJobResultDiscoveryTypeNew,
	"MODIFIED": DiscoveryJobResultDiscoveryTypeModified,
	"DELETED":  DiscoveryJobResultDiscoveryTypeDeleted,
}

var mappingDiscoveryJobResultDiscoveryTypeEnumLowerCase = map[string]DiscoveryJobResultDiscoveryTypeEnum{
	"new":      DiscoveryJobResultDiscoveryTypeNew,
	"modified": DiscoveryJobResultDiscoveryTypeModified,
	"deleted":  DiscoveryJobResultDiscoveryTypeDeleted,
}

// GetDiscoveryJobResultDiscoveryTypeEnumValues Enumerates the set of values for DiscoveryJobResultDiscoveryTypeEnum
func GetDiscoveryJobResultDiscoveryTypeEnumValues() []DiscoveryJobResultDiscoveryTypeEnum {
	values := make([]DiscoveryJobResultDiscoveryTypeEnum, 0)
	for _, v := range mappingDiscoveryJobResultDiscoveryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryJobResultDiscoveryTypeEnumStringValues Enumerates the set of values in String for DiscoveryJobResultDiscoveryTypeEnum
func GetDiscoveryJobResultDiscoveryTypeEnumStringValues() []string {
	return []string{
		"NEW",
		"MODIFIED",
		"DELETED",
	}
}

// GetMappingDiscoveryJobResultDiscoveryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryJobResultDiscoveryTypeEnum(val string) (DiscoveryJobResultDiscoveryTypeEnum, bool) {
	enum, ok := mappingDiscoveryJobResultDiscoveryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DiscoveryJobResultObjectTypeEnum Enum with underlying type: string
type DiscoveryJobResultObjectTypeEnum string

// Set of constants representing the allowable values for DiscoveryJobResultObjectTypeEnum
const (
	DiscoveryJobResultObjectTypeTable          DiscoveryJobResultObjectTypeEnum = "TABLE"
	DiscoveryJobResultObjectTypeEditioningView DiscoveryJobResultObjectTypeEnum = "EDITIONING_VIEW"
)

var mappingDiscoveryJobResultObjectTypeEnum = map[string]DiscoveryJobResultObjectTypeEnum{
	"TABLE":           DiscoveryJobResultObjectTypeTable,
	"EDITIONING_VIEW": DiscoveryJobResultObjectTypeEditioningView,
}

var mappingDiscoveryJobResultObjectTypeEnumLowerCase = map[string]DiscoveryJobResultObjectTypeEnum{
	"table":           DiscoveryJobResultObjectTypeTable,
	"editioning_view": DiscoveryJobResultObjectTypeEditioningView,
}

// GetDiscoveryJobResultObjectTypeEnumValues Enumerates the set of values for DiscoveryJobResultObjectTypeEnum
func GetDiscoveryJobResultObjectTypeEnumValues() []DiscoveryJobResultObjectTypeEnum {
	values := make([]DiscoveryJobResultObjectTypeEnum, 0)
	for _, v := range mappingDiscoveryJobResultObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryJobResultObjectTypeEnumStringValues Enumerates the set of values in String for DiscoveryJobResultObjectTypeEnum
func GetDiscoveryJobResultObjectTypeEnumStringValues() []string {
	return []string{
		"TABLE",
		"EDITIONING_VIEW",
	}
}

// GetMappingDiscoveryJobResultObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryJobResultObjectTypeEnum(val string) (DiscoveryJobResultObjectTypeEnum, bool) {
	enum, ok := mappingDiscoveryJobResultObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DiscoveryJobResultRelationTypeEnum Enum with underlying type: string
type DiscoveryJobResultRelationTypeEnum string

// Set of constants representing the allowable values for DiscoveryJobResultRelationTypeEnum
const (
	DiscoveryJobResultRelationTypeNone       DiscoveryJobResultRelationTypeEnum = "NONE"
	DiscoveryJobResultRelationTypeAppDefined DiscoveryJobResultRelationTypeEnum = "APP_DEFINED"
	DiscoveryJobResultRelationTypeDbDefined  DiscoveryJobResultRelationTypeEnum = "DB_DEFINED"
)

var mappingDiscoveryJobResultRelationTypeEnum = map[string]DiscoveryJobResultRelationTypeEnum{
	"NONE":        DiscoveryJobResultRelationTypeNone,
	"APP_DEFINED": DiscoveryJobResultRelationTypeAppDefined,
	"DB_DEFINED":  DiscoveryJobResultRelationTypeDbDefined,
}

var mappingDiscoveryJobResultRelationTypeEnumLowerCase = map[string]DiscoveryJobResultRelationTypeEnum{
	"none":        DiscoveryJobResultRelationTypeNone,
	"app_defined": DiscoveryJobResultRelationTypeAppDefined,
	"db_defined":  DiscoveryJobResultRelationTypeDbDefined,
}

// GetDiscoveryJobResultRelationTypeEnumValues Enumerates the set of values for DiscoveryJobResultRelationTypeEnum
func GetDiscoveryJobResultRelationTypeEnumValues() []DiscoveryJobResultRelationTypeEnum {
	values := make([]DiscoveryJobResultRelationTypeEnum, 0)
	for _, v := range mappingDiscoveryJobResultRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryJobResultRelationTypeEnumStringValues Enumerates the set of values in String for DiscoveryJobResultRelationTypeEnum
func GetDiscoveryJobResultRelationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingDiscoveryJobResultRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryJobResultRelationTypeEnum(val string) (DiscoveryJobResultRelationTypeEnum, bool) {
	enum, ok := mappingDiscoveryJobResultRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DiscoveryJobResultPlannedActionEnum Enum with underlying type: string
type DiscoveryJobResultPlannedActionEnum string

// Set of constants representing the allowable values for DiscoveryJobResultPlannedActionEnum
const (
	DiscoveryJobResultPlannedActionNone       DiscoveryJobResultPlannedActionEnum = "NONE"
	DiscoveryJobResultPlannedActionAccept     DiscoveryJobResultPlannedActionEnum = "ACCEPT"
	DiscoveryJobResultPlannedActionInvalidate DiscoveryJobResultPlannedActionEnum = "INVALIDATE"
	DiscoveryJobResultPlannedActionReject     DiscoveryJobResultPlannedActionEnum = "REJECT"
)

var mappingDiscoveryJobResultPlannedActionEnum = map[string]DiscoveryJobResultPlannedActionEnum{
	"NONE":       DiscoveryJobResultPlannedActionNone,
	"ACCEPT":     DiscoveryJobResultPlannedActionAccept,
	"INVALIDATE": DiscoveryJobResultPlannedActionInvalidate,
	"REJECT":     DiscoveryJobResultPlannedActionReject,
}

var mappingDiscoveryJobResultPlannedActionEnumLowerCase = map[string]DiscoveryJobResultPlannedActionEnum{
	"none":       DiscoveryJobResultPlannedActionNone,
	"accept":     DiscoveryJobResultPlannedActionAccept,
	"invalidate": DiscoveryJobResultPlannedActionInvalidate,
	"reject":     DiscoveryJobResultPlannedActionReject,
}

// GetDiscoveryJobResultPlannedActionEnumValues Enumerates the set of values for DiscoveryJobResultPlannedActionEnum
func GetDiscoveryJobResultPlannedActionEnumValues() []DiscoveryJobResultPlannedActionEnum {
	values := make([]DiscoveryJobResultPlannedActionEnum, 0)
	for _, v := range mappingDiscoveryJobResultPlannedActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryJobResultPlannedActionEnumStringValues Enumerates the set of values in String for DiscoveryJobResultPlannedActionEnum
func GetDiscoveryJobResultPlannedActionEnumStringValues() []string {
	return []string{
		"NONE",
		"ACCEPT",
		"INVALIDATE",
		"REJECT",
	}
}

// GetMappingDiscoveryJobResultPlannedActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryJobResultPlannedActionEnum(val string) (DiscoveryJobResultPlannedActionEnum, bool) {
	enum, ok := mappingDiscoveryJobResultPlannedActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
