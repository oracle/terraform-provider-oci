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

// LaunchInstanceAgentConfigDetails Instance agent configuration options to choose for launching the instance
type LaunchInstanceAgentConfigDetails struct {

	// Whether the agent running on the instance can gather performance metrics and monitor the instance.
	// Default value is false.
	IsMonitoringDisabled *bool `mandatory:"false" json:"isMonitoringDisabled"`
}

func (m LaunchInstanceAgentConfigDetails) String() string {
	return common.PointerString(m)
}
