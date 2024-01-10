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

// LicenseRecordSummary The license record summary.
type LicenseRecordSummary struct {

	// The license record OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// License record display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Specifies if the license count is unlimited.
	IsUnlimited *bool `mandatory:"true" json:"isUnlimited"`

	// Specifies if the license record term is perpertual.
	IsPerpetual *bool `mandatory:"true" json:"isPerpetual"`

	// The product license OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) with which the license record is associated.
	ProductLicenseId *string `mandatory:"false" json:"productLicenseId"`

	// The compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) where the license record is created.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The license record product ID.
	ProductId *string `mandatory:"false" json:"productId"`

	// The number of license record units added by the user for the given license record.
	// Default 1
	LicenseCount *int `mandatory:"false" json:"licenseCount"`

	// The license record end date in RFC 3339 (https://tools.ietf.org/html/rfc3339) format.
	// date format.
	// Example: `2018-09-12`
	ExpirationDate *common.SDKTime `mandatory:"false" json:"expirationDate"`

	// The license record support end date in RFC 3339 (https://tools.ietf.org/html/rfc3339) format.
	// date format.
	// Example: `2018-09-12`
	SupportEndDate *common.SDKTime `mandatory:"false" json:"supportEndDate"`

	// The time the license record was created. An RFC 3339 (https://tools.ietf.org/html/rfc3339)-formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the license record was updated. An RFC 3339 (https://tools.ietf.org/html/rfc3339)-formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current license record state.
	LifecycleState LifeCycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The product license unit.
	LicenseUnit LicenseUnitEnum `mandatory:"false" json:"licenseUnit,omitempty"`

	// The product license name with which the license record is associated.
	ProductLicense *string `mandatory:"false" json:"productLicense"`

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

func (m LicenseRecordSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LicenseRecordSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifeCycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifeCycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLicenseUnitEnum(string(m.LicenseUnit)); !ok && m.LicenseUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseUnit: %s. Supported values are: %s.", m.LicenseUnit, strings.Join(GetLicenseUnitEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
