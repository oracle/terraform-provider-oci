// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceProfileImpactedResourceSummary Resource Profile impacted resource summary.
type ResourceProfileImpactedResourceSummary struct {

	// Unique identifier for impacted resource
	Id *string `mandatory:"true" json:"id"`

	// Resource profile Id associated with the imacted resource
	ResourceProfileId *string `mandatory:"true" json:"resourceProfileId"`

	// Compartment Id for impacted resource
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Impacted resource Id
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Resource name
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// Resource type
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Identifier for the sighting type
	SightingType *string `mandatory:"true" json:"sightingType"`

	// Name of the sighting type
	SightingTypeDisplayName *string `mandatory:"true" json:"sightingTypeDisplayName"`

	// Region for impacted resource
	Region *string `mandatory:"true" json:"region"`

	// Time when the impacted resource is identified for given resource profile.
	TimeIdentified *common.SDKTime `mandatory:"true" json:"timeIdentified"`

	// Problem Id for impacted resource
	ProblemId *string `mandatory:"false" json:"problemId"`
}

func (m ResourceProfileImpactedResourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceProfileImpactedResourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
