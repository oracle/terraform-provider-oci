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

// DeployPluginsDetails The information required to deploy new Management Agent Plugins.
type DeployPluginsDetails struct {

	// Plugin Id
	PluginIds []string `mandatory:"true" json:"pluginIds"`

	// Management Agent Compartment Identifier
	AgentCompartmentId *string `mandatory:"true" json:"agentCompartmentId"`

	// List of Agent identifiers
	AgentIds []string `mandatory:"true" json:"agentIds"`
}

func (m DeployPluginsDetails) String() string {
	return common.PointerString(m)
}
