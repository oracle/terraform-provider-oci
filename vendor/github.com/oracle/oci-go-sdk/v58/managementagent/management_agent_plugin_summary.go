// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementAgentPluginSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}

	for _, val := range m.SupportedPlatformTypes {
		if _, ok := GetMappingPlatformTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SupportedPlatformTypes: %s. Supported values are: %s.", val, strings.Join(GetPlatformTypesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
