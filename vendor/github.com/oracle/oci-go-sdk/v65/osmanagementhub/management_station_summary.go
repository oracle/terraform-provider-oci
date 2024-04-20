// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagementStationSummary Summary of the Management Station.
type ManagementStationSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the management station.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// User-friendly name for the management station.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Hostname of the management station.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance that is acting as the management station.
	ManagedInstanceId *string `mandatory:"false" json:"managedInstanceId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the registration profile used for the management station.
	ProfileId *string `mandatory:"false" json:"profileId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the scheduled job for the mirror sync.
	ScheduledJobId *string `mandatory:"false" json:"scheduledJobId"`

	// The date and time of the next scheduled mirror sync (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeNextExecution *common.SDKTime `mandatory:"false" json:"timeNextExecution"`

	// User-specified description of the management station.
	Description *string `mandatory:"false" json:"description"`

	// Current state of the mirror sync for the management station.
	OverallState OverallStateEnum `mandatory:"false" json:"overallState,omitempty"`

	// Overall health status of the managment station.
	HealthState HealthStateEnum `mandatory:"false" json:"healthState,omitempty"`

	// A decimal number representing the progress of the current mirror sync.
	OverallPercentage *int `mandatory:"false" json:"overallPercentage"`

	// A decimal number representing the amount of mirror capacity used by the sync.
	MirrorCapacity *int `mandatory:"false" json:"mirrorCapacity"`

	// The current state of the management station.
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
	if _, ok := GetMappingHealthStateEnum(string(m.HealthState)); !ok && m.HealthState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HealthState: %s. Supported values are: %s.", m.HealthState, strings.Join(GetHealthStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingManagementStationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetManagementStationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
