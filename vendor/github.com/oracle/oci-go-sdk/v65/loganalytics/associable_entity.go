// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// AssociableEntity Entity details including whether or not it is eligible for association with the source.
type AssociableEntity struct {

	// The entity OCID.
	EntityId *string `mandatory:"false" json:"entityId"`

	// The name of the entity.
	EntityName *string `mandatory:"false" json:"entityName"`

	// The type name of the entity.
	EntityTypeName *string `mandatory:"false" json:"entityTypeName"`

	// The display name of the entity type.
	EntityTypeDisplayName *string `mandatory:"false" json:"entityTypeDisplayName"`

	// The entity host.
	Host *string `mandatory:"false" json:"host"`

	// The OCID of the Management Agent.
	AgentId *string `mandatory:"false" json:"agentId"`

	// This field indicates whether the entity is (in)eligible to be associated with this source.
	EligibilityStatus AssociableEntityEligibilityStatusEnum `mandatory:"false" json:"eligibilityStatus,omitempty"`

	// The reason the entity is not eligible for association.
	IneligibilityDetails *string `mandatory:"false" json:"ineligibilityDetails"`
}

func (m AssociableEntity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociableEntity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAssociableEntityEligibilityStatusEnum(string(m.EligibilityStatus)); !ok && m.EligibilityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EligibilityStatus: %s. Supported values are: %s.", m.EligibilityStatus, strings.Join(GetAssociableEntityEligibilityStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AssociableEntityEligibilityStatusEnum Enum with underlying type: string
type AssociableEntityEligibilityStatusEnum string

// Set of constants representing the allowable values for AssociableEntityEligibilityStatusEnum
const (
	AssociableEntityEligibilityStatusEligible   AssociableEntityEligibilityStatusEnum = "ELIGIBLE"
	AssociableEntityEligibilityStatusIneligible AssociableEntityEligibilityStatusEnum = "INELIGIBLE"
)

var mappingAssociableEntityEligibilityStatusEnum = map[string]AssociableEntityEligibilityStatusEnum{
	"ELIGIBLE":   AssociableEntityEligibilityStatusEligible,
	"INELIGIBLE": AssociableEntityEligibilityStatusIneligible,
}

var mappingAssociableEntityEligibilityStatusEnumLowerCase = map[string]AssociableEntityEligibilityStatusEnum{
	"eligible":   AssociableEntityEligibilityStatusEligible,
	"ineligible": AssociableEntityEligibilityStatusIneligible,
}

// GetAssociableEntityEligibilityStatusEnumValues Enumerates the set of values for AssociableEntityEligibilityStatusEnum
func GetAssociableEntityEligibilityStatusEnumValues() []AssociableEntityEligibilityStatusEnum {
	values := make([]AssociableEntityEligibilityStatusEnum, 0)
	for _, v := range mappingAssociableEntityEligibilityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAssociableEntityEligibilityStatusEnumStringValues Enumerates the set of values in String for AssociableEntityEligibilityStatusEnum
func GetAssociableEntityEligibilityStatusEnumStringValues() []string {
	return []string{
		"ELIGIBLE",
		"INELIGIBLE",
	}
}

// GetMappingAssociableEntityEligibilityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssociableEntityEligibilityStatusEnum(val string) (AssociableEntityEligibilityStatusEnum, bool) {
	enum, ok := mappingAssociableEntityEligibilityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
