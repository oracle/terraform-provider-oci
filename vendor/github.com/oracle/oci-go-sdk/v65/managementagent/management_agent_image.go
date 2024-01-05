// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagementAgentImage Supported Agent downloads
type ManagementAgentImage struct {

	// Agent image resource id
	Id *string `mandatory:"true" json:"id"`

	// Agent image platform type
	PlatformType PlatformTypesEnum `mandatory:"true" json:"platformType"`

	// Agent image version
	Version *string `mandatory:"true" json:"version"`

	// Agent image platform display name
	PlatformName *string `mandatory:"false" json:"platformName"`

	// The installation package type
	PackageType PackageTypesEnum `mandatory:"false" json:"packageType,omitempty"`

	// The installation package target architecture type
	PackageArchitectureType ArchitectureTypesEnum `mandatory:"false" json:"packageArchitectureType,omitempty"`

	// Agent image size in bytes
	Size *float32 `mandatory:"false" json:"size"`

	// Agent image content SHA256 Hash
	Checksum *string `mandatory:"false" json:"checksum"`

	// Object storage URL for download
	ObjectUrl *string `mandatory:"false" json:"objectUrl"`

	ImageObjectStorageDetails *ObjectDetails `mandatory:"false" json:"imageObjectStorageDetails"`

	// The current state of Management Agent Image
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m ManagementAgentImage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementAgentImage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPlatformTypesEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetPlatformTypesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPackageTypesEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetPackageTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArchitectureTypesEnum(string(m.PackageArchitectureType)); !ok && m.PackageArchitectureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageArchitectureType: %s. Supported values are: %s.", m.PackageArchitectureType, strings.Join(GetArchitectureTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
