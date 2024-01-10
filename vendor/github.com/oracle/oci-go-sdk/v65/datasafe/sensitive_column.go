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

// SensitiveColumn A sensitive column is a resource corresponding to a database column that is considered sensitive.
// It's a subresource of sensitive data model resource and is always associated with a sensitive data model.
// Note that referential relationships are also managed as part of sensitive columns.
type SensitiveColumn struct {

	// The unique key that identifies the sensitive column. It's numeric and unique within a sensitive data model.
	Key *string `mandatory:"true" json:"key"`

	// The OCID of the sensitive data model that contains the sensitive column.
	SensitiveDataModelId *string `mandatory:"true" json:"sensitiveDataModelId"`

	// The date and time, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339),
	// the sensitive column was created in the sensitive data model.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339),
	// the sensitive column was last updated in the sensitive data model.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the sensitive column.
	LifecycleState SensitiveColumnLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The name of the application associated with the sensitive column. It's useful when the application name is
	// different from the schema name. Otherwise, it can be ignored.
	AppName *string `mandatory:"true" json:"appName"`

	// The database schema that contains the sensitive column.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The database object that contains the sensitive column.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The name of the sensitive column.
	ColumnName *string `mandatory:"true" json:"columnName"`

	// The type of the database object that contains the sensitive column.
	ObjectType SensitiveColumnObjectTypeEnum `mandatory:"true" json:"objectType"`

	// The data type of the sensitive column.
	DataType *string `mandatory:"true" json:"dataType"`

	// The status of the sensitive column. VALID means the column is considered sensitive. INVALID means the column
	// is not considered sensitive. Tracking invalid columns in a sensitive data model helps ensure that an
	// incremental data discovery job does not identify these columns as sensitive again.
	Status SensitiveColumnStatusEnum `mandatory:"true" json:"status"`

	// The source of the sensitive column. DISCOVERY indicates that the column was added to the sensitive data model
	// using a data discovery job. MANUAL indicates that the column was added manually.
	Source SensitiveColumnSourceEnum `mandatory:"true" json:"source"`

	// The type of referential relationship the sensitive column has with its parent. NONE indicates that the
	// sensitive column does not have a parent. DB_DEFINED indicates that the relationship is defined in the database
	// dictionary. APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary.
	RelationType SensitiveColumnRelationTypeEnum `mandatory:"true" json:"relationType"`

	// The estimated number of data values the column has in the associated database.
	EstimatedDataValueCount *int64 `mandatory:"true" json:"estimatedDataValueCount"`

	// Details about the current state of the sensitive column.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

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

	// The composite key groups to which the sensitive column belongs. If the column is part of a composite key,
	// it's assigned a column group. It helps identify and manage referential relationships that involve composite keys.
	ColumnGroups []string `mandatory:"false" json:"columnGroups"`
}

