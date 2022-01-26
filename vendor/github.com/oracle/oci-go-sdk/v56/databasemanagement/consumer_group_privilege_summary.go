// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ConsumerGroupPrivilegeSummary Summary of consumerGroupPrivileges.
type ConsumerGroupPrivilegeSummary struct {

	// The name of granted consumer group.
	Name *string `mandatory:"false" json:"name"`

	// Indicates whether the grant was with the GRANT option (YES) or not (NO).
	GrantOption ConsumerGroupPrivilegeSummaryGrantOptionEnum `mandatory:"false" json:"grantOption,omitempty"`

	// Indicates whether the consumer group is designated as the default for this user or role (YES) or not (NO)
	InitialGroup ConsumerGroupPrivilegeSummaryInitialGroupEnum `mandatory:"false" json:"initialGroup,omitempty"`
}

func (m ConsumerGroupPrivilegeSummary) String() string {
	return common.PointerString(m)
}

// ConsumerGroupPrivilegeSummaryGrantOptionEnum Enum with underlying type: string
type ConsumerGroupPrivilegeSummaryGrantOptionEnum string

// Set of constants representing the allowable values for ConsumerGroupPrivilegeSummaryGrantOptionEnum
const (
	ConsumerGroupPrivilegeSummaryGrantOptionYes ConsumerGroupPrivilegeSummaryGrantOptionEnum = "YES"
	ConsumerGroupPrivilegeSummaryGrantOptionNo  ConsumerGroupPrivilegeSummaryGrantOptionEnum = "NO"
)

var mappingConsumerGroupPrivilegeSummaryGrantOption = map[string]ConsumerGroupPrivilegeSummaryGrantOptionEnum{
	"YES": ConsumerGroupPrivilegeSummaryGrantOptionYes,
	"NO":  ConsumerGroupPrivilegeSummaryGrantOptionNo,
}

// GetConsumerGroupPrivilegeSummaryGrantOptionEnumValues Enumerates the set of values for ConsumerGroupPrivilegeSummaryGrantOptionEnum
func GetConsumerGroupPrivilegeSummaryGrantOptionEnumValues() []ConsumerGroupPrivilegeSummaryGrantOptionEnum {
	values := make([]ConsumerGroupPrivilegeSummaryGrantOptionEnum, 0)
	for _, v := range mappingConsumerGroupPrivilegeSummaryGrantOption {
		values = append(values, v)
	}
	return values
}

// ConsumerGroupPrivilegeSummaryInitialGroupEnum Enum with underlying type: string
type ConsumerGroupPrivilegeSummaryInitialGroupEnum string

// Set of constants representing the allowable values for ConsumerGroupPrivilegeSummaryInitialGroupEnum
const (
	ConsumerGroupPrivilegeSummaryInitialGroupYes ConsumerGroupPrivilegeSummaryInitialGroupEnum = "YES"
	ConsumerGroupPrivilegeSummaryInitialGroupNo  ConsumerGroupPrivilegeSummaryInitialGroupEnum = "NO"
)

var mappingConsumerGroupPrivilegeSummaryInitialGroup = map[string]ConsumerGroupPrivilegeSummaryInitialGroupEnum{
	"YES": ConsumerGroupPrivilegeSummaryInitialGroupYes,
	"NO":  ConsumerGroupPrivilegeSummaryInitialGroupNo,
}

// GetConsumerGroupPrivilegeSummaryInitialGroupEnumValues Enumerates the set of values for ConsumerGroupPrivilegeSummaryInitialGroupEnum
func GetConsumerGroupPrivilegeSummaryInitialGroupEnumValues() []ConsumerGroupPrivilegeSummaryInitialGroupEnum {
	values := make([]ConsumerGroupPrivilegeSummaryInitialGroupEnum, 0)
	for _, v := range mappingConsumerGroupPrivilegeSummaryInitialGroup {
		values = append(values, v)
	}
	return values
}
