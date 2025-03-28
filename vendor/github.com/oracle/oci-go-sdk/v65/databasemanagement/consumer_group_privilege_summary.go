// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConsumerGroupPrivilegeSummary A summary of consumer group privileges.
type ConsumerGroupPrivilegeSummary struct {

	// The name of the granted consumer group privilege.
	Name *string `mandatory:"false" json:"name"`

	// Indicates whether the privilege is granted with the GRANT option (YES) or not (NO).
	GrantOption ConsumerGroupPrivilegeSummaryGrantOptionEnum `mandatory:"false" json:"grantOption,omitempty"`

	// Indicates whether the consumer group is designated as the default for this user or role (YES) or not (NO).
	InitialGroup ConsumerGroupPrivilegeSummaryInitialGroupEnum `mandatory:"false" json:"initialGroup,omitempty"`
}

func (m ConsumerGroupPrivilegeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConsumerGroupPrivilegeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConsumerGroupPrivilegeSummaryGrantOptionEnum(string(m.GrantOption)); !ok && m.GrantOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GrantOption: %s. Supported values are: %s.", m.GrantOption, strings.Join(GetConsumerGroupPrivilegeSummaryGrantOptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConsumerGroupPrivilegeSummaryInitialGroupEnum(string(m.InitialGroup)); !ok && m.InitialGroup != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InitialGroup: %s. Supported values are: %s.", m.InitialGroup, strings.Join(GetConsumerGroupPrivilegeSummaryInitialGroupEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConsumerGroupPrivilegeSummaryGrantOptionEnum Enum with underlying type: string
type ConsumerGroupPrivilegeSummaryGrantOptionEnum string

// Set of constants representing the allowable values for ConsumerGroupPrivilegeSummaryGrantOptionEnum
const (
	ConsumerGroupPrivilegeSummaryGrantOptionYes ConsumerGroupPrivilegeSummaryGrantOptionEnum = "YES"
	ConsumerGroupPrivilegeSummaryGrantOptionNo  ConsumerGroupPrivilegeSummaryGrantOptionEnum = "NO"
)

var mappingConsumerGroupPrivilegeSummaryGrantOptionEnum = map[string]ConsumerGroupPrivilegeSummaryGrantOptionEnum{
	"YES": ConsumerGroupPrivilegeSummaryGrantOptionYes,
	"NO":  ConsumerGroupPrivilegeSummaryGrantOptionNo,
}

var mappingConsumerGroupPrivilegeSummaryGrantOptionEnumLowerCase = map[string]ConsumerGroupPrivilegeSummaryGrantOptionEnum{
	"yes": ConsumerGroupPrivilegeSummaryGrantOptionYes,
	"no":  ConsumerGroupPrivilegeSummaryGrantOptionNo,
}

// GetConsumerGroupPrivilegeSummaryGrantOptionEnumValues Enumerates the set of values for ConsumerGroupPrivilegeSummaryGrantOptionEnum
func GetConsumerGroupPrivilegeSummaryGrantOptionEnumValues() []ConsumerGroupPrivilegeSummaryGrantOptionEnum {
	values := make([]ConsumerGroupPrivilegeSummaryGrantOptionEnum, 0)
	for _, v := range mappingConsumerGroupPrivilegeSummaryGrantOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetConsumerGroupPrivilegeSummaryGrantOptionEnumStringValues Enumerates the set of values in String for ConsumerGroupPrivilegeSummaryGrantOptionEnum
func GetConsumerGroupPrivilegeSummaryGrantOptionEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingConsumerGroupPrivilegeSummaryGrantOptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConsumerGroupPrivilegeSummaryGrantOptionEnum(val string) (ConsumerGroupPrivilegeSummaryGrantOptionEnum, bool) {
	enum, ok := mappingConsumerGroupPrivilegeSummaryGrantOptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConsumerGroupPrivilegeSummaryInitialGroupEnum Enum with underlying type: string
type ConsumerGroupPrivilegeSummaryInitialGroupEnum string

// Set of constants representing the allowable values for ConsumerGroupPrivilegeSummaryInitialGroupEnum
const (
	ConsumerGroupPrivilegeSummaryInitialGroupYes ConsumerGroupPrivilegeSummaryInitialGroupEnum = "YES"
	ConsumerGroupPrivilegeSummaryInitialGroupNo  ConsumerGroupPrivilegeSummaryInitialGroupEnum = "NO"
)

var mappingConsumerGroupPrivilegeSummaryInitialGroupEnum = map[string]ConsumerGroupPrivilegeSummaryInitialGroupEnum{
	"YES": ConsumerGroupPrivilegeSummaryInitialGroupYes,
	"NO":  ConsumerGroupPrivilegeSummaryInitialGroupNo,
}

var mappingConsumerGroupPrivilegeSummaryInitialGroupEnumLowerCase = map[string]ConsumerGroupPrivilegeSummaryInitialGroupEnum{
	"yes": ConsumerGroupPrivilegeSummaryInitialGroupYes,
	"no":  ConsumerGroupPrivilegeSummaryInitialGroupNo,
}

// GetConsumerGroupPrivilegeSummaryInitialGroupEnumValues Enumerates the set of values for ConsumerGroupPrivilegeSummaryInitialGroupEnum
func GetConsumerGroupPrivilegeSummaryInitialGroupEnumValues() []ConsumerGroupPrivilegeSummaryInitialGroupEnum {
	values := make([]ConsumerGroupPrivilegeSummaryInitialGroupEnum, 0)
	for _, v := range mappingConsumerGroupPrivilegeSummaryInitialGroupEnum {
		values = append(values, v)
	}
	return values
}

// GetConsumerGroupPrivilegeSummaryInitialGroupEnumStringValues Enumerates the set of values in String for ConsumerGroupPrivilegeSummaryInitialGroupEnum
func GetConsumerGroupPrivilegeSummaryInitialGroupEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingConsumerGroupPrivilegeSummaryInitialGroupEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConsumerGroupPrivilegeSummaryInitialGroupEnum(val string) (ConsumerGroupPrivilegeSummaryInitialGroupEnum, bool) {
	enum, ok := mappingConsumerGroupPrivilegeSummaryInitialGroupEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
