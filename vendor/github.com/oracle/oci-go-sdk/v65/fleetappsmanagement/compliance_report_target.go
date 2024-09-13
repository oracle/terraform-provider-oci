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

// ComplianceReportTarget Details of the target and patches.
type ComplianceReportTarget struct {

	// Target Identifier.Can be the target name if a separate ID is not available.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Target Name.
	TargetName *string `mandatory:"true" json:"targetName"`

	// The last known compliance state of the target.
	ComplianceState ComplianceStateEnum `mandatory:"true" json:"complianceState"`

	// Current version of the target.
	Version *string `mandatory:"false" json:"version"`

	// Installed Patches for the Target.
	InstalledPatches []ComplianceReportPatchDetail `mandatory:"false" json:"installedPatches"`

	// Recommended Patches for the Target.
	RecommendedPatches []ComplianceReportPatchDetail `mandatory:"false" json:"recommendedPatches"`
}

func (m ComplianceReportTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComplianceReportTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingComplianceStateEnum(string(m.ComplianceState)); !ok && m.ComplianceState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComplianceState: %s. Supported values are: %s.", m.ComplianceState, strings.Join(GetComplianceStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
