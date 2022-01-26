// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// AppMgmt Control API
//
// AppMgmt Control API
//

package appmgmtcontrol

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// MonitoredInstanceSummary Summary of the monitored instance.
type MonitoredInstanceSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of monitored instance.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name of the monitored instance. It is binded to Compute Instance (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
	// DisplayName is fetched from Core Service API (https://docs.cloud.oracle.com/api/#/en/iaas/20160918/Instance/).
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Management Agent Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	// The current state of the monitored instance.
	LifecycleState MonitoredInstanceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Monitoring status. Can be either enabled or disabled.
	MonitoringState MonitoredInstanceMonitoringStateEnum `mandatory:"false" json:"monitoringState,omitempty"`
}

func (m MonitoredInstanceSummary) String() string {
	return common.PointerString(m)
}
