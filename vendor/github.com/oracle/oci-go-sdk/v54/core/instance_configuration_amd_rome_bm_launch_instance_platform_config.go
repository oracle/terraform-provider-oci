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
	"github.com/oracle/oci-go-sdk/v54/common"
	"strings"
)

// InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig The platform configuration used when launching a bare metal instance with the AMD Rome platform.
type InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig struct {

	// Whether Secure Boot is enabled on the instance.
	IsSecureBootEnabled *bool `mandatory:"false" json:"isSecureBootEnabled"`

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	IsTrustedPlatformModuleEnabled *bool `mandatory:"false" json:"isTrustedPlatformModuleEnabled"`

	// Whether the Measured Boot feature is enabled on the instance.
	IsMeasuredBootEnabled *bool `mandatory:"false" json:"isMeasuredBootEnabled"`

	// Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`.
	IsMemoryEncryptionEnabled *bool `mandatory:"false" json:"isMemoryEncryptionEnabled"`

	// The manufacturer specific technology used for memory encryption.
	MemoryEncryptionTechnology InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum `mandatory:"false" json:"memoryEncryptionTechnology,omitempty"`
}

//GetIsSecureBootEnabled returns IsSecureBootEnabled
func (m InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig) GetIsSecureBootEnabled() *bool {
	return m.IsSecureBootEnabled
}

//GetIsTrustedPlatformModuleEnabled returns IsTrustedPlatformModuleEnabled
func (m InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig) GetIsTrustedPlatformModuleEnabled() *bool {
	return m.IsTrustedPlatformModuleEnabled
}

//GetIsMeasuredBootEnabled returns IsMeasuredBootEnabled
func (m InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig) GetIsMeasuredBootEnabled() *bool {
	return m.IsMeasuredBootEnabled
}

//GetIsMemoryEncryptionEnabled returns IsMemoryEncryptionEnabled
func (m InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig) GetIsMemoryEncryptionEnabled() *bool {
	return m.IsMemoryEncryptionEnabled
}

func (m InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingInstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum[string(m.MemoryEncryptionTechnology)]; !ok && m.MemoryEncryptionTechnology != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MemoryEncryptionTechnology: %s. Supported values are: %s.", m.MemoryEncryptionTechnology, strings.Join(GetInstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeInstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig
	}{
		"AMD_ROME_BM",
		(MarshalTypeInstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig)(m),
	}

	return json.Marshal(&s)
}

// InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum Enum with underlying type: string
type InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum string

// Set of constants representing the allowable values for InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum
const (
	InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyTsme InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum = "TSME"
	InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologySmee InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum = "SMEE"
)

var mappingInstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum = map[string]InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum{
	"TSME": InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyTsme,
	"SMEE": InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologySmee,
}

// GetInstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnumValues Enumerates the set of values for InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum
func GetInstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnumValues() []InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum {
	values := make([]InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum, 0)
	for _, v := range mappingInstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnumStringValues Enumerates the set of values in String for InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum
func GetInstanceConfigurationAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnumStringValues() []string {
	return []string{
		"TSME",
		"SMEE",
	}
}
