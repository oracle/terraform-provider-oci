// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagedInstanceModuleSummary Summary information pertaining to a module on a managed instance.
type ManagedInstanceModuleSummary struct {

	// The module name.
	Name *string `mandatory:"true" json:"name"`

	// The stream that is enabled in the module.
	EnabledStream *string `mandatory:"false" json:"enabledStream"`

	// List of installed profiles in the enabled stream of the module.
	InstalledProfiles []string `mandatory:"false" json:"installedProfiles"`

	// List of streams that are active in the module.
	ActiveStreams []string `mandatory:"false" json:"activeStreams"`

	// List of streams that are disabled in the module.
	DisabledStreams []string `mandatory:"false" json:"disabledStreams"`

	// The OCID of the software source that provides this module and the associated streams.
	SoftwareSourceId *string `mandatory:"false" json:"softwareSourceId"`
}

func (m ManagedInstanceModuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedInstanceModuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
