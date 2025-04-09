// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternalRequestQueueSummary An internal representation of Request Queue Entry.
type InternalRequestQueueSummary struct {

	// The id of the resource queue
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the resource.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The type of the resource.
	ResourceType DatabaseSubTypeEnum `mandatory:"true" json:"resourceType"`

	// The entry type of resource.
	EntryType RequestQueueEntryTypeEnum `mandatory:"true" json:"entryType"`

	// The action type.
	ActionType RequestQueueActionTypeEnum `mandatory:"true" json:"actionType"`

	// The action mode.
	ActionMode RequestQueueActionModeEnum `mandatory:"true" json:"actionMode"`

	// The action target.
	ActionTarget RequestQueueActionTargetEnum `mandatory:"true" json:"actionTarget"`

	// The input details specific to a type of resource defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	InputDetails map[string]string `mandatory:"true" json:"inputDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Resource parent.
	ParentResourceId *string `mandatory:"false" json:"parentResourceId"`

	// The name of the Resource parent.
	ParentResourceName *string `mandatory:"false" json:"parentResourceName"`

	// The current lifecycle state of the request queue entry.
	LifecycleState RequestQueueLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The details of the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the request queue entry was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the request queue entry was last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The additional details specific to a type of database defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`
}

func (m InternalRequestQueueSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalRequestQueueSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseSubTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetDatabaseSubTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestQueueEntryTypeEnum(string(m.EntryType)); !ok && m.EntryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntryType: %s. Supported values are: %s.", m.EntryType, strings.Join(GetRequestQueueEntryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestQueueActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetRequestQueueActionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestQueueActionModeEnum(string(m.ActionMode)); !ok && m.ActionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionMode: %s. Supported values are: %s.", m.ActionMode, strings.Join(GetRequestQueueActionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestQueueActionTargetEnum(string(m.ActionTarget)); !ok && m.ActionTarget != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionTarget: %s. Supported values are: %s.", m.ActionTarget, strings.Join(GetRequestQueueActionTargetEnumStringValues(), ",")))
	}

	if _, ok := GetMappingRequestQueueLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRequestQueueLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
