// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceMaintenanceEventSummary It is the event in which the maintenance action will be be performed on the customer instance on the scheduled date and time.
type InstanceMaintenanceEventSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance event.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the instance.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The OCID of the compartment that contains the instance.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// They are the instance actions performed during the scheduled maintenance event.
	// Maintenance actions are reboot, livemigrate, disable, terminate, rebuild-in-place, start, stop.
	InstanceAction *string `mandatory:"true" json:"instanceAction"`

	// It is the scheduled soft due date and time of the maintenance event which is going to start after this time.
	TimeNotBefore *common.SDKTime `mandatory:"true" json:"timeNotBefore"`

	// It is the scheduled soft due date and time of the maintenance event which is not going to start after this time.
	TimeNotAfter *common.SDKTime `mandatory:"true" json:"timeNotAfter"`

	// The date and time the maintenance event was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the maintenance event.
	LifecycleState InstanceMaintenanceEventLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The creator of the maintenance event.
	CreatedBy InstanceMaintenanceEventCreatedByEnum `mandatory:"true" json:"createdBy"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// It is the scheduled hard due date and time of the maintenance event.
	// The maintenance event will happen at this time and the due date will not be extended.
	TimeHardDueDate *common.SDKTime `mandatory:"false" json:"timeHardDueDate"`

	// It is the descriptive information about the maintenance taking place on the customer instance.
	Description *string `mandatory:"false" json:"description"`

	// Additional details of the maintenance in the form of json.
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`
}

func (m InstanceMaintenanceEventSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceMaintenanceEventSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInstanceMaintenanceEventLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInstanceMaintenanceEventLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInstanceMaintenanceEventCreatedByEnum(string(m.CreatedBy)); !ok && m.CreatedBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreatedBy: %s. Supported values are: %s.", m.CreatedBy, strings.Join(GetInstanceMaintenanceEventCreatedByEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
