// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementAgentAggregationDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAvailabilityStatusEnum(string(m.AvailabilityStatus)); !ok && m.AvailabilityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailabilityStatus: %s. Supported values are: %s.", m.AvailabilityStatus, strings.Join(GetAvailabilityStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPlatformTypesEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetPlatformTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInstallTypesEnum(string(m.InstallType)); !ok && m.InstallType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstallType: %s. Supported values are: %s.", m.InstallType, strings.Join(GetInstallTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
