// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AdditionalConfigPropertyDefinition Detector rule additional property field
type AdditionalConfigPropertyDefinition struct {

	// Property Type
	PropertyType AdditionalConfigPropertyDefinitionPropertyTypeEnum `mandatory:"false" json:"propertyType,omitempty"`

	// Name for Additional Property, for example, "interpreter", "router"
	Key *string `mandatory:"false" json:"key"`

	// Value for Property Name, for example, "generic", "cloudguard"
	Value *string `mandatory:"false" json:"value"`
}

func (m AdditionalConfigPropertyDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdditionalConfigPropertyDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAdditionalConfigPropertyDefinitionPropertyTypeEnum(string(m.PropertyType)); !ok && m.PropertyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PropertyType: %s. Supported values are: %s.", m.PropertyType, strings.Join(GetAdditionalConfigPropertyDefinitionPropertyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AdditionalConfigPropertyDefinitionPropertyTypeEnum Enum with underlying type: string
type AdditionalConfigPropertyDefinitionPropertyTypeEnum string

// Set of constants representing the allowable values for AdditionalConfigPropertyDefinitionPropertyTypeEnum
const (
	AdditionalConfigPropertyDefinitionPropertyTypeHint  AdditionalConfigPropertyDefinitionPropertyTypeEnum = "HINT"
	AdditionalConfigPropertyDefinitionPropertyTypeRange AdditionalConfigPropertyDefinitionPropertyTypeEnum = "RANGE"
)

var mappingAdditionalConfigPropertyDefinitionPropertyTypeEnum = map[string]AdditionalConfigPropertyDefinitionPropertyTypeEnum{
	"HINT":  AdditionalConfigPropertyDefinitionPropertyTypeHint,
	"RANGE": AdditionalConfigPropertyDefinitionPropertyTypeRange,
}

var mappingAdditionalConfigPropertyDefinitionPropertyTypeEnumLowerCase = map[string]AdditionalConfigPropertyDefinitionPropertyTypeEnum{
	"hint":  AdditionalConfigPropertyDefinitionPropertyTypeHint,
	"range": AdditionalConfigPropertyDefinitionPropertyTypeRange,
}

// GetAdditionalConfigPropertyDefinitionPropertyTypeEnumValues Enumerates the set of values for AdditionalConfigPropertyDefinitionPropertyTypeEnum
func GetAdditionalConfigPropertyDefinitionPropertyTypeEnumValues() []AdditionalConfigPropertyDefinitionPropertyTypeEnum {
	values := make([]AdditionalConfigPropertyDefinitionPropertyTypeEnum, 0)
	for _, v := range mappingAdditionalConfigPropertyDefinitionPropertyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAdditionalConfigPropertyDefinitionPropertyTypeEnumStringValues Enumerates the set of values in String for AdditionalConfigPropertyDefinitionPropertyTypeEnum
func GetAdditionalConfigPropertyDefinitionPropertyTypeEnumStringValues() []string {
	return []string{
		"HINT",
		"RANGE",
	}
}

// GetMappingAdditionalConfigPropertyDefinitionPropertyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdditionalConfigPropertyDefinitionPropertyTypeEnum(val string) (AdditionalConfigPropertyDefinitionPropertyTypeEnum, bool) {
	enum, ok := mappingAdditionalConfigPropertyDefinitionPropertyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
