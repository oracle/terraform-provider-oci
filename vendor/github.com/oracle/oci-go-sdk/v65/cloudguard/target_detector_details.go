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

// TargetDetectorDetails Overriden settings of a detector rule in recipe attached to target.
type TargetDetectorDetails struct {

	// Enablement state of the detector rule
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The risk level of the detector rule
	RiskLevel RiskLevelEnum `mandatory:"false" json:"riskLevel,omitempty"`

	// List of detector rule configurations
	Configurations []DetectorConfiguration `mandatory:"false" json:"configurations"`

	// Condition group corresponding to each compartment
	ConditionGroups []ConditionGroup `mandatory:"false" json:"conditionGroups"`

	// User-defined labels for a detector rule
	Labels []string `mandatory:"false" json:"labels"`

	// Configuration allowed or not
	IsConfigurationAllowed *bool `mandatory:"false" json:"isConfigurationAllowed"`

	// Point at which an elevated resource risk score creates a problem
	ProblemThreshold *int `mandatory:"false" json:"problemThreshold"`

	// List of target types for which the detector rule is applicable
	TargetTypes []string `mandatory:"false" json:"targetTypes"`

	// List of sighting types
	SightingTypes []SightingType `mandatory:"false" json:"sightingTypes"`
}

func (m TargetDetectorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetDetectorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRiskLevelEnum(string(m.RiskLevel)); !ok && m.RiskLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RiskLevel: %s. Supported values are: %s.", m.RiskLevel, strings.Join(GetRiskLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
