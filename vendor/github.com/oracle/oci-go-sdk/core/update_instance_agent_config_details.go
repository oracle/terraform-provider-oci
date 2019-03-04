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

// UpdateInstanceAgentConfigDetails Instance agent configuration options to choose for updating the instance
type UpdateInstanceAgentConfigDetails struct {

	// Whether the agent running on the instance can gather performance metrics and monitor the instance.
	IsMonitoringDisabled *bool `mandatory:"false" json:"isMonitoringDisabled"`
}

func (m UpdateInstanceAgentConfigDetails) String() string {
	return common.PointerString(m)
}
