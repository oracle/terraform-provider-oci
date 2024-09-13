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

// FleetResourceSummary Summary of the FleetResource.
type FleetResourceSummary struct {

	// The unique id of the resource.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the tenancy to which the resource belongs to.
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Type of the Resource.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The current state of the FleetResource.
	LifecycleState FleetResourceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Associated region
	ResourceRegion *string `mandatory:"false" json:"resourceRegion"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Compliance State of the Resource.
	ComplianceState ComplianceStateEnum `mandatory:"false" json:"complianceState,omitempty"`

	// Resource Tenancy Name.
	TenancyName *string `mandatory:"false" json:"tenancyName"`

	// Resource Compartment name.
	Compartment *string `mandatory:"false" json:"compartment"`

	// Count of products within the resource.
	ProductCount *int `mandatory:"false" json:"productCount"`

	// Count of targets within the resource.
	TargetCount *int `mandatory:"false" json:"targetCount"`

	// Product associated with the resource when the resource type is fleet.
	// Will only be returned for PRODUCT fleets that are part of a GROUP Fleet.
	Product *string `mandatory:"false" json:"product"`

	// Application Type associated with the resource when the resource type is fleet.
	// Will only be returned for ENVIRONMENT fleets that are part of a GROUP Fleet.
	ApplicationType *string `mandatory:"false" json:"applicationType"`

	// Environment Type associated with the Fleet when the resource type is fleet.
	// Will only be returned for ENVIRONMENT fleets that are part of a GROUP Fleet.
	EnvironmentType *string `mandatory:"false" json:"environmentType"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m FleetResourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FleetResourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFleetResourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFleetResourceLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingComplianceStateEnum(string(m.ComplianceState)); !ok && m.ComplianceState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComplianceState: %s. Supported values are: %s.", m.ComplianceState, strings.Join(GetComplianceStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
