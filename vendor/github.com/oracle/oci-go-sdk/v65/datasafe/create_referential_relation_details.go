// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateReferentialRelationDetails A sensitive column is a resource corresponding to a database column that is considered sensitive.
// It's a subresource of sensitive data model resource and is always associated with a sensitive data model.
// Note that referential relationships are also managed as part of sensitive columns.
type CreateReferentialRelationDetails struct {

	// The type of referential relationship the sensitive column has with its parent.
	// DB_DEFINED indicates that the relationship is defined in the database dictionary.
	// APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary.
	RelationType CreateReferentialRelationDetailsRelationTypeEnum `mandatory:"true" json:"relationType"`

	Parent *ColumnsInfo `mandatory:"true" json:"parent"`

	Child *ColumnsInfo `mandatory:"true" json:"child"`

	// Add to sensitive data model if passed true. If false is passed, then the
	// columns will not be added in the sensitive data model as sensitive columns and
	// if sensitive type OCIDs are assigned to the columns, then the sensitive type
	// OCIDs will not be retained.
	IsSensitive *bool `mandatory:"false" json:"isSensitive"`
}

func (m CreateReferentialRelationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateReferentialRelationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateReferentialRelationDetailsRelationTypeEnum(string(m.RelationType)); !ok && m.RelationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", m.RelationType, strings.Join(GetCreateReferentialRelationDetailsRelationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateReferentialRelationDetailsRelationTypeEnum Enum with underlying type: string
type CreateReferentialRelationDetailsRelationTypeEnum string

// Set of constants representing the allowable values for CreateReferentialRelationDetailsRelationTypeEnum
const (
	CreateReferentialRelationDetailsRelationTypeAppDefined CreateReferentialRelationDetailsRelationTypeEnum = "APP_DEFINED"
	CreateReferentialRelationDetailsRelationTypeDbDefined  CreateReferentialRelationDetailsRelationTypeEnum = "DB_DEFINED"
)

var mappingCreateReferentialRelationDetailsRelationTypeEnum = map[string]CreateReferentialRelationDetailsRelationTypeEnum{
	"APP_DEFINED": CreateReferentialRelationDetailsRelationTypeAppDefined,
	"DB_DEFINED":  CreateReferentialRelationDetailsRelationTypeDbDefined,
}

var mappingCreateReferentialRelationDetailsRelationTypeEnumLowerCase = map[string]CreateReferentialRelationDetailsRelationTypeEnum{
	"app_defined": CreateReferentialRelationDetailsRelationTypeAppDefined,
	"db_defined":  CreateReferentialRelationDetailsRelationTypeDbDefined,
}

// GetCreateReferentialRelationDetailsRelationTypeEnumValues Enumerates the set of values for CreateReferentialRelationDetailsRelationTypeEnum
func GetCreateReferentialRelationDetailsRelationTypeEnumValues() []CreateReferentialRelationDetailsRelationTypeEnum {
	values := make([]CreateReferentialRelationDetailsRelationTypeEnum, 0)
	for _, v := range mappingCreateReferentialRelationDetailsRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateReferentialRelationDetailsRelationTypeEnumStringValues Enumerates the set of values in String for CreateReferentialRelationDetailsRelationTypeEnum
func GetCreateReferentialRelationDetailsRelationTypeEnumStringValues() []string {
	return []string{
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingCreateReferentialRelationDetailsRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateReferentialRelationDetailsRelationTypeEnum(val string) (CreateReferentialRelationDetailsRelationTypeEnum, bool) {
	enum, ok := mappingCreateReferentialRelationDetailsRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
