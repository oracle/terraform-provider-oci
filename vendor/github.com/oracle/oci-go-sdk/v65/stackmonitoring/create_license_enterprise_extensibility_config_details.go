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

// CreateLicenseEnterpriseExtensibilityConfigDetails The details of a LICENSE_ENTERPRISE_EXTENSIBILITY configuration.
type CreateLicenseEnterpriseExtensibilityConfigDetails struct {

	// Compartment in which the configuration is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// True if enterprise extensibility is enabled, false if it is not enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The display name of the configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetDisplayName returns DisplayName
func (m CreateLicenseEnterpriseExtensibilityConfigDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateLicenseEnterpriseExtensibilityConfigDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateLicenseEnterpriseExtensibilityConfigDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateLicenseEnterpriseExtensibilityConfigDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateLicenseEnterpriseExtensibilityConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLicenseEnterpriseExtensibilityConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateLicenseEnterpriseExtensibilityConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateLicenseEnterpriseExtensibilityConfigDetails CreateLicenseEnterpriseExtensibilityConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateLicenseEnterpriseExtensibilityConfigDetails
	}{
		"LICENSE_ENTERPRISE_EXTENSIBILITY",
		(MarshalTypeCreateLicenseEnterpriseExtensibilityConfigDetails)(m),
	}

	return json.Marshal(&s)
}
