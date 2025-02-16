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

// IntelIcelakeBmUpdateInstancePlatformConfig The platform configuration used when updating a bare metal instance with the BM.Standard3.64 shape
// or the BM.Optimized3.36 shape (the Intel Ice Lake platform).
type IntelIcelakeBmUpdateInstancePlatformConfig struct {

	// The number of NUMA nodes per socket (NPS).
	NumaNodesPerSocket IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum `mandatory:"false" json:"numaNodesPerSocket,omitempty"`
}

func (m IntelIcelakeBmUpdateInstancePlatformConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IntelIcelakeBmUpdateInstancePlatformConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum(string(m.NumaNodesPerSocket)); !ok && m.NumaNodesPerSocket != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NumaNodesPerSocket: %s. Supported values are: %s.", m.NumaNodesPerSocket, strings.Join(GetIntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m IntelIcelakeBmUpdateInstancePlatformConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIntelIcelakeBmUpdateInstancePlatformConfig IntelIcelakeBmUpdateInstancePlatformConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeIntelIcelakeBmUpdateInstancePlatformConfig
	}{
		"INTEL_ICELAKE_BM",
		(MarshalTypeIntelIcelakeBmUpdateInstancePlatformConfig)(m),
	}

	return json.Marshal(&s)
}

// IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
const (
	IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1 IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS1"
	IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2 IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS2"
)

var mappingIntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = map[string]IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum{
	"NPS1": IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1,
	"NPS2": IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2,
}

var mappingIntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumLowerCase = map[string]IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum{
	"nps1": IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1,
	"nps2": IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2,
}

// GetIntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumValues Enumerates the set of values for IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
func GetIntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumValues() []IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum {
	values := make([]IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum, 0)
	for _, v := range mappingIntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum {
		values = append(values, v)
	}
	return values
}

// GetIntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues Enumerates the set of values in String for IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
func GetIntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues() []string {
	return []string{
		"NPS1",
		"NPS2",
	}
}

// GetMappingIntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum(val string) (IntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum, bool) {
	enum, ok := mappingIntelIcelakeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
