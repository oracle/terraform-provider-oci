// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AmdMilanBmGpuUpdateInstancePlatformConfig The platform configuration used when updating a bare metal GPU instance with the following shape: BM.GPU.GM4.8 (also
// named BM.GPU.A100-v2.8) (the AMD Milan platform).
type AmdMilanBmGpuUpdateInstancePlatformConfig struct {

	// The number of NUMA nodes per socket (NPS).
	NumaNodesPerSocket AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum `mandatory:"false" json:"numaNodesPerSocket,omitempty"`
}

func (m AmdMilanBmGpuUpdateInstancePlatformConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AmdMilanBmGpuUpdateInstancePlatformConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum(string(m.NumaNodesPerSocket)); !ok && m.NumaNodesPerSocket != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NumaNodesPerSocket: %s. Supported values are: %s.", m.NumaNodesPerSocket, strings.Join(GetAmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AmdMilanBmGpuUpdateInstancePlatformConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAmdMilanBmGpuUpdateInstancePlatformConfig AmdMilanBmGpuUpdateInstancePlatformConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAmdMilanBmGpuUpdateInstancePlatformConfig
	}{
		"AMD_MILAN_BM_GPU",
		(MarshalTypeAmdMilanBmGpuUpdateInstancePlatformConfig)(m),
	}

	return json.Marshal(&s)
}

// AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum
const (
	AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps0 AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS0"
	AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps1 AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS1"
	AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps2 AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS2"
	AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps4 AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS4"
)

var mappingAmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum = map[string]AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum{
	"NPS0": AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps0,
	"NPS1": AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps1,
	"NPS2": AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps2,
	"NPS4": AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps4,
}

var mappingAmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnumLowerCase = map[string]AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum{
	"nps0": AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps0,
	"nps1": AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps1,
	"nps2": AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps2,
	"nps4": AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps4,
}

// GetAmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnumValues Enumerates the set of values for AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum
func GetAmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnumValues() []AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum {
	values := make([]AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum, 0)
	for _, v := range mappingAmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum {
		values = append(values, v)
	}
	return values
}

// GetAmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues Enumerates the set of values in String for AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum
func GetAmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues() []string {
	return []string{
		"NPS0",
		"NPS1",
		"NPS2",
		"NPS4",
	}
}

// GetMappingAmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum(val string) (AmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum, bool) {
	enum, ok := mappingAmdMilanBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
