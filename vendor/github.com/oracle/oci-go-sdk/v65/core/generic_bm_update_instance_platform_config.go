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

// GenericBmUpdateInstancePlatformConfig The standard platform configuration to be used when updating a bare metal instance.
type GenericBmUpdateInstancePlatformConfig struct {

	// The number of NUMA nodes per socket (NPS).
	NumaNodesPerSocket GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum `mandatory:"false" json:"numaNodesPerSocket,omitempty"`
}

func (m GenericBmUpdateInstancePlatformConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenericBmUpdateInstancePlatformConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum(string(m.NumaNodesPerSocket)); !ok && m.NumaNodesPerSocket != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NumaNodesPerSocket: %s. Supported values are: %s.", m.NumaNodesPerSocket, strings.Join(GetGenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GenericBmUpdateInstancePlatformConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGenericBmUpdateInstancePlatformConfig GenericBmUpdateInstancePlatformConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGenericBmUpdateInstancePlatformConfig
	}{
		"GENERIC_BM",
		(MarshalTypeGenericBmUpdateInstancePlatformConfig)(m),
	}

	return json.Marshal(&s)
}

// GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
const (
	GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketNps0 GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS0"
	GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1 GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS1"
	GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2 GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS2"
	GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketNps4 GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = "NPS4"
)

var mappingGenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum = map[string]GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum{
	"NPS0": GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketNps0,
	"NPS1": GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1,
	"NPS2": GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2,
	"NPS4": GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketNps4,
}

var mappingGenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumLowerCase = map[string]GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum{
	"nps0": GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketNps0,
	"nps1": GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketNps1,
	"nps2": GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketNps2,
	"nps4": GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketNps4,
}

// GetGenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumValues Enumerates the set of values for GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
func GetGenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumValues() []GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum {
	values := make([]GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum, 0)
	for _, v := range mappingGenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum {
		values = append(values, v)
	}
	return values
}

// GetGenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues Enumerates the set of values in String for GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum
func GetGenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumStringValues() []string {
	return []string{
		"NPS0",
		"NPS1",
		"NPS2",
		"NPS4",
	}
}

// GetMappingGenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum(val string) (GenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnum, bool) {
	enum, ok := mappingGenericBmUpdateInstancePlatformConfigNumaNodesPerSocketEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
