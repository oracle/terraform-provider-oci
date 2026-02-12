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

// PendingPatches Pending patches details.
type PendingPatches struct {

	// The patch OCID.
	PatchId *string `mandatory:"true" json:"patchId"`

	// Patch name.
	PatchName *string `mandatory:"false" json:"patchName"`

	// Patch description.
	PatchDescription *string `mandatory:"false" json:"patchDescription"`

	// Patch type.
	PatchType *string `mandatory:"false" json:"patchType"`

	// Severity.
	Severity *string `mandatory:"false" json:"severity"`

	// Patch release date.
	ReleaseDate *common.SDKTime `mandatory:"false" json:"releaseDate"`

	// A value determining if patch needs manual installation.
	IsManualInstallationOnly *bool `mandatory:"false" json:"isManualInstallationOnly"`
}

func (m PendingPatches) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PendingPatches) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
