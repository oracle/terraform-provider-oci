// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Configuration API
//
// Use the Application Performance Monitoring Configuration API to query and set Application Performance Monitoring
// configuration. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmconfig

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AgentConfigOverride Agent configuration overrides that should apply to a subset of the agents associated with an Agent Config object.
type AgentConfigOverride struct {

	// The string that defines the Agent Filter expression.
	AgentFilter *string `mandatory:"false" json:"agentFilter"`

	// A map whose key is a substitution variable specified within the configuration's body. For example, if below was specified in the configuration's body
	// {{ isJfrEnabled | default false }}
	// Then a valid map key would be "isJfrEnabled". The value is typically different than the default specified in the configuration's body.
	// Thus, in this example, the map entry could be "isJfrEnabled": true
	OverrideMap map[string]string `mandatory:"false" json:"overrideMap"`
}

func (m AgentConfigOverride) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AgentConfigOverride) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
