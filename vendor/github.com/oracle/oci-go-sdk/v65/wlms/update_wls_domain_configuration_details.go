// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateWlsDomainConfigurationDetails The WebLogic domain configuration for update operation.
type UpdateWlsDomainConfigurationDetails struct {

	// Whether or not the WebLogic domain is enabled for patching.
	IsPatchEnabled *bool `mandatory:"false" json:"isPatchEnabled"`

	// Whether or not to rollback on failure during patching of WebLogic domain.
	IsRollbackOnFailure *bool `mandatory:"false" json:"isRollbackOnFailure"`

	// Servers shutdown timeout in seconds. If set to 0 seconds, it means there is no timeout.
	ServersShutdownTimeout *int `mandatory:"false" json:"serversShutdownTimeout"`

	// Whether to manage the admin server using Node Manager or scripts.
	AdminServerControlMode ServerControlModeEnum `mandatory:"false" json:"adminServerControlMode,omitempty"`

	// Whether to manage the managed server using Node Manager or scripts.
	ManagedServerControlMode ServerControlModeEnum `mandatory:"false" json:"managedServerControlMode,omitempty"`

	// Path to admin server start script.
	AdminServerStartScriptPath *string `mandatory:"false" json:"adminServerStartScriptPath"`

	// Path to admin server stop script.
	AdminServerStopScriptPath *string `mandatory:"false" json:"adminServerStopScriptPath"`

	// Path to managed server start script.
	ManagedServerStartScriptPath *string `mandatory:"false" json:"managedServerStartScriptPath"`

	// Path to managed server stop script.
	ManagedServerStopScriptPath *string `mandatory:"false" json:"managedServerStopScriptPath"`
}

func (m UpdateWlsDomainConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateWlsDomainConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingServerControlModeEnum(string(m.AdminServerControlMode)); !ok && m.AdminServerControlMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdminServerControlMode: %s. Supported values are: %s.", m.AdminServerControlMode, strings.Join(GetServerControlModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingServerControlModeEnum(string(m.ManagedServerControlMode)); !ok && m.ManagedServerControlMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedServerControlMode: %s. Supported values are: %s.", m.ManagedServerControlMode, strings.Join(GetServerControlModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
