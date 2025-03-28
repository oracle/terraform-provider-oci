// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cluster Placement Groups API
//
// API for managing cluster placement groups.
//

package clusterplacementgroups

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateClusterPlacementGroupDetails Detailed information about the new cluster placement group.
type CreateClusterPlacementGroupDetails struct {

	// The friendly name of the cluster placement group.
	Name *string `mandatory:"true" json:"name"`

	// ClusterPlacementGroup Identifier.
	ClusterPlacementGroupType ClusterPlacementGroupTypeEnum `mandatory:"true" json:"clusterPlacementGroupType"`

	// A description of the cluster placement group.
	Description *string `mandatory:"true" json:"description"`

	// The availability domain where you want to create the cluster placement group.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the cluster placement group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	PlacementInstruction *PlacementInstructionDetails `mandatory:"false" json:"placementInstruction"`

	Capabilities *CapabilitiesCollection `mandatory:"false" json:"capabilities"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateClusterPlacementGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateClusterPlacementGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingClusterPlacementGroupTypeEnum(string(m.ClusterPlacementGroupType)); !ok && m.ClusterPlacementGroupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterPlacementGroupType: %s. Supported values are: %s.", m.ClusterPlacementGroupType, strings.Join(GetClusterPlacementGroupTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
