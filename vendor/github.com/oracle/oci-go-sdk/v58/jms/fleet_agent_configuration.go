// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// FleetAgentConfiguration Management Agent Configuration for a Fleet. Includes JRE scanning frequency and list of include/exclude file system paths.
type FleetAgentConfiguration struct {

	// The frequency (in minutes) of JRE scanning. (That is, how often should JMS scan for JRE installations.)
	JreScanFrequencyInMinutes *int `mandatory:"true" json:"jreScanFrequencyInMinutes"`

	// The frequency (in minutes) of Java Usage Tracker processing. (That is, how often should JMS process data from the Java Usage Tracker.)
	JavaUsageTrackerProcessingFrequencyInMinutes *int `mandatory:"true" json:"javaUsageTrackerProcessingFrequencyInMinutes"`

	LinuxConfiguration *FleetAgentOsConfiguration `mandatory:"true" json:"linuxConfiguration"`

	WindowsConfiguration *FleetAgentOsConfiguration `mandatory:"true" json:"windowsConfiguration"`

	// The date and time of the last modification to the Fleet Agent Configuration (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeLastModified *common.SDKTime `mandatory:"true" json:"timeLastModified"`
}

func (m FleetAgentConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FleetAgentConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
