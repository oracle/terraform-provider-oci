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

// UpdateLicenseRecordDetails The details about updates in the license record.
type UpdateLicenseRecordDetails struct {

	// License record name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Specifies if the license record term is perpertual.
	IsPerpetual *bool `mandatory:"true" json:"isPerpetual"`

	// Specifies if the license count is unlimited.
	IsUnlimited *bool `mandatory:"true" json:"isUnlimited"`

	// The license record end date in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// date format.
	// Example: `2018-09-12`
	ExpirationDate *common.SDKTime `mandatory:"false" json:"expirationDate"`

	// The license record support end date in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// date format.
	// Example: `2018-09-12`
	SupportEndDate *common.SDKTime `mandatory:"false" json:"supportEndDate"`

	// The number of license units added by a user in a license record.
	// Default 1
	LicenseCount *int `mandatory:"false" json:"licenseCount"`

	// The license record product ID.
	ProductId *string `mandatory:"false" json:"productId"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateLicenseRecordDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateLicenseRecordDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
