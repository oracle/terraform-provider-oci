// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TablespaceUsageTrendAggregation Usage data per tablespace for a Pluggable database
type TablespaceUsageTrendAggregation struct {

	// The name of tablespace.
	TablespaceName *string `mandatory:"true" json:"tablespaceName"`

	// Type of tablespace
	TablespaceType *string `mandatory:"true" json:"tablespaceType"`

	// List of usage data samples for a tablespace
	UsageData []TablespaceUsageTrend `mandatory:"true" json:"usageData"`
}

func (m TablespaceUsageTrendAggregation) String() string {
	return common.PointerString(m)
}
