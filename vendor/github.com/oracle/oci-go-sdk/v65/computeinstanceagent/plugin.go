// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Agent API
//
// API for the Oracle Cloud Agent software running on compute instances. Oracle Cloud Agent
// is a lightweight process that monitors and manages compute instances.
//

package computeinstanceagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Plugin) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
