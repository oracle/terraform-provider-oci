// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LaunchInstanceAgentConfigDetails Configuration options for the Oracle Cloud Agent software running on the instance.
type LaunchInstanceAgentConfigDetails struct {

	// Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the
	// monitoring plugins. By default, the value is false (monitoring plugins are enabled).
	// These are the monitoring plugins: Compute instance monitoring
	// and Custom logs monitoring.
	// The monitoring plugins are controlled by this parameter and by the per-plugin
	// configuration in the `pluginsConfig` object.
	// - If `isMonitoringDisabled` is true, all the monitoring plugins are disabled, regardless of
	// the per-plugin configuration.
	// - If `isMonitoringDisabled` is false, all the monitoring plugins are enabled. You
	// can optionally disable individual monitoring plugins by providing a value in the `pluginsConfig`
	// object.
	IsMonitoringDisabled *bool `mandatory:"false" json:"isMonitoringDisabled"`

	// Whether Oracle Cloud Agent can run all the available management plugins.
	// By default, the value is false (management plugins are enabled).
	// These are the management plugins: OS Management Service Agent and Compute instance
	// run command.
	// The management plugins are controlled by this parameter and the per-plugin
	// configuration in the `pluginsConfig` object.
	// - If `isManagementDisabled` is true, all the management plugins are disabled, regardless of
	// the per-plugin configuration.
	// - If `isManagementDisabled` is false, all the management plugins are enabled. You
	// can optionally disable individual management plugins by providing a value in the `pluginsConfig`
	// object.
	IsManagementDisabled *bool `mandatory:"false" json:"isManagementDisabled"`

	// Whether Oracle Cloud Agent can run all the available plugins.
	// This includes the management and monitoring plugins.
	// To get a list of available plugins, use the
	// ListInstanceagentAvailablePlugins
	// operation in the Oracle Cloud Agent API. For more information about the available plugins, see
	// Managing Plugins with Oracle Cloud Agent (https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm).
	AreAllPluginsDisabled *bool `mandatory:"false" json:"areAllPluginsDisabled"`

	// The configuration of plugins associated with this instance.
	PluginsConfig []InstanceAgentPluginConfigDetails `mandatory:"false" json:"pluginsConfig"`
}

func (m LaunchInstanceAgentConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LaunchInstanceAgentConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
