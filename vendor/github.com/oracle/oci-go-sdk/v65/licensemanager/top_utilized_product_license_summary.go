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

// TopUtilizedProductLicenseSummary A summary of the top utilized product licenses.
type TopUtilizedProductLicenseSummary struct {

	// The product license OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ProductLicenseId *string `mandatory:"true" json:"productLicenseId"`

	// The product type.
	ProductType *string `mandatory:"true" json:"productType"`

	// The product license unit.
	UnitType LicenseUnitEnum `mandatory:"true" json:"unitType"`

	// Number of license units consumed.
	TotalUnitsConsumed *float64 `mandatory:"true" json:"totalUnitsConsumed"`

	// Total number of license units in the product license provided by the user.
	TotalLicenseUnitCount *int `mandatory:"true" json:"totalLicenseUnitCount"`

	// Specifies if the license unit count is unlimited.
	IsUnlimited *bool `mandatory:"true" json:"isUnlimited"`

	// The current product license status.
	Status StatusEnum `mandatory:"true" json:"status"`
}

func (m TopUtilizedProductLicenseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TopUtilizedProductLicenseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLicenseUnitEnum(string(m.UnitType)); !ok && m.UnitType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UnitType: %s. Supported values are: %s.", m.UnitType, strings.Join(GetLicenseUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
