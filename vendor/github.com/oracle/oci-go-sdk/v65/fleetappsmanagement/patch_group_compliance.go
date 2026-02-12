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

// PatchGroupCompliance Patch group details.
type PatchGroupCompliance struct {

	// Target Id.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Target name.
	TargetName *string `mandatory:"false" json:"targetName"`

	// The resource OCID.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Resource name.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// Product Id.
	ProductId *string `mandatory:"false" json:"productId"`

	// Product name.
	ProductName *string `mandatory:"false" json:"productName"`

	// Product version.
	ProductVersion *string `mandatory:"false" json:"productVersion"`

	// Compliance status.
	ComplianceStatus PatchGroupComplianceComplianceStatusEnum `mandatory:"false" json:"complianceStatus,omitempty"`

	// List of pending patches.
	PendingPatches []PendingPatches `mandatory:"false" json:"pendingPatches"`
}

func (m PatchGroupCompliance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchGroupCompliance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchGroupComplianceComplianceStatusEnum(string(m.ComplianceStatus)); !ok && m.ComplianceStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComplianceStatus: %s. Supported values are: %s.", m.ComplianceStatus, strings.Join(GetPatchGroupComplianceComplianceStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchGroupComplianceComplianceStatusEnum Enum with underlying type: string
type PatchGroupComplianceComplianceStatusEnum string

// Set of constants representing the allowable values for PatchGroupComplianceComplianceStatusEnum
const (
	PatchGroupComplianceComplianceStatusCompliant    PatchGroupComplianceComplianceStatusEnum = "COMPLIANT"
	PatchGroupComplianceComplianceStatusNonCompliant PatchGroupComplianceComplianceStatusEnum = "NON_COMPLIANT"
)

var mappingPatchGroupComplianceComplianceStatusEnum = map[string]PatchGroupComplianceComplianceStatusEnum{
	"COMPLIANT":     PatchGroupComplianceComplianceStatusCompliant,
	"NON_COMPLIANT": PatchGroupComplianceComplianceStatusNonCompliant,
}

var mappingPatchGroupComplianceComplianceStatusEnumLowerCase = map[string]PatchGroupComplianceComplianceStatusEnum{
	"compliant":     PatchGroupComplianceComplianceStatusCompliant,
	"non_compliant": PatchGroupComplianceComplianceStatusNonCompliant,
}

// GetPatchGroupComplianceComplianceStatusEnumValues Enumerates the set of values for PatchGroupComplianceComplianceStatusEnum
func GetPatchGroupComplianceComplianceStatusEnumValues() []PatchGroupComplianceComplianceStatusEnum {
	values := make([]PatchGroupComplianceComplianceStatusEnum, 0)
	for _, v := range mappingPatchGroupComplianceComplianceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchGroupComplianceComplianceStatusEnumStringValues Enumerates the set of values in String for PatchGroupComplianceComplianceStatusEnum
func GetPatchGroupComplianceComplianceStatusEnumStringValues() []string {
	return []string{
		"COMPLIANT",
		"NON_COMPLIANT",
	}
}

// GetMappingPatchGroupComplianceComplianceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchGroupComplianceComplianceStatusEnum(val string) (PatchGroupComplianceComplianceStatusEnum, bool) {
	enum, ok := mappingPatchGroupComplianceComplianceStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
