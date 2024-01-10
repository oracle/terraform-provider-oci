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

// ManagedInstanceSummary Summary of the ManagedInstance.
type ManagedInstanceSummary struct {

	// The OCID for the managed instance.
	Id *string `mandatory:"true" json:"id"`

	// Managed instance identifier.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID for the tenancy this managed instance resides in.
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// The OCID for the compartment this managed instance resides in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// status of the managed instance.
	Status ManagedInstanceStatusEnum `mandatory:"true" json:"status"`

	// Information specified by the user about the managed instance.
	Description *string `mandatory:"false" json:"description"`

	// Location of the managed instance.
	Location ManagedInstanceLocationEnum `mandatory:"false" json:"location,omitempty"`

	// The CPU architecture type of the managed instance.
	Architecture ArchTypeEnum `mandatory:"false" json:"architecture,omitempty"`

	// The Operating System type of the managed instance.
	OsFamily OsFamilyEnum `mandatory:"false" json:"osFamily,omitempty"`

	ManagedInstanceGroup *Id `mandatory:"false" json:"managedInstanceGroup"`

	LifecycleEnvironment *Id `mandatory:"false" json:"lifecycleEnvironment"`

	LifecycleStage *Id `mandatory:"false" json:"lifecycleStage"`

	// Indicates whether a reboot is required to complete installation of updates.
	IsRebootRequired *bool `mandatory:"false" json:"isRebootRequired"`

	// Number of updates available to be installed.
	UpdatesAvailable *int `mandatory:"false" json:"updatesAvailable"`

	// Whether this managed instance is acting as an on-premise management station.
	IsManagementStation *bool `mandatory:"false" json:"isManagementStation"`
}

func (m ManagedInstanceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedInstanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingManagedInstanceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetManagedInstanceStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingManagedInstanceLocationEnum(string(m.Location)); !ok && m.Location != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Location: %s. Supported values are: %s.", m.Location, strings.Join(GetManagedInstanceLocationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArchTypeEnum(string(m.Architecture)); !ok && m.Architecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Architecture: %s. Supported values are: %s.", m.Architecture, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOsFamilyEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamilyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
