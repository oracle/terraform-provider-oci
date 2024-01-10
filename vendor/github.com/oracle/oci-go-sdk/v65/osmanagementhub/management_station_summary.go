// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagementStationSummary Summary of the Management Station.
type ManagementStationSummary struct {

	// OCID for the Management Station
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy containing the Management Station.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// ManagementStation name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Name of the host
	Hostname *string `mandatory:"true" json:"hostname"`

	// OCID for the Instance associated with the Management Station
	ManagedInstanceId *string `mandatory:"false" json:"managedInstanceId"`

	// OCID of the Registration Profile associated with the Management Station
	ProfileId *string `mandatory:"false" json:"profileId"`

	// OCID of the Scheduled Job for mirror sync
	ScheduledJobId *string `mandatory:"false" json:"scheduledJobId"`

	// the time/date of the next scheduled execution of the Scheduled Job
	TimeNextExecution *common.SDKTime `mandatory:"false" json:"timeNextExecution"`

	// Details describing the Management Station config.
	Description *string `mandatory:"false" json:"description"`

	// Current state of the mirroring
	OverallState OverallStateEnum `mandatory:"false" json:"overallState,omitempty"`

	// A decimal number representing the completeness percentage
	OverallPercentage *int `mandatory:"false" json:"overallPercentage"`

	// A decimal number representing the mirror capacity
	MirrorCapacity *int `mandatory:"false" json:"mirrorCapacity"`

	// The current state of the Management Station config.
	LifecycleState ManagementStationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ManagementStationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementStationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOverallStateEnum(string(m.OverallState)); !ok && m.OverallState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OverallState: %s. Supported values are: %s.", m.OverallState, strings.Join(GetOverallStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingManagementStationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetManagementStationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
