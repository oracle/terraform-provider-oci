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

// ComplianceRecordDimension Aggregated summary information for ComplianceRecord
type ComplianceRecordDimension struct {

	// Last known compliance state.
	ComplianceState ComplianceStateEnum `mandatory:"true" json:"complianceState"`

	// Level at which the compliance is calculated.
	ComplianceLevel ComplianceLevelEnum `mandatory:"true" json:"complianceLevel"`
}

func (m ComplianceRecordDimension) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComplianceRecordDimension) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingComplianceStateEnum(string(m.ComplianceState)); !ok && m.ComplianceState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComplianceState: %s. Supported values are: %s.", m.ComplianceState, strings.Join(GetComplianceStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingComplianceLevelEnum(string(m.ComplianceLevel)); !ok && m.ComplianceLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComplianceLevel: %s. Supported values are: %s.", m.ComplianceLevel, strings.Join(GetComplianceLevelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
