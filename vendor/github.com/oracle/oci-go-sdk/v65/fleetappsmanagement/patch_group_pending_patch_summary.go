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

// PatchGroupPendingPatchSummary Summary of the PatchGroup Pending Patches.
type PatchGroupPendingPatchSummary struct {

	// The OCID of the resource.
	PatchId *string `mandatory:"true" json:"patchId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	PatchName *string `mandatory:"true" json:"patchName"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	PatchDescription *string `mandatory:"true" json:"patchDescription"`

	// Patch type.
	PatchType *string `mandatory:"true" json:"patchType"`

	// Patch Severity.
	Severity PatchSeverityEnum `mandatory:"true" json:"severity"`

	// Date when the patch was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// A value determining if patch needs manual installation.
	IsManualInstallationOnly *bool `mandatory:"false" json:"isManualInstallationOnly"`
}

func (m PatchGroupPendingPatchSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchGroupPendingPatchSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetPatchSeverityEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
