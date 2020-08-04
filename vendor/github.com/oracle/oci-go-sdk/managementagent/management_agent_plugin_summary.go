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

// ManagementAgentPluginSummary Summary of the ManagementAgentPlugin.
type ManagementAgentPluginSummary struct {

	// Management Agent Plugin Id
	Id *string `mandatory:"true" json:"id"`

	// Management Agent Plugin Name
	Name *string `mandatory:"true" json:"name"`

	// The current state of Management Agent Plugin
	LifecycleState LifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Management Agent Plugin Version
	Version *int `mandatory:"false" json:"version"`

	// Supported Platform Types
	SupportedPlatformTypes []PlatformTypesEnum `mandatory:"false" json:"supportedPlatformTypes,omitempty"`

	// Management Agent Plugin Display Name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Management Agent Plugin description
	Description *string `mandatory:"false" json:"description"`

	// A flag to indicate whether a given plugin can be deployed from Agent Console UI or not.
	IsConsoleDeployable *bool `mandatory:"false" json:"isConsoleDeployable"`
}

func (m ManagementAgentPluginSummary) String() string {
	return common.PointerString(m)
}