func (m SensitiveColumn) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SensitiveColumn) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSensitiveColumnLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSensitiveColumnLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSensitiveColumnObjectTypeEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetSensitiveColumnObjectTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSensitiveColumnStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSensitiveColumnStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSensitiveColumnSourceEnum(string(m.Source)); !ok && m.Source != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Source: %s. Supported values are: %s.", m.Source, strings.Join(GetSensitiveColumnSourceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSensitiveColumnRelationTypeEnum(string(m.RelationType)); !ok && m.RelationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", m.RelationType, strings.Join(GetSensitiveColumnRelationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SensitiveColumnObjectTypeEnum Enum with underlying type: string
type SensitiveColumnObjectTypeEnum string

// Set of constants representing the allowable values for SensitiveColumnObjectTypeEnum
const (
	SensitiveColumnObjectTypeTable          SensitiveColumnObjectTypeEnum = "TABLE"
	SensitiveColumnObjectTypeEditioningView SensitiveColumnObjectTypeEnum = "EDITIONING_VIEW"
)

var mappingSensitiveColumnObjectTypeEnum = map[string]SensitiveColumnObjectTypeEnum{
	"TABLE":           SensitiveColumnObjectTypeTable,
	"EDITIONING_VIEW": SensitiveColumnObjectTypeEditioningView,
}

var mappingSensitiveColumnObjectTypeEnumLowerCase = map[string]SensitiveColumnObjectTypeEnum{
	"table":           SensitiveColumnObjectTypeTable,
	"editioning_view": SensitiveColumnObjectTypeEditioningView,
}

// GetSensitiveColumnObjectTypeEnumValues Enumerates the set of values for SensitiveColumnObjectTypeEnum
func GetSensitiveColumnObjectTypeEnumValues() []SensitiveColumnObjectTypeEnum {
	values := make([]SensitiveColumnObjectTypeEnum, 0)
	for _, v := range mappingSensitiveColumnObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSensitiveColumnObjectTypeEnumStringValues Enumerates the set of values in String for SensitiveColumnObjectTypeEnum
func GetSensitiveColumnObjectTypeEnumStringValues() []string {
	return []string{
		"TABLE",
		"EDITIONING_VIEW",
	}
}

// GetMappingSensitiveColumnObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSensitiveColumnObjectTypeEnum(val string) (SensitiveColumnObjectTypeEnum, bool) {
	enum, ok := mappingSensitiveColumnObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SensitiveColumnStatusEnum Enum with underlying type: string
type SensitiveColumnStatusEnum string

// Set of constants representing the allowable values for SensitiveColumnStatusEnum
const (
	SensitiveColumnStatusValid   SensitiveColumnStatusEnum = "VALID"
	SensitiveColumnStatusInvalid SensitiveColumnStatusEnum = "INVALID"
)

var mappingSensitiveColumnStatusEnum = map[string]SensitiveColumnStatusEnum{
	"VALID":   SensitiveColumnStatusValid,
	"INVALID": SensitiveColumnStatusInvalid,
}

var mappingSensitiveColumnStatusEnumLowerCase = map[string]SensitiveColumnStatusEnum{
	"valid":   SensitiveColumnStatusValid,
	"invalid": SensitiveColumnStatusInvalid,
}

// GetSensitiveColumnStatusEnumValues Enumerates the set of values for SensitiveColumnStatusEnum
func GetSensitiveColumnStatusEnumValues() []SensitiveColumnStatusEnum {
	values := make([]SensitiveColumnStatusEnum, 0)
	for _, v := range mappingSensitiveColumnStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSensitiveColumnStatusEnumStringValues Enumerates the set of values in String for SensitiveColumnStatusEnum
func GetSensitiveColumnStatusEnumStringValues() []string {
	return []string{
		"VALID",
		"INVALID",
	}
}

// GetMappingSensitiveColumnStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSensitiveColumnStatusEnum(val string) (SensitiveColumnStatusEnum, bool) {
	enum, ok := mappingSensitiveColumnStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SensitiveColumnSourceEnum Enum with underlying type: string
type SensitiveColumnSourceEnum string

// Set of constants representing the allowable values for SensitiveColumnSourceEnum
const (
	SensitiveColumnSourceManual    SensitiveColumnSourceEnum = "MANUAL"
	SensitiveColumnSourceDiscovery SensitiveColumnSourceEnum = "DISCOVERY"
)

var mappingSensitiveColumnSourceEnum = map[string]SensitiveColumnSourceEnum{
	"MANUAL":    SensitiveColumnSourceManual,
	"DISCOVERY": SensitiveColumnSourceDiscovery,
}

var mappingSensitiveColumnSourceEnumLowerCase = map[string]SensitiveColumnSourceEnum{
	"manual":    SensitiveColumnSourceManual,
	"discovery": SensitiveColumnSourceDiscovery,
}

// GetSensitiveColumnSourceEnumValues Enumerates the set of values for SensitiveColumnSourceEnum
func GetSensitiveColumnSourceEnumValues() []SensitiveColumnSourceEnum {
	values := make([]SensitiveColumnSourceEnum, 0)
	for _, v := range mappingSensitiveColumnSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetSensitiveColumnSourceEnumStringValues Enumerates the set of values in String for SensitiveColumnSourceEnum
func GetSensitiveColumnSourceEnumStringValues() []string {
	return []string{
		"MANUAL",
		"DISCOVERY",
	}
}

// GetMappingSensitiveColumnSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSensitiveColumnSourceEnum(val string) (SensitiveColumnSourceEnum, bool) {
	enum, ok := mappingSensitiveColumnSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SensitiveColumnRelationTypeEnum Enum with underlying type: string
type SensitiveColumnRelationTypeEnum string

// Set of constants representing the allowable values for SensitiveColumnRelationTypeEnum
const (
	SensitiveColumnRelationTypeNone       SensitiveColumnRelationTypeEnum = "NONE"
	SensitiveColumnRelationTypeAppDefined SensitiveColumnRelationTypeEnum = "APP_DEFINED"
	SensitiveColumnRelationTypeDbDefined  SensitiveColumnRelationTypeEnum = "DB_DEFINED"
)

var mappingSensitiveColumnRelationTypeEnum = map[string]SensitiveColumnRelationTypeEnum{
	"NONE":        SensitiveColumnRelationTypeNone,
	"APP_DEFINED": SensitiveColumnRelationTypeAppDefined,
	"DB_DEFINED":  SensitiveColumnRelationTypeDbDefined,
}

var mappingSensitiveColumnRelationTypeEnumLowerCase = map[string]SensitiveColumnRelationTypeEnum{
	"none":        SensitiveColumnRelationTypeNone,
	"app_defined": SensitiveColumnRelationTypeAppDefined,
	"db_defined":  SensitiveColumnRelationTypeDbDefined,
}

// GetSensitiveColumnRelationTypeEnumValues Enumerates the set of values for SensitiveColumnRelationTypeEnum
func GetSensitiveColumnRelationTypeEnumValues() []SensitiveColumnRelationTypeEnum {
	values := make([]SensitiveColumnRelationTypeEnum, 0)
	for _, v := range mappingSensitiveColumnRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSensitiveColumnRelationTypeEnumStringValues Enumerates the set of values in String for SensitiveColumnRelationTypeEnum
func GetSensitiveColumnRelationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingSensitiveColumnRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSensitiveColumnRelationTypeEnum(val string) (SensitiveColumnRelationTypeEnum, bool) {
	enum, ok := mappingSensitiveColumnRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
