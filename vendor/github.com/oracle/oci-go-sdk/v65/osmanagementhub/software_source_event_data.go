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

// SoftwareSourceEventData Provides additional information for a software source event.
type SoftwareSourceEventData struct {

	// Type of software source operation.
	OperationType SoftwareSourceEventDataOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of the software source operation.
	Status EventStatusEnum `mandatory:"true" json:"status"`

	AdditionalDetails *WorkRequestEventDataAdditionalDetails `mandatory:"false" json:"additionalDetails"`
}

func (m SoftwareSourceEventData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoftwareSourceEventData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSoftwareSourceEventDataOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetSoftwareSourceEventDataOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEventStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetEventStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SoftwareSourceEventDataOperationTypeEnum Enum with underlying type: string
type SoftwareSourceEventDataOperationTypeEnum string

// Set of constants representing the allowable values for SoftwareSourceEventDataOperationTypeEnum
const (
	SoftwareSourceEventDataOperationTypeEnableModuleStreams  SoftwareSourceEventDataOperationTypeEnum = "ENABLE_MODULE_STREAMS"
	SoftwareSourceEventDataOperationTypeDisableModuleStreams SoftwareSourceEventDataOperationTypeEnum = "DISABLE_MODULE_STREAMS"
	SoftwareSourceEventDataOperationTypeSwitchModuleStream   SoftwareSourceEventDataOperationTypeEnum = "SWITCH_MODULE_STREAM"
	SoftwareSourceEventDataOperationTypeInstallModuleProfile SoftwareSourceEventDataOperationTypeEnum = "INSTALL_MODULE_PROFILE"
	SoftwareSourceEventDataOperationTypeRemoveModuleProfiles SoftwareSourceEventDataOperationTypeEnum = "REMOVE_MODULE_PROFILES"
	SoftwareSourceEventDataOperationTypeSetSoftwareSources   SoftwareSourceEventDataOperationTypeEnum = "SET_SOFTWARE_SOURCES"
	SoftwareSourceEventDataOperationTypeUpdateSoftwareSource SoftwareSourceEventDataOperationTypeEnum = "UPDATE_SOFTWARE_SOURCE"
	SoftwareSourceEventDataOperationTypeCreateSoftwareSource SoftwareSourceEventDataOperationTypeEnum = "CREATE_SOFTWARE_SOURCE"
	SoftwareSourceEventDataOperationTypeLifecyclePromotion   SoftwareSourceEventDataOperationTypeEnum = "LIFECYCLE_PROMOTION"
)

var mappingSoftwareSourceEventDataOperationTypeEnum = map[string]SoftwareSourceEventDataOperationTypeEnum{
	"ENABLE_MODULE_STREAMS":  SoftwareSourceEventDataOperationTypeEnableModuleStreams,
	"DISABLE_MODULE_STREAMS": SoftwareSourceEventDataOperationTypeDisableModuleStreams,
	"SWITCH_MODULE_STREAM":   SoftwareSourceEventDataOperationTypeSwitchModuleStream,
	"INSTALL_MODULE_PROFILE": SoftwareSourceEventDataOperationTypeInstallModuleProfile,
	"REMOVE_MODULE_PROFILES": SoftwareSourceEventDataOperationTypeRemoveModuleProfiles,
	"SET_SOFTWARE_SOURCES":   SoftwareSourceEventDataOperationTypeSetSoftwareSources,
	"UPDATE_SOFTWARE_SOURCE": SoftwareSourceEventDataOperationTypeUpdateSoftwareSource,
	"CREATE_SOFTWARE_SOURCE": SoftwareSourceEventDataOperationTypeCreateSoftwareSource,
	"LIFECYCLE_PROMOTION":    SoftwareSourceEventDataOperationTypeLifecyclePromotion,
}

var mappingSoftwareSourceEventDataOperationTypeEnumLowerCase = map[string]SoftwareSourceEventDataOperationTypeEnum{
	"enable_module_streams":  SoftwareSourceEventDataOperationTypeEnableModuleStreams,
	"disable_module_streams": SoftwareSourceEventDataOperationTypeDisableModuleStreams,
	"switch_module_stream":   SoftwareSourceEventDataOperationTypeSwitchModuleStream,
	"install_module_profile": SoftwareSourceEventDataOperationTypeInstallModuleProfile,
	"remove_module_profiles": SoftwareSourceEventDataOperationTypeRemoveModuleProfiles,
	"set_software_sources":   SoftwareSourceEventDataOperationTypeSetSoftwareSources,
	"update_software_source": SoftwareSourceEventDataOperationTypeUpdateSoftwareSource,
	"create_software_source": SoftwareSourceEventDataOperationTypeCreateSoftwareSource,
	"lifecycle_promotion":    SoftwareSourceEventDataOperationTypeLifecyclePromotion,
}

// GetSoftwareSourceEventDataOperationTypeEnumValues Enumerates the set of values for SoftwareSourceEventDataOperationTypeEnum
func GetSoftwareSourceEventDataOperationTypeEnumValues() []SoftwareSourceEventDataOperationTypeEnum {
	values := make([]SoftwareSourceEventDataOperationTypeEnum, 0)
	for _, v := range mappingSoftwareSourceEventDataOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwareSourceEventDataOperationTypeEnumStringValues Enumerates the set of values in String for SoftwareSourceEventDataOperationTypeEnum
func GetSoftwareSourceEventDataOperationTypeEnumStringValues() []string {
	return []string{
		"ENABLE_MODULE_STREAMS",
		"DISABLE_MODULE_STREAMS",
		"SWITCH_MODULE_STREAM",
		"INSTALL_MODULE_PROFILE",
		"REMOVE_MODULE_PROFILES",
		"SET_SOFTWARE_SOURCES",
		"UPDATE_SOFTWARE_SOURCE",
		"CREATE_SOFTWARE_SOURCE",
		"LIFECYCLE_PROMOTION",
	}
}

// GetMappingSoftwareSourceEventDataOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwareSourceEventDataOperationTypeEnum(val string) (SoftwareSourceEventDataOperationTypeEnum, bool) {
	enum, ok := mappingSoftwareSourceEventDataOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
