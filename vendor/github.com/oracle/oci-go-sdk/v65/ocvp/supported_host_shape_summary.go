// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SupportedHostShapeSummary A specific compute shape supported by the Oracle Cloud VMware Solution.
type SupportedHostShapeSummary struct {

	// The name of the supported compute shape.
	Name *string `mandatory:"true" json:"name"`

	// The operations where you can use the shape. The operations can be CREATE_SDDC or CREATE_ESXI_HOST.
	SupportedOperations []OperationTypesEnum `mandatory:"true" json:"supportedOperations"`

	// The family of the shape. ESXi hosts of one SDDC must have the same shape family.
	ShapeFamily *string `mandatory:"true" json:"shapeFamily"`

	// The default OCPU count of the shape.
	DefaultOcpuCount *float32 `mandatory:"false" json:"defaultOcpuCount"`

	// Support OCPU count of the shape.
	SupportedOcpuCount []float32 `mandatory:"false" json:"supportedOcpuCount"`

	// Indicates whether the shape supports single host SDDCs.
	IsSingleHostSddcSupported *bool `mandatory:"false" json:"isSingleHostSddcSupported"`

	// The VMware software versions supported by the shape.
	SupportedVmwareSoftwareVersions []string `mandatory:"false" json:"supportedVmwareSoftwareVersions"`

	// Description of the shape.
	Description *string `mandatory:"false" json:"description"`

	// Indicates whether the shape supports shielded instances.
	IsSupportShieldedInstances *bool `mandatory:"false" json:"isSupportShieldedInstances"`

	// Whether the shape supports "MONTH" Commitment.
	IsSupportMonthlyCommitment *bool `mandatory:"false" json:"isSupportMonthlyCommitment"`
}

func (m SupportedHostShapeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SupportedHostShapeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.SupportedOperations {
		if _, ok := GetMappingOperationTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SupportedOperations: %s. Supported values are: %s.", val, strings.Join(GetOperationTypesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
