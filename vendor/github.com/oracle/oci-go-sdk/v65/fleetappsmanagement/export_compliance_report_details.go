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

// ExportComplianceReportDetails Details for generating Compliance Report
type ExportComplianceReportDetails struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the entity for which the compliance is calculated.Ex.FleetId
	EntityId *string `mandatory:"false" json:"entityId"`

	// The OCID to identify the resource.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Product Stack.
	ProductStack *string `mandatory:"false" json:"productStack"`

	// Product Name.
	ProductName *string `mandatory:"false" json:"productName"`

	// Last known compliance state of target.
	ComplianceState ComplianceStateEnum `mandatory:"false" json:"complianceState,omitempty"`
}

func (m ExportComplianceReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportComplianceReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingComplianceStateEnum(string(m.ComplianceState)); !ok && m.ComplianceState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComplianceState: %s. Supported values are: %s.", m.ComplianceState, strings.Join(GetComplianceStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
