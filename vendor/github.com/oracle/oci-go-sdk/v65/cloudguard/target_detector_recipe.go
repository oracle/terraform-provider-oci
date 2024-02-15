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

// TargetDetectorRecipe A TargetDetectorRecipe resource contains a specific instance of one of the
// supported detector types (for example, activity, configuration, or threat)
// in which some settings can be modified specifically for a single target.
// A TargetDetectorRecipe resource:
// * Is effectively a copy of a DetectorRecipe resource in which users can make
// very limited changes if it’s Oracle-managed, and more changes if it’s user-managed.
// * Is visible on the Cloud Guard Targets, Target Details page.
// * Is located in a specific OCI compartment.
// * Can be modified by users, programmatically or through the UI.
// * Changes that can be made here override any settings in the corresponding
// DetectorRecipe, of which the TargetDetectorRecipe resource is effectively a copy,
// created when the detector recipe is attached to the target.
type TargetDetectorRecipe struct {

	// OCID for the detector recipe
	Id *string `mandatory:"true" json:"id"`

	// Display name of the detector recipe
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment OCID of the detector recipe
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Unique identifier for of original Oracle-managed detector recipe on which the TargetDetectorRecipe is based
	DetectorRecipeId *string `mandatory:"true" json:"detectorRecipeId"`

	// Owner of the detector recipe
	Owner OwnerTypeEnum `mandatory:"true" json:"owner"`

	// Type of detector
	Detector DetectorEnumEnum `mandatory:"true" json:"detector"`

	// Detector recipe description.
	Description *string `mandatory:"false" json:"description"`

	// List of detector rules for the detector recipe - user input
	DetectorRules []TargetDetectorRecipeDetectorRule `mandatory:"false" json:"detectorRules"`

	// List of currently enabled detector rules for the detector type for recipe after applying defaults
	EffectiveDetectorRules []TargetDetectorRecipeDetectorRule `mandatory:"false" json:"effectiveDetectorRules"`

	// The date and time the target detector recipe was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target detector recipe was last updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current lifecycle state of the resource
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Recipe type ( STANDARD, ENTERPRISE )
	DetectorRecipeType DetectorRecipeEnumEnum `mandatory:"false" json:"detectorRecipeType,omitempty"`

	// The number of days for which source data is retained
	SourceDataRetention *int `mandatory:"false" json:"sourceDataRetention"`
}

func (m TargetDetectorRecipe) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetDetectorRecipe) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOwnerTypeEnum(string(m.Owner)); !ok && m.Owner != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Owner: %s. Supported values are: %s.", m.Owner, strings.Join(GetOwnerTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDetectorEnumEnum(string(m.Detector)); !ok && m.Detector != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Detector: %s. Supported values are: %s.", m.Detector, strings.Join(GetDetectorEnumEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDetectorRecipeEnumEnum(string(m.DetectorRecipeType)); !ok && m.DetectorRecipeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DetectorRecipeType: %s. Supported values are: %s.", m.DetectorRecipeType, strings.Join(GetDetectorRecipeEnumEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
