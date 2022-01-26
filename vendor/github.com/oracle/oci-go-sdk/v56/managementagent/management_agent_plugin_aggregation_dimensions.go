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

// ManagementAgentPluginAggregationDimensions The Aggregation of Management Agent Plugin Dimensions
type ManagementAgentPluginAggregationDimensions struct {

	// Management Agent Plugin Name
	PluginName *string `mandatory:"false" json:"pluginName"`

	// Management Agent Plugin Display Name
	PluginDisplayName *string `mandatory:"false" json:"pluginDisplayName"`
}

func (m ManagementAgentPluginAggregationDimensions) String() string {
	return common.PointerString(m)
}
