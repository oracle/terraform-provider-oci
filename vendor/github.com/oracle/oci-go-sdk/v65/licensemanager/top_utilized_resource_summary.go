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

// TopUtilizedResourceSummary A summary of a top utlized resource.
type TopUtilizedResourceSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Resource canonical name.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that contains the resource.
	ResourceCompartmentId *string `mandatory:"true" json:"resourceCompartmentId"`

	// The display name of the compartment that contains the resource.
	ResourceCompartmentName *string `mandatory:"true" json:"resourceCompartmentName"`

	// Number of license units consumed by the resource.
	TotalUnits *float64 `mandatory:"true" json:"totalUnits"`

	// The resource unit.
	UnitType ResourceUnitEnum `mandatory:"true" json:"unitType"`
}

func (m TopUtilizedResourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TopUtilizedResourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourceUnitEnum(string(m.UnitType)); !ok && m.UnitType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UnitType: %s. Supported values are: %s.", m.UnitType, strings.Join(GetResourceUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
