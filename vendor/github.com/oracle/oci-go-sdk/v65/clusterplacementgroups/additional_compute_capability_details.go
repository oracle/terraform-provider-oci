// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cluster Placement Groups API
//
// API for managing cluster placement groups.
//

package clusterplacementgroups

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AdditionalComputeCapabilityDetails Additional details about the COMPUTE capability.
type AdditionalComputeCapabilityDetails struct {

	// The amount of memory (in GBs) needed in the instance.
	MemoryInGBs *float32 `mandatory:"false" json:"memoryInGBs"`

	// The number of OCPUs needed in the instance.
	Ocpus *float32 `mandatory:"false" json:"ocpus"`

	// The number of NVMe drives to use for storage.
	Nvmes *int `mandatory:"false" json:"nvmes"`

	// The number of instances or size of the resource.
	Count *int `mandatory:"false" json:"count"`
}

func (m AdditionalComputeCapabilityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdditionalComputeCapabilityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AdditionalComputeCapabilityDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAdditionalComputeCapabilityDetails AdditionalComputeCapabilityDetails
	s := struct {
		DiscriminatorParam string `json:"serviceType"`
		MarshalTypeAdditionalComputeCapabilityDetails
	}{
		"COMPUTE",
		(MarshalTypeAdditionalComputeCapabilityDetails)(m),
	}

	return json.Marshal(&s)
}
