// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SensitiveColumnSummary Summary of a sensitive column present in a sensitive data model.
type SensitiveColumnSummary struct {

	// The unique key that identifies the sensitive column. It's numeric and unique within a sensitive data model.
	Key *string `mandatory:"true" json:"key"`

	// The OCID of the sensitive data model that contains the sensitive column.
	SensitiveDataModelId *string `mandatory:"true" json:"sensitiveDataModelId"`

	// The current state of the sensitive column.
	LifecycleState SensitiveColumnLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339),
	// the sensitive column was created in the sensitive data model.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339),
	// the sensitive column was last updated in the sensitive data model.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

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
	ObjectType SensitiveColumnSummaryObjectTypeEnum `mandatory:"true" json:"objectType"`

	// The data type of the sensitive column.
	DataType *string `mandatory:"true" json:"dataType"`

	// The status of the sensitive column. VALID means the column is considered sensitive. INVALID means the column
	// is not considered sensitive. Tracking invalid columns in a sensitive data model helps ensure that an
	// incremental data discovery job does not identify these columns as sensitive again.
	Status SensitiveColumnSummaryStatusEnum `mandatory:"true" json:"status"`

	// The source of the sensitive column. DISCOVERY indicates that the column was added to the sensitive data model
	// using a data discovery job. MANUAL indicates that the column was added manually.
	Source SensitiveColumnSummarySourceEnum `mandatory:"true" json:"source"`

	// The type of referential relationship the sensitive column has with its parent. NONE indicates that the
	// sensitive column does not have a parent. DB_DEFINED indicates that the relationship is defined in the database
	// dictionary. APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary.
	RelationType SensitiveColumnSummaryRelationTypeEnum `mandatory:"true" json:"relationType"`

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
}

