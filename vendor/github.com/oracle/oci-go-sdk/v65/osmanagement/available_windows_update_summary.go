// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AvailableWindowsUpdateSummary An update available for installation on the Windows managed instance.
type AvailableWindowsUpdateSummary struct {

	// Windows Update name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier for the Windows update. NOTE - This is not an OCID,
	// but is a unique identifier assigned by Microsoft.
	// Example: `6981d463-cd91-4a26-b7c4-ea4ded9183ed`
	Name *string `mandatory:"true" json:"name"`

	// The purpose of this update.
	UpdateType UpdateTypesEnum `mandatory:"true" json:"updateType"`

	// Indicates whether the update can be installed using OSMS.
	IsEligibleForInstallation IsEligibleForInstallationEnum `mandatory:"false" json:"isEligibleForInstallation,omitempty"`

	// Indicates whether a reboot may be required to complete installation of this update.
	IsRebootRequiredForInstallation *bool `mandatory:"false" json:"isRebootRequiredForInstallation"`
}

func (m AvailableWindowsUpdateSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AvailableWindowsUpdateSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateTypesEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetUpdateTypesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingIsEligibleForInstallationEnum(string(m.IsEligibleForInstallation)); !ok && m.IsEligibleForInstallation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsEligibleForInstallation: %s. Supported values are: %s.", m.IsEligibleForInstallation, strings.Join(GetIsEligibleForInstallationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
