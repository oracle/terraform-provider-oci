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

// LogAnalyticsEntityTopologyLink Log Analytics entity relationship link used in hierarchical representation of entity relationships topology.
type LogAnalyticsEntityTopologyLink struct {

	// The log analytics entity OCID. This ID is a reference used by log analytics features and it represents
	// a resource that is provisioned and managed by the customer on their premises or on the cloud.
	SourceEntityId *string `mandatory:"true" json:"sourceEntityId"`

	// The log analytics entity OCID. This ID is a reference used by log analytics features and it represents
	// a resource that is provisioned and managed by the customer on their premises or on the cloud.
	DestinationEntityId *string `mandatory:"true" json:"destinationEntityId"`
}

func (m LogAnalyticsEntityTopologyLink) String() string {
	return common.PointerString(m)
}
