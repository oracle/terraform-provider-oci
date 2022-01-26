// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// LogAnalyticsEmBridgeSummaryReport Log-Analytics EM Bridge counts summary.
type LogAnalyticsEmBridgeSummaryReport struct {

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Total number of ACTIVE enterprise manager bridges.
	ActiveEmBridgeCount *int `mandatory:"true" json:"activeEmBridgeCount"`

	// Number of enterprise manager bridges in CREATING state.
	CreatingEmBridgeCount *int `mandatory:"true" json:"creatingEmBridgeCount"`

	// Number of enterprise manager bridges in NEEDS_ATTENTION state.
	NeedsAttentionEmBridgeCount *int `mandatory:"true" json:"needsAttentionEmBridgeCount"`

	// Number of enterprise manager bridges in DELETED state.
	DeletedEmBridgeCount *int `mandatory:"true" json:"deletedEmBridgeCount"`

	// Total number of enterprise manager bridges.
	TotalEmBridgeCount *int `mandatory:"true" json:"totalEmBridgeCount"`
}

func (m LogAnalyticsEmBridgeSummaryReport) String() string {
	return common.PointerString(m)
}
