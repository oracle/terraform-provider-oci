// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSchedulingPolicyDetails Request to create Scheduling Policy resource with Maintenance windows per cadence.
type CreateSchedulingPolicyDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Scheduling Policy. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The cadence period.
	Cadence CreateSchedulingPolicyDetailsCadenceEnum `mandatory:"true" json:"cadence"`

	// Start of the month to be followed during the cadence period.
	CadenceStartMonth *Month `mandatory:"false" json:"cadenceStartMonth"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSchedulingPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSchedulingPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateSchedulingPolicyDetailsCadenceEnum(string(m.Cadence)); !ok && m.Cadence != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Cadence: %s. Supported values are: %s.", m.Cadence, strings.Join(GetCreateSchedulingPolicyDetailsCadenceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateSchedulingPolicyDetailsCadenceEnum Enum with underlying type: string
type CreateSchedulingPolicyDetailsCadenceEnum string

// Set of constants representing the allowable values for CreateSchedulingPolicyDetailsCadenceEnum
const (
	CreateSchedulingPolicyDetailsCadenceHalfyearly CreateSchedulingPolicyDetailsCadenceEnum = "HALFYEARLY"
	CreateSchedulingPolicyDetailsCadenceQuarterly  CreateSchedulingPolicyDetailsCadenceEnum = "QUARTERLY"
	CreateSchedulingPolicyDetailsCadenceMonthly    CreateSchedulingPolicyDetailsCadenceEnum = "MONTHLY"
)

var mappingCreateSchedulingPolicyDetailsCadenceEnum = map[string]CreateSchedulingPolicyDetailsCadenceEnum{
	"HALFYEARLY": CreateSchedulingPolicyDetailsCadenceHalfyearly,
	"QUARTERLY":  CreateSchedulingPolicyDetailsCadenceQuarterly,
	"MONTHLY":    CreateSchedulingPolicyDetailsCadenceMonthly,
}

var mappingCreateSchedulingPolicyDetailsCadenceEnumLowerCase = map[string]CreateSchedulingPolicyDetailsCadenceEnum{
	"halfyearly": CreateSchedulingPolicyDetailsCadenceHalfyearly,
	"quarterly":  CreateSchedulingPolicyDetailsCadenceQuarterly,
	"monthly":    CreateSchedulingPolicyDetailsCadenceMonthly,
}

// GetCreateSchedulingPolicyDetailsCadenceEnumValues Enumerates the set of values for CreateSchedulingPolicyDetailsCadenceEnum
func GetCreateSchedulingPolicyDetailsCadenceEnumValues() []CreateSchedulingPolicyDetailsCadenceEnum {
	values := make([]CreateSchedulingPolicyDetailsCadenceEnum, 0)
	for _, v := range mappingCreateSchedulingPolicyDetailsCadenceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSchedulingPolicyDetailsCadenceEnumStringValues Enumerates the set of values in String for CreateSchedulingPolicyDetailsCadenceEnum
func GetCreateSchedulingPolicyDetailsCadenceEnumStringValues() []string {
	return []string{
		"HALFYEARLY",
		"QUARTERLY",
		"MONTHLY",
	}
}

// GetMappingCreateSchedulingPolicyDetailsCadenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSchedulingPolicyDetailsCadenceEnum(val string) (CreateSchedulingPolicyDetailsCadenceEnum, bool) {
	enum, ok := mappingCreateSchedulingPolicyDetailsCadenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
