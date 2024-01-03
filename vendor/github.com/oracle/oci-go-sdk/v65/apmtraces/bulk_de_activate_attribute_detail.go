// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// BulkDeActivateAttributeDetail Object that contains the details about a single attribute to be deactivated.
type BulkDeActivateAttributeDetail struct {

	// Name of the attribute to be deactivated.
	AttributeName *string `mandatory:"true" json:"attributeName"`

	// Namespace of the attribute to be deactivated.  The attributeNameSpace will default to TRACES if it is
	// not passed in.
	AttributeNameSpace BulkDeActivateAttributeDetailAttributeNameSpaceEnum `mandatory:"false" json:"attributeNameSpace,omitempty"`
}

func (m BulkDeActivateAttributeDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkDeActivateAttributeDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBulkDeActivateAttributeDetailAttributeNameSpaceEnum(string(m.AttributeNameSpace)); !ok && m.AttributeNameSpace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeNameSpace: %s. Supported values are: %s.", m.AttributeNameSpace, strings.Join(GetBulkDeActivateAttributeDetailAttributeNameSpaceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkDeActivateAttributeDetailAttributeNameSpaceEnum Enum with underlying type: string
type BulkDeActivateAttributeDetailAttributeNameSpaceEnum string

// Set of constants representing the allowable values for BulkDeActivateAttributeDetailAttributeNameSpaceEnum
const (
	BulkDeActivateAttributeDetailAttributeNameSpaceTraces    BulkDeActivateAttributeDetailAttributeNameSpaceEnum = "TRACES"
	BulkDeActivateAttributeDetailAttributeNameSpaceSynthetic BulkDeActivateAttributeDetailAttributeNameSpaceEnum = "SYNTHETIC"
)

var mappingBulkDeActivateAttributeDetailAttributeNameSpaceEnum = map[string]BulkDeActivateAttributeDetailAttributeNameSpaceEnum{
	"TRACES":    BulkDeActivateAttributeDetailAttributeNameSpaceTraces,
	"SYNTHETIC": BulkDeActivateAttributeDetailAttributeNameSpaceSynthetic,
}

var mappingBulkDeActivateAttributeDetailAttributeNameSpaceEnumLowerCase = map[string]BulkDeActivateAttributeDetailAttributeNameSpaceEnum{
	"traces":    BulkDeActivateAttributeDetailAttributeNameSpaceTraces,
	"synthetic": BulkDeActivateAttributeDetailAttributeNameSpaceSynthetic,
}

// GetBulkDeActivateAttributeDetailAttributeNameSpaceEnumValues Enumerates the set of values for BulkDeActivateAttributeDetailAttributeNameSpaceEnum
func GetBulkDeActivateAttributeDetailAttributeNameSpaceEnumValues() []BulkDeActivateAttributeDetailAttributeNameSpaceEnum {
	values := make([]BulkDeActivateAttributeDetailAttributeNameSpaceEnum, 0)
	for _, v := range mappingBulkDeActivateAttributeDetailAttributeNameSpaceEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkDeActivateAttributeDetailAttributeNameSpaceEnumStringValues Enumerates the set of values in String for BulkDeActivateAttributeDetailAttributeNameSpaceEnum
func GetBulkDeActivateAttributeDetailAttributeNameSpaceEnumStringValues() []string {
	return []string{
		"TRACES",
		"SYNTHETIC",
	}
}

// GetMappingBulkDeActivateAttributeDetailAttributeNameSpaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkDeActivateAttributeDetailAttributeNameSpaceEnum(val string) (BulkDeActivateAttributeDetailAttributeNameSpaceEnum, bool) {
	enum, ok := mappingBulkDeActivateAttributeDetailAttributeNameSpaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
