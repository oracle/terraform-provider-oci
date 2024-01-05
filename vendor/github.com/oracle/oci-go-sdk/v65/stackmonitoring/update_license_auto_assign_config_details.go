// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateLicenseAutoAssignConfigDetails Change the details of a LICENSE_AUTO_ASSIGN configuration.
type UpdateLicenseAutoAssignConfigDetails struct {

	// The display name of the configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// License edition.
	License LicenseTypeEnum `mandatory:"false" json:"license,omitempty"`
}

// GetDisplayName returns DisplayName
func (m UpdateLicenseAutoAssignConfigDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m UpdateLicenseAutoAssignConfigDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateLicenseAutoAssignConfigDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateLicenseAutoAssignConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateLicenseAutoAssignConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLicenseTypeEnum(string(m.License)); !ok && m.License != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for License: %s. Supported values are: %s.", m.License, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateLicenseAutoAssignConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateLicenseAutoAssignConfigDetails UpdateLicenseAutoAssignConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeUpdateLicenseAutoAssignConfigDetails
	}{
		"LICENSE_AUTO_ASSIGN",
		(MarshalTypeUpdateLicenseAutoAssignConfigDetails)(m),
	}

	return json.Marshal(&s)
}
