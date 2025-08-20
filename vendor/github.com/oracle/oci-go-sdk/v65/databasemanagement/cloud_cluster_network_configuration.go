// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudClusterNetworkConfiguration The details of a network address configuration in a cloud cluster.
type CloudClusterNetworkConfiguration struct {

	// The network number.
	NetworkNumber *int `mandatory:"false" json:"networkNumber"`

	// The network type.
	NetworkType CloudClusterNetworkConfigurationNetworkTypeEnum `mandatory:"false" json:"networkType,omitempty"`

	// The subnet for the network.
	Subnet *string `mandatory:"false" json:"subnet"`
}

func (m CloudClusterNetworkConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudClusterNetworkConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCloudClusterNetworkConfigurationNetworkTypeEnum(string(m.NetworkType)); !ok && m.NetworkType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkType: %s. Supported values are: %s.", m.NetworkType, strings.Join(GetCloudClusterNetworkConfigurationNetworkTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudClusterNetworkConfigurationNetworkTypeEnum Enum with underlying type: string
type CloudClusterNetworkConfigurationNetworkTypeEnum string

// Set of constants representing the allowable values for CloudClusterNetworkConfigurationNetworkTypeEnum
const (
	CloudClusterNetworkConfigurationNetworkTypeAutoconfig CloudClusterNetworkConfigurationNetworkTypeEnum = "AUTOCONFIG"
	CloudClusterNetworkConfigurationNetworkTypeDhcp       CloudClusterNetworkConfigurationNetworkTypeEnum = "DHCP"
	CloudClusterNetworkConfigurationNetworkTypeStatic     CloudClusterNetworkConfigurationNetworkTypeEnum = "STATIC"
	CloudClusterNetworkConfigurationNetworkTypeMixed      CloudClusterNetworkConfigurationNetworkTypeEnum = "MIXED"
)

var mappingCloudClusterNetworkConfigurationNetworkTypeEnum = map[string]CloudClusterNetworkConfigurationNetworkTypeEnum{
	"AUTOCONFIG": CloudClusterNetworkConfigurationNetworkTypeAutoconfig,
	"DHCP":       CloudClusterNetworkConfigurationNetworkTypeDhcp,
	"STATIC":     CloudClusterNetworkConfigurationNetworkTypeStatic,
	"MIXED":      CloudClusterNetworkConfigurationNetworkTypeMixed,
}

var mappingCloudClusterNetworkConfigurationNetworkTypeEnumLowerCase = map[string]CloudClusterNetworkConfigurationNetworkTypeEnum{
	"autoconfig": CloudClusterNetworkConfigurationNetworkTypeAutoconfig,
	"dhcp":       CloudClusterNetworkConfigurationNetworkTypeDhcp,
	"static":     CloudClusterNetworkConfigurationNetworkTypeStatic,
	"mixed":      CloudClusterNetworkConfigurationNetworkTypeMixed,
}

// GetCloudClusterNetworkConfigurationNetworkTypeEnumValues Enumerates the set of values for CloudClusterNetworkConfigurationNetworkTypeEnum
func GetCloudClusterNetworkConfigurationNetworkTypeEnumValues() []CloudClusterNetworkConfigurationNetworkTypeEnum {
	values := make([]CloudClusterNetworkConfigurationNetworkTypeEnum, 0)
	for _, v := range mappingCloudClusterNetworkConfigurationNetworkTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudClusterNetworkConfigurationNetworkTypeEnumStringValues Enumerates the set of values in String for CloudClusterNetworkConfigurationNetworkTypeEnum
func GetCloudClusterNetworkConfigurationNetworkTypeEnumStringValues() []string {
	return []string{
		"AUTOCONFIG",
		"DHCP",
		"STATIC",
		"MIXED",
	}
}

// GetMappingCloudClusterNetworkConfigurationNetworkTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudClusterNetworkConfigurationNetworkTypeEnum(val string) (CloudClusterNetworkConfigurationNetworkTypeEnum, bool) {
	enum, ok := mappingCloudClusterNetworkConfigurationNetworkTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
