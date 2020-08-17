// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LogSavedSearchSummary A summary of a log saved search that can be used to save and share a given search result.
type LogSavedSearchSummary struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of a user-friendly name. It has to be unique within enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	// True if the LogSavedSearch should be show as quickstart in the UI
	IsQuickStart *bool `mandatory:"true" json:"isQuickStart"`

	// Time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the resource was last modified.
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`

	// The state of the LogSavedSearch
	LifecycleState LogSavedSearchLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m LogSavedSearchSummary) String() string {
	return common.PointerString(m)
}
