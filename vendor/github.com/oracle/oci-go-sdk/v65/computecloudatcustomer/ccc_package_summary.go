// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CccPackageSummary A base object for all types of listing packages.
type CccPackageSummary interface {

	// The identifier of this package.
	GetId() *string

	// Compute Cloud@Customer package display name.
	GetDisplayName() *string

	// The identifier of the listing this package belongs to.
	GetCccListingId() *string

	// The version of the specified package.
	GetCccPackageVersion() *string

	GetPricing() *PricingModel

	// Description of this package.
	GetDescription() *string

	// A list of agreements that apply to this version of a package.
	GetCccAgreementIds() []string

	// The date and time this listing package was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	GetTimeCreated() *common.SDKTime

	// List of operating systems supported by this package.
	GetOperatingSystems() []OperatingSystem

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type cccpackagesummary struct {
	JsonData          []byte
	Description       *string                           `mandatory:"false" json:"description"`
	CccAgreementIds   []string                          `mandatory:"false" json:"cccAgreementIds"`
	TimeCreated       *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	OperatingSystems  []OperatingSystem                 `mandatory:"false" json:"operatingSystems"`
	SystemTags        map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                *string                           `mandatory:"true" json:"id"`
	DisplayName       *string                           `mandatory:"true" json:"displayName"`
	CccListingId      *string                           `mandatory:"true" json:"cccListingId"`
	CccPackageVersion *string                           `mandatory:"true" json:"cccPackageVersion"`
	Pricing           *PricingModel                     `mandatory:"true" json:"pricing"`
	PackageType       string                            `json:"packageType"`
}

// UnmarshalJSON unmarshals json
func (m *cccpackagesummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercccpackagesummary cccpackagesummary
	s := struct {
		Model Unmarshalercccpackagesummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CccListingId = s.Model.CccListingId
	m.CccPackageVersion = s.Model.CccPackageVersion
	m.Pricing = s.Model.Pricing
	m.Description = s.Model.Description
	m.CccAgreementIds = s.Model.CccAgreementIds
	m.TimeCreated = s.Model.TimeCreated
	m.OperatingSystems = s.Model.OperatingSystems
	m.SystemTags = s.Model.SystemTags
	m.PackageType = s.Model.PackageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *cccpackagesummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PackageType {
	case "IMAGE":
		mm := CccImageListingPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CccPackageSummary: %s.", m.PackageType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m cccpackagesummary) GetDescription() *string {
	return m.Description
}

// GetCccAgreementIds returns CccAgreementIds
func (m cccpackagesummary) GetCccAgreementIds() []string {
	return m.CccAgreementIds
}

// GetTimeCreated returns TimeCreated
func (m cccpackagesummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetOperatingSystems returns OperatingSystems
func (m cccpackagesummary) GetOperatingSystems() []OperatingSystem {
	return m.OperatingSystems
}

// GetSystemTags returns SystemTags
func (m cccpackagesummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m cccpackagesummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m cccpackagesummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCccListingId returns CccListingId
func (m cccpackagesummary) GetCccListingId() *string {
	return m.CccListingId
}

// GetCccPackageVersion returns CccPackageVersion
func (m cccpackagesummary) GetCccPackageVersion() *string {
	return m.CccPackageVersion
}

// GetPricing returns Pricing
func (m cccpackagesummary) GetPricing() *PricingModel {
	return m.Pricing
}

func (m cccpackagesummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m cccpackagesummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
