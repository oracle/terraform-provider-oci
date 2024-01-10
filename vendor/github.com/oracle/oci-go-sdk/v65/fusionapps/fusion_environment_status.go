// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FusionEnvironmentStatus The health status of the Fusion Applications environment. For more information, see Environment Status (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/manage-environment.htm#environment-status).
type FusionEnvironmentStatus struct {

	// The data plane status of FusionEnvironment.
	Status FusionEnvironmentStatusStatusEnum `mandatory:"true" json:"status"`
}

func (m FusionEnvironmentStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FusionEnvironmentStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFusionEnvironmentStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetFusionEnvironmentStatusStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FusionEnvironmentStatusStatusEnum Enum with underlying type: string
type FusionEnvironmentStatusStatusEnum string

// Set of constants representing the allowable values for FusionEnvironmentStatusStatusEnum
const (
	FusionEnvironmentStatusStatusAvailable             FusionEnvironmentStatusStatusEnum = "AVAILABLE"
	FusionEnvironmentStatusStatusUnavailable           FusionEnvironmentStatusStatusEnum = "UNAVAILABLE"
	FusionEnvironmentStatusStatusNotApplicable         FusionEnvironmentStatusStatusEnum = "NOT_APPLICABLE"
	FusionEnvironmentStatusStatusMaintenanceInProgress FusionEnvironmentStatusStatusEnum = "MAINTENANCE_IN_PROGRESS"
	FusionEnvironmentStatusStatusRefreshInProgress     FusionEnvironmentStatusStatusEnum = "REFRESH_IN_PROGRESS"
	FusionEnvironmentStatusStatusUnknown               FusionEnvironmentStatusStatusEnum = "UNKNOWN"
)

var mappingFusionEnvironmentStatusStatusEnum = map[string]FusionEnvironmentStatusStatusEnum{
	"AVAILABLE":               FusionEnvironmentStatusStatusAvailable,
	"UNAVAILABLE":             FusionEnvironmentStatusStatusUnavailable,
	"NOT_APPLICABLE":          FusionEnvironmentStatusStatusNotApplicable,
	"MAINTENANCE_IN_PROGRESS": FusionEnvironmentStatusStatusMaintenanceInProgress,
	"REFRESH_IN_PROGRESS":     FusionEnvironmentStatusStatusRefreshInProgress,
	"UNKNOWN":                 FusionEnvironmentStatusStatusUnknown,
}

var mappingFusionEnvironmentStatusStatusEnumLowerCase = map[string]FusionEnvironmentStatusStatusEnum{
	"available":               FusionEnvironmentStatusStatusAvailable,
	"unavailable":             FusionEnvironmentStatusStatusUnavailable,
	"not_applicable":          FusionEnvironmentStatusStatusNotApplicable,
	"maintenance_in_progress": FusionEnvironmentStatusStatusMaintenanceInProgress,
	"refresh_in_progress":     FusionEnvironmentStatusStatusRefreshInProgress,
	"unknown":                 FusionEnvironmentStatusStatusUnknown,
}

// GetFusionEnvironmentStatusStatusEnumValues Enumerates the set of values for FusionEnvironmentStatusStatusEnum
func GetFusionEnvironmentStatusStatusEnumValues() []FusionEnvironmentStatusStatusEnum {
	values := make([]FusionEnvironmentStatusStatusEnum, 0)
	for _, v := range mappingFusionEnvironmentStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetFusionEnvironmentStatusStatusEnumStringValues Enumerates the set of values in String for FusionEnvironmentStatusStatusEnum
func GetFusionEnvironmentStatusStatusEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"UNAVAILABLE",
		"NOT_APPLICABLE",
		"MAINTENANCE_IN_PROGRESS",
		"REFRESH_IN_PROGRESS",
		"UNKNOWN",
	}
}

// GetMappingFusionEnvironmentStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFusionEnvironmentStatusStatusEnum(val string) (FusionEnvironmentStatusStatusEnum, bool) {
	enum, ok := mappingFusionEnvironmentStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
