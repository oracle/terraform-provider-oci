// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// TargetDetectorRecipeSummary Summary of DetectorRecipe
type TargetDetectorRecipeSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// DetectorRecipe Identifier Name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// DetectorRecipe Description
	Description *string `mandatory:"true" json:"description"`

	// Owner of DetectorRecipe
	Owner OwnerTypeEnum `mandatory:"true" json:"owner"`

	// Unique identifier for Detector Recipe of which this is an extension
	DetectorRecipeId *string `mandatory:"true" json:"detectorRecipeId"`

	// Type of detector
	Detector DetectorEnumEnum `mandatory:"false" json:"detector,omitempty"`

	// The current state of the resource.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the target detector recipe was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target detector recipe was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m TargetDetectorRecipeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetDetectorRecipeSummary) ValidateEnumValue() (bool, error) {
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
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
