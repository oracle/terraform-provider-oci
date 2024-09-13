// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ComplianceReportPatchDetail Details of the Patch.
type ComplianceReportPatchDetail struct {

	// Patch name.
	PatchName *string `mandatory:"true" json:"patchName"`

	// Type of patch.
	PatchType *string `mandatory:"true" json:"patchType"`

	// Patch description.
	PatchDescription *string `mandatory:"false" json:"patchDescription"`

	// Time the patch was applied.
	TimeApplied *common.SDKTime `mandatory:"false" json:"timeApplied"`

	// The date on which patch was released.
	TimeReleased *common.SDKTime `mandatory:"false" json:"timeReleased"`
}

func (m ComplianceReportPatchDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComplianceReportPatchDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
