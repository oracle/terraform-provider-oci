// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeClusterPlacementConstraintDetails The details for providing placement constraints for a compute cluster.
type ComputeClusterPlacementConstraintDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the HPC island for the compute cluster.
	// This field cannot be updated after creation of the compute cluster.
	HpcIslandId *string `mandatory:"false" json:"hpcIslandId"`

	// The list of target network block OCIDs to constrain placement.
	// If `targetNetworkBlockIds` is provided, the `hpcIslandId` must be set on the compute cluster,
	// and the provided network blocks must belong to that same HPC island.
	// The ordering of the array will be preserved. Ensure that all items in the array are unique.
	TargetNetworkBlockIds []string `mandatory:"false" json:"targetNetworkBlockIds"`

	// The list of target GPU memory fabric OCIDs to constrain placement.
	// If GMFs are passed in, the `hpcIslandId` must be set on the compute cluster, and
	// the provided GMFs must belong to that same HPC island.
	// The ordering of the array will be preserved. Ensure that all items in the array are unique.
	TargetMemoryFabricIds []string `mandatory:"false" json:"targetMemoryFabricIds"`

	// The logical placement strategy to apply.
	// Allowed values are `SINGLE_TIER`, `SINGLE_BLOCK`, and `PACKED_DISTRIBUTION_MULTI_BLOCK`.
	LogicalPlacementConstraint ComputeClusterLogicalPlacementConstraintEnum `mandatory:"false" json:"logicalPlacementConstraint,omitempty"`
}

func (m ComputeClusterPlacementConstraintDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeClusterPlacementConstraintDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingComputeClusterLogicalPlacementConstraintEnum(string(m.LogicalPlacementConstraint)); !ok && m.LogicalPlacementConstraint != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogicalPlacementConstraint: %s. Supported values are: %s.", m.LogicalPlacementConstraint, strings.Join(GetComputeClusterLogicalPlacementConstraintEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ComputeClusterPlacementConstraintDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeClusterPlacementConstraintDetails ComputeClusterPlacementConstraintDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeComputeClusterPlacementConstraintDetails
	}{
		"COMPUTE_CLUSTER",
		(MarshalTypeComputeClusterPlacementConstraintDetails)(m),
	}

	return json.Marshal(&s)
}
