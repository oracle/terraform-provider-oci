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

// LogAnalyticsEntityTopologySummary Log Analytics Entity topology that contains a set of log analytics entities and a set of relationships between those.
type LogAnalyticsEntityTopologySummary struct {
	Nodes *LogAnalyticsEntityCollection `mandatory:"true" json:"nodes"`

	Links *LogAnalyticsEntityTopologyLinkCollection `mandatory:"true" json:"links"`
}

func (m LogAnalyticsEntityTopologySummary) String() string {
	return common.PointerString(m)
}
