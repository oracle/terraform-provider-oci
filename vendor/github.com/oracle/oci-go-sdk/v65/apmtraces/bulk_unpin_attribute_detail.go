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

// BulkUnpinAttributeDetail Object that contains the details about a single attribute in the bulk request to be unpinned.
type BulkUnpinAttributeDetail struct {

	// Name of the attribute to be unpinned.
	AttributeName *string `mandatory:"true" json:"attributeName"`

	// Namespace of the attribute to be unpinned.  The attributeNameSpace will default to TRACES if it is
	// not passed in.
	AttributeNameSpace BulkUnpinAttributeDetailAttributeNameSpaceEnum `mandatory:"false" json:"attributeNameSpace,omitempty"`
}

func (m BulkUnpinAttributeDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkUnpinAttributeDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBulkUnpinAttributeDetailAttributeNameSpaceEnum(string(m.AttributeNameSpace)); !ok && m.AttributeNameSpace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeNameSpace: %s. Supported values are: %s.", m.AttributeNameSpace, strings.Join(GetBulkUnpinAttributeDetailAttributeNameSpaceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkUnpinAttributeDetailAttributeNameSpaceEnum Enum with underlying type: string
type BulkUnpinAttributeDetailAttributeNameSpaceEnum string

// Set of constants representing the allowable values for BulkUnpinAttributeDetailAttributeNameSpaceEnum
const (
	BulkUnpinAttributeDetailAttributeNameSpaceTraces    BulkUnpinAttributeDetailAttributeNameSpaceEnum = "TRACES"
	BulkUnpinAttributeDetailAttributeNameSpaceSynthetic BulkUnpinAttributeDetailAttributeNameSpaceEnum = "SYNTHETIC"
)

var mappingBulkUnpinAttributeDetailAttributeNameSpaceEnum = map[string]BulkUnpinAttributeDetailAttributeNameSpaceEnum{
	"TRACES":    BulkUnpinAttributeDetailAttributeNameSpaceTraces,
	"SYNTHETIC": BulkUnpinAttributeDetailAttributeNameSpaceSynthetic,
}

var mappingBulkUnpinAttributeDetailAttributeNameSpaceEnumLowerCase = map[string]BulkUnpinAttributeDetailAttributeNameSpaceEnum{
	"traces":    BulkUnpinAttributeDetailAttributeNameSpaceTraces,
	"synthetic": BulkUnpinAttributeDetailAttributeNameSpaceSynthetic,
}

// GetBulkUnpinAttributeDetailAttributeNameSpaceEnumValues Enumerates the set of values for BulkUnpinAttributeDetailAttributeNameSpaceEnum
func GetBulkUnpinAttributeDetailAttributeNameSpaceEnumValues() []BulkUnpinAttributeDetailAttributeNameSpaceEnum {
	values := make([]BulkUnpinAttributeDetailAttributeNameSpaceEnum, 0)
	for _, v := range mappingBulkUnpinAttributeDetailAttributeNameSpaceEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkUnpinAttributeDetailAttributeNameSpaceEnumStringValues Enumerates the set of values in String for BulkUnpinAttributeDetailAttributeNameSpaceEnum
func GetBulkUnpinAttributeDetailAttributeNameSpaceEnumStringValues() []string {
	return []string{
		"TRACES",
		"SYNTHETIC",
	}
}

// GetMappingBulkUnpinAttributeDetailAttributeNameSpaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkUnpinAttributeDetailAttributeNameSpaceEnum(val string) (BulkUnpinAttributeDetailAttributeNameSpaceEnum, bool) {
	enum, ok := mappingBulkUnpinAttributeDetailAttributeNameSpaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
