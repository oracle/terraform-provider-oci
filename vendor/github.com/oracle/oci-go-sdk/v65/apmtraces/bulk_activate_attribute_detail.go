// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BulkActivateAttributeDetail Object that contains the details about a single attribute in the bulk request to be activated.
type BulkActivateAttributeDetail struct {

	// Name of the attribute to be activated.
	AttributeName *string `mandatory:"true" json:"attributeName"`

	// Type of the attribute to be activated.
	AttributeType BulkActivateAttributeDetailAttributeTypeEnum `mandatory:"true" json:"attributeType"`
}

func (m BulkActivateAttributeDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkActivateAttributeDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBulkActivateAttributeDetailAttributeTypeEnum(string(m.AttributeType)); !ok && m.AttributeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeType: %s. Supported values are: %s.", m.AttributeType, strings.Join(GetBulkActivateAttributeDetailAttributeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkActivateAttributeDetailAttributeTypeEnum Enum with underlying type: string
type BulkActivateAttributeDetailAttributeTypeEnum string

// Set of constants representing the allowable values for BulkActivateAttributeDetailAttributeTypeEnum
const (
	BulkActivateAttributeDetailAttributeTypeNumeric BulkActivateAttributeDetailAttributeTypeEnum = "NUMERIC"
	BulkActivateAttributeDetailAttributeTypeString  BulkActivateAttributeDetailAttributeTypeEnum = "STRING"
)

var mappingBulkActivateAttributeDetailAttributeTypeEnum = map[string]BulkActivateAttributeDetailAttributeTypeEnum{
	"NUMERIC": BulkActivateAttributeDetailAttributeTypeNumeric,
	"STRING":  BulkActivateAttributeDetailAttributeTypeString,
}

var mappingBulkActivateAttributeDetailAttributeTypeEnumLowerCase = map[string]BulkActivateAttributeDetailAttributeTypeEnum{
	"numeric": BulkActivateAttributeDetailAttributeTypeNumeric,
	"string":  BulkActivateAttributeDetailAttributeTypeString,
}

// GetBulkActivateAttributeDetailAttributeTypeEnumValues Enumerates the set of values for BulkActivateAttributeDetailAttributeTypeEnum
func GetBulkActivateAttributeDetailAttributeTypeEnumValues() []BulkActivateAttributeDetailAttributeTypeEnum {
	values := make([]BulkActivateAttributeDetailAttributeTypeEnum, 0)
	for _, v := range mappingBulkActivateAttributeDetailAttributeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkActivateAttributeDetailAttributeTypeEnumStringValues Enumerates the set of values in String for BulkActivateAttributeDetailAttributeTypeEnum
func GetBulkActivateAttributeDetailAttributeTypeEnumStringValues() []string {
	return []string{
		"NUMERIC",
		"STRING",
	}
}

// GetMappingBulkActivateAttributeDetailAttributeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkActivateAttributeDetailAttributeTypeEnum(val string) (BulkActivateAttributeDetailAttributeTypeEnum, bool) {
	enum, ok := mappingBulkActivateAttributeDetailAttributeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
