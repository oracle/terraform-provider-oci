// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
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

	// Type of the Fleet.
	// PRODUCT - A fleet of product-specific resources for a product type.
	// ENVIRONMENT - A fleet of environment-specific resources for a product stack.
	// GROUP - A fleet of a fleet of either environment or product fleets.
	// GENERIC - A fleet of resources selected dynamically or manually for reporting purposes
	FleetType FleetFleetTypeEnum `mandatory:"true" json:"fleetType"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Products associated with the Fleet.
	Products []string `mandatory:"false" json:"products"`

	// Product stack associated with the Fleet.
	// Applicable for ENVIRONMENT fleet types.
	ApplicationType *string `mandatory:"false" json:"applicationType"`

	// Environment Type associated with the Fleet.
	// Applicable for ENVIRONMENT fleet types.
	EnvironmentType *string `mandatory:"false" json:"environmentType"`

	// Group Type associated with Group Fleet.
	GroupType FleetGroupTypeEnum `mandatory:"false" json:"groupType,omitempty"`

	// Type of resource selection in a Fleet.
	// Select resources manually or select resources based on rules.
	ResourceSelectionType FleetResourceSelectionTypeEnum `mandatory:"false" json:"resourceSelectionType,omitempty"`

	RuleSelectionCriteria *SelectionCriteria `mandatory:"false" json:"ruleSelectionCriteria"`

	NotificationPreferences *NotificationPreferences `mandatory:"false" json:"notificationPreferences"`

	// Resources associated with the Fleet if resourceSelectionType is MANUAL.
	Resources []AssociatedFleetResourceDetails `mandatory:"false" json:"resources"`

	// Credentials associated with the Fleet.
	Credentials []AssociatedFleetCredentialDetails `mandatory:"false" json:"credentials"`

	// A value that represents if auto-confirming of the targets can be enabled.
	// This will allow targets to be auto-confirmed in the fleet without manual intervention.
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
