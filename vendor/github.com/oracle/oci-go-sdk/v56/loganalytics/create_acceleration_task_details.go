// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateAccelerationTaskDetails Details for creating a scheduled task to accelerate a saved search.
// The client must specify the savedSearchId, and the service will supply other details.
// The resulting scheduled task will have TaskType ACCELERATION.
type CreateAccelerationTaskDetails struct {

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The ManagementSavedSearch id [OCID] to be accelerated.
	SavedSearchId *string `mandatory:"true" json:"savedSearchId"`

	// A user-friendly name that is changeable and that does not have to be unique.
	// Format: a leading alphanumeric, followed by zero or more
	// alphanumerics, underscores, spaces, backslashes, or hyphens in any order).
	// No trailing spaces allowed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetCompartmentId returns CompartmentId
func (m CreateAccelerationTaskDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m CreateAccelerationTaskDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m CreateAccelerationTaskDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateAccelerationTaskDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateAccelerationTaskDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateAccelerationTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAccelerationTaskDetails CreateAccelerationTaskDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeCreateAccelerationTaskDetails
	}{
		"ACCELERATION",
		(MarshalTypeCreateAccelerationTaskDetails)(m),
	}

	return json.Marshal(&s)
}
