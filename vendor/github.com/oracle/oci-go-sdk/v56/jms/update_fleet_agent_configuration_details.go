// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateFleetAgentConfigurationDetails Attributes to update a Fleet Agent Configuration.
type UpdateFleetAgentConfigurationDetails struct {

	// The frequency (in minutes) of JRE scanning. (That is, how often should JMS scan for JRE installations.)
	JreScanFrequencyInMinutes *int `mandatory:"false" json:"jreScanFrequencyInMinutes"`

	// The frequency (in minutes) of Java Usage Tracker processing. (That is, how often should JMS process data from the Java Usage Tracker.)
	JavaUsageTrackerProcessingFrequencyInMinutes *int `mandatory:"false" json:"javaUsageTrackerProcessingFrequencyInMinutes"`

	LinuxConfiguration *FleetAgentOsConfiguration `mandatory:"false" json:"linuxConfiguration"`

	WindowsConfiguration *FleetAgentOsConfiguration `mandatory:"false" json:"windowsConfiguration"`
}

func (m UpdateFleetAgentConfigurationDetails) String() string {
	return common.PointerString(m)
}
