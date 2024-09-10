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

// AmdRomeBmGpuUpdateInstancePlatformConfig The platform configuration used when updating a bare metal GPU instance with the BM.GPU4.8 shape
// (the AMD Rome platform).
type AmdRomeBmGpuUpdateInstancePlatformConfig struct {

	// The number of NUMA nodes per socket (NPS).
	NumaNodesPerSocket AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum `mandatory:"false" json:"numaNodesPerSocket,omitempty"`
}

func (m AmdRomeBmGpuUpdateInstancePlatformConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AmdRomeBmGpuUpdateInstancePlatformConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum(string(m.NumaNodesPerSocket)); !ok && m.NumaNodesPerSocket != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NumaNodesPerSocket: %s. Supported values are: %s.", m.NumaNodesPerSocket, strings.Join(GetAmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AmdRomeBmGpuUpdateInstancePlatformConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAmdRomeBmGpuUpdateInstancePlatformConfig AmdRomeBmGpuUpdateInstancePlatformConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAmdRomeBmGpuUpdateInstancePlatformConfig
	}{
		"AMD_ROME_BM_GPU",
		(MarshalTypeAmdRomeBmGpuUpdateInstancePlatformConfig)(m),
	}

	return json.Marshal(&s)
}

// AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum
const (
	AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps0 AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS0"
	AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps1 AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS1"
	AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps2 AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS2"
	AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps4 AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS4"
)

var mappingAmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum = map[string]AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum{
	"NPS0": AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps0,
	"NPS1": AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps1,
	"NPS2": AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps2,
	"NPS4": AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps4,
}

var mappingAmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnumLowerCase = map[string]AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum{
	"nps0": AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps0,
	"nps1": AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps1,
	"nps2": AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps2,
	"nps4": AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketNps4,
}

// GetAmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnumValues Enumerates the set of values for AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum
func GetAmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnumValues() []AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum {
	values := make([]AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum, 0)
	for _, v := range mappingAmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum {
		values = append(values, v)
	}
	return values
}

// GetAmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues Enumerates the set of values in String for AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum
func GetAmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues() []string {
	return []string{
		"NPS0",
		"NPS1",
		"NPS2",
		"NPS4",
	}
}

// GetMappingAmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum(val string) (AmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnum, bool) {
	enum, ok := mappingAmdRomeBmGpuUpdateInstancePlatformConfigNumaNodesPerSocketEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
