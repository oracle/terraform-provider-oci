// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ManagementAgentInstallKey The details of the Agent install Key
type ManagementAgentInstallKey struct {

	// Agent install Key identifier
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Management Agent Install Key Name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Management Agent Install Key
	Key *string `mandatory:"false" json:"key"`

	// Principal id of user who created the Agent Install key
	CreatedByPrincipalId *string `mandatory:"false" json:"createdByPrincipalId"`

	// Total number of install for this keys
	AllowedKeyInstallCount *int `mandatory:"false" json:"allowedKeyInstallCount"`

	// Total number of install for this keys
	CurrentKeyInstallCount *int `mandatory:"false" json:"currentKeyInstallCount"`

	// Status of Key
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// date after which key would expire after creation
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// The time when Management Agent install Key was created. An RFC3339 formatted date time string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time when Management Agent install Key was updated. An RFC3339 formatted date time string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m ManagementAgentInstallKey) String() string {
	return common.PointerString(m)
}
