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

// LogAnalyticsEntitySummaryReport Log-Analytics entity counts summary.
type LogAnalyticsEntitySummaryReport struct {

	// Total number of ACTIVE entities
	ActiveEntitiesCount *int `mandatory:"true" json:"activeEntitiesCount"`

	// Entities with log collection enabled
	EntitiesWithHasLogsCollectedCount *int `mandatory:"true" json:"entitiesWithHasLogsCollectedCount"`

	// Entities with management agent
	EntitiesWithManagementAgentCount *int `mandatory:"true" json:"entitiesWithManagementAgentCount"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m LogAnalyticsEntitySummaryReport) String() string {
	return common.PointerString(m)
}
