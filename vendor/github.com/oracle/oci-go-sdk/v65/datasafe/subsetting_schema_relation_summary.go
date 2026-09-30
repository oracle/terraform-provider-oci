// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SubsettingSchemaRelationSummary Summary of a relationship between tables in the subsetting for a subsetting policy.
type SubsettingSchemaRelationSummary struct {

	// The unique key that identifies a relation between subsetting tables. The key is numeric and unique within a subsetting policy.
	Key *string `mandatory:"true" json:"key"`

	// The key that identifies the parent subsetting table in this relation.
	ParentObjectKey *string `mandatory:"true" json:"parentObjectKey"`

	// The database schema that contains the parent subsetting table
	ParentSchemaName *string `mandatory:"true" json:"parentSchemaName"`

	// The name of the parent subsetting table
	ParentObjectName *string `mandatory:"true" json:"parentObjectName"`

	// Unique identifiers identifying the parents columns in the relation.
	ParentColumns []string `mandatory:"true" json:"parentColumns"`

	// The key that identifies the child subsetting table in this relation.
	ChildObjectKey *string `mandatory:"true" json:"childObjectKey"`

	// The database schema that contains the child subsetting table
	ChildSchemaName *string `mandatory:"true" json:"childSchemaName"`

	// The name of the child subsetting table
	ChildObjectName *string `mandatory:"true" json:"childObjectName"`

	// Unique identifiers identifying the child columns in the relation.
	ChildColumns []string `mandatory:"true" json:"childColumns"`

	// The type of referential relationship the column has with its parent. NONE indicates that the
	// sensitive column does not have a parent. DB_DEFINED indicates that the relationship is defined in the database
	// dictionary. APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary.
	RelationType SubsettingSchemaRelationSummaryRelationTypeEnum `mandatory:"true" json:"relationType"`

	// The date and time the subsetting relation was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the subsetting relation was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m SubsettingSchemaRelationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubsettingSchemaRelationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubsettingSchemaRelationSummaryRelationTypeEnum(string(m.RelationType)); !ok && m.RelationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", m.RelationType, strings.Join(GetSubsettingSchemaRelationSummaryRelationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubsettingSchemaRelationSummaryRelationTypeEnum Enum with underlying type: string
type SubsettingSchemaRelationSummaryRelationTypeEnum string

// Set of constants representing the allowable values for SubsettingSchemaRelationSummaryRelationTypeEnum
const (
	SubsettingSchemaRelationSummaryRelationTypeNone       SubsettingSchemaRelationSummaryRelationTypeEnum = "NONE"
	SubsettingSchemaRelationSummaryRelationTypeAppDefined SubsettingSchemaRelationSummaryRelationTypeEnum = "APP_DEFINED"
	SubsettingSchemaRelationSummaryRelationTypeDbDefined  SubsettingSchemaRelationSummaryRelationTypeEnum = "DB_DEFINED"
)

var mappingSubsettingSchemaRelationSummaryRelationTypeEnum = map[string]SubsettingSchemaRelationSummaryRelationTypeEnum{
	"NONE":        SubsettingSchemaRelationSummaryRelationTypeNone,
	"APP_DEFINED": SubsettingSchemaRelationSummaryRelationTypeAppDefined,
	"DB_DEFINED":  SubsettingSchemaRelationSummaryRelationTypeDbDefined,
}

var mappingSubsettingSchemaRelationSummaryRelationTypeEnumLowerCase = map[string]SubsettingSchemaRelationSummaryRelationTypeEnum{
	"none":        SubsettingSchemaRelationSummaryRelationTypeNone,
	"app_defined": SubsettingSchemaRelationSummaryRelationTypeAppDefined,
	"db_defined":  SubsettingSchemaRelationSummaryRelationTypeDbDefined,
}

// GetSubsettingSchemaRelationSummaryRelationTypeEnumValues Enumerates the set of values for SubsettingSchemaRelationSummaryRelationTypeEnum
func GetSubsettingSchemaRelationSummaryRelationTypeEnumValues() []SubsettingSchemaRelationSummaryRelationTypeEnum {
	values := make([]SubsettingSchemaRelationSummaryRelationTypeEnum, 0)
	for _, v := range mappingSubsettingSchemaRelationSummaryRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingSchemaRelationSummaryRelationTypeEnumStringValues Enumerates the set of values in String for SubsettingSchemaRelationSummaryRelationTypeEnum
func GetSubsettingSchemaRelationSummaryRelationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingSubsettingSchemaRelationSummaryRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingSchemaRelationSummaryRelationTypeEnum(val string) (SubsettingSchemaRelationSummaryRelationTypeEnum, bool) {
	enum, ok := mappingSubsettingSchemaRelationSummaryRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
