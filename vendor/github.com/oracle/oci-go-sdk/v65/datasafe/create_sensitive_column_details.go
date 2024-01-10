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

// CreateSensitiveColumnDetails Details to create a new sensitive column in a sensitive data model.
type CreateSensitiveColumnDetails struct {

	// The database schema that contains the sensitive column.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The database object that contains the sensitive column.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The name of the sensitive column.
	ColumnName *string `mandatory:"true" json:"columnName"`

	// The name of the application associated with the sensitive column. It's useful when the application name is
	// different from the schema name. Otherwise, it can be ignored. If this attribute is not provided, it's automatically
	// populated with the value provided for the schemaName attribute.
	AppName *string `mandatory:"false" json:"appName"`

	// The type of the database object that contains the sensitive column.
	ObjectType CreateSensitiveColumnDetailsObjectTypeEnum `mandatory:"false" json:"objectType,omitempty"`

	// The data type of the sensitive column.
	DataType *string `mandatory:"false" json:"dataType"`

	// The status of the sensitive column. VALID means the column is considered sensitive. INVALID means the column
	// is not considered sensitive. Tracking invalid columns in a sensitive data model helps ensure that an incremental
	// data discovery job does not identify these columns as sensitive.
	Status CreateSensitiveColumnDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The OCID of the sensitive type to be associated with the sensitive column.
	SensitiveTypeId *string `mandatory:"false" json:"sensitiveTypeId"`

	// Unique keys identifying the columns that are parents of the sensitive column. At present, it accepts only one
	// parent column key. This attribute can be used to establish relationship between columns in a sensitive data model.
	// Note that the parent column must be added to the sensitive data model before its key can be specified here.
	// If this attribute is provided, the appDefinedChildColumnKeys or dbDefinedChildColumnKeys attribute of the parent
	// column is automatically updated to reflect the relationship.
	ParentColumnKeys []string `mandatory:"false" json:"parentColumnKeys"`

	// The type of referential relationship the sensitive column has with its parent. NONE indicates that the sensitive
	// column does not have a parent. DB_DEFINED indicates that the relationship is defined in the database dictionary.
	// APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary.
	RelationType CreateSensitiveColumnDetailsRelationTypeEnum `mandatory:"false" json:"relationType,omitempty"`

	// Unique keys identifying the columns that are application-level (non-dictionary) children of the sensitive column.
	// This attribute can be used to establish relationship between columns in a sensitive data model. Note that the
	// child columns must be added to the sensitive data model before their keys can be specified here. If this attribute
	// is provided, the parentColumnKeys and relationType attributes of the child columns are automatically updated to reflect the relationship.
	AppDefinedChildColumnKeys []string `mandatory:"false" json:"appDefinedChildColumnKeys"`

	// Unique keys identifying the columns that are database-level (dictionary-defined) children of the sensitive column.
	// This attribute can be used to establish relationship between columns in a sensitive data model. Note that the
	// child columns must be added to the sensitive data model before their keys can be specified here. If this attribute
	// is provided, the parentColumnKeys and relationType attributes of the child columns are automatically updated to reflect the relationship.
	DbDefinedChildColumnKeys []string `mandatory:"false" json:"dbDefinedChildColumnKeys"`
}

