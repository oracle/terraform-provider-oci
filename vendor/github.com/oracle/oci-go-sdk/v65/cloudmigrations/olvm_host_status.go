// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmHostStatus Type representing a host status.
type OlvmHostStatus struct {

	// Type representing a host status.
	Status OlvmHostStatusStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m OlvmHostStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmHostStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmHostStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOlvmHostStatusStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmHostStatusStatusEnum Enum with underlying type: string
type OlvmHostStatusStatusEnum string

// Set of constants representing the allowable values for OlvmHostStatusStatusEnum
const (
	OlvmHostStatusStatusConnecting              OlvmHostStatusStatusEnum = "CONNECTING"
	OlvmHostStatusStatusDown                    OlvmHostStatusStatusEnum = "DOWN"
	OlvmHostStatusStatusError                   OlvmHostStatusStatusEnum = "ERROR"
	OlvmHostStatusStatusInitializing            OlvmHostStatusStatusEnum = "INITIALIZING"
	OlvmHostStatusStatusInstallFailed           OlvmHostStatusStatusEnum = "INSTALL_FAILED"
	OlvmHostStatusStatusInstalling              OlvmHostStatusStatusEnum = "INSTALLING"
	OlvmHostStatusStatusInstallingOs            OlvmHostStatusStatusEnum = "INSTALLING_OS"
	OlvmHostStatusStatusKdumping                OlvmHostStatusStatusEnum = "KDUMPING"
	OlvmHostStatusStatusMaintenance             OlvmHostStatusStatusEnum = "MAINTENANCE"
	OlvmHostStatusStatusNonOperational          OlvmHostStatusStatusEnum = "NON_OPERATIONAL"
	OlvmHostStatusStatusNonResponsive           OlvmHostStatusStatusEnum = "NON_RESPONSIVE"
	OlvmHostStatusStatusPendingApproval         OlvmHostStatusStatusEnum = "PENDING_APPROVAL"
	OlvmHostStatusStatusPreparingForMaintenance OlvmHostStatusStatusEnum = "PREPARING_FOR_MAINTENANCE"
	OlvmHostStatusStatusReboot                  OlvmHostStatusStatusEnum = "REBOOT"
	OlvmHostStatusStatusUnassigned              OlvmHostStatusStatusEnum = "UNASSIGNED"
	OlvmHostStatusStatusUp                      OlvmHostStatusStatusEnum = "UP"
)

var mappingOlvmHostStatusStatusEnum = map[string]OlvmHostStatusStatusEnum{
	"CONNECTING":                OlvmHostStatusStatusConnecting,
	"DOWN":                      OlvmHostStatusStatusDown,
	"ERROR":                     OlvmHostStatusStatusError,
	"INITIALIZING":              OlvmHostStatusStatusInitializing,
	"INSTALL_FAILED":            OlvmHostStatusStatusInstallFailed,
	"INSTALLING":                OlvmHostStatusStatusInstalling,
	"INSTALLING_OS":             OlvmHostStatusStatusInstallingOs,
	"KDUMPING":                  OlvmHostStatusStatusKdumping,
	"MAINTENANCE":               OlvmHostStatusStatusMaintenance,
	"NON_OPERATIONAL":           OlvmHostStatusStatusNonOperational,
	"NON_RESPONSIVE":            OlvmHostStatusStatusNonResponsive,
	"PENDING_APPROVAL":          OlvmHostStatusStatusPendingApproval,
	"PREPARING_FOR_MAINTENANCE": OlvmHostStatusStatusPreparingForMaintenance,
	"REBOOT":                    OlvmHostStatusStatusReboot,
	"UNASSIGNED":                OlvmHostStatusStatusUnassigned,
	"UP":                        OlvmHostStatusStatusUp,
}

var mappingOlvmHostStatusStatusEnumLowerCase = map[string]OlvmHostStatusStatusEnum{
	"connecting":                OlvmHostStatusStatusConnecting,
	"down":                      OlvmHostStatusStatusDown,
	"error":                     OlvmHostStatusStatusError,
	"initializing":              OlvmHostStatusStatusInitializing,
	"install_failed":            OlvmHostStatusStatusInstallFailed,
	"installing":                OlvmHostStatusStatusInstalling,
	"installing_os":             OlvmHostStatusStatusInstallingOs,
	"kdumping":                  OlvmHostStatusStatusKdumping,
	"maintenance":               OlvmHostStatusStatusMaintenance,
	"non_operational":           OlvmHostStatusStatusNonOperational,
	"non_responsive":            OlvmHostStatusStatusNonResponsive,
	"pending_approval":          OlvmHostStatusStatusPendingApproval,
	"preparing_for_maintenance": OlvmHostStatusStatusPreparingForMaintenance,
	"reboot":                    OlvmHostStatusStatusReboot,
	"unassigned":                OlvmHostStatusStatusUnassigned,
	"up":                        OlvmHostStatusStatusUp,
}

// GetOlvmHostStatusStatusEnumValues Enumerates the set of values for OlvmHostStatusStatusEnum
func GetOlvmHostStatusStatusEnumValues() []OlvmHostStatusStatusEnum {
	values := make([]OlvmHostStatusStatusEnum, 0)
	for _, v := range mappingOlvmHostStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmHostStatusStatusEnumStringValues Enumerates the set of values in String for OlvmHostStatusStatusEnum
func GetOlvmHostStatusStatusEnumStringValues() []string {
	return []string{
		"CONNECTING",
		"DOWN",
		"ERROR",
		"INITIALIZING",
		"INSTALL_FAILED",
		"INSTALLING",
		"INSTALLING_OS",
		"KDUMPING",
		"MAINTENANCE",
		"NON_OPERATIONAL",
		"NON_RESPONSIVE",
		"PENDING_APPROVAL",
		"PREPARING_FOR_MAINTENANCE",
		"REBOOT",
		"UNASSIGNED",
		"UP",
	}
}

// GetMappingOlvmHostStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmHostStatusStatusEnum(val string) (OlvmHostStatusStatusEnum, bool) {
	enum, ok := mappingOlvmHostStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
