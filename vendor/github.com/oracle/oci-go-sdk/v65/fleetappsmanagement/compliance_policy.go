// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CompliancePolicy Define software patch compliance policies for various products running in OCI resources.
// A compliance policy is a configuration you set up for various products to report compliance by defining the schedule and patch baseline
type CompliancePolicy struct {

	// The OCID of the CompliancePolicy.
	Id *string `mandatory:"true" json:"id"`

	// Display name for the CompliancePolicy.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// platformConfiguration OCID corresponding to the Product.
	ProductId *string `mandatory:"true" json:"productId"`

	// The OCID of the compartment the CompliancePolicy belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the CompliancePolicy was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the CompliancePolicy.
	LifecycleState CompliancePolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The date and time the CompliancePolicy was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the CompliancePolicy in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CompliancePolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CompliancePolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCompliancePolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCompliancePolicyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CompliancePolicyLifecycleStateEnum Enum with underlying type: string
type CompliancePolicyLifecycleStateEnum string

// Set of constants representing the allowable values for CompliancePolicyLifecycleStateEnum
const (
	CompliancePolicyLifecycleStateCreating CompliancePolicyLifecycleStateEnum = "CREATING"
	CompliancePolicyLifecycleStateUpdating CompliancePolicyLifecycleStateEnum = "UPDATING"
	CompliancePolicyLifecycleStateActive   CompliancePolicyLifecycleStateEnum = "ACTIVE"
	CompliancePolicyLifecycleStateDeleting CompliancePolicyLifecycleStateEnum = "DELETING"
	CompliancePolicyLifecycleStateDeleted  CompliancePolicyLifecycleStateEnum = "DELETED"
	CompliancePolicyLifecycleStateFailed   CompliancePolicyLifecycleStateEnum = "FAILED"
)

var mappingCompliancePolicyLifecycleStateEnum = map[string]CompliancePolicyLifecycleStateEnum{
	"CREATING": CompliancePolicyLifecycleStateCreating,
	"UPDATING": CompliancePolicyLifecycleStateUpdating,
	"ACTIVE":   CompliancePolicyLifecycleStateActive,
	"DELETING": CompliancePolicyLifecycleStateDeleting,
	"DELETED":  CompliancePolicyLifecycleStateDeleted,
	"FAILED":   CompliancePolicyLifecycleStateFailed,
}

var mappingCompliancePolicyLifecycleStateEnumLowerCase = map[string]CompliancePolicyLifecycleStateEnum{
	"creating": CompliancePolicyLifecycleStateCreating,
	"updating": CompliancePolicyLifecycleStateUpdating,
	"active":   CompliancePolicyLifecycleStateActive,
	"deleting": CompliancePolicyLifecycleStateDeleting,
	"deleted":  CompliancePolicyLifecycleStateDeleted,
	"failed":   CompliancePolicyLifecycleStateFailed,
}

// GetCompliancePolicyLifecycleStateEnumValues Enumerates the set of values for CompliancePolicyLifecycleStateEnum
func GetCompliancePolicyLifecycleStateEnumValues() []CompliancePolicyLifecycleStateEnum {
	values := make([]CompliancePolicyLifecycleStateEnum, 0)
	for _, v := range mappingCompliancePolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCompliancePolicyLifecycleStateEnumStringValues Enumerates the set of values in String for CompliancePolicyLifecycleStateEnum
func GetCompliancePolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCompliancePolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCompliancePolicyLifecycleStateEnum(val string) (CompliancePolicyLifecycleStateEnum, bool) {
	enum, ok := mappingCompliancePolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
