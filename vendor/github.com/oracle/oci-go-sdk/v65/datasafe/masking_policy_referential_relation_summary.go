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

// MaskingPolicyReferentialRelationSummary A referential relation is a resource corresponding to database columns.
// It is always associated with a masking policy.
type MaskingPolicyReferentialRelationSummary struct {

	// The OCID of the masking policy that contains the column.
	MaskingPolicyId *string `mandatory:"true" json:"maskingPolicyId"`

	// The type of referential relationship the column has with its parent. DB_DEFINED indicates that the relationship is defined in the database
	// dictionary. APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary.
	RelationType MaskingPolicyReferentialRelationSummaryRelationTypeEnum `mandatory:"true" json:"relationType"`

	Parent *MaskingPolicyColumnsInfo `mandatory:"true" json:"parent"`

	Child *MaskingPolicyColumnsInfo `mandatory:"true" json:"child"`

	// The masking format associated with the parent column.
	MaskingFormat []string `mandatory:"false" json:"maskingFormat"`
}

func (m MaskingPolicyReferentialRelationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaskingPolicyReferentialRelationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaskingPolicyReferentialRelationSummaryRelationTypeEnum(string(m.RelationType)); !ok && m.RelationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", m.RelationType, strings.Join(GetMaskingPolicyReferentialRelationSummaryRelationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MaskingPolicyReferentialRelationSummaryRelationTypeEnum Enum with underlying type: string
type MaskingPolicyReferentialRelationSummaryRelationTypeEnum string

// Set of constants representing the allowable values for MaskingPolicyReferentialRelationSummaryRelationTypeEnum
const (
	MaskingPolicyReferentialRelationSummaryRelationTypeNone       MaskingPolicyReferentialRelationSummaryRelationTypeEnum = "NONE"
	MaskingPolicyReferentialRelationSummaryRelationTypeAppDefined MaskingPolicyReferentialRelationSummaryRelationTypeEnum = "APP_DEFINED"
	MaskingPolicyReferentialRelationSummaryRelationTypeDbDefined  MaskingPolicyReferentialRelationSummaryRelationTypeEnum = "DB_DEFINED"
)

var mappingMaskingPolicyReferentialRelationSummaryRelationTypeEnum = map[string]MaskingPolicyReferentialRelationSummaryRelationTypeEnum{
	"NONE":        MaskingPolicyReferentialRelationSummaryRelationTypeNone,
	"APP_DEFINED": MaskingPolicyReferentialRelationSummaryRelationTypeAppDefined,
	"DB_DEFINED":  MaskingPolicyReferentialRelationSummaryRelationTypeDbDefined,
}

var mappingMaskingPolicyReferentialRelationSummaryRelationTypeEnumLowerCase = map[string]MaskingPolicyReferentialRelationSummaryRelationTypeEnum{
	"none":        MaskingPolicyReferentialRelationSummaryRelationTypeNone,
	"app_defined": MaskingPolicyReferentialRelationSummaryRelationTypeAppDefined,
	"db_defined":  MaskingPolicyReferentialRelationSummaryRelationTypeDbDefined,
}

// GetMaskingPolicyReferentialRelationSummaryRelationTypeEnumValues Enumerates the set of values for MaskingPolicyReferentialRelationSummaryRelationTypeEnum
func GetMaskingPolicyReferentialRelationSummaryRelationTypeEnumValues() []MaskingPolicyReferentialRelationSummaryRelationTypeEnum {
	values := make([]MaskingPolicyReferentialRelationSummaryRelationTypeEnum, 0)
	for _, v := range mappingMaskingPolicyReferentialRelationSummaryRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaskingPolicyReferentialRelationSummaryRelationTypeEnumStringValues Enumerates the set of values in String for MaskingPolicyReferentialRelationSummaryRelationTypeEnum
func GetMaskingPolicyReferentialRelationSummaryRelationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingMaskingPolicyReferentialRelationSummaryRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaskingPolicyReferentialRelationSummaryRelationTypeEnum(val string) (MaskingPolicyReferentialRelationSummaryRelationTypeEnum, bool) {
	enum, ok := mappingMaskingPolicyReferentialRelationSummaryRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
