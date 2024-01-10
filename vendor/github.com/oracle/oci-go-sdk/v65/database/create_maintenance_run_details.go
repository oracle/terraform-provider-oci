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

// CreateMaintenanceRunDetails Details to schedule Maintenance Run with Latest Release Update along TimeZone File Update for the specified resource.
type CreateMaintenanceRunDetails struct {

	// The ID of the target resource for which the maintenance run should be created.
	TargetResourceId *string `mandatory:"true" json:"targetResourceId"`

	// The date and time that update should be scheduled.
	TimeScheduled *common.SDKTime `mandatory:"true" json:"timeScheduled"`

	// Patch type, either "QUARTERLY" or "TIMEZONE".
	PatchType CreateMaintenanceRunDetailsPatchTypeEnum `mandatory:"true" json:"patchType"`

	// Indicates if an automatic DST Time Zone file update is enabled for the Autonomous Container Database. If enabled along with Release Update, patching will be done in a Non-Rolling manner.
	IsDstFileUpdateEnabled *bool `mandatory:"false" json:"isDstFileUpdateEnabled"`

	// Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.
	// *IMPORTANT*: Non-rolling infrastructure patching involves system down time. See Oracle-Managed Infrastructure Maintenance Updates (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information.
	PatchingMode CreateMaintenanceRunDetailsPatchingModeEnum `mandatory:"false" json:"patchingMode,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the Maintenance Run.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m CreateMaintenanceRunDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMaintenanceRunDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateMaintenanceRunDetailsPatchTypeEnum(string(m.PatchType)); !ok && m.PatchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchType: %s. Supported values are: %s.", m.PatchType, strings.Join(GetCreateMaintenanceRunDetailsPatchTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCreateMaintenanceRunDetailsPatchingModeEnum(string(m.PatchingMode)); !ok && m.PatchingMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchingMode: %s. Supported values are: %s.", m.PatchingMode, strings.Join(GetCreateMaintenanceRunDetailsPatchingModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateMaintenanceRunDetailsPatchingModeEnum Enum with underlying type: string
type CreateMaintenanceRunDetailsPatchingModeEnum string

// Set of constants representing the allowable values for CreateMaintenanceRunDetailsPatchingModeEnum
const (
	CreateMaintenanceRunDetailsPatchingModeRolling    CreateMaintenanceRunDetailsPatchingModeEnum = "ROLLING"
	CreateMaintenanceRunDetailsPatchingModeNonrolling CreateMaintenanceRunDetailsPatchingModeEnum = "NONROLLING"
)

var mappingCreateMaintenanceRunDetailsPatchingModeEnum = map[string]CreateMaintenanceRunDetailsPatchingModeEnum{
	"ROLLING":    CreateMaintenanceRunDetailsPatchingModeRolling,
	"NONROLLING": CreateMaintenanceRunDetailsPatchingModeNonrolling,
}

var mappingCreateMaintenanceRunDetailsPatchingModeEnumLowerCase = map[string]CreateMaintenanceRunDetailsPatchingModeEnum{
	"rolling":    CreateMaintenanceRunDetailsPatchingModeRolling,
	"nonrolling": CreateMaintenanceRunDetailsPatchingModeNonrolling,
}

// GetCreateMaintenanceRunDetailsPatchingModeEnumValues Enumerates the set of values for CreateMaintenanceRunDetailsPatchingModeEnum
func GetCreateMaintenanceRunDetailsPatchingModeEnumValues() []CreateMaintenanceRunDetailsPatchingModeEnum {
	values := make([]CreateMaintenanceRunDetailsPatchingModeEnum, 0)
	for _, v := range mappingCreateMaintenanceRunDetailsPatchingModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateMaintenanceRunDetailsPatchingModeEnumStringValues Enumerates the set of values in String for CreateMaintenanceRunDetailsPatchingModeEnum
func GetCreateMaintenanceRunDetailsPatchingModeEnumStringValues() []string {
	return []string{
		"ROLLING",
		"NONROLLING",
	}
}

// GetMappingCreateMaintenanceRunDetailsPatchingModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateMaintenanceRunDetailsPatchingModeEnum(val string) (CreateMaintenanceRunDetailsPatchingModeEnum, bool) {
	enum, ok := mappingCreateMaintenanceRunDetailsPatchingModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateMaintenanceRunDetailsPatchTypeEnum Enum with underlying type: string
type CreateMaintenanceRunDetailsPatchTypeEnum string

// Set of constants representing the allowable values for CreateMaintenanceRunDetailsPatchTypeEnum
const (
	CreateMaintenanceRunDetailsPatchTypeQuarterly CreateMaintenanceRunDetailsPatchTypeEnum = "QUARTERLY"
	CreateMaintenanceRunDetailsPatchTypeTimezone  CreateMaintenanceRunDetailsPatchTypeEnum = "TIMEZONE"
)

var mappingCreateMaintenanceRunDetailsPatchTypeEnum = map[string]CreateMaintenanceRunDetailsPatchTypeEnum{
	"QUARTERLY": CreateMaintenanceRunDetailsPatchTypeQuarterly,
	"TIMEZONE":  CreateMaintenanceRunDetailsPatchTypeTimezone,
}

var mappingCreateMaintenanceRunDetailsPatchTypeEnumLowerCase = map[string]CreateMaintenanceRunDetailsPatchTypeEnum{
	"quarterly": CreateMaintenanceRunDetailsPatchTypeQuarterly,
	"timezone":  CreateMaintenanceRunDetailsPatchTypeTimezone,
}

// GetCreateMaintenanceRunDetailsPatchTypeEnumValues Enumerates the set of values for CreateMaintenanceRunDetailsPatchTypeEnum
func GetCreateMaintenanceRunDetailsPatchTypeEnumValues() []CreateMaintenanceRunDetailsPatchTypeEnum {
	values := make([]CreateMaintenanceRunDetailsPatchTypeEnum, 0)
	for _, v := range mappingCreateMaintenanceRunDetailsPatchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateMaintenanceRunDetailsPatchTypeEnumStringValues Enumerates the set of values in String for CreateMaintenanceRunDetailsPatchTypeEnum
func GetCreateMaintenanceRunDetailsPatchTypeEnumStringValues() []string {
	return []string{
		"QUARTERLY",
		"TIMEZONE",
	}
}

// GetMappingCreateMaintenanceRunDetailsPatchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateMaintenanceRunDetailsPatchTypeEnum(val string) (CreateMaintenanceRunDetailsPatchTypeEnum, bool) {
	enum, ok := mappingCreateMaintenanceRunDetailsPatchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
