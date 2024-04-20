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

// InstallWindowsUpdatesOnManagedInstanceDetails Provides the information used to install Windows updates on a managed instance.
type InstallWindowsUpdatesOnManagedInstanceDetails struct {

	// The list of Windows update unique identifiers.
	// Note that this is not an OCID, but is a unique identifier assigned by Microsoft.
	// Example: '6981d463-cd91-4a26-b7c4-ea4ded9183ed'
	WindowsUpdateName []string `mandatory:"false" json:"windowsUpdateName"`

	// The types of Windows updates to be installed.
	WindowsUpdateTypes []WindowsUpdateTypesEnum `mandatory:"false" json:"windowsUpdateTypes,omitempty"`

	WorkRequestDetails *WorkRequestDetails `mandatory:"false" json:"workRequestDetails"`
}

func (m InstallWindowsUpdatesOnManagedInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstallWindowsUpdatesOnManagedInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.WindowsUpdateTypes {
		if _, ok := GetMappingWindowsUpdateTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WindowsUpdateTypes: %s. Supported values are: %s.", val, strings.Join(GetWindowsUpdateTypesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
