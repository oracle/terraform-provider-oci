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

// ComplianceRecord The compliance status reports.
type ComplianceRecord struct {

	// The OCID of the ComplianceRecord.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the entity for which the compliance is calculated.Ex.FleetId
	EntityId *string `mandatory:"true" json:"entityId"`

	// The displayName of the entity for which the compliance is calculated.Ex.DisplayName for the Fleet
	EntityDisplayName *string `mandatory:"true" json:"entityDisplayName"`

	Resource *ComplianceDetailResource `mandatory:"true" json:"resource"`

	Target *ComplianceDetailTarget `mandatory:"true" json:"target"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	Policy *ComplianceDetailPolicy `mandatory:"false" json:"policy"`

	Patch *CompliancePatchDetail `mandatory:"false" json:"patch"`

	// Last known compliance state of target.
	ComplianceState ComplianceStateEnum `mandatory:"false" json:"complianceState,omitempty"`

	// The date and time the ComplianceRecord was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the ComplianceRecord was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the ComplianceRecord.
	LifecycleState ComplianceRecordLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m ComplianceRecord) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComplianceRecord) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingComplianceStateEnum(string(m.ComplianceState)); !ok && m.ComplianceState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComplianceState: %s. Supported values are: %s.", m.ComplianceState, strings.Join(GetComplianceStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingComplianceRecordLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetComplianceRecordLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComplianceRecordLifecycleStateEnum Enum with underlying type: string
type ComplianceRecordLifecycleStateEnum string

// Set of constants representing the allowable values for ComplianceRecordLifecycleStateEnum
const (
	ComplianceRecordLifecycleStateActive  ComplianceRecordLifecycleStateEnum = "ACTIVE"
	ComplianceRecordLifecycleStateDeleted ComplianceRecordLifecycleStateEnum = "DELETED"
	ComplianceRecordLifecycleStateFailed  ComplianceRecordLifecycleStateEnum = "FAILED"
)

var mappingComplianceRecordLifecycleStateEnum = map[string]ComplianceRecordLifecycleStateEnum{
	"ACTIVE":  ComplianceRecordLifecycleStateActive,
	"DELETED": ComplianceRecordLifecycleStateDeleted,
	"FAILED":  ComplianceRecordLifecycleStateFailed,
}

var mappingComplianceRecordLifecycleStateEnumLowerCase = map[string]ComplianceRecordLifecycleStateEnum{
	"active":  ComplianceRecordLifecycleStateActive,
	"deleted": ComplianceRecordLifecycleStateDeleted,
	"failed":  ComplianceRecordLifecycleStateFailed,
}

// GetComplianceRecordLifecycleStateEnumValues Enumerates the set of values for ComplianceRecordLifecycleStateEnum
func GetComplianceRecordLifecycleStateEnumValues() []ComplianceRecordLifecycleStateEnum {
	values := make([]ComplianceRecordLifecycleStateEnum, 0)
	for _, v := range mappingComplianceRecordLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetComplianceRecordLifecycleStateEnumStringValues Enumerates the set of values in String for ComplianceRecordLifecycleStateEnum
func GetComplianceRecordLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"FAILED",
	}
}

// GetMappingComplianceRecordLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComplianceRecordLifecycleStateEnum(val string) (ComplianceRecordLifecycleStateEnum, bool) {
	enum, ok := mappingComplianceRecordLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
