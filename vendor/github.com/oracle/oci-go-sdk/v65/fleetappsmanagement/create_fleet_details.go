// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateFleetDetails The information about new Fleet.
type CreateFleetDetails struct {

	// Tenancy OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of the Fleet
	FleetType FleetFleetTypeEnum `mandatory:"true" json:"fleetType"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Products associated with the Fleet
	Products []string `mandatory:"false" json:"products"`

	// Application Type associated with the Fleet.Applicable for Environment fleet types.
	ApplicationType *string `mandatory:"false" json:"applicationType"`

	// Environment Type associated with the Fleet.Applicable for Environment fleet types.
	EnvironmentType *string `mandatory:"false" json:"environmentType"`

	// Group Type associated with Group Fleet.Applicable for Group fleet types.
	GroupType FleetGroupTypeEnum `mandatory:"false" json:"groupType,omitempty"`

	// Type of resource selection in a fleet
	ResourceSelectionType FleetResourceSelectionTypeEnum `mandatory:"false" json:"resourceSelectionType,omitempty"`

	RuleSelectionCriteria *SelectionCriteria `mandatory:"false" json:"ruleSelectionCriteria"`

	NotificationPreferences *NotificationPreferences `mandatory:"false" json:"notificationPreferences"`

	// Resources to be added during fleet creation when Resource selection type is Manual.
	Resources []AssociatedFleetResourceDetails `mandatory:"false" json:"resources"`

	// A value which represents if auto confirming of the targets can be enabled
	IsTargetAutoConfirm *bool `mandatory:"false" json:"isTargetAutoConfirm"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateFleetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateFleetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFleetFleetTypeEnum(string(m.FleetType)); !ok && m.FleetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FleetType: %s. Supported values are: %s.", m.FleetType, strings.Join(GetFleetFleetTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingFleetGroupTypeEnum(string(m.GroupType)); !ok && m.GroupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupType: %s. Supported values are: %s.", m.GroupType, strings.Join(GetFleetGroupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFleetResourceSelectionTypeEnum(string(m.ResourceSelectionType)); !ok && m.ResourceSelectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceSelectionType: %s. Supported values are: %s.", m.ResourceSelectionType, strings.Join(GetFleetResourceSelectionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
