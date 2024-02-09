// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VirtualNodePoolCyclingDetails Virtual Node Pool Cycling Details
type VirtualNodePoolCyclingDetails struct {

	// Maximum active virtual nodes that would be terminated from virtual nodepool during the cycling virtual nodepool process.
	// OKE supports both integer and percentage input.
	// Defaults to 0, Ranges from 0 to Virtual Nodepool size or 0% to 100%
	MaximumUnavailable *string `mandatory:"false" json:"maximumUnavailable"`

	// Maximum additional new virtual nodes that would be temporarily created and added to virtual nodepool during the cycling virtual nodepool process.
	// OKE supports both integer and percentage input.
	// Defaults to 1, Ranges from 0 to Virtual Nodepool size or 0% to 100%
	MaximumSurge *string `mandatory:"false" json:"maximumSurge"`

	// If virtual nodes in the virtual nodepool will be cycled to have new changes.
	IsVirtualNodeCyclingEnabled *bool `mandatory:"false" json:"isVirtualNodeCyclingEnabled"`
}

func (m VirtualNodePoolCyclingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VirtualNodePoolCyclingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
