// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchResourcesDetails Detailed specification about resources patched in a patch operation.
type PatchResourcesDetails struct {

	// PatchOperation Identifier
	DisplayName *string `mandatory:"false" json:"displayName"`

	// If true, only validates that prerequisites are met
	IsPrerequisitesOnly *bool `mandatory:"false" json:"isPrerequisitesOnly"`

	// If true, only quick validations are run. This is applicable only when isPrerequisitesOnly=true.
	IsQuickPrerequisitesCheck *bool `mandatory:"false" json:"isQuickPrerequisitesCheck"`

	// Working directory for staging binaries and temporary files
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	DeployHomesSpecification *DeployHomesSpecification `mandatory:"false" json:"deployHomesSpecification"`

	MigrateListenerSpecification *MigrateListenerSpecification `mandatory:"false" json:"migrateListenerSpecification"`

	UpdateSpecification *UpdateSpecification `mandatory:"false" json:"updateSpecification"`

	CleanupHomesSpecification *CleanupHomesSpecification `mandatory:"false" json:"cleanupHomesSpecification"`

	RollbackListenerSpecification *RollbackListenerSpecification `mandatory:"false" json:"rollbackListenerSpecification"`

	RollbackSpecification *RollbackSpecification `mandatory:"false" json:"rollbackSpecification"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m PatchResourcesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchResourcesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
