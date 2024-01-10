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

// ProductLicenseSummary The product license summary.
type ProductLicenseSummary struct {

	// The product license OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// The compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) where the product license is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current product license status.
	Status StatusEnum `mandatory:"true" json:"status"`

	// The product license unit.
	LicenseUnit LicenseUnitEnum `mandatory:"true" json:"licenseUnit"`

	// Specifies whether the vendor is Oracle or a third party.
	IsVendorOracle *bool `mandatory:"true" json:"isVendorOracle"`

	// License record name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Status description for the current product license status.
	StatusDescription *string `mandatory:"false" json:"statusDescription"`

	// The current product license state.
	LifecycleState LifeCycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The total number of licenses available for the product license, calculated by adding up all the license counts for active license records associated with the product license.
	TotalActiveLicenseUnitCount *int `mandatory:"false" json:"totalActiveLicenseUnitCount"`

	// The number of license units consumed. Updated after each allocation run.
	TotalLicenseUnitsConsumed *float64 `mandatory:"false" json:"totalLicenseUnitsConsumed"`

	// The number of license records associated with the product license.
	TotalLicenseRecordCount *int `mandatory:"false" json:"totalLicenseRecordCount"`

	// The number of active license records associated with the product license.
	ActiveLicenseRecordCount *int `mandatory:"false" json:"activeLicenseRecordCount"`

	// Specifies whether or not the product license is oversubscribed.
	IsOverSubscribed *bool `mandatory:"false" json:"isOverSubscribed"`

	// Specifies if the license unit count is unlimited.
	IsUnlimited *bool `mandatory:"false" json:"isUnlimited"`

	// The vendor of the ProductLicense
	VendorName *string `mandatory:"false" json:"vendorName"`

	// The time the product license was created. An RFC 3339 (https://tools.ietf.org/html/rfc3339)-formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the product license was updated. An RFC 3339 (https://tools.ietf.org/html/rfc3339)-formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The images associated with the product license.
	Images []ImageResponse `mandatory:"false" json:"images"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ProductLicenseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProductLicenseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLicenseUnitEnum(string(m.LicenseUnit)); !ok && m.LicenseUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseUnit: %s. Supported values are: %s.", m.LicenseUnit, strings.Join(GetLicenseUnitEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLifeCycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifeCycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
