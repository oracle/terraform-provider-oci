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

// ProductLicenseConsumerSummary Details of a resource that is consuming a particular product license.
type ProductLicenseConsumerSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The display name of the resource.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The resource product name.
	ProductName *string `mandatory:"true" json:"productName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the resource.
	ResourceCompartmentId *string `mandatory:"true" json:"resourceCompartmentId"`

	// The display name of the compartment that contains the resource.
	ResourceCompartmentName *string `mandatory:"true" json:"resourceCompartmentName"`

	// The unit type for the resource.
	ResourceUnitType ResourceUnitEnum `mandatory:"true" json:"resourceUnitType"`

	// Number of units of the resource
	ResourceUnitCount *float64 `mandatory:"true" json:"resourceUnitCount"`

	// The product license unit.
	LicenseUnitType LicenseUnitEnum `mandatory:"true" json:"licenseUnitType"`

	// Number of license units consumed by the resource.
	LicenseUnitsConsumed *float64 `mandatory:"true" json:"licenseUnitsConsumed"`

	// Specifies if the base license is available.
	IsBaseLicenseAvailable *bool `mandatory:"true" json:"isBaseLicenseAvailable"`

	// Specifies if all options are available.
	AreAllOptionsAvailable *bool `mandatory:"true" json:"areAllOptionsAvailable"`

	// Collection of missing product licenses.
	MissingProducts []Product `mandatory:"true" json:"missingProducts"`
}

func (m ProductLicenseConsumerSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProductLicenseConsumerSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourceUnitEnum(string(m.ResourceUnitType)); !ok && m.ResourceUnitType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceUnitType: %s. Supported values are: %s.", m.ResourceUnitType, strings.Join(GetResourceUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLicenseUnitEnum(string(m.LicenseUnitType)); !ok && m.LicenseUnitType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseUnitType: %s. Supported values are: %s.", m.LicenseUnitType, strings.Join(GetLicenseUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
