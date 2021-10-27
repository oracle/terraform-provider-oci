// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// LaunchInstancePlatformConfig The platform configuration requested for the instance.
// If you provide the parameter, the instance is created with the platform configuration that you specify.
// For any values that you omit, the instance uses the default configuration values for the `shape` that you
// specify. If you don't provide the parameter, the default values for the `shape` are used.
// Each shape only supports certain configurable values. If the values that you provide are not valid for the
// specified `shape`, an error is returned.
type LaunchInstancePlatformConfig interface {

	// Whether Secure Boot is enabled on the instance.
	GetIsSecureBootEnabled() *bool

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	GetIsTrustedPlatformModuleEnabled() *bool

	// Whether the Measured Boot feature is enabled on the instance.
	GetIsMeasuredBootEnabled() *bool
}

type launchinstanceplatformconfig struct {
	JsonData                       []byte
	IsSecureBootEnabled            *bool  `mandatory:"false" json:"isSecureBootEnabled"`
	IsTrustedPlatformModuleEnabled *bool  `mandatory:"false" json:"isTrustedPlatformModuleEnabled"`
	IsMeasuredBootEnabled          *bool  `mandatory:"false" json:"isMeasuredBootEnabled"`
	Type                           string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *launchinstanceplatformconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerlaunchinstanceplatformconfig launchinstanceplatformconfig
	s := struct {
		Model Unmarshalerlaunchinstanceplatformconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsSecureBootEnabled = s.Model.IsSecureBootEnabled
	m.IsTrustedPlatformModuleEnabled = s.Model.IsTrustedPlatformModuleEnabled
	m.IsMeasuredBootEnabled = s.Model.IsMeasuredBootEnabled
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *launchinstanceplatformconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "AMD_ROME_BM":
		mm := AmdRomeBmLaunchInstancePlatformConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMD_VM":
		mm := AmdVmLaunchInstancePlatformConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTEL_VM":
		mm := IntelVmLaunchInstancePlatformConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTEL_SKYLAKE_BM":
		mm := IntelSkylakeBmLaunchInstancePlatformConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMD_MILAN_BM":
		mm := AmdMilanBmLaunchInstancePlatformConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetIsSecureBootEnabled returns IsSecureBootEnabled
func (m launchinstanceplatformconfig) GetIsSecureBootEnabled() *bool {
	return m.IsSecureBootEnabled
}

//GetIsTrustedPlatformModuleEnabled returns IsTrustedPlatformModuleEnabled
func (m launchinstanceplatformconfig) GetIsTrustedPlatformModuleEnabled() *bool {
	return m.IsTrustedPlatformModuleEnabled
}

//GetIsMeasuredBootEnabled returns IsMeasuredBootEnabled
func (m launchinstanceplatformconfig) GetIsMeasuredBootEnabled() *bool {
	return m.IsMeasuredBootEnabled
}

func (m launchinstanceplatformconfig) String() string {
	return common.PointerString(m)
}

// LaunchInstancePlatformConfigTypeEnum Enum with underlying type: string
type LaunchInstancePlatformConfigTypeEnum string

// Set of constants representing the allowable values for LaunchInstancePlatformConfigTypeEnum
const (
	LaunchInstancePlatformConfigTypeAmdMilanBm     LaunchInstancePlatformConfigTypeEnum = "AMD_MILAN_BM"
	LaunchInstancePlatformConfigTypeAmdRomeBm      LaunchInstancePlatformConfigTypeEnum = "AMD_ROME_BM"
	LaunchInstancePlatformConfigTypeIntelSkylakeBm LaunchInstancePlatformConfigTypeEnum = "INTEL_SKYLAKE_BM"
	LaunchInstancePlatformConfigTypeAmdVm          LaunchInstancePlatformConfigTypeEnum = "AMD_VM"
	LaunchInstancePlatformConfigTypeIntelVm        LaunchInstancePlatformConfigTypeEnum = "INTEL_VM"
)

var mappingLaunchInstancePlatformConfigType = map[string]LaunchInstancePlatformConfigTypeEnum{
	"AMD_MILAN_BM":     LaunchInstancePlatformConfigTypeAmdMilanBm,
	"AMD_ROME_BM":      LaunchInstancePlatformConfigTypeAmdRomeBm,
	"INTEL_SKYLAKE_BM": LaunchInstancePlatformConfigTypeIntelSkylakeBm,
	"AMD_VM":           LaunchInstancePlatformConfigTypeAmdVm,
	"INTEL_VM":         LaunchInstancePlatformConfigTypeIntelVm,
}

// GetLaunchInstancePlatformConfigTypeEnumValues Enumerates the set of values for LaunchInstancePlatformConfigTypeEnum
func GetLaunchInstancePlatformConfigTypeEnumValues() []LaunchInstancePlatformConfigTypeEnum {
	values := make([]LaunchInstancePlatformConfigTypeEnum, 0)
	for _, v := range mappingLaunchInstancePlatformConfigType {
		values = append(values, v)
	}
	return values
}
