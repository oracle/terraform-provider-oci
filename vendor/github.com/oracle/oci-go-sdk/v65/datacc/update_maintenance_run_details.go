// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateMaintenanceRunDetails Request to update maintenance run resource.
type UpdateMaintenanceRunDetails struct {

	// Database Infrastructure patching mode, either "ROLLING" or "NONROLLING". Default value is ROLLING.
	PatchingMode PatchingModeEnumEnum `mandatory:"true" json:"patchingMode"`

	// Determines the amount of time the system will wait before the start of each compute server patching operation.
	// Supported values are 15 to 120 minutes.
	CustomActionTimeoutInMins *int `mandatory:"true" json:"customActionTimeoutInMins"`

	// If true, enables configuration of customActionTimeoutInMins for this maintenance run update request.
	IsCustomActionTimeoutEnabled *bool `mandatory:"false" json:"isCustomActionTimeoutEnabled"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m UpdateMaintenanceRunDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMaintenanceRunDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchingModeEnumEnum(string(m.PatchingMode)); !ok && m.PatchingMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchingMode: %s. Supported values are: %s.", m.PatchingMode, strings.Join(GetPatchingModeEnumEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
