// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagementApplianceSummary Information about management appliance.
type ManagementApplianceSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of management appliance.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartment in OCI, that this appliance is going to be created in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of SDDC in OCI, that this appliance is going to be registered in.
	SddcId *string `mandatory:"true" json:"sddcId"`

	// A descriptive name for the management appliance. It must be unique, start with a letter, and contain only letters, digits, whitespaces, dashes and underscores. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Current state of the management appliance.
	LifecycleState ManagementApplianceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	Configuration *ManagementApplianceConfiguration `mandatory:"true" json:"configuration"`

	// Array of connections for management appliance.
	Connections []ManagementApplianceConnection `mandatory:"true" json:"connections"`

	// The date and time the management appliance was created in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compute instance of management appliance in OCI.
	ComputeInstanceId *string `mandatory:"false" json:"computeInstanceId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of management agent, that this appliance is running in.
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	// Information about current lifecycleState. For FAILED and NEEDS_ATTENTION contains explanations. For other states may contain some details about their progress.
	LifecycleDetails ManagementApplianceLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The date and time the management appliance was last updated in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time the configuration of management appliance was last updated in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeConfigurationUpdated *common.SDKTime `mandatory:"false" json:"timeConfigurationUpdated"`

	// The date and time the management appliance has last received heartbeat in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeLastHeartbeat *common.SDKTime `mandatory:"false" json:"timeLastHeartbeat"`

	// Current states of connections.
	HeartbeatConnectionStates []ManagementApplianceConnectionStatus `mandatory:"false" json:"heartbeatConnectionStates"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ManagementApplianceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementApplianceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingManagementApplianceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetManagementApplianceLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingManagementApplianceLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetManagementApplianceLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
