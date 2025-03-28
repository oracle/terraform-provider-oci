// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagementAgentPluginDetails The information about the current management agent plugins that agent is having.
type ManagementAgentPluginDetails struct {

	// Management Agent Plugin Name
	PluginName *string `mandatory:"true" json:"pluginName"`

	// Plugin Id
	PluginId *string `mandatory:"false" json:"pluginId"`

	// Management Agent Plugin Identifier, can be renamed
	PluginDisplayName *string `mandatory:"false" json:"pluginDisplayName"`

	// Plugin Version
	PluginVersion *string `mandatory:"false" json:"pluginVersion"`

	// Plugin Status
	PluginStatus PluginStatusEnum `mandatory:"false" json:"pluginStatus,omitempty"`

	// Status message of the Plugin
	PluginStatusMessage *string `mandatory:"false" json:"pluginStatusMessage"`

	// flag indicating whether the plugin is in enabled mode or disabled mode.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

func (m ManagementAgentPluginDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementAgentPluginDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPluginStatusEnum(string(m.PluginStatus)); !ok && m.PluginStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PluginStatus: %s. Supported values are: %s.", m.PluginStatus, strings.Join(GetPluginStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
