// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ManagedInstanceSummary An OCI Compute instance that is being managed
type ManagedInstanceSummary struct {

	// user settable name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID for the managed instance
	Id *string `mandatory:"true" json:"id"`

	// OCID for the Compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Time at which the instance last checked in
	LastCheckin *string `mandatory:"false" json:"lastCheckin"`

	// Time at which the instance last booted
	LastBoot *string `mandatory:"false" json:"lastBoot"`

	// Number of updates available to be installed
	UpdatesAvailable *int `mandatory:"false" json:"updatesAvailable"`

	// Information specified by the user about the managed instance
	Description *string `mandatory:"false" json:"description"`

	// status of the managed instance.
	Status ManagedInstanceSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The Operating System type of the managed instance.
	OsFamily OsFamiliesEnum `mandatory:"false" json:"osFamily,omitempty"`

	// Indicates whether a reboot is required to complete installation of updates.
	IsRebootRequired *bool `mandatory:"false" json:"isRebootRequired"`
}

func (m ManagedInstanceSummary) String() string {
	return common.PointerString(m)
}

// ManagedInstanceSummaryStatusEnum Enum with underlying type: string
type ManagedInstanceSummaryStatusEnum string

// Set of constants representing the allowable values for ManagedInstanceSummaryStatusEnum
const (
	ManagedInstanceSummaryStatusNormal      ManagedInstanceSummaryStatusEnum = "NORMAL"
	ManagedInstanceSummaryStatusUnreachable ManagedInstanceSummaryStatusEnum = "UNREACHABLE"
	ManagedInstanceSummaryStatusError       ManagedInstanceSummaryStatusEnum = "ERROR"
	ManagedInstanceSummaryStatusWarning     ManagedInstanceSummaryStatusEnum = "WARNING"
)

var mappingManagedInstanceSummaryStatus = map[string]ManagedInstanceSummaryStatusEnum{
	"NORMAL":      ManagedInstanceSummaryStatusNormal,
	"UNREACHABLE": ManagedInstanceSummaryStatusUnreachable,
	"ERROR":       ManagedInstanceSummaryStatusError,
	"WARNING":     ManagedInstanceSummaryStatusWarning,
}

// GetManagedInstanceSummaryStatusEnumValues Enumerates the set of values for ManagedInstanceSummaryStatusEnum
func GetManagedInstanceSummaryStatusEnumValues() []ManagedInstanceSummaryStatusEnum {
	values := make([]ManagedInstanceSummaryStatusEnum, 0)
	for _, v := range mappingManagedInstanceSummaryStatus {
		values = append(values, v)
	}
	return values
}
