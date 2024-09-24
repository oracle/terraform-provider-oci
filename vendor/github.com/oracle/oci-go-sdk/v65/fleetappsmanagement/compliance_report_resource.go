// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComplianceReportResource Details of the Resource
type ComplianceReportResource struct {

	// The OCID to identify the resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Display name of the resource.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// Type of the resource.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Last known compliance state of fleet.
	ComplianceState ComplianceStateEnum `mandatory:"true" json:"complianceState"`

	// TenancyId of the resource.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// Tenancy the resource belongs to.
	TenancyName *string `mandatory:"false" json:"tenancyName"`

	// Compartment the resource belongs to.
	Compartment *string `mandatory:"false" json:"compartment"`

	// Region the resource belongs to.
	ResourceRegion *string `mandatory:"false" json:"resourceRegion"`

	// Products assocaited with the Fleet.Only products belonging to managed targets will be shown.
	Products []ComplianceReportProduct `mandatory:"false" json:"products"`
}

func (m ComplianceReportResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComplianceReportResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingComplianceStateEnum(string(m.ComplianceState)); !ok && m.ComplianceState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComplianceState: %s. Supported values are: %s.", m.ComplianceState, strings.Join(GetComplianceStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
