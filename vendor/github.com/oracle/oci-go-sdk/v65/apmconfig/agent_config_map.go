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

// AgentConfigMap Collection of agent configuration files.
// For agents that use a single configuration file, this SHOULD contain a single entry and the key MAY be an empty string.
// To apply a different configuration in a subset of the agents, put this block anywhere in the body of the configuration and edit <some variable> and <some content>
// {{ <some variable> | default <some content> }}
// Example:
// com.oracle.apm.agent.tracer.enable.jfr = {{ isJfrEnabled | default false }}
// Then, in the configuration's overrides, specify a different value for <some variable> along with the desired agent filter.
// Example:
// "agentFilter": "ApplicationType='Tomcat'"
//
//	"overrideMap": {
//	    "isJfrEnabled": true
//	}
type AgentConfigMap struct {

	// Map of agent configuration files, where keys are file names.
	ConfigMap map[string]AgentConfigFile `mandatory:"false" json:"configMap"`
}

func (m AgentConfigMap) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AgentConfigMap) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
