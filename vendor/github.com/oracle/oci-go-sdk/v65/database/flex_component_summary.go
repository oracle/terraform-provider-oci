// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FlexComponentSummary The Flex Components for a DB system. The Flex Component determines resources to allocate to the DB system -  CPU cores, memory and storage for Flex shapes.
// For Exadata flexible shapes, detailed specifications can be found in https://docs.oracle.com/en/engineered-systems/exadata-cloud-service/ecscm/exa-service-desc.html#GUID-9E090174-5C57-4EB1-9243-B470F9F10D6B
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator.
// If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type FlexComponentSummary struct {

	// The name of the Flex Component used for the DB system.
	Name *string `mandatory:"true" json:"name"`

	// The minimum number of CPU cores that can be enabled on the DB Server for this Flex Component.
	MinimumCoreCount *int `mandatory:"false" json:"minimumCoreCount"`

	// The maximum number of CPU cores that can ben enabled on the DB Server for this Flex Component.
	AvailableCoreCount *int `mandatory:"false" json:"availableCoreCount"`

	// The maximum  storage that can be enabled on the Storage Server for this Flex Component.
	AvailableDbStorageInGBs *int `mandatory:"false" json:"availableDbStorageInGBs"`

	// The runtime minimum number of CPU cores that can be enabled for this Flex Component.
	RuntimeMinimumCoreCount *int `mandatory:"false" json:"runtimeMinimumCoreCount"`

	// The name of the DB system shape for this Flex Component.
	Shape *string `mandatory:"false" json:"shape"`

	// The maximum memory size that can be enabled on the DB Server for this Flex Component.
	AvailableMemoryInGBs *int `mandatory:"false" json:"availableMemoryInGBs"`

	// The maximum local storage that can be enabled on the DB Server for this Flex Component.
	AvailableLocalStorageInGBs *int `mandatory:"false" json:"availableLocalStorageInGBs"`

	// The compute model of the DB Server for this Flex Component.
	ComputeModel *string `mandatory:"false" json:"computeModel"`

	// The hardware type of the DB (Compute) or Storage (Cell) Server for this Flex Component.
	HardwareType FlexComponentSummaryHardwareTypeEnum `mandatory:"false" json:"hardwareType,omitempty"`

	// The description summary for this Flex Component.
	DescriptionSummary *string `mandatory:"false" json:"descriptionSummary"`
}

func (m FlexComponentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FlexComponentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFlexComponentSummaryHardwareTypeEnum(string(m.HardwareType)); !ok && m.HardwareType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HardwareType: %s. Supported values are: %s.", m.HardwareType, strings.Join(GetFlexComponentSummaryHardwareTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FlexComponentSummaryHardwareTypeEnum Enum with underlying type: string
type FlexComponentSummaryHardwareTypeEnum string

// Set of constants representing the allowable values for FlexComponentSummaryHardwareTypeEnum
const (
	FlexComponentSummaryHardwareTypeCompute FlexComponentSummaryHardwareTypeEnum = "COMPUTE"
	FlexComponentSummaryHardwareTypeCell    FlexComponentSummaryHardwareTypeEnum = "CELL"
)

var mappingFlexComponentSummaryHardwareTypeEnum = map[string]FlexComponentSummaryHardwareTypeEnum{
	"COMPUTE": FlexComponentSummaryHardwareTypeCompute,
	"CELL":    FlexComponentSummaryHardwareTypeCell,
}

var mappingFlexComponentSummaryHardwareTypeEnumLowerCase = map[string]FlexComponentSummaryHardwareTypeEnum{
	"compute": FlexComponentSummaryHardwareTypeCompute,
	"cell":    FlexComponentSummaryHardwareTypeCell,
}

// GetFlexComponentSummaryHardwareTypeEnumValues Enumerates the set of values for FlexComponentSummaryHardwareTypeEnum
func GetFlexComponentSummaryHardwareTypeEnumValues() []FlexComponentSummaryHardwareTypeEnum {
	values := make([]FlexComponentSummaryHardwareTypeEnum, 0)
	for _, v := range mappingFlexComponentSummaryHardwareTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFlexComponentSummaryHardwareTypeEnumStringValues Enumerates the set of values in String for FlexComponentSummaryHardwareTypeEnum
func GetFlexComponentSummaryHardwareTypeEnumStringValues() []string {
	return []string{
		"COMPUTE",
		"CELL",
	}
}

// GetMappingFlexComponentSummaryHardwareTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFlexComponentSummaryHardwareTypeEnum(val string) (FlexComponentSummaryHardwareTypeEnum, bool) {
	enum, ok := mappingFlexComponentSummaryHardwareTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
