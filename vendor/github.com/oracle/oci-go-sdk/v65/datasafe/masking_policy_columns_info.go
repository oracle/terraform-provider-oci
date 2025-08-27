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

// MaskingPolicyColumnsInfo maskingPolicyColumnsInfo object has details of column group with schema details.
type MaskingPolicyColumnsInfo struct {

	// The name of the schema that contains the database column(s).
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The type of the database object that contains the masking policy.
	ObjectType MaskingPolicyColumnsInfoObjectTypeEnum `mandatory:"true" json:"objectType"`

	// The name of the object (table or editioning view) that contains the database column(s).
	ObjectName *string `mandatory:"true" json:"objectName"`

	// Group of columns in referential relation. Order needs to be maintained in the elements of the parent/child array listing.
	ReferentialColumnGroup []string `mandatory:"true" json:"referentialColumnGroup"`
}

func (m MaskingPolicyColumnsInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaskingPolicyColumnsInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaskingPolicyColumnsInfoObjectTypeEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetMaskingPolicyColumnsInfoObjectTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MaskingPolicyColumnsInfoObjectTypeEnum Enum with underlying type: string
type MaskingPolicyColumnsInfoObjectTypeEnum string

// Set of constants representing the allowable values for MaskingPolicyColumnsInfoObjectTypeEnum
const (
	MaskingPolicyColumnsInfoObjectTypeTable MaskingPolicyColumnsInfoObjectTypeEnum = "TABLE"
)

var mappingMaskingPolicyColumnsInfoObjectTypeEnum = map[string]MaskingPolicyColumnsInfoObjectTypeEnum{
	"TABLE": MaskingPolicyColumnsInfoObjectTypeTable,
}

var mappingMaskingPolicyColumnsInfoObjectTypeEnumLowerCase = map[string]MaskingPolicyColumnsInfoObjectTypeEnum{
	"table": MaskingPolicyColumnsInfoObjectTypeTable,
}

// GetMaskingPolicyColumnsInfoObjectTypeEnumValues Enumerates the set of values for MaskingPolicyColumnsInfoObjectTypeEnum
func GetMaskingPolicyColumnsInfoObjectTypeEnumValues() []MaskingPolicyColumnsInfoObjectTypeEnum {
	values := make([]MaskingPolicyColumnsInfoObjectTypeEnum, 0)
	for _, v := range mappingMaskingPolicyColumnsInfoObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaskingPolicyColumnsInfoObjectTypeEnumStringValues Enumerates the set of values in String for MaskingPolicyColumnsInfoObjectTypeEnum
func GetMaskingPolicyColumnsInfoObjectTypeEnumStringValues() []string {
	return []string{
		"TABLE",
	}
}

// GetMappingMaskingPolicyColumnsInfoObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaskingPolicyColumnsInfoObjectTypeEnum(val string) (MaskingPolicyColumnsInfoObjectTypeEnum, bool) {
	enum, ok := mappingMaskingPolicyColumnsInfoObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
