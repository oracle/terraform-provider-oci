// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// License Manager API
//
// Use the License Manager API to manage product licenses and license records. For more information, see License Manager Overview (https://docs.cloud.oracle.com/iaas/Content/LicenseManager/Concepts/licensemanageroverview.htm).
//

package licensemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LicenseMetric Overview of product license and resources usage.
type LicenseMetric struct {

	// Total number of product licenses in a particular compartment.
	TotalProductLicenseCount *int `mandatory:"true" json:"totalProductLicenseCount"`

	// Total number of BYOL instances in a particular compartment.
	TotalByolInstanceCount *int `mandatory:"true" json:"totalByolInstanceCount"`

	// Total number of License Included (LI) instances in a particular compartment.
	TotalLicenseIncludedInstanceCount *int `mandatory:"true" json:"totalLicenseIncludedInstanceCount"`

	// Total number of license records that will expire within 90 days in a particular compartment.
	LicenseRecordExpiringSoonCount *int `mandatory:"true" json:"licenseRecordExpiringSoonCount"`
}

func (m LicenseMetric) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LicenseMetric) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
