// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Analytics API
//
// Use the Resource Analytics API to manage Resource Analytics Instances.
//

package resourceanalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateResourceAnalyticsInstanceDetails The data to create a ResourceAnalyticsInstance.
type CreateResourceAnalyticsInstanceDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the ResourceAnalyticsInstance in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	AdwAdminPassword AdwAdminPasswordDetails `mandatory:"true" json:"adwAdminPassword"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description of the ResourceAnalyticsInstance instance.
	Description *string `mandatory:"false" json:"description"`

	// Require mutual TLS (mTLS) when authenticating connections to the ADW database.
	IsMutualTlsRequired *bool `mandatory:"false" json:"isMutualTlsRequired"`

	// List of Network Security Group OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)'s.
	// Example: `["ocid...", "ocid..."]`
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The Oracle license model that applies to the ADW instance.
	LicenseModel CreateResourceAnalyticsInstanceDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateResourceAnalyticsInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateResourceAnalyticsInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateResourceAnalyticsInstanceDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCreateResourceAnalyticsInstanceDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateResourceAnalyticsInstanceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName         *string                                                `json:"displayName"`
		Description         *string                                                `json:"description"`
		IsMutualTlsRequired *bool                                                  `json:"isMutualTlsRequired"`
		NsgIds              []string                                               `json:"nsgIds"`
		LicenseModel        CreateResourceAnalyticsInstanceDetailsLicenseModelEnum `json:"licenseModel"`
		FreeformTags        map[string]string                                      `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{}                      `json:"definedTags"`
		CompartmentId       *string                                                `json:"compartmentId"`
		AdwAdminPassword    adwadminpassworddetails                                `json:"adwAdminPassword"`
		SubnetId            *string                                                `json:"subnetId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.IsMutualTlsRequired = model.IsMutualTlsRequired

	m.NsgIds = make([]string, len(model.NsgIds))
	copy(m.NsgIds, model.NsgIds)
	m.LicenseModel = model.LicenseModel

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	nn, e = model.AdwAdminPassword.UnmarshalPolymorphicJSON(model.AdwAdminPassword.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AdwAdminPassword = nn.(AdwAdminPasswordDetails)
	} else {
		m.AdwAdminPassword = nil
	}

	m.SubnetId = model.SubnetId

	return
}

// CreateResourceAnalyticsInstanceDetailsLicenseModelEnum Enum with underlying type: string
type CreateResourceAnalyticsInstanceDetailsLicenseModelEnum string

// Set of constants representing the allowable values for CreateResourceAnalyticsInstanceDetailsLicenseModelEnum
const (
	CreateResourceAnalyticsInstanceDetailsLicenseModelLicenseIncluded     CreateResourceAnalyticsInstanceDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	CreateResourceAnalyticsInstanceDetailsLicenseModelBringYourOwnLicense CreateResourceAnalyticsInstanceDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateResourceAnalyticsInstanceDetailsLicenseModelEnum = map[string]CreateResourceAnalyticsInstanceDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateResourceAnalyticsInstanceDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateResourceAnalyticsInstanceDetailsLicenseModelBringYourOwnLicense,
}

var mappingCreateResourceAnalyticsInstanceDetailsLicenseModelEnumLowerCase = map[string]CreateResourceAnalyticsInstanceDetailsLicenseModelEnum{
	"license_included":       CreateResourceAnalyticsInstanceDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": CreateResourceAnalyticsInstanceDetailsLicenseModelBringYourOwnLicense,
}

// GetCreateResourceAnalyticsInstanceDetailsLicenseModelEnumValues Enumerates the set of values for CreateResourceAnalyticsInstanceDetailsLicenseModelEnum
func GetCreateResourceAnalyticsInstanceDetailsLicenseModelEnumValues() []CreateResourceAnalyticsInstanceDetailsLicenseModelEnum {
	values := make([]CreateResourceAnalyticsInstanceDetailsLicenseModelEnum, 0)
	for _, v := range mappingCreateResourceAnalyticsInstanceDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateResourceAnalyticsInstanceDetailsLicenseModelEnumStringValues Enumerates the set of values in String for CreateResourceAnalyticsInstanceDetailsLicenseModelEnum
func GetCreateResourceAnalyticsInstanceDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCreateResourceAnalyticsInstanceDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateResourceAnalyticsInstanceDetailsLicenseModelEnum(val string) (CreateResourceAnalyticsInstanceDetailsLicenseModelEnum, bool) {
	enum, ok := mappingCreateResourceAnalyticsInstanceDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
