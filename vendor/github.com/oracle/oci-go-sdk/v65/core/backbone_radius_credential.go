// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BackboneRadiusCredential RadiusCredential for edgepop devices. For C3 internal use only.
type BackboneRadiusCredential struct {

	// The device type.
	DeviceType BackboneRadiusCredentialDeviceTypeEnum `mandatory:"true" json:"deviceType"`

	// The primary vip of radius server.
	PrimaryVip *string `mandatory:"true" json:"primaryVip"`

	// The primary radius credential for edgepop devices.
	PrimaryCredential *string `mandatory:"true" json:"primaryCredential"`

	// The secondary vip of radius server.
	SecondaryVip *string `mandatory:"true" json:"secondaryVip"`

	// The secondary radius credential for edgepop devices.
	SecondaryCredential *string `mandatory:"true" json:"secondaryCredential"`
}

func (m BackboneRadiusCredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackboneRadiusCredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBackboneRadiusCredentialDeviceTypeEnum(string(m.DeviceType)); !ok && m.DeviceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeviceType: %s. Supported values are: %s.", m.DeviceType, strings.Join(GetBackboneRadiusCredentialDeviceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackboneRadiusCredentialDeviceTypeEnum Enum with underlying type: string
type BackboneRadiusCredentialDeviceTypeEnum string

// Set of constants representing the allowable values for BackboneRadiusCredentialDeviceTypeEnum
const (
	BackboneRadiusCredentialDeviceTypeJuniper BackboneRadiusCredentialDeviceTypeEnum = "JUNIPER"
	BackboneRadiusCredentialDeviceTypeArista  BackboneRadiusCredentialDeviceTypeEnum = "ARISTA"
)

var mappingBackboneRadiusCredentialDeviceTypeEnum = map[string]BackboneRadiusCredentialDeviceTypeEnum{
	"JUNIPER": BackboneRadiusCredentialDeviceTypeJuniper,
	"ARISTA":  BackboneRadiusCredentialDeviceTypeArista,
}

var mappingBackboneRadiusCredentialDeviceTypeEnumLowerCase = map[string]BackboneRadiusCredentialDeviceTypeEnum{
	"juniper": BackboneRadiusCredentialDeviceTypeJuniper,
	"arista":  BackboneRadiusCredentialDeviceTypeArista,
}

// GetBackboneRadiusCredentialDeviceTypeEnumValues Enumerates the set of values for BackboneRadiusCredentialDeviceTypeEnum
func GetBackboneRadiusCredentialDeviceTypeEnumValues() []BackboneRadiusCredentialDeviceTypeEnum {
	values := make([]BackboneRadiusCredentialDeviceTypeEnum, 0)
	for _, v := range mappingBackboneRadiusCredentialDeviceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackboneRadiusCredentialDeviceTypeEnumStringValues Enumerates the set of values in String for BackboneRadiusCredentialDeviceTypeEnum
func GetBackboneRadiusCredentialDeviceTypeEnumStringValues() []string {
	return []string{
		"JUNIPER",
		"ARISTA",
	}
}

// GetMappingBackboneRadiusCredentialDeviceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackboneRadiusCredentialDeviceTypeEnum(val string) (BackboneRadiusCredentialDeviceTypeEnum, bool) {
	enum, ok := mappingBackboneRadiusCredentialDeviceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
