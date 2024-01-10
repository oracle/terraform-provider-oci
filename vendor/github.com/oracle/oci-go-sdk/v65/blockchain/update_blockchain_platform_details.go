// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateBlockchainPlatformDetails Blockchain Platform details for updating a service.
type UpdateBlockchainPlatformDetails struct {

	// Platform Description
	Description *string `mandatory:"false" json:"description"`

	// Storage size in TBs
	StorageSizeInTBs *float64 `mandatory:"false" json:"storageSizeInTBs"`

	Replicas *ReplicaDetails `mandatory:"false" json:"replicas"`

	// Number of total OCPUs to allocate
	TotalOcpuCapacity *int `mandatory:"false" json:"totalOcpuCapacity"`

	// Type of Load Balancer shape - LB_100_MBPS or LB_400_MBPS. Default is LB_100_MBPS.
	LoadBalancerShape BlockchainPlatformLoadBalancerShapeEnum `mandatory:"false" json:"loadBalancerShape,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateBlockchainPlatformDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateBlockchainPlatformDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBlockchainPlatformLoadBalancerShapeEnum(string(m.LoadBalancerShape)); !ok && m.LoadBalancerShape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LoadBalancerShape: %s. Supported values are: %s.", m.LoadBalancerShape, strings.Join(GetBlockchainPlatformLoadBalancerShapeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
