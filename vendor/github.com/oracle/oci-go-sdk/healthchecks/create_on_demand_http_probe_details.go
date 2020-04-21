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

// CreateOnDemandHttpProbeDetails The request body used to create an on-demand HTTP probe.
type CreateOnDemandHttpProbeDetails struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A list of targets (hostnames or IP addresses) of the probe.
	Targets []string `mandatory:"true" json:"targets"`

	Protocol HttpProbeProtocolEnum `mandatory:"true" json:"protocol"`

	// A list of names of vantage points from which to execute the probe.
	VantagePointNames []string `mandatory:"false" json:"vantagePointNames"`

	// The port on which to probe endpoints. If unspecified, probes will use the
	// default port of their protocol.
	Port *int `mandatory:"false" json:"port"`

	// The probe timeout in seconds. Valid values: 10, 20, 30, and 60.
	// The probe timeout must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	Method HttpProbeMethodEnum `mandatory:"false" json:"method,omitempty"`

	// The optional URL path to probe, including query parameters.
	Path *string `mandatory:"false" json:"path"`

	// A dictionary of HTTP request headers.
	// *Note:* Monitors and probes do not support the use of the `Authorization` HTTP header.
	Headers map[string]string `mandatory:"false" json:"headers"`
}

func (m CreateOnDemandHttpProbeDetails) String() string {
	return common.PointerString(m)
}

// CreateOnDemandHttpProbeDetailsProtocolEnum is an alias to type: HttpProbeProtocolEnum
// Consider using HttpProbeProtocolEnum instead
// Deprecated
type CreateOnDemandHttpProbeDetailsProtocolEnum = HttpProbeProtocolEnum

// Set of constants representing the allowable values for HttpProbeProtocolEnum
// Deprecated
const (
	CreateOnDemandHttpProbeDetailsProtocolHttp  HttpProbeProtocolEnum = "HTTP"
	CreateOnDemandHttpProbeDetailsProtocolHttps HttpProbeProtocolEnum = "HTTPS"
)

// GetCreateOnDemandHttpProbeDetailsProtocolEnumValues Enumerates the set of values for HttpProbeProtocolEnum
// Consider using GetHttpProbeProtocolEnumValue
// Deprecated
var GetCreateOnDemandHttpProbeDetailsProtocolEnumValues = GetHttpProbeProtocolEnumValues

// CreateOnDemandHttpProbeDetailsMethodEnum is an alias to type: HttpProbeMethodEnum
// Consider using HttpProbeMethodEnum instead
// Deprecated
type CreateOnDemandHttpProbeDetailsMethodEnum = HttpProbeMethodEnum

// Set of constants representing the allowable values for HttpProbeMethodEnum
// Deprecated
const (
	CreateOnDemandHttpProbeDetailsMethodGet  HttpProbeMethodEnum = "GET"
	CreateOnDemandHttpProbeDetailsMethodHead HttpProbeMethodEnum = "HEAD"
)

// GetCreateOnDemandHttpProbeDetailsMethodEnumValues Enumerates the set of values for HttpProbeMethodEnum
// Consider using GetHttpProbeMethodEnumValue
// Deprecated
var GetCreateOnDemandHttpProbeDetailsMethodEnumValues = GetHttpProbeMethodEnumValues
