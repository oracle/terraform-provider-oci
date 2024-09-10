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

// AmdMilanBmUpdateInstancePlatformConfig The platform configuration used when updating a bare metal instance with one of the following shapes: BM.Standard.E4.128
// or BM.DenseIO.E4.128 (the AMD Milan platform).
type AmdMilanBmUpdateInstancePlatformConfig struct {

	// The number of NUMA nodes per socket (NPS).
	NumaNodesPerSocket AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum `mandatory:"false" json:"numaNodesPerSocket,omitempty"`
}

func (m AmdMilanBmUpdateInstancePlatformConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AmdMilanBmUpdateInstancePlatformConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum(string(m.NumaNodesPerSocket)); !ok && m.NumaNodesPerSocket != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NumaNodesPerSocket: %s. Supported values are: %s.", m.NumaNodesPerSocket, strings.Join(GetAmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AmdMilanBmUpdateInstancePlatformConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAmdMilanBmUpdateInstancePlatformConfig AmdMilanBmUpdateInstancePlatformConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAmdMilanBmUpdateInstancePlatformConfig
	}{
		"AMD_MILAN_BM",
		(MarshalTypeAmdMilanBmUpdateInstancePlatformConfig)(m),
	}

	return json.Marshal(&s)
}

// AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
const (
	AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketNps0 AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS0"
	AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1 AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS1"
	AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2 AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS2"
	AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketNps4 AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS4"
)

var mappingAmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = map[string]AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum{
	"NPS0": AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketNps0,
	"NPS1": AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1,
	"NPS2": AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2,
	"NPS4": AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketNps4,
}

var mappingAmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumLowerCase = map[string]AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum{
	"nps0": AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketNps0,
	"nps1": AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1,
	"nps2": AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2,
	"nps4": AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketNps4,
}

// GetAmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumValues Enumerates the set of values for AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
func GetAmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumValues() []AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum {
	values := make([]AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum, 0)
	for _, v := range mappingAmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum {
		values = append(values, v)
	}
	return values
}

// GetAmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues Enumerates the set of values in String for AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
func GetAmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues() []string {
	return []string{
		"NPS0",
		"NPS1",
		"NPS2",
		"NPS4",
	}
}

// GetMappingAmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum(val string) (AmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum, bool) {
	enum, ok := mappingAmdMilanBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
