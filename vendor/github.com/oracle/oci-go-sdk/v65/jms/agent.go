// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Agent Information about the agent.
type Agent struct {

	// The name of the agent.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The agent type.
	Type AgentTypeEnum `mandatory:"true" json:"type"`

	// The java version.
	JavaVersion *string `mandatory:"true" json:"javaVersion"`

	// The security status of the Java Runtime.
	JavaSecurityStatus JreSecurityStatusEnum `mandatory:"true" json:"javaSecurityStatus"`

	// A list of plugins installed on this agent.
	Plugins []Plugin `mandatory:"true" json:"plugins"`
}

func (m Agent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Agent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAgentTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAgentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJreSecurityStatusEnum(string(m.JavaSecurityStatus)); !ok && m.JavaSecurityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JavaSecurityStatus: %s. Supported values are: %s.", m.JavaSecurityStatus, strings.Join(GetJreSecurityStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
