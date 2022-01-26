// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Agent API
//
// API for the Oracle Cloud Agent software running on compute instances. Oracle Cloud Agent
// is a lightweight process that monitors and manages compute instances.
//

package computeinstanceagent

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Plugin The agent plugin
type Plugin struct {

	// The plugin name
	Name *string `mandatory:"true" json:"name"`

	// The plugin version
	Version *string `mandatory:"true" json:"version"`

	// The plugin status
	Status *string `mandatory:"true" json:"status"`

	// The last update time of the plugin
	LastUpdateTime *common.SDKTime `mandatory:"false" json:"lastUpdateTime"`

	// The optional message from the agent plugin
	Message *string `mandatory:"false" json:"message"`
}

func (m Plugin) String() string {
	return common.PointerString(m)
}
