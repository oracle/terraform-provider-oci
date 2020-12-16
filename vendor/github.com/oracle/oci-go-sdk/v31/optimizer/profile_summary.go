// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// ProfileSummary The metadata associated with the profile summary.
type ProfileSummary struct {

	// The unique OCID of the profile.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy. The tenancy is the root compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name assigned to the profile.
	Name *string `mandatory:"true" json:"name"`

	// Text describing the profile.
	Description *string `mandatory:"true" json:"description"`

	// The profile's current state.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the profile was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the profile was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair applied without any predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	LevelsConfiguration *LevelsConfiguration `mandatory:"false" json:"levelsConfiguration"`
}

func (m ProfileSummary) String() string {
	return common.PointerString(m)
}
