// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// IntelSkylakeBmUpdateInstancePlatformConfig The platform configuration used when updating a bare metal instance with an Intel X7-based processor
// (the Intel Skylake platform).
type IntelSkylakeBmUpdateInstancePlatformConfig struct {

	// The number of NUMA nodes per socket (NPS).
	NumaNodesPerSocket IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum `mandatory:"false" json:"numaNodesPerSocket,omitempty"`
}

func (m IntelSkylakeBmUpdateInstancePlatformConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IntelSkylakeBmUpdateInstancePlatformConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum(string(m.NumaNodesPerSocket)); !ok && m.NumaNodesPerSocket != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NumaNodesPerSocket: %s. Supported values are: %s.", m.NumaNodesPerSocket, strings.Join(GetIntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m IntelSkylakeBmUpdateInstancePlatformConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIntelSkylakeBmUpdateInstancePlatformConfig IntelSkylakeBmUpdateInstancePlatformConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeIntelSkylakeBmUpdateInstancePlatformConfig
	}{
		"INTEL_SKYLAKE_BM",
		(MarshalTypeIntelSkylakeBmUpdateInstancePlatformConfig)(m),
	}

	return json.Marshal(&s)
}

// IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
const (
	IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1 IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS1"
	IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2 IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS2"
)

var mappingIntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = map[string]IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum{
	"NPS1": IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1,
	"NPS2": IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2,
}

var mappingIntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumLowerCase = map[string]IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum{
	"nps1": IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1,
	"nps2": IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2,
}

// GetIntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumValues Enumerates the set of values for IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
func GetIntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumValues() []IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum {
	values := make([]IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum, 0)
	for _, v := range mappingIntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum {
		values = append(values, v)
	}
	return values
}

// GetIntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues Enumerates the set of values in String for IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
func GetIntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues() []string {
	return []string{
		"NPS1",
		"NPS2",
	}
}

// GetMappingIntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum(val string) (IntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum, bool) {
	enum, ok := mappingIntelSkylakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
