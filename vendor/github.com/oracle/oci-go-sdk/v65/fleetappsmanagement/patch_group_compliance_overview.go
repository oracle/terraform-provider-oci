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

// PatchGroupComplianceOverview Summarized Patch Group Compliance.
type PatchGroupComplianceOverview struct {

	// The patch OCID.
	ComplianceStatus PatchGroupComplianceComplianceStatusEnum `mandatory:"true" json:"complianceStatus"`

	// Total applicable patches.
	TotalApplicablePatches *int64 `mandatory:"false" json:"totalApplicablePatches"`

	// Total pending patches.
	TotalPendingPatching *int64 `mandatory:"false" json:"totalPendingPatching"`
}

func (m PatchGroupComplianceOverview) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchGroupComplianceOverview) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchGroupComplianceComplianceStatusEnum(string(m.ComplianceStatus)); !ok && m.ComplianceStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComplianceStatus: %s. Supported values are: %s.", m.ComplianceStatus, strings.Join(GetPatchGroupComplianceComplianceStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
