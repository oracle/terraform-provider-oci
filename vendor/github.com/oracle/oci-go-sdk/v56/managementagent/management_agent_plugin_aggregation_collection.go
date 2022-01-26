// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ManagementAgentPluginAggregationCollection The summary of Management Agent Plugin count items
type ManagementAgentPluginAggregationCollection struct {

	// List in which each item describes an aggregation of Managment Agent Plugins
	Items []ManagementAgentPluginAggregation `mandatory:"true" json:"items"`
}

func (m ManagementAgentPluginAggregationCollection) String() string {
	return common.PointerString(m)
}
