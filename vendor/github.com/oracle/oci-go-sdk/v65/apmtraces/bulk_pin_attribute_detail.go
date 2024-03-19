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

// BulkPinAttributeDetail Object that contains the details about a single attribute in the bulk request to be pinned.
type BulkPinAttributeDetail struct {

	// Name of the attribute to be pinned.
	AttributeName *string `mandatory:"true" json:"attributeName"`

	// Namespace of the attribute to be pinned.  The attributeNameSpace will default to TRACES if it is
	// not passed in.
	AttributeNameSpace BulkPinAttributeDetailAttributeNameSpaceEnum `mandatory:"false" json:"attributeNameSpace,omitempty"`
}

func (m BulkPinAttributeDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkPinAttributeDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBulkPinAttributeDetailAttributeNameSpaceEnum(string(m.AttributeNameSpace)); !ok && m.AttributeNameSpace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeNameSpace: %s. Supported values are: %s.", m.AttributeNameSpace, strings.Join(GetBulkPinAttributeDetailAttributeNameSpaceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkPinAttributeDetailAttributeNameSpaceEnum Enum with underlying type: string
type BulkPinAttributeDetailAttributeNameSpaceEnum string

// Set of constants representing the allowable values for BulkPinAttributeDetailAttributeNameSpaceEnum
const (
	BulkPinAttributeDetailAttributeNameSpaceTraces    BulkPinAttributeDetailAttributeNameSpaceEnum = "TRACES"
	BulkPinAttributeDetailAttributeNameSpaceSynthetic BulkPinAttributeDetailAttributeNameSpaceEnum = "SYNTHETIC"
)

var mappingBulkPinAttributeDetailAttributeNameSpaceEnum = map[string]BulkPinAttributeDetailAttributeNameSpaceEnum{
	"TRACES":    BulkPinAttributeDetailAttributeNameSpaceTraces,
	"SYNTHETIC": BulkPinAttributeDetailAttributeNameSpaceSynthetic,
}

var mappingBulkPinAttributeDetailAttributeNameSpaceEnumLowerCase = map[string]BulkPinAttributeDetailAttributeNameSpaceEnum{
	"traces":    BulkPinAttributeDetailAttributeNameSpaceTraces,
	"synthetic": BulkPinAttributeDetailAttributeNameSpaceSynthetic,
}

// GetBulkPinAttributeDetailAttributeNameSpaceEnumValues Enumerates the set of values for BulkPinAttributeDetailAttributeNameSpaceEnum
func GetBulkPinAttributeDetailAttributeNameSpaceEnumValues() []BulkPinAttributeDetailAttributeNameSpaceEnum {
	values := make([]BulkPinAttributeDetailAttributeNameSpaceEnum, 0)
	for _, v := range mappingBulkPinAttributeDetailAttributeNameSpaceEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkPinAttributeDetailAttributeNameSpaceEnumStringValues Enumerates the set of values in String for BulkPinAttributeDetailAttributeNameSpaceEnum
func GetBulkPinAttributeDetailAttributeNameSpaceEnumStringValues() []string {
	return []string{
		"TRACES",
		"SYNTHETIC",
	}
}

// GetMappingBulkPinAttributeDetailAttributeNameSpaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkPinAttributeDetailAttributeNameSpaceEnum(val string) (BulkPinAttributeDetailAttributeNameSpaceEnum, bool) {
	enum, ok := mappingBulkPinAttributeDetailAttributeNameSpaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
