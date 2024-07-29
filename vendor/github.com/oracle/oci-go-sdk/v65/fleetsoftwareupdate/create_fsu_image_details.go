// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateFsuImageDetails The information about the new Exadata Fleet Update Image
type CreateFsuImageDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment to contain the
	// resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Unique name of the image.
	// This name has to be specified in the exa_map file that will be available on the bucket specified in "imageDetails".
	// It cannot be the same as an existing Exadata Fleet Update Image resource.
	ImageName *string `mandatory:"true" json:"imageName"`

	ImageDetails CreateImageDetails `mandatory:"true" json:"imageDetails"`

	// Exadata Fleet Update Image display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateFsuImageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateFsuImageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateFsuImageDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName   *string                           `json:"displayName"`
		FreeformTags  map[string]string                 `json:"freeformTags"`
		DefinedTags   map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId *string                           `json:"compartmentId"`
		ImageName     *string                           `json:"imageName"`
		ImageDetails  createimagedetails                `json:"imageDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.ImageName = model.ImageName

	nn, e = model.ImageDetails.UnmarshalPolymorphicJSON(model.ImageDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ImageDetails = nn.(CreateImageDetails)
	} else {
		m.ImageDetails = nil
	}

	return
}
