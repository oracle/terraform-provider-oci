// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReleaseVersionSummary Customer-facing release version catalog entry.
type ReleaseVersionSummary struct {

	// Release version key
	ReleaseVersionKey *string `mandatory:"true" json:"releaseVersionKey"`

	// Display name for the release version.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Major version family grouping (for example, EXA_OL_8).
	MajorFamily *string `mandatory:"true" json:"majorFamily"`

	// Status of the release version.
	Status ReleaseVersionStatusEnum `mandatory:"true" json:"status"`

	// Whether this entry is the latest supported release version for the category.
	IsLatest *bool `mandatory:"true" json:"isLatest"`
}

func (m ReleaseVersionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReleaseVersionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReleaseVersionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetReleaseVersionStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
