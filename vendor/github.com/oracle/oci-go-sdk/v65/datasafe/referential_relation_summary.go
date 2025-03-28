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

// ReferentialRelationSummary A referential relation is a resource corresponding to a database columns.
// It's a subresource of sensitive data model resource and is always associated with a sensitive data model.
type ReferentialRelationSummary struct {

	// The unique key that identifies the referential relation. It's numeric and unique within a sensitive data model.
	Key *string `mandatory:"true" json:"key"`

	// The current state of the referential relation.
	LifecycleState ReferentialRelationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the sensitive data model that contains the sensitive column.
	SensitiveDataModelId *string `mandatory:"true" json:"sensitiveDataModelId"`

	// The type of referential relationship the sensitive column has with its parent. NONE indicates that the
	// sensitive column does not have a parent. DB_DEFINED indicates that the relationship is defined in the database
	// dictionary. APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary.
	RelationType ReferentialRelationSummaryRelationTypeEnum `mandatory:"true" json:"relationType"`

	Parent *ColumnsInfo `mandatory:"true" json:"parent"`

	Child *ColumnsInfo `mandatory:"true" json:"child"`

	// Determines if the columns present in the referential relation is present in the sensitive data model
	IsSensitive *bool `mandatory:"false" json:"isSensitive"`
}

func (m ReferentialRelationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReferentialRelationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReferentialRelationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetReferentialRelationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReferentialRelationSummaryRelationTypeEnum(string(m.RelationType)); !ok && m.RelationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", m.RelationType, strings.Join(GetReferentialRelationSummaryRelationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReferentialRelationSummaryRelationTypeEnum Enum with underlying type: string
type ReferentialRelationSummaryRelationTypeEnum string

// Set of constants representing the allowable values for ReferentialRelationSummaryRelationTypeEnum
const (
	ReferentialRelationSummaryRelationTypeNone       ReferentialRelationSummaryRelationTypeEnum = "NONE"
	ReferentialRelationSummaryRelationTypeAppDefined ReferentialRelationSummaryRelationTypeEnum = "APP_DEFINED"
	ReferentialRelationSummaryRelationTypeDbDefined  ReferentialRelationSummaryRelationTypeEnum = "DB_DEFINED"
)

var mappingReferentialRelationSummaryRelationTypeEnum = map[string]ReferentialRelationSummaryRelationTypeEnum{
	"NONE":        ReferentialRelationSummaryRelationTypeNone,
	"APP_DEFINED": ReferentialRelationSummaryRelationTypeAppDefined,
	"DB_DEFINED":  ReferentialRelationSummaryRelationTypeDbDefined,
}

var mappingReferentialRelationSummaryRelationTypeEnumLowerCase = map[string]ReferentialRelationSummaryRelationTypeEnum{
	"none":        ReferentialRelationSummaryRelationTypeNone,
	"app_defined": ReferentialRelationSummaryRelationTypeAppDefined,
	"db_defined":  ReferentialRelationSummaryRelationTypeDbDefined,
}

// GetReferentialRelationSummaryRelationTypeEnumValues Enumerates the set of values for ReferentialRelationSummaryRelationTypeEnum
func GetReferentialRelationSummaryRelationTypeEnumValues() []ReferentialRelationSummaryRelationTypeEnum {
	values := make([]ReferentialRelationSummaryRelationTypeEnum, 0)
	for _, v := range mappingReferentialRelationSummaryRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReferentialRelationSummaryRelationTypeEnumStringValues Enumerates the set of values in String for ReferentialRelationSummaryRelationTypeEnum
func GetReferentialRelationSummaryRelationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingReferentialRelationSummaryRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReferentialRelationSummaryRelationTypeEnum(val string) (ReferentialRelationSummaryRelationTypeEnum, bool) {
	enum, ok := mappingReferentialRelationSummaryRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
