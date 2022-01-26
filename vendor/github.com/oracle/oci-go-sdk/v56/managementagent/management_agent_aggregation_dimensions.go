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

// ManagementAgentAggregationDimensions The Aggregation of Management Agent Dimensions
type ManagementAgentAggregationDimensions struct {

	// The availability status of managementAgent
	AvailabilityStatus AvailabilityStatusEnum `mandatory:"false" json:"availabilityStatus,omitempty"`

	// Platform Type
	PlatformType PlatformTypesEnum `mandatory:"false" json:"platformType,omitempty"`

	// Agent image version
	Version *string `mandatory:"false" json:"version"`

	// Whether or not a managementAgent has at least one plugin
	HasPlugins *bool `mandatory:"false" json:"hasPlugins"`

	// The install type, either AGENT or GATEWAY
	InstallType InstallTypesEnum `mandatory:"false" json:"installType,omitempty"`
}

func (m ManagementAgentAggregationDimensions) String() string {
	return common.PointerString(m)
}
