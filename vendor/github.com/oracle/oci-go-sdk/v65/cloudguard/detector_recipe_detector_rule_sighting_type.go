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

// DetectorRecipeDetectorRuleSightingType SightingType associated to a Threat detector rule.
type DetectorRecipeDetectorRuleSightingType struct {

	// The unique identifier of sighting type
	Id *string `mandatory:"true" json:"id"`

	// Name of the sighting type
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Mitre Att&ck tactic associated with a Sighting Type
	Tactic *string `mandatory:"true" json:"tactic"`

	// Mitre Att&ck Techniques associated with a Sighting Type
	Techniques []string `mandatory:"true" json:"techniques"`

	// Description of the sighting type
	Description *string `mandatory:"false" json:"description"`

	// Link of the sighting type
	MitreLink *string `mandatory:"false" json:"mitreLink"`

	// Identifier of the associated DataSource
	DataSourceId *string `mandatory:"false" json:"dataSourceId"`

	// Data Source additional entities mapping for a Detector Rule
	AdditionalEntitiesMapping []EntitiesMapping `mandatory:"false" json:"additionalEntitiesMapping"`

	// Owner type for the sighting type.
	Owner OwnerTypeEnum `mandatory:"false" json:"owner,omitempty"`

	MitreDetails *MitreDetails `mandatory:"false" json:"mitreDetails"`

	// The date and time the Sighting Type was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the Sighting Type was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m DetectorRecipeDetectorRuleSightingType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DetectorRecipeDetectorRuleSightingType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOwnerTypeEnum(string(m.Owner)); !ok && m.Owner != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Owner: %s. Supported values are: %s.", m.Owner, strings.Join(GetOwnerTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
