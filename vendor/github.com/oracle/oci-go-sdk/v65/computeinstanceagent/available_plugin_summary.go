// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// AvailablePluginSummary Information about where a plugin is supported.
type AvailablePluginSummary struct {

	// The plugin name.
	Name *string `mandatory:"true" json:"name"`

	// Whether the plugin is supported.
	IsSupported *bool `mandatory:"true" json:"isSupported"`

	// Whether the plugin is enabled or disabled by default.
	IsEnabledByDefault *bool `mandatory:"true" json:"isEnabledByDefault"`

	// A brief description of the plugin's functionality.
	Summary *string `mandatory:"false" json:"summary"`
}

func (m AvailablePluginSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AvailablePluginSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
