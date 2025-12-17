// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateBatchTaskEnvironmentDetails The data to create a batch task environment.
// If the value for a collection is absent or is explicitly provided as null, it will be converted to an empty value, i.e. "[]" or "{}" in json notation. This applies to nested collections as well.
type CreateBatchTaskEnvironmentDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The URL of the ocir image.
	ImageUrl *string `mandatory:"true" json:"imageUrl"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// If not specified or provided as null or empty string, it be generated as "<resourceType><timeCreated>", where timeCreated corresponds with the resource creation time in ISO 8601 basic format, i.e. omitting separating punctuation, at second-level precision and no UTC offset. Example: batchtaskenvironment20250914115623.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The batch task environment description.
	Description *string `mandatory:"false" json:"description"`

	SecurityContext *SecurityContext `mandatory:"false" json:"securityContext"`

	// Container's working directory.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// List of volumes attached to the image.
	// The use cases of the volumes are but not limited to: read the input of the task and write the output.
	Volumes []BatchTaskEnvironmentVolume `mandatory:"false" json:"volumes"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateBatchTaskEnvironmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBatchTaskEnvironmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateBatchTaskEnvironmentDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                           `json:"displayName"`
		Description      *string                           `json:"description"`
		SecurityContext  *SecurityContext                  `json:"securityContext"`
		WorkingDirectory *string                           `json:"workingDirectory"`
		Volumes          []batchtaskenvironmentvolume      `json:"volumes"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		CompartmentId    *string                           `json:"compartmentId"`
		ImageUrl         *string                           `json:"imageUrl"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.SecurityContext = model.SecurityContext

	m.WorkingDirectory = model.WorkingDirectory

	m.Volumes = make([]BatchTaskEnvironmentVolume, len(model.Volumes))
	for i, n := range model.Volumes {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Volumes[i] = nn.(BatchTaskEnvironmentVolume)
		} else {
			m.Volumes[i] = nil
		}
	}
	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.CompartmentId = model.CompartmentId

	m.ImageUrl = model.ImageUrl

	return
}
