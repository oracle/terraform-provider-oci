// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateTagDetails The representation of UpdateTagDetails
type UpdateTagDetails struct {

	// The description you assign to the tag during creation.
	Description *string `mandatory:"false" json:"description"`

	// Whether the tag is retired.
	// See Retiring Key Definitions and Namespace Definitions (https://docs.oracle.com/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm#retiringkeys).
	IsRetired *bool `mandatory:"false" json:"isRetired"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Indicates whether the tag is enabled for cost tracking.
	IsCostTracking *bool `mandatory:"false" json:"isCostTracking"`

	Validator BaseTagDefinitionValidator `mandatory:"false" json:"validator"`
}

func (m UpdateTagDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTagDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateTagDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description    *string                           `json:"description"`
		IsRetired      *bool                             `json:"isRetired"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
		IsCostTracking *bool                             `json:"isCostTracking"`
		Validator      basetagdefinitionvalidator        `json:"validator"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.IsRetired = model.IsRetired

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.IsCostTracking = model.IsCostTracking

	nn, e = model.Validator.UnmarshalPolymorphicJSON(model.Validator.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Validator = nn.(BaseTagDefinitionValidator)
	} else {
		m.Validator = nil
	}

	return
}
