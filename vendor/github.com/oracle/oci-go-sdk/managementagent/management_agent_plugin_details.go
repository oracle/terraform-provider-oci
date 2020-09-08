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
}

func (m ManagementAgentPluginDetails) String() string {
	return common.PointerString(m)
}
