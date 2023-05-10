// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateDetectorRecipeDetectorRuleSightingTypeDetails Details for updating new Sighting Type
type UpdateDetectorRecipeDetectorRuleSightingTypeDetails struct {

	// Name of the sighting type
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the sighting type
	Description *string `mandatory:"false" json:"description"`

	// Mitre attack tactic associated with a sighting Type
	Tactic *string `mandatory:"false" json:"tactic"`

	// Mitre Attack technique associated with a sighting type
	Technique *string `mandatory:"false" json:"technique"`

	// Identifier of associated dataSource
	DataSourceId *string `mandatory:"false" json:"dataSourceId"`

	// Data Source additional entities mapping for a Detector Rule
	AdditionalEntitiesMapping []EntitiesMapping `mandatory:"false" json:"additionalEntitiesMapping"`
}

func (m UpdateDetectorRecipeDetectorRuleSightingTypeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDetectorRecipeDetectorRuleSightingTypeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
