// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Analytics API
//
// Use the Resource Analytics API to manage Resource Analytics Instances.
//

package resourceanalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceAnalyticsInstanceOacAttachmentDetails Additional details needed when attaching the OAC instance.
// Example: `{"idcsDomainId":"ocid...","networkDetails":{...}, ...}`
type ResourceAnalyticsInstanceOacAttachmentDetails struct {

	// IDCS domain OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) identifying a stripe and service administrator user.
	IdcsDomainId *string `mandatory:"true" json:"idcsDomainId"`

	NetworkDetails *ResourceAnalyticsInstanceOacNetworkDetails `mandatory:"false" json:"networkDetails"`

	// Deprecated. Use `networkDetails.subnetId` instead.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Deprecated. Use `networkDetails.nsgIds` instead.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The Oracle license model that applies to the OAC instance.
	LicenseModel ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`
}

func (m ResourceAnalyticsInstanceOacAttachmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceAnalyticsInstanceOacAttachmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum Enum with underlying type: string
type ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum string

// Set of constants representing the allowable values for ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum
const (
	ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelLicenseIncluded     ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelBringYourOwnLicense ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum = map[string]ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelBringYourOwnLicense,
}

var mappingResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnumLowerCase = map[string]ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum{
	"license_included":       ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelBringYourOwnLicense,
}

// GetResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnumValues Enumerates the set of values for ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum
func GetResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnumValues() []ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum {
	values := make([]ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum, 0)
	for _, v := range mappingResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnumStringValues Enumerates the set of values in String for ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum
func GetResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum(val string) (ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum, bool) {
	enum, ok := mappingResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
