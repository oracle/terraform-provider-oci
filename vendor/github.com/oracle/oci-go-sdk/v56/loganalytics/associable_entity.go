// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

// AssociableEntityEligibilityStatusEnum Enum with underlying type: string
type AssociableEntityEligibilityStatusEnum string

// Set of constants representing the allowable values for AssociableEntityEligibilityStatusEnum
const (
	AssociableEntityEligibilityStatusEligible   AssociableEntityEligibilityStatusEnum = "ELIGIBLE"
	AssociableEntityEligibilityStatusIneligible AssociableEntityEligibilityStatusEnum = "INELIGIBLE"
)

var mappingAssociableEntityEligibilityStatus = map[string]AssociableEntityEligibilityStatusEnum{
	"ELIGIBLE":   AssociableEntityEligibilityStatusEligible,
	"INELIGIBLE": AssociableEntityEligibilityStatusIneligible,
}

// GetAssociableEntityEligibilityStatusEnumValues Enumerates the set of values for AssociableEntityEligibilityStatusEnum
func GetAssociableEntityEligibilityStatusEnumValues() []AssociableEntityEligibilityStatusEnum {
	values := make([]AssociableEntityEligibilityStatusEnum, 0)
	for _, v := range mappingAssociableEntityEligibilityStatus {
		values = append(values, v)
	}
	return values
}