func (m SensitiveColumnSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SensitiveColumnSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSensitiveColumnLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSensitiveColumnLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSensitiveColumnSummaryObjectTypeEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetSensitiveColumnSummaryObjectTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSensitiveColumnSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSensitiveColumnSummaryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSensitiveColumnSummarySourceEnum(string(m.Source)); !ok && m.Source != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Source: %s. Supported values are: %s.", m.Source, strings.Join(GetSensitiveColumnSummarySourceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSensitiveColumnSummaryRelationTypeEnum(string(m.RelationType)); !ok && m.RelationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", m.RelationType, strings.Join(GetSensitiveColumnSummaryRelationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SensitiveColumnSummaryObjectTypeEnum Enum with underlying type: string
type SensitiveColumnSummaryObjectTypeEnum string

// Set of constants representing the allowable values for SensitiveColumnSummaryObjectTypeEnum
const (
	SensitiveColumnSummaryObjectTypeTable          SensitiveColumnSummaryObjectTypeEnum = "TABLE"
	SensitiveColumnSummaryObjectTypeEditioningView SensitiveColumnSummaryObjectTypeEnum = "EDITIONING_VIEW"
)

var mappingSensitiveColumnSummaryObjectTypeEnum = map[string]SensitiveColumnSummaryObjectTypeEnum{
	"TABLE":           SensitiveColumnSummaryObjectTypeTable,
	"EDITIONING_VIEW": SensitiveColumnSummaryObjectTypeEditioningView,
}

// GetSensitiveColumnSummaryObjectTypeEnumValues Enumerates the set of values for SensitiveColumnSummaryObjectTypeEnum
func GetSensitiveColumnSummaryObjectTypeEnumValues() []SensitiveColumnSummaryObjectTypeEnum {
	values := make([]SensitiveColumnSummaryObjectTypeEnum, 0)
	for _, v := range mappingSensitiveColumnSummaryObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSensitiveColumnSummaryObjectTypeEnumStringValues Enumerates the set of values in String for SensitiveColumnSummaryObjectTypeEnum
func GetSensitiveColumnSummaryObjectTypeEnumStringValues() []string {
	return []string{
		"TABLE",
		"EDITIONING_VIEW",
	}
}

// GetMappingSensitiveColumnSummaryObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSensitiveColumnSummaryObjectTypeEnum(val string) (SensitiveColumnSummaryObjectTypeEnum, bool) {
	mappingSensitiveColumnSummaryObjectTypeEnumIgnoreCase := make(map[string]SensitiveColumnSummaryObjectTypeEnum)
	for k, v := range mappingSensitiveColumnSummaryObjectTypeEnum {
		mappingSensitiveColumnSummaryObjectTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSensitiveColumnSummaryObjectTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SensitiveColumnSummaryStatusEnum Enum with underlying type: string
type SensitiveColumnSummaryStatusEnum string

// Set of constants representing the allowable values for SensitiveColumnSummaryStatusEnum
const (
	SensitiveColumnSummaryStatusValid   SensitiveColumnSummaryStatusEnum = "VALID"
	SensitiveColumnSummaryStatusInvalid SensitiveColumnSummaryStatusEnum = "INVALID"
)

var mappingSensitiveColumnSummaryStatusEnum = map[string]SensitiveColumnSummaryStatusEnum{
	"VALID":   SensitiveColumnSummaryStatusValid,
	"INVALID": SensitiveColumnSummaryStatusInvalid,
}

// GetSensitiveColumnSummaryStatusEnumValues Enumerates the set of values for SensitiveColumnSummaryStatusEnum
func GetSensitiveColumnSummaryStatusEnumValues() []SensitiveColumnSummaryStatusEnum {
	values := make([]SensitiveColumnSummaryStatusEnum, 0)
	for _, v := range mappingSensitiveColumnSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSensitiveColumnSummaryStatusEnumStringValues Enumerates the set of values in String for SensitiveColumnSummaryStatusEnum
func GetSensitiveColumnSummaryStatusEnumStringValues() []string {
	return []string{
		"VALID",
		"INVALID",
	}
}

// GetMappingSensitiveColumnSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSensitiveColumnSummaryStatusEnum(val string) (SensitiveColumnSummaryStatusEnum, bool) {
	mappingSensitiveColumnSummaryStatusEnumIgnoreCase := make(map[string]SensitiveColumnSummaryStatusEnum)
	for k, v := range mappingSensitiveColumnSummaryStatusEnum {
		mappingSensitiveColumnSummaryStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSensitiveColumnSummaryStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SensitiveColumnSummarySourceEnum Enum with underlying type: string
type SensitiveColumnSummarySourceEnum string

// Set of constants representing the allowable values for SensitiveColumnSummarySourceEnum
const (
	SensitiveColumnSummarySourceManual    SensitiveColumnSummarySourceEnum = "MANUAL"
	SensitiveColumnSummarySourceDiscovery SensitiveColumnSummarySourceEnum = "DISCOVERY"
)

var mappingSensitiveColumnSummarySourceEnum = map[string]SensitiveColumnSummarySourceEnum{
	"MANUAL":    SensitiveColumnSummarySourceManual,
	"DISCOVERY": SensitiveColumnSummarySourceDiscovery,
}

// GetSensitiveColumnSummarySourceEnumValues Enumerates the set of values for SensitiveColumnSummarySourceEnum
func GetSensitiveColumnSummarySourceEnumValues() []SensitiveColumnSummarySourceEnum {
	values := make([]SensitiveColumnSummarySourceEnum, 0)
	for _, v := range mappingSensitiveColumnSummarySourceEnum {
		values = append(values, v)
	}
	return values
}

// GetSensitiveColumnSummarySourceEnumStringValues Enumerates the set of values in String for SensitiveColumnSummarySourceEnum
func GetSensitiveColumnSummarySourceEnumStringValues() []string {
	return []string{
		"MANUAL",
		"DISCOVERY",
	}
}

// GetMappingSensitiveColumnSummarySourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSensitiveColumnSummarySourceEnum(val string) (SensitiveColumnSummarySourceEnum, bool) {
	mappingSensitiveColumnSummarySourceEnumIgnoreCase := make(map[string]SensitiveColumnSummarySourceEnum)
	for k, v := range mappingSensitiveColumnSummarySourceEnum {
		mappingSensitiveColumnSummarySourceEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSensitiveColumnSummarySourceEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SensitiveColumnSummaryRelationTypeEnum Enum with underlying type: string
type SensitiveColumnSummaryRelationTypeEnum string

// Set of constants representing the allowable values for SensitiveColumnSummaryRelationTypeEnum
const (
	SensitiveColumnSummaryRelationTypeNone       SensitiveColumnSummaryRelationTypeEnum = "NONE"
	SensitiveColumnSummaryRelationTypeAppDefined SensitiveColumnSummaryRelationTypeEnum = "APP_DEFINED"
	SensitiveColumnSummaryRelationTypeDbDefined  SensitiveColumnSummaryRelationTypeEnum = "DB_DEFINED"
)

var mappingSensitiveColumnSummaryRelationTypeEnum = map[string]SensitiveColumnSummaryRelationTypeEnum{
	"NONE":        SensitiveColumnSummaryRelationTypeNone,
	"APP_DEFINED": SensitiveColumnSummaryRelationTypeAppDefined,
	"DB_DEFINED":  SensitiveColumnSummaryRelationTypeDbDefined,
}

// GetSensitiveColumnSummaryRelationTypeEnumValues Enumerates the set of values for SensitiveColumnSummaryRelationTypeEnum
func GetSensitiveColumnSummaryRelationTypeEnumValues() []SensitiveColumnSummaryRelationTypeEnum {
	values := make([]SensitiveColumnSummaryRelationTypeEnum, 0)
	for _, v := range mappingSensitiveColumnSummaryRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSensitiveColumnSummaryRelationTypeEnumStringValues Enumerates the set of values in String for SensitiveColumnSummaryRelationTypeEnum
func GetSensitiveColumnSummaryRelationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingSensitiveColumnSummaryRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSensitiveColumnSummaryRelationTypeEnum(val string) (SensitiveColumnSummaryRelationTypeEnum, bool) {
	mappingSensitiveColumnSummaryRelationTypeEnumIgnoreCase := make(map[string]SensitiveColumnSummaryRelationTypeEnum)
	for k, v := range mappingSensitiveColumnSummaryRelationTypeEnum {
		mappingSensitiveColumnSummaryRelationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSensitiveColumnSummaryRelationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
