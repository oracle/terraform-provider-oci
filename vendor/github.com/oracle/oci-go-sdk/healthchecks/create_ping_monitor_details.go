// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.cloud.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreatePingMonitorDetails The request body used to create a Ping monitor.
type CreatePingMonitorDetails struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A list of targets (hostnames or IP addresses) of the probe.
	Targets []string `mandatory:"true" json:"targets"`

	Protocol PingProbeProtocolEnum `mandatory:"true" json:"protocol"`

	// A user-friendly and mutable name suitable for display in a user interface.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The monitor interval in seconds. Valid values: 10, 30, and 60.
	IntervalInSeconds *int `mandatory:"true" json:"intervalInSeconds"`

	// A list of names of vantage points from which to execute the probe.
	VantagePointNames []string `mandatory:"false" json:"vantagePointNames"`

	// The port on which to probe endpoints. If unspecified, probes will use the
	// default port of their protocol.
	Port *int `mandatory:"false" json:"port"`

	// The probe timeout in seconds. Valid values: 10, 20, 30, and 60.
	// The probe timeout must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	// Enables or disables the monitor. Set to 'true' to launch monitoring.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace.  For more information,
	// see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreatePingMonitorDetails) String() string {
	return common.PointerString(m)
}

// CreatePingMonitorDetailsProtocolEnum is an alias to type: PingProbeProtocolEnum
// Consider using PingProbeProtocolEnum instead
// Deprecated
type CreatePingMonitorDetailsProtocolEnum = PingProbeProtocolEnum

// Set of constants representing the allowable values for PingProbeProtocolEnum
// Deprecated
const (
	CreatePingMonitorDetailsProtocolIcmp PingProbeProtocolEnum = "ICMP"
	CreatePingMonitorDetailsProtocolTcp  PingProbeProtocolEnum = "TCP"
)

// GetCreatePingMonitorDetailsProtocolEnumValues Enumerates the set of values for PingProbeProtocolEnum
// Consider using GetPingProbeProtocolEnumValue
// Deprecated
var GetCreatePingMonitorDetailsProtocolEnumValues = GetPingProbeProtocolEnumValues
