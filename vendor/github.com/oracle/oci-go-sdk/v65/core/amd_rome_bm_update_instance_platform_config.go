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

// AmdRomeBmUpdateInstancePlatformConfig The platform configuration used when updating a bare metal instance with the BM.Standard.E3.128 shape
// (the AMD Rome platform).
type AmdRomeBmUpdateInstancePlatformConfig struct {

	// The number of NUMA nodes per socket (NPS).
	NumaNodesPerSocket AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum `mandatory:"false" json:"numaNodesPerSocket,omitempty"`
}

func (m AmdRomeBmUpdateInstancePlatformConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AmdRomeBmUpdateInstancePlatformConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum(string(m.NumaNodesPerSocket)); !ok && m.NumaNodesPerSocket != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NumaNodesPerSocket: %s. Supported values are: %s.", m.NumaNodesPerSocket, strings.Join(GetAmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AmdRomeBmUpdateInstancePlatformConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAmdRomeBmUpdateInstancePlatformConfig AmdRomeBmUpdateInstancePlatformConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAmdRomeBmUpdateInstancePlatformConfig
	}{
		"AMD_ROME_BM",
		(MarshalTypeAmdRomeBmUpdateInstancePlatformConfig)(m),
	}

	return json.Marshal(&s)
}

// AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
const (
	AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps0 AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS0"
	AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1 AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS1"
	AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2 AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS2"
	AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps4 AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS4"
)

var mappingAmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = map[string]AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum{
	"NPS0": AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps0,
	"NPS1": AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1,
	"NPS2": AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2,
	"NPS4": AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps4,
}

var mappingAmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumLowerCase = map[string]AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum{
	"nps0": AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps0,
	"nps1": AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1,
	"nps2": AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2,
	"nps4": AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketNps4,
}

// GetAmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumValues Enumerates the set of values for AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
func GetAmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumValues() []AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum {
	values := make([]AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum, 0)
	for _, v := range mappingAmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum {
		values = append(values, v)
	}
	return values
}

// GetAmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues Enumerates the set of values in String for AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
func GetAmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues() []string {
	return []string{
		"NPS0",
		"NPS1",
		"NPS2",
		"NPS4",
	}
}

// GetMappingAmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum(val string) (AmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum, bool) {
	enum, ok := mappingAmdRomeBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
