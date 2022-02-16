// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Content and Experience API
//
// Oracle Content and Experience is a cloud-based content hub to drive omni-channel content management and accelerate experience delivery
//

package oce

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateOceInstanceDetails The information to be updated.
type UpdateOceInstanceDetails struct {

	// OceInstance description
	Description *string `mandatory:"false" json:"description"`

	// Web Application Firewall(WAF) primary domain
	WafPrimaryDomain *string `mandatory:"false" json:"wafPrimaryDomain"`

	// Flag indicating whether the instance license is new cloud or bring your own license
	InstanceLicenseType LicenseTypeEnum `mandatory:"false" json:"instanceLicenseType,omitempty"`

	// Instance type based on its usage
	InstanceUsageType UpdateOceInstanceDetailsInstanceUsageTypeEnum `mandatory:"false" json:"instanceUsageType,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateOceInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOceInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLicenseTypeEnum(string(m.InstanceLicenseType)); !ok && m.InstanceLicenseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstanceLicenseType: %s. Supported values are: %s.", m.InstanceLicenseType, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateOceInstanceDetailsInstanceUsageTypeEnum(string(m.InstanceUsageType)); !ok && m.InstanceUsageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstanceUsageType: %s. Supported values are: %s.", m.InstanceUsageType, strings.Join(GetUpdateOceInstanceDetailsInstanceUsageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateOceInstanceDetailsInstanceUsageTypeEnum Enum with underlying type: string
type UpdateOceInstanceDetailsInstanceUsageTypeEnum string

// Set of constants representing the allowable values for UpdateOceInstanceDetailsInstanceUsageTypeEnum
const (
	UpdateOceInstanceDetailsInstanceUsageTypePrimary    UpdateOceInstanceDetailsInstanceUsageTypeEnum = "PRIMARY"
	UpdateOceInstanceDetailsInstanceUsageTypeNonprimary UpdateOceInstanceDetailsInstanceUsageTypeEnum = "NONPRIMARY"
)

var mappingUpdateOceInstanceDetailsInstanceUsageTypeEnum = map[string]UpdateOceInstanceDetailsInstanceUsageTypeEnum{
	"PRIMARY":    UpdateOceInstanceDetailsInstanceUsageTypePrimary,
	"NONPRIMARY": UpdateOceInstanceDetailsInstanceUsageTypeNonprimary,
}

// GetUpdateOceInstanceDetailsInstanceUsageTypeEnumValues Enumerates the set of values for UpdateOceInstanceDetailsInstanceUsageTypeEnum
func GetUpdateOceInstanceDetailsInstanceUsageTypeEnumValues() []UpdateOceInstanceDetailsInstanceUsageTypeEnum {
	values := make([]UpdateOceInstanceDetailsInstanceUsageTypeEnum, 0)
	for _, v := range mappingUpdateOceInstanceDetailsInstanceUsageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateOceInstanceDetailsInstanceUsageTypeEnumStringValues Enumerates the set of values in String for UpdateOceInstanceDetailsInstanceUsageTypeEnum
func GetUpdateOceInstanceDetailsInstanceUsageTypeEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"NONPRIMARY",
	}
}

// GetMappingUpdateOceInstanceDetailsInstanceUsageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateOceInstanceDetailsInstanceUsageTypeEnum(val string) (UpdateOceInstanceDetailsInstanceUsageTypeEnum, bool) {
	mappingUpdateOceInstanceDetailsInstanceUsageTypeEnumIgnoreCase := make(map[string]UpdateOceInstanceDetailsInstanceUsageTypeEnum)
	for k, v := range mappingUpdateOceInstanceDetailsInstanceUsageTypeEnum {
		mappingUpdateOceInstanceDetailsInstanceUsageTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateOceInstanceDetailsInstanceUsageTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
