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

// BulkUpdateAttributeNotesDetail Object that contains the details about a single attribute in the bulk request for which notes are to be updated.
type BulkUpdateAttributeNotesDetail struct {

	// Name of the attribute for which notes are to be updated.
	AttributeName *string `mandatory:"true" json:"attributeName"`

	// Notes to be updated.  The size of notes cannot exceed 1000 chars.
	Notes *string `mandatory:"true" json:"notes"`

	// Namespace of the attribute for which the notes are to be updated.  The attributeNameSpace will default to TRACES if it is
	// not passed in.
	AttributeNameSpace BulkUpdateAttributeNotesDetailAttributeNameSpaceEnum `mandatory:"false" json:"attributeNameSpace,omitempty"`
}

func (m BulkUpdateAttributeNotesDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkUpdateAttributeNotesDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBulkUpdateAttributeNotesDetailAttributeNameSpaceEnum(string(m.AttributeNameSpace)); !ok && m.AttributeNameSpace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeNameSpace: %s. Supported values are: %s.", m.AttributeNameSpace, strings.Join(GetBulkUpdateAttributeNotesDetailAttributeNameSpaceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkUpdateAttributeNotesDetailAttributeNameSpaceEnum Enum with underlying type: string
type BulkUpdateAttributeNotesDetailAttributeNameSpaceEnum string

// Set of constants representing the allowable values for BulkUpdateAttributeNotesDetailAttributeNameSpaceEnum
const (
	BulkUpdateAttributeNotesDetailAttributeNameSpaceTraces    BulkUpdateAttributeNotesDetailAttributeNameSpaceEnum = "TRACES"
	BulkUpdateAttributeNotesDetailAttributeNameSpaceSynthetic BulkUpdateAttributeNotesDetailAttributeNameSpaceEnum = "SYNTHETIC"
)

var mappingBulkUpdateAttributeNotesDetailAttributeNameSpaceEnum = map[string]BulkUpdateAttributeNotesDetailAttributeNameSpaceEnum{
	"TRACES":    BulkUpdateAttributeNotesDetailAttributeNameSpaceTraces,
	"SYNTHETIC": BulkUpdateAttributeNotesDetailAttributeNameSpaceSynthetic,
}

var mappingBulkUpdateAttributeNotesDetailAttributeNameSpaceEnumLowerCase = map[string]BulkUpdateAttributeNotesDetailAttributeNameSpaceEnum{
	"traces":    BulkUpdateAttributeNotesDetailAttributeNameSpaceTraces,
	"synthetic": BulkUpdateAttributeNotesDetailAttributeNameSpaceSynthetic,
}

// GetBulkUpdateAttributeNotesDetailAttributeNameSpaceEnumValues Enumerates the set of values for BulkUpdateAttributeNotesDetailAttributeNameSpaceEnum
func GetBulkUpdateAttributeNotesDetailAttributeNameSpaceEnumValues() []BulkUpdateAttributeNotesDetailAttributeNameSpaceEnum {
	values := make([]BulkUpdateAttributeNotesDetailAttributeNameSpaceEnum, 0)
	for _, v := range mappingBulkUpdateAttributeNotesDetailAttributeNameSpaceEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkUpdateAttributeNotesDetailAttributeNameSpaceEnumStringValues Enumerates the set of values in String for BulkUpdateAttributeNotesDetailAttributeNameSpaceEnum
func GetBulkUpdateAttributeNotesDetailAttributeNameSpaceEnumStringValues() []string {
	return []string{
		"TRACES",
		"SYNTHETIC",
	}
}

// GetMappingBulkUpdateAttributeNotesDetailAttributeNameSpaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkUpdateAttributeNotesDetailAttributeNameSpaceEnum(val string) (BulkUpdateAttributeNotesDetailAttributeNameSpaceEnum, bool) {
	enum, ok := mappingBulkUpdateAttributeNotesDetailAttributeNameSpaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
