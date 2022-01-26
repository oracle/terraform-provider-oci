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

// ManagementAgentPluginAggregation A count of Management Agents Plugins sharing the values for specified dimensions.
type ManagementAgentPluginAggregation struct {
	Dimensions *ManagementAgentPluginAggregationDimensions `mandatory:"false" json:"dimensions"`

	// The number of Management Agent Plugins in this group
	Count *int `mandatory:"false" json:"count"`
}

func (m ManagementAgentPluginAggregation) String() string {
	return common.PointerString(m)
}
