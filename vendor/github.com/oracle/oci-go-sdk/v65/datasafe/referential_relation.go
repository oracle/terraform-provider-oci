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

// ReferentialRelation A referential relation is a resource corresponding to database columns.
// It's a subresource of sensitive data model resource and is always associated with a sensitive data model.
type ReferentialRelation struct {

	// The unique key that identifies the referential relation. It's numeric and unique within a sensitive data model.
	Key *string `mandatory:"true" json:"key"`

	// The current state of the referential relation.
	LifecycleState ReferentialRelationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the sensitive data model that contains the sensitive column.
	SensitiveDataModelId *string `mandatory:"true" json:"sensitiveDataModelId"`

	// The type of referential relationship the sensitive column has with its parent. NONE indicates that the
	// sensitive column does not have a parent. DB_DEFINED indicates that the relationship is defined in the database
	// dictionary. APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary.
	RelationType ReferentialRelationRelationTypeEnum `mandatory:"true" json:"relationType"`

	Parent *ColumnsInfo `mandatory:"true" json:"parent"`

	Child *ColumnsInfo `mandatory:"true" json:"child"`

	// Determines if the columns present in the referential relation is present in the sensitive data model
	IsSensitive *bool `mandatory:"false" json:"isSensitive"`
}

func (m ReferentialRelation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReferentialRelation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReferentialRelationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetReferentialRelationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReferentialRelationRelationTypeEnum(string(m.RelationType)); !ok && m.RelationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", m.RelationType, strings.Join(GetReferentialRelationRelationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReferentialRelationRelationTypeEnum Enum with underlying type: string
type ReferentialRelationRelationTypeEnum string

// Set of constants representing the allowable values for ReferentialRelationRelationTypeEnum
const (
	ReferentialRelationRelationTypeNone       ReferentialRelationRelationTypeEnum = "NONE"
	ReferentialRelationRelationTypeAppDefined ReferentialRelationRelationTypeEnum = "APP_DEFINED"
	ReferentialRelationRelationTypeDbDefined  ReferentialRelationRelationTypeEnum = "DB_DEFINED"
)

var mappingReferentialRelationRelationTypeEnum = map[string]ReferentialRelationRelationTypeEnum{
	"NONE":        ReferentialRelationRelationTypeNone,
	"APP_DEFINED": ReferentialRelationRelationTypeAppDefined,
	"DB_DEFINED":  ReferentialRelationRelationTypeDbDefined,
}

var mappingReferentialRelationRelationTypeEnumLowerCase = map[string]ReferentialRelationRelationTypeEnum{
	"none":        ReferentialRelationRelationTypeNone,
	"app_defined": ReferentialRelationRelationTypeAppDefined,
	"db_defined":  ReferentialRelationRelationTypeDbDefined,
}

// GetReferentialRelationRelationTypeEnumValues Enumerates the set of values for ReferentialRelationRelationTypeEnum
func GetReferentialRelationRelationTypeEnumValues() []ReferentialRelationRelationTypeEnum {
	values := make([]ReferentialRelationRelationTypeEnum, 0)
	for _, v := range mappingReferentialRelationRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReferentialRelationRelationTypeEnumStringValues Enumerates the set of values in String for ReferentialRelationRelationTypeEnum
func GetReferentialRelationRelationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingReferentialRelationRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReferentialRelationRelationTypeEnum(val string) (ReferentialRelationRelationTypeEnum, bool) {
	enum, ok := mappingReferentialRelationRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
