// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

	// The capacity model to use for the Analytics Instance.
	CapacityType ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum `mandatory:"false" json:"capacityType,omitempty"`

	// The capacity value selected, either the number of OCPUs (OLPU_COUNT) or the number of users (USER_COUNT). This parameter affects the number of OCPUs, amount of memory, and other resources allocated to the Analytics Instance.
	CapacityValue *int `mandatory:"false" json:"capacityValue"`
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
	if _, ok := GetMappingResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum(string(m.CapacityType)); !ok && m.CapacityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CapacityType: %s. Supported values are: %s.", m.CapacityType, strings.Join(GetResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnumStringValues(), ",")))
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

// ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum Enum with underlying type: string
type ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum string

// Set of constants representing the allowable values for ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum
const (
	ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeOlpuCount ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum = "OLPU_COUNT"
	ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeUserCount ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum = "USER_COUNT"
)

var mappingResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum = map[string]ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum{
	"OLPU_COUNT": ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeOlpuCount,
	"USER_COUNT": ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeUserCount,
}

var mappingResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnumLowerCase = map[string]ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum{
	"olpu_count": ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeOlpuCount,
	"user_count": ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeUserCount,
}

// GetResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnumValues Enumerates the set of values for ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum
func GetResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnumValues() []ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum {
	values := make([]ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum, 0)
	for _, v := range mappingResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnumStringValues Enumerates the set of values in String for ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum
func GetResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnumStringValues() []string {
	return []string{
		"OLPU_COUNT",
		"USER_COUNT",
	}
}

// GetMappingResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum(val string) (ResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnum, bool) {
	enum, ok := mappingResourceAnalyticsInstanceOacAttachmentDetailsCapacityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
