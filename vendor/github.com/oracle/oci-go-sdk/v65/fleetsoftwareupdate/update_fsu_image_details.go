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

// UpdateFsuImageDetails The information to update an Exadata Fleet Update Image resource.
type UpdateFsuImageDetails struct {

	// Unique Name of the image.
	// This name has to be specified in the exa_map file that will be available on the bucket specified in "imageDetails".
	ImageName *string `mandatory:"false" json:"imageName"`

	ImageDetails UpdateImageDetails `mandatory:"false" json:"imageDetails"`

	// Exadata Fleet Update Image resource display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateFsuImageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateFsuImageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateFsuImageDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ImageName    *string                           `json:"imageName"`
		ImageDetails updateimagedetails                `json:"imageDetails"`
		DisplayName  *string                           `json:"displayName"`
		FreeformTags map[string]string                 `json:"freeformTags"`
		DefinedTags  map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ImageName = model.ImageName

	nn, e = model.ImageDetails.UnmarshalPolymorphicJSON(model.ImageDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ImageDetails = nn.(UpdateImageDetails)
	} else {
		m.ImageDetails = nil
	}

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
