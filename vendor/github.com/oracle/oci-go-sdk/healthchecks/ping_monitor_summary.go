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

// PingMonitorSummary This model contains all of the mutable and immutable summary properties for an HTTP monitor.
type PingMonitorSummary struct {

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

	// A user-friendly and mutable name suitable for display in a user interface.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The monitor interval in seconds. Valid values: 10, 30, and 60.
	IntervalInSeconds *int `mandatory:"false" json:"intervalInSeconds"`

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

	Protocol PingProbeProtocolEnum `mandatory:"false" json:"protocol,omitempty"`
}

func (m PingMonitorSummary) String() string {
	return common.PointerString(m)
}

// PingMonitorSummaryProtocolEnum is an alias to type: PingProbeProtocolEnum
// Consider using PingProbeProtocolEnum instead
// Deprecated
type PingMonitorSummaryProtocolEnum = PingProbeProtocolEnum

// Set of constants representing the allowable values for PingProbeProtocolEnum
// Deprecated
const (
	PingMonitorSummaryProtocolIcmp PingProbeProtocolEnum = "ICMP"
	PingMonitorSummaryProtocolTcp  PingProbeProtocolEnum = "TCP"
)

// GetPingMonitorSummaryProtocolEnumValues Enumerates the set of values for PingProbeProtocolEnum
// Consider using GetPingProbeProtocolEnumValue
// Deprecated
var GetPingMonitorSummaryProtocolEnumValues = GetPingProbeProtocolEnumValues
