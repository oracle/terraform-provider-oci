// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CccImageListingPackage A package for image listings.
type CccImageListingPackage struct {

	// The identifier of this package.
	Id *string `mandatory:"true" json:"id"`

	// Compute Cloud@Customer package display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The identifier of the listing this package belongs to.
	CccListingId *string `mandatory:"true" json:"cccListingId"`

	// The version of the specified package.
	CccPackageVersion *string `mandatory:"true" json:"cccPackageVersion"`

	Pricing *PricingModel `mandatory:"true" json:"pricing"`

	// Description of this package.
	Description *string `mandatory:"false" json:"description"`

	// A list of agreements that apply to this version of a package.
	CccAgreementIds []string `mandatory:"false" json:"cccAgreementIds"`

	// The date and time this listing package was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// List of operating systems supported by this package.
	OperatingSystems []OperatingSystem `mandatory:"false" json:"operatingSystems"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The regions where you can deploy the listing package. (Some packages have restrictions that limit their deployment to United States regions only.)
	Regions []Region `mandatory:"false" json:"regions"`
}

// GetId returns Id
func (m CccImageListingPackage) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m CccImageListingPackage) GetDisplayName() *string {
	return m.DisplayName
}

// GetCccListingId returns CccListingId
func (m CccImageListingPackage) GetCccListingId() *string {
	return m.CccListingId
}

// GetCccPackageVersion returns CccPackageVersion
func (m CccImageListingPackage) GetCccPackageVersion() *string {
	return m.CccPackageVersion
}

// GetDescription returns Description
func (m CccImageListingPackage) GetDescription() *string {
	return m.Description
}

// GetPricing returns Pricing
func (m CccImageListingPackage) GetPricing() *PricingModel {
	return m.Pricing
}

// GetCccAgreementIds returns CccAgreementIds
func (m CccImageListingPackage) GetCccAgreementIds() []string {
	return m.CccAgreementIds
}

// GetTimeCreated returns TimeCreated
func (m CccImageListingPackage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetOperatingSystems returns OperatingSystems
func (m CccImageListingPackage) GetOperatingSystems() []OperatingSystem {
	return m.OperatingSystems
}

// GetSystemTags returns SystemTags
func (m CccImageListingPackage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m CccImageListingPackage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CccImageListingPackage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CccImageListingPackage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCccImageListingPackage CccImageListingPackage
	s := struct {
		DiscriminatorParam string `json:"packageType"`
		MarshalTypeCccImageListingPackage
	}{
		"IMAGE",
		(MarshalTypeCccImageListingPackage)(m),
	}

	return json.Marshal(&s)
}
