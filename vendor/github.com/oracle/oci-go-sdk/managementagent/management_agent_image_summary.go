// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"github.com/oracle/oci-go-sdk/common"
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
