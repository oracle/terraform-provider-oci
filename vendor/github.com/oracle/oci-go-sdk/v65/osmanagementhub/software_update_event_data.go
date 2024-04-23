// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SoftwareUpdateEventData Provides additional information for a software update event.
type SoftwareUpdateEventData struct {

	// Type of software update operation.
	OperationType SoftwareUpdateEventDataOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of the software update.
	Status EventStatusEnum `mandatory:"true" json:"status"`

	AdditionalDetails *WorkRequestEventDataAdditionalDetails `mandatory:"false" json:"additionalDetails"`
}

func (m SoftwareUpdateEventData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoftwareUpdateEventData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSoftwareUpdateEventDataOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetSoftwareUpdateEventDataOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEventStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetEventStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SoftwareUpdateEventDataOperationTypeEnum Enum with underlying type: string
type SoftwareUpdateEventDataOperationTypeEnum string

// Set of constants representing the allowable values for SoftwareUpdateEventDataOperationTypeEnum
const (
	SoftwareUpdateEventDataOperationTypeUpdateAllPackages SoftwareUpdateEventDataOperationTypeEnum = "UPDATE_ALL_PACKAGES"
	SoftwareUpdateEventDataOperationTypeInstallPackages   SoftwareUpdateEventDataOperationTypeEnum = "INSTALL_PACKAGES"
	SoftwareUpdateEventDataOperationTypeRemovePackages    SoftwareUpdateEventDataOperationTypeEnum = "REMOVE_PACKAGES"
	SoftwareUpdateEventDataOperationTypeUpdatePackages    SoftwareUpdateEventDataOperationTypeEnum = "UPDATE_PACKAGES"
	SoftwareUpdateEventDataOperationTypeUpdateSecurity    SoftwareUpdateEventDataOperationTypeEnum = "UPDATE_SECURITY"
	SoftwareUpdateEventDataOperationTypeUpdateBugfix      SoftwareUpdateEventDataOperationTypeEnum = "UPDATE_BUGFIX"
	SoftwareUpdateEventDataOperationTypeUpdateEnhancement SoftwareUpdateEventDataOperationTypeEnum = "UPDATE_ENHANCEMENT"
	SoftwareUpdateEventDataOperationTypeUpdateOther       SoftwareUpdateEventDataOperationTypeEnum = "UPDATE_OTHER"
)

var mappingSoftwareUpdateEventDataOperationTypeEnum = map[string]SoftwareUpdateEventDataOperationTypeEnum{
	"UPDATE_ALL_PACKAGES": SoftwareUpdateEventDataOperationTypeUpdateAllPackages,
	"INSTALL_PACKAGES":    SoftwareUpdateEventDataOperationTypeInstallPackages,
	"REMOVE_PACKAGES":     SoftwareUpdateEventDataOperationTypeRemovePackages,
	"UPDATE_PACKAGES":     SoftwareUpdateEventDataOperationTypeUpdatePackages,
	"UPDATE_SECURITY":     SoftwareUpdateEventDataOperationTypeUpdateSecurity,
	"UPDATE_BUGFIX":       SoftwareUpdateEventDataOperationTypeUpdateBugfix,
	"UPDATE_ENHANCEMENT":  SoftwareUpdateEventDataOperationTypeUpdateEnhancement,
	"UPDATE_OTHER":        SoftwareUpdateEventDataOperationTypeUpdateOther,
}

var mappingSoftwareUpdateEventDataOperationTypeEnumLowerCase = map[string]SoftwareUpdateEventDataOperationTypeEnum{
	"update_all_packages": SoftwareUpdateEventDataOperationTypeUpdateAllPackages,
	"install_packages":    SoftwareUpdateEventDataOperationTypeInstallPackages,
	"remove_packages":     SoftwareUpdateEventDataOperationTypeRemovePackages,
	"update_packages":     SoftwareUpdateEventDataOperationTypeUpdatePackages,
	"update_security":     SoftwareUpdateEventDataOperationTypeUpdateSecurity,
	"update_bugfix":       SoftwareUpdateEventDataOperationTypeUpdateBugfix,
	"update_enhancement":  SoftwareUpdateEventDataOperationTypeUpdateEnhancement,
	"update_other":        SoftwareUpdateEventDataOperationTypeUpdateOther,
}

// GetSoftwareUpdateEventDataOperationTypeEnumValues Enumerates the set of values for SoftwareUpdateEventDataOperationTypeEnum
func GetSoftwareUpdateEventDataOperationTypeEnumValues() []SoftwareUpdateEventDataOperationTypeEnum {
	values := make([]SoftwareUpdateEventDataOperationTypeEnum, 0)
	for _, v := range mappingSoftwareUpdateEventDataOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwareUpdateEventDataOperationTypeEnumStringValues Enumerates the set of values in String for SoftwareUpdateEventDataOperationTypeEnum
func GetSoftwareUpdateEventDataOperationTypeEnumStringValues() []string {
	return []string{
		"UPDATE_ALL_PACKAGES",
		"INSTALL_PACKAGES",
		"REMOVE_PACKAGES",
		"UPDATE_PACKAGES",
		"UPDATE_SECURITY",
		"UPDATE_BUGFIX",
		"UPDATE_ENHANCEMENT",
		"UPDATE_OTHER",
	}
}

// GetMappingSoftwareUpdateEventDataOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwareUpdateEventDataOperationTypeEnum(val string) (SoftwareUpdateEventDataOperationTypeEnum, bool) {
	enum, ok := mappingSoftwareUpdateEventDataOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
