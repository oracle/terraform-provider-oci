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

// UpdateSensitiveColumnDetails Details to update a sensitive column in a sensitive data model.
type UpdateSensitiveColumnDetails struct {

	// The data type of the sensitive column.
	DataType *string `mandatory:"false" json:"dataType"`

	// The status of the sensitive column. VALID means the column is considered sensitive. INVALID means the column
	// is not considered sensitive. Tracking invalid columns in a sensitive data model helps ensure that an incremental
	// data discovery job does not identify these columns as sensitive.
	Status UpdateSensitiveColumnDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The OCID of the sensitive type to be associated with the sensitive column.
	SensitiveTypeId *string `mandatory:"false" json:"sensitiveTypeId"`

	// Unique keys identifying the columns that are parents of the sensitive column. At present, it accepts only one
	// parent column key. This attribute can be used to establish relationship between columns in a sensitive data model.
	// Note that the parent column must be added to the sensitive data model before its key can be specified here.
	// If this attribute is provided, the appDefinedChildColumnKeys or dbDefinedChildColumnKeys attribute of the
	// parent column is automatically updated to reflect the relationship.
	ParentColumnKeys []string `mandatory:"false" json:"parentColumnKeys"`

	// The type of referential relationship the sensitive column has with its parent. NONE indicates that the sensitive
	// column does not have a parent. DB_DEFINED indicates that the relationship is defined in the database dictionary.
	// APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary.
	RelationType UpdateSensitiveColumnDetailsRelationTypeEnum `mandatory:"false" json:"relationType,omitempty"`

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

func (m UpdateSensitiveColumnDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSensitiveColumnDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateSensitiveColumnDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUpdateSensitiveColumnDetailsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateSensitiveColumnDetailsRelationTypeEnum(string(m.RelationType)); !ok && m.RelationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", m.RelationType, strings.Join(GetUpdateSensitiveColumnDetailsRelationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateSensitiveColumnDetailsStatusEnum Enum with underlying type: string
type UpdateSensitiveColumnDetailsStatusEnum string

// Set of constants representing the allowable values for UpdateSensitiveColumnDetailsStatusEnum
const (
	UpdateSensitiveColumnDetailsStatusValid   UpdateSensitiveColumnDetailsStatusEnum = "VALID"
	UpdateSensitiveColumnDetailsStatusInvalid UpdateSensitiveColumnDetailsStatusEnum = "INVALID"
)

var mappingUpdateSensitiveColumnDetailsStatusEnum = map[string]UpdateSensitiveColumnDetailsStatusEnum{
	"VALID":   UpdateSensitiveColumnDetailsStatusValid,
	"INVALID": UpdateSensitiveColumnDetailsStatusInvalid,
}

// GetUpdateSensitiveColumnDetailsStatusEnumValues Enumerates the set of values for UpdateSensitiveColumnDetailsStatusEnum
func GetUpdateSensitiveColumnDetailsStatusEnumValues() []UpdateSensitiveColumnDetailsStatusEnum {
	values := make([]UpdateSensitiveColumnDetailsStatusEnum, 0)
	for _, v := range mappingUpdateSensitiveColumnDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSensitiveColumnDetailsStatusEnumStringValues Enumerates the set of values in String for UpdateSensitiveColumnDetailsStatusEnum
func GetUpdateSensitiveColumnDetailsStatusEnumStringValues() []string {
	return []string{
		"VALID",
		"INVALID",
	}
}

// GetMappingUpdateSensitiveColumnDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSensitiveColumnDetailsStatusEnum(val string) (UpdateSensitiveColumnDetailsStatusEnum, bool) {
	mappingUpdateSensitiveColumnDetailsStatusEnumIgnoreCase := make(map[string]UpdateSensitiveColumnDetailsStatusEnum)
	for k, v := range mappingUpdateSensitiveColumnDetailsStatusEnum {
		mappingUpdateSensitiveColumnDetailsStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateSensitiveColumnDetailsStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateSensitiveColumnDetailsRelationTypeEnum Enum with underlying type: string
type UpdateSensitiveColumnDetailsRelationTypeEnum string

// Set of constants representing the allowable values for UpdateSensitiveColumnDetailsRelationTypeEnum
const (
	UpdateSensitiveColumnDetailsRelationTypeNone       UpdateSensitiveColumnDetailsRelationTypeEnum = "NONE"
	UpdateSensitiveColumnDetailsRelationTypeAppDefined UpdateSensitiveColumnDetailsRelationTypeEnum = "APP_DEFINED"
	UpdateSensitiveColumnDetailsRelationTypeDbDefined  UpdateSensitiveColumnDetailsRelationTypeEnum = "DB_DEFINED"
)

var mappingUpdateSensitiveColumnDetailsRelationTypeEnum = map[string]UpdateSensitiveColumnDetailsRelationTypeEnum{
	"NONE":        UpdateSensitiveColumnDetailsRelationTypeNone,
	"APP_DEFINED": UpdateSensitiveColumnDetailsRelationTypeAppDefined,
	"DB_DEFINED":  UpdateSensitiveColumnDetailsRelationTypeDbDefined,
}

// GetUpdateSensitiveColumnDetailsRelationTypeEnumValues Enumerates the set of values for UpdateSensitiveColumnDetailsRelationTypeEnum
func GetUpdateSensitiveColumnDetailsRelationTypeEnumValues() []UpdateSensitiveColumnDetailsRelationTypeEnum {
	values := make([]UpdateSensitiveColumnDetailsRelationTypeEnum, 0)
	for _, v := range mappingUpdateSensitiveColumnDetailsRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSensitiveColumnDetailsRelationTypeEnumStringValues Enumerates the set of values in String for UpdateSensitiveColumnDetailsRelationTypeEnum
func GetUpdateSensitiveColumnDetailsRelationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingUpdateSensitiveColumnDetailsRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSensitiveColumnDetailsRelationTypeEnum(val string) (UpdateSensitiveColumnDetailsRelationTypeEnum, bool) {
	mappingUpdateSensitiveColumnDetailsRelationTypeEnumIgnoreCase := make(map[string]UpdateSensitiveColumnDetailsRelationTypeEnum)
	for k, v := range mappingUpdateSensitiveColumnDetailsRelationTypeEnum {
		mappingUpdateSensitiveColumnDetailsRelationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateSensitiveColumnDetailsRelationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
