// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceProfileSummary Summary information for a resource profile.
type ResourceProfileSummary struct {

	// Unique identifier for the resource profile
	Id *string `mandatory:"true" json:"id"`

	// Unique identifier for the resource associated with the resource profile
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Display name for the resource profile
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Resource type for the resource profile
	Type *string `mandatory:"true" json:"type"`

	// Risk score for the resource profile
	RiskScore *float64 `mandatory:"true" json:"riskScore"`

	// List of tactic summaries associated with the resource profile
	Tactics []TacticSummary `mandatory:"true" json:"tactics"`

	// Time the activities were first detected. Format defined by RFC3339.
	TimeFirstDetected *common.SDKTime `mandatory:"true" json:"timeFirstDetected"`

	// Time the activities were last detected. Format defined by RFC3339.
	TimeLastDetected *common.SDKTime `mandatory:"true" json:"timeLastDetected"`

	// Number of sightings associated with the resource profile
	SightingsCount *int `mandatory:"false" json:"sightingsCount"`

	// Time the activities were first performed. Format defined by RFC3339.
	TimeFirstOccurred *common.SDKTime `mandatory:"false" json:"timeFirstOccurred"`

	// Time the activities were last performed. Format defined by RFC3339.
	TimeLastOccurred *common.SDKTime `mandatory:"false" json:"timeLastOccurred"`

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
