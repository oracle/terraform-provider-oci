// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// AppMgmt Control API
//
// AppMgmt Control API
//

package appmgmtcontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoredInstanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMonitoredInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMonitoredInstanceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMonitoredInstanceMonitoringStateEnum(string(m.MonitoringState)); !ok && m.MonitoringState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MonitoringState: %s. Supported values are: %s.", m.MonitoringState, strings.Join(GetMonitoredInstanceMonitoringStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
