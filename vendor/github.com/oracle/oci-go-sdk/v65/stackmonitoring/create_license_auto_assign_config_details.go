// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// CreateLicenseAutoAssignConfigDetails The details of a LICENSE_AUTO_ASSIGN configuration.
type CreateLicenseAutoAssignConfigDetails struct {

	// Compartment in which the configuration is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// License edition.
	License LicenseTypeEnum `mandatory:"true" json:"license"`
}

// GetDisplayName returns DisplayName
func (m CreateLicenseAutoAssignConfigDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateLicenseAutoAssignConfigDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateLicenseAutoAssignConfigDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateLicenseAutoAssignConfigDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateLicenseAutoAssignConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLicenseAutoAssignConfigDetails) ValidateEnumValue() (bool, error) {
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
func (m CreateLicenseAutoAssignConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateLicenseAutoAssignConfigDetails CreateLicenseAutoAssignConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateLicenseAutoAssignConfigDetails
	}{
		"LICENSE_AUTO_ASSIGN",
		(MarshalTypeCreateLicenseAutoAssignConfigDetails)(m),
	}

	return json.Marshal(&s)
}
