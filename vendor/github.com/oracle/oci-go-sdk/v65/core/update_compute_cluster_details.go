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

// UpdateComputeClusterDetails The data to update a compute cluster (https://docs.oracle.com/iaas/Content/Compute/Tasks/compute-clusters.htm).
// Use `COMPUTE_CLUSTER` type when using placementConstraintDetails.
// `placementConstraintDetails.hpcIslandId` is create-only and cannot be part of this update request.
// All other fields in `placementConstraintDetails` are optional, and only the fields provided will be updated.
// If `placementConstraintDetails.targetNetworkBlockIds` or `placementConstraintDetails.targetMemoryFabricIds` is
// provided, then the target compute cluster must already have `hpcIslandId` persisted.
type UpdateComputeClusterDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	PlacementConstraintDetails PlacementConstraintDetails `mandatory:"false" json:"placementConstraintDetails"`
}

func (m UpdateComputeClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateComputeClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateComputeClusterDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                *string                           `json:"displayName"`
		DefinedTags                map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags               map[string]string                 `json:"freeformTags"`
		PlacementConstraintDetails placementconstraintdetails        `json:"placementConstraintDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	nn, e = model.PlacementConstraintDetails.UnmarshalPolymorphicJSON(model.PlacementConstraintDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PlacementConstraintDetails = nn.(PlacementConstraintDetails)
	} else {
		m.PlacementConstraintDetails = nil
	}

	return
}
