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

// ResourceProfileSummary Resource profile summary.
type ResourceProfileSummary struct {

	// Unique identifier for resource profile
	Id *string `mandatory:"true" json:"id"`

	// Unique identifier for resource profile
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Resource name for resource profile
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Resource type for resource profile
	Type *string `mandatory:"true" json:"type"`

	// Risk Score for the resource profile
	RiskScore *float64 `mandatory:"true" json:"riskScore"`

	// List of tactic summary associated with the resource profile.
	Tactics []TacticSummary `mandatory:"true" json:"tactics"`

	// The date and time the resource profile was first detected. Format defined by RFC3339.
	TimeFirstDetected *common.SDKTime `mandatory:"true" json:"timeFirstDetected"`

	// The date and time the resource profile was last detected. Format defined by RFC3339.
	TimeLastDetected *common.SDKTime `mandatory:"true" json:"timeLastDetected"`

	// Number of sightings associated with this resource profile
	SightingsCount *int `mandatory:"false" json:"sightingsCount"`

	// Number of problems associated with this resource profile
	ProblemsCount *int `mandatory:"false" json:"problemsCount"`
}

func (m ResourceProfileSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceProfileSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
