// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ComputeInstanceDetails Compute Instance Configuration instance details.
type ComputeInstanceDetails struct {
	BlockVolumes []InstanceConfigurationBlockVolumeDetails `mandatory:"false" json:"blockVolumes"`

	LaunchDetails *InstanceConfigurationLaunchInstanceDetails `mandatory:"false" json:"launchDetails"`

	SecondaryVnics []InstanceConfigurationAttachVnicDetails `mandatory:"false" json:"secondaryVnics"`
}

func (m ComputeInstanceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ComputeInstanceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeInstanceDetails ComputeInstanceDetails
	s := struct {
		DiscriminatorParam string `json:"instanceType"`
		MarshalTypeComputeInstanceDetails
	}{
		"compute",
		(MarshalTypeComputeInstanceDetails)(m),
	}

	return json.Marshal(&s)
}