func (m CreateSensitiveColumnDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSensitiveColumnDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateSensitiveColumnDetailsObjectTypeEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetCreateSensitiveColumnDetailsObjectTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateSensitiveColumnDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCreateSensitiveColumnDetailsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateSensitiveColumnDetailsRelationTypeEnum(string(m.RelationType)); !ok && m.RelationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", m.RelationType, strings.Join(GetCreateSensitiveColumnDetailsRelationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateSensitiveColumnDetailsObjectTypeEnum Enum with underlying type: string
type CreateSensitiveColumnDetailsObjectTypeEnum string

// Set of constants representing the allowable values for CreateSensitiveColumnDetailsObjectTypeEnum
const (
	CreateSensitiveColumnDetailsObjectTypeTable          CreateSensitiveColumnDetailsObjectTypeEnum = "TABLE"
	CreateSensitiveColumnDetailsObjectTypeEditioningView CreateSensitiveColumnDetailsObjectTypeEnum = "EDITIONING_VIEW"
)

var mappingCreateSensitiveColumnDetailsObjectTypeEnum = map[string]CreateSensitiveColumnDetailsObjectTypeEnum{
	"TABLE":           CreateSensitiveColumnDetailsObjectTypeTable,
	"EDITIONING_VIEW": CreateSensitiveColumnDetailsObjectTypeEditioningView,
}

var mappingCreateSensitiveColumnDetailsObjectTypeEnumLowerCase = map[string]CreateSensitiveColumnDetailsObjectTypeEnum{
	"table":           CreateSensitiveColumnDetailsObjectTypeTable,
	"editioning_view": CreateSensitiveColumnDetailsObjectTypeEditioningView,
}

// GetCreateSensitiveColumnDetailsObjectTypeEnumValues Enumerates the set of values for CreateSensitiveColumnDetailsObjectTypeEnum
func GetCreateSensitiveColumnDetailsObjectTypeEnumValues() []CreateSensitiveColumnDetailsObjectTypeEnum {
	values := make([]CreateSensitiveColumnDetailsObjectTypeEnum, 0)
	for _, v := range mappingCreateSensitiveColumnDetailsObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSensitiveColumnDetailsObjectTypeEnumStringValues Enumerates the set of values in String for CreateSensitiveColumnDetailsObjectTypeEnum
func GetCreateSensitiveColumnDetailsObjectTypeEnumStringValues() []string {
	return []string{
		"TABLE",
		"EDITIONING_VIEW",
	}
}

// GetMappingCreateSensitiveColumnDetailsObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSensitiveColumnDetailsObjectTypeEnum(val string) (CreateSensitiveColumnDetailsObjectTypeEnum, bool) {
	enum, ok := mappingCreateSensitiveColumnDetailsObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateSensitiveColumnDetailsStatusEnum Enum with underlying type: string
type CreateSensitiveColumnDetailsStatusEnum string

// Set of constants representing the allowable values for CreateSensitiveColumnDetailsStatusEnum
const (
	CreateSensitiveColumnDetailsStatusValid   CreateSensitiveColumnDetailsStatusEnum = "VALID"
	CreateSensitiveColumnDetailsStatusInvalid CreateSensitiveColumnDetailsStatusEnum = "INVALID"
)

var mappingCreateSensitiveColumnDetailsStatusEnum = map[string]CreateSensitiveColumnDetailsStatusEnum{
	"VALID":   CreateSensitiveColumnDetailsStatusValid,
	"INVALID": CreateSensitiveColumnDetailsStatusInvalid,
}

var mappingCreateSensitiveColumnDetailsStatusEnumLowerCase = map[string]CreateSensitiveColumnDetailsStatusEnum{
	"valid":   CreateSensitiveColumnDetailsStatusValid,
	"invalid": CreateSensitiveColumnDetailsStatusInvalid,
}

// GetCreateSensitiveColumnDetailsStatusEnumValues Enumerates the set of values for CreateSensitiveColumnDetailsStatusEnum
func GetCreateSensitiveColumnDetailsStatusEnumValues() []CreateSensitiveColumnDetailsStatusEnum {
	values := make([]CreateSensitiveColumnDetailsStatusEnum, 0)
	for _, v := range mappingCreateSensitiveColumnDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSensitiveColumnDetailsStatusEnumStringValues Enumerates the set of values in String for CreateSensitiveColumnDetailsStatusEnum
func GetCreateSensitiveColumnDetailsStatusEnumStringValues() []string {
	return []string{
		"VALID",
		"INVALID",
	}
}

// GetMappingCreateSensitiveColumnDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSensitiveColumnDetailsStatusEnum(val string) (CreateSensitiveColumnDetailsStatusEnum, bool) {
	enum, ok := mappingCreateSensitiveColumnDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateSensitiveColumnDetailsRelationTypeEnum Enum with underlying type: string
type CreateSensitiveColumnDetailsRelationTypeEnum string

// Set of constants representing the allowable values for CreateSensitiveColumnDetailsRelationTypeEnum
const (
	CreateSensitiveColumnDetailsRelationTypeNone       CreateSensitiveColumnDetailsRelationTypeEnum = "NONE"
	CreateSensitiveColumnDetailsRelationTypeAppDefined CreateSensitiveColumnDetailsRelationTypeEnum = "APP_DEFINED"
	CreateSensitiveColumnDetailsRelationTypeDbDefined  CreateSensitiveColumnDetailsRelationTypeEnum = "DB_DEFINED"
)

var mappingCreateSensitiveColumnDetailsRelationTypeEnum = map[string]CreateSensitiveColumnDetailsRelationTypeEnum{
	"NONE":        CreateSensitiveColumnDetailsRelationTypeNone,
	"APP_DEFINED": CreateSensitiveColumnDetailsRelationTypeAppDefined,
	"DB_DEFINED":  CreateSensitiveColumnDetailsRelationTypeDbDefined,
}

var mappingCreateSensitiveColumnDetailsRelationTypeEnumLowerCase = map[string]CreateSensitiveColumnDetailsRelationTypeEnum{
	"none":        CreateSensitiveColumnDetailsRelationTypeNone,
	"app_defined": CreateSensitiveColumnDetailsRelationTypeAppDefined,
	"db_defined":  CreateSensitiveColumnDetailsRelationTypeDbDefined,
}

// GetCreateSensitiveColumnDetailsRelationTypeEnumValues Enumerates the set of values for CreateSensitiveColumnDetailsRelationTypeEnum
func GetCreateSensitiveColumnDetailsRelationTypeEnumValues() []CreateSensitiveColumnDetailsRelationTypeEnum {
	values := make([]CreateSensitiveColumnDetailsRelationTypeEnum, 0)
	for _, v := range mappingCreateSensitiveColumnDetailsRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSensitiveColumnDetailsRelationTypeEnumStringValues Enumerates the set of values in String for CreateSensitiveColumnDetailsRelationTypeEnum
func GetCreateSensitiveColumnDetailsRelationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingCreateSensitiveColumnDetailsRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSensitiveColumnDetailsRelationTypeEnum(val string) (CreateSensitiveColumnDetailsRelationTypeEnum, bool) {
	enum, ok := mappingCreateSensitiveColumnDetailsRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
