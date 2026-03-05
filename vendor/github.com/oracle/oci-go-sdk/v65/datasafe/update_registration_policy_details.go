// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateRegistrationPolicyDetails The details required to update an existing registration policy.
// This object contains the necessary information to update a registration policy, including the display name, features, and OnPrem connection IDs
type UpdateRegistrationPolicyDetails struct {

	// The display name of the registration policy.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description of the registration policy.
	Description *string `mandatory:"false" json:"description"`

	// The Data Safe features granted to the databases registering under the registration policy.
	Features []RegistrationPolicyFeaturesEnum `mandatory:"false" json:"features,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	ConnectionOption *PolicyConnectionOption `mandatory:"false" json:"connectionOption"`
}

func (m UpdateRegistrationPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateRegistrationPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.Features {
		if _, ok := GetMappingRegistrationPolicyFeaturesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Features: %s. Supported values are: %s.", val, strings.Join(GetRegistrationPolicyFeaturesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
