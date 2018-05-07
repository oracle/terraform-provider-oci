// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// InstanceConfigurationInstanceDetails Instance Configuration instance details.
type InstanceConfigurationInstanceDetails struct {
	BlockVolumes []InstanceConfigurationBlockVolumeDetails `mandatory:"false" json:"blockVolumes"`

	LaunchDetails *InstanceConfigurationLaunchInstanceDetails `mandatory:"false" json:"launchDetails"`
}

func (m InstanceConfigurationInstanceDetails) String() string {
	return common.PointerString(m)
}
