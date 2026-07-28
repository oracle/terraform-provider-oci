// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RetentionDuration This field is used to define the retention period for backups. This is an optional field. If it is not specified, it is set to null, no retention period will be applied to the backups.
type RetentionDuration struct {

	// The value you can assign to the Time Unit property for this Duration may be either "YEARS" or "DAYS".
	RetentionTimeUnit RetentionDurationRetentionTimeUnitEnum `mandatory:"true" json:"retentionTimeUnit"`

	// The value to enter for the amount of retention time should be a numerical figure (such as 1, 7, 30, etc.) that corresponds to the period specified in the retention time unit property (such as YEARS, DAYS). The combination of these two properties determines the total length of the retention period.
	RetentionTimeAmount *int `mandatory:"true" json:"retentionTimeAmount"`
}

func (m RetentionDuration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RetentionDuration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRetentionDurationRetentionTimeUnitEnum(string(m.RetentionTimeUnit)); !ok && m.RetentionTimeUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RetentionTimeUnit: %s. Supported values are: %s.", m.RetentionTimeUnit, strings.Join(GetRetentionDurationRetentionTimeUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RetentionDurationRetentionTimeUnitEnum Enum with underlying type: string
type RetentionDurationRetentionTimeUnitEnum string

// Set of constants representing the allowable values for RetentionDurationRetentionTimeUnitEnum
const (
	RetentionDurationRetentionTimeUnitYears RetentionDurationRetentionTimeUnitEnum = "YEARS"
	RetentionDurationRetentionTimeUnitDays  RetentionDurationRetentionTimeUnitEnum = "DAYS"
)

var mappingRetentionDurationRetentionTimeUnitEnum = map[string]RetentionDurationRetentionTimeUnitEnum{
	"YEARS": RetentionDurationRetentionTimeUnitYears,
	"DAYS":  RetentionDurationRetentionTimeUnitDays,
}

var mappingRetentionDurationRetentionTimeUnitEnumLowerCase = map[string]RetentionDurationRetentionTimeUnitEnum{
	"years": RetentionDurationRetentionTimeUnitYears,
	"days":  RetentionDurationRetentionTimeUnitDays,
}

// GetRetentionDurationRetentionTimeUnitEnumValues Enumerates the set of values for RetentionDurationRetentionTimeUnitEnum
func GetRetentionDurationRetentionTimeUnitEnumValues() []RetentionDurationRetentionTimeUnitEnum {
	values := make([]RetentionDurationRetentionTimeUnitEnum, 0)
	for _, v := range mappingRetentionDurationRetentionTimeUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetRetentionDurationRetentionTimeUnitEnumStringValues Enumerates the set of values in String for RetentionDurationRetentionTimeUnitEnum
func GetRetentionDurationRetentionTimeUnitEnumStringValues() []string {
	return []string{
		"YEARS",
		"DAYS",
	}
}

// GetMappingRetentionDurationRetentionTimeUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRetentionDurationRetentionTimeUnitEnum(val string) (RetentionDurationRetentionTimeUnitEnum, bool) {
	enum, ok := mappingRetentionDurationRetentionTimeUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
