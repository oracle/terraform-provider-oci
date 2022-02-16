// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ManagementAgentImageSummary Supported Agent downloads
type ManagementAgentImageSummary struct {

	// Agent image resource id
	Id *string `mandatory:"true" json:"id"`

	// Agent image platform type
	PlatformType PlatformTypesEnum `mandatory:"true" json:"platformType"`

	// Agent image version
	Version *string `mandatory:"true" json:"version"`

	// Agent image platform display name
	PlatformName *string `mandatory:"false" json:"platformName"`

	// Agent image size in bytes
	Size *float32 `mandatory:"false" json:"size"`

	// Agent image content SHA256 Hash
	Checksum *string `mandatory:"false" json:"checksum"`

	// Object storage URL for download
	ObjectUrl *string `mandatory:"false" json:"objectUrl"`

	// The current state of Management Agent Image
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m ManagementAgentImageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementAgentImageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPlatformTypesEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetPlatformTypesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
