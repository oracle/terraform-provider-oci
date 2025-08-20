// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Link The model for links.
type Link struct {

	// Reference links to the previous page, next page, and other pages.
	Rel LinkRelEnum `mandatory:"true" json:"rel"`

	// The anchor tag.
	Href *string `mandatory:"true" json:"href"`
}

func (m Link) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Link) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLinkRelEnum(string(m.Rel)); !ok && m.Rel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Rel: %s. Supported values are: %s.", m.Rel, strings.Join(GetLinkRelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LinkRelEnum Enum with underlying type: string
type LinkRelEnum string

// Set of constants representing the allowable values for LinkRelEnum
const (
	LinkRelSelf      LinkRelEnum = "SELF"
	LinkRelCanonical LinkRelEnum = "CANONICAL"
	LinkRelNext      LinkRelEnum = "NEXT"
	LinkRelTemplate  LinkRelEnum = "TEMPLATE"
	LinkRelPrev      LinkRelEnum = "PREV"
)

var mappingLinkRelEnum = map[string]LinkRelEnum{
	"SELF":      LinkRelSelf,
	"CANONICAL": LinkRelCanonical,
	"NEXT":      LinkRelNext,
	"TEMPLATE":  LinkRelTemplate,
	"PREV":      LinkRelPrev,
}

var mappingLinkRelEnumLowerCase = map[string]LinkRelEnum{
	"self":      LinkRelSelf,
	"canonical": LinkRelCanonical,
	"next":      LinkRelNext,
	"template":  LinkRelTemplate,
	"prev":      LinkRelPrev,
}

// GetLinkRelEnumValues Enumerates the set of values for LinkRelEnum
func GetLinkRelEnumValues() []LinkRelEnum {
	values := make([]LinkRelEnum, 0)
	for _, v := range mappingLinkRelEnum {
		values = append(values, v)
	}
	return values
}

// GetLinkRelEnumStringValues Enumerates the set of values in String for LinkRelEnum
func GetLinkRelEnumStringValues() []string {
	return []string{
		"SELF",
		"CANONICAL",
		"NEXT",
		"TEMPLATE",
		"PREV",
	}
}

// GetMappingLinkRelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLinkRelEnum(val string) (LinkRelEnum, bool) {
	enum, ok := mappingLinkRelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
