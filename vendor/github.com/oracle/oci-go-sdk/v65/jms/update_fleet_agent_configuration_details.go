// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateFleetAgentConfigurationDetails Attributes to update a Fleet Agent Configuration.
type UpdateFleetAgentConfigurationDetails struct {

	// The frequency (in minutes) of JRE scanning. (That is, how often should JMS scan for JRE installations.)
	JreScanFrequencyInMinutes *int `mandatory:"false" json:"jreScanFrequencyInMinutes"`

	// The frequency (in minutes) of Java Usage Tracker processing. (That is, how often should JMS process data from the Java Usage Tracker.)
	JavaUsageTrackerProcessingFrequencyInMinutes *int `mandatory:"false" json:"javaUsageTrackerProcessingFrequencyInMinutes"`

	// The validity period in days for work requests.
	WorkRequestValidityPeriodInDays *int `mandatory:"false" json:"workRequestValidityPeriodInDays"`

	// Agent polling interval in minutes
	AgentPollingIntervalInMinutes *int `mandatory:"false" json:"agentPollingIntervalInMinutes"`

	// Collect JMS agent metrics on all managed instances in the fleet.
	IsCollectingManagedInstanceMetricsEnabled *bool `mandatory:"false" json:"isCollectingManagedInstanceMetricsEnabled"`

	// Collect username for application invocations for all managed instances in the fleet.
	IsCollectingUsernamesEnabled *bool `mandatory:"false" json:"isCollectingUsernamesEnabled"`

	LinuxConfiguration *FleetAgentOsConfiguration `mandatory:"false" json:"linuxConfiguration"`

	WindowsConfiguration *FleetAgentOsConfiguration `mandatory:"false" json:"windowsConfiguration"`

	MacOsConfiguration *FleetAgentOsConfiguration `mandatory:"false" json:"macOsConfiguration"`
}

func (m UpdateFleetAgentConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateFleetAgentConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
