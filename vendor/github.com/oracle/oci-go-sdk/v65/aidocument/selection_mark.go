// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SelectionMark A single checkbox with selection mark.
type SelectionMark struct {

	// String to display if checkbox is selected or not selected.
	State SelectionMarkStateEnum `mandatory:"true" json:"state"`

	// The confidence score, a float value between 0 and 1.
	Confidence *float32 `mandatory:"true" json:"confidence"`

	BoundingPolygon *BoundingPolygon `mandatory:"true" json:"boundingPolygon"`
}

func (m SelectionMark) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SelectionMark) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSelectionMarkStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetSelectionMarkStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SelectionMarkStateEnum Enum with underlying type: string
type SelectionMarkStateEnum string

// Set of constants representing the allowable values for SelectionMarkStateEnum
const (
	SelectionMarkStateUnselected SelectionMarkStateEnum = "UNSELECTED"
	SelectionMarkStateSelected   SelectionMarkStateEnum = "SELECTED"
)

var mappingSelectionMarkStateEnum = map[string]SelectionMarkStateEnum{
	"UNSELECTED": SelectionMarkStateUnselected,
	"SELECTED":   SelectionMarkStateSelected,
}

var mappingSelectionMarkStateEnumLowerCase = map[string]SelectionMarkStateEnum{
	"unselected": SelectionMarkStateUnselected,
	"selected":   SelectionMarkStateSelected,
}

// GetSelectionMarkStateEnumValues Enumerates the set of values for SelectionMarkStateEnum
func GetSelectionMarkStateEnumValues() []SelectionMarkStateEnum {
	values := make([]SelectionMarkStateEnum, 0)
	for _, v := range mappingSelectionMarkStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSelectionMarkStateEnumStringValues Enumerates the set of values in String for SelectionMarkStateEnum
func GetSelectionMarkStateEnumStringValues() []string {
	return []string{
		"UNSELECTED",
		"SELECTED",
	}
}

// GetMappingSelectionMarkStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSelectionMarkStateEnum(val string) (SelectionMarkStateEnum, bool) {
	enum, ok := mappingSelectionMarkStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
