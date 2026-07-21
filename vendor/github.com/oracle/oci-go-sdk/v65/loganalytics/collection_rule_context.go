// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CollectionRuleContext Context of the collection rule.
type CollectionRuleContext struct {

	// The context type to which the collection rule applies.
	Type CollectionRuleContextTypeEnum `mandatory:"true" json:"type"`

	// The context value to which the collection rule applies.
	Value *string `mandatory:"true" json:"value"`
}

func (m CollectionRuleContext) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CollectionRuleContext) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCollectionRuleContextTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCollectionRuleContextTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CollectionRuleContextTypeEnum Enum with underlying type: string
type CollectionRuleContextTypeEnum string

// Set of constants representing the allowable values for CollectionRuleContextTypeEnum
const (
	CollectionRuleContextTypeEntityType CollectionRuleContextTypeEnum = "ENTITY_TYPE"
	CollectionRuleContextTypeSource     CollectionRuleContextTypeEnum = "SOURCE"
)

var mappingCollectionRuleContextTypeEnum = map[string]CollectionRuleContextTypeEnum{
	"ENTITY_TYPE": CollectionRuleContextTypeEntityType,
	"SOURCE":      CollectionRuleContextTypeSource,
}

var mappingCollectionRuleContextTypeEnumLowerCase = map[string]CollectionRuleContextTypeEnum{
	"entity_type": CollectionRuleContextTypeEntityType,
	"source":      CollectionRuleContextTypeSource,
}

// GetCollectionRuleContextTypeEnumValues Enumerates the set of values for CollectionRuleContextTypeEnum
func GetCollectionRuleContextTypeEnumValues() []CollectionRuleContextTypeEnum {
	values := make([]CollectionRuleContextTypeEnum, 0)
	for _, v := range mappingCollectionRuleContextTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCollectionRuleContextTypeEnumStringValues Enumerates the set of values in String for CollectionRuleContextTypeEnum
func GetCollectionRuleContextTypeEnumStringValues() []string {
	return []string{
		"ENTITY_TYPE",
		"SOURCE",
	}
}

// GetMappingCollectionRuleContextTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCollectionRuleContextTypeEnum(val string) (CollectionRuleContextTypeEnum, bool) {
	enum, ok := mappingCollectionRuleContextTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
