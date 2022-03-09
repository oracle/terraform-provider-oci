// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig The platform configuration used when launching a bare metal instance with the Intel Icelake platform.
type InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig struct {

	// Whether Secure Boot is enabled on the instance.
	IsSecureBootEnabled *bool `mandatory:"false" json:"isSecureBootEnabled"`

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	IsTrustedPlatformModuleEnabled *bool `mandatory:"false" json:"isTrustedPlatformModuleEnabled"`

	// Whether the Measured Boot feature is enabled on the instance.
	IsMeasuredBootEnabled *bool `mandatory:"false" json:"isMeasuredBootEnabled"`

	// Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`.
	IsMemoryEncryptionEnabled *bool `mandatory:"false" json:"isMemoryEncryptionEnabled"`

	// Whether symmetric multi-threading is enabled on the instance.
	IsSymmetricMultiThreadingEnabled *bool `mandatory:"false" json:"isSymmetricMultiThreadingEnabled"`

	// Whether the input-output memory management unit is enabled.
	IsInputOutputMemoryManagementUnitEnabled *bool `mandatory:"false" json:"isInputOutputMemoryManagementUnitEnabled"`

	// The percentage of cores enabled.
	PercentageOfCoresEnabled *int `mandatory:"false" json:"percentageOfCoresEnabled"`

	// The number of NUMA nodes per socket (NPS).
	NumaNodesPerSocket InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum `mandatory:"false" json:"numaNodesPerSocket,omitempty"`
}

//GetIsSecureBootEnabled returns IsSecureBootEnabled
func (m InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig) GetIsSecureBootEnabled() *bool {
	return m.IsSecureBootEnabled
}

//GetIsTrustedPlatformModuleEnabled returns IsTrustedPlatformModuleEnabled
func (m InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig) GetIsTrustedPlatformModuleEnabled() *bool {
	return m.IsTrustedPlatformModuleEnabled
}

//GetIsMeasuredBootEnabled returns IsMeasuredBootEnabled
func (m InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig) GetIsMeasuredBootEnabled() *bool {
	return m.IsMeasuredBootEnabled
}

//GetIsMemoryEncryptionEnabled returns IsMemoryEncryptionEnabled
func (m InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig) GetIsMemoryEncryptionEnabled() *bool {
	return m.IsMemoryEncryptionEnabled
}

func (m InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum(string(m.NumaNodesPerSocket)); !ok && m.NumaNodesPerSocket != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NumaNodesPerSocket: %s. Supported values are: %s.", m.NumaNodesPerSocket, strings.Join(GetInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig
	}{
		"INTEL_ICELAKE_BM",
		(MarshalTypeInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfig)(m),
	}

	return json.Marshal(&s)
}

// InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum
const (
	InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketNps1 InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum = "NPS1"
	InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketNps2 InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum = "NPS2"
)

var mappingInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum = map[string]InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum{
	"NPS1": InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketNps1,
	"NPS2": InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketNps2,
}

var mappingInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnumLowerCase = map[string]InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum{
	"nps1": InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketNps1,
	"nps2": InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketNps2,
}

// GetInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnumValues Enumerates the set of values for InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum
func GetInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnumValues() []InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum {
	values := make([]InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum, 0)
	for _, v := range mappingInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnumStringValues Enumerates the set of values in String for InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum
func GetInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnumStringValues() []string {
	return []string{
		"NPS1",
		"NPS2",
	}
}

// GetMappingInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum(val string) (InstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum, bool) {
	enum, ok := mappingInstanceConfigurationIntelIcelakeBmLaunchInstancePlatformConfigNumaNodesPerSocketEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
