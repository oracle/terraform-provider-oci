// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddPatchesByConditionDetails Details for adding Patches to a Patch Group based on condition.
type AddPatchesByConditionDetails struct {

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ReleaseDate *ReleaseDate `mandatory:"true" json:"releaseDate"`

	// When set to true and the `compartmentId` field is set to a tenancy OCID, this option includes all patches from
	// the root compartment (the tenancy) and all its subcompartments based on access level.
	// The default value is false, which limits the results to the specified compartment only.
	CompartmentIdInSubtree *bool `mandatory:"false" json:"compartmentIdInSubtree"`

	// Filters and adds patches for the specified Product Ids (PlatformConfiguration Ids corresponding to the Products).
	Products []string `mandatory:"false" json:"products"`

	// Filters and adds patches matching the given Patch Type Id (PlatformConfiguration Id).
	PatchType *string `mandatory:"false" json:"patchType"`

	// Filters and adds patches matching the specified severity level.
	PatchSeverity PatchSeverityEnum `mandatory:"false" json:"patchSeverity,omitempty"`
}

func (m AddPatchesByConditionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddPatchesByConditionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchSeverityEnum(string(m.PatchSeverity)); !ok && m.PatchSeverity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchSeverity: %s. Supported values are: %s.", m.PatchSeverity, strings.Join(GetPatchSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
