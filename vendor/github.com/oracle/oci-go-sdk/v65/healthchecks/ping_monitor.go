// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PingMonitor A summary containing all of the mutable and immutable properties for a ping monitor.
type PingMonitor struct {

	// The OCID of the resource.
	Id *string `mandatory:"false" json:"id"`

	// A URL for fetching the probe results.
	ResultsUrl *string `mandatory:"false" json:"resultsUrl"`

	// The region where updates must be made and where results must be fetched from.
	HomeRegion *string `mandatory:"false" json:"homeRegion"`

	// The RFC 3339-formatted creation date and time of the probe.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// A list of targets (hostnames or IP addresses) of the probe.
	Targets []string `mandatory:"false" json:"targets"`

	// A list of names of vantage points from which to execute the probe.
	VantagePointNames []string `mandatory:"false" json:"vantagePointNames"`

	// The port on which to probe endpoints. If unspecified, probes will use the
	// default port of their protocol.
	Port *int `mandatory:"false" json:"port"`

	// The probe timeout in seconds. Valid values: 10, 20, 30, and 60.
	// The probe timeout must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	Protocol PingProbeProtocolEnum `mandatory:"false" json:"protocol,omitempty"`

	// A user-friendly and mutable name suitable for display in a user interface.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The monitor interval in seconds. Valid values: 10, 30, and 60.
	IntervalInSeconds *int `mandatory:"false" json:"intervalInSeconds"`

	// Enables or disables the monitor. Set to 'true' to launch monitoring.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace.  For more information,
	// see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m PingMonitor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PingMonitor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPingProbeProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetPingProbeProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PingMonitorProtocolEnum is an alias to type: PingProbeProtocolEnum
// Consider using PingProbeProtocolEnum instead
// Deprecated
type PingMonitorProtocolEnum = PingProbeProtocolEnum

// Set of constants representing the allowable values for PingProbeProtocolEnum
// Deprecated
const (
	PingMonitorProtocolIcmp PingProbeProtocolEnum = "ICMP"
	PingMonitorProtocolTcp  PingProbeProtocolEnum = "TCP"
)

// GetPingMonitorProtocolEnumValues Enumerates the set of values for PingProbeProtocolEnum
// Consider using GetPingProbeProtocolEnumValue
// Deprecated
var GetPingMonitorProtocolEnumValues = GetPingProbeProtocolEnumValues
