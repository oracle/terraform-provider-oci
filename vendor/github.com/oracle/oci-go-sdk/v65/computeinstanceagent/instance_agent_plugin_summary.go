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

// InstanceAgentPluginSummary The agent plugin information
type InstanceAgentPluginSummary struct {

	// The plugin name
	Name *string `mandatory:"true" json:"name"`

	// The plugin status Specified the plugin state on the instance * `RUNNING` - The plugin is in running state * `STOPPED` - The plugin is in stopped state * `NOT_SUPPORTED` - The plugin is not supported on this platform * `INVALID` - The plugin state is not recognizable by the service
	Status InstanceAgentPluginSummaryStatusEnum `mandatory:"true" json:"status"`

	// The last update time of the plugin in UTC
	TimeLastUpdatedUtc *common.SDKTime `mandatory:"true" json:"timeLastUpdatedUtc"`
}

func (m InstanceAgentPluginSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceAgentPluginSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInstanceAgentPluginSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetInstanceAgentPluginSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstanceAgentPluginSummaryStatusEnum Enum with underlying type: string
type InstanceAgentPluginSummaryStatusEnum string

// Set of constants representing the allowable values for InstanceAgentPluginSummaryStatusEnum
const (
	InstanceAgentPluginSummaryStatusRunning      InstanceAgentPluginSummaryStatusEnum = "RUNNING"
	InstanceAgentPluginSummaryStatusStopped      InstanceAgentPluginSummaryStatusEnum = "STOPPED"
	InstanceAgentPluginSummaryStatusNotSupported InstanceAgentPluginSummaryStatusEnum = "NOT_SUPPORTED"
	InstanceAgentPluginSummaryStatusInvalid      InstanceAgentPluginSummaryStatusEnum = "INVALID"
)

var mappingInstanceAgentPluginSummaryStatusEnum = map[string]InstanceAgentPluginSummaryStatusEnum{
	"RUNNING":       InstanceAgentPluginSummaryStatusRunning,
	"STOPPED":       InstanceAgentPluginSummaryStatusStopped,
	"NOT_SUPPORTED": InstanceAgentPluginSummaryStatusNotSupported,
	"INVALID":       InstanceAgentPluginSummaryStatusInvalid,
}

var mappingInstanceAgentPluginSummaryStatusEnumLowerCase = map[string]InstanceAgentPluginSummaryStatusEnum{
	"running":       InstanceAgentPluginSummaryStatusRunning,
	"stopped":       InstanceAgentPluginSummaryStatusStopped,
	"not_supported": InstanceAgentPluginSummaryStatusNotSupported,
	"invalid":       InstanceAgentPluginSummaryStatusInvalid,
}

// GetInstanceAgentPluginSummaryStatusEnumValues Enumerates the set of values for InstanceAgentPluginSummaryStatusEnum
func GetInstanceAgentPluginSummaryStatusEnumValues() []InstanceAgentPluginSummaryStatusEnum {
	values := make([]InstanceAgentPluginSummaryStatusEnum, 0)
	for _, v := range mappingInstanceAgentPluginSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceAgentPluginSummaryStatusEnumStringValues Enumerates the set of values in String for InstanceAgentPluginSummaryStatusEnum
func GetInstanceAgentPluginSummaryStatusEnumStringValues() []string {
	return []string{
		"RUNNING",
		"STOPPED",
		"NOT_SUPPORTED",
		"INVALID",
	}
}

// GetMappingInstanceAgentPluginSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceAgentPluginSummaryStatusEnum(val string) (InstanceAgentPluginSummaryStatusEnum, bool) {
	enum, ok := mappingInstanceAgentPluginSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
